package awsgo

import (

	"github.com/jbvmio/awsctl"
	"github.com/jbvmio/awsctl/cli/x/out"
)

// GetEC2Instances returns AWS Instances
func GetEC2Instances(flags EC2Flags, ids ...string) []awsctl.Instance {
	defaultFlags := EC2Flags.GetDefaults(EC2Flags{}, client.AWSContext().DefaultConfigDir, flags)
	if defaultFlags != nil {
		client.AddConfig(defaultFlags)
	}
	//client.AddConfig(flags)
	instances := client.GetInstances(ids...)
	if len(instances) < 1 {
		out.Failf("No Results Found.")
	}
	return instances
}
