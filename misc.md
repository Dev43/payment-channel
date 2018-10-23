package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/Dev43/payment-channel/cmd/bindings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func generateKeyPair() (*ecdsa.PrivateKey, *ecdsa.PrivateKey) {
	aPk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	bPk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	return aPk, bPk
}

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}

	aPk, _ := generateKeyPair()
	authAlice := bind.NewKeyedTransactor(aPk)
	// authBob := bind.NewKeyedTransactor(aPk)
	address, tx, paymentChannel, err := bindings.DeploySinglePaymentChannel(authAlice, client)
	fmt.Println("we have a connection")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(address, tx, paymentChannel, err)
	// // Setting nil gives you the last block number
	// balance, err := client.BalanceAt(context.Background(), address, nil)
	// fmt.Println(balance)
	// fbalance := new(big.Float)
	// fbalance.SetString(balance.String())
	// ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	// fmt.Println(ethValue)

	// privateKey, err := crypto.GenerateKey()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// privateKeyBytes := crypto.FromECDSA(privateKey)
	// // Encodes it to hex
	// fmt.Println(hexutil.Encode(privateKeyBytes)[2:])
	// publicKey := privateKey.Public()
	// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// if !ok {
	// 	log.Fatal("error casting public key to ECDSA")
	// }

	// publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	// fmt.Println(hexutil.Encode(publicKeyBytes)[4:])
	// address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	// fmt.Println(address)
	// // Deriving the address from pubkey directly
	// hash := sha3.NewKeccak256()
	// hash.Write(publicKeyBytes[1:])
	// fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 0x96216849c49358b10257cb55b28ea603c874b05e

	// // Address check regex
	// re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	// privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// publicKey := privateKey.Public()
	// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// if !ok {
	// 	log.Fatal("error casting public key to ECDSA")
	// }

	// fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// playiung with bignum
	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10)
	// Pad them to 32 bytes
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))
}





///////////////////////////////////////
// // Parse all the flags that we sent
	// flag.Parse()

	// var err error
	// chOpen, ok := new(big.Int).SetString(parsedDepositValue, 10)
	// if !ok {
	// 	log.Fatal(errors.New("Could not set value string"))
	// }
	// channelOpenValue = chOpen
	// value, ok := new(big.Int).SetString("100000000000000000", 10)
	// if !ok {
	// 	log.Fatal(errors.New("Could not set value string"))
	// }
	// nonce = big.NewInt(1)

	// // var err error
	// // var mnemonic = "toward effort adult vacant pink outdoor alcohol noise enhance term ozone slogan"
	// // var aliceKey = "a736684a498b593af152d062399aca27e51d44ae5dc652bdda049e381fb21afc"
	// // var bobKey = "0b37674d89c3418e8293a0db9eb69b3c16f1b0438793ab7ac3b2301f570017ec"

	// // alice, alicePriv := GenerateKeyPair()
	// // bob, bobPriv := GenerateKeyPair()
	// alice = common.HexToAddress("0x29535aB060046D7020f3B3464527eE24b802c871")
	// bob = common.HexToAddress("0x1b3960dB2F02C23Ed3b816750Dc4BD688B325792")
	// alicePriv, err = crypto.HexToECDSA("d1ea7553648eea5e58f22abf8b03055415d121cdb5c6e7c099e0ff232214c5f6")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// bobPriv, err = crypto.HexToECDSA("86feb87fe87829d0519b569e6ef503099d55ee4fb2c5e6a753975f4cef590461")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// client, err := connect(alice, bob, true)

	// authAlice := bind.NewKeyedTransactor(alicePriv)
	// authBob := bind.NewKeyedTransactor(bobPriv)
	// _ = authBob

	// contractAddress, _, paymentChannelAlice, err := bindings.DeploySinglePaymentChannel(authAlice, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Open the channel between alice and bob
	// authAlice.Value = channelOpenValue
	// tx, err := paymentChannelAlice.OpenChannel(authAlice, bob)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// authAlice.Value = big.NewInt(0)
	// receipt, err := client.TransactionReceipt(context.TODO(), tx.Hash())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if receipt.Status != 1 {
	// 	fmt.Println("there was a problem with the transaction")
	// }

	// // Create signatures from both

	// pr, sig, err := CreateNewMessage(alice, contractAddress, value, nonce, alicePriv)
	// prb, sigb, err := CreateNewMessage(bob, contractAddress, value, nonce, bobPriv)
	// paymentProof := PaymentProof{
	// 	Signatures: []Signature{{Sig: hexutil.Encode(sig), From: "alice"}, {Sig: hexutil.Encode(sigb), From: "bob"}},
	// 	Amount:     value.String(),
	// 	Date:       time.Now().String(),
	// 	Nonce:      nonce.String(),
	// }
	// proofs = append(proofs, paymentProof)
	// ok, err = ValidateMessage(paymentChannelAlice, sig, pr, alice, value, nonce)
	// ok, err = ValidateMessage(paymentChannelAlice, sigb, prb, bob, value, nonce)

	// // close
	// sigA, err := hexutil.Decode(proofs[len(proofs)-1].Signatures[0].Sig)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// sigBob, err := hexutil.Decode(proofs[len(proofs)-1].Signatures[1].Sig)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// r, s, v := ExtractRSVFromSignature(sigA)
	// rb, sb, vb := ExtractRSVFromSignature(sigBob)
	// tx, err = paymentChannelAlice.CloseChannel(authAlice, pr, [2]uint8{v, vb}, [2][32]byte{r, rb}, [2][32]byte{s, sb}, value, nonce)
	// receipt, err = client.TransactionReceipt(context.TODO(), tx.Hash())
	// fmt.Println(receipt.Status)

	// // challenge
	// tx, err = paymentChannelAlice.Challenge(authAlice, pr, [2]uint8{v, vb}, [2][32]byte{r, rb}, [2][32]byte{s, sb}, value, nonce)
	// receipt, err = client.TransactionReceipt(context.TODO(), tx.Hash())
	// fmt.Println(receipt.Status)

	// time, err := ethtest.EvmIncreaseTime(16 * 60)
	// fmt.Println(time, err)

	// // finalize
	// authAlice.GasLimit = 300000
	// tx, err = paymentChannelAlice.FinalizeChannel(authAlice)
	// fmt.Println(tx)
	// fmt.Println(err)
	// receipt, err = client.TransactionReceipt(context.TODO(), tx.Hash())
	// fmt.Println(receipt.Status)
	// var infor bool
	// flag.BoolVar(&infor, "info", false, "information")
	// for {
	// 	flag.Parse()
	// 	if infor {
	// 		info()
	// 	}
	// 	reader := bufio.NewReader(os.Stdin)
	// 	fmt.Print("Enter text: ")
	// 	text, _ := reader.ReadString('\n')
	// 	fmt.Println(text)
	// 	if text == "info\n" {
	// 		info()
	// 	}
	// }
	// info()

	// 
	// // Parse all the flags that we sent
	// flag.Parse()

	// var err error
	// chOpen, ok := new(big.Int).SetString(parsedDepositValue, 10)
	// if !ok {
	// 	log.Fatal(errors.New("Could not set value string"))
	// }
	// channelOpenValue = chOpen
	// value, ok := new(big.Int).SetString("100000000000000000", 10)
	// if !ok {
	// 	log.Fatal(errors.New("Could not set value string"))
	// }
	// nonce = big.NewInt(1)

	// // var err error
	// // var mnemonic = "toward effort adult vacant pink outdoor alcohol noise enhance term ozone slogan"
	// // var aliceKey = "a736684a498b593af152d062399aca27e51d44ae5dc652bdda049e381fb21afc"
	// // var bobKey = "0b37674d89c3418e8293a0db9eb69b3c16f1b0438793ab7ac3b2301f570017ec"

	// // alice, alicePriv := GenerateKeyPair()
	// // bob, bobPriv := GenerateKeyPair()
	// alice = common.HexToAddress("0x29535aB060046D7020f3B3464527eE24b802c871")
	// bob = common.HexToAddress("0x1b3960dB2F02C23Ed3b816750Dc4BD688B325792")
	// alicePriv, err = crypto.HexToECDSA("d1ea7553648eea5e58f22abf8b03055415d121cdb5c6e7c099e0ff232214c5f6")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// bobPriv, err = crypto.HexToECDSA("86feb87fe87829d0519b569e6ef503099d55ee4fb2c5e6a753975f4cef590461")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// client, err := connect(alice, bob, true)

	// authAlice := bind.NewKeyedTransactor(alicePriv)
	// authBob := bind.NewKeyedTransactor(bobPriv)
	// _ = authBob

	// contractAddress, _, paymentChannelAlice, err := bindings.DeploySinglePaymentChannel(authAlice, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Open the channel between alice and bob
	// authAlice.Value = channelOpenValue
	// tx, err := paymentChannelAlice.OpenChannel(authAlice, bob)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// authAlice.Value = big.NewInt(0)
	// receipt, err := client.TransactionReceipt(context.TODO(), tx.Hash())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if receipt.Status != 1 {
	// 	fmt.Println("there was a problem with the transaction")
	// }

	// // Create signatures from both

	// pr, sig, err := CreateNewMessage(alice, contractAddress, value, nonce, alicePriv)
	// prb, sigb, err := CreateNewMessage(bob, contractAddress, value, nonce, bobPriv)
	// paymentProof := PaymentProof{
	// 	Signatures: []Signature{{Sig: hexutil.Encode(sig), From: "alice"}, {Sig: hexutil.Encode(sigb), From: "bob"}},
	// 	Amount:     value.String(),
	// 	Date:       time.Now().String(),
	// 	Nonce:      nonce.String(),
	// }
	// proofs = append(proofs, paymentProof)
	// ok, err = ValidateMessage(paymentChannelAlice, sig, pr, alice, value, nonce)
	// ok, err = ValidateMessage(paymentChannelAlice, sigb, prb, bob, value, nonce)

	// // close
	// sigA, err := hexutil.Decode(proofs[len(proofs)-1].Signatures[0].Sig)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// sigBob, err := hexutil.Decode(proofs[len(proofs)-1].Signatures[1].Sig)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// r, s, v := ExtractRSVFromSignature(sigA)
	// rb, sb, vb := ExtractRSVFromSignature(sigBob)
	// tx, err = paymentChannelAlice.CloseChannel(authAlice, pr, [2]uint8{v, vb}, [2][32]byte{r, rb}, [2][32]byte{s, sb}, value, nonce)
	// receipt, err = client.TransactionReceipt(context.TODO(), tx.Hash())
	// fmt.Println(receipt.Status)

	// // challenge
	// tx, err = paymentChannelAlice.Challenge(authAlice, pr, [2]uint8{v, vb}, [2][32]byte{r, rb}, [2][32]byte{s, sb}, value, nonce)
	// receipt, err = client.TransactionReceipt(context.TODO(), tx.Hash())
	// fmt.Println(receipt.Status)

	// time, err := ethtest.EvmIncreaseTime(16 * 60)
	// fmt.Println(time, err)

	// // finalize
	// authAlice.GasLimit = 300000
	// tx, err = paymentChannelAlice.FinalizeChannel(authAlice)
	// fmt.Println(tx)
	// fmt.Println(err)
	// receipt, err = client.TransactionReceipt(context.TODO(), tx.Hash())
	// fmt.Println(receipt.Status)
	// var infor bool
	// flag.BoolVar(&infor, "info", false, "information")
	// for {
	// 	flag.Parse()
	// 	if infor {
	// 		info()
	// 	}
	// 	reader := bufio.NewReader(os.Stdin)
	// 	fmt.Print("Enter text: ")
	// 	text, _ := reader.ReadString('\n')
	// 	fmt.Println(text)
	// 	if text == "info\n" {
	// 		info()
	// 	}
	// }
	// info()