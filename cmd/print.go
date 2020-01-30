package cmd

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/fatih/color"
	"github.com/jbvmio/awsgo"
	"github.com/rodaine/table"
)

const (
	timeLayout = `2006-01-02 15:04:05`
)

// OutFlags .
type OutFlags struct {
	Format string
	Header bool
}

// PrintAWS .
func PrintAWS(i interface{}, format ...string) {
	var f string
	if len(format) > 0 {
		f = format[0]
	}
	if f == "" {
		printAws(i)
		return
	}
	switch f {
	case "wide":
		printAwsWide(i)
	case "long":
		printAwsLong(i)
	default:
		IfErrWarnf(Marshal(i, f))
	}
}

func printAws(i interface{}) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	var tbl table.Table
	switch i := i.(type) {
	case []awsgo.Instance:
		tbl = table.New("INDEX", "ID", "STATE", "IP", "HOSTNAME")
		for _, v := range i {
			tbl.AddRow(v.Index, v.ID, v.State, v.PublicIP, v.PublicDNSName)
		}
	case []awsgo.InstanceStateChange:
		tbl = table.New("ID", "PREVIOUS", "CURRENT")
		for _, v := range i {
			tbl.AddRow(v.ID, v.PreviousState, v.CurrentState)
		}
	case *cloudwatch.ListMetricsOutput:
		tbl = table.New("NAMESPACE", "METRIC", "TAGS")
		for _, v := range i.Metrics {
			t := makeTagList(v)
			tbl.AddRow(*v.Namespace, *v.MetricName, t)
		}
	case *cloudwatch.GetMetricDataOutput:
		tbl = table.New("METRIC", "TAGS", "VALUE", "TIMESTAMP")
		for _, v := range i.MetricDataResults {
			for x := 0; x < len(v.Timestamps); x++ {
				mt := strings.Split(decodeString(*v.Id), `:`)
				if len(mt) > 1 {
					tbl.AddRow(mt[0], mt[1], *v.Values[x], *v.Timestamps[x])
				}
			}
		}
	case []*cloudwatchlogs.LogGroup:
		tbl = table.New("LOGGROUP", "BYTES", "RETENTIONDAYS", "KMSKEY")
		for _, v := range i {
			retention := aws.Int64(-1)
			kmsKey := aws.String(`NA`)
			if v.RetentionInDays != nil {
				retention = v.RetentionInDays
			}
			if v.KmsKeyId != nil {
				kmsKey = v.KmsKeyId
			}
			tbl.AddRow(*v.LogGroupName, *v.StoredBytes, *retention, *kmsKey)
		}
	case []*cloudwatchlogs.LogStream:
		tbl = table.New("STREAM", "FIRSTEVENT", "LASTINGEST", "BYTES")
		for _, v := range i {
			var first, final string
			if *v.StoredBytes > 0 {
				first = time.Unix((*v.FirstEventTimestamp / 1000), 0).Format(timeLayout)
				//last = time.Unix((*v.LastEventTimestamp / 1000), 0).Format(timeLayout)
				final = time.Unix((*v.LastIngestionTime / 1000), 0).Format(timeLayout)
			}
			tbl.AddRow(*v.LogStreamName, first, final, *v.StoredBytes)
		}
	case []*cloudwatchlogs.OutputLogEvent:
		switch {
		case logsRaw:
			for _, v := range i {
				Infof("%s", *v.Message)
			}
		default:
			for _, v := range i {
				Infof("%s %s", printYR(true, "%s", time.Unix((*v.Timestamp/1000), 0).Format(timeLayout)), *v.Message)
			}
		}
		return
	case ECSServiceList:
		tbl = table.New("CLUSTER", "SERVICE")
		for k, v := range i {
			for _, s := range v {
				tbl.AddRow(parseArn(k), parseArn(s))
			}
		}
	case *ecr.DescribeRepositoriesOutput:
		tbl = table.New("REPO", "URI")
		for _, v := range i.Repositories {
			tbl.AddRow(*v.RepositoryName, *v.RepositoryUri)
		}
	case []*ecs.Service:
		tbl = table.New("SERVICE", "STATUS", "DESIRED", "RUNNING", "PENDING", "EVENTS")
		for _, v := range i {
			tbl.AddRow(*v.ServiceName, *v.Status, *v.DesiredCount, *v.RunningCount, *v.PendingCount, len(v.Events))
		}
	case map[string][]*ecr.ImageIdentifier:
		switch {
		case showImgTags:
			tbl = table.New("IMAGE", "TAG")
			for k, t := range i {
				for _, tag := range t {
					if tag.ImageTag != nil {
						tbl.AddRow(k, *tag.ImageTag)
					}
				}
			}
		default:
			tbl = table.New("IMAGE", "TAGS")
			for k, v := range i {
				tbl.AddRow(k, len(v))
			}
		}
	case *ecs.DescribeClustersOutput:
		tbl = table.New("CLUSTER", "SERVICES", "TASKS", "PENDING")
		for _, v := range i.Clusters {
			tbl.AddRow(*v.ClusterName, *v.ActiveServicesCount, *v.RunningTasksCount, *v.PendingTasksCount)
		}
	}
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	tbl.Print()
	fmt.Println()
}

func printAwsWide(i interface{}) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	var tbl table.Table
	switch i := i.(type) {
	case []awsgo.Instance:
		tbl = table.New("INDEX", "NAME", "ID", "TYPE", "AZ", "VPC", "STATE", "IP", "HOSTNAME", "KEY", "TAGS")
		for _, v := range i {
			tbl.AddRow(v.Index, v.Name, v.ID, v.Type, v.AZ, v.VPC, v.State, v.PublicIP, v.PublicDNSName, v.KeyName, v.TagCount)
		}
	}
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	tbl.Print()
	fmt.Println()
}

func printAwsLong(i interface{}) {
	switch i := i.(type) {
	case []awsgo.Instance:
		for _, v := range i {
			longPrint(v, v.ID)
		}
	}
}

func longPrint(i interface{}, header string) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	var tbl table.Table
	switch i := i.(type) {
	case awsgo.Instance:
		tbl = table.New("INSTANCEID:", header)
		longMap := convertToLong(i)
		for l := range longMap {
			tbl.AddRow(l, longMap[l])
		}
	}
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	tbl.Print()
	fmt.Println()
}

func convertToLong(i interface{}) map[string]interface{} {
	longMap := make(map[string]interface{})
	j, _ := json.Marshal(i)
	json.Unmarshal(j, &longMap)
	return longMap
}

func parseArn(arn string) string {
	values := strings.Split(arn, `/`)
	if len(values) == 2 {
		return values[1]
	}
	return arn
}

func printYR(yellow bool, format string, a ...interface{}) string {
	var pFmt func(format string, a ...interface{}) string
	if yellow {
		pFmt = color.New(color.FgGreen).SprintfFunc()
	} else {
		pFmt = color.New(color.FgRed).SprintfFunc()
	}
	return pFmt(format, a...)
}
