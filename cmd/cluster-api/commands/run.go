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

package commands

import (
	"github.com/nalej/cluster-api/internal/pkg/server"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var config = server.Config{}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run Cluster API",
	Long:  `Run Cluster API`,
	Run: func(cmd *cobra.Command, args []string) {
		SetupLogging()
		log.Info().Msg("Launching API!")
		server := server.NewService(config)
		server.Run()
	},
}

func init() {
	runCmd.Flags().IntVar(&config.Port, "port", 8280, "Port to launch the Public gRPC API")
	runCmd.PersistentFlags().StringVar(&config.NetworkManagerAddress, "networkManagerAddress", "localhost:8000",
		"Network Manager address (host:port)")
	runCmd.PersistentFlags().StringVar(&config.ConductorAddress, "conductorAddress", "localhost:5000",
		"Conductor address (host:port)")
	runCmd.PersistentFlags().StringVar(&config.AuthHeader, "authHeader", "", "Authorization Header")
	runCmd.PersistentFlags().StringVar(&config.AuthSecret, "authSecret", "", "Authorization secret")
	runCmd.PersistentFlags().StringVar(&config.AuthConfigPath, "authConfigPath", "", "Authorization config path")
	runCmd.PersistentFlags().StringVar(&config.DeviceManagerAddress, "deviceManagerAddress", "", "localhost:6010")
	runCmd.PersistentFlags().StringVar(&config.AuthxAddress, "authxAddress", "localhost:8810",
		"Authx address (host:port)")
	runCmd.PersistentFlags().StringVar(&config.QueueAddress, "queueAddress", "localhost:6650", "localhost:6650")
	runCmd.PersistentFlags().StringVar(&config.ClusterWatcherAddress, "clusterWatcherAddress", "localhost:7777", "localhost:7777")
	runCmd.PersistentFlags().StringVar(&config.ConnectivityManagerAddress, "connectivityManagerAddress", "localhost:8384", "localhost:8384")
	rootCmd.AddCommand(runCmd)
}
