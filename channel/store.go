package channel

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	"github.com/Dev43/payment-channel/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ethereum/go-ethereum/crypto"
)

type JsonStorage struct {
	PaymentProofs   []PaymentProof  `json:"payment_proof"`
	TotalBalance    string          `json:"total_balance"`
	LatestBalance   string          `json:"latest_balance"`
	Accounts        map[string]Keys `json:"accounts"`
	ContractAddress string          `json:"contract_address"`
}

type Keys struct {
	Address    string `json:"address"`
	PrivateKey string `json:"private_key"`
}

func CreateStorage(safe bool) error {
	if ok := exists(); !ok && safe {
		return errors.New("Storage already exists")
	}

	alice := common.HexToAddress("0x29535aB060046D7020f3B3464527eE24b802c871")
	bob := common.HexToAddress("0x1b3960dB2F02C23Ed3b816750Dc4BD688B325792")
	alicePriv, err := crypto.HexToECDSA("d1ea7553648eea5e58f22abf8b03055415d121cdb5c6e7c099e0ff232214c5f6")
	if err != nil {
		log.Fatal(err)
	}

	bobPriv, err := crypto.HexToECDSA("86feb87fe87829d0519b569e6ef503099d55ee4fb2c5e6a753975f4cef590461")
	if err != nil {
		log.Fatal(err)
	}
	aliceKeys := Keys{
		Address:    alice.String(),
		PrivateKey: hexutil.Encode(crypto.FromECDSA(alicePriv)),
	}
	bobKeys := Keys{
		Address:    bob.String(),
		PrivateKey: hexutil.Encode(crypto.FromECDSA(bobPriv)),
	}
	// aliceKey, alicePriv := cryptoutil.GenerateKeyPair()
	// bobKey, bobPriv := cryptoutil.GenerateKeyPair()
	// aliceKeys := Keys{
	// 	Address:    aliceKey.String(),
	// 	PrivateKey: hexutil.Encode(crypto.FromECDSA(alicePriv)),
	// }
	// bobKeys := Keys{
	// 	Address:    bobKey.String(),
	// 	PrivateKey: hexutil.Encode(crypto.FromECDSA(bobPriv)),
	// }
	j := JsonStorage{
		TotalBalance:  "0",
		LatestBalance: "0",
		Accounts: map[string]Keys{
			"alice": aliceKeys,
			"bob":   bobKeys,
		},
		ContractAddress: util.ZeroAddress,
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
	ok := exists()
	if !ok {
		return &JsonStorage{}, errors.New("Storage does not exist, run init")
	}
	data, err := ioutil.ReadFile("storage.json")
	if err != nil {
		return &JsonStorage{}, err
	}
	json.Unmarshal(data, &s)
	return &s, nil
}

// see if it exists
func exists() bool {
	f, err := os.Open("storage.json")
	if err != nil {
		return false
		// TO DO -- add a way to remove with current file with a flag
	}
	defer f.Close()
	return true
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
