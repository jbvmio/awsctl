package main

import (
	"fmt"
	"local/jbvmio/awsctl"
)

func main() {
	client, _ := awsctl.NewClient("us-east-2")
	instances := client.GetInstances()
	fmt.Printf("%+v", instances)
}
