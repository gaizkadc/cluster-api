/*
 * Copyright (C) 2018 Nalej - All Rights Reserved
 */

package conductor

import (
	"context"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-conductor-go"
	"github.com/rs/zerolog/log"
)

// Handler structure for the conductor requests.
type Handler struct {
	Manager Manager
}

// NewHandler creates a new Handler with a linked manager.
func NewHandler(manager Manager) *Handler {
	return &Handler{manager}
}

// Request to update a fragment plan status with all the corresponding services
func (h *Handler) UpdateDeploymentFragmentStatus(ctx context.Context, updateRequest *grpc_conductor_go.DeploymentFragmentUpdateRequest) (*grpc_common_go.Success, error) {
	log.Debug().Interface("request", updateRequest).Msg("update deployment fragment status")
	return h.Manager.UpdateDeploymentFragmentStatus(updateRequest)
}

// Update the status of a set of services during a given fragment deployment
func (h *Handler) UpdateServiceStatus(ctx context.Context, updateRequest *grpc_conductor_go.DeploymentServiceUpdateRequest) (*grpc_common_go.Success, error) {
	log.Debug().Interface("request", updateRequest).Msg("update service status")
	return h.Manager.UpdateServiceStatus(updateRequest)
}

