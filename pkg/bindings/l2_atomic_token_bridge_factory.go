// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// L2RuntimeCode is an auto generated low-level Go binding around an user-defined struct.
type L2RuntimeCode struct {
	Router          []byte
	StandardGateway []byte
	CustomGateway   []byte
	WethGateway     []byte
	AeWeth          []byte
	UpgradeExecutor []byte
	Multicall       []byte
}

// L2AtomicTokenBridgeFactoryMetaData contains all meta data concerning the L2AtomicTokenBridgeFactory contract.
var L2AtomicTokenBridgeFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"L2AtomicTokenBridgeFactory_AlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"router\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"standardGateway\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"customGateway\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"wethGateway\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"aeWeth\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"upgradeExecutor\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"multicall\",\"type\":\"bytes\"}],\"internalType\":\"structL2RuntimeCode\",\"name\":\"l2Code\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"l1Router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l1StandardGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l1CustomGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l1WethGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l1Weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l2StandardGatewayCanonicalAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rollupOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"aliasedL1UpgradeExecutor\",\"type\":\"address\"}],\"name\":\"deployL2Contracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// L2AtomicTokenBridgeFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use L2AtomicTokenBridgeFactoryMetaData.ABI instead.
var L2AtomicTokenBridgeFactoryABI = L2AtomicTokenBridgeFactoryMetaData.ABI

// L2AtomicTokenBridgeFactory is an auto generated Go binding around an Ethereum contract.
type L2AtomicTokenBridgeFactory struct {
	L2AtomicTokenBridgeFactoryCaller     // Read-only binding to the contract
	L2AtomicTokenBridgeFactoryTransactor // Write-only binding to the contract
	L2AtomicTokenBridgeFactoryFilterer   // Log filterer for contract events
}

// L2AtomicTokenBridgeFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2AtomicTokenBridgeFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2AtomicTokenBridgeFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2AtomicTokenBridgeFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2AtomicTokenBridgeFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2AtomicTokenBridgeFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2AtomicTokenBridgeFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2AtomicTokenBridgeFactorySession struct {
	Contract     *L2AtomicTokenBridgeFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// L2AtomicTokenBridgeFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2AtomicTokenBridgeFactoryCallerSession struct {
	Contract *L2AtomicTokenBridgeFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// L2AtomicTokenBridgeFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2AtomicTokenBridgeFactoryTransactorSession struct {
	Contract     *L2AtomicTokenBridgeFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// L2AtomicTokenBridgeFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2AtomicTokenBridgeFactoryRaw struct {
	Contract *L2AtomicTokenBridgeFactory // Generic contract binding to access the raw methods on
}

// L2AtomicTokenBridgeFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2AtomicTokenBridgeFactoryCallerRaw struct {
	Contract *L2AtomicTokenBridgeFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// L2AtomicTokenBridgeFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2AtomicTokenBridgeFactoryTransactorRaw struct {
	Contract *L2AtomicTokenBridgeFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2AtomicTokenBridgeFactory creates a new instance of L2AtomicTokenBridgeFactory, bound to a specific deployed contract.
func NewL2AtomicTokenBridgeFactory(address common.Address, backend bind.ContractBackend) (*L2AtomicTokenBridgeFactory, error) {
	contract, err := bindL2AtomicTokenBridgeFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2AtomicTokenBridgeFactory{L2AtomicTokenBridgeFactoryCaller: L2AtomicTokenBridgeFactoryCaller{contract: contract}, L2AtomicTokenBridgeFactoryTransactor: L2AtomicTokenBridgeFactoryTransactor{contract: contract}, L2AtomicTokenBridgeFactoryFilterer: L2AtomicTokenBridgeFactoryFilterer{contract: contract}}, nil
}

// NewL2AtomicTokenBridgeFactoryCaller creates a new read-only instance of L2AtomicTokenBridgeFactory, bound to a specific deployed contract.
func NewL2AtomicTokenBridgeFactoryCaller(address common.Address, caller bind.ContractCaller) (*L2AtomicTokenBridgeFactoryCaller, error) {
	contract, err := bindL2AtomicTokenBridgeFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2AtomicTokenBridgeFactoryCaller{contract: contract}, nil
}

// NewL2AtomicTokenBridgeFactoryTransactor creates a new write-only instance of L2AtomicTokenBridgeFactory, bound to a specific deployed contract.
func NewL2AtomicTokenBridgeFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*L2AtomicTokenBridgeFactoryTransactor, error) {
	contract, err := bindL2AtomicTokenBridgeFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2AtomicTokenBridgeFactoryTransactor{contract: contract}, nil
}

// NewL2AtomicTokenBridgeFactoryFilterer creates a new log filterer instance of L2AtomicTokenBridgeFactory, bound to a specific deployed contract.
func NewL2AtomicTokenBridgeFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*L2AtomicTokenBridgeFactoryFilterer, error) {
	contract, err := bindL2AtomicTokenBridgeFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2AtomicTokenBridgeFactoryFilterer{contract: contract}, nil
}

// bindL2AtomicTokenBridgeFactory binds a generic wrapper to an already deployed contract.
func bindL2AtomicTokenBridgeFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2AtomicTokenBridgeFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2AtomicTokenBridgeFactory *L2AtomicTokenBridgeFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2AtomicTokenBridgeFactory.Contract.L2AtomicTokenBridgeFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2AtomicTokenBridgeFactory *L2AtomicTokenBridgeFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2AtomicTokenBridgeFactory.Contract.L2AtomicTokenBridgeFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2AtomicTokenBridgeFactory *L2AtomicTokenBridgeFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2AtomicTokenBridgeFactory.Contract.L2AtomicTokenBridgeFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2AtomicTokenBridgeFactory *L2AtomicTokenBridgeFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2AtomicTokenBridgeFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2AtomicTokenBridgeFactory *L2AtomicTokenBridgeFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2AtomicTokenBridgeFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2AtomicTokenBridgeFactory *L2AtomicTokenBridgeFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2AtomicTokenBridgeFactory.Contract.contract.Transact(opts, method, params...)
}

// DeployL2Contracts is a paid mutator transaction binding the contract method 0xb1c7a870.
//
// Solidity: function deployL2Contracts((bytes,bytes,bytes,bytes,bytes,bytes,bytes) l2Code, address l1Router, address l1StandardGateway, address l1CustomGateway, address l1WethGateway, address l1Weth, address l2StandardGatewayCanonicalAddress, address rollupOwner, address aliasedL1UpgradeExecutor) returns()
func (_L2AtomicTokenBridgeFactory *L2AtomicTokenBridgeFactoryTransactor) DeployL2Contracts(opts *bind.TransactOpts, l2Code L2RuntimeCode, l1Router common.Address, l1StandardGateway common.Address, l1CustomGateway common.Address, l1WethGateway common.Address, l1Weth common.Address, l2StandardGatewayCanonicalAddress common.Address, rollupOwner common.Address, aliasedL1UpgradeExecutor common.Address) (*types.Transaction, error) {
	return _L2AtomicTokenBridgeFactory.contract.Transact(opts, "deployL2Contracts", l2Code, l1Router, l1StandardGateway, l1CustomGateway, l1WethGateway, l1Weth, l2StandardGatewayCanonicalAddress, rollupOwner, aliasedL1UpgradeExecutor)
}

// DeployL2Contracts is a paid mutator transaction binding the contract method 0xb1c7a870.
//
// Solidity: function deployL2Contracts((bytes,bytes,bytes,bytes,bytes,bytes,bytes) l2Code, address l1Router, address l1StandardGateway, address l1CustomGateway, address l1WethGateway, address l1Weth, address l2StandardGatewayCanonicalAddress, address rollupOwner, address aliasedL1UpgradeExecutor) returns()
func (_L2AtomicTokenBridgeFactory *L2AtomicTokenBridgeFactorySession) DeployL2Contracts(l2Code L2RuntimeCode, l1Router common.Address, l1StandardGateway common.Address, l1CustomGateway common.Address, l1WethGateway common.Address, l1Weth common.Address, l2StandardGatewayCanonicalAddress common.Address, rollupOwner common.Address, aliasedL1UpgradeExecutor common.Address) (*types.Transaction, error) {
	return _L2AtomicTokenBridgeFactory.Contract.DeployL2Contracts(&_L2AtomicTokenBridgeFactory.TransactOpts, l2Code, l1Router, l1StandardGateway, l1CustomGateway, l1WethGateway, l1Weth, l2StandardGatewayCanonicalAddress, rollupOwner, aliasedL1UpgradeExecutor)
}

// DeployL2Contracts is a paid mutator transaction binding the contract method 0xb1c7a870.
//
// Solidity: function deployL2Contracts((bytes,bytes,bytes,bytes,bytes,bytes,bytes) l2Code, address l1Router, address l1StandardGateway, address l1CustomGateway, address l1WethGateway, address l1Weth, address l2StandardGatewayCanonicalAddress, address rollupOwner, address aliasedL1UpgradeExecutor) returns()
func (_L2AtomicTokenBridgeFactory *L2AtomicTokenBridgeFactoryTransactorSession) DeployL2Contracts(l2Code L2RuntimeCode, l1Router common.Address, l1StandardGateway common.Address, l1CustomGateway common.Address, l1WethGateway common.Address, l1Weth common.Address, l2StandardGatewayCanonicalAddress common.Address, rollupOwner common.Address, aliasedL1UpgradeExecutor common.Address) (*types.Transaction, error) {
	return _L2AtomicTokenBridgeFactory.Contract.DeployL2Contracts(&_L2AtomicTokenBridgeFactory.TransactOpts, l2Code, l1Router, l1StandardGateway, l1CustomGateway, l1WethGateway, l1Weth, l2StandardGatewayCanonicalAddress, rollupOwner, aliasedL1UpgradeExecutor)
}
