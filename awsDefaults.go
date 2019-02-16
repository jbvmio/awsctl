package awsctl

// Default Config Templates
const (
	EC2File = `ec2_defaults`
)

// ConfigDirectory contains all default config template files
var ConfigDirectory string

type ConfigOptions interface {
	ConfigRegion() *string
	GetDefaults(defaultDir string, overrides ConfigOptions) ConfigOptions
}

func GetDefault(dir, key string) {

}
