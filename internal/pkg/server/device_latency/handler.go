package device_latency

import (
	"context"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-device-controller-go"
)

type Handler struct {
	Manager Manager
}

// NewHandler creates a new Handler with a linked manager.
func NewHandler(manager Manager) *Handler {
	return &Handler{manager}
}

func (h * Handler) 	RegisterLatency(ctx context.Context, request *grpc_device_controller_go.RegisterLatencyRequest) (*grpc_common_go.Success, error) {
	return h.Manager.RegisterLatency(request)
}