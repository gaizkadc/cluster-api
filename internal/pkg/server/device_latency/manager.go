package device_latency

import (
	"context"
	"github.com/nalej/grpc-authx-go"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-device-controller-go"
	"github.com/nalej/grpc-device-go"
	"github.com/nalej/grpc-device-manager-go"
)

type Manager struct {
  DeviceClient grpc_device_manager_go.LatencyClient
	Authx grpc_authx_go.AuthxClient
}

func NewManager(dLatencyClient grpc_device_manager_go.LatencyClient, authxClient grpc_authx_go.AuthxClient) Manager {
	return Manager{
		DeviceClient: dLatencyClient,
		Authx:authxClient,
	}
}

func (m * Manager ) RegisterLatency(request *grpc_device_controller_go.RegisterLatencyRequest) (*grpc_common_go.Success, error) {
	return m.DeviceClient.RegisterLatency(context.Background(), request)
}

func (m * Manager ) GetDeviceGroupSecret(deviceGroupId *grpc_device_go.DeviceGroupId) (*grpc_authx_go.DeviceGroupSecret, error){
	return m.Authx.GetDeviceGroupSecret(context.Background(), deviceGroupId)
}