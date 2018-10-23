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
	Use:   "init",
	Short: "Initialize a payment channel",
	Long:  "This function initializes a payment channel and 2 people's accounts: Alice and Bob",
	Run: func(cmd *cobra.Command, args []string) {
		err := channel.InitStorage()
		if err != nil {
			log.Fatal("Error: ", err)
		}
		fmt.Println("New payment channel initalized")
	},
}
