package cmd

import (
	"context"
	"encoding/base64"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/docker/cli/cli/config"
	"github.com/docker/cli/cli/config/credentials"
	ctypes "github.com/docker/cli/cli/config/types"
	"github.com/docker/docker/api/types"
	dclient "github.com/docker/docker/client"
	"github.com/jbvmio/awsgo"
)

// LoginAWSECR - Currently not being used?
func LoginAWSECR(aFlags AWSFlags) {
	client.AddConfig(awsgo.SvcTypeECR, aFlags)
	input := ecr.GetAuthorizationTokenInput{
		RegistryIds: []*string{
			aws.String(`0007778889997`),
		},
	}
	token, err := client.ECR().GetAuthorizationToken(&input)
	if err != nil {
		Failf("error generating token: %v", err)
	}
	if len(token.AuthorizationData) < 1 {
		Exitf(0, "no data found")
	}
	u, p := parseTokenAuthData(*token.AuthorizationData[0].AuthorizationToken)
	c, err := dclient.NewEnvClient()
	if err != nil {
		Failf("is the Docker Daemon running? Error: %s\n", err.Error())
	}
	authConf := types.AuthConfig{
		Username:      u,
		Password:      p,
		ServerAddress: *token.AuthorizationData[0].ProxyEndpoint,
	}
	resp, err := c.RegistryLogin(context.Background(), authConf)
	if err != nil {
		Failf("error logging into registry: %v\n", err)
	}
	dir := config.Dir()
	conf, err := config.Load(dir)
	if err != nil {
		Failf("error obtaining config dir: %v\n", err)
	}
	cAuthConf := ctypes.AuthConfig{
		Username:      u,
		Password:      p,
		ServerAddress: *token.AuthorizationData[0].ProxyEndpoint,
	}
	store := credentials.NewNativeStore(conf, `desktop`)
	err = store.Store(cAuthConf)
	if err != nil {
		Failf("error saving auth config: %v\n", err)
	}
	Infof("%s", resp.Status)
}

// GenECRToken generates an Auth token for pushing to ECR.
func GenECRToken(aFlags AWSFlags) *ecr.GetAuthorizationTokenOutput {
	client.AddConfig(awsgo.SvcTypeECR, aFlags)
	output, err := client.ECR().GetAuthorizationToken(nil)
	if err != nil {
		Failf("error generating token: %v", err)
	}
	return output
}

// LoginECR generates a token and performs a docker login.
func LoginECR(aFlags AWSFlags) {
	token := GenECRToken(aFlags)
	if len(token.AuthorizationData) < 1 {
		Exitf(0, "no data found")
	}
	u, p := parseTokenAuthData(*token.AuthorizationData[0].AuthorizationToken)
	c, err := dclient.NewEnvClient()
	if err != nil {
		Failf("is the Docker Daemon running? Error: %s\n", err.Error())
	}
	authConf := types.AuthConfig{
		Username:      u,
		Password:      p,
		ServerAddress: *token.AuthorizationData[0].ProxyEndpoint,
	}
	resp, err := c.RegistryLogin(context.Background(), authConf)
	if err != nil {
		Failf("error logging into registry: %v\n", err)
	}
	dir := config.Dir()
	conf, err := config.Load(dir)
	if err != nil {
		Failf("error obtaining config dir: %v\n", err)
	}
	cAuthConf := ctypes.AuthConfig{
		Username:      u,
		Password:      p,
		ServerAddress: *token.AuthorizationData[0].ProxyEndpoint,
	}
	store := credentials.NewNativeStore(conf, `desktop`)
	err = store.Store(cAuthConf)
	if err != nil {
		Failf("error saving auth config: %v\n", err)
	}
	Infof("%s", resp.Status)
}

func parseTokenAuthData(authToken string) (username, password string) {
	b, err := base64.StdEncoding.DecodeString(authToken)
	if err != nil {
		Failf("error parsing auth token: %v", err)
	}
	up := strings.Split(string(b), `:`)
	switch {
	case len(up) != 2:
		Failf("error decoding correct login parameters")
	default:
		username = up[0]
		password = up[1]
	}
	return
}

// ListImages returns a list of Images for the default registry for the account.
func ListImages(aFlags AWSFlags) map[string][]*ecr.ImageIdentifier {
	imgTags := make(map[string][]*ecr.ImageIdentifier)
	repos := ListRepos(aFlags)
	for _, repo := range repos.Repositories {
		img, err := client.ECR().ListImages(&ecr.ListImagesInput{RepositoryName: repo.RepositoryName})
		if err != nil {
			Failf("error obtaining images for repo %s: %v", repo.RepositoryName, err)
		}
		imgTags[*repo.RepositoryUri] = img.ImageIds
	}
	return imgTags
}

// ListRepos lists all repositories in the default registry for the account.
func ListRepos(aFlags AWSFlags) *ecr.DescribeRepositoriesOutput {
	client.AddConfig(awsgo.SvcTypeECR, aFlags)
	output, err := client.ECR().DescribeRepositories(nil)
	if err != nil {
		Failf("error listing repositories: %v", err)
	}
	Infof("%+v\n", output)
	return output
}

// TODO:
/*
https://github.com/moby/moby/issues/33429
*/
