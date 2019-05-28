/*
 * Copyright (C) 2018 Nalej - All Rights Reserved
 */

package queue

import (
    "context"
    "github.com/golang/protobuf/proto"
    "github.com/nalej/derrors"
    "github.com/nalej/nalej-bus/pkg/bus"
    "github.com/nalej/nalej-bus/pkg/queue/network/ops"
)

// Structures and operations designed to manipulate the networks operation queue.

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

