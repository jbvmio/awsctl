package cmd

import (
	"github.com/spf13/cobra"
)

// CmdCW starts the cloudwatch command:
var cmdCW = &cobra.Command{
	Use:     "cw",
	Aliases: []string{"cloudwatch"},
	Short:   "CloudWatch Operations",
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
	cmdCW.PersistentFlags().StringVar(&awsFlags.overrides.region, "region", "", "Desired Region.")
	cmdCW.AddCommand(cmdCWMetrics)
	cmdCW.AddCommand(cmdCWLogs)
}
