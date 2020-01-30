package cmd

import (
	"github.com/jbvmio/awsgo"
	"github.com/spf13/cobra"
)

var awsFlags AWSFlags

var cmdGetEc2 = &cobra.Command{
	Use:     "ec2",
	Aliases: []string{"inst", "instance", "instances"},
	Short:   "Get EC2 Information",
	Run: func(cmd *cobra.Command, args []string) {
		var instances []awsgo.Instance
		switch {
		default:
			instances = GetEC2Instances(awsFlags, args...)
		}
		PrintAWS(instances)
	},
}

func init() {
	cmdGetEc2.PersistentFlags().StringVar(&awsFlags.overrides.region, "region", "", "Desired Region.")
	cmdGetEc2.Flags().StringVarP(&outFlags.Format, "out", "o", "", "Change Output Format - wide|long|yaml|json.")
}
