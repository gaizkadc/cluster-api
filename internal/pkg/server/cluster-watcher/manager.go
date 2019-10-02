/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package cluster_watcher

import (
	"context"
	"github.com/nalej/grpc-cluster-watcher-go"
	"github.com/nalej/grpc-common-go"
)



// Manager structure with the required clients for network operations.
type Manager struct {
	ClusterWatcherClient grpc_cluster_watcher_go.ClusterWatcherMasterClient
}

// NewManager creates a Manager using a set of clients.
func NewManager(clusterWatcherMasterclient grpc_cluster_watcher_go.ClusterWatcherMasterClient) Manager {
	return Manager{ClusterWatcherClient: clusterWatcherMasterclient}
}

// AddClusterInfo
func (m *Manager) AddClusterInfo(request *grpc_cluster_watcher_go.ClusterWatchInfo) (*grpc_common_go.Success, error) {
	return m.ClusterWatcherClient.AddClusterInfo(context.Background(), request)
}
