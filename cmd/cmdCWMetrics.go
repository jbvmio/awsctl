package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

var metricFlags MetricFlags
var cmdCWMetrics = &cobra.Command{
	Use:     "metrics",
	Aliases: []string{"metric"},
	Short:   "CloudWatch Metrics",
	Run: func(cmd *cobra.Command, args []string) {
		metrics := ListMetrics(awsFlags, metricFlags)
		if metricFlags.ShowData {
			data := GetMetricData(metricFlags, metrics.Metrics)
			PrintAWS(data)
			return
		}
		switch {
		case cmd.Flags().Changed("out"):
			outFmt, err := cmd.Flags().GetString("out")
			if err != nil {
				Warnf("WARN: %v", err)
			}
			PrintAWS(metrics, outFmt)
		default:
			PrintAWS(metrics)
		}
	},
}

func init() {
	cmdCWMetrics.Flags().StringVarP(&outFlags.Format, "out", "o", "", "Change Output Format - yaml|json.")
	cmdCWMetrics.Flags().StringVarP(&metricFlags.Namespace, "namespace", "N", "", "Filter by Namespace.")
	cmdCWMetrics.Flags().StringVarP(&metricFlags.MetricName, "metric", "M", "", "Filter by Metric Name.")
	cmdCWMetrics.Flags().BoolVarP(&metricFlags.ShowData, "data", "D", false, "Return Metric Data.")
	cmdCWMetrics.Flags().StringVarP(&metricFlags.Tags, "tags", "T", "", `Filter by Comma Delimited Tags, EX: "Type=API,Resource=ListMetrics"`)
	cmdCWMetrics.Flags().DurationVarP(&metricFlags.Last, "last", "L", time.Duration(time.Hour*3), "Last Duration.")
}
