package message

import (
	"fmt"
	"testing"
)

func TestTxx(t *testing.T) {
	err := Txx()
	if err != nil {
		fmt.Println("err: ", err.Error())
		t.Fail()
	}
}
