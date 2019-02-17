package awsctl

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Instance holds metadata of an ec2 instance.
type Instance struct {
	AZ             string
	ID             string
	Image          string
	Index          int64
	KeyName        string
	Name           string
	PrivateDnsName string
	PrivateIP      string
	PublicDnsName  string
	PublicIP       string
	State          string
	Type           string
	VPC            string
	Tags           map[string]string
	TagCount       int
}

func (cl *Client) GetInstances(ids ...string) []Instance {
	return cl.GetInstanceMap().GetInstances(ids...)
}

func (i InstanceMap) GetInstances(ids ...string) []Instance {
	var instances []Instance
	switch {
	case len(ids) > 0:
		for _, id := range ids {
			if i[id] != nil {
				var name string
				var tags map[string]string
				if i[id].Tags != nil {
					tags = make(map[string]string, len(i[id].Tags))
					for _, tag := range i[id].Tags {
						tags[*tag.Key] = *tag.Value
					}
					name = tags["Name"]
				}
				inst := Instance{
					AZ:             *i[id].Placement.AvailabilityZone,
					ID:             *i[id].InstanceId,
					Image:          *i[id].ImageId,
					Index:          *i[id].AmiLaunchIndex,
					KeyName:        *i[id].KeyName,
					Name:           name,
					PrivateDnsName: *i[id].PrivateDnsName,
					PrivateIP:      *i[id].PrivateIpAddress,
					PublicDnsName:  *i[id].PublicDnsName,
					PublicIP:       *i[id].PublicIpAddress,
					State:          *i[id].State.Name,
					Type:           *i[id].InstanceType,
					VPC:            *i[id].VpcId,
					Tags:           tags,
					TagCount:       len(tags),
				}
				instances = append(instances, inst)
			}
		}
	default:
		for id := range i {
			var name string
			var tags map[string]string
			if i[id].Tags != nil {
				tags = make(map[string]string, len(i[id].Tags))
				for _, tag := range i[id].Tags {
					tags[*tag.Key] = *tag.Value
				}
				name = tags["Name"]
			}
			inst := Instance{
				AZ:             *i[id].Placement.AvailabilityZone,
				ID:             *i[id].InstanceId,
				Image:          *i[id].ImageId,
				Index:          *i[id].AmiLaunchIndex,
				KeyName:        *i[id].KeyName,
				Name:           name,
				PrivateDnsName: *i[id].PrivateDnsName,
				PrivateIP:      *i[id].PrivateIpAddress,
				PublicDnsName:  *i[id].PublicDnsName,
				PublicIP:       *i[id].PublicIpAddress,
				State:          *i[id].State.Name,
				Type:           *i[id].InstanceType,
				VPC:            *i[id].VpcId,
				Tags:           tags,
				TagCount:       len(tags),
			}
			instances = append(instances, inst)
		}
	}
	return instances
}

// StartEC2Instances starts AWS Instances by id.
func (cl *Client) StartEC2Instances(ids ...string) []awsctl.Instance {
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
	var instances []*string
	for _, i := range ids {
		instances = append(instances, &i)
	}
	input := &ec2.StartInstancesInput{
		InstanceIds: []*string{instances},
		DryRun: cl.dryrunMode,
	}
	result, err := cl.EC2().StartInstances(input)
	if aerr, ok := err.(awserr.Error); ok {
		fmt.Printf("Error: %v\n%v\n", aerr.Message(), err)
	}
	if err == nil {

	}



	return instances
}