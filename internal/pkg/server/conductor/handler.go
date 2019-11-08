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
