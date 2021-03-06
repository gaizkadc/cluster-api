/*
 * Copyright 2019 Nalej
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package server

import (
	"fmt"
	"github.com/nalej/authx/pkg/interceptor"
	"github.com/nalej/cluster-api/internal/pkg/server/cluster-watcher"
	"github.com/nalej/cluster-api/internal/pkg/server/conductor"
	"github.com/nalej/cluster-api/internal/pkg/server/connectivity-checker"
	"github.com/nalej/cluster-api/internal/pkg/server/device_latency"
	"github.com/nalej/cluster-api/internal/pkg/server/network"
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-authx-go"
	"github.com/nalej/grpc-cluster-api-go"
	"github.com/nalej/grpc-cluster-watcher-go"
	"github.com/nalej/grpc-conductor-go"
	"github.com/nalej/grpc-device-manager-go"
	"github.com/nalej/grpc-network-go"
	"github.com/nalej/nalej-bus/pkg/bus"
	"github.com/nalej/nalej-bus/pkg/bus/pulsar-comcast"
	"github.com/nalej/nalej-bus/pkg/queue/infrastructure/events"
	"github.com/nalej/nalej-bus/pkg/queue/network/ops"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

const (
	ClusterAPIInfraEventsProducerName = "cluster-api-infrastructure-events"
	ClusterAPINetworkOpsProducerName  = "cluster-api-network-ops"
)

// Service structure with the configuration and the gRPC server.
type Service struct {
	Configuration Config
}

// NewService creates a new system model service.
func NewService(conf Config) *Service {
	return &Service{
		conf,
	}
}

// Clients structure with the gRPC clients for remote services.
type Clients struct {
	NetworkManager grpc_network_go.NetworksClient
	DNSClient      grpc_network_go.DNSClient
	Conductor      grpc_conductor_go.ConductorMonitorClient
	DeviceLatency  grpc_device_manager_go.LatencyClient
	Authx          grpc_authx_go.AuthxClient
	QueueClient    bus.NalejClient
	ClusterWatcher grpc_cluster_watcher_go.ClusterWatcherMasterClient
}

type BusClients struct {
	NetworkOpsProducer           *ops.NetworkOpsProducer
	InfrastructureEventsProducer *events.InfrastructureEventsProducer
}

// GetBusClients creates the required connections with the bus
func (s *Service) GetBusClients() (*BusClients, derrors.Error) {
	queueClient := pulsar_comcast.NewClient(s.Configuration.QueueAddress, nil)

	netOpsProducer, err := ops.NewNetworkOpsProducer(queueClient, ClusterAPINetworkOpsProducerName)
	if err != nil {
		return nil, err
	}

	infraEvProducer, err := events.NewInfrastructureEventsProducer(queueClient, ClusterAPIInfraEventsProducerName)
	if err != nil {
		return nil, err
	}

	return &BusClients{
		NetworkOpsProducer:           netOpsProducer,
		InfrastructureEventsProducer: infraEvProducer,
	}, nil
}

// GetClients creates the required connections with the remote clients.
func (s *Service) GetClients() (*Clients, derrors.Error) {
	nmConn, err := grpc.Dial(s.Configuration.NetworkManagerAddress, grpc.WithInsecure())
	if err != nil {
		return nil, derrors.AsError(err, "cannot create connection with the network manager")
	}
	cConn, err := grpc.Dial(s.Configuration.ConductorAddress, grpc.WithInsecure())
	if err != nil {
		return nil, derrors.AsError(err, "cannot create connection with conductor")
	}
	dConn, err := grpc.Dial(s.Configuration.DeviceManagerAddress, grpc.WithInsecure())
	if err != nil {
		return nil, derrors.AsError(err, "cannot create connection with device manager")
	}
	aConn, err := grpc.Dial(s.Configuration.AuthxAddress, grpc.WithInsecure())
	if err != nil {
		return nil, derrors.AsError(err, "cannot create connection with authx")
	}

	cwConn, err := grpc.Dial(s.Configuration.ClusterWatcherAddress, grpc.WithInsecure())
	if err != nil {
		return nil, derrors.AsError(err, "cannot create connection with cluster watcher")
	}

	qClient := pulsar_comcast.NewClient(s.Configuration.QueueAddress, nil)

	nClient := grpc_network_go.NewNetworksClient(nmConn)
	dnsClient := grpc_network_go.NewDNSClient(nmConn)
	cClient := grpc_conductor_go.NewConductorMonitorClient(cConn)
	dClient := grpc_device_manager_go.NewLatencyClient(dConn)
	aClient := grpc_authx_go.NewAuthxClient(aConn)
	cwClient := grpc_cluster_watcher_go.NewClusterWatcherMasterClient(cwConn)

	return &Clients{
		NetworkManager: nClient,
		DNSClient:      dnsClient,
		Conductor:      cClient,
		DeviceLatency:  dClient,
		Authx:          aClient,
		QueueClient:    qClient,
		ClusterWatcher: cwClient,
	}, nil
}

// Run the service, launch the REST service handler.
func (s *Service) Run() error {
	cErr := s.Configuration.Validate()
	if cErr != nil {
		log.Fatal().Str("err", cErr.DebugReport()).Msg("invalid configuration")
	}
	s.Configuration.Print()

	authConfig, authErr := s.Configuration.LoadAuthConfig()
	if authErr != nil {
		log.Fatal().Str("err", authErr.DebugReport()).Msg("cannot load authx config")
	}

	clients, cErr := s.GetClients()
	if cErr != nil {
		log.Fatal().Str("err", cErr.DebugReport()).Msg("Cannot create clients")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Configuration.Port))
	if err != nil {
		log.Fatal().Errs("failed to listen: %v", []error{err})
	}

	conductorManager := conductor.NewManager(clients.Conductor)
	conductorHandler := conductor.NewHandler(conductorManager)

	// BusClients
	busClients, bErr := s.GetBusClients()
	if err != nil {
		log.Fatal().Str("err", bErr.DebugReport()).Msg("Cannot create bus clients")
	}

	networkManager := network.NewManager(clients.NetworkManager, clients.DNSClient, busClients.NetworkOpsProducer)
	networkHandler := network.NewHandler(networkManager)

	deviceLatencyManager := device_latency.NewManager(clients.DeviceLatency, clients.Authx)
	deviceLatencyHandler := device_latency.NewHandler(deviceLatencyManager)

	clusterWatcherManager := cluster_watcher.NewManager(clients.ClusterWatcher)
	clusterWatcherHandler := cluster_watcher.NewHandler(clusterWatcherManager)
	connectivityCheckerManager := connectivity_checker.NewManager(busClients.InfrastructureEventsProducer)
	connectivityCheckerHandler := connectivity_checker.NewHandler(connectivityCheckerManager)

	// Create handlers
	grpcServer := grpc.NewServer(interceptor.WithServerAuthxInterceptor(
		interceptor.NewConfig(authConfig, s.Configuration.AuthSecret, s.Configuration.AuthHeader)))

	grpc_cluster_api_go.RegisterConductorServer(grpcServer, conductorHandler)
	grpc_cluster_api_go.RegisterNetworkManagerServer(grpcServer, networkHandler)
	grpc_cluster_api_go.RegisterDeviceManagerServer(grpcServer, deviceLatencyHandler)
	grpc_cluster_api_go.RegisterClusterWatcherMasterServer(grpcServer, clusterWatcherHandler)
	grpc_cluster_api_go.RegisterConnectivityCheckerServer(grpcServer, connectivityCheckerHandler)

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)
	log.Info().Int("port", s.Configuration.Port).Msg("Launching gRPC server")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal().Errs("failed to serve: %v", []error{err})
	}
	return nil
}
