package cmd

import (
	"fmt"
	"log"

	"github.com/Dev43/payment-channel/channel"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(resetCmd)
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the current storage",
	Long:  "Reset the current storage",
	Run: func(cmd *cobra.Command, args []string) {
		err := channel.ResetStorage()
		if err != nil {
			log.Fatal("Error: ", err)
		}
		fmt.Println("Storage reset")
	},
}
