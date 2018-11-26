/*
 * Copyright (C) 2018 Nalej - All Rights Reserved
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
	// AuthSecret contains the shared authx secret.
	AuthSecret string
	// AuthHeader contains the name of the target header.
	AuthHeader string
	// AuthConfigPath contains the path of the file with the authentication configuration.
	AuthConfigPath string
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

	if conf.AuthHeader == "" || conf.AuthSecret == "" {
		return derrors.NewInvalidArgumentError("Authorization header and secret must be set")
	}

	if conf.AuthConfigPath == "" {
		return derrors.NewInvalidArgumentError("authConfigPath must be set")
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
	log.Info().Str("header", conf.AuthHeader).Str("secret", strings.Repeat("*", len(conf.AuthSecret))).Msg("Authorization")
	log.Info().Str("path", conf.AuthConfigPath).Msg("Permissions file")
}
