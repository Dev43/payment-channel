package channel

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"

	"github.com/Dev43/payment-channel/cryptoutil"
	"github.com/Dev43/payment-channel/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"

	"github.com/ethereum/go-ethereum/crypto"
)

// JsonStorage is our JSON struct that will be marshalled into our file storage
type JsonStorage struct {
	PaymentProofs   []PaymentProof  `json:"payment_proof"`
	TotalBalance    string          `json:"total_balance"`
	LatestBalance   string          `json:"latest_balance"`
	Accounts        map[string]Keys `json:"accounts"`
	ContractAddress string          `json:"contract_address"`
	ContractState   string          `json:"contract_state"`
	FinalizeTime    int64           `json:"finalize_time"`
}

// Signature holds our signature
type Signature struct {
	Sig  string `json:"sig"`
	From string `json:"from"`
}

// Keys are the address and private key of an account
type Keys struct {
	Address    string `json:"address"`
	PrivateKey string `json:"private_key"`
}

// NewStorage initiates a JsonStorage pointer
func NewStorage() *JsonStorage {
	return &JsonStorage{}
}

// Create initializes our Json storage struct and saves it to the filesystem
// takes in a valid mnemonic with derivation path "m/44'/60'/0'/0/[INDEX]"
func (j *JsonStorage) Create(mnemonic string) (*Channel, error) {
	if ok := exists(); !ok {
		return nil, errors.New("Storage already exists")
	}

	wallet, err := cryptoutil.MnemonicToKeys(mnemonic)
	if err != nil {
		return nil, err
	}
	people := []string{"alice", "bob"}
	accts := make(map[string]Keys)
	for i, person := range people {

		path := hdwallet.MustParseDerivationPath(fmt.Sprint("m/44'/60'/0'/0/", i))
		account, err := wallet.Derive(path, true)
		if err != nil {
			return nil, err
		}
		priv, err := wallet.PrivateKeyBytes(account)
		if err != nil {
			return nil, err
		}

		accts[person] = Keys{
			PrivateKey: hexutil.Encode(priv),
			Address:    account.Address.String(),
		}
	}

	s := JsonStorage{
		TotalBalance:    "0",
		LatestBalance:   "0",
		Accounts:        accts,
		ContractAddress: util.ZeroAddress,
		ContractState:   "init",
		FinalizeTime:    0,
	}
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	ioutil.WriteFile("storage.json", b, 0666)
	return nil, nil
}

// Load takes looks to see if a storage file exists and if it does loads it into the program and
// converts all necessary properties and injects it into a new channel instance
func (j *JsonStorage) Load() (*Channel, error) {
	ok := exists()
	if !ok {
		return nil, errors.New("Storage does not exist, run init")
	}
	data, err := ioutil.ReadFile("storage.json")
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &j)

	tb, ok := new(big.Int).SetString(j.TotalBalance, 10)
	if !ok {
		return nil, errors.New("Could not set big.Int")
	}
	lb, ok := new(big.Int).SetString(j.LatestBalance, 10)
	if !ok {
		return nil, errors.New("Could not set big.Int")
	}
	// Default to alice's private key for now
	priv, err := hexutil.Decode(j.Accounts["alice"].PrivateKey)
	if err != nil {
		return nil, err
	}
	privEcdsa, err := crypto.ToECDSA(priv)
	if err != nil {
		return nil, err
	}
	// Default to bobs private key for now
	privB, err := hexutil.Decode(j.Accounts["bob"].PrivateKey)
	if err != nil {
		return nil, err
	}
	privEcdsaB, err := crypto.ToECDSA(privB)
	if err != nil {
		return nil, err
	}
	cAddr := common.HexToAddress(util.ZeroAddress)
	if j.ContractAddress != util.ZeroAddress {
		cAddr = common.HexToAddress(j.ContractAddress)
	}
	accts := map[string]Account{
		"alice": {
			address: common.HexToAddress(j.Accounts["alice"].Address),
			privKey: privEcdsa,
		},
		"bob": {
			address: common.HexToAddress(j.Accounts["bob"].Address),
			privKey: privEcdsaB,
		},
	}
	c := Channel{
		store:         j,
		totalBalance:  tb,
		latestBalance: lb,
		privKeyA:      privEcdsa,
		privKeyB:      privEcdsaB,
		address:       cAddr,
		paymentProofs: j.PaymentProofs,
		accounts:      accts,
		state:         j.ContractState,
		finalizeTime:  j.FinalizeTime,
	}

	return &c, nil
}

// see if it exists
func exists() bool {
	f, err := os.Open("storage.json")
	if err != nil {
		return false
	}
	defer f.Close()
	return true
}

// Save takes our current channel pointer and saves its properties in the filesystem, overwriting the
// last storage file
func (j *JsonStorage) Save(c *Channel) error {
	j.TotalBalance = c.totalBalance.String()
	j.LatestBalance = c.latestBalance.String()
	j.ContractAddress = c.address.String()
	j.ContractState = c.state
	j.FinalizeTime = c.finalizeTime
	accts := make(map[string]Keys)
	for key, value := range c.accounts {
		priv := hexutil.Encode(crypto.FromECDSA(value.privKey))
		accts[key] = Keys{
			Address:    value.address.Hex(),
			PrivateKey: priv,
		}
	}
	j.Accounts = accts
	j.PaymentProofs = c.paymentProofs

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
