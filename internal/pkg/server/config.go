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

package server

import (
	"github.com/nalej/authx/pkg/interceptor"
	"github.com/nalej/cluster-api/version"
	"github.com/nalej/derrors"
	"github.com/rs/zerolog/log"
	"strings"
)

type Config struct {
	// Port where the gRPC API service will listen requests.
	Port int
	// NetworkManagerAddress with the Network Manager address.
	NetworkManagerAddress string
	// ConductorAddress with the Conductor Address.
	ConductorAddress string
	// DeviceManagerAddress with the Device Manager address
	DeviceManagerAddress string
	// AuthxAddress with the host:port to connect to the Authx manager.
	AuthxAddress string
	// ClusterWatcher with the host:port to connect to the cluster watcher.
	ClusterWatcherAddress string
	// AuthSecret contains the shared authx secret.
	AuthSecret string
	// AuthHeader contains the name of the target header.
	AuthHeader string
	// AuthConfigPath contains the path of the file with the authentication configuration.
	AuthConfigPath string
	// QueueAddress contains the URL of the message service queue
	QueueAddress string
	// ConnectivityManager address
	ConnectivityManagerAddress string
}

func (conf *Config) Validate() derrors.Error {

	if conf.Port <= 0 {
		return derrors.NewInvalidArgumentError("ports must be valid")
	}

	if conf.NetworkManagerAddress == "" {
		return derrors.NewInvalidArgumentError("networkManagerAddress must be set")
	}

	if conf.ConductorAddress == "" {
		return derrors.NewInvalidArgumentError("conductorAddress must be set")
	}

	if conf.DeviceManagerAddress == "" {
		return derrors.NewInvalidArgumentError("deviceManagerAddress must be set")
	}

	if conf.AuthxAddress == "" {
		return derrors.NewInvalidArgumentError("authxAddress must be set")
	}

	if conf.AuthHeader == "" || conf.AuthSecret == "" {
		return derrors.NewInvalidArgumentError("Authorization header and secret must be set")
	}

	if conf.AuthConfigPath == "" {
		return derrors.NewInvalidArgumentError("authConfigPath must be set")
	}

	if conf.QueueAddress == "" {
		return derrors.NewInvalidArgumentError("queueAddress must be set")
	}

	if conf.ClusterWatcherAddress == "" {
		return derrors.NewInvalidArgumentError("clusterWatcherAddress must be set")
	}

	if conf.ConnectivityManagerAddress == "" {
		return derrors.NewInvalidArgumentError("connectivityManagerAddress must be set")

	}

	return nil
}

// LoadAuthConfig loads the security configuration.
func (conf *Config) LoadAuthConfig() (*interceptor.AuthorizationConfig, derrors.Error) {
	return interceptor.LoadAuthorizationConfig(conf.AuthConfigPath)
}

func (conf *Config) Print() {
	log.Info().Str("app", version.AppVersion).Str("commit", version.Commit).Msg("Version")
	log.Info().Int("port", conf.Port).Msg("gRPC port")
	log.Info().Str("URL", conf.NetworkManagerAddress).Msg("Network Manager Service")
	log.Info().Str("URL", conf.ConductorAddress).Msg("Conductor Service")
	log.Info().Str("URL", conf.DeviceManagerAddress).Msg("Device Manager Service")
	log.Info().Str("URL", conf.AuthxAddress).Msg("Authx")
	log.Info().Str("header", conf.AuthHeader).Str("secret", strings.Repeat("*", len(conf.AuthSecret))).Msg("Authorization")
	log.Info().Str("path", conf.AuthConfigPath).Msg("Permissions file")
	log.Info().Str("queueAddress", conf.QueueAddress).Msg("Queue address")
	log.Info().Str("clusterWatcherAddress", conf.ClusterWatcherAddress).Msg("Cluster watcher address")
	log.Info().Str("connectivityManagerAddress", conf.ConnectivityManagerAddress).Msg("Connectivity Manager Address")

}
