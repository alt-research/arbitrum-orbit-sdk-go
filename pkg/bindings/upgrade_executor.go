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

// UpgradeExecutorMetaData contains all meta data concerning the UpgradeExecutor contract.
var UpgradeExecutorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"TargetCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"upgrade\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"UpgradeExecuted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXECUTOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"upgrade\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"upgradeCallData\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"targetCallData\",\"type\":\"bytes\"}],\"name\":\"executeCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"executors\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UpgradeExecutorABI is the input ABI used to generate the binding from.
// Deprecated: Use UpgradeExecutorMetaData.ABI instead.
var UpgradeExecutorABI = UpgradeExecutorMetaData.ABI

// UpgradeExecutor is an auto generated Go binding around an Ethereum contract.
type UpgradeExecutor struct {
	UpgradeExecutorCaller     // Read-only binding to the contract
	UpgradeExecutorTransactor // Write-only binding to the contract
	UpgradeExecutorFilterer   // Log filterer for contract events
}

// UpgradeExecutorCaller is an auto generated read-only Go binding around an Ethereum contract.
type UpgradeExecutorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeExecutorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UpgradeExecutorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeExecutorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UpgradeExecutorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeExecutorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UpgradeExecutorSession struct {
	Contract     *UpgradeExecutor  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UpgradeExecutorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UpgradeExecutorCallerSession struct {
	Contract *UpgradeExecutorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// UpgradeExecutorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UpgradeExecutorTransactorSession struct {
	Contract     *UpgradeExecutorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// UpgradeExecutorRaw is an auto generated low-level Go binding around an Ethereum contract.
type UpgradeExecutorRaw struct {
	Contract *UpgradeExecutor // Generic contract binding to access the raw methods on
}

// UpgradeExecutorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UpgradeExecutorCallerRaw struct {
	Contract *UpgradeExecutorCaller // Generic read-only contract binding to access the raw methods on
}

// UpgradeExecutorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UpgradeExecutorTransactorRaw struct {
	Contract *UpgradeExecutorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpgradeExecutor creates a new instance of UpgradeExecutor, bound to a specific deployed contract.
func NewUpgradeExecutor(address common.Address, backend bind.ContractBackend) (*UpgradeExecutor, error) {
	contract, err := bindUpgradeExecutor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutor{UpgradeExecutorCaller: UpgradeExecutorCaller{contract: contract}, UpgradeExecutorTransactor: UpgradeExecutorTransactor{contract: contract}, UpgradeExecutorFilterer: UpgradeExecutorFilterer{contract: contract}}, nil
}

// NewUpgradeExecutorCaller creates a new read-only instance of UpgradeExecutor, bound to a specific deployed contract.
func NewUpgradeExecutorCaller(address common.Address, caller bind.ContractCaller) (*UpgradeExecutorCaller, error) {
	contract, err := bindUpgradeExecutor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorCaller{contract: contract}, nil
}

// NewUpgradeExecutorTransactor creates a new write-only instance of UpgradeExecutor, bound to a specific deployed contract.
func NewUpgradeExecutorTransactor(address common.Address, transactor bind.ContractTransactor) (*UpgradeExecutorTransactor, error) {
	contract, err := bindUpgradeExecutor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorTransactor{contract: contract}, nil
}

// NewUpgradeExecutorFilterer creates a new log filterer instance of UpgradeExecutor, bound to a specific deployed contract.
func NewUpgradeExecutorFilterer(address common.Address, filterer bind.ContractFilterer) (*UpgradeExecutorFilterer, error) {
	contract, err := bindUpgradeExecutor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorFilterer{contract: contract}, nil
}

// bindUpgradeExecutor binds a generic wrapper to an already deployed contract.
func bindUpgradeExecutor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UpgradeExecutorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeExecutor *UpgradeExecutorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeExecutor.Contract.UpgradeExecutorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeExecutor *UpgradeExecutorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.UpgradeExecutorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeExecutor *UpgradeExecutorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.UpgradeExecutorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeExecutor *UpgradeExecutorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeExecutor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeExecutor *UpgradeExecutorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeExecutor *UpgradeExecutorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_UpgradeExecutor *UpgradeExecutorCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UpgradeExecutor.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_UpgradeExecutor *UpgradeExecutorSession) ADMINROLE() ([32]byte, error) {
	return _UpgradeExecutor.Contract.ADMINROLE(&_UpgradeExecutor.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_UpgradeExecutor *UpgradeExecutorCallerSession) ADMINROLE() ([32]byte, error) {
	return _UpgradeExecutor.Contract.ADMINROLE(&_UpgradeExecutor.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_UpgradeExecutor *UpgradeExecutorCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UpgradeExecutor.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_UpgradeExecutor *UpgradeExecutorSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _UpgradeExecutor.Contract.DEFAULTADMINROLE(&_UpgradeExecutor.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_UpgradeExecutor *UpgradeExecutorCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _UpgradeExecutor.Contract.DEFAULTADMINROLE(&_UpgradeExecutor.CallOpts)
}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_UpgradeExecutor *UpgradeExecutorCaller) EXECUTORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UpgradeExecutor.contract.Call(opts, &out, "EXECUTOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_UpgradeExecutor *UpgradeExecutorSession) EXECUTORROLE() ([32]byte, error) {
	return _UpgradeExecutor.Contract.EXECUTORROLE(&_UpgradeExecutor.CallOpts)
}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_UpgradeExecutor *UpgradeExecutorCallerSession) EXECUTORROLE() ([32]byte, error) {
	return _UpgradeExecutor.Contract.EXECUTORROLE(&_UpgradeExecutor.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_UpgradeExecutor *UpgradeExecutorCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _UpgradeExecutor.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_UpgradeExecutor *UpgradeExecutorSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _UpgradeExecutor.Contract.GetRoleAdmin(&_UpgradeExecutor.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_UpgradeExecutor *UpgradeExecutorCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _UpgradeExecutor.Contract.GetRoleAdmin(&_UpgradeExecutor.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_UpgradeExecutor *UpgradeExecutorCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _UpgradeExecutor.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_UpgradeExecutor *UpgradeExecutorSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _UpgradeExecutor.Contract.HasRole(&_UpgradeExecutor.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_UpgradeExecutor *UpgradeExecutorCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _UpgradeExecutor.Contract.HasRole(&_UpgradeExecutor.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UpgradeExecutor *UpgradeExecutorCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _UpgradeExecutor.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UpgradeExecutor *UpgradeExecutorSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UpgradeExecutor.Contract.SupportsInterface(&_UpgradeExecutor.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_UpgradeExecutor *UpgradeExecutorCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _UpgradeExecutor.Contract.SupportsInterface(&_UpgradeExecutor.CallOpts, interfaceId)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address upgrade, bytes upgradeCallData) payable returns()
func (_UpgradeExecutor *UpgradeExecutorTransactor) Execute(opts *bind.TransactOpts, upgrade common.Address, upgradeCallData []byte) (*types.Transaction, error) {
	return _UpgradeExecutor.contract.Transact(opts, "execute", upgrade, upgradeCallData)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address upgrade, bytes upgradeCallData) payable returns()
func (_UpgradeExecutor *UpgradeExecutorSession) Execute(upgrade common.Address, upgradeCallData []byte) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.Execute(&_UpgradeExecutor.TransactOpts, upgrade, upgradeCallData)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address upgrade, bytes upgradeCallData) payable returns()
func (_UpgradeExecutor *UpgradeExecutorTransactorSession) Execute(upgrade common.Address, upgradeCallData []byte) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.Execute(&_UpgradeExecutor.TransactOpts, upgrade, upgradeCallData)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0xbca8c7b5.
//
// Solidity: function executeCall(address target, bytes targetCallData) payable returns()
func (_UpgradeExecutor *UpgradeExecutorTransactor) ExecuteCall(opts *bind.TransactOpts, target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _UpgradeExecutor.contract.Transact(opts, "executeCall", target, targetCallData)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0xbca8c7b5.
//
// Solidity: function executeCall(address target, bytes targetCallData) payable returns()
func (_UpgradeExecutor *UpgradeExecutorSession) ExecuteCall(target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.ExecuteCall(&_UpgradeExecutor.TransactOpts, target, targetCallData)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0xbca8c7b5.
//
// Solidity: function executeCall(address target, bytes targetCallData) payable returns()
func (_UpgradeExecutor *UpgradeExecutorTransactorSession) ExecuteCall(target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.ExecuteCall(&_UpgradeExecutor.TransactOpts, target, targetCallData)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_UpgradeExecutor *UpgradeExecutorTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutor.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_UpgradeExecutor *UpgradeExecutorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.GrantRole(&_UpgradeExecutor.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_UpgradeExecutor *UpgradeExecutorTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.GrantRole(&_UpgradeExecutor.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address admin, address[] executors) returns()
func (_UpgradeExecutor *UpgradeExecutorTransactor) Initialize(opts *bind.TransactOpts, admin common.Address, executors []common.Address) (*types.Transaction, error) {
	return _UpgradeExecutor.contract.Transact(opts, "initialize", admin, executors)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address admin, address[] executors) returns()
func (_UpgradeExecutor *UpgradeExecutorSession) Initialize(admin common.Address, executors []common.Address) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.Initialize(&_UpgradeExecutor.TransactOpts, admin, executors)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address admin, address[] executors) returns()
func (_UpgradeExecutor *UpgradeExecutorTransactorSession) Initialize(admin common.Address, executors []common.Address) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.Initialize(&_UpgradeExecutor.TransactOpts, admin, executors)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_UpgradeExecutor *UpgradeExecutorTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutor.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_UpgradeExecutor *UpgradeExecutorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.RenounceRole(&_UpgradeExecutor.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_UpgradeExecutor *UpgradeExecutorTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.RenounceRole(&_UpgradeExecutor.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_UpgradeExecutor *UpgradeExecutorTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutor.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_UpgradeExecutor *UpgradeExecutorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.RevokeRole(&_UpgradeExecutor.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_UpgradeExecutor *UpgradeExecutorTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _UpgradeExecutor.Contract.RevokeRole(&_UpgradeExecutor.TransactOpts, role, account)
}

// UpgradeExecutorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the UpgradeExecutor contract.
type UpgradeExecutorInitializedIterator struct {
	Event *UpgradeExecutorInitialized // Event containing the contract specifics and raw log

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
func (it *UpgradeExecutorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeExecutorInitialized)
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
		it.Event = new(UpgradeExecutorInitialized)
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
func (it *UpgradeExecutorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeExecutorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeExecutorInitialized represents a Initialized event raised by the UpgradeExecutor contract.
type UpgradeExecutorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UpgradeExecutor *UpgradeExecutorFilterer) FilterInitialized(opts *bind.FilterOpts) (*UpgradeExecutorInitializedIterator, error) {

	logs, sub, err := _UpgradeExecutor.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorInitializedIterator{contract: _UpgradeExecutor.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UpgradeExecutor *UpgradeExecutorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *UpgradeExecutorInitialized) (event.Subscription, error) {

	logs, sub, err := _UpgradeExecutor.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeExecutorInitialized)
				if err := _UpgradeExecutor.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_UpgradeExecutor *UpgradeExecutorFilterer) ParseInitialized(log types.Log) (*UpgradeExecutorInitialized, error) {
	event := new(UpgradeExecutorInitialized)
	if err := _UpgradeExecutor.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeExecutorRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the UpgradeExecutor contract.
type UpgradeExecutorRoleAdminChangedIterator struct {
	Event *UpgradeExecutorRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *UpgradeExecutorRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeExecutorRoleAdminChanged)
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
		it.Event = new(UpgradeExecutorRoleAdminChanged)
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
func (it *UpgradeExecutorRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeExecutorRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeExecutorRoleAdminChanged represents a RoleAdminChanged event raised by the UpgradeExecutor contract.
type UpgradeExecutorRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_UpgradeExecutor *UpgradeExecutorFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*UpgradeExecutorRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _UpgradeExecutor.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorRoleAdminChangedIterator{contract: _UpgradeExecutor.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_UpgradeExecutor *UpgradeExecutorFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *UpgradeExecutorRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _UpgradeExecutor.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeExecutorRoleAdminChanged)
				if err := _UpgradeExecutor.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_UpgradeExecutor *UpgradeExecutorFilterer) ParseRoleAdminChanged(log types.Log) (*UpgradeExecutorRoleAdminChanged, error) {
	event := new(UpgradeExecutorRoleAdminChanged)
	if err := _UpgradeExecutor.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeExecutorRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the UpgradeExecutor contract.
type UpgradeExecutorRoleGrantedIterator struct {
	Event *UpgradeExecutorRoleGranted // Event containing the contract specifics and raw log

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
func (it *UpgradeExecutorRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeExecutorRoleGranted)
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
		it.Event = new(UpgradeExecutorRoleGranted)
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
func (it *UpgradeExecutorRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeExecutorRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeExecutorRoleGranted represents a RoleGranted event raised by the UpgradeExecutor contract.
type UpgradeExecutorRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_UpgradeExecutor *UpgradeExecutorFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*UpgradeExecutorRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _UpgradeExecutor.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorRoleGrantedIterator{contract: _UpgradeExecutor.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_UpgradeExecutor *UpgradeExecutorFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *UpgradeExecutorRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _UpgradeExecutor.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeExecutorRoleGranted)
				if err := _UpgradeExecutor.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_UpgradeExecutor *UpgradeExecutorFilterer) ParseRoleGranted(log types.Log) (*UpgradeExecutorRoleGranted, error) {
	event := new(UpgradeExecutorRoleGranted)
	if err := _UpgradeExecutor.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeExecutorRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the UpgradeExecutor contract.
type UpgradeExecutorRoleRevokedIterator struct {
	Event *UpgradeExecutorRoleRevoked // Event containing the contract specifics and raw log

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
func (it *UpgradeExecutorRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeExecutorRoleRevoked)
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
		it.Event = new(UpgradeExecutorRoleRevoked)
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
func (it *UpgradeExecutorRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeExecutorRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeExecutorRoleRevoked represents a RoleRevoked event raised by the UpgradeExecutor contract.
type UpgradeExecutorRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_UpgradeExecutor *UpgradeExecutorFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*UpgradeExecutorRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _UpgradeExecutor.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorRoleRevokedIterator{contract: _UpgradeExecutor.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_UpgradeExecutor *UpgradeExecutorFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *UpgradeExecutorRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _UpgradeExecutor.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeExecutorRoleRevoked)
				if err := _UpgradeExecutor.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_UpgradeExecutor *UpgradeExecutorFilterer) ParseRoleRevoked(log types.Log) (*UpgradeExecutorRoleRevoked, error) {
	event := new(UpgradeExecutorRoleRevoked)
	if err := _UpgradeExecutor.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeExecutorTargetCallExecutedIterator is returned from FilterTargetCallExecuted and is used to iterate over the raw logs and unpacked data for TargetCallExecuted events raised by the UpgradeExecutor contract.
type UpgradeExecutorTargetCallExecutedIterator struct {
	Event *UpgradeExecutorTargetCallExecuted // Event containing the contract specifics and raw log

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
func (it *UpgradeExecutorTargetCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeExecutorTargetCallExecuted)
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
		it.Event = new(UpgradeExecutorTargetCallExecuted)
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
func (it *UpgradeExecutorTargetCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeExecutorTargetCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeExecutorTargetCallExecuted represents a TargetCallExecuted event raised by the UpgradeExecutor contract.
type UpgradeExecutorTargetCallExecuted struct {
	Target common.Address
	Value  *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTargetCallExecuted is a free log retrieval operation binding the contract event 0x4d7dbdcc249630ec373f584267f10abf44938de920c32562f5aee93959c25258.
//
// Solidity: event TargetCallExecuted(address indexed target, uint256 value, bytes data)
func (_UpgradeExecutor *UpgradeExecutorFilterer) FilterTargetCallExecuted(opts *bind.FilterOpts, target []common.Address) (*UpgradeExecutorTargetCallExecutedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _UpgradeExecutor.contract.FilterLogs(opts, "TargetCallExecuted", targetRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorTargetCallExecutedIterator{contract: _UpgradeExecutor.contract, event: "TargetCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTargetCallExecuted is a free log subscription operation binding the contract event 0x4d7dbdcc249630ec373f584267f10abf44938de920c32562f5aee93959c25258.
//
// Solidity: event TargetCallExecuted(address indexed target, uint256 value, bytes data)
func (_UpgradeExecutor *UpgradeExecutorFilterer) WatchTargetCallExecuted(opts *bind.WatchOpts, sink chan<- *UpgradeExecutorTargetCallExecuted, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _UpgradeExecutor.contract.WatchLogs(opts, "TargetCallExecuted", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeExecutorTargetCallExecuted)
				if err := _UpgradeExecutor.contract.UnpackLog(event, "TargetCallExecuted", log); err != nil {
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

// ParseTargetCallExecuted is a log parse operation binding the contract event 0x4d7dbdcc249630ec373f584267f10abf44938de920c32562f5aee93959c25258.
//
// Solidity: event TargetCallExecuted(address indexed target, uint256 value, bytes data)
func (_UpgradeExecutor *UpgradeExecutorFilterer) ParseTargetCallExecuted(log types.Log) (*UpgradeExecutorTargetCallExecuted, error) {
	event := new(UpgradeExecutorTargetCallExecuted)
	if err := _UpgradeExecutor.contract.UnpackLog(event, "TargetCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeExecutorUpgradeExecutedIterator is returned from FilterUpgradeExecuted and is used to iterate over the raw logs and unpacked data for UpgradeExecuted events raised by the UpgradeExecutor contract.
type UpgradeExecutorUpgradeExecutedIterator struct {
	Event *UpgradeExecutorUpgradeExecuted // Event containing the contract specifics and raw log

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
func (it *UpgradeExecutorUpgradeExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeExecutorUpgradeExecuted)
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
		it.Event = new(UpgradeExecutorUpgradeExecuted)
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
func (it *UpgradeExecutorUpgradeExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeExecutorUpgradeExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeExecutorUpgradeExecuted represents a UpgradeExecuted event raised by the UpgradeExecutor contract.
type UpgradeExecutorUpgradeExecuted struct {
	Upgrade common.Address
	Value   *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpgradeExecuted is a free log retrieval operation binding the contract event 0x49f6851d1cd01a518db5bdea5cffbbe90276baa2595f74250b7472b96806302e.
//
// Solidity: event UpgradeExecuted(address indexed upgrade, uint256 value, bytes data)
func (_UpgradeExecutor *UpgradeExecutorFilterer) FilterUpgradeExecuted(opts *bind.FilterOpts, upgrade []common.Address) (*UpgradeExecutorUpgradeExecutedIterator, error) {

	var upgradeRule []interface{}
	for _, upgradeItem := range upgrade {
		upgradeRule = append(upgradeRule, upgradeItem)
	}

	logs, sub, err := _UpgradeExecutor.contract.FilterLogs(opts, "UpgradeExecuted", upgradeRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeExecutorUpgradeExecutedIterator{contract: _UpgradeExecutor.contract, event: "UpgradeExecuted", logs: logs, sub: sub}, nil
}

// WatchUpgradeExecuted is a free log subscription operation binding the contract event 0x49f6851d1cd01a518db5bdea5cffbbe90276baa2595f74250b7472b96806302e.
//
// Solidity: event UpgradeExecuted(address indexed upgrade, uint256 value, bytes data)
func (_UpgradeExecutor *UpgradeExecutorFilterer) WatchUpgradeExecuted(opts *bind.WatchOpts, sink chan<- *UpgradeExecutorUpgradeExecuted, upgrade []common.Address) (event.Subscription, error) {

	var upgradeRule []interface{}
	for _, upgradeItem := range upgrade {
		upgradeRule = append(upgradeRule, upgradeItem)
	}

	logs, sub, err := _UpgradeExecutor.contract.WatchLogs(opts, "UpgradeExecuted", upgradeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeExecutorUpgradeExecuted)
				if err := _UpgradeExecutor.contract.UnpackLog(event, "UpgradeExecuted", log); err != nil {
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

// ParseUpgradeExecuted is a log parse operation binding the contract event 0x49f6851d1cd01a518db5bdea5cffbbe90276baa2595f74250b7472b96806302e.
//
// Solidity: event UpgradeExecuted(address indexed upgrade, uint256 value, bytes data)
func (_UpgradeExecutor *UpgradeExecutorFilterer) ParseUpgradeExecuted(log types.Log) (*UpgradeExecutorUpgradeExecuted, error) {
	event := new(UpgradeExecutorUpgradeExecuted)
	if err := _UpgradeExecutor.contract.UnpackLog(event, "UpgradeExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
