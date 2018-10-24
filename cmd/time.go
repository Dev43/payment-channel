package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Dev43/ethtest"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(timeCmd)
}

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Advance time",
	Long:  "Bring time forward",
	Run: func(cmd *cobra.Command, args []string) {
		time, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		t, err := ethtest.EvmIncreaseTime(time)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Time accelerated by ", t, " seconds")
	},
}
