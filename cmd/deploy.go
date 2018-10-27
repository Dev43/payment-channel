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
	Short: "deploy the payment channel contract",
	Long:  "this function deploys the payment channel contract to a blockchain",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := channel.NewChannel()
		if err != nil {
			log.Fatal(err)
		}
		address, err := c.Deploy()
		if err != nil {
			log.Fatal(err)
		}
		c.Info()
		fmt.Printf("Contract deployed at address %s, you may now open a payment channel\n", address.String())
	},
}
