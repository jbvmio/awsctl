package cmd

import (
	"github.com/spf13/cobra"
)

// CmdCW starts the cloudwatch command:
var cmdKinesis = &cobra.Command{
	Use:     "kinesis",
	Aliases: []string{"kin"},
	Short:   "Kinesis Operations",
	Run: func(cmd *cobra.Command, args []string) {
		listKinesisStreams(awsFlags)
		//Infof("%v", getKinesisStreams(awsFlags))
		//getShardIterator(awsFlags)
		//getKinesisRecords(awsFlags)

		//consumeKinesis(awsFlags)

		//allKinesisStreams(awsFlags)
	},
}

func init() {
	cmdKinesis.PersistentFlags().StringVar(&awsFlags.overrides.region, "region", "", "Desired Region.")
}
