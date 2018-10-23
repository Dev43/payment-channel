package cmd

import (
	"fmt"
	"log"

	"github.com/Dev43/payment-channel/channel"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(infoCmd)
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		ch, err := channel.LoadChannel()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(ch.Info())
	},
}
