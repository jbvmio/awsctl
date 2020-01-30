package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmdUseContext = &cobra.Command{
	Use:     "use-context",
	Aliases: []string{"use"},
	Short:   "switch active context",
	Long:    `command to switch active context`,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		context := strings.Join(args, " ")
		ctx := GetContext(context)
		viper.Set("current-context", ctx.Name)
		if err := viper.WriteConfig(); err != nil {
			Failf("Unable to write config: %v", err)
		}
		Infof("Context changed to %v", ctx.Name)
	},
}

func init() {
}
