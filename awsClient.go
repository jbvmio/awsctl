package awsctl

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var initErr error

// Client makes calls into AWS
type Client struct {
	config     *aws.Config
	session    *session.Session
	dryrunMode bool
	svc        *SVC
	awsContext *AWSContext
}

// SVC contains available AWS service clients
type SVC struct {
	ec2Svc *ec2.EC2
}

// NewClient creates a new Client
func NewClient(awsContext *AWSContext) (*Client, error) {
	var client Client
	creds, err := awsContext.Retrieve()
	if err != nil {
		return &client, err
	}
	awsConfig := aws.Config{
		Credentials: credentials.NewStaticCredentialsFromCreds(creds),
	}
	sess, err := session.NewSession(&awsConfig)
	if err != nil {
		return &client, err
	}
	client.awsContext = awsContext
	client.session = sess //sess.Copy()
	client.svc = &SVC{}
	return &client, nil
}

// AWSContext returns the client's AWSContext.
func (cl *Client) AWSContext() *AWSContext {
	return cl.awsContext
}

// AddConfig changes the underlying session with new Config options.
func (cl *Client) AddConfig(options ConfigOptions) {
	cl.session = cl.session.Copy(&aws.Config{
		Region: options.ConfigRegion(),
	})
}

// DryRunMode sets the DryRun bool
func (cl *Client) DryRunMode(enabled bool) {
	cl.dryrunMode = enabled
}

func (cl *Client) InitEC2() {
	cl.svc.ec2Svc = ec2.New(cl.session)
}

func (cl *Client) EC2() *ec2.EC2 {
	if cl.svc.ec2Svc == nil {
		cl.InitEC2()
	}
	return cl.svc.ec2Svc
}
