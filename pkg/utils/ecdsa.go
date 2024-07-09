package utils

import (
	"crypto/ecdsa"
	"crypto/rand"

	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateECDSAKeys() ([]byte, []byte, string, error) {
	privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		return nil, nil, "", err
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.PublicKey
	publicKeyBytes := crypto.FromECDSAPub(&publicKey)
	address := crypto.PubkeyToAddress(publicKey).Hex()

	return privateKeyBytes, publicKeyBytes, address, nil
}
