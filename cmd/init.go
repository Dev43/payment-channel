package cmd

import (
	"fmt"
	"log"

	"github.com/Dev43/payment-channel/channel"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init [mnemonic]",
	Short: "Initialize a payment channel",
	Long:  "This function initializes a payment channel and 2 people's accounts: Alice and Bob",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := channel.InitStorage(args[0])
		if err != nil {
			log.Fatal("Error: ", err)
		}

		fmt.Println("New payment channel initalized, you may now deploy the contract")
	},
}
