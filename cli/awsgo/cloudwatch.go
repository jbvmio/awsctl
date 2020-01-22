package awsgo

import (
	"github.com/jbvmio/awsctl/cli/x/out"
)

// ListMetrics returns CloudWatch Metrics.
func ListMetrics(flags EC2Flags) {
	defaultFlags := EC2Flags{}.GetDefaults(client.AWSContext().DefaultConfigDir, flags)
	if defaultFlags != nil {
		client.AddConfig(defaultFlags)
	}
	output, err := client.CW().ListMetrics(nil)
	if err != nil {
		out.Failf("error listing metrics: %v", err)
	}
}
