/*
 * Copyright (C) 2018 Nalej - All Rights Reserved
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
	runCmd.PersistentFlags().StringVar(&config.ConnectivityManagerAddress, "connectivityManagerAddress", "localhost:8384", "localhost:8384")
	rootCmd.AddCommand(runCmd)
}
