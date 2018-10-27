package cmd

import (
	"log"

	"github.com/Dev43/payment-channel/channel"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(infoCmd)
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print important information about the payment channel",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := channel.NewChannel()
		if err != nil {
			log.Fatal(err)
		}
		c.Info()
		c.Balance()
	},
}
