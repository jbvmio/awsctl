package config

import (
	"github.com/jbvmio/awsctl/cli/x/out"
	"github.com/spf13/cobra"
)

var outFlags out.OutFlags
var showSample bool

var CmdConfig = &cobra.Command{
	Use:   "config",
	Short: "Show and Edit awsctl config",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(cmd *cobra.Command, args []string) {
		if showSample {
			genSample()
			return
		}
	},
}

func init() {
	CmdConfig.PersistentFlags().StringVarP(&outFlags.Format, "out", "o", "yaml", "Output Format - yaml|json.")
	CmdConfig.Flags().BoolVar(&showSample, "sample", false, "Display a sample config file.")

	CmdConfig.AddCommand(cmdView)
	CmdConfig.AddCommand(cmdShowContext)
	CmdConfig.AddCommand(cmdUseContext)
}
