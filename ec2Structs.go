package awsctl

import (
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

func (i *Instance) convertFrom(awsI *ec2.Instance) {
	var name string
	var tags map[string]string
	if awsI.Tags != nil {
		tags = make(map[string]string, len(awsI.Tags))
		for _, tag := range awsI.Tags {
			tags[*tag.Key] = *tag.Value
		}
		name = tags["Name"]
	}

	i.AZ = *awsI.Placement.AvailabilityZone
	i.ID = *awsI.InstanceId
	i.Image = *awsI.ImageId
	i.Index = *awsI.AmiLaunchIndex
	i.KeyName = *awsI.KeyName
	i.Name = name
	i.PrivateDnsName = *awsI.PrivateDnsName
	i.PrivateIP = *awsI.PrivateIpAddress
	i.PublicDnsName = *awsI.PublicDnsName
	i.PublicIP = *awsI.PublicIpAddress
	i.State = *awsI.State.Name
	i.Type = *awsI.InstanceType
	i.VPC = *awsI.VpcId
	i.Tags = tags
	i.TagCount = len(tags)

}

// InstanceStateChange holds state changes for an ec2 instance.
type InstanceStateChange struct {
	ID            string
	CurrentCode   int64
	CurrentState  string
	PreviousCode  int64
	PreviousState string
}

func (isc *InstanceStateChange) convertFrom(awsISC *ec2.InstanceStateChange) {
	isc.ID = *awsISC.InstanceId
	isc.CurrentCode = *awsISC.CurrentState.Code
	isc.CurrentState = *awsISC.CurrentState.Name
	isc.PreviousCode = *awsISC.PreviousState.Code
	isc.PreviousState = *awsISC.PreviousState.Name
}
