package cmd

import (
	"strings"

	"github.com/jbvmio/awsgo"
	"github.com/spf13/cobra"
)

var cmdRestartEc2 = &cobra.Command{
	Use:     "restart",
	Aliases: []string{"stop", "start", "reboot"},
	Short:   "Restart an EC2 Instance",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var isc []awsgo.InstanceStateChange
		switch {
		case strings.Contains(cmd.CalledAs(), "reboot"):
			RebootEC2Instances(awsFlags, args...)
			return
		case strings.Contains(cmd.CalledAs(), "start"):
			isc = StartEC2Instances(awsFlags, args...)
		case strings.Contains(cmd.CalledAs(), "stop"):
			isc = StopEC2Instances(awsFlags, args...)
		}
		switch {
		case cmd.Flags().Changed("out"):
			outFmt, err := cmd.Flags().GetString("out")
			if err != nil {
				Warnf("WARN: %v", err)
			}
			PrintAWS(isc, outFmt)
		default:
			PrintAWS(isc)
		}
	},
}

func init() {
	cmdRestartEc2.Flags().StringVar(&awsFlags.overrides.region, "region", "", "Desired Region.")
}
