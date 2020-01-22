package cw

import (
	"github.com/jbvmio/awsctl"
	"github.com/jbvmio/awsctl/cli/awsgo"
	"github.com/jbvmio/awsctl/cli/x/out"
)

// CmdCW starts the cloudwatch command:
var CmdCW = &cobra.Command{
	Use:     "cw",
	Aliases: []string{"cloudwatch"},
	Short:   "CloudWatch Operations",
	Run: func(cmd *cobra.Command, args []string) {

		awsgo.ListMetrics(ec2Flags)
		/*
			switch true {
			case len(args) > 0:
				out.Failf("No such resource: %v", args[0])
			default:
				cmd.Help()
			}
		*/
	},
}

func init() {
	CmdCW.Flags().StringVar(&ec2Flags.Region, "region", "", "Desired Region.")
}
