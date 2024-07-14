package utils

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBasicAuth(client *ethclient.Client, privateKey string) (*bind.TransactOpts, error) {
	key, fromAddress, err := GetAddressFromPrivateKey(privateKey)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.From = fromAddress
	// todo
	auth.GasLimit = uint64(9268689)
	auth.GasPrice = gasPrice
	return auth, nil
}
