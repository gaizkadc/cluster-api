package device_latency

import (
	"context"
	"github.com/nalej/grpc-authx-go"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-device-controller-go"
	"github.com/nalej/grpc-device-go"
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

// GetDeviceGroupSecret retrieves the secret associated with a device group
func (h * Handler) 	GetDeviceGroupSecret(ctx context.Context, deviceGroupId *grpc_device_go.DeviceGroupId) (*grpc_authx_go.DeviceGroupSecret, error){
	return h.Manager.GetDeviceGroupSecret(deviceGroupId)
}