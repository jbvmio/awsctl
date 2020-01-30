package cmd

import (
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/jbvmio/awsgo"
)

// MetricFlags defines options for CloudWatch Metric Ops.
type MetricFlags struct {
	Namespace  string
	MetricName string
	Tags       string
	Last       time.Duration
	ShowData   bool
}

// GetMetricData returns CloudWatch Metric Data.
func GetMetricData(mFlags MetricFlags, metrics []*cloudwatch.Metric) *cloudwatch.GetMetricDataOutput {
	timeNow := time.Now()
	timeThen := timeNow.Add(-mFlags.Last)
	var Q []*cloudwatch.MetricDataQuery
	for _, m := range metrics {
		tags := makeTagList(m)
		mID := encodeString(*m.MetricName + `:` + tags)
		q := cloudwatch.MetricDataQuery{
			Id: aws.String(mID),
			MetricStat: &cloudwatch.MetricStat{
				Metric: m,
				Period: aws.Int64(60),
				Stat:   aws.String("Average"),
			},
		}
		Q = append(Q, &q)
	}
	input := cloudwatch.GetMetricDataInput{
		EndTime:           &timeNow,
		StartTime:         &timeThen,
		MetricDataQueries: Q,
	}
	output, err := client.CW().GetMetricData(&input)
	if err != nil {
		Failf("error retrieving metric statistics: %v", err)
	}
	return output
}

// GetMetricStats returns CloudWatch Metric Data.
func GetMetricStats(aFlags AWSFlags, mFlags MetricFlags) []*cloudwatch.GetMetricStatisticsOutput {
	var output []*cloudwatch.GetMetricStatisticsOutput
	timeNow := time.Now()
	timeThen := timeNow.Add(-mFlags.Last)
	metrics := ListMetrics(aFlags, mFlags).Metrics
	for _, m := range metrics {
		input := cloudwatch.GetMetricStatisticsInput{
			Dimensions: m.Dimensions,
			MetricName: m.MetricName,
			Namespace:  m.Namespace,
			Period:     aws.Int64(60),
			StartTime:  &timeThen,
			EndTime:    &timeNow,
			Statistics: []*string{
				aws.String("SampleCount"),
				aws.String("Average"),
				aws.String("Sum"),
				aws.String("Minimum"),
				aws.String("Maximum"),
			},
		}
		out, err := client.CW().GetMetricStatistics(&input)
		if err != nil {
			Failf("error retrieving metric statistics: %v", err)
		}
		output = append(output, out)
	}
	return output
}

// ListMetrics returns CloudWatch Metrics.
func ListMetrics(aFlags AWSFlags, mFlags MetricFlags) *cloudwatch.ListMetricsOutput {
	client.AddConfig(awsgo.SvcTypeCW, aFlags)
	var input *cloudwatch.ListMetricsInput
	switch {
	case mFlags.Namespace == "" && mFlags.MetricName == "" && mFlags.Tags == "":
		input = nil
	default:
		i := cloudwatch.ListMetricsInput{}
		if mFlags.Namespace != "" {
			i.Namespace = &mFlags.Namespace
		}
		if mFlags.MetricName != "" {
			i.MetricName = &mFlags.MetricName
		}
		if mFlags.Tags != "" {
			i.Dimensions = makeDimensions(mFlags.Tags)
		}
		input = &i
	}
	output, err := client.CW().ListMetrics(input)
	if err != nil {
		Failf("error listing metrics: %v", err)
	}
	return output
}

func makeDimensions(tags string) (D []*cloudwatch.DimensionFilter) {
	pairs := strings.Split(tags, `,`)
	for _, pair := range pairs {
		kv := strings.Split(pair, `=`)
		switch len(kv) {
		case 1:
			d := cloudwatch.DimensionFilter{
				Name: &kv[0],
			}
			D = append(D, &d)
		case 2:
			d := cloudwatch.DimensionFilter{
				Name:  &kv[0],
				Value: &kv[1],
			}
			D = append(D, &d)
		}
	}
	return D
}

func makeTagList(metric *cloudwatch.Metric) string {
	var list string
	switch len(metric.Dimensions) {
	case 0:
		return ""
	default:
		list = *metric.Dimensions[0].Name
		list += `=`
		list += *metric.Dimensions[0].Value
		for _, m := range metric.Dimensions[1:] {
			list += `,` + *m.Name + `=` + *m.Value
		}
	}
	return list
}

func encodeString(value string) string {
	return `metric_` + fmt.Sprintf("%x", []byte(value))
}

func decodeString(value string) string {
	data := strings.TrimPrefix(value, `metric_`)
	decode, err := hex.DecodeString(data)
	if err != nil {
		Failf("error decoding string [%s]: %v", value, err)
	}
	return string(decode)
}
