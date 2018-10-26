package cmd

import (
	"fmt"
	"log"

	"github.com/Dev43/payment-channel/channel"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deployCmd)
}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy the payment channel contract",
	Long:  "This function deploys the payment channel contract to a blockchain",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := channel.NewChannel()
		if err != nil {
			log.Fatal(err)
		}
		address, err := c.Deploy()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Contract deployed at address %s\n", address.String())
	},
}
