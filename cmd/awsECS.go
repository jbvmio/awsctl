package cmd

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/jbvmio/awsgo"
)

// ECSFlags defines arguments for ECS Ops.
type ECSFlags struct {
	Cluster string
	Service string
}

// ECSServiceList maps cluster names with services.
type ECSServiceList map[string][]string

// DescribeECSServices returns ECS Service details for a ECS Cluster.
func DescribeECSServices(services ECSServiceList) []*ecs.Service {
	var S []*ecs.Service
	for k, v := range services {
		var svc []*string
		for _, s := range v {
			//x := parseArn(s)
			//Infof(">>: %s", x)
			svc = append(svc, &s)
		}
		input := &ecs.DescribeServicesInput{
			Cluster:  aws.String(k),
			Services: svc,
		}
		output, err := client.ECS().DescribeServices(input)
		if err != nil {
			Failf("error retieving ECS services: %v", err)
		}
		S = append(S, output.Services...)
	}
	return S
}

// ListECSServices lists available ECS Services.
func ListECSServices(aFlags AWSFlags, eFlags ECSFlags) ECSServiceList {
	var clusters []*string
	services := make(map[string][]string)
	switch {
	case eFlags.Cluster != "":
		client.AddConfig(awsgo.SvcTypeECS, aFlags)
		clusters = []*string{
			aws.String(eFlags.Cluster),
		}
	default:
		C := ListECSClusters(aFlags)
		for _, c := range C.ClusterArns {
			clusters = append(clusters, c)
		}
	}
	for _, c := range clusters {
		output, err := client.ECS().ListServices(&ecs.ListServicesInput{
			Cluster: c,
		})
		if err != nil {
			Failf("error listing ECS services: %v", err)
		}
		for _, s := range output.ServiceArns {
			services[*c] = append(services[*c], *s)
		}
	}
	if eFlags.Service != "" {
		for k, v := range services {
			var tmp []string
			for _, s := range v {
				if strings.Contains(s, eFlags.Service) {
					tmp = append(tmp, s)
				}
			}
			switch {
			case len(tmp) < 1:
				delete(services, k)
			default:
				services[k] = tmp
			}
		}
	}
	if len(services) < 1 {
		Exitf(0, "no results found.")
	}
	return services
}

// DescribeClusters returns Cluster details for ECS.
func DescribeClusters(aFlags AWSFlags, eFlags ECSFlags) *ecs.DescribeClustersOutput {
	var clusters []*string

	switch {
	case eFlags.Cluster != "":
		client.AddConfig(awsgo.SvcTypeECS, aFlags)
		clusters = []*string{
			aws.String(eFlags.Cluster),
		}
	default:
		C := ListECSClusters(aFlags)
		for _, c := range C.ClusterArns {
			clusters = append(clusters, c)
		}
	}
	output, err := client.ECS().DescribeClusters(&ecs.DescribeClustersInput{
		Clusters: clusters,
		Include: []*string{
			aws.String(`ATTACHMENTS`),
			aws.String(`SETTINGS`),
			aws.String(`STATISTICS`),
			aws.String(`TAGS`),
		},
	})
	if err != nil {
		Failf("error listing ECS clusters: %v", err)
	}
	return output
}

// ListECSClusters lists available ECS Clusters.
func ListECSClusters(aFlags AWSFlags) *ecs.ListClustersOutput {
	client.AddConfig(awsgo.SvcTypeECS, aFlags)
	output, err := client.ECS().ListClusters(nil)
	if err != nil {
		Failf("error listing ECS clusters: %v", err)
	}
	return output
}
