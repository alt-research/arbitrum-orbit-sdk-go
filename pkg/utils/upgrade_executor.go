package utils

import (
	"crypto/ecdsa"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/renlulu/arbitrum-orbit-sdk-go/pkg/bindings"
	"github.com/renlulu/arbitrum-orbit-sdk-go/pkg/bindings/sequencerinboxgen"
)

func GenerateWhitelistDAKeySetCalldata(sequencerInbox common.Address, keyset []byte) ([]byte, error) {
	parsedABI, err := abi.JSON(strings.NewReader(sequencerinboxgen.SequencerInboxABI))
	if err != nil {
		return nil, err
	}
	calldata, err := parsedABI.Pack("setValidKeyset", keyset)
	if err != nil {
		return nil, err
	}
	return calldata, nil
}

func WhiteListDAKeySet(
	upgradeExecutor common.Address,
	sequencerInbox common.Address,
	url string,
	chainId *big.Int,
	key *ecdsa.PrivateKey,
	keyset []byte,
) (*types.Transaction, error) {
	calldata, err := GenerateWhitelistDAKeySetCalldata(sequencerInbox, keyset)
	if err != nil {
		return nil, err
	}
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainId)
	if err != nil {
		return nil, err
	}
	instance, err := bindings.NewUpgradeExecutor(upgradeExecutor, client)
	if err != nil {
		return nil, err
	}
	return instance.ExecuteCall(auth, sequencerInbox, calldata)

}
