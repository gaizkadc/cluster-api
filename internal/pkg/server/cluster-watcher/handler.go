/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package cluster_watcher

import (
	"context"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-cluster-watcher-go"
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
func (h *Handler) AddClusterInfo(ctx context.Context, request *grpc_cluster_watcher_go.ClusterWatcherInfo) (*grpc_common_go.Success, error) {
	log.Debug().Interface("request", request).Msg("add cluster infor")
	return h.Manager.AddClusterInfo(request)
}
