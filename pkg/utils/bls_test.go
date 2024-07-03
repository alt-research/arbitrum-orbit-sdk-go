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
	_, public0, err := GenerateBLSKeys()
	if err != nil {
		t.FailNow()
	}
	_, public1, err := GenerateBLSKeys()
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}

	var publicKeys []string
	publicKeys = append(publicKeys, string(public0))
	publicKeys = append(publicKeys, string(public1))

	ret, err := PrepareKeyset(publicKeys, 2)
	if err != nil {
		fmt.Println(err.Error())
		t.FailNow()
	}

	fmt.Println(string(ret))
}
