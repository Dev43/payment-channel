package channel

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/Dev43/payment-channel/bindings"
	"github.com/Dev43/payment-channel/cryptoutil"
	"github.com/Dev43/payment-channel/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Channel struct {
	store         Storage
	client        *ethclient.Client
	latestBalance *big.Int
	totalBalance  *big.Int
	privKeyA      *ecdsa.PrivateKey
	privKeyB      *ecdsa.PrivateKey
	address       common.Address
	paymentProofs []PaymentProof
	accounts      map[string]Account
}

type Signature struct {
	Sig  string `json:"sig"`
	From string `json:"from"`
}

type PaymentProof struct {
	Signatures []Signature `json:"signatures"`
	Amount     string      `json:"amount"`
	Date       string      `json:"date"`
	Nonce      string      `json:"nonce"`
	Proof      string      `json:"proof"`
}

type Account struct {
	address common.Address
	privKey *ecdsa.PrivateKey
}

type Storage interface {
	Create() (*Channel, error)
	Load() (*Channel, error)
	Save(*Channel) error
}

func InitStorage() error {
	s := NewStorage()
	_, err := s.Create()
	if err != nil {
		return err
	}
	return nil
}

func NewChannel() (*Channel, error) {
	s := NewStorage()
	// TODO add diff url
	cli, err := connect()
	if err != nil {
		return &Channel{}, err
	}
	c, err := s.Load()
	if err != nil {
		return &Channel{}, err
	}
	c.client = cli
	return c, nil
}

func (c *Channel) ValidateNonce(nonce *big.Int) error {
	sigs := c.paymentProofs
	// If there are no signatures, then we can use whatever nonce we want
	if len(sigs) == 0 {
		return nil
	}

	ns := sigs[len(sigs)-1].Nonce
	lastNonce, _ := new(big.Int).SetString(ns, 10)
	if nonce.Cmp(lastNonce) < 0 {
		return errors.New("Invalid nonce, it needs to be equal or higher to the last nonce")
	}
	return nil
}

func (c *Channel) Deploy() (common.Address, error) {

	auth := bind.NewKeyedTransactor(c.privKeyA)

	contractAddress, _, _, err := bindings.DeploySinglePaymentChannel(auth, c.client)
	if err != nil {
		return common.Address{}, err
	}
	c.address = contractAddress
	err = c.store.Save(c)
	if err != nil {
		return common.Address{}, err
	}
	return contractAddress, nil
}

func (c *Channel) Open(openingValue *big.Int, counterParty common.Address) error {
	// For now alice's priv key
	cp := counterParty
	if cp == common.HexToAddress(util.ZeroAddress) {
		cp = c.accounts["bob"].address
	}
	if c.address.String() == util.ZeroAddress {
		return errors.New("You need to deploy a contract first")
	}
	auth := bind.NewKeyedTransactor(c.privKeyA)
	paymentChannel, err := bindings.NewSinglePaymentChannel(c.address, c.client)
	if err != nil {
		return err
	}
	// Open the channel between alice and bob with value openingValue
	auth.Value = openingValue
	tx, err := paymentChannel.OpenChannel(auth, cp)
	if err != nil {
		return err
	}
	// Remove any further sending of value
	auth.Value = big.NewInt(0)
	// Ensure our transaction went through
	receipt, err := c.client.TransactionReceipt(context.TODO(), tx.Hash())
	if err != nil {
		return err
	}
	if receipt.Status != 1 {
		return fmt.Errorf("Problem at the EVM execution level for transaction %s ", receipt.TxHash.String())
	}
	c.totalBalance = openingValue
	c.store.Save(c)
	return nil
}

func (c *Channel) Close() error {
	var sigs [2][]byte
	var r, s [32]byte
	var v uint8
	if len(c.paymentProofs) == 0 {
		return errors.New("No proofs")
	}
	latestProof := c.paymentProofs[len(c.paymentProofs)-1]
	for i, sig := range latestProof.Signatures {
		s, err := hexutil.Decode(sig.Sig)
		if err != nil {
			return err
		}
		sigs[i] = s
	}

	r, s, v = cryptoutil.ExtractRSVFromSignature(sigs[0])

	value, ok := new(big.Int).SetString(latestProof.Amount, 10)
	if !ok {
		log.Fatal(errors.New("Could not set the string inputted to a big.Int"))
	}
	nonce, ok := new(big.Int).SetString(latestProof.Nonce, 10)
	if !ok {
		log.Fatal(errors.New("Could not set the string inputted to a big.Int"))
	}

	auth := bind.NewKeyedTransactor(c.privKeyA)
	paymentChannel, err := bindings.NewSinglePaymentChannel(c.address, c.client)
	if err != nil {
		return err
	}

	tx, err := paymentChannel.CloseChannel(auth, common.HexToHash(latestProof.Proof), v, r, s, value, nonce)
	if err != nil {
		return err
	}
	receipt, err := c.client.TransactionReceipt(context.TODO(), tx.Hash())
	if err != nil {
		return err
	}
	if receipt.Status != 1 {
		return errors.New("Undefined error when closing the channel")
	}
	return nil
}

// finalize
func (c *Channel) Finalize() error {
	auth := bind.NewKeyedTransactor(c.privKeyA)
	auth.GasLimit = 300000
	paymentChannel, err := bindings.NewSinglePaymentChannel(c.address, c.client)
	if err != nil {
		return err
	}

	tx, err := paymentChannel.FinalizeChannel(auth)
	if err != nil {
		return err
	}
	receipt, err := c.client.TransactionReceipt(context.TODO(), tx.Hash())
	if err != nil {
		return err
	}
	if receipt.Status != 1 {
		return errors.New("Undefined error when closing the channel")
	}
	return nil
}

func (c *Channel) Challenge() error {
	var sigs [2][]byte
	var r, s [32]byte
	var v uint8
	if len(c.paymentProofs) == 0 {
		return errors.New("No proofs")
	}
	latestProof := c.paymentProofs[len(c.paymentProofs)-1]
	for i, sig := range latestProof.Signatures {
		s, err := hexutil.Decode(sig.Sig)
		if err != nil {
			return err
		}
		sigs[i] = s
	}

	r, s, v = cryptoutil.ExtractRSVFromSignature(sigs[0])

	value, ok := new(big.Int).SetString(latestProof.Amount, 10)
	if !ok {
		log.Fatal(errors.New("Could not set the string inputted to a big.Int"))
	}
	nonce, ok := new(big.Int).SetString(latestProof.Nonce, 10)
	if !ok {
		log.Fatal(errors.New("Could not set the string inputted to a big.Int"))
	}

	auth := bind.NewKeyedTransactor(c.privKeyA)
	paymentChannel, err := bindings.NewSinglePaymentChannel(c.address, c.client)
	if err != nil {
		return err
	}

	tx, err := paymentChannel.Challenge(auth, common.HexToHash(latestProof.Proof), v, r, s, value, nonce)
	if err != nil {
		return err
	}
	receipt, err := c.client.TransactionReceipt(context.TODO(), tx.Hash())
	if err != nil {
		return err
	}
	if receipt.Status != 1 {
		return errors.New("Undefined error when challenging the channel")
	}
	return nil
}

func CreateNewMessage(address common.Address, cAddress common.Address, value *big.Int, nonce *big.Int, priv *ecdsa.PrivateKey) (common.Hash, []byte, error) {
	data, err := formatData(address, cAddress, value, nonce)
	if err != nil {
		return common.Hash{}, nil, err
	}
	proof := crypto.Keccak256Hash(data)
	signature, err := cryptoutil.Sign(proof, priv)
	if err != nil {
		return common.Hash{}, nil, err
	}
	return proof, signature, nil
}

// TODO split this function so one can decide who is creating the signature
func (c *Channel) CreateSignatures(value *big.Int, nonce *big.Int) error {
	alice := c.accounts["alice"].address
	pr, sig, err := CreateNewMessage(alice, c.address, value, nonce, c.privKeyA)
	if err != nil {
		return err
	}

	paymentProof := PaymentProof{
		Signatures: []Signature{{Sig: hexutil.Encode(sig), From: "alice"}},
		Amount:     value.String(),
		Date:       time.Now().String(),
		Nonce:      nonce.String(),
		Proof:      pr.String(),
	}
	c.paymentProofs = append(c.paymentProofs, paymentProof)
	c.latestBalance = value
	err = c.store.Save(c)
	if err != nil {
		return err
	}
	return nil
}

func (c *Channel) VerifyMessages() error {
	paymentChannel, err := bindings.NewSinglePaymentChannel(c.address, c.client)
	if err != nil {
		return err
	}
	pproofs := c.paymentProofs
	if len(pproofs) == 0 {
		return errors.New("No signatures to verify")
	}
	latestProof := pproofs[len(pproofs)-1]
	for _, sig := range latestProof.Signatures {
		sigb, err := hexutil.Decode(sig.Sig)
		if err != nil {
			return err
		}
		v, ok := new(big.Int).SetString(latestProof.Amount, 10)
		if !ok {
			return errors.New("could not switch to big.Int")
		}
		n, ok := new(big.Int).SetString(latestProof.Nonce, 10)
		if !ok {
			return errors.New("could not switch to big.Int")
		}
		proof := common.HexToHash(latestProof.Proof)
		// TODO change this
		err = validateMessage(paymentChannel, sigb, proof, v, n)
		if err != nil {
			return err
		}

	}
	return nil

}

// TODO
// - Add closing time, timeout time, challenge time to info
// - Add a transaction viewing function that tracks the transactions from alice to bob
func (c *Channel) Info() {
	fmt.Println(fmt.Sprintf(`

	So far:

	Globals:
		totalBalance: %s,
		latestBalance:        %s,
		address:      %s,

		PaymentProofs: %+v
	
	`, c.totalBalance.String(), c.latestBalance.String(), c.address.String(), c.paymentProofs))
}

func (c *Channel) Balance() {
	fmt.Println(fmt.Sprintf(`
	-------------------	
	|Alice: %s|Bob: %s|
	-------------------
	
	`, util.ToDecimal(new(big.Int).Sub(c.totalBalance, c.latestBalance)), util.ToDecimal(c.latestBalance)))
}

func validateMessage(paymentChannel *bindings.SinglePaymentChannel, signature []byte, proof common.Hash, value *big.Int, nonce *big.Int) error {
	// Extract the r,s,v of the signature
	r, s, v := cryptoutil.ExtractRSVFromSignature(signature)
	// Let's verify our signature is correct
	ok, err := paymentChannel.SinglePaymentChannelCaller.VerifyValidityOfMessage(nil, proof, v, r, s, value, nonce)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("There was an error verifying the validity of the Message")
	}
	return nil
}

func connect() (*ethclient.Client, error) {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}
	return client, err
}

func formatData(address common.Address, contractAddress common.Address, value *big.Int, nonce *big.Int) ([]byte, error) {

	paddedAddress := (contractAddress.Bytes())
	// Let's pad the values so they are uint256
	paddedValue := common.LeftPadBytes(value.Bytes(), 32)
	paddedNonce := common.LeftPadBytes(nonce.Bytes(), 32)

	var data []byte
	data = append(data, paddedAddress...)
	data = append(data, paddedValue...)
	data = append(data, paddedNonce...)

	return data, nil
}
