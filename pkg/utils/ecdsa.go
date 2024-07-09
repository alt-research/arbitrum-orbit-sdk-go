package utils

import (
	"crypto/ecdsa"
	"crypto/rand"

	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateECDSAKeys() ([]byte, []byte, error) {
	privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.PublicKey
	publicKeyBytes := crypto.FromECDSAPub(&publicKey)
	return privateKeyBytes, publicKeyBytes, nil
}
