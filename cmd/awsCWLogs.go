package cmd

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/jbvmio/awsgo"
)

// LogsFlags contains arg options for CloudWatch Logs.
type LogsFlags struct {
	LogGroup string
	Stream   string
	Pages    int
	Last     time.Duration
}

// GetLogStream returns logs from the specified Log Stream.
func GetLogStream(aFlags AWSFlags, lFlags LogsFlags) (events []*cloudwatchlogs.OutputLogEvent) {
	client.AddConfig(awsgo.SvcTypeCWLogs, aFlags)
	input := &cloudwatchlogs.GetLogEventsInput{
		LogGroupName:  &lFlags.LogGroup,
		LogStreamName: &lFlags.Stream,
	}

	if lFlags.Last != time.Duration(0) {
		timeNow := time.Now().UTC()
		timeThen := timeNow.Add(-lFlags.Last).UTC().Unix() * 1000
		input.SetStartTime(timeThen)
	}

	//Infof("%v", lFlags.Last)
	//Exitf(0, ">>>: %v", *input.StartTime)

	pageNum := 0
	err := client.CWLogs().GetLogEventsPages(input, func(output *cloudwatchlogs.GetLogEventsOutput, lastPage bool) bool {
		pageNum++
		events = append(events, output.Events...)
		return pageNum <= lFlags.Pages
	})
	if err != nil {
		Failf("error retrieving logs from stream: %v", err)
	}
	return
}

// ListLogStreams lists the log streams for the specified LogGroup.
func ListLogStreams(aFlags AWSFlags, lFlags LogsFlags) (streams []*cloudwatchlogs.LogStream) {
	client.AddConfig(awsgo.SvcTypeCWLogs, aFlags)
	input := &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: &lFlags.LogGroup,
		OrderBy:      aws.String(`LastEventTime`),
	}
	err := client.CWLogs().DescribeLogStreamsPages(input, func(output *cloudwatchlogs.DescribeLogStreamsOutput, more bool) bool {
		if len(output.LogStreams) < 1 {
			return false
		}
		streams = append(streams, output.LogStreams...)
		return true
	})
	if err != nil {
		Failf("error retrieving log streams: %v", err)
	}
	return
}

// ListLogGroups returns the list of LogGroups.
func ListLogGroups(aFlags AWSFlags) (groups []*cloudwatchlogs.LogGroup) {
	client.AddConfig(awsgo.SvcTypeCWLogs, aFlags)
	output, err := client.CWLogs().DescribeLogGroups(nil)
	if err != nil {
		Failf("error listing LogGroups: %v", err)
	}
	groups = output.LogGroups
	if output.NextToken != nil {
		for output.NextToken != nil {
			input := &cloudwatchlogs.DescribeLogGroupsInput{
				NextToken: output.NextToken,
			}
			output, err = client.CWLogs().DescribeLogGroups(input)
			if err != nil {
				Failf("error using NextToken listing LogGroups: %v", err)
			}
			groups = append(groups, output.LogGroups...)
		}
	}
	return
}
