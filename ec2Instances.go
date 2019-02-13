package awsctl

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// GetInstances retrieves instances based on the entered id string. All instances returned if no id entered.
func (cl *Client) GetInstances(ids ...string) {
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
	instances, err := cl.EC2().DescribeInstances(input)
	if aerr, ok := err.(awserr.Error); ok {
		fmt.Println("Error:", aerr.Message())
	}
	if err == nil {
		fmt.Println(instances)
	}
}
