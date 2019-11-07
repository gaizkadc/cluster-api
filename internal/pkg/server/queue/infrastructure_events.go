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

package queue

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-bus-go"
	"github.com/nalej/grpc-connectivity-manager-go"
	"github.com/nalej/nalej-bus/pkg/bus"
	"github.com/nalej/nalej-bus/pkg/queue"
)

const (
	InfrastructureEventsTopic = "nalej/infrastructure/events"
)

type InfrastructureEventsProducer struct {
	producer bus.NalejProducer
}

func NewInfrastructureEventsProducer(client bus.NalejClient, name string) (*InfrastructureEventsProducer, derrors.Error) {
	prod, err := client.BuildProducer(name, InfrastructureEventsTopic)
	if err != nil {
		return nil, err
	}
	return &InfrastructureEventsProducer{producer: prod}, nil
}

// Send messages to the network ops queue.
func (i *InfrastructureEventsProducer) Send(ctx context.Context, msg proto.Message) derrors.Error {
	var wrapper grpc_bus_go.InfrastructureEvents

	switch x := msg.(type) {
	case *grpc_connectivity_manager_go.ClusterAlive:
		wrapper = grpc_bus_go.InfrastructureEvents{Event: &grpc_bus_go.InfrastructureEvents_ClusterAlive{x}}
	default:
		return derrors.NewInvalidArgumentError("invalid proto message type")
	}

	payload, err := queue.MarshallPbMsg(&wrapper)
	if err != nil {
		return err
	}

	err = i.producer.Send(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}
