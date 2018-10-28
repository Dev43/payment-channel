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
	"github.com/k0kubun/pp"

	"github.com/Dev43/payment-channel/bindings"
	"github.com/Dev43/payment-channel/cryptoutil"
	"github.com/Dev43/payment-channel/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Channel keeps all the information necessary for our channel
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
	state         string // State of the current contract
	finalizeTime  int64
}

// Payment proofs are appended to into an array. We keep all of the payment proofs created
// They aggregate all information neccessary to prove that a payment is valid
type PaymentProof struct {
	Signatures []Signature `json:"signatures"`
	Amount     string      `json:"amount"`
	Date       string      `json:"date"`
	Nonce      string      `json:"nonce"`
	Proof      string      `json:"proof"`
}

// Account struct is where the user accounts are held
type Account struct {
	address common.Address
	privKey *ecdsa.PrivateKey
}

// Storage is our defined interface to astract our storage layer
type Storage interface {
	Create(string) (*Channel, error)
	Load() (*Channel, error)
	Save(*Channel) error
}

// InitStorage initializes our storage, whatever it is on the backend
func InitStorage(mnemonic string) error {
	s := NewStorage()
	_, err := s.Create(mnemonic)
	if err != nil {
		return err
	}
	return nil
}

// NewChannel initiates our connection to the blockchain and loads all of
// the data into our channel pointer
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

// ValidateNonce ensures that the nonce is higher than the last nonce
func (c *Channel) ValidateNonce(nonce *big.Int) error {
	sigs := c.paymentProofs

	// Cannot have a signature with a nonce of 0
	if nonce.Cmp(big.NewInt(0)) == 0 {
		return errors.New("Invalid nonce, it needs to be higher than the last nonce")
	}

	// If there are no signatures, then we can use whatever nonce we want
	if len(sigs) == 0 {
		return nil
	}

	ns := sigs[len(sigs)-1].Nonce
	lastNonce, _ := new(big.Int).SetString(ns, 10)
	if nonce.Cmp(lastNonce) <= 0 {
		return errors.New("Invalid nonce, it needs to be higher than the last nonce")
	}
	return nil
}

// Deploy is used to deploy our contract to the blockchain
func (c *Channel) Deploy() (common.Address, error) {

	auth := bind.NewKeyedTransactor(c.privKeyA)

	contractAddress, _, _, err := bindings.DeploySinglePaymentChannel(auth, c.client)
	if err != nil {
		return common.Address{}, err
	}
	c.address = contractAddress
	c.state = "deployed"
	err = c.store.Save(c)
	if err != nil {
		return common.Address{}, err
	}
	return contractAddress, nil
}

// Open opens our payment channel with an inital value
func (c *Channel) Open(openingValue *big.Int) error {
	// For now alice's priv key
	cp := c.accounts["bob"].address
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
	c.state = "opened"
	c.store.Save(c)
	if err != nil {
		return err
	}
	return nil
}

// Close initiates the closure of our channel, which is not immediate as there is a challenge period
func (c *Channel) Close(index int) error {
	var sigs [2][]byte
	var r, rb, s, sb [32]byte
	var v, vb uint8
	var proofIndex int
	if len(c.paymentProofs) == 0 {
		return errors.New("No proofs")
	}
	if proofIndex > len(c.paymentProofs)-1 {
		return errors.New("proof does not exist at this index")
	}
	proofIndex = index
	if proofIndex == 0 {
		proofIndex = len(c.paymentProofs) - 1
	}
	latestProof := c.paymentProofs[proofIndex]
	for i, sig := range latestProof.Signatures {
		s, err := hexutil.Decode(sig.Sig)
		if err != nil {
			return err
		}
		sigs[i] = s
	}

	r, s, v = cryptoutil.ExtractRSVFromSignature(sigs[0])
	rb, sb, vb = cryptoutil.ExtractRSVFromSignature(sigs[1])

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

	tx, err := paymentChannel.CloseChannel(auth, common.HexToHash(latestProof.Proof), [2]uint8{v, vb}, [2][32]byte{r, rb}, [2][32]byte{s, sb}, value, nonce)
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
	chStart, err := paymentChannel.StartChallengePeriod(nil)
	if err != nil {
		return err
	}
	chPeriodLength, err := paymentChannel.ChallengePeriodLength(nil)
	if err != nil {
		return err
	}
	// Set the challenge period
	finalizeTime := new(big.Int).Add(chStart, chPeriodLength).Int64()
	c.state = "closed"
	c.finalizeTime = finalizeTime
	c.store.Save(c)
	if err != nil {
		return err
	}
	return nil
}

// Finalize is called only after the finalize time has passed.
// It distributes the to bob and alice
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
	c.state = "finalized"
	c.store.Save(c)
	if err != nil {
		return err
	}
	return nil
}

// Challenge can be done after a channel is closed but before it is finalized
// it needs the opponent's signature for it to work
func (c *Channel) Challenge(from string) error {
	_, ok := c.accounts[from]
	if !ok {
		return errors.New("account does not exist")
	}
	var opponentSig []byte
	var r, s [32]byte
	var v uint8
	if len(c.paymentProofs) == 0 {
		return errors.New("No proofs")
	}
	latestProof := c.paymentProofs[len(c.paymentProofs)-1]
	// Get the opponent's signature
	for _, sig := range latestProof.Signatures {
		if sig.From != from {
			s, err := hexutil.Decode(sig.Sig)
			if err != nil {
				return err
			}
			opponentSig = s
		}
	}

	r, s, v = cryptoutil.ExtractRSVFromSignature(opponentSig)

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
	c.state = "challenged"
	c.latestBalance = value
	c.store.Save(c)
	if err != nil {
		return err
	}
	return nil
}

// CreateSignatures creates both alice's and bob's signatures
func (c *Channel) CreateSignatures(value *big.Int, nonce *big.Int) error {
	bob := c.accounts["bob"].address
	alice := c.accounts["alice"].address
	pr, sig, err := createNewMessage(alice, c.address, value, nonce, c.privKeyA)
	if err != nil {
		return err
	}
	pr, sigb, err := createNewMessage(bob, c.address, value, nonce, c.privKeyB)
	if err != nil {
		return err
	}

	paymentProof := PaymentProof{
		Signatures: []Signature{{Sig: hexutil.Encode(sig), From: "alice"}, {Sig: hexutil.Encode(sigb), From: "bob"}},
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

// CreateSignature creates only 1 signature based on who it's from
func (c *Channel) CreateSignature(from string, value *big.Int, nonce *big.Int) (Signature, error) {
	acct, ok := c.accounts[from]
	if !ok {
		return Signature{}, errors.New("account does not exist")
	}
	_, sig, err := createNewMessage(acct.address, c.address, value, nonce, c.privKeyA)
	if err != nil {
		return Signature{}, err
	}
	return Signature{Sig: hexutil.Encode(sig), From: from}, nil
}

// VerifyMessages verifies the last two messages sent and ensures that everything is correct
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

		originator := c.accounts[sig.From].address
		err = validateMessage(paymentChannel, sigb, proof, v, n, originator)
		if err != nil {
			return err
		}

	}
	return nil

}

// Info simply outputs some useful information to visualize
func (c *Channel) Info() {
	var timeRemaining int64
	b, _ := c.client.HeaderByNumber(context.TODO(), nil)
	pp.Printf("Channel state: %s\n", c.state)
	if c.finalizeTime == 0 {
		pp.Printf("Channel finalize time: %s\n", "not set")
	} else {
		pp.Printf("Channel finalize time: %s\n", c.finalizeTime)
		timeRemaining = c.finalizeTime - b.Time.Int64()
	}

	pp.Printf("Current Block time %s\n", b.Time.Int64())
	pp.Printf("Time remaining before finalized channel %s seconds\n", timeRemaining)
	pp.Printf("Total payment channel balance %s\n", c.totalBalance.String())
	pp.Printf("Total balance sent from Alice %s\n", c.latestBalance.String())
	pp.Printf("Contract address %s\n", c.address.String())
	pp.Print("Latest proof: ")
	if len(c.paymentProofs) > 0 {
		pp.Println(c.paymentProofs[len(c.paymentProofs)-1])
	} else {
		pp.Println("No payment proof yet")
	}
	pChannel, _ := bindings.NewSinglePaymentChannelCaller(c.address, c.client)
	latestProof, _ := pChannel.LastPaymentProof(&bind.CallOpts{})
	pp.Println("Closing proof:")
	pp.Printf("value: %s\n", latestProof.Value.String())
	pp.Printf("nonce: %s\n", latestProof.Nonce.String())
}

// Balance outputs the balance of both alice and Bob
func (c *Channel) Balance() {
	if c.state != "finalized" {
		pp.Printf("Alice's channel balance: %s\nBob's channel balance: %s\n", util.ToDecimal(new(big.Int).Sub(c.totalBalance, c.latestBalance)).StringFixed(18), util.ToDecimal(c.latestBalance).StringFixed(18))
	} else {
		pp.Printf("Alice balance: %s\nBob's balance: %s\n", util.ToDecimal(big.NewInt(0)).StringFixed(18), util.ToDecimal(big.NewInt(0)).StringFixed(18))
	}
	for key, val := range c.accounts {
		bal, _ := c.client.BalanceAt(context.TODO(), val.address, nil)
		pp.Printf("Current balance for %s: %s\n", key, util.ToDecimal(bal).StringFixed(18))
	}
}

func validateMessage(paymentChannel *bindings.SinglePaymentChannel, signature []byte, proof common.Hash, value *big.Int, nonce *big.Int, originator common.Address) error {
	// Extract the r,s,v of the signature
	r, s, v := cryptoutil.ExtractRSVFromSignature(signature)
	// Let's verify our signature is correct
	ok, err := paymentChannel.SinglePaymentChannelCaller.VerifyValidityOfMessage(nil, proof, v, r, s, value, nonce, originator)
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

func createNewMessage(address common.Address, cAddress common.Address, value *big.Int, nonce *big.Int, priv *ecdsa.PrivateKey) (common.Hash, []byte, error) {
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
