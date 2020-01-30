package cmd

import "github.com/jbvmio/awsgo"

// package variables
var (
	exact bool
)

// client variables
var (
	client        *awsgo.Client
	errd          error
	clientContext *awsgo.AWSContext
)

// GlobalFlags holds Global options.
type GlobalFlags struct {
	DryRun bool
}

// LaunchAWSClient launches the AWS Client
func LaunchAWSClient(context *awsgo.AWSContext, flags GlobalFlags) {
	client, errd = awsgo.NewClient(context)
	if errd != nil {
		Failf("Error Launching AWS Client: %v", errd)
	}
	client.DryRunMode(flags.DryRun)
	clientContext = context
}
