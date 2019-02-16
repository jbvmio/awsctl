package ec2

import (
	"github.com/jbvmio/awsctl"
	"github.com/jbvmio/awsctl/cli/awsgo"
	"github.com/jbvmio/awsctl/cli/x/out"
	"github.com/spf13/cobra"
)

var ec2Flags awsgo.EC2Flags

var CmdGetEc2 = &cobra.Command{
	Use:     "ec2",
	Aliases: []string{"inst", "instance", "instances"},
	Short:   "Get EC2 Information",
	Run: func(cmd *cobra.Command, args []string) {
		var instances []awsctl.Instance
		switch {
		default:
			instances = awsgo.GetEC2Instances(ec2Flags, args...)
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
	CmdGetEc2.Flags().StringVar(&ec2Flags.Region, "region", "", "Desired Region.")
	//CmdGetEc2.PersistentFlags().StringVarP(&outFlags.Format, "out", "o", "", "Change Output Format - wide|long|yaml|json.")
}
