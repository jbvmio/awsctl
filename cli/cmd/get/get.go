package get

import (
	"github.com/jbvmio/awsctl/cli/cmd/ec2"
	"github.com/jbvmio/awsctl/cli/x/out"
	"github.com/spf13/cobra"
)

var outFlags out.OutFlags

var CmdGet = &cobra.Command{
	Use:   "get",
	Short: "Get AWS Information",
	Run: func(cmd *cobra.Command, args []string) {
		switch true {
		case len(args) > 0:
			out.Failf("No such resource: %v", args[0])
		default:
			cmd.Help()
		}
	},
}

func init() {
	CmdGet.PersistentFlags().StringVarP(&outFlags.Format, "out", "o", "", "Change Output Format - yaml|json.")

	CmdGet.AddCommand(ec2.CmdGetEc2)
}
