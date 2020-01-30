package cmd

import (
	"github.com/spf13/cobra"
)

var genToken bool
var cmdECR = &cobra.Command{
	Use:   "ecr",
	Short: "Elastic Container Registry Ops",
	Run: func(cmd *cobra.Command, args []string) {
		if genToken {
			token := GenECRToken(awsFlags)
			Infof("%+v", token)
			return
		}
		PrintAWS(ListRepos(awsFlags), outFlags.Format)
	},
}

func init() {
	cmdECR.PersistentFlags().StringVar(&awsFlags.overrides.region, "region", "", "Desired Region.")
	cmdECR.PersistentFlags().StringVarP(&outFlags.Format, "out", "o", "", "Change Output Format - yaml|json.")
	cmdECR.Flags().BoolVarP(&genToken, "gen-token", "G", false, "Generate Auth Token for ECR.")
	cmdECR.AddCommand(cmdECRImages)
	cmdECR.AddCommand(cmdECRLogin)
}
