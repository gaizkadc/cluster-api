package connectivity_checker

import (
	"context"
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-infrastructure-go"
)

// Manager structure with the required clients for connectivity-checker operations.
type Manager struct {
}

// NewManager creates a Manager using a set of clients.
func NewManager() Manager {
	return Manager{}
}

func (m *Manager) ClusterAlive (ctx context.Context, clusterId *grpc_infrastructure_go.ClusterId) (*grpc_common_go.Success, derrors.Error) {
	return &grpc_common_go.Success{}, nil
}
