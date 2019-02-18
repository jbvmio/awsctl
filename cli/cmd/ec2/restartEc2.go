package ec2

import (
	"fmt"
	"strings"

	"github.com/jbvmio/awsctl"
	"github.com/jbvmio/awsctl/cli/awsgo"
	"github.com/jbvmio/awsctl/cli/x/out"
	"github.com/spf13/cobra"
)

var CmdRestartEc2 = &cobra.Command{
	Use:     "restart",
	Aliases: []string{"stop", "start"},
	Short:   "Restart an EC2 Instance",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var isc []awsctl.InstanceStateChange
		switch {
		case strings.Contains(cmd.CalledAs(), "restart"):
			fmt.Println("restart Called")
		case strings.Contains(cmd.CalledAs(), "start"):
			fmt.Println("start Called")
			isc = awsgo.StartEC2Instances(ec2Flags, args...)
		case strings.Contains(cmd.CalledAs(), "stop"):
			fmt.Println("stop Called")
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
