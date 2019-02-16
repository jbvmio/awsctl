package awsctl

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/credentials"
)

// AWSContext contains AWS credentials and details for a given context.
type AWSContext struct {
	Name             string `yaml:"name"`
	DefaultConfigDir string `yaml:"default_config_dir"`
	AccessKeyID      string `yaml:"aws_access_key_id"`
	SecretAccessKey  string `yaml:"aws_secret_access_key"`
	SessionToken     string `yaml:"aws_session_token"`
	ProviderName     string `yaml:"aws_provider_name"`
}

// Value holds AWS specific values.
type Value struct {
	AccessKeyID     string `yaml:"aws_access_key_id"`
	SecretAccessKey string `yaml:"aws_secret_access_key"`
	SessionToken    string `yaml:"aws_session_token"`
	ProviderName    string `yaml:"aws_provider_name"`
}

// CreateAWSContext constructs an AWSContext from a map[string]string containing the AWSContext values.
func CreateAWSContext(values map[string]string) *AWSContext {
	return &AWSContext{
		Name:             values["name"],
		DefaultConfigDir: values["default_config_dir"],
		AccessKeyID:      values["aws_access_key_id"],
		SecretAccessKey:  values["aws_secret_access_key"],
		SessionToken:     values["aws_session_token"],
		ProviderName:     values["aws_provider_name"],
	}
}

// Retrieve return AWS Credential Values
func (ctx *AWSContext) Retrieve() (credentials.Value, error) {
	switch {
	case ctx.AccessKeyID == "" || ctx.SecretAccessKey == "":
		return credentials.Value{}, fmt.Errorf("Error: Empty Context.")
	}
	return credentials.Value{
		AccessKeyID:     ctx.AccessKeyID,
		SecretAccessKey: ctx.SecretAccessKey,
		SessionToken:    ctx.SessionToken,
		ProviderName:    ctx.ProviderName,
	}, nil
}

func (ctx *AWSContext) IsExpired() bool {
	return false
}
