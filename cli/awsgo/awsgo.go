package awsgo

import (
	"github.com/jbvmio/awsctl"
	"github.com/jbvmio/awsctl/cli/x/out"
)

// package variables
var (
	exact bool
)

// client variables
var (
	client        *awsctl.Client
	errd          error
	clientContext *awsctl.AWSContext
)

// LaunchAWSClient launches the AWS Client
func LaunchAWSClient(context *awsctl.AWSContext) {
	client, errd = awsctl.NewClient(context)
	if errd != nil {
		out.Failf("Error Launching AWS Client: %v", errd)
	}
	clientContext = context
}
