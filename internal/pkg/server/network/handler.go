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

package network

import (
	"context"
	"github.com/nalej/derrors"
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
// DEPRECATED. This function is no longer available in the cluster-api
func (h *Handler) AddDNSEntry(ctx context.Context, addRequest *grpc_network_go.AddDNSEntryRequest) (*grpc_common_go.Success, error) {
	return nil, derrors.NewUnimplementedError("AddDNSEntry operation is no loger available in the cluster-api")

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

// AuthorizeZTConnection A pod requests authorization to join a secondary ZT Network
func (h *Handler) AuthorizeZTConnection(ctx context.Context, request *grpc_network_go.AuthorizeZTConnectionRequest) (*grpc_common_go.Success, error) {
	log.Debug().Interface("request", request).Msg("authorize zt connection")
	return h.Manager.AuthorizeZTConnection(request)
}

// RegisterZTConnection operation to indicate that the inbound or outbound  are within the ztNetwork
func (h *Handler) RegisterZTConnection(ctx context.Context, request *grpc_network_go.RegisterZTConnectionRequest) (*grpc_common_go.Success, error) {
	log.Debug().Interface("request", request).Msg("register zt connection")
	return h.Manager.RegisterZTConnection(request)
}
