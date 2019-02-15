package ec2

import (
	"github.com/jbvmio/awsctl"
	"github.com/jbvmio/awsctl/cli/aws"
	"github.com/jbvmio/awsctl/cli/x/out"
	"github.com/spf13/cobra"
)

var CmdGetEc2 = &cobra.Command{
	Use:     "ec2",
	Aliases: []string{"inst", "instance", "instances"},
	Short:   "Get EC2 Information",
	Run: func(cmd *cobra.Command, args []string) {
		var instances []awsctl.Instance
		switch {
		default:
			instances = aws.GetEC2Instances(args...)
		}
		switch {
		case cmd.Flags().Changed("out"):
			outFmt, err := cmd.Flags().GetString("out")
			if err != nil {
				out.Warnf("WARN: %v", err)
			}
			out.PrintAWS(instances, outFmt)
		default:
			out.PrintAWS(instances)
		}
	},
}

func init() {
	//CmdGetEc2.PersistentFlags().StringVarP(&outFlags.Format, "out", "o", "", "Change Output Format - wide|long|yaml|json.")
}
