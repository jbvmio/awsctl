package awsctl

// Default Config Templates
const (
	EC2File = `ec2_defaults`
)

// ConfigDirectory contains all default config template files
var ConfigDirectory string

// ConfigOptions contains the options that deliver configuration details.
type ConfigOptions interface {
	ConfigRegion() *string
	GetDefaults(defaultDir string, overrides ConfigOptions) ConfigOptions
}

// GetDefault returns the default value for the given key (TODO).
func GetDefault(dir, key string) {

}
