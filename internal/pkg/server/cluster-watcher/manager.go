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
