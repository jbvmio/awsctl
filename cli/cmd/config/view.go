package config

import (
	"github.com/jbvmio/awsctl/cli/x/out"
	"github.com/spf13/cobra"
)

var cmdView = &cobra.Command{
	Use:     "view",
	Aliases: []string{"show"},
	Short:   "Display awsctl configurations",
	Run: func(cmd *cobra.Command, args []string) {
		out.Marshal(GetConfig(), outFlags.Format)
	},
}

func init() {
}
