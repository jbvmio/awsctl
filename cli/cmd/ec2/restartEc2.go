package ec2

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var CmdRestartEc2 = &cobra.Command{
	Use:     "restart",
	Aliases: []string{"stop", "start"},
	Short:   "Restart an EC2 Instance",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case strings.Contains(cmd.CalledAs(), "restart"):
			fmt.Println("restart Called")
		case strings.Contains(cmd.CalledAs(), "start"):
			fmt.Println("start Called")
		case strings.Contains(cmd.CalledAs(), "stop"):
			fmt.Println("stop Called")
		}

		/*
			switch {
			case cmd.Flags().Changed("out"):
				outFmt, err := cmd.Flags().GetString("out")
				if err != nil {
					out.Warnf("WARN: %v", err)
				}
				out.PrintAWS(instances, outFmt)
			default:
				out.PrintAWS(instances)
			}
		*/
	},
}

func init() {
	//CmdRestartEc2.PersistentFlags().StringVarP(&outFlags.Format, "out", "o", "", "Change Output Format - wide|long|yaml|json.")
}
