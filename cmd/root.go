package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const AppVer = "0.1.0"

var RootCmd = &cobra.Command{
	Use:   "starter-project",
	Short: "Change the project name",
	Long:  `You really should change the project name`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringP("config", "c", "", "path to config file")
	RootCmd.PersistentFlags().StringP("port", "p", "8190", "server listen port")
	// RootCmd.Flags().BoolP("toggle", "t", false, "help for toggle")
}
