package cmd

import (
	"strings"

	"github.com/jbvmio/awsgo"
	"github.com/spf13/viper"
)

// Config holds all values for a given context.
type Config struct {
	Contexts       map[string]map[string]string `yaml:"profiles"`
	CurrentContext string                       `yaml:"current-profile"`
	ConfigVersion  int                          `yaml:"config-version"`
}

// GetConfig .
func GetConfig() *Config {
	var config Config
	config.Contexts = make(map[string]map[string]string)
	//viper.Unmarshal(&config)
	profiles := viper.GetStringMap(`profiles`)
	for profile, values := range profiles {
		config.Contexts[profile] = make(map[string]string)
		config.Contexts[profile][`name`] = profile
		config.Contexts[profile][`default_config_dir`] = defaultConfigDir
		if vals, ok := values.(map[string]interface{}); ok {
			for k, val := range vals {
				if v, good := val.(string); good {
					config.Contexts[profile][k] = v
				}
			}
		}
	}
	config.CurrentContext = viper.GetString("current-profile")
	config.ConfigVersion = viper.GetInt("config-version")
	return &config
}

// GetContextList .
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
func GetContext(context ...string) *awsgo.AWSContext {
	switch true {
	case len(context) > 1:
		Failf("Error: too many contexts specified, only 1 allowed")
	case len(context) < 1:
		return getCurrentCtx()
	case context[0] == "":
		return getCurrentCtx()
	}
	config := GetConfig()
	ctx := config.Contexts[context[0]]
	if ctx["name"] == "" {
		Failf("Error: no context named %v", context[0])
	}
	return awsgo.CreateAWSContext(ctx)
}

func getCurrentCtx() *awsgo.AWSContext {
	current := viper.GetString("current-profile")
	current = strings.ToLower(current)
	config := GetConfig()
	ctx := config.Contexts[current]
	if ctx["name"] == "" {
		Failf("Error: invalid config or profile")
	}
	return awsgo.CreateAWSContext(ctx)
}

func genSample() {
	Infof(`
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
