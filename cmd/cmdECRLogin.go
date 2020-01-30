package cmd

import (
	"github.com/spf13/cobra"
)

var cmdECRLogin = &cobra.Command{
	Use:   "login",
	Short: "Docker Login to ECR Registry",
	Run: func(cmd *cobra.Command, args []string) {
		LoginECR(awsFlags)
	},
}

func init() {
}
