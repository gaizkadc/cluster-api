/*
 * Copyright (C) 2018 Nalej - All Rights Reserved
 */

package server

import (
	"fmt"
	"github.com/nalej/authx/pkg/interceptor"
	"github.com/nalej/cluster-api/internal/pkg/server/conductor"
	"github.com/nalej/cluster-api/internal/pkg/server/device_latency"
	"github.com/nalej/cluster-api/internal/pkg/server/network"
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-cluster-api-go"
	"github.com/nalej/grpc-conductor-go"
	"github.com/nalej/grpc-device-manager-go"
	"github.com/nalej/grpc-network-go"
	"github.com/nalej/grpc-authx-go"
	"github.com/nalej/grpc-utils/pkg/tools"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

// Service structure with the configuration and the gRPC server.
type Service struct {
	Configuration Config
	Server        *tools.GenericGRPCServer
}

// NewService creates a new system model service.
func NewService(conf Config) *Service {
	return &Service{
		conf,
		tools.NewGenericGRPCServer(uint32(conf.Port)),
	}
}

// Clients structure with the gRPC clients for remote services.
type Clients struct {
	NetworkManager grpc_network_go.NetworksClient
	DNSClient      grpc_network_go.DNSClient
	Conductor      grpc_conductor_go.ConductorMonitorClient
	DeviceLatency  grpc_device_manager_go.LatencyClient
	Authx grpc_authx_go.AuthxClient
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
	nClient := grpc_network_go.NewNetworksClient(nmConn)
	dnsClient := grpc_network_go.NewDNSClient(nmConn)
	cClient := grpc_conductor_go.NewConductorMonitorClient(cConn)
	dClient := grpc_device_manager_go.NewLatencyClient(dConn)
	aClient := grpc_authx_go.NewAuthxClient(aConn)
	return &Clients{nClient, dnsClient, cClient, dClient, aClient}, nil
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

	networkManager := network.NewManager(clients.NetworkManager, clients.DNSClient)
	networkHandler := network.NewHandler(networkManager)

	deviceLatencyManager := device_latency.NewManager(clients.DeviceLatency, clients.Authx)
	deviceLatencyHandler := device_latency.NewHandler(deviceLatencyManager)

	// Create handlers
	grpcServer := grpc.NewServer(interceptor.WithServerAuthxInterceptor(
		interceptor.NewConfig(authConfig, s.Configuration.AuthSecret, s.Configuration.AuthHeader)))

	grpc_cluster_api_go.RegisterConductorServer(grpcServer, conductorHandler)
	grpc_cluster_api_go.RegisterNetworkManagerServer(grpcServer, networkHandler)
	grpc_cluster_api_go.RegisterDeviceManagerServer(grpcServer, deviceLatencyHandler)

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)
	log.Info().Int("port", s.Configuration.Port).Msg("Launching gRPC server")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal().Errs("failed to serve: %v", []error{err})
	}
	return nil
}
