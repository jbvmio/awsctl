package cmd

import (
	"github.com/spf13/cobra"
)

//var showDefaults bool
var cmdView = &cobra.Command{
	Use:     "view",
	Aliases: []string{"show"},
	Short:   "Display awsctl config",
	Run: func(cmd *cobra.Command, args []string) {
		Marshal(GetConfig(), oFlags.Format)
	},
}

func init() {
	//cmdView.Flags().StringVarP(&oFlags.Format, "out", "o", "yaml", "Output Format - yaml|json.")
	//cmdView.Flags().BoolVar(&showDefaults, "defaults", false, "Display Configured AWS Defaults.")
}
