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

package network

import (
	"context"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-network-go"
	"github.com/nalej/grpc-utils/pkg/conversions"
	"github.com/nalej/nalej-bus/pkg/queue/network/ops"
	"github.com/rs/zerolog/log"
	"time"
)

// Timeout to send packages to the network operations queue
const NetworkOpsTimeout = time.Second * 5

// Manager structure with the required clients for network operations.
type Manager struct {
	NetworkClient      grpc_network_go.NetworksClient
	DNSClient          grpc_network_go.DNSClient
	NetworkOpsProducer *ops.NetworkOpsProducer
}

// NewManager creates a Manager using a set of clients.
func NewManager(networkClient grpc_network_go.NetworksClient, dnsClient grpc_network_go.DNSClient, netProducer *ops.NetworkOpsProducer) Manager {
	return Manager{NetworkClient: networkClient, DNSClient: dnsClient, NetworkOpsProducer: netProducer}
}

// Authorize member
func (m *Manager) AuthorizeMember(request *grpc_network_go.AuthorizeMemberRequest) (*grpc_common_go.Success, error) {
	// send an asynchronous message and return success if no error is detected
	ctx, cancel := context.WithTimeout(context.Background(), NetworkOpsTimeout)
	defer cancel()
	err := m.NetworkOpsProducer.Send(ctx, request)
	if err != nil {
		return nil, err
	}
	return &grpc_common_go.Success{}, nil
}

// AddDNSEntry creates a new DNSEntry on the system.
func (m *Manager) AddDNSEntry(request *grpc_network_go.AddDNSEntryRequest) (*grpc_common_go.Success, error) {
	// send an asynchronous message and return success if no error is detected
	ctx, cancel := context.WithTimeout(context.Background(), NetworkOpsTimeout)
	defer cancel()
	err := m.NetworkOpsProducer.Send(ctx, request)
	if err != nil {
		return nil, err
	}
	return &grpc_common_go.Success{}, nil
}

func (m *Manager) RegisterInboundServiceProxy(request *grpc_network_go.InboundServiceProxy) (*grpc_common_go.Success, error) {
	ctx, cancel := context.WithTimeout(context.Background(), NetworkOpsTimeout)
	defer cancel()
	err := m.NetworkOpsProducer.Send(ctx, request)
	if err != nil {
		return nil, err
	}
	return &grpc_common_go.Success{}, nil
}

func (m *Manager) RegisterOutboundProxy(request *grpc_network_go.OutboundService) (*grpc_common_go.Success, error) {
	ctx, cancel := context.WithTimeout(context.Background(), NetworkOpsTimeout)
	defer cancel()
	err := m.NetworkOpsProducer.Send(ctx, request)
	if err != nil {
		return nil, err
	}
	return &grpc_common_go.Success{}, nil
}

func (m *Manager) AuthorizeZTConnection(request *grpc_network_go.AuthorizeZTConnectionRequest) (*grpc_common_go.Success, error) {
	ctx, cancel := context.WithTimeout(context.Background(), NetworkOpsTimeout)
	defer cancel()
	err := m.NetworkOpsProducer.Send(ctx, request)
	if err != nil {
		log.Debug().Str("trace", conversions.ToDerror(err).DebugReport()).Msg("ERROR AuthorizeZTConnection")
		return nil, err
	}
	return &grpc_common_go.Success{}, nil
}

func (m *Manager) RegisterZTConnection(request *grpc_network_go.RegisterZTConnectionRequest) (*grpc_common_go.Success, error) {
	ctx, cancel := context.WithTimeout(context.Background(), NetworkOpsTimeout)
	defer cancel()
	err := m.NetworkOpsProducer.Send(ctx, request)
	if err != nil {
		return nil, err
	}
	return &grpc_common_go.Success{}, nil
}
