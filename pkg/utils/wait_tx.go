package utils

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
)

func WaitTx(ctx context.Context, client *ethclient.Client, txn *ethtypes.Transaction, waitForSafePoll bool) error {
	receipt, err := bind.WaitMined(context.Background(), client, txn)
	if err != nil {
		return err
	}

	if waitForSafePoll {
		for {
			safeBlock, err := client.BlockByNumber(ctx, big.NewInt(int64(rpc.SafeBlockNumber)))
			if err != nil {
				log.Info("get safe block error: ", err.Error())
				continue
			}
			if safeBlock.NumberU64() < receipt.BlockNumber.Uint64() {
				log.Info("parent chain: waiting for safe block (see wait-for-tx-approval-safe-poll)", "waiting", receipt.BlockNumber.Uint64(), "safe", safeBlock.NumberU64())
				time.Sleep(time.Second * 50)
				continue
			}
		}
	}
	return nil

}
