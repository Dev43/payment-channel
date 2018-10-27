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

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
