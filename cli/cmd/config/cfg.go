package config

import (
	"github.com/jbvmio/awsctl"
	"github.com/jbvmio/awsctl/cli/x/out"
	"github.com/spf13/viper"
)

// Config holds all values for a given context.
type Config struct {
	Contexts       map[string]map[string]string `yaml:"contexts"`
	CurrentContext string                       `yaml:"current-context"`
	ConfigVersion  int                          `yaml:"config-version"`
}

func GetConfig() *Config {
	var config Config
	viper.Unmarshal(&config)
	config.CurrentContext = viper.GetString("current-context")
	config.ConfigVersion = viper.GetInt("config-version")
	return &config
}

func GetContextList() map[string][]string {
	config := GetConfig()
	contexts := make(map[string][]string, len(config.Contexts))
	for k := range config.Contexts {
		if k == config.CurrentContext {
			contexts["contexts"] = append(contexts["contexts"], string(k+" [current-context]"))
		} else {
			contexts["contexts"] = append(contexts["contexts"], k)
		}
	}
	return contexts
}

// GetContext returns the configuration for the given context, or the current context if none is specified.
func GetContext(context ...string) *awsctl.AWSContext {
	switch true {
	case len(context) > 1:
		out.Failf("Error: too many contexts specified, only 1 allowed")
	case len(context) < 1:
		return getCurrentCtx()
	case context[0] == "":
		return getCurrentCtx()
	}
	config := GetConfig()
	ctx := config.Contexts[context[0]]
	if ctx["name"] == "" {
		out.Failf("Error: no context named %v", context[0])
	}
	return awsctl.CreateAWSContext(ctx)
}

func getCurrentCtx() *awsctl.AWSContext {
	current := viper.GetString("current-context")
	config := GetConfig()
	ctx := config.Contexts[current]
	if ctx["name"] == "" {
		out.Failf("Error: invalid config or context")
	}
	return awsctl.CreateAWSContext(ctx)
}

func genSample() {
	out.Infof(`
contexts:
  default:
    name: default
    default_config_dir: .aws/configs/default
    aws_access_key_id: accessKeyHere
    aws_secret_access_key: secretAccessKeyHere
    aws_session_token: ""
    aws_provider_name: ""
  dev:
    name: dev
    default_config_dir: .aws/configs/dev
    aws_access_key_id: accessKeyHere
    aws_secret_access_key: secretAccessKeyHere
    aws_session_token: ""
    aws_provider_name: ""
current-context: default
config-version: 1
`)
}
