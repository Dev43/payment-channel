package channel

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/Dev43/payment-channel/cryptoutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type JsonStorage struct {
	PaymentProofs []PaymentProof  `json:"payment_proof"`
	TotalBalance  string          `json:"total_balance"`
	LatestBalance string          `json:"latest_balance"`
	Accounts      map[string]Keys `json:"accounts"`
}

type Keys struct {
	Address    string `json:"address"`
	PrivateKey string `json:"private_key"`
}

func CreateStorage() error {
	fmt.Println(exists())
	if err := exists(); err != nil {
		return err
	}
	aliceKey, alicePriv := cryptoutil.GenerateKeyPair()
	bobKey, bobPriv := cryptoutil.GenerateKeyPair()
	aliceKeys := Keys{
		Address:    aliceKey.String(),
		PrivateKey: hexutil.Encode(crypto.FromECDSA(alicePriv)),
	}
	bobKeys := Keys{
		Address:    bobKey.String(),
		PrivateKey: hexutil.Encode(crypto.FromECDSA(bobPriv)),
	}
	j := JsonStorage{
		TotalBalance:  "0",
		LatestBalance: "0",
		Accounts: map[string]Keys{
			"alice": aliceKeys,
			"bob":   bobKeys,
		},
	}
	b, err := json.Marshal(j)
	if err != nil {
		return err
	}
	ioutil.WriteFile("storage.json", b, 0666)
	return nil
}

func LoadStorage() (*JsonStorage, error) {
	s := JsonStorage{}
	err := exists()
	if err != nil {
		return &JsonStorage{}, err
	}
	data, err := ioutil.ReadFile("storage.json")
	if err != nil {
		return &JsonStorage{}, err
	}
	json.Unmarshal(data, &s)
	return &s, nil
}

// see if it exists
func exists() error {
	f, err := os.Open("storage.json")
	if err == nil {
		return errors.New("Storage file already exists")
		// TO DO -- add a way to remove with current file with a flag
	}
	defer f.Close()
	return nil
}

func (j *JsonStorage) Save() error {

	b, err := json.Marshal(j)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("storage.json", b, 0666)
	if err != nil {
		return err
	}
	return nil
}

func (j *JsonStorage) AddToProofs(proof PaymentProof) bool {
	j.PaymentProofs = append(j.PaymentProofs, proof)
	return true
}

func (j *JsonStorage) Balance() (*big.Int, error) {
	totalBalance, ok := new(big.Int).SetString(j.TotalBalance, 10)
	if !ok {
		return big.NewInt(0), errors.New("could not transform string into big number")
	}

	latestBalance, ok := new(big.Int).SetString(j.LatestBalance, 10)
	if !ok {
		return big.NewInt(0), errors.New("could not transform string into big number")
	}

	return new(big.Int).Sub(totalBalance, latestBalance), nil
}
