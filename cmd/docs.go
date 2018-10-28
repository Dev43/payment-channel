package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func init() {
	rootCmd.AddCommand(docCmd)
}

var docCmd = &cobra.Command{
	Use:   "docs",
	Short: "Generate markdown docs for payment_channel",
	Run: func(cmd *cobra.Command, args []string) {
		// Generate markdown tree command dynamically
		err := doc.GenMarkdownTree(rootCmd, "./docs")
		if err != nil {
			log.Fatal(err)
		}
	},
}
