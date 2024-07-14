package utils

import (
	"crypto/ecdsa"
	"crypto/rand"
	"errors"

	"github.com/ethereum/go-ethereum/common"
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

func GetAddressFromPrivateKey(privateKey string) (*ecdsa.PrivateKey, common.Address, error) {
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, common.Address{}, nil
	}
	publicKey := key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, common.Address{}, errors.New("Falied to cast public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return key, fromAddress, nil
}
