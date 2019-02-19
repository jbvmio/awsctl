package ec2

import (
	"strings"

	"github.com/jbvmio/awsctl"
	"github.com/jbvmio/awsctl/cli/awsgo"
	"github.com/jbvmio/awsctl/cli/x/out"
	"github.com/spf13/cobra"
)

var CmdRestartEc2 = &cobra.Command{
	Use:     "restart",
	Aliases: []string{"stop", "start", "reboot"},
	Short:   "Restart an EC2 Instance",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var isc []awsctl.InstanceStateChange
		switch {
		case strings.Contains(cmd.CalledAs(), "reboot"):
			awsgo.RebootEC2Instances(ec2Flags, args...)
			return
		case strings.Contains(cmd.CalledAs(), "start"):
			isc = awsgo.StartEC2Instances(ec2Flags, args...)
		case strings.Contains(cmd.CalledAs(), "stop"):
			isc = awsgo.StopEC2Instances(ec2Flags, args...)
		}
		switch {
		case cmd.Flags().Changed("out"):
			outFmt, err := cmd.Flags().GetString("out")
			if err != nil {
				out.Warnf("WARN: %v", err)
			}
			out.PrintAWS(isc, outFmt)
		default:
			out.PrintAWS(isc)
		}
	},
}

func init() {
	CmdRestartEc2.Flags().StringVar(&ec2Flags.Region, "region", "", "Desired Region.")
	//CmdRestartEc2.PersistentFlags().StringVarP(&outFlags.Format, "out", "o", "", "Change Output Format - wide|long|yaml|json.")
}
