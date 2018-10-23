package cryptoutil

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core"
)

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

func SignHash(data string) ([]byte, string) {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n32" + data)
	fmt.Println("the msg is", msg)
	return crypto.Keccak256([]byte(msg)), msg
}

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

// ExtractRSVFromSignature signatures R S V returned as arrays
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
