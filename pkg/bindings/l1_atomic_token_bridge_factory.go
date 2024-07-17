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

// L1AtomicTokenBridgeCreatorL1Templates is an auto generated low-level Go binding around an user-defined struct.
type L1AtomicTokenBridgeCreatorL1Templates struct {
	RouterTemplate                       common.Address
	StandardGatewayTemplate              common.Address
	CustomGatewayTemplate                common.Address
	WethGatewayTemplate                  common.Address
	FeeTokenBasedRouterTemplate          common.Address
	FeeTokenBasedStandardGatewayTemplate common.Address
	FeeTokenBasedCustomGatewayTemplate   common.Address
	UpgradeExecutor                      common.Address
}

// L1DeploymentAddresses is an auto generated low-level Go binding around an user-defined struct.
type L1DeploymentAddresses struct {
	Router          common.Address
	StandardGateway common.Address
	CustomGateway   common.Address
	WethGateway     common.Address
	Weth            common.Address
}

// L2DeploymentAddresses is an auto generated low-level Go binding around an user-defined struct.
type L2DeploymentAddresses struct {
	Router             common.Address
	StandardGateway    common.Address
	CustomGateway      common.Address
	WethGateway        common.Address
	Weth               common.Address
	ProxyAdmin         common.Address
	BeaconProxyFactory common.Address
	UpgradeExecutor    common.Address
	Multicall          common.Address
}

// L1AtomicTokenBridgeFactoryMetaData contains all meta data concerning the L1AtomicTokenBridgeFactory contract.
var L1AtomicTokenBridgeFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"L1AtomicTokenBridgeCreator_L2FactoryCannotBeChanged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1AtomicTokenBridgeCreator_OnlyRollupOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1AtomicTokenBridgeCreator_ProxyAdminNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1AtomicTokenBridgeCreator_RollupOwnershipMisconfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1AtomicTokenBridgeCreator_TemplatesNotSet\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"standardGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"customGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wethGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"weth\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structL1DeploymentAddresses\",\"name\":\"l1Deployment\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"standardGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"customGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wethGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proxyAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beaconProxyFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"upgradeExecutor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"multicall\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structL2DeploymentAddresses\",\"name\":\"l2Deployment\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxyAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"upgradeExecutor\",\"type\":\"address\"}],\"name\":\"OrbitTokenBridgeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"standardGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"customGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wethGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"weth\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structL1DeploymentAddresses\",\"name\":\"l1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"standardGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"customGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wethGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proxyAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beaconProxyFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"upgradeExecutor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"multicall\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structL2DeploymentAddresses\",\"name\":\"l2\",\"type\":\"tuple\"}],\"name\":\"OrbitTokenBridgeDeploymentSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"OrbitTokenBridgeTemplatesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"canonicalL2FactoryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rollupOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxGasForContracts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"}],\"name\":\"createTokenBridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasLimitForL2FactoryDeployment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"}],\"name\":\"getRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"inboxToL1Deployment\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"standardGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"customGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wethGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"weth\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"inboxToL2Deployment\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"standardGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"customGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wethGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proxyAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beaconProxyFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"upgradeExecutor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"multicall\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractL1TokenBridgeRetryableSender\",\"name\":\"_retryableSender\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Multicall\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Templates\",\"outputs\":[{\"internalType\":\"contractL1GatewayRouter\",\"name\":\"routerTemplate\",\"type\":\"address\"},{\"internalType\":\"contractL1ERC20Gateway\",\"name\":\"standardGatewayTemplate\",\"type\":\"address\"},{\"internalType\":\"contractL1CustomGateway\",\"name\":\"customGatewayTemplate\",\"type\":\"address\"},{\"internalType\":\"contractL1WethGateway\",\"name\":\"wethGatewayTemplate\",\"type\":\"address\"},{\"internalType\":\"contractL1OrbitGatewayRouter\",\"name\":\"feeTokenBasedRouterTemplate\",\"type\":\"address\"},{\"internalType\":\"contractL1OrbitERC20Gateway\",\"name\":\"feeTokenBasedStandardGatewayTemplate\",\"type\":\"address\"},{\"internalType\":\"contractL1OrbitCustomGateway\",\"name\":\"feeTokenBasedCustomGatewayTemplate\",\"type\":\"address\"},{\"internalType\":\"contractIUpgradeExecutor\",\"name\":\"upgradeExecutor\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2CustomGatewayTemplate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2MulticallTemplate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2RouterTemplate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2StandardGatewayTemplate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2TokenBridgeFactoryTemplate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2WethGatewayTemplate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2WethTemplate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"retryableSender\",\"outputs\":[{\"internalType\":\"contractL1TokenBridgeRetryableSender\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"standardGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"customGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wethGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"weth\",\"type\":\"address\"}],\"internalType\":\"structL1DeploymentAddresses\",\"name\":\"l1Deployment\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"standardGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"customGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"wethGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proxyAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beaconProxyFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"upgradeExecutor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"multicall\",\"type\":\"address\"}],\"internalType\":\"structL2DeploymentAddresses\",\"name\":\"l2Deployment\",\"type\":\"tuple\"}],\"name\":\"setDeployment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractL1GatewayRouter\",\"name\":\"routerTemplate\",\"type\":\"address\"},{\"internalType\":\"contractL1ERC20Gateway\",\"name\":\"standardGatewayTemplate\",\"type\":\"address\"},{\"internalType\":\"contractL1CustomGateway\",\"name\":\"customGatewayTemplate\",\"type\":\"address\"},{\"internalType\":\"contractL1WethGateway\",\"name\":\"wethGatewayTemplate\",\"type\":\"address\"},{\"internalType\":\"contractL1OrbitGatewayRouter\",\"name\":\"feeTokenBasedRouterTemplate\",\"type\":\"address\"},{\"internalType\":\"contractL1OrbitERC20Gateway\",\"name\":\"feeTokenBasedStandardGatewayTemplate\",\"type\":\"address\"},{\"internalType\":\"contractL1OrbitCustomGateway\",\"name\":\"feeTokenBasedCustomGatewayTemplate\",\"type\":\"address\"},{\"internalType\":\"contractIUpgradeExecutor\",\"name\":\"upgradeExecutor\",\"type\":\"address\"}],\"internalType\":\"structL1AtomicTokenBridgeCreator.L1Templates\",\"name\":\"_l1Templates\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_l2TokenBridgeFactoryTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2RouterTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2StandardGatewayTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2CustomGatewayTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2WethGatewayTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2WethTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2MulticallTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Multicall\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimitForL2FactoryDeployment\",\"type\":\"uint256\"}],\"name\":\"setTemplates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// L1AtomicTokenBridgeFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use L1AtomicTokenBridgeFactoryMetaData.ABI instead.
var L1AtomicTokenBridgeFactoryABI = L1AtomicTokenBridgeFactoryMetaData.ABI

// L1AtomicTokenBridgeFactory is an auto generated Go binding around an Ethereum contract.
type L1AtomicTokenBridgeFactory struct {
	L1AtomicTokenBridgeFactoryCaller     // Read-only binding to the contract
	L1AtomicTokenBridgeFactoryTransactor // Write-only binding to the contract
	L1AtomicTokenBridgeFactoryFilterer   // Log filterer for contract events
}

// L1AtomicTokenBridgeFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1AtomicTokenBridgeFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1AtomicTokenBridgeFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1AtomicTokenBridgeFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1AtomicTokenBridgeFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1AtomicTokenBridgeFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1AtomicTokenBridgeFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1AtomicTokenBridgeFactorySession struct {
	Contract     *L1AtomicTokenBridgeFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// L1AtomicTokenBridgeFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1AtomicTokenBridgeFactoryCallerSession struct {
	Contract *L1AtomicTokenBridgeFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// L1AtomicTokenBridgeFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1AtomicTokenBridgeFactoryTransactorSession struct {
	Contract     *L1AtomicTokenBridgeFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// L1AtomicTokenBridgeFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1AtomicTokenBridgeFactoryRaw struct {
	Contract *L1AtomicTokenBridgeFactory // Generic contract binding to access the raw methods on
}

// L1AtomicTokenBridgeFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1AtomicTokenBridgeFactoryCallerRaw struct {
	Contract *L1AtomicTokenBridgeFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// L1AtomicTokenBridgeFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1AtomicTokenBridgeFactoryTransactorRaw struct {
	Contract *L1AtomicTokenBridgeFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1AtomicTokenBridgeFactory creates a new instance of L1AtomicTokenBridgeFactory, bound to a specific deployed contract.
func NewL1AtomicTokenBridgeFactory(address common.Address, backend bind.ContractBackend) (*L1AtomicTokenBridgeFactory, error) {
	contract, err := bindL1AtomicTokenBridgeFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1AtomicTokenBridgeFactory{L1AtomicTokenBridgeFactoryCaller: L1AtomicTokenBridgeFactoryCaller{contract: contract}, L1AtomicTokenBridgeFactoryTransactor: L1AtomicTokenBridgeFactoryTransactor{contract: contract}, L1AtomicTokenBridgeFactoryFilterer: L1AtomicTokenBridgeFactoryFilterer{contract: contract}}, nil
}

// NewL1AtomicTokenBridgeFactoryCaller creates a new read-only instance of L1AtomicTokenBridgeFactory, bound to a specific deployed contract.
func NewL1AtomicTokenBridgeFactoryCaller(address common.Address, caller bind.ContractCaller) (*L1AtomicTokenBridgeFactoryCaller, error) {
	contract, err := bindL1AtomicTokenBridgeFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1AtomicTokenBridgeFactoryCaller{contract: contract}, nil
}

// NewL1AtomicTokenBridgeFactoryTransactor creates a new write-only instance of L1AtomicTokenBridgeFactory, bound to a specific deployed contract.
func NewL1AtomicTokenBridgeFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*L1AtomicTokenBridgeFactoryTransactor, error) {
	contract, err := bindL1AtomicTokenBridgeFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1AtomicTokenBridgeFactoryTransactor{contract: contract}, nil
}

// NewL1AtomicTokenBridgeFactoryFilterer creates a new log filterer instance of L1AtomicTokenBridgeFactory, bound to a specific deployed contract.
func NewL1AtomicTokenBridgeFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*L1AtomicTokenBridgeFactoryFilterer, error) {
	contract, err := bindL1AtomicTokenBridgeFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1AtomicTokenBridgeFactoryFilterer{contract: contract}, nil
}

// bindL1AtomicTokenBridgeFactory binds a generic wrapper to an already deployed contract.
func bindL1AtomicTokenBridgeFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L1AtomicTokenBridgeFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1AtomicTokenBridgeFactory.Contract.L1AtomicTokenBridgeFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1AtomicTokenBridgeFactory.Contract.L1AtomicTokenBridgeFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1AtomicTokenBridgeFactory.Contract.L1AtomicTokenBridgeFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1AtomicTokenBridgeFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1AtomicTokenBridgeFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1AtomicTokenBridgeFactory.Contract.contract.Transact(opts, method, params...)
}

// CanonicalL2FactoryAddress is a free data retrieval call binding the contract method 0xbfd3e518.
//
// Solidity: function canonicalL2FactoryAddress() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCaller) CanonicalL2FactoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1AtomicTokenBridgeFactory.contract.Call(opts, &out, "canonicalL2FactoryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CanonicalL2FactoryAddress is a free data retrieval call binding the contract method 0xbfd3e518.
//
// Solidity: function canonicalL2FactoryAddress() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactorySession) CanonicalL2FactoryAddress() (common.Address, error) {
	return _L1AtomicTokenBridgeFactory.Contract.CanonicalL2FactoryAddress(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// CanonicalL2FactoryAddress is a free data retrieval call binding the contract method 0xbfd3e518.
//
// Solidity: function canonicalL2FactoryAddress() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCallerSession) CanonicalL2FactoryAddress() (common.Address, error) {
	return _L1AtomicTokenBridgeFactory.Contract.CanonicalL2FactoryAddress(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// GasLimitForL2FactoryDeployment is a free data retrieval call binding the contract method 0x888139d4.
//
// Solidity: function gasLimitForL2FactoryDeployment() view returns(uint256)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCaller) GasLimitForL2FactoryDeployment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1AtomicTokenBridgeFactory.contract.Call(opts, &out, "gasLimitForL2FactoryDeployment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasLimitForL2FactoryDeployment is a free data retrieval call binding the contract method 0x888139d4.
//
// Solidity: function gasLimitForL2FactoryDeployment() view returns(uint256)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactorySession) GasLimitForL2FactoryDeployment() (*big.Int, error) {
	return _L1AtomicTokenBridgeFactory.Contract.GasLimitForL2FactoryDeployment(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// GasLimitForL2FactoryDeployment is a free data retrieval call binding the contract method 0x888139d4.
//
// Solidity: function gasLimitForL2FactoryDeployment() view returns(uint256)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCallerSession) GasLimitForL2FactoryDeployment() (*big.Int, error) {
	return _L1AtomicTokenBridgeFactory.Contract.GasLimitForL2FactoryDeployment(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// GetRouter is a free data retrieval call binding the contract method 0x8369166d.
//
// Solidity: function getRouter(address inbox) view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCaller) GetRouter(opts *bind.CallOpts, inbox common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1AtomicTokenBridgeFactory.contract.Call(opts, &out, "getRouter", inbox)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRouter is a free data retrieval call binding the contract method 0x8369166d.
//
// Solidity: function getRouter(address inbox) view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactorySession) GetRouter(inbox common.Address) (common.Address, error) {
	return _L1AtomicTokenBridgeFactory.Contract.GetRouter(&_L1AtomicTokenBridgeFactory.CallOpts, inbox)
}

// GetRouter is a free data retrieval call binding the contract method 0x8369166d.
//
// Solidity: function getRouter(address inbox) view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCallerSession) GetRouter(inbox common.Address) (common.Address, error) {
	return _L1AtomicTokenBridgeFactory.Contract.GetRouter(&_L1AtomicTokenBridgeFactory.CallOpts, inbox)
}

// InboxToL1Deployment is a free data retrieval call binding the contract method 0xd9ce0ef9.
//
// Solidity: function inboxToL1Deployment(address ) view returns(address router, address standardGateway, address customGateway, address wethGateway, address weth)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCaller) InboxToL1Deployment(opts *bind.CallOpts, arg0 common.Address) (struct {
	Router          common.Address
	StandardGateway common.Address
	CustomGateway   common.Address
	WethGateway     common.Address
	Weth            common.Address
}, error) {
	var out []interface{}
	err := _L1AtomicTokenBridgeFactory.contract.Call(opts, &out, "inboxToL1Deployment", arg0)

	outstruct := new(struct {
		Router          common.Address
		StandardGateway common.Address
		CustomGateway   common.Address
		WethGateway     common.Address
		Weth            common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Router = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.StandardGateway = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.CustomGateway = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.WethGateway = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Weth = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// InboxToL1Deployment is a free data retrieval call binding the contract method 0xd9ce0ef9.
//
// Solidity: function inboxToL1Deployment(address ) view returns(address router, address standardGateway, address customGateway, address wethGateway, address weth)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactorySession) InboxToL1Deployment(arg0 common.Address) (struct {
	Router          common.Address
	StandardGateway common.Address
	CustomGateway   common.Address
	WethGateway     common.Address
	Weth            common.Address
}, error) {
	return _L1AtomicTokenBridgeFactory.Contract.InboxToL1Deployment(&_L1AtomicTokenBridgeFactory.CallOpts, arg0)
}

// InboxToL1Deployment is a free data retrieval call binding the contract method 0xd9ce0ef9.
//
// Solidity: function inboxToL1Deployment(address ) view returns(address router, address standardGateway, address customGateway, address wethGateway, address weth)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCallerSession) InboxToL1Deployment(arg0 common.Address) (struct {
	Router          common.Address
	StandardGateway common.Address
	CustomGateway   common.Address
	WethGateway     common.Address
	Weth            common.Address
}, error) {
	return _L1AtomicTokenBridgeFactory.Contract.InboxToL1Deployment(&_L1AtomicTokenBridgeFactory.CallOpts, arg0)
}

// InboxToL2Deployment is a free data retrieval call binding the contract method 0x46052706.
//
// Solidity: function inboxToL2Deployment(address ) view returns(address router, address standardGateway, address customGateway, address wethGateway, address weth, address proxyAdmin, address beaconProxyFactory, address upgradeExecutor, address multicall)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCaller) InboxToL2Deployment(opts *bind.CallOpts, arg0 common.Address) (struct {
	Router             common.Address
	StandardGateway    common.Address
	CustomGateway      common.Address
	WethGateway        common.Address
	Weth               common.Address
	ProxyAdmin         common.Address
	BeaconProxyFactory common.Address
	UpgradeExecutor    common.Address
	Multicall          common.Address
}, error) {
	var out []interface{}
	err := _L1AtomicTokenBridgeFactory.contract.Call(opts, &out, "inboxToL2Deployment", arg0)

	outstruct := new(struct {
		Router             common.Address
		StandardGateway    common.Address
		CustomGateway      common.Address
		WethGateway        common.Address
		Weth               common.Address
		ProxyAdmin         common.Address
		BeaconProxyFactory common.Address
		UpgradeExecutor    common.Address
		Multicall          common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Router = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.StandardGateway = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.CustomGateway = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.WethGateway = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Weth = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.ProxyAdmin = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.BeaconProxyFactory = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.UpgradeExecutor = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.Multicall = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// InboxToL2Deployment is a free data retrieval call binding the contract method 0x46052706.
//
// Solidity: function inboxToL2Deployment(address ) view returns(address router, address standardGateway, address customGateway, address wethGateway, address weth, address proxyAdmin, address beaconProxyFactory, address upgradeExecutor, address multicall)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactorySession) InboxToL2Deployment(arg0 common.Address) (struct {
	Router             common.Address
	StandardGateway    common.Address
	CustomGateway      common.Address
	WethGateway        common.Address
	Weth               common.Address
	ProxyAdmin         common.Address
	BeaconProxyFactory common.Address
	UpgradeExecutor    common.Address
	Multicall          common.Address
}, error) {
	return _L1AtomicTokenBridgeFactory.Contract.InboxToL2Deployment(&_L1AtomicTokenBridgeFactory.CallOpts, arg0)
}

// InboxToL2Deployment is a free data retrieval call binding the contract method 0x46052706.
//
// Solidity: function inboxToL2Deployment(address ) view returns(address router, address standardGateway, address customGateway, address wethGateway, address weth, address proxyAdmin, address beaconProxyFactory, address upgradeExecutor, address multicall)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCallerSession) InboxToL2Deployment(arg0 common.Address) (struct {
	Router             common.Address
	StandardGateway    common.Address
	CustomGateway      common.Address
	WethGateway        common.Address
	Weth               common.Address
	ProxyAdmin         common.Address
	BeaconProxyFactory common.Address
	UpgradeExecutor    common.Address
	Multicall          common.Address
}, error) {
	return _L1AtomicTokenBridgeFactory.Contract.InboxToL2Deployment(&_L1AtomicTokenBridgeFactory.CallOpts, arg0)
}

// L1Multicall is a free data retrieval call binding the contract method 0xb1460a71.
//
// Solidity: function l1Multicall() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCaller) L1Multicall(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1AtomicTokenBridgeFactory.contract.Call(opts, &out, "l1Multicall")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Multicall is a free data retrieval call binding the contract method 0xb1460a71.
//
// Solidity: function l1Multicall() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactorySession) L1Multicall() (common.Address, error) {
	return _L1AtomicTokenBridgeFactory.Contract.L1Multicall(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// L1Multicall is a free data retrieval call binding the contract method 0xb1460a71.
//
// Solidity: function l1Multicall() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCallerSession) L1Multicall() (common.Address, error) {
	return _L1AtomicTokenBridgeFactory.Contract.L1Multicall(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// L1Templates is a free data retrieval call binding the contract method 0xa5595da9.
//
// Solidity: function l1Templates() view returns(address routerTemplate, address standardGatewayTemplate, address customGatewayTemplate, address wethGatewayTemplate, address feeTokenBasedRouterTemplate, address feeTokenBasedStandardGatewayTemplate, address feeTokenBasedCustomGatewayTemplate, address upgradeExecutor)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCaller) L1Templates(opts *bind.CallOpts) (struct {
	RouterTemplate                       common.Address
	StandardGatewayTemplate              common.Address
	CustomGatewayTemplate                common.Address
	WethGatewayTemplate                  common.Address
	FeeTokenBasedRouterTemplate          common.Address
	FeeTokenBasedStandardGatewayTemplate common.Address
	FeeTokenBasedCustomGatewayTemplate   common.Address
	UpgradeExecutor                      common.Address
}, error) {
	var out []interface{}
	err := _L1AtomicTokenBridgeFactory.contract.Call(opts, &out, "l1Templates")

	outstruct := new(struct {
		RouterTemplate                       common.Address
		StandardGatewayTemplate              common.Address
		CustomGatewayTemplate                common.Address
		WethGatewayTemplate                  common.Address
		FeeTokenBasedRouterTemplate          common.Address
		FeeTokenBasedStandardGatewayTemplate common.Address
		FeeTokenBasedCustomGatewayTemplate   common.Address
		UpgradeExecutor                      common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RouterTemplate = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.StandardGatewayTemplate = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.CustomGatewayTemplate = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.WethGatewayTemplate = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.FeeTokenBasedRouterTemplate = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.FeeTokenBasedStandardGatewayTemplate = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.FeeTokenBasedCustomGatewayTemplate = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.UpgradeExecutor = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// L1Templates is a free data retrieval call binding the contract method 0xa5595da9.
//
// Solidity: function l1Templates() view returns(address routerTemplate, address standardGatewayTemplate, address customGatewayTemplate, address wethGatewayTemplate, address feeTokenBasedRouterTemplate, address feeTokenBasedStandardGatewayTemplate, address feeTokenBasedCustomGatewayTemplate, address upgradeExecutor)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactorySession) L1Templates() (struct {
	RouterTemplate                       common.Address
	StandardGatewayTemplate              common.Address
	CustomGatewayTemplate                common.Address
	WethGatewayTemplate                  common.Address
	FeeTokenBasedRouterTemplate          common.Address
	FeeTokenBasedStandardGatewayTemplate common.Address
	FeeTokenBasedCustomGatewayTemplate   common.Address
	UpgradeExecutor                      common.Address
}, error) {
	return _L1AtomicTokenBridgeFactory.Contract.L1Templates(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// L1Templates is a free data retrieval call binding the contract method 0xa5595da9.
//
// Solidity: function l1Templates() view returns(address routerTemplate, address standardGatewayTemplate, address customGatewayTemplate, address wethGatewayTemplate, address feeTokenBasedRouterTemplate, address feeTokenBasedStandardGatewayTemplate, address feeTokenBasedCustomGatewayTemplate, address upgradeExecutor)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCallerSession) L1Templates() (struct {
	RouterTemplate                       common.Address
	StandardGatewayTemplate              common.Address
	CustomGatewayTemplate                common.Address
	WethGatewayTemplate                  common.Address
	FeeTokenBasedRouterTemplate          common.Address
	FeeTokenBasedStandardGatewayTemplate common.Address
	FeeTokenBasedCustomGatewayTemplate   common.Address
	UpgradeExecutor                      common.Address
}, error) {
	return _L1AtomicTokenBridgeFactory.Contract.L1Templates(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// L1Weth is a free data retrieval call binding the contract method 0x146bf4b1.
//
// Solidity: function l1Weth() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCaller) L1Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1AtomicTokenBridgeFactory.contract.Call(opts, &out, "l1Weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Weth is a free data retrieval call binding the contract method 0x146bf4b1.
//
// Solidity: function l1Weth() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactorySession) L1Weth() (common.Address, error) {
	return _L1AtomicTokenBridgeFactory.Contract.L1Weth(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// L1Weth is a free data retrieval call binding the contract method 0x146bf4b1.
//
// Solidity: function l1Weth() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCallerSession) L1Weth() (common.Address, error) {
	return _L1AtomicTokenBridgeFactory.Contract.L1Weth(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// L2CustomGatewayTemplate is a free data retrieval call binding the contract method 0x41083186.
//
// Solidity: function l2CustomGatewayTemplate() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCaller) L2CustomGatewayTemplate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1AtomicTokenBridgeFactory.contract.Call(opts, &out, "l2CustomGatewayTemplate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2CustomGatewayTemplate is a free data retrieval call binding the contract method 0x41083186.
//
// Solidity: function l2CustomGatewayTemplate() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactorySession) L2CustomGatewayTemplate() (common.Address, error) {
	return _L1AtomicTokenBridgeFactory.Contract.L2CustomGatewayTemplate(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// L2CustomGatewayTemplate is a free data retrieval call binding the contract method 0x41083186.
//
// Solidity: function l2CustomGatewayTemplate() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCallerSession) L2CustomGatewayTemplate() (common.Address, error) {
	return _L1AtomicTokenBridgeFactory.Contract.L2CustomGatewayTemplate(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// L2MulticallTemplate is a free data retrieval call binding the contract method 0x8c99e31c.
//
// Solidity: function l2MulticallTemplate() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCaller) L2MulticallTemplate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1AtomicTokenBridgeFactory.contract.Call(opts, &out, "l2MulticallTemplate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2MulticallTemplate is a free data retrieval call binding the contract method 0x8c99e31c.
//
// Solidity: function l2MulticallTemplate() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactorySession) L2MulticallTemplate() (common.Address, error) {
	return _L1AtomicTokenBridgeFactory.Contract.L2MulticallTemplate(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// L2MulticallTemplate is a free data retrieval call binding the contract method 0x8c99e31c.
//
// Solidity: function l2MulticallTemplate() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCallerSession) L2MulticallTemplate() (common.Address, error) {
	return _L1AtomicTokenBridgeFactory.Contract.L2MulticallTemplate(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// L2RouterTemplate is a free data retrieval call binding the contract method 0x381c9d99.
//
// Solidity: function l2RouterTemplate() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCaller) L2RouterTemplate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1AtomicTokenBridgeFactory.contract.Call(opts, &out, "l2RouterTemplate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2RouterTemplate is a free data retrieval call binding the contract method 0x381c9d99.
//
// Solidity: function l2RouterTemplate() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactorySession) L2RouterTemplate() (common.Address, error) {
	return _L1AtomicTokenBridgeFactory.Contract.L2RouterTemplate(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// L2RouterTemplate is a free data retrieval call binding the contract method 0x381c9d99.
//
// Solidity: function l2RouterTemplate() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCallerSession) L2RouterTemplate() (common.Address, error) {
	return _L1AtomicTokenBridgeFactory.Contract.L2RouterTemplate(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// L2StandardGatewayTemplate is a free data retrieval call binding the contract method 0xd7eee6ca.
//
// Solidity: function l2StandardGatewayTemplate() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCaller) L2StandardGatewayTemplate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1AtomicTokenBridgeFactory.contract.Call(opts, &out, "l2StandardGatewayTemplate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2StandardGatewayTemplate is a free data retrieval call binding the contract method 0xd7eee6ca.
//
// Solidity: function l2StandardGatewayTemplate() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactorySession) L2StandardGatewayTemplate() (common.Address, error) {
	return _L1AtomicTokenBridgeFactory.Contract.L2StandardGatewayTemplate(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// L2StandardGatewayTemplate is a free data retrieval call binding the contract method 0xd7eee6ca.
//
// Solidity: function l2StandardGatewayTemplate() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCallerSession) L2StandardGatewayTemplate() (common.Address, error) {
	return _L1AtomicTokenBridgeFactory.Contract.L2StandardGatewayTemplate(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// L2TokenBridgeFactoryTemplate is a free data retrieval call binding the contract method 0x1aeef2e2.
//
// Solidity: function l2TokenBridgeFactoryTemplate() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCaller) L2TokenBridgeFactoryTemplate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1AtomicTokenBridgeFactory.contract.Call(opts, &out, "l2TokenBridgeFactoryTemplate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TokenBridgeFactoryTemplate is a free data retrieval call binding the contract method 0x1aeef2e2.
//
// Solidity: function l2TokenBridgeFactoryTemplate() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactorySession) L2TokenBridgeFactoryTemplate() (common.Address, error) {
	return _L1AtomicTokenBridgeFactory.Contract.L2TokenBridgeFactoryTemplate(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// L2TokenBridgeFactoryTemplate is a free data retrieval call binding the contract method 0x1aeef2e2.
//
// Solidity: function l2TokenBridgeFactoryTemplate() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCallerSession) L2TokenBridgeFactoryTemplate() (common.Address, error) {
	return _L1AtomicTokenBridgeFactory.Contract.L2TokenBridgeFactoryTemplate(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// L2WethGatewayTemplate is a free data retrieval call binding the contract method 0x9095765e.
//
// Solidity: function l2WethGatewayTemplate() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCaller) L2WethGatewayTemplate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1AtomicTokenBridgeFactory.contract.Call(opts, &out, "l2WethGatewayTemplate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2WethGatewayTemplate is a free data retrieval call binding the contract method 0x9095765e.
//
// Solidity: function l2WethGatewayTemplate() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactorySession) L2WethGatewayTemplate() (common.Address, error) {
	return _L1AtomicTokenBridgeFactory.Contract.L2WethGatewayTemplate(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// L2WethGatewayTemplate is a free data retrieval call binding the contract method 0x9095765e.
//
// Solidity: function l2WethGatewayTemplate() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCallerSession) L2WethGatewayTemplate() (common.Address, error) {
	return _L1AtomicTokenBridgeFactory.Contract.L2WethGatewayTemplate(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// L2WethTemplate is a free data retrieval call binding the contract method 0xfd40ad85.
//
// Solidity: function l2WethTemplate() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCaller) L2WethTemplate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1AtomicTokenBridgeFactory.contract.Call(opts, &out, "l2WethTemplate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2WethTemplate is a free data retrieval call binding the contract method 0xfd40ad85.
//
// Solidity: function l2WethTemplate() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactorySession) L2WethTemplate() (common.Address, error) {
	return _L1AtomicTokenBridgeFactory.Contract.L2WethTemplate(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// L2WethTemplate is a free data retrieval call binding the contract method 0xfd40ad85.
//
// Solidity: function l2WethTemplate() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCallerSession) L2WethTemplate() (common.Address, error) {
	return _L1AtomicTokenBridgeFactory.Contract.L2WethTemplate(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1AtomicTokenBridgeFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactorySession) Owner() (common.Address, error) {
	return _L1AtomicTokenBridgeFactory.Contract.Owner(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCallerSession) Owner() (common.Address, error) {
	return _L1AtomicTokenBridgeFactory.Contract.Owner(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// RetryableSender is a free data retrieval call binding the contract method 0x36dddb97.
//
// Solidity: function retryableSender() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCaller) RetryableSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1AtomicTokenBridgeFactory.contract.Call(opts, &out, "retryableSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RetryableSender is a free data retrieval call binding the contract method 0x36dddb97.
//
// Solidity: function retryableSender() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactorySession) RetryableSender() (common.Address, error) {
	return _L1AtomicTokenBridgeFactory.Contract.RetryableSender(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// RetryableSender is a free data retrieval call binding the contract method 0x36dddb97.
//
// Solidity: function retryableSender() view returns(address)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryCallerSession) RetryableSender() (common.Address, error) {
	return _L1AtomicTokenBridgeFactory.Contract.RetryableSender(&_L1AtomicTokenBridgeFactory.CallOpts)
}

// CreateTokenBridge is a paid mutator transaction binding the contract method 0x8277742b.
//
// Solidity: function createTokenBridge(address inbox, address rollupOwner, uint256 maxGasForContracts, uint256 gasPriceBid) payable returns()
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryTransactor) CreateTokenBridge(opts *bind.TransactOpts, inbox common.Address, rollupOwner common.Address, maxGasForContracts *big.Int, gasPriceBid *big.Int) (*types.Transaction, error) {
	return _L1AtomicTokenBridgeFactory.contract.Transact(opts, "createTokenBridge", inbox, rollupOwner, maxGasForContracts, gasPriceBid)
}

// CreateTokenBridge is a paid mutator transaction binding the contract method 0x8277742b.
//
// Solidity: function createTokenBridge(address inbox, address rollupOwner, uint256 maxGasForContracts, uint256 gasPriceBid) payable returns()
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactorySession) CreateTokenBridge(inbox common.Address, rollupOwner common.Address, maxGasForContracts *big.Int, gasPriceBid *big.Int) (*types.Transaction, error) {
	return _L1AtomicTokenBridgeFactory.Contract.CreateTokenBridge(&_L1AtomicTokenBridgeFactory.TransactOpts, inbox, rollupOwner, maxGasForContracts, gasPriceBid)
}

// CreateTokenBridge is a paid mutator transaction binding the contract method 0x8277742b.
//
// Solidity: function createTokenBridge(address inbox, address rollupOwner, uint256 maxGasForContracts, uint256 gasPriceBid) payable returns()
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryTransactorSession) CreateTokenBridge(inbox common.Address, rollupOwner common.Address, maxGasForContracts *big.Int, gasPriceBid *big.Int) (*types.Transaction, error) {
	return _L1AtomicTokenBridgeFactory.Contract.CreateTokenBridge(&_L1AtomicTokenBridgeFactory.TransactOpts, inbox, rollupOwner, maxGasForContracts, gasPriceBid)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _retryableSender) returns()
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryTransactor) Initialize(opts *bind.TransactOpts, _retryableSender common.Address) (*types.Transaction, error) {
	return _L1AtomicTokenBridgeFactory.contract.Transact(opts, "initialize", _retryableSender)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _retryableSender) returns()
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactorySession) Initialize(_retryableSender common.Address) (*types.Transaction, error) {
	return _L1AtomicTokenBridgeFactory.Contract.Initialize(&_L1AtomicTokenBridgeFactory.TransactOpts, _retryableSender)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _retryableSender) returns()
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryTransactorSession) Initialize(_retryableSender common.Address) (*types.Transaction, error) {
	return _L1AtomicTokenBridgeFactory.Contract.Initialize(&_L1AtomicTokenBridgeFactory.TransactOpts, _retryableSender)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1AtomicTokenBridgeFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _L1AtomicTokenBridgeFactory.Contract.RenounceOwnership(&_L1AtomicTokenBridgeFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1AtomicTokenBridgeFactory.Contract.RenounceOwnership(&_L1AtomicTokenBridgeFactory.TransactOpts)
}

// SetDeployment is a paid mutator transaction binding the contract method 0x4c149671.
//
// Solidity: function setDeployment(address inbox, (address,address,address,address,address) l1Deployment, (address,address,address,address,address,address,address,address,address) l2Deployment) returns()
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryTransactor) SetDeployment(opts *bind.TransactOpts, inbox common.Address, l1Deployment L1DeploymentAddresses, l2Deployment L2DeploymentAddresses) (*types.Transaction, error) {
	return _L1AtomicTokenBridgeFactory.contract.Transact(opts, "setDeployment", inbox, l1Deployment, l2Deployment)
}

// SetDeployment is a paid mutator transaction binding the contract method 0x4c149671.
//
// Solidity: function setDeployment(address inbox, (address,address,address,address,address) l1Deployment, (address,address,address,address,address,address,address,address,address) l2Deployment) returns()
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactorySession) SetDeployment(inbox common.Address, l1Deployment L1DeploymentAddresses, l2Deployment L2DeploymentAddresses) (*types.Transaction, error) {
	return _L1AtomicTokenBridgeFactory.Contract.SetDeployment(&_L1AtomicTokenBridgeFactory.TransactOpts, inbox, l1Deployment, l2Deployment)
}

// SetDeployment is a paid mutator transaction binding the contract method 0x4c149671.
//
// Solidity: function setDeployment(address inbox, (address,address,address,address,address) l1Deployment, (address,address,address,address,address,address,address,address,address) l2Deployment) returns()
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryTransactorSession) SetDeployment(inbox common.Address, l1Deployment L1DeploymentAddresses, l2Deployment L2DeploymentAddresses) (*types.Transaction, error) {
	return _L1AtomicTokenBridgeFactory.Contract.SetDeployment(&_L1AtomicTokenBridgeFactory.TransactOpts, inbox, l1Deployment, l2Deployment)
}

// SetTemplates is a paid mutator transaction binding the contract method 0x81fb9184.
//
// Solidity: function setTemplates((address,address,address,address,address,address,address,address) _l1Templates, address _l2TokenBridgeFactoryTemplate, address _l2RouterTemplate, address _l2StandardGatewayTemplate, address _l2CustomGatewayTemplate, address _l2WethGatewayTemplate, address _l2WethTemplate, address _l2MulticallTemplate, address _l1Weth, address _l1Multicall, uint256 _gasLimitForL2FactoryDeployment) returns()
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryTransactor) SetTemplates(opts *bind.TransactOpts, _l1Templates L1AtomicTokenBridgeCreatorL1Templates, _l2TokenBridgeFactoryTemplate common.Address, _l2RouterTemplate common.Address, _l2StandardGatewayTemplate common.Address, _l2CustomGatewayTemplate common.Address, _l2WethGatewayTemplate common.Address, _l2WethTemplate common.Address, _l2MulticallTemplate common.Address, _l1Weth common.Address, _l1Multicall common.Address, _gasLimitForL2FactoryDeployment *big.Int) (*types.Transaction, error) {
	return _L1AtomicTokenBridgeFactory.contract.Transact(opts, "setTemplates", _l1Templates, _l2TokenBridgeFactoryTemplate, _l2RouterTemplate, _l2StandardGatewayTemplate, _l2CustomGatewayTemplate, _l2WethGatewayTemplate, _l2WethTemplate, _l2MulticallTemplate, _l1Weth, _l1Multicall, _gasLimitForL2FactoryDeployment)
}

// SetTemplates is a paid mutator transaction binding the contract method 0x81fb9184.
//
// Solidity: function setTemplates((address,address,address,address,address,address,address,address) _l1Templates, address _l2TokenBridgeFactoryTemplate, address _l2RouterTemplate, address _l2StandardGatewayTemplate, address _l2CustomGatewayTemplate, address _l2WethGatewayTemplate, address _l2WethTemplate, address _l2MulticallTemplate, address _l1Weth, address _l1Multicall, uint256 _gasLimitForL2FactoryDeployment) returns()
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactorySession) SetTemplates(_l1Templates L1AtomicTokenBridgeCreatorL1Templates, _l2TokenBridgeFactoryTemplate common.Address, _l2RouterTemplate common.Address, _l2StandardGatewayTemplate common.Address, _l2CustomGatewayTemplate common.Address, _l2WethGatewayTemplate common.Address, _l2WethTemplate common.Address, _l2MulticallTemplate common.Address, _l1Weth common.Address, _l1Multicall common.Address, _gasLimitForL2FactoryDeployment *big.Int) (*types.Transaction, error) {
	return _L1AtomicTokenBridgeFactory.Contract.SetTemplates(&_L1AtomicTokenBridgeFactory.TransactOpts, _l1Templates, _l2TokenBridgeFactoryTemplate, _l2RouterTemplate, _l2StandardGatewayTemplate, _l2CustomGatewayTemplate, _l2WethGatewayTemplate, _l2WethTemplate, _l2MulticallTemplate, _l1Weth, _l1Multicall, _gasLimitForL2FactoryDeployment)
}

// SetTemplates is a paid mutator transaction binding the contract method 0x81fb9184.
//
// Solidity: function setTemplates((address,address,address,address,address,address,address,address) _l1Templates, address _l2TokenBridgeFactoryTemplate, address _l2RouterTemplate, address _l2StandardGatewayTemplate, address _l2CustomGatewayTemplate, address _l2WethGatewayTemplate, address _l2WethTemplate, address _l2MulticallTemplate, address _l1Weth, address _l1Multicall, uint256 _gasLimitForL2FactoryDeployment) returns()
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryTransactorSession) SetTemplates(_l1Templates L1AtomicTokenBridgeCreatorL1Templates, _l2TokenBridgeFactoryTemplate common.Address, _l2RouterTemplate common.Address, _l2StandardGatewayTemplate common.Address, _l2CustomGatewayTemplate common.Address, _l2WethGatewayTemplate common.Address, _l2WethTemplate common.Address, _l2MulticallTemplate common.Address, _l1Weth common.Address, _l1Multicall common.Address, _gasLimitForL2FactoryDeployment *big.Int) (*types.Transaction, error) {
	return _L1AtomicTokenBridgeFactory.Contract.SetTemplates(&_L1AtomicTokenBridgeFactory.TransactOpts, _l1Templates, _l2TokenBridgeFactoryTemplate, _l2RouterTemplate, _l2StandardGatewayTemplate, _l2CustomGatewayTemplate, _l2WethGatewayTemplate, _l2WethTemplate, _l2MulticallTemplate, _l1Weth, _l1Multicall, _gasLimitForL2FactoryDeployment)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1AtomicTokenBridgeFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1AtomicTokenBridgeFactory.Contract.TransferOwnership(&_L1AtomicTokenBridgeFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1AtomicTokenBridgeFactory.Contract.TransferOwnership(&_L1AtomicTokenBridgeFactory.TransactOpts, newOwner)
}

// L1AtomicTokenBridgeFactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L1AtomicTokenBridgeFactory contract.
type L1AtomicTokenBridgeFactoryInitializedIterator struct {
	Event *L1AtomicTokenBridgeFactoryInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *L1AtomicTokenBridgeFactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1AtomicTokenBridgeFactoryInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(L1AtomicTokenBridgeFactoryInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *L1AtomicTokenBridgeFactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1AtomicTokenBridgeFactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1AtomicTokenBridgeFactoryInitialized represents a Initialized event raised by the L1AtomicTokenBridgeFactory contract.
type L1AtomicTokenBridgeFactoryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*L1AtomicTokenBridgeFactoryInitializedIterator, error) {

	logs, sub, err := _L1AtomicTokenBridgeFactory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L1AtomicTokenBridgeFactoryInitializedIterator{contract: _L1AtomicTokenBridgeFactory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L1AtomicTokenBridgeFactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _L1AtomicTokenBridgeFactory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1AtomicTokenBridgeFactoryInitialized)
				if err := _L1AtomicTokenBridgeFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryFilterer) ParseInitialized(log types.Log) (*L1AtomicTokenBridgeFactoryInitialized, error) {
	event := new(L1AtomicTokenBridgeFactoryInitialized)
	if err := _L1AtomicTokenBridgeFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1AtomicTokenBridgeFactoryOrbitTokenBridgeCreatedIterator is returned from FilterOrbitTokenBridgeCreated and is used to iterate over the raw logs and unpacked data for OrbitTokenBridgeCreated events raised by the L1AtomicTokenBridgeFactory contract.
type L1AtomicTokenBridgeFactoryOrbitTokenBridgeCreatedIterator struct {
	Event *L1AtomicTokenBridgeFactoryOrbitTokenBridgeCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *L1AtomicTokenBridgeFactoryOrbitTokenBridgeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1AtomicTokenBridgeFactoryOrbitTokenBridgeCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(L1AtomicTokenBridgeFactoryOrbitTokenBridgeCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *L1AtomicTokenBridgeFactoryOrbitTokenBridgeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1AtomicTokenBridgeFactoryOrbitTokenBridgeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1AtomicTokenBridgeFactoryOrbitTokenBridgeCreated represents a OrbitTokenBridgeCreated event raised by the L1AtomicTokenBridgeFactory contract.
type L1AtomicTokenBridgeFactoryOrbitTokenBridgeCreated struct {
	Inbox           common.Address
	Owner           common.Address
	L1Deployment    L1DeploymentAddresses
	L2Deployment    L2DeploymentAddresses
	ProxyAdmin      common.Address
	UpgradeExecutor common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOrbitTokenBridgeCreated is a free log retrieval operation binding the contract event 0x9a9203aa9ddcf21d8523e422e009214f0447efca13201ecdd802d8663092de7e.
//
// Solidity: event OrbitTokenBridgeCreated(address indexed inbox, address indexed owner, (address,address,address,address,address) l1Deployment, (address,address,address,address,address,address,address,address,address) l2Deployment, address proxyAdmin, address upgradeExecutor)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryFilterer) FilterOrbitTokenBridgeCreated(opts *bind.FilterOpts, inbox []common.Address, owner []common.Address) (*L1AtomicTokenBridgeFactoryOrbitTokenBridgeCreatedIterator, error) {

	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _L1AtomicTokenBridgeFactory.contract.FilterLogs(opts, "OrbitTokenBridgeCreated", inboxRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &L1AtomicTokenBridgeFactoryOrbitTokenBridgeCreatedIterator{contract: _L1AtomicTokenBridgeFactory.contract, event: "OrbitTokenBridgeCreated", logs: logs, sub: sub}, nil
}

// WatchOrbitTokenBridgeCreated is a free log subscription operation binding the contract event 0x9a9203aa9ddcf21d8523e422e009214f0447efca13201ecdd802d8663092de7e.
//
// Solidity: event OrbitTokenBridgeCreated(address indexed inbox, address indexed owner, (address,address,address,address,address) l1Deployment, (address,address,address,address,address,address,address,address,address) l2Deployment, address proxyAdmin, address upgradeExecutor)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryFilterer) WatchOrbitTokenBridgeCreated(opts *bind.WatchOpts, sink chan<- *L1AtomicTokenBridgeFactoryOrbitTokenBridgeCreated, inbox []common.Address, owner []common.Address) (event.Subscription, error) {

	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _L1AtomicTokenBridgeFactory.contract.WatchLogs(opts, "OrbitTokenBridgeCreated", inboxRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1AtomicTokenBridgeFactoryOrbitTokenBridgeCreated)
				if err := _L1AtomicTokenBridgeFactory.contract.UnpackLog(event, "OrbitTokenBridgeCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrbitTokenBridgeCreated is a log parse operation binding the contract event 0x9a9203aa9ddcf21d8523e422e009214f0447efca13201ecdd802d8663092de7e.
//
// Solidity: event OrbitTokenBridgeCreated(address indexed inbox, address indexed owner, (address,address,address,address,address) l1Deployment, (address,address,address,address,address,address,address,address,address) l2Deployment, address proxyAdmin, address upgradeExecutor)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryFilterer) ParseOrbitTokenBridgeCreated(log types.Log) (*L1AtomicTokenBridgeFactoryOrbitTokenBridgeCreated, error) {
	event := new(L1AtomicTokenBridgeFactoryOrbitTokenBridgeCreated)
	if err := _L1AtomicTokenBridgeFactory.contract.UnpackLog(event, "OrbitTokenBridgeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1AtomicTokenBridgeFactoryOrbitTokenBridgeDeploymentSetIterator is returned from FilterOrbitTokenBridgeDeploymentSet and is used to iterate over the raw logs and unpacked data for OrbitTokenBridgeDeploymentSet events raised by the L1AtomicTokenBridgeFactory contract.
type L1AtomicTokenBridgeFactoryOrbitTokenBridgeDeploymentSetIterator struct {
	Event *L1AtomicTokenBridgeFactoryOrbitTokenBridgeDeploymentSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *L1AtomicTokenBridgeFactoryOrbitTokenBridgeDeploymentSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1AtomicTokenBridgeFactoryOrbitTokenBridgeDeploymentSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(L1AtomicTokenBridgeFactoryOrbitTokenBridgeDeploymentSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *L1AtomicTokenBridgeFactoryOrbitTokenBridgeDeploymentSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1AtomicTokenBridgeFactoryOrbitTokenBridgeDeploymentSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1AtomicTokenBridgeFactoryOrbitTokenBridgeDeploymentSet represents a OrbitTokenBridgeDeploymentSet event raised by the L1AtomicTokenBridgeFactory contract.
type L1AtomicTokenBridgeFactoryOrbitTokenBridgeDeploymentSet struct {
	Inbox common.Address
	L1    L1DeploymentAddresses
	L2    L2DeploymentAddresses
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOrbitTokenBridgeDeploymentSet is a free log retrieval operation binding the contract event 0x003661d67ef6fa28d5937e796b7701a68fbf54c16d9434eb705715ebc28f424b.
//
// Solidity: event OrbitTokenBridgeDeploymentSet(address indexed inbox, (address,address,address,address,address) l1, (address,address,address,address,address,address,address,address,address) l2)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryFilterer) FilterOrbitTokenBridgeDeploymentSet(opts *bind.FilterOpts, inbox []common.Address) (*L1AtomicTokenBridgeFactoryOrbitTokenBridgeDeploymentSetIterator, error) {

	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}

	logs, sub, err := _L1AtomicTokenBridgeFactory.contract.FilterLogs(opts, "OrbitTokenBridgeDeploymentSet", inboxRule)
	if err != nil {
		return nil, err
	}
	return &L1AtomicTokenBridgeFactoryOrbitTokenBridgeDeploymentSetIterator{contract: _L1AtomicTokenBridgeFactory.contract, event: "OrbitTokenBridgeDeploymentSet", logs: logs, sub: sub}, nil
}

// WatchOrbitTokenBridgeDeploymentSet is a free log subscription operation binding the contract event 0x003661d67ef6fa28d5937e796b7701a68fbf54c16d9434eb705715ebc28f424b.
//
// Solidity: event OrbitTokenBridgeDeploymentSet(address indexed inbox, (address,address,address,address,address) l1, (address,address,address,address,address,address,address,address,address) l2)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryFilterer) WatchOrbitTokenBridgeDeploymentSet(opts *bind.WatchOpts, sink chan<- *L1AtomicTokenBridgeFactoryOrbitTokenBridgeDeploymentSet, inbox []common.Address) (event.Subscription, error) {

	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}

	logs, sub, err := _L1AtomicTokenBridgeFactory.contract.WatchLogs(opts, "OrbitTokenBridgeDeploymentSet", inboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1AtomicTokenBridgeFactoryOrbitTokenBridgeDeploymentSet)
				if err := _L1AtomicTokenBridgeFactory.contract.UnpackLog(event, "OrbitTokenBridgeDeploymentSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrbitTokenBridgeDeploymentSet is a log parse operation binding the contract event 0x003661d67ef6fa28d5937e796b7701a68fbf54c16d9434eb705715ebc28f424b.
//
// Solidity: event OrbitTokenBridgeDeploymentSet(address indexed inbox, (address,address,address,address,address) l1, (address,address,address,address,address,address,address,address,address) l2)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryFilterer) ParseOrbitTokenBridgeDeploymentSet(log types.Log) (*L1AtomicTokenBridgeFactoryOrbitTokenBridgeDeploymentSet, error) {
	event := new(L1AtomicTokenBridgeFactoryOrbitTokenBridgeDeploymentSet)
	if err := _L1AtomicTokenBridgeFactory.contract.UnpackLog(event, "OrbitTokenBridgeDeploymentSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1AtomicTokenBridgeFactoryOrbitTokenBridgeTemplatesUpdatedIterator is returned from FilterOrbitTokenBridgeTemplatesUpdated and is used to iterate over the raw logs and unpacked data for OrbitTokenBridgeTemplatesUpdated events raised by the L1AtomicTokenBridgeFactory contract.
type L1AtomicTokenBridgeFactoryOrbitTokenBridgeTemplatesUpdatedIterator struct {
	Event *L1AtomicTokenBridgeFactoryOrbitTokenBridgeTemplatesUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *L1AtomicTokenBridgeFactoryOrbitTokenBridgeTemplatesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1AtomicTokenBridgeFactoryOrbitTokenBridgeTemplatesUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(L1AtomicTokenBridgeFactoryOrbitTokenBridgeTemplatesUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *L1AtomicTokenBridgeFactoryOrbitTokenBridgeTemplatesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1AtomicTokenBridgeFactoryOrbitTokenBridgeTemplatesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1AtomicTokenBridgeFactoryOrbitTokenBridgeTemplatesUpdated represents a OrbitTokenBridgeTemplatesUpdated event raised by the L1AtomicTokenBridgeFactory contract.
type L1AtomicTokenBridgeFactoryOrbitTokenBridgeTemplatesUpdated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOrbitTokenBridgeTemplatesUpdated is a free log retrieval operation binding the contract event 0x8a040c53d83c1e62b7c4c7c88774e55775a95bdaadc2c0ef9f63b8d1a118ffdc.
//
// Solidity: event OrbitTokenBridgeTemplatesUpdated()
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryFilterer) FilterOrbitTokenBridgeTemplatesUpdated(opts *bind.FilterOpts) (*L1AtomicTokenBridgeFactoryOrbitTokenBridgeTemplatesUpdatedIterator, error) {

	logs, sub, err := _L1AtomicTokenBridgeFactory.contract.FilterLogs(opts, "OrbitTokenBridgeTemplatesUpdated")
	if err != nil {
		return nil, err
	}
	return &L1AtomicTokenBridgeFactoryOrbitTokenBridgeTemplatesUpdatedIterator{contract: _L1AtomicTokenBridgeFactory.contract, event: "OrbitTokenBridgeTemplatesUpdated", logs: logs, sub: sub}, nil
}

// WatchOrbitTokenBridgeTemplatesUpdated is a free log subscription operation binding the contract event 0x8a040c53d83c1e62b7c4c7c88774e55775a95bdaadc2c0ef9f63b8d1a118ffdc.
//
// Solidity: event OrbitTokenBridgeTemplatesUpdated()
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryFilterer) WatchOrbitTokenBridgeTemplatesUpdated(opts *bind.WatchOpts, sink chan<- *L1AtomicTokenBridgeFactoryOrbitTokenBridgeTemplatesUpdated) (event.Subscription, error) {

	logs, sub, err := _L1AtomicTokenBridgeFactory.contract.WatchLogs(opts, "OrbitTokenBridgeTemplatesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1AtomicTokenBridgeFactoryOrbitTokenBridgeTemplatesUpdated)
				if err := _L1AtomicTokenBridgeFactory.contract.UnpackLog(event, "OrbitTokenBridgeTemplatesUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrbitTokenBridgeTemplatesUpdated is a log parse operation binding the contract event 0x8a040c53d83c1e62b7c4c7c88774e55775a95bdaadc2c0ef9f63b8d1a118ffdc.
//
// Solidity: event OrbitTokenBridgeTemplatesUpdated()
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryFilterer) ParseOrbitTokenBridgeTemplatesUpdated(log types.Log) (*L1AtomicTokenBridgeFactoryOrbitTokenBridgeTemplatesUpdated, error) {
	event := new(L1AtomicTokenBridgeFactoryOrbitTokenBridgeTemplatesUpdated)
	if err := _L1AtomicTokenBridgeFactory.contract.UnpackLog(event, "OrbitTokenBridgeTemplatesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1AtomicTokenBridgeFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L1AtomicTokenBridgeFactory contract.
type L1AtomicTokenBridgeFactoryOwnershipTransferredIterator struct {
	Event *L1AtomicTokenBridgeFactoryOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *L1AtomicTokenBridgeFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1AtomicTokenBridgeFactoryOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(L1AtomicTokenBridgeFactoryOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *L1AtomicTokenBridgeFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1AtomicTokenBridgeFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1AtomicTokenBridgeFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the L1AtomicTokenBridgeFactory contract.
type L1AtomicTokenBridgeFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L1AtomicTokenBridgeFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1AtomicTokenBridgeFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L1AtomicTokenBridgeFactoryOwnershipTransferredIterator{contract: _L1AtomicTokenBridgeFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L1AtomicTokenBridgeFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1AtomicTokenBridgeFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1AtomicTokenBridgeFactoryOwnershipTransferred)
				if err := _L1AtomicTokenBridgeFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1AtomicTokenBridgeFactory *L1AtomicTokenBridgeFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*L1AtomicTokenBridgeFactoryOwnershipTransferred, error) {
	event := new(L1AtomicTokenBridgeFactoryOwnershipTransferred)
	if err := _L1AtomicTokenBridgeFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
