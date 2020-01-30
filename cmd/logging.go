package cmd


import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tidwall/pretty"
	"gopkg.in/yaml.v2"
)

// Failf .
func Failf(msg string, args ...interface{}) {
	Exitf(1, msg, args...)
}

// Warnf .
func Warnf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
}

// IfErrWarnf .
func IfErrWarnf(err error) {
	if err != nil {
		Warnf("Error %v", err)
	}
}

// Infof .
func Infof(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, msg+"\n", args...)
}

// Exitf .
func Exitf(code int, msg string, args ...interface{}) {
	if code == 0 {
		fmt.Fprintf(os.Stdout, msg+"\n", args...)
	} else {
		fmt.Fprintf(os.Stderr, msg+"\n", args...)
	}
	os.Exit(code)
}

// Marshal .
func Marshal(object interface{}, format string) error {
	var fmtString []byte
	var err error
	if format == "yaml" {
		fmtString, err = yaml.Marshal(object)
		if err != nil {
			err = fmt.Errorf("unable to format yaml: %v", err)
		}
		fmt.Println(string(fmtString))
	} else if format == "json" {
		fmtString, err = json.Marshal(object)
		if err != nil {
			err = fmt.Errorf("unable to format json: %v", err)
		}
		fmt.Printf("%s", pretty.Pretty(fmtString))
	} else {
		err = fmt.Errorf("unknown format: %v", format)
	}
	return err
}

// PrintStrings .
func PrintStrings(args ...string) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}
