package cmd

import (
	"io/ioutil"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/jbvmio/awsgo"
	"gopkg.in/yaml.v2"
)

type overrides struct {
	region string
}

// AWSFlags contains flag options for EC2.
type AWSFlags struct {
	Region    *string `yaml:"Region"`
	ctx       *awsgo.AWSContext
	overrides overrides
}

// GetDefaults returns default configurations defined with overrides provided by flags.
func (flags AWSFlags) GetDefaults(svcType awsgo.ServiceType) *aws.Config {
	flags.ctx = client.AWSContext()
	var conf aws.Config
	var defaultFlags AWSFlags
	defaultDir := flags.ctx.DefaultConfigDir
	var path string
	if fileExists(defaultDir) {
		ok, err := regexp.MatchString(`.*/$`, defaultDir)
		if err == nil {
			if !ok {
				defaultDir = defaultDir + "/"
			}
		}
		switch {
		case fileExists(defaultDir + svcType.DefaultConfig()):
			path = defaultDir + svcType.DefaultConfig()
		default:
			path = defaultDir + awsgo.DefaultConfigName
		}
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		Failf("%v", err)
	}
	err = yaml.Unmarshal(data, &defaultFlags)
	if err != nil {
		Failf("%v", err)
	}
	conf.Region = defaultFlags.Region
	if flags.overrides.region != "" {
		conf.Region = &flags.overrides.region
	}
	return &conf
}
