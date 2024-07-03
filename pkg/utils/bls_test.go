package utils

import (
	"fmt"
	"testing"
)

func TestGenerateBLSKeys(t *testing.T) {
	private, public, err := GenerateBLSKeys()
	if err != nil {
		t.FailNow()
	}
	fmt.Println("encoded private key: ", string(private))
	fmt.Println("encoded public key: ", string(public))
}

func TestPrepareKeyset(t *testing.T) {
	private0, public0, err := GenerateBLSKeys()
	if err != nil {
		t.FailNow()
	}
	fmt.Println("private key: ", string(private0))
	fmt.Println("public key: ", string(public0))

	var publicKeys []string
	publicKeys = append(publicKeys, string(public0))

	ret, err := PrepareKeyset(publicKeys, 1)
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}

	fmt.Println("keyset bytes: ", string(ret))
}
