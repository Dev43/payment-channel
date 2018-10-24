package cmd

import (
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/Dev43/payment-channel/channel"
	"github.com/Dev43/payment-channel/util"
	"github.com/ethereum/go-ethereum/common"

	"github.com/spf13/cobra"
)

// For open
var OpenValue string
var With string

// For Sign
var Nonce string
var From string // not used for now

func init() {
	// open
	openCmd.Flags().StringVarP(&OpenValue, "value", "v", "0", "value of Ether used to open the payment channel")
	openCmd.MarkFlagRequired("value")
	openCmd.Flags().StringVarP(&With, "with", "w", util.ZeroAddress, "With which address to open the payment channel")
	channelCmd.AddCommand(openCmd)

	// sign
	signCmd.Flags().StringVarP(&Nonce, "nonce", "n", "0", "Nonce to use")
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
}

var channelCmd = &cobra.Command{
	Use:   "channel",
	Short: "Channel functions",
	Long:  `All channel functions`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You need to give me arguments")
	},
}

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a channel",
	Long:  "Open a channel",
	Run: func(cmd *cobra.Command, args []string) {
		value, ok := new(big.Int).SetString(OpenValue, 10)
		if !ok {
			log.Fatal(errors.New("Could not set the string inputted to a big.Int"))
		}
		otherAddress := common.HexToAddress(With)
		c, err := channel.NewChannel()
		if err != nil {
			log.Fatal(err)
		}
		err = c.Open(value, otherAddress)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Payment channel successfully opened")
	},
}

var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign a message",
	Long:  "Sign a message",
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
		c.Info()
	},
}

// TO DO, ask the user which signature to verify
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "verify both messages",
	Long:  "verify both messages",
	Run: func(cmd *cobra.Command, args []string) {
		// the first argument is the amount to sign
		// can change the nonce if needed
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
	Short: "close the channel",
	Long:  "close the channel",
	Run: func(cmd *cobra.Command, args []string) {
		// the first argument is the amount to sign
		// can change the nonce if needed
		c, err := channel.NewChannel()
		if err != nil {
			log.Fatal(err)
		}
		err = c.Close()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Channel successfully closed")
	},
}

var challengeCmd = &cobra.Command{
	Use:   "challenge",
	Short: "challenge the channel",
	Long:  "challenge the channel",
	Run: func(cmd *cobra.Command, args []string) {
		// the first argument is the amount to sign
		// can change the nonce if needed
		c, err := channel.NewChannel()
		if err != nil {
			log.Fatal(err)
		}
		err = c.Challenge()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Channel successfully challenged")
	},
}

var finalizeCmd = &cobra.Command{
	Use:   "finalize",
	Short: "finalize the channel",
	Long:  "finalize the channel",
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
