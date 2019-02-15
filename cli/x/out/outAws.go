package out

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
	"github.com/jbvmio/awsctl"
	"github.com/rodaine/table"
)

func PrintAWS(i interface{}, format ...string) {
	var f string
	switch {
	case len(format) > 0:
		f = format[0]
	default:
		printAws(i)
		return
	}
	switch f {
	case "wide":
		printAwsWide(i)
	case "long":
		printAwsLong(i)
	default:
		IfErrWarnf(Marshal(i, f))
	}
}

func printAws(i interface{}) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	var tbl table.Table
	switch i := i.(type) {
	case []awsctl.Instance:
		tbl = table.New("INDEX", "ID", "STATE", "IP", "HOSTNAME")
		for _, v := range i {
			tbl.AddRow(v.Index, v.ID, v.State, v.PublicIP, v.PublicDnsName)
		}
	}
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	tbl.Print()
	fmt.Println()
}

func printAwsWide(i interface{}) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	var tbl table.Table
	switch i := i.(type) {
	case []awsctl.Instance:
		tbl = table.New("INDEX", "NAME", "ID", "TYPE", "AZ", "VPC", "STATE", "IP", "HOSTNAME", "KEY", "TAGS")
		for _, v := range i {
			tbl.AddRow(v.Index, v.Name, v.ID, v.Type, v.AZ, v.VPC, v.State, v.PublicIP, v.PublicDnsName, v.KeyName, v.TagCount)
		}
	}
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	tbl.Print()
	fmt.Println()
}

func printAwsLong(i interface{}) {
	switch i := i.(type) {
	case []awsctl.Instance:
		for _, v := range i {
			longPrint(v, v.ID)
		}
	}
}

func longPrint(i interface{}, header string) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	var tbl table.Table
	switch i := i.(type) {
	case awsctl.Instance:
		tbl = table.New("INSTANCEID:", header)
		longMap := convertToLong(i)
		for l := range longMap {
			tbl.AddRow(l, longMap[l])
		}
	}
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	tbl.Print()
	fmt.Println()
}

func convertToLong(i interface{}) map[string]interface{} {
	longMap := make(map[string]interface{})
	j, _ := json.Marshal(i)
	json.Unmarshal(j, &longMap)
	return longMap
}
