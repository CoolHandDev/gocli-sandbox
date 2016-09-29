// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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

package main

import (
	"fmt"
	"os"

	//"github.com/coolhanddev/gocli-sandbox/cmd"
	"github.com/spf13/cobra"
)

var userName string
var passWord string
var server string
var database string

func main() {
	//cmd.Execute()

	//root command
	var rootCmd = &cobra.Command{
		Use:   "gocli",
		Short: "a sandbox for cli",
		Long:  "This is a sandbox for the cobra command line library",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("please pass a flag")
			}
		},
	}
	rootCmd.Flags().StringVarP(&userName, "username", "u", "", "The username for SQL Server")
	rootCmd.Flags().StringVarP(&passWord, "passWord", "p", "", "The password for SQL Server")
	rootCmd.Flags().StringVarP(&server, "server", "s", "", "The SQL Server hostname. use localhost if local")
	rootCmd.Flags().StringVarP(&database, "database", "d", "", "The database to use")
	cobra.MarkFlagRequired(rootCmd.Flags(), "userName")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println("the value of username:", userName)
}
