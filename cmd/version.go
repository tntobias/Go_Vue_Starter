package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "The version of the app",
	Long:  `Return the current running version of the app`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version ", AppVer)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
