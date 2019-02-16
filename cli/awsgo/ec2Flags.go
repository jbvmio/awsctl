package awsgo

import (
	"fmt"
	"io/ioutil"
	"regexp"

	yaml "gopkg.in/yaml.v2"

	"github.com/jbvmio/awsctl"
	"github.com/jbvmio/awsctl/cli/x/ops"
)

type EC2Flags struct {
	Region string `yaml:"Region"`
}

func (flags EC2Flags) ConfigRegion() *string {
	return &flags.Region
}

func (flags EC2Flags) GetDefaults(defaultDir string, overrides awsctl.ConfigOptions) awsctl.ConfigOptions {
	var defaultFlags EC2Flags
	if ops.FileExists(defaultDir) {
		var path string
		ok, err := regexp.MatchString(`.*/$`, defaultDir)
		if err == nil {
			switch {
			case !ok:
				path = defaultDir + "/" + awsctl.EC2File
			default:
				path = defaultDir + awsctl.EC2File
			}
			data, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println(err)
				return &defaultFlags
			}
			err = yaml.Unmarshal(data, &defaultFlags)
			if err != nil {
				fmt.Println(err)
				return &defaultFlags
			}
			switch {
			case *overrides.ConfigRegion() != "":
				defaultFlags.Region = *overrides.ConfigRegion()
			}
		}
	}
	return &defaultFlags
}