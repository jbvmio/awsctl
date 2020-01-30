package cmd

import (
	"github.com/spf13/cobra"
)

var logsFlags LogsFlags
var logsRaw bool
var cmdCWLogs = &cobra.Command{
	Use:     "logs",
	Aliases: []string{"log"},
	Short:   "CloudWatch Logs",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			logsFlags.LogGroup = args[0]
			switch {
			case logsFlags.Stream != "":
				logs := GetLogStream(awsFlags, logsFlags)
				PrintAWS(logs)
				return
			default:
				streams := ListLogStreams(awsFlags, logsFlags)
				PrintAWS(streams, outFlags.Format)
				return
			}
		}
		groups := ListLogGroups(awsFlags)
		PrintAWS(groups, outFlags.Format)
	},
}

func init() {
	cmdCWLogs.Flags().StringVarP(&logsFlags.LogGroup, "group", "G", "", "Log Group to Target.")
	cmdCWLogs.Flags().StringVarP(&logsFlags.Stream, "stream", "S", "", "Log Stream to Target.")
	cmdCWLogs.Flags().BoolVar(&logsRaw, "raw", false, "Display Log Message Output Only, No Formatting.")
	cmdCWLogs.Flags().IntVarP(&logsFlags.Pages, "pages", "P", 1, "Pages to Return.")
	cmdCWLogs.Flags().DurationVar(&logsFlags.Last, "last", logsFlags.Last, "")
}
