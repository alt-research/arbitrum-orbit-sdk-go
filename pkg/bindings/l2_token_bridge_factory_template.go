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

// L2TokenBridgeFactoryTemplateMetaData contains all meta data concerning the L2TokenBridgeFactoryTemplate contract.
var L2TokenBridgeFactoryTemplateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"L2AtomicTokenBridgeFactory_AlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"router\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"standardGateway\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"customGateway\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"wethGateway\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"aeWeth\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"upgradeExecutor\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"multicall\",\"type\":\"bytes\"}],\"internalType\":\"structL2RuntimeCode\",\"name\":\"l2Code\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"l1Router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l1StandardGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l1CustomGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l1WethGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l1Weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l2StandardGatewayCanonicalAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rollupOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"aliasedL1UpgradeExecutor\",\"type\":\"address\"}],\"name\":\"deployL2Contracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// L2TokenBridgeFactoryTemplateABI is the input ABI used to generate the binding from.
// Deprecated: Use L2TokenBridgeFactoryTemplateMetaData.ABI instead.
var L2TokenBridgeFactoryTemplateABI = L2TokenBridgeFactoryTemplateMetaData.ABI

// L2TokenBridgeFactoryTemplate is an auto generated Go binding around an Ethereum contract.
type L2TokenBridgeFactoryTemplate struct {
	L2TokenBridgeFactoryTemplateCaller     // Read-only binding to the contract
	L2TokenBridgeFactoryTemplateTransactor // Write-only binding to the contract
	L2TokenBridgeFactoryTemplateFilterer   // Log filterer for contract events
}

// L2TokenBridgeFactoryTemplateCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2TokenBridgeFactoryTemplateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2TokenBridgeFactoryTemplateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2TokenBridgeFactoryTemplateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2TokenBridgeFactoryTemplateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2TokenBridgeFactoryTemplateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2TokenBridgeFactoryTemplateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2TokenBridgeFactoryTemplateSession struct {
	Contract     *L2TokenBridgeFactoryTemplate // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// L2TokenBridgeFactoryTemplateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2TokenBridgeFactoryTemplateCallerSession struct {
	Contract *L2TokenBridgeFactoryTemplateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// L2TokenBridgeFactoryTemplateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2TokenBridgeFactoryTemplateTransactorSession struct {
	Contract     *L2TokenBridgeFactoryTemplateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// L2TokenBridgeFactoryTemplateRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2TokenBridgeFactoryTemplateRaw struct {
	Contract *L2TokenBridgeFactoryTemplate // Generic contract binding to access the raw methods on
}

// L2TokenBridgeFactoryTemplateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2TokenBridgeFactoryTemplateCallerRaw struct {
	Contract *L2TokenBridgeFactoryTemplateCaller // Generic read-only contract binding to access the raw methods on
}

// L2TokenBridgeFactoryTemplateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2TokenBridgeFactoryTemplateTransactorRaw struct {
	Contract *L2TokenBridgeFactoryTemplateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2TokenBridgeFactoryTemplate creates a new instance of L2TokenBridgeFactoryTemplate, bound to a specific deployed contract.
func NewL2TokenBridgeFactoryTemplate(address common.Address, backend bind.ContractBackend) (*L2TokenBridgeFactoryTemplate, error) {
	contract, err := bindL2TokenBridgeFactoryTemplate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2TokenBridgeFactoryTemplate{L2TokenBridgeFactoryTemplateCaller: L2TokenBridgeFactoryTemplateCaller{contract: contract}, L2TokenBridgeFactoryTemplateTransactor: L2TokenBridgeFactoryTemplateTransactor{contract: contract}, L2TokenBridgeFactoryTemplateFilterer: L2TokenBridgeFactoryTemplateFilterer{contract: contract}}, nil
}

// NewL2TokenBridgeFactoryTemplateCaller creates a new read-only instance of L2TokenBridgeFactoryTemplate, bound to a specific deployed contract.
func NewL2TokenBridgeFactoryTemplateCaller(address common.Address, caller bind.ContractCaller) (*L2TokenBridgeFactoryTemplateCaller, error) {
	contract, err := bindL2TokenBridgeFactoryTemplate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2TokenBridgeFactoryTemplateCaller{contract: contract}, nil
}

// NewL2TokenBridgeFactoryTemplateTransactor creates a new write-only instance of L2TokenBridgeFactoryTemplate, bound to a specific deployed contract.
func NewL2TokenBridgeFactoryTemplateTransactor(address common.Address, transactor bind.ContractTransactor) (*L2TokenBridgeFactoryTemplateTransactor, error) {
	contract, err := bindL2TokenBridgeFactoryTemplate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2TokenBridgeFactoryTemplateTransactor{contract: contract}, nil
}

// NewL2TokenBridgeFactoryTemplateFilterer creates a new log filterer instance of L2TokenBridgeFactoryTemplate, bound to a specific deployed contract.
func NewL2TokenBridgeFactoryTemplateFilterer(address common.Address, filterer bind.ContractFilterer) (*L2TokenBridgeFactoryTemplateFilterer, error) {
	contract, err := bindL2TokenBridgeFactoryTemplate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2TokenBridgeFactoryTemplateFilterer{contract: contract}, nil
}

// bindL2TokenBridgeFactoryTemplate binds a generic wrapper to an already deployed contract.
func bindL2TokenBridgeFactoryTemplate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2TokenBridgeFactoryTemplateMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2TokenBridgeFactoryTemplate *L2TokenBridgeFactoryTemplateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2TokenBridgeFactoryTemplate.Contract.L2TokenBridgeFactoryTemplateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2TokenBridgeFactoryTemplate *L2TokenBridgeFactoryTemplateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2TokenBridgeFactoryTemplate.Contract.L2TokenBridgeFactoryTemplateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2TokenBridgeFactoryTemplate *L2TokenBridgeFactoryTemplateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2TokenBridgeFactoryTemplate.Contract.L2TokenBridgeFactoryTemplateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2TokenBridgeFactoryTemplate *L2TokenBridgeFactoryTemplateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2TokenBridgeFactoryTemplate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2TokenBridgeFactoryTemplate *L2TokenBridgeFactoryTemplateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2TokenBridgeFactoryTemplate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2TokenBridgeFactoryTemplate *L2TokenBridgeFactoryTemplateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2TokenBridgeFactoryTemplate.Contract.contract.Transact(opts, method, params...)
}

// DeployL2Contracts is a paid mutator transaction binding the contract method 0xb1c7a870.
//
// Solidity: function deployL2Contracts((bytes,bytes,bytes,bytes,bytes,bytes,bytes) l2Code, address l1Router, address l1StandardGateway, address l1CustomGateway, address l1WethGateway, address l1Weth, address l2StandardGatewayCanonicalAddress, address rollupOwner, address aliasedL1UpgradeExecutor) returns()
func (_L2TokenBridgeFactoryTemplate *L2TokenBridgeFactoryTemplateTransactor) DeployL2Contracts(opts *bind.TransactOpts, l2Code L2RuntimeCode, l1Router common.Address, l1StandardGateway common.Address, l1CustomGateway common.Address, l1WethGateway common.Address, l1Weth common.Address, l2StandardGatewayCanonicalAddress common.Address, rollupOwner common.Address, aliasedL1UpgradeExecutor common.Address) (*types.Transaction, error) {
	return _L2TokenBridgeFactoryTemplate.contract.Transact(opts, "deployL2Contracts", l2Code, l1Router, l1StandardGateway, l1CustomGateway, l1WethGateway, l1Weth, l2StandardGatewayCanonicalAddress, rollupOwner, aliasedL1UpgradeExecutor)
}

// DeployL2Contracts is a paid mutator transaction binding the contract method 0xb1c7a870.
//
// Solidity: function deployL2Contracts((bytes,bytes,bytes,bytes,bytes,bytes,bytes) l2Code, address l1Router, address l1StandardGateway, address l1CustomGateway, address l1WethGateway, address l1Weth, address l2StandardGatewayCanonicalAddress, address rollupOwner, address aliasedL1UpgradeExecutor) returns()
func (_L2TokenBridgeFactoryTemplate *L2TokenBridgeFactoryTemplateSession) DeployL2Contracts(l2Code L2RuntimeCode, l1Router common.Address, l1StandardGateway common.Address, l1CustomGateway common.Address, l1WethGateway common.Address, l1Weth common.Address, l2StandardGatewayCanonicalAddress common.Address, rollupOwner common.Address, aliasedL1UpgradeExecutor common.Address) (*types.Transaction, error) {
	return _L2TokenBridgeFactoryTemplate.Contract.DeployL2Contracts(&_L2TokenBridgeFactoryTemplate.TransactOpts, l2Code, l1Router, l1StandardGateway, l1CustomGateway, l1WethGateway, l1Weth, l2StandardGatewayCanonicalAddress, rollupOwner, aliasedL1UpgradeExecutor)
}

// DeployL2Contracts is a paid mutator transaction binding the contract method 0xb1c7a870.
//
// Solidity: function deployL2Contracts((bytes,bytes,bytes,bytes,bytes,bytes,bytes) l2Code, address l1Router, address l1StandardGateway, address l1CustomGateway, address l1WethGateway, address l1Weth, address l2StandardGatewayCanonicalAddress, address rollupOwner, address aliasedL1UpgradeExecutor) returns()
func (_L2TokenBridgeFactoryTemplate *L2TokenBridgeFactoryTemplateTransactorSession) DeployL2Contracts(l2Code L2RuntimeCode, l1Router common.Address, l1StandardGateway common.Address, l1CustomGateway common.Address, l1WethGateway common.Address, l1Weth common.Address, l2StandardGatewayCanonicalAddress common.Address, rollupOwner common.Address, aliasedL1UpgradeExecutor common.Address) (*types.Transaction, error) {
	return _L2TokenBridgeFactoryTemplate.Contract.DeployL2Contracts(&_L2TokenBridgeFactoryTemplate.TransactOpts, l2Code, l1Router, l1StandardGateway, l1CustomGateway, l1WethGateway, l1Weth, l2StandardGatewayCanonicalAddress, rollupOwner, aliasedL1UpgradeExecutor)
}
