/*
 * Copyright (C) 2018 Nalej - All Rights Reserved
 */

package network

import (
	"context"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-network-go"
)

// Manager structure with the required clients for network operations.
type Manager struct {
	NetworkClient grpc_network_go.NetworksClient
	DNSClient     grpc_network_go.DNSClient
}

// NewManager creates a Manager using a set of clients.
func NewManager(networkClient grpc_network_go.NetworksClient, dnsClient grpc_network_go.DNSClient) Manager {
	return Manager{NetworkClient: networkClient, DNSClient: dnsClient}
}

// Authorize member
func (m *Manager) AuthorizeMember(request *grpc_network_go.AuthorizeMemberRequest) (*grpc_common_go.Success, error) {
	return m.NetworkClient.AuthorizeMember(context.Background(), request)
}

// AddDNSEntry creates a new DNSEntry on the system.
func (m *Manager) AddDNSEntry(addRequest *grpc_network_go.AddDNSEntryRequest) (*grpc_common_go.Success, error) {
	return m.DNSClient.AddDNSEntry(context.Background(), addRequest)
}
