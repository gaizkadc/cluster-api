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
	"github.com/nalej/grpc-connectivity-manager-go"
	"github.com/nalej/grpc-infrastructure-go"
	"github.com/nalej/nalej-bus/pkg/queue/infrastructure/events"
	"github.com/rs/zerolog/log"
	"time"
)

const (
	DefaultTimeout =  time.Minute
)

// Manager structure with the required clients for connectivity-checker operations.
type Manager struct {
	infrastructureEventsProducer *events.InfrastructureEventsProducer
}

// NewManager creates a Manager using a set of clients.
func NewManager(infraEventsProducer *events.InfrastructureEventsProducer) Manager {
	return Manager{
		infrastructureEventsProducer: infraEventsProducer,
	}
}

func (m *Manager) ClusterAlive (ctx context.Context, clusterId *grpc_infrastructure_go.ClusterId) (*grpc_common_go.Success, derrors.Error) {

	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	clusterAlive := clusterIdToClusterAlive(clusterId)
	err := m.infrastructureEventsProducer.Send(ctx, clusterAlive)
	if err != nil {
		log.Error().Err(err).Str("clusterId", clusterId.ClusterId).
			Msg("error when sending cluster alive check to the queue")
		return nil, err
	}

	return &grpc_common_go.Success{}, nil
}

func clusterIdToClusterAlive (clusterId *grpc_infrastructure_go.ClusterId) *grpc_connectivity_manager_go.ClusterAlive {
	return &grpc_connectivity_manager_go.ClusterAlive{
		OrganizationId:       clusterId.OrganizationId,
		ClusterId:            clusterId.ClusterId,
		Timestamp:            time.Now().Unix(),
	}
}