package cmd

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/Dev43/payment-channel/channel"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// Nonce flag
var Nonce string

func init() {
	// open
	channelCmd.AddCommand(openCmd)

	// sign
	signCmd.Flags().StringVarP(&Nonce, "nonce", "n", "0", "Nonce to use")
	signCmd.MarkFlagRequired("nonce")
	channelCmd.AddCommand(signCmd)

	// verify
	channelCmd.AddCommand(verifyCmd)

	// close
	channelCmd.AddCommand(closeCmd)

	// close
	channelCmd.AddCommand(challengeCmd)

	// finalize
	channelCmd.AddCommand(finalizeCmd)

	// rootCmd
	rootCmd.AddCommand(channelCmd)

	// Generate markdown tree command dynamically
	err := doc.GenMarkdownTree(rootCmd, "./docs")
	if err != nil {
		log.Fatal(err)
	}
}

var channelCmd = &cobra.Command{
	Use:   "channel",
	Short: "Channel functions",
	Long:  `All channel functions`,
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		c, err := channel.NewChannel()
		if err != nil {
			log.Fatal(err)
		}
		c.Info()
		c.Balance()
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You need to give me arguments")
	},
}

var openCmd = &cobra.Command{
	Use:   "open [value]",
	Short: "Open a payment channel between Alice and Bob",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		value, ok := new(big.Int).SetString(args[0], 10)
		if !ok {
			log.Fatal(errors.New("Could not set the string inputted to a big.Int"))
		}
		c, err := channel.NewChannel()
		if err != nil {
			log.Fatal(err)
		}
		err = c.Open(value)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Payment channel successfully opened")
	},
}

var signCmd = &cobra.Command{
	Use:   "sign [value]",
	Short: "Signs messages with a specific value attached to it",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// the first argument is the amount to sign
		// can change the nonce if needed
		value, ok := new(big.Int).SetString(args[0], 10)
		if !ok {
			log.Fatal(errors.New("Could not set the string inputted to a big.Int"))
		}
		nonce, ok := new(big.Int).SetString(Nonce, 10)
		if !ok {
			log.Fatal(errors.New("Could not set the string inputted to a big.Int"))
		}
		c, err := channel.NewChannel()
		if err != nil {
			log.Fatal(err)
		}
		err = c.ValidateNonce(nonce)
		if err != nil {
			log.Fatal(err)
		}
		err = c.CreateSignatures(value, nonce)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Signatures verified")
	},
}

// TO DO, ask the user which signature to verify
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "verifies the messages that are stored",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := channel.NewChannel()
		if err != nil {
			log.Fatal(err)
		}
		err = c.VerifyMessages()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Signatures verified!")
	},
}

var closeCmd = &cobra.Command{
	Use:   "close",
	Short: "closes the payment channel",
	Run: func(cmd *cobra.Command, args []string) {
		// the first argument is the amount to sign
		// can change the nonce if needed
		c, err := channel.NewChannel()
		if err != nil {
			log.Fatal(err)
		}
		if len(args) > 0 {
			index, err := strconv.ParseInt(args[0], 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			err = c.Close(int(index))
		} else {
			err = c.Close(0)
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Channel successfully closed")
	},
}

var challengeCmd = &cobra.Command{
	Use:       "challenge [from]",
	Short:     "challenges the payment channel",
	Long:      "challenges the payment channel, from can either be alice or bob",
	ValidArgs: []string{"alice", "bob"},
	Args:      cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// the first argument is the amount to sign
		// can change the nonce if needed
		c, err := channel.NewChannel()
		if err != nil {
			log.Fatal(err)
		}
		err = c.Challenge(args[0])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Channel successfully challenged")
	},
}

var finalizeCmd = &cobra.Command{
	Use:   "finalize",
	Short: "finalizes the payment channel",
	Run: func(cmd *cobra.Command, args []string) {
		// the first argument is the amount to sign
		// can change the nonce if needed
		c, err := channel.NewChannel()
		if err != nil {
			log.Fatal(err)
		}
		err = c.Finalize()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Channel successfully finalized")
	},
}
