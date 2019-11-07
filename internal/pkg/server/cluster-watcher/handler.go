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
	"github.com/rs/zerolog/log"
)

// Handler structure for the network requests.
type Handler struct {
	Manager Manager
}

// NewHandler creates a new Handler with a linked manager.
func NewHandler(manager Manager) *Handler {
	return &Handler{manager}
}

// Add this cluster info
func (h *Handler) AddClusterInfo(ctx context.Context, request *grpc_cluster_watcher_go.ClusterWatchInfo) (*grpc_common_go.Success, error) {
	log.Debug().Interface("request", request).Msg("add cluster information")
	return h.Manager.AddClusterInfo(request)
}
