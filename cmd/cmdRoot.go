package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile          string
	defaultConfigDir string
)

var globalFlags GlobalFlags
var outFlags OutFlags

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "awsctl",
	Short: "awsctl: AWS Management Tool",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		LaunchAWSClient(GetContext(), globalFlags)
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute starts here.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolVar(&globalFlags.DryRun, "dry-run", false, "Enable Dry Run Mode.")

	rootCmd.AddCommand(cmdGetEc2)
	rootCmd.AddCommand(cmdCW)
	rootCmd.AddCommand(cmdConfig)
	//rootCmd.AddCommand(cmdGet)
	rootCmd.AddCommand(cmdRestartEc2)
	rootCmd.AddCommand(cmdECR)
	rootCmd.AddCommand(cmdECS)
	rootCmd.AddCommand(cmdKinesis)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defaultConfigDir = home + `/.aws/awsctl/configs/default`
		viper.AddConfigPath(home)
		//viper.SetConfigName(".awsctl")
		viper.SetConfigName(".aws/.condorctl")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
	}
}
