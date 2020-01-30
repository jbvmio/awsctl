package cmd

import (
	"github.com/spf13/cobra"
)

var ecsSvcDescribe bool
var cmdECSServices = &cobra.Command{
	Use:     "services",
	Aliases: []string{"service", "svc"},
	Short:   "ECS/Fargate Service Operations",
	Run: func(cmd *cobra.Command, args []string) {
		svc := ListECSServices(awsFlags, ecsFlags)
		if ecsSvcDescribe {
			PrintAWS(DescribeECSServices(svc), outFlags.Format)
			return
		}
		PrintAWS(svc, outFlags.Format)
	},
}

func init() {
	cmdECSServices.Flags().StringVarP(&ecsFlags.Service, "service", "S", "", "Service to Target.")
	cmdECSServices.Flags().BoolVarP(&ecsSvcDescribe, "describe", "D", false, "Additional Details for a Targeted Service.")
}
