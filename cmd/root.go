package cmd

import (
	"fmt"
	"os"

	"github.com/krobus00/nexus-service/internal/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nexus-service",
	Short: "nexus",
	Long:  `"nexus" or hub that connects multiple APIs together`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func Init() {
	if err := config.LoadConfig(); err != nil {
		log.Fatalln(err.Error())
	}

	log.Info(fmt.Sprintf("starting %s:%s...", config.ServiceName(), config.ServiceVersion()))
}
