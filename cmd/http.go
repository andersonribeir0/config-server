/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/andersonribeir0/config-server/webserver"

	"github.com/spf13/cobra"
)

var port string
var consulURL string
var consulPort string
var consulPrefix string
var consulAutoRefresh bool
var consulAutoRefreshDuration int64
var appName string

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Http config server",
	Long:  `Example of an http config server`,
	Run: func(cmd *cobra.Command, args []string) {
		server := webserver.Server{
			HttpPort:          		  port,
			ConsulURL:         		  consulURL,
			ConsulPort:        		  consulPort,
			ConsulPrefix:             consulPrefix,
			ConsulAutoRefresh:        consulAutoRefresh,
			ConsulAutoRefreshSeconds: consulAutoRefreshDuration,
			AppName:                  appName,
		}
		server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
	// Here you will define your flags and configuration settings.
	httpCmd.
		Flags().
		StringVarP(&port, "port", "p", ":4040", "Port to be used on http server.")
	httpCmd.
		Flags().
		StringVarP(&consulURL, "consul-url", "", "consul", "Consul URL.")
	httpCmd.
		Flags().
		StringVarP(&consulPort, "consul-port", "", ":8500", "Port to be used on consul.")
	httpCmd.
		Flags().
		StringVarP(&consulPrefix, "consul-prefix", "", "config", "Consul key value prefix.")
	httpCmd.
		Flags().
		BoolVarP(&consulAutoRefresh, "auto-refresh", "", true,
			"Refresh key-value pairs from consult integration.")
	httpCmd.
		Flags().
		Int64VarP(&consulAutoRefreshDuration, "auto-refresh-duration", "", 5,
			"Interval between key-value pairs refreshes in seconds.")
	httpCmd.
		Flags().
		StringVarP(&appName, "app-name", "n", "config-server", "Application name.")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
