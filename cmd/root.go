package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "payment_channel",
	Short: "Payment channel visualizer",
	Long:  `A way to visualize payment channels and create signatures`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Payment channel visualizer")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
