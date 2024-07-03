package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/renlulu/arbitrum-orbit-sdk-go/pkg/utils"
)

func main() {
	upgradeExecutor := "0x38a6afaafdA3961b79b179CCeB0fC799d9638E46"
	sequencerInbox := "0x08425a8d2C0ECFBD7b58AFD750cE79608141D04a"
	keyset := "AAAAAAAAAAEAAAAAAAAAAQEhYBQUKaNOBbtqhhmQ1nSRitv4jsBi2qQ8WPyoPp5Elq31oQNGE8bPiV9PqJrnHTqxGwEqPchDf0gb2xtevXPufSPgwN0FLGT7bu1ONRm6AqNRvaZ6ADOiU37uU3QzK6h7ORGS4dYFFsuaW5hr5LjtyTutMYzob/z00FKhGJcq227poU1Cj+tio84E4fyMuChvDgDiqWluU6kZBleetXxg58XWa0SRbQ82DplekxypBvS8DAPuFDAtT8kBCH2gu6pLYwlGSNX/I0doNU+V86CN6Pm0m8BTGEYgCvK+GU7xxh3wTepSln2cirAtsvCWTgHsPAaACMUKFyjeA3BJnQ7XBIFXGjRSY6LBFfOjPsYPgxJMSfWJzemsMG+TaVB3EhoPSA=="
	parentChainId := 421614
	privateKey := ""
	l1conn := ""
	privateKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pk, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Decode the Base64 string into bytes
	ks, err := base64.StdEncoding.DecodeString(keyset)
	if err != nil {
		log.Fatalf("Failed to decode Base64 string: %v", err)
	}
	txn, err := utils.WhiteListDAKeySet(
		common.HexToAddress(upgradeExecutor),
		common.HexToAddress(sequencerInbox),
		l1conn,
		big.NewInt(int64(parentChainId)),
		pk,
		ks,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(txn.Hash())

}
