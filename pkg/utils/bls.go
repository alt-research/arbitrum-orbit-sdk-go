package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"io"
)

const BinSize = 64 * 1024 // 64 kB

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

func PrepareKeyset(pubKeys []string, assumedHonest uint64) (string, error) {
	var parsedKeys []PublicKey
	for _, key := range pubKeys {
		decoded, err := base64.StdEncoding.DecodeString(key)
		if err != nil {
			return "", err
		}
		p, err := PublicKeyFromBytes(decoded, true)
		if err != nil {
			return "", err
		}
		parsedKeys = append(parsedKeys, p)
	}
	keyset := &DataAvailabilityKeyset{
		AssumedHonest: uint64(assumedHonest),
		PubKeys:       parsedKeys,
	}
	ksBuf := bytes.NewBuffer([]byte{})
	if err := keyset.Serialize(ksBuf); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ksBuf.Bytes()), nil
}

type DataAvailabilityKeyset struct {
	AssumedHonest uint64
	PubKeys       []PublicKey
}

func (keyset *DataAvailabilityKeyset) Serialize(wr io.Writer) error {
	if err := Uint64ToWriter(keyset.AssumedHonest, wr); err != nil {
		return err
	}
	if err := Uint64ToWriter(uint64(len(keyset.PubKeys)), wr); err != nil {
		return err
	}
	for _, pk := range keyset.PubKeys {
		pkBuf := PublicKeyToBytes(pk)
		buf := []byte{byte(len(pkBuf) / 256), byte(len(pkBuf) % 256)}
		_, err := wr.Write(append(buf, pkBuf...))
		if err != nil {
			return err
		}
	}
	return nil
}

func Uint64ToWriter(val uint64, wr io.Writer) error {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], val)
	_, err := wr.Write(buf[:])
	return err
}
