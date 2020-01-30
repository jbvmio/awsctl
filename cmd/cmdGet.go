package cmd

import (
	"github.com/spf13/cobra"
)

var cmdGet = &cobra.Command{
	Use:   "get",
	Short: "Get AWS Information",
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case len(args) > 0:
			Failf("No such resource: %v", args[0])
		default:
			cmd.Help()
		}
	},
}

func init() {
	cmdGet.PersistentFlags().StringVar(&awsFlags.overrides.region, "region", "", "Desired Region.")
	cmdGet.AddCommand(cmdGetEc2)
	cmdGet.AddCommand(cmdCW)
}
