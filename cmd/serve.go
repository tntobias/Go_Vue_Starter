package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/tntobias/Go_Vue_Starter/config"
	"github.com/tntobias/Go_Vue_Starter/server"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start app Server",
	Long:  `Start the application server`,
	Run:   run,
}

func init() {
	RootCmd.AddCommand(serveCmd)
}

func run(cmd *cobra.Command, args []string) {
	cfg, err := config.LoadConfig(RootCmd)
	if err != nil {
		log.Fatal("Failed to load config: " + err.Error())
	}

	srv := server.New(cfg)
	if err := srv.Start(); err != nil {
		os.Exit(1)
	}
}
