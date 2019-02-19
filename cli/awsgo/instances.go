package awsgo

import (
	"github.com/jbvmio/awsctl"
	"github.com/jbvmio/awsctl/cli/x/out"
)

// GetEC2Instances returns AWS Instances by id.
func GetEC2Instances(flags EC2Flags, ids ...string) []awsctl.Instance {
	defaultFlags := EC2Flags{}.GetDefaults(client.AWSContext().DefaultConfigDir, flags)
	if defaultFlags != nil {
		client.AddConfig(defaultFlags)
	}
	instances := client.GetInstances(ids...)
	if len(instances) < 1 {
		out.Failf("No Results Found.")
	}
	return instances
}

// StartEC2Instances stops AWS Instances by id.
func StartEC2Instances(flags EC2Flags, ids ...string) []awsctl.InstanceStateChange {
	defaultFlags := EC2Flags{}.GetDefaults(client.AWSContext().DefaultConfigDir, flags)
	if defaultFlags != nil {
		client.AddConfig(defaultFlags)
	}
	output := client.StartEC2Instances(ids...)
	if len(output) < 1 {
		out.Failf("No Results Found.")
	}
	return output
}

// StopEC2Instances stops AWS Instances by id.
func StopEC2Instances(flags EC2Flags, ids ...string) []awsctl.InstanceStateChange {
	defaultFlags := EC2Flags{}.GetDefaults(client.AWSContext().DefaultConfigDir, flags)
	if defaultFlags != nil {
		client.AddConfig(defaultFlags)
	}
	output := client.StopEC2Instances(ids...)
	if len(output) < 1 {
		out.Failf("No Results Found.")
	}
	return output
}

// RebootEC2Instances reboots AWS Instances by id.
func RebootEC2Instances(flags EC2Flags, ids ...string) { //[]awsctl.InstanceStateChange {
	defaultFlags := EC2Flags{}.GetDefaults(client.AWSContext().DefaultConfigDir, flags)
	if defaultFlags != nil {
		client.AddConfig(defaultFlags)
	}
	if ok := client.RebootEC2Instances(ids...); ok {
		out.Infof("Reboot Successfully Sent.")
	}
}
