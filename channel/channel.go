package channel

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/Dev43/payment-channel/bindings"
	"github.com/Dev43/payment-channel/cryptoutil"
	"github.com/Dev43/payment-channel/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Signature struct {
	Sig  string
	From string
}

type PaymentProof struct {
	Signatures []Signature
	Amount     string
	Date       string
	Nonce      string
}

type Channel struct {
	store        *JsonStorage
	client       *ethclient.Client
	value        *big.Int
	totalBalance *big.Int
}

func InitStorage() error {
	err := CreateStorage()
	if err != nil {
		return err
	}
	return nil
}

func NewChannel() (*Channel, error) {
	j, err := LoadStorage()
	if err != nil {
		return &Channel{}, err
	}
	cli, err := Connect(common.HexToAddress(j.Accounts["alice"].Address), common.HexToAddress(j.Accounts["bob"].Address), false)
	if err != nil {
		return &Channel{}, err
	}
	tb, ok := new(big.Int).SetString(j.TotalBalance, 10)
	if !ok {
		return &Channel{}, errors.New("Could not set big.Int")
	}
	lb, ok := new(big.Int).SetString(j.LatestBalance, 10)
	if !ok {
		return &Channel{}, errors.New("Could not set big.Int")
	}
	c := Channel{
		store:        j,
		client:       cli,
		totalBalance: tb,
		value:        lb,
	}
	return &c, nil
}

func LoadChannel() (*Channel, error) {
	j, err := LoadStorage()
	if err != nil {
		return &Channel{}, err
	}
	totalB, ok := new(big.Int).SetString(j.TotalBalance, 10)
	if !ok {
		return &Channel{}, errors.New("could not set big int")
	}
	v, ok := new(big.Int).SetString(j.LatestBalance, 10)
	if !ok {
		return &Channel{}, errors.New("could not set big int")
	}

	c := Channel{
		store:        j,
		value:        v,
		totalBalance: totalB,
	}
	return &c, nil
}

func Connect(alice common.Address, bob common.Address, isSim bool) (*ethclient.Client, error) {
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

func (c *Channel) CreateNewMessage(address common.Address, cAddress common.Address, value *big.Int, nonce *big.Int, priv *ecdsa.PrivateKey) (common.Hash, []byte, error) {
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

func (c *Channel) ValidateMessage(paymentChannel *bindings.SinglePaymentChannel, signature []byte, proof common.Hash, originator common.Address, value *big.Int, nonce *big.Int) (bool, error) {
	// Extract the r,s,v of the signature
	r, s, v := cryptoutil.ExtractRSVFromSignature(signature)
	// Let's verify our signature is correct
	ok, err := paymentChannel.SinglePaymentChannelCaller.VerifyValidityOfMessage(nil, proof, v, r, s, value, nonce, originator)
	if err != nil {
		return false, err
	}
	return ok, nil
}

func (c *Channel) Info() string {
	return fmt.Sprintf(`

	Alice signatures: %+v Amount: %s
	Bob signatures:  Amount: %s

	`, c.store.PaymentProofs, util.ToDecimal(c.totalBalance), util.ToDecimal(new(big.Int).Sub(c.totalBalance, c.value)))
}
