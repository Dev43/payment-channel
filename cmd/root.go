package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "payment_channel",
	Short: "Payment channel visualizer",
	Long:  `CLI tool to understand and visualize how payment channels work`,
}

// Execute is our top line function for all CLI commands
func Execute() {

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
