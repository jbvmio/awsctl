package awsctl

import (
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// ServiceType defines a specific AWS Service available for a Client.
type ServiceType uint

// Available Service Types:
const (
	SvcTypeEC2        ServiceType = 1
	SvcTypeCloudWatch ServiceType = 2
)

// SVC contains available AWS service clients
type SVC struct {
	ec2Svc *ec2.EC2
	cwSvc  *cloudwatch.CloudWatch
}
