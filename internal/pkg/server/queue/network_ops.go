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
    "github.com/nalej/nalej-bus/pkg/bus"
    "github.com/nalej/nalej-bus/pkg/queue/network/ops"
)

// Structures and operations designed to manipulate the infrastructure events queue.
type NetworkOpsBusManager struct {
    producer *ops.NetworkOpsProducer
}

// Create a new manager for the network ops topic.
// params:
//  client the implementation of que queueing protocol
//  name of the producer to be generated
// return:
//  instance and error if any
func NewNetworkOpsBusManager(client bus.NalejClient, name string) (*NetworkOpsBusManager, derrors.Error) {
    producer, err := ops.NewNetworkOpsProducer(client, name)
    if err != nil {
        return nil, err
    }
    return &NetworkOpsBusManager{producer: producer}, nil
}

// Send messages to the network ops queue.
func (n *NetworkOpsBusManager) Send(ctx context.Context, msg proto.Message) derrors.Error {
    return n.producer.Send(ctx,msg)
}

