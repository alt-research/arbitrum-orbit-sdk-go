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
