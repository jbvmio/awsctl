package awsctl

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Instances hold a mapping of ec2 instances by instance id.
type Instances map[string]*ec2.Instance

// GetInstances retrieves instances based on the entered id string. All instances returned if no id entered.
func (cl *Client) GetInstances(ids ...string) *Instances {
	var input *ec2.DescribeInstancesInput
	switch {
	case len(ids) > 0:
		var targets []*string
		for _, id := range ids {
			targets = append(targets, aws.String(id))
		}
		input = &ec2.DescribeInstancesInput{
			DryRun: aws.Bool(cl.dryrunMode),
			Filters: []*ec2.Filter{
				&ec2.Filter{
					Name:   aws.String("instance-id"),
					Values: targets,
				},
			},
			InstanceIds: targets,
		}
		//MaxResults: aws.Int64(6),
		//NextToken:  aws.String("String"),
	default:
		input = nil
	}
	Insts, err := cl.EC2().DescribeInstances(input)
	if aerr, ok := err.(awserr.Error); ok {
		fmt.Println("Error:", aerr.Message())
	}
	var instances Instances = make(map[string]*ec2.Instance)
	if err == nil {
		for _, res := range Insts.Reservations {
			for _, inst := range res.Instances {
				instances[*inst.InstanceId] = inst
			}
		}
		//fmt.Printf("%s", instances.Reservations)
	}
	return &instances
}
