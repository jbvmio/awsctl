package main

import (
	"fmt"
	"local/jbvmio/awsctl"
)

func main() {
	client, _ := awsctl.NewClient("us-east-2")
	//instMap := client.GetEC2Instances()
	instances := client.GetInstances("i-01df69b0a6b87d929", "i-01df69b0a6b87d929")
	fmt.Printf("%+v", instances)
}
