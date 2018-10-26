package cmd

import (
	"log"

	"github.com/Dev43/payment-channel/channel"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(balanceCmd)
}

var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Prints the current balance of the two accounts",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := channel.NewChannel()
		if err != nil {
			log.Fatal(err)
		}
		c.Balance()
	},
}
