// Copyright © 2018 NAME HERE <jbonds@jbvm.io>
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
	"os"

	"github.com/jbvmio/awsctl/cli/awsgo"
	"github.com/jbvmio/awsctl/cli/cmd/config"
	"github.com/jbvmio/awsctl/cli/cmd/ec2"
	"github.com/jbvmio/awsctl/cli/cmd/get"
	"github.com/jbvmio/awsctl/cli/x/out"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

var outFlags out.OutFlags

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "awsctl",
	Short: "awsctl: AWS Management Tool",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		awsgo.LaunchAWSClient(config.GetContext())
	},
	Run: func(cmd *cobra.Command, args []string) {
		switch true {
		case outFlags.Format != "":
			fmt.Println("outFlag")
		default:
			fmt.Println("awsctl")
		}
	},
}

// Execute starts here.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVarP(&outFlags.Format, "out", "o", "", "Change Output Format - wide|long|yaml|json.")

	rootCmd.AddCommand(config.CmdConfig)
	rootCmd.AddCommand(get.CmdGet)
	rootCmd.AddCommand(ec2.CmdRestartEc2)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".awsctl")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
	}
}
