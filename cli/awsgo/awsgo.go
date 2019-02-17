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

// GlobalFlags holds Global options.
type GlobalFlags struct {
	DryRun bool
}

// LaunchAWSClient launches the AWS Client
func LaunchAWSClient(context *awsctl.AWSContext, flags GlobalFlags) {
	client, errd = awsctl.NewClient(context)
	if errd != nil {
		out.Failf("Error Launching AWS Client: %v", errd)
	}
	clientContext = context
}
