package cmd

import (
	"github.com/spf13/cobra"
)

var showImgTags bool
var cmdECRImages = &cobra.Command{
	Use:   "images",
	Short: "Elastic Container Registry Images",
	Run: func(cmd *cobra.Command, args []string) {
		PrintAWS(ListImages(awsFlags), outFlags.Format)
	},
}

func init() {
	cmdECRImages.Flags().BoolVarP(&showImgTags, "tags", "T", false, "Display Image Tags.")
}
