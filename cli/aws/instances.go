package aws

import (
	"github.com/jbvmio/awsctl"
)

func GetEC2Instances(ids ...string) []awsctl.Instance {
	client, _ := awsctl.NewClient("us-east-2")
	instances := client.GetInstances(ids...)
	return instances
}
