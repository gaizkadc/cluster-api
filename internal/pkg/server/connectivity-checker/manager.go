package connectivity_checker

import (
	"context"
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-cluster-api-go"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-infrastructure-go"
)

// Manager structure with the required clients for connectivity-checker operations.
type Manager struct {
	ConnectivityCheckerClient grpc_cluster_api_go.ConnectivityCheckerClient
}

// NewManager creates a Manager using a set of clients.
func NewManager(ConnectivityCheckerClient grpc_cluster_api_go.ConnectivityCheckerClient) Manager {
	return Manager{ConnectivityCheckerClient: ConnectivityCheckerClient}
}

func (m *Manager) ClusterAlive (ctx context.Context, clusterId *grpc_infrastructure_go.ClusterId) (*grpc_common_go.Success, derrors.Error) {
	return &grpc_common_go.Success{}, nil
}