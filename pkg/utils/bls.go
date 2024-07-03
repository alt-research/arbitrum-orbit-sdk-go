package utils

import (
	"encoding/base64"
)

func GenerateBLSKeys() ([]byte, []byte, error) {
	pubKey, privKey, err := GenerateKeys()
	if err != nil {
		return nil, nil, err
	}
	pubKeyBytes := PublicKeyToBytes(pubKey)
	encodedPubKey := make([]byte, base64.StdEncoding.EncodedLen(len(pubKeyBytes)))
	base64.StdEncoding.Encode(encodedPubKey, pubKeyBytes)
	privKeyBytes := PrivateKeyToBytes(privKey)
	encodedPrivKey := make([]byte, base64.StdEncoding.EncodedLen(len(privKeyBytes)))
	base64.StdEncoding.Encode(encodedPrivKey, privKeyBytes)
	return encodedPrivKey, encodedPubKey, nil
}
