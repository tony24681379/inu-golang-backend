// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tony24681379/inu-golang-backend/server"
)

const (
	defaultElasticSearchIP = "0.0.0.0:9200"
	flagElasticSearchIP    = "elastic-search-ip"
	flagElasticSearchPort  = "elastic-search-port"
	flagAkkaIP             = "akka-ip"
)

type serverOptions struct {
	elasticSearchIP   string
	elasticSearchPort string
	akkaIP            string
}

func newServerCommand() {
	opts := serverOptions{}
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "Backend server serve frontend",
		Long:  `A golang backend server serve frontend which connects to ElasticSearch and Akka`,
		Run: func(cmd *cobra.Command, args []string) {
			runServer(opts)
		},
	}
	RootCmd.AddCommand(serverCmd)

	flags := serverCmd.Flags()
	flags.StringVar(&opts.elasticSearchIP, flagElasticSearchIP, "", "ElasticSearch IP address")
	flags.StringVar(&opts.elasticSearchPort, flagElasticSearchPort, "", "ElasticSearch port")
}

func runServer(opts serverOptions) {
	fmt.Println(flagElasticSearchIP, opts.elasticSearchIP)
	server.Server(opts.elasticSearchIP, opts.elasticSearchPort)
}
