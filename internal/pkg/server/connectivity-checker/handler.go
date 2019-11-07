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

package connectivity_checker

import (
	"context"
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-infrastructure-go"
	"github.com/nalej/grpc-utils/pkg/conversions"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	Manager Manager
}

// NewHandler creates a new Handler with a linked manager.
func NewHandler(manager Manager) *Handler {
	return &Handler{manager}
}

func (h *Handler) ClusterAlive(ctx context.Context, clusterId *grpc_infrastructure_go.ClusterId) (*grpc_common_go.Success, error) {
	log.Debug().Str("clusterId", clusterId.ClusterId).Str("organizationId", clusterId.OrganizationId).Msg("cluster alive")
	err := h.ValidClusterId(clusterId)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	clusterAliveResult, err := h.Manager.ClusterAlive(ctx, clusterId)
	if err != nil {
		return nil, conversions.ToGRPCError(err)
	}
	return clusterAliveResult, nil
}

func (h *Handler) ValidClusterId(clusterId *grpc_infrastructure_go.ClusterId) derrors.Error {
	if clusterId.ClusterId == "" {
		return derrors.NewInvalidArgumentError("expecting ClusterId")
	}

	if clusterId.OrganizationId == "" {
		return derrors.NewInvalidArgumentError("expecting OrganizationId")
	}

	return nil
}
