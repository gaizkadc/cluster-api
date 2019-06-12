/*
 * Copyright (C) 2018 Nalej - All Rights Reserved
 */

package network

import (
	"context"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-network-go"
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

// Authorize member
func (h *Handler) AuthorizeMember(ctx context.Context, request *grpc_network_go.AuthorizeMemberRequest) (*grpc_common_go.Success, error) {
	log.Debug().Interface("request", request).Msg("authorize member")
	return h.Manager.AuthorizeMember(request)
}

// AddDNSEntry creates a new DNSEntry on the system.
func (h *Handler) AddDNSEntry(ctx context.Context, addRequest *grpc_network_go.AddDNSEntryRequest) (*grpc_common_go.Success, error) {
	log.Debug().Interface("request", addRequest).Msg("add DNS Entry")
	return h.Manager.AddDNSEntry(addRequest)
}

// RegisterInboundServiceProxy operation to update rules based on new service proxy being created.
func (h *Handler) RegisterInboundServiceProxy(ctx context.Context, request *grpc_network_go.InboundServiceProxy) (*grpc_common_go.Success, error) {
	log.Debug().Interface("request", request).Msg("register inbound service proxy")
	return h.Manager.RegisterInboundServiceProxy(request)
}
// RegisterOutboundProxy operation to retrieve existing networking rules.
func (h *Handler) RegisterOutboundProxy(ctx context.Context, request *grpc_network_go.OutboundService) (*grpc_common_go.Success, error) {
	return h.Manager.RegisterOutboundProxy(request)
}