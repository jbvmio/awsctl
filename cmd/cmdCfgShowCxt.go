package cmd

import (
	"github.com/spf13/cobra"
)

var cmdShowContext = &cobra.Command{
	Use:     "get-context",
	Aliases: []string{"current-context", "get-contexts", "get"},
	Short:   "Display current and available context details",
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case cmd.CalledAs() == "current-context":
			Marshal(GetContext(), oFlags.Format)
		case len(args) < 1:
			Marshal(GetContextList(), oFlags.Format)
		default:
			Marshal(GetContext(args...), oFlags.Format)
		}
	},
}

func init() {
}
