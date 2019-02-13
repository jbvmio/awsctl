package main

import (
	"local/jbvmio/awsctl"
)

func main() {
	client, _ := awsctl.NewClient("us-east-2")
	client.GetInstances()
}
