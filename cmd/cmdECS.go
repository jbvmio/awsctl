package cmd

import (
	"github.com/spf13/cobra"
)

var ecsFlags ECSFlags
var cmdECS = &cobra.Command{
	Use:     "ecs",
	Aliases: []string{"fargate"},
	Short:   "ECS/Fargate Operations",
	Run: func(cmd *cobra.Command, args []string) {
		//Infof("%+v\n", ListECSClusters(awsFlags))
		//PrintAWS(DescribeClusters(awsFlags, ecsFlags), outFlags.Format)
		//DescribeECSServices(awsFlags, )
		PrintAWS(ListECSServices(awsFlags, ecsFlags))
		//DescribeECSServices(awsFlags, ecsFlags)
	},
}

func init() {
	cmdECS.PersistentFlags().StringVar(&awsFlags.overrides.region, "region", "", "Desired Region.")
	cmdECS.PersistentFlags().StringVarP(&outFlags.Format, "out", "o", "", "Change Output Format - yaml|json.")
	cmdECS.PersistentFlags().StringVarP(&ecsFlags.Cluster, "cluster", "C", "", "Cluster to Target.")
	cmdECS.AddCommand(cmdECSServices)
}
