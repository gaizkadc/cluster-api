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