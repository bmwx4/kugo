// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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

	"github.com/bmwx4/kugo/pkg/controller"
	"github.com/bmwx4/kugo/pkg/server"
	"github.com/spf13/cobra"
)

var (
	https bool
)

// httpdCmd represents the httpd command
var httpdCmd = &cobra.Command{
	Use:   "httpd",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("httpd called")
		ctl := controller.New()
		if https {
			go func() {
				httpsServer := server.NewHTTPSServer()
				httpsServer.Register(server.Healthz)
				httpsServer.Register(ctl)
				httpsServer.ListenAndServe()
			}()
		}
		httpServer := server.NewHTTPServer()
		httpServer.Register(server.Healthz)
		httpServer.Register(ctl)
		httpServer.ListenAndServe()
	},
}

func init() {
	RootCmd.AddCommand(httpdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//httpdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	httpdCmd.PersistentFlags().BoolVarP(&https, "https", "", false, "If enable the https server")
}
