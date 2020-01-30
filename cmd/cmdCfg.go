package cmd

import (
	"github.com/spf13/cobra"
)

var oFlags OutFlags
var showSample bool
var cmdConfig = &cobra.Command{
	Use:   "config",
	Short: "Show and Edit awsctl config",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	Run: func(cmd *cobra.Command, args []string) {
		if showSample {
			genSample()
			return
		}
		cmd.Help()
	},
}

func init() {
	cmdConfig.PersistentFlags().StringVarP(&oFlags.Format, "out", "o", "yaml", "Output Format - yaml|json.")
	cmdConfig.Flags().BoolVar(&showSample, "sample", false, "Display a sample config file.")
	cmdConfig.AddCommand(cmdView)
	cmdConfig.AddCommand(cmdShowContext)
	cmdConfig.AddCommand(cmdUseContext)
}
