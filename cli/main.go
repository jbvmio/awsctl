package main

import (
	"fmt"
	"local/jbvmio/awsctl"
)

func main() {
	client, _ := awsctl.NewClient("us-east-2")
	//instMap := client.GetEC2Instances()
	instances := client.GetInstances()
	fmt.Printf("%+v", instances)
}
