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

// StartEC2Instances starts AWS Instances by id.
func StartEC2Instances(flags EC2Flags, ids ...string) []awsctl.Instance {
	/*
		defaultFlags := EC2Flags{}.GetDefaults(client.AWSContext().DefaultConfigDir, flags)
		if defaultFlags != nil {
			client.AddConfig(defaultFlags)
		}
		instances := client.GetInstances(ids...)
		if len(instances) < 1 {
			out.Failf("No Results Found.")
		}
	*/
	instances := GetEC2Instances(flags, ids...)
	var insts []*string
	for _, inst := range instances {
		insts = append(insts, &inst.ID)
	}
	return instances
}

// StopEC2Instances stops AWS Instances by id.
func StopEC2Instances(flags EC2Flags, ids ...string) []awsctl.Instance {
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
