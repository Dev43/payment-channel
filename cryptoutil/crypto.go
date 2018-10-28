package cryptoutil

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

// GenerateKeyPair is a utility function to generate and format a new key pair for our use
func GenerateKeyPair() (common.Address, *ecdsa.PrivateKey) {
	priv, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	publicKey := priv.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return address, priv
}

// SignHash prefixes our string to hash with "\x19Ethereum Signed Message:\n32" and takes the
// keccack256 hash of it
func SignHash(data string) ([]byte, string) {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n32" + data)
	fmt.Println("the msg is", msg)
	return crypto.Keccak256([]byte(msg)), msg
}

// Sign does all the necessary steps to correctly hash and sign an ethereum message
func Sign(proof common.Hash, priv *ecdsa.PrivateKey) ([]byte, error) {
	// We need to hash the proof  with "\x19Ethereum Signed Message:\n%d%s"
	hash, _ := core.SignHash(proof.Bytes())
	// Sign using a private key
	signature, err := crypto.Sign(hash, priv)
	if err != nil {
		return nil, err
	}
	return signature, nil
}

// ExtractRSVFromSignature extracts the R,S,V from the signature
func ExtractRSVFromSignature(sig []byte) ([32]byte, [32]byte, uint8) {
	sigstr := common.Bytes2Hex(sig)
	rS := sigstr[0:64]
	sS := sigstr[64:128]
	R := [32]byte{}
	S := [32]byte{}
	copy(R[:], common.FromHex(rS))
	copy(S[:], common.FromHex(sS))
	vStr := sigstr[128:130]
	vI, _ := strconv.Atoi(vStr)
	V := uint8(vI + 27)

	return R, S, V
}

// MnemonicToKeys takes a mnemonic and returns a pointer to a wallet
func MnemonicToKeys(mnemonic string) (*hdwallet.Wallet, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}
