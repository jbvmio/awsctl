package awsctl

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// InstanceMap hold a mapping of ec2 instances by instance id.
type InstanceMap map[string]*ec2.Instance

func (i InstanceMap) ListIDs() []string {
	var ids []string
	for k := range i {
		ids = append(ids, k)
	}
	return ids
}

func (i InstanceMap) ListSG() []string {
	var ids []string
	for k := range i {
		for _, sg := range i[k].SecurityGroups {
			ids = append(ids, *sg.GroupId)
		}
	}
	return ids
}

// GetInstances returns an InstanceMap based on the entered id string. All instances returned if no id entered.
func (cl *Client) GetInstanceMap(ids ...string) InstanceMap {
	var instanceMap InstanceMap
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
	if err == nil {
		instanceMap = make(map[string]*ec2.Instance)
		for _, res := range Insts.Reservations {
			for _, inst := range res.Instances {
				instanceMap[*inst.InstanceId] = inst
			}
		}
	}
	return instanceMap
}
