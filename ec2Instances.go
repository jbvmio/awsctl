package awsctl

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

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

// To replace in above ^.
/*
func makeThisSoonJB() {
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
*/

// StartEC2Instances starts AWS Instances by id.
func (cl *Client) StartEC2Instances(ids ...string) []InstanceStateChange {
	var stateChanges []InstanceStateChange //*ec2.StartInstancesOutput
	var instances []*string
	for _, i := range ids {
		instances = append(instances, &i)
	}
	input := &ec2.StartInstancesInput{
		InstanceIds: instances,
		DryRun:      aws.Bool(cl.dryrunMode),
	}
	output, err := cl.EC2().StartInstances(input)
	if aerr, ok := err.(awserr.Error); ok {
		fmt.Printf("Error: %v\n%v\n", aerr.Message(), err)
	}
	if err == nil {
		for _, out := range output.StartingInstances {
			var isc InstanceStateChange
			isc.convertFrom(out)
			stateChanges = append(stateChanges, isc)
		}
	}
	return stateChanges
}

// StopEC2Instances starts AWS Instances by id.
func (cl *Client) StopEC2Instances(ids ...string) []InstanceStateChange {
	var stateChanges []InstanceStateChange //*ec2.StartInstancesOutput
	var instances []*string
	for _, i := range ids {
		instances = append(instances, &i)
	}
	input := &ec2.StopInstancesInput{
		InstanceIds: instances,
		DryRun:      aws.Bool(cl.dryrunMode),
	}
	output, err := cl.EC2().StopInstances(input)
	if aerr, ok := err.(awserr.Error); ok {
		fmt.Printf("Error: %v\n%v\n", aerr.Message(), err)
	}
	if err == nil {
		for _, out := range output.StoppingInstances {
			var isc InstanceStateChange
			isc.convertFrom(out)
			stateChanges = append(stateChanges, isc)
		}
	}
	return stateChanges
}
