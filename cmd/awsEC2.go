package cmd

import (
	"github.com/jbvmio/awsgo"
)

// GetEC2Instances returns AWS Instances by id.
func GetEC2Instances(flags AWSFlags, ids ...string) []awsgo.Instance {
	client.AddConfig(awsgo.SvcTypeEC2, flags)
	instances := client.GetInstances(ids...)
	if len(instances) < 1 {
		Failf("No Results Found.")
	}
	return instances
}

// StartEC2Instances stops AWS Instances by id.
func StartEC2Instances(flags AWSFlags, ids ...string) []awsgo.InstanceStateChange {
	client.AddConfig(awsgo.SvcTypeEC2, flags)
	output := client.StartEC2Instances(ids...)
	if len(output) < 1 {
		Failf("No Results Found.")
	}
	return output
}

// StopEC2Instances stops AWS Instances by id.
func StopEC2Instances(flags AWSFlags, ids ...string) []awsgo.InstanceStateChange {
	client.AddConfig(awsgo.SvcTypeEC2, flags)
	output := client.StopEC2Instances(ids...)
	if len(output) < 1 {
		Failf("No Results Found.")
	}
	return output
}

// RebootEC2Instances reboots AWS Instances by id.
func RebootEC2Instances(flags AWSFlags, ids ...string) { //[]awsgo.InstanceStateChange {
	client.AddConfig(awsgo.SvcTypeEC2, flags)
	if ok := client.RebootEC2Instances(ids...); ok {
		Infof("Reboot Successfully Sent.")
	}
}
