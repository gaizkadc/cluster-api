package device_latency

import (
	"context"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-device-controller-go"
	"github.com/nalej/grpc-device-manager-go"
)

type Manager struct {
  DeviceClient grpc_device_manager_go.LatencyClient
}

func NewManager(dLatencyClient grpc_device_manager_go.LatencyClient) Manager {
	return Manager{
		DeviceClient: dLatencyClient,
	}
}

func (m * Manager ) RegisterLatency(request *grpc_device_controller_go.RegisterLatencyRequest) (*grpc_common_go.Success, error) {
	return m.DeviceClient.RegisterLatency(context.Background(), request)
}