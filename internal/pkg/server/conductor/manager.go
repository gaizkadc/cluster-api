/*
 * Copyright (C) 2018 Nalej - All Rights Reserved
 */

package conductor

import (
	"context"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-conductor-go"
)

// Manager structure with the required clients for conductor operations.
type Manager struct {
	ConductorClient grpc_conductor_go.ConductorMonitorClient
}

// NewManager creates a Manager using a set of clients.
func NewManager(conductorClient grpc_conductor_go.ConductorMonitorClient) Manager {
	return Manager{ConductorClient: conductorClient}
}

// Request to update a fragment plan status with all the corresponding services
func (m *Manager) UpdateDeploymentFragmentStatus(updateRequest *grpc_conductor_go.DeploymentFragmentUpdateRequest) (*grpc_common_go.Success, error) {
	return m.ConductorClient.UpdateDeploymentFragmentStatus(context.Background(), updateRequest)
}

// Update the status of a set of services during a given fragment deployment
func (m *Manager) UpdateServiceStatus(updateRequest *grpc_conductor_go.DeploymentServiceUpdateRequest) (*grpc_common_go.Success, error) {
	return m.ConductorClient.UpdateServiceStatus(context.Background(), updateRequest)
}
