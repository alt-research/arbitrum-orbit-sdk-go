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

// SequencerInboxMetaData contains all meta data concerning the SequencerInbox contract.
var SequencerInboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"TargetCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"upgrade\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"UpgradeExecuted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXECUTOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"upgrade\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"upgradeCallData\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"targetCallData\",\"type\":\"bytes\"}],\"name\":\"executeCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"executors\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SequencerInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use SequencerInboxMetaData.ABI instead.
var SequencerInboxABI = SequencerInboxMetaData.ABI

// SequencerInbox is an auto generated Go binding around an Ethereum contract.
type SequencerInbox struct {
	SequencerInboxCaller     // Read-only binding to the contract
	SequencerInboxTransactor // Write-only binding to the contract
	SequencerInboxFilterer   // Log filterer for contract events
}

// SequencerInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type SequencerInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SequencerInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SequencerInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SequencerInboxSession struct {
	Contract     *SequencerInbox   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SequencerInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SequencerInboxCallerSession struct {
	Contract *SequencerInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SequencerInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SequencerInboxTransactorSession struct {
	Contract     *SequencerInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SequencerInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type SequencerInboxRaw struct {
	Contract *SequencerInbox // Generic contract binding to access the raw methods on
}

// SequencerInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SequencerInboxCallerRaw struct {
	Contract *SequencerInboxCaller // Generic read-only contract binding to access the raw methods on
}

// SequencerInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SequencerInboxTransactorRaw struct {
	Contract *SequencerInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSequencerInbox creates a new instance of SequencerInbox, bound to a specific deployed contract.
func NewSequencerInbox(address common.Address, backend bind.ContractBackend) (*SequencerInbox, error) {
	contract, err := bindSequencerInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SequencerInbox{SequencerInboxCaller: SequencerInboxCaller{contract: contract}, SequencerInboxTransactor: SequencerInboxTransactor{contract: contract}, SequencerInboxFilterer: SequencerInboxFilterer{contract: contract}}, nil
}

// NewSequencerInboxCaller creates a new read-only instance of SequencerInbox, bound to a specific deployed contract.
func NewSequencerInboxCaller(address common.Address, caller bind.ContractCaller) (*SequencerInboxCaller, error) {
	contract, err := bindSequencerInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxCaller{contract: contract}, nil
}

// NewSequencerInboxTransactor creates a new write-only instance of SequencerInbox, bound to a specific deployed contract.
func NewSequencerInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*SequencerInboxTransactor, error) {
	contract, err := bindSequencerInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxTransactor{contract: contract}, nil
}

// NewSequencerInboxFilterer creates a new log filterer instance of SequencerInbox, bound to a specific deployed contract.
func NewSequencerInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*SequencerInboxFilterer, error) {
	contract, err := bindSequencerInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxFilterer{contract: contract}, nil
}

// bindSequencerInbox binds a generic wrapper to an already deployed contract.
func bindSequencerInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SequencerInboxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequencerInbox *SequencerInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequencerInbox.Contract.SequencerInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequencerInbox *SequencerInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerInbox.Contract.SequencerInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequencerInbox *SequencerInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequencerInbox.Contract.SequencerInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequencerInbox *SequencerInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequencerInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequencerInbox *SequencerInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequencerInbox *SequencerInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequencerInbox.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_SequencerInbox *SequencerInboxCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_SequencerInbox *SequencerInboxSession) ADMINROLE() ([32]byte, error) {
	return _SequencerInbox.Contract.ADMINROLE(&_SequencerInbox.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_SequencerInbox *SequencerInboxCallerSession) ADMINROLE() ([32]byte, error) {
	return _SequencerInbox.Contract.ADMINROLE(&_SequencerInbox.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SequencerInbox *SequencerInboxCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SequencerInbox *SequencerInboxSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SequencerInbox.Contract.DEFAULTADMINROLE(&_SequencerInbox.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SequencerInbox *SequencerInboxCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SequencerInbox.Contract.DEFAULTADMINROLE(&_SequencerInbox.CallOpts)
}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_SequencerInbox *SequencerInboxCaller) EXECUTORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "EXECUTOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_SequencerInbox *SequencerInboxSession) EXECUTORROLE() ([32]byte, error) {
	return _SequencerInbox.Contract.EXECUTORROLE(&_SequencerInbox.CallOpts)
}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_SequencerInbox *SequencerInboxCallerSession) EXECUTORROLE() ([32]byte, error) {
	return _SequencerInbox.Contract.EXECUTORROLE(&_SequencerInbox.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SequencerInbox *SequencerInboxCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SequencerInbox *SequencerInboxSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SequencerInbox.Contract.GetRoleAdmin(&_SequencerInbox.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SequencerInbox *SequencerInboxCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SequencerInbox.Contract.GetRoleAdmin(&_SequencerInbox.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SequencerInbox *SequencerInboxCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SequencerInbox *SequencerInboxSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SequencerInbox.Contract.HasRole(&_SequencerInbox.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SequencerInbox *SequencerInboxCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SequencerInbox.Contract.HasRole(&_SequencerInbox.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SequencerInbox *SequencerInboxCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SequencerInbox.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SequencerInbox *SequencerInboxSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SequencerInbox.Contract.SupportsInterface(&_SequencerInbox.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SequencerInbox *SequencerInboxCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SequencerInbox.Contract.SupportsInterface(&_SequencerInbox.CallOpts, interfaceId)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address upgrade, bytes upgradeCallData) payable returns()
func (_SequencerInbox *SequencerInboxTransactor) Execute(opts *bind.TransactOpts, upgrade common.Address, upgradeCallData []byte) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "execute", upgrade, upgradeCallData)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address upgrade, bytes upgradeCallData) payable returns()
func (_SequencerInbox *SequencerInboxSession) Execute(upgrade common.Address, upgradeCallData []byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.Execute(&_SequencerInbox.TransactOpts, upgrade, upgradeCallData)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address upgrade, bytes upgradeCallData) payable returns()
func (_SequencerInbox *SequencerInboxTransactorSession) Execute(upgrade common.Address, upgradeCallData []byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.Execute(&_SequencerInbox.TransactOpts, upgrade, upgradeCallData)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0xbca8c7b5.
//
// Solidity: function executeCall(address target, bytes targetCallData) payable returns()
func (_SequencerInbox *SequencerInboxTransactor) ExecuteCall(opts *bind.TransactOpts, target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "executeCall", target, targetCallData)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0xbca8c7b5.
//
// Solidity: function executeCall(address target, bytes targetCallData) payable returns()
func (_SequencerInbox *SequencerInboxSession) ExecuteCall(target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.ExecuteCall(&_SequencerInbox.TransactOpts, target, targetCallData)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0xbca8c7b5.
//
// Solidity: function executeCall(address target, bytes targetCallData) payable returns()
func (_SequencerInbox *SequencerInboxTransactorSession) ExecuteCall(target common.Address, targetCallData []byte) (*types.Transaction, error) {
	return _SequencerInbox.Contract.ExecuteCall(&_SequencerInbox.TransactOpts, target, targetCallData)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SequencerInbox *SequencerInboxTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SequencerInbox *SequencerInboxSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SequencerInbox.Contract.GrantRole(&_SequencerInbox.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SequencerInbox.Contract.GrantRole(&_SequencerInbox.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address admin, address[] executors) returns()
func (_SequencerInbox *SequencerInboxTransactor) Initialize(opts *bind.TransactOpts, admin common.Address, executors []common.Address) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "initialize", admin, executors)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address admin, address[] executors) returns()
func (_SequencerInbox *SequencerInboxSession) Initialize(admin common.Address, executors []common.Address) (*types.Transaction, error) {
	return _SequencerInbox.Contract.Initialize(&_SequencerInbox.TransactOpts, admin, executors)
}

// Initialize is a paid mutator transaction binding the contract method 0x946d9204.
//
// Solidity: function initialize(address admin, address[] executors) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) Initialize(admin common.Address, executors []common.Address) (*types.Transaction, error) {
	return _SequencerInbox.Contract.Initialize(&_SequencerInbox.TransactOpts, admin, executors)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SequencerInbox *SequencerInboxTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SequencerInbox *SequencerInboxSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SequencerInbox.Contract.RenounceRole(&_SequencerInbox.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SequencerInbox.Contract.RenounceRole(&_SequencerInbox.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SequencerInbox *SequencerInboxTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SequencerInbox.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SequencerInbox *SequencerInboxSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SequencerInbox.Contract.RevokeRole(&_SequencerInbox.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SequencerInbox *SequencerInboxTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SequencerInbox.Contract.RevokeRole(&_SequencerInbox.TransactOpts, role, account)
}

// SequencerInboxInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SequencerInbox contract.
type SequencerInboxInitializedIterator struct {
	Event *SequencerInboxInitialized // Event containing the contract specifics and raw log

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
func (it *SequencerInboxInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxInitialized)
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
		it.Event = new(SequencerInboxInitialized)
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
func (it *SequencerInboxInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxInitialized represents a Initialized event raised by the SequencerInbox contract.
type SequencerInboxInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SequencerInbox *SequencerInboxFilterer) FilterInitialized(opts *bind.FilterOpts) (*SequencerInboxInitializedIterator, error) {

	logs, sub, err := _SequencerInbox.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SequencerInboxInitializedIterator{contract: _SequencerInbox.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SequencerInbox *SequencerInboxFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SequencerInboxInitialized) (event.Subscription, error) {

	logs, sub, err := _SequencerInbox.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxInitialized)
				if err := _SequencerInbox.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SequencerInbox *SequencerInboxFilterer) ParseInitialized(log types.Log) (*SequencerInboxInitialized, error) {
	event := new(SequencerInboxInitialized)
	if err := _SequencerInbox.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the SequencerInbox contract.
type SequencerInboxRoleAdminChangedIterator struct {
	Event *SequencerInboxRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *SequencerInboxRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxRoleAdminChanged)
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
		it.Event = new(SequencerInboxRoleAdminChanged)
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
func (it *SequencerInboxRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxRoleAdminChanged represents a RoleAdminChanged event raised by the SequencerInbox contract.
type SequencerInboxRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SequencerInbox *SequencerInboxFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SequencerInboxRoleAdminChangedIterator, error) {

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

	logs, sub, err := _SequencerInbox.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxRoleAdminChangedIterator{contract: _SequencerInbox.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SequencerInbox *SequencerInboxFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SequencerInboxRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _SequencerInbox.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxRoleAdminChanged)
				if err := _SequencerInbox.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_SequencerInbox *SequencerInboxFilterer) ParseRoleAdminChanged(log types.Log) (*SequencerInboxRoleAdminChanged, error) {
	event := new(SequencerInboxRoleAdminChanged)
	if err := _SequencerInbox.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the SequencerInbox contract.
type SequencerInboxRoleGrantedIterator struct {
	Event *SequencerInboxRoleGranted // Event containing the contract specifics and raw log

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
func (it *SequencerInboxRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxRoleGranted)
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
		it.Event = new(SequencerInboxRoleGranted)
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
func (it *SequencerInboxRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxRoleGranted represents a RoleGranted event raised by the SequencerInbox contract.
type SequencerInboxRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SequencerInbox *SequencerInboxFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SequencerInboxRoleGrantedIterator, error) {

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

	logs, sub, err := _SequencerInbox.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxRoleGrantedIterator{contract: _SequencerInbox.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SequencerInbox *SequencerInboxFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SequencerInboxRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SequencerInbox.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxRoleGranted)
				if err := _SequencerInbox.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_SequencerInbox *SequencerInboxFilterer) ParseRoleGranted(log types.Log) (*SequencerInboxRoleGranted, error) {
	event := new(SequencerInboxRoleGranted)
	if err := _SequencerInbox.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the SequencerInbox contract.
type SequencerInboxRoleRevokedIterator struct {
	Event *SequencerInboxRoleRevoked // Event containing the contract specifics and raw log

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
func (it *SequencerInboxRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxRoleRevoked)
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
		it.Event = new(SequencerInboxRoleRevoked)
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
func (it *SequencerInboxRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxRoleRevoked represents a RoleRevoked event raised by the SequencerInbox contract.
type SequencerInboxRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SequencerInbox *SequencerInboxFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SequencerInboxRoleRevokedIterator, error) {

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

	logs, sub, err := _SequencerInbox.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxRoleRevokedIterator{contract: _SequencerInbox.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SequencerInbox *SequencerInboxFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SequencerInboxRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SequencerInbox.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxRoleRevoked)
				if err := _SequencerInbox.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_SequencerInbox *SequencerInboxFilterer) ParseRoleRevoked(log types.Log) (*SequencerInboxRoleRevoked, error) {
	event := new(SequencerInboxRoleRevoked)
	if err := _SequencerInbox.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxTargetCallExecutedIterator is returned from FilterTargetCallExecuted and is used to iterate over the raw logs and unpacked data for TargetCallExecuted events raised by the SequencerInbox contract.
type SequencerInboxTargetCallExecutedIterator struct {
	Event *SequencerInboxTargetCallExecuted // Event containing the contract specifics and raw log

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
func (it *SequencerInboxTargetCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxTargetCallExecuted)
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
		it.Event = new(SequencerInboxTargetCallExecuted)
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
func (it *SequencerInboxTargetCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxTargetCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxTargetCallExecuted represents a TargetCallExecuted event raised by the SequencerInbox contract.
type SequencerInboxTargetCallExecuted struct {
	Target common.Address
	Value  *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTargetCallExecuted is a free log retrieval operation binding the contract event 0x4d7dbdcc249630ec373f584267f10abf44938de920c32562f5aee93959c25258.
//
// Solidity: event TargetCallExecuted(address indexed target, uint256 value, bytes data)
func (_SequencerInbox *SequencerInboxFilterer) FilterTargetCallExecuted(opts *bind.FilterOpts, target []common.Address) (*SequencerInboxTargetCallExecutedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _SequencerInbox.contract.FilterLogs(opts, "TargetCallExecuted", targetRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxTargetCallExecutedIterator{contract: _SequencerInbox.contract, event: "TargetCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTargetCallExecuted is a free log subscription operation binding the contract event 0x4d7dbdcc249630ec373f584267f10abf44938de920c32562f5aee93959c25258.
//
// Solidity: event TargetCallExecuted(address indexed target, uint256 value, bytes data)
func (_SequencerInbox *SequencerInboxFilterer) WatchTargetCallExecuted(opts *bind.WatchOpts, sink chan<- *SequencerInboxTargetCallExecuted, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _SequencerInbox.contract.WatchLogs(opts, "TargetCallExecuted", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxTargetCallExecuted)
				if err := _SequencerInbox.contract.UnpackLog(event, "TargetCallExecuted", log); err != nil {
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
func (_SequencerInbox *SequencerInboxFilterer) ParseTargetCallExecuted(log types.Log) (*SequencerInboxTargetCallExecuted, error) {
	event := new(SequencerInboxTargetCallExecuted)
	if err := _SequencerInbox.contract.UnpackLog(event, "TargetCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxUpgradeExecutedIterator is returned from FilterUpgradeExecuted and is used to iterate over the raw logs and unpacked data for UpgradeExecuted events raised by the SequencerInbox contract.
type SequencerInboxUpgradeExecutedIterator struct {
	Event *SequencerInboxUpgradeExecuted // Event containing the contract specifics and raw log

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
func (it *SequencerInboxUpgradeExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxUpgradeExecuted)
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
		it.Event = new(SequencerInboxUpgradeExecuted)
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
func (it *SequencerInboxUpgradeExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxUpgradeExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxUpgradeExecuted represents a UpgradeExecuted event raised by the SequencerInbox contract.
type SequencerInboxUpgradeExecuted struct {
	Upgrade common.Address
	Value   *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpgradeExecuted is a free log retrieval operation binding the contract event 0x49f6851d1cd01a518db5bdea5cffbbe90276baa2595f74250b7472b96806302e.
//
// Solidity: event UpgradeExecuted(address indexed upgrade, uint256 value, bytes data)
func (_SequencerInbox *SequencerInboxFilterer) FilterUpgradeExecuted(opts *bind.FilterOpts, upgrade []common.Address) (*SequencerInboxUpgradeExecutedIterator, error) {

	var upgradeRule []interface{}
	for _, upgradeItem := range upgrade {
		upgradeRule = append(upgradeRule, upgradeItem)
	}

	logs, sub, err := _SequencerInbox.contract.FilterLogs(opts, "UpgradeExecuted", upgradeRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxUpgradeExecutedIterator{contract: _SequencerInbox.contract, event: "UpgradeExecuted", logs: logs, sub: sub}, nil
}

// WatchUpgradeExecuted is a free log subscription operation binding the contract event 0x49f6851d1cd01a518db5bdea5cffbbe90276baa2595f74250b7472b96806302e.
//
// Solidity: event UpgradeExecuted(address indexed upgrade, uint256 value, bytes data)
func (_SequencerInbox *SequencerInboxFilterer) WatchUpgradeExecuted(opts *bind.WatchOpts, sink chan<- *SequencerInboxUpgradeExecuted, upgrade []common.Address) (event.Subscription, error) {

	var upgradeRule []interface{}
	for _, upgradeItem := range upgrade {
		upgradeRule = append(upgradeRule, upgradeItem)
	}

	logs, sub, err := _SequencerInbox.contract.WatchLogs(opts, "UpgradeExecuted", upgradeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxUpgradeExecuted)
				if err := _SequencerInbox.contract.UnpackLog(event, "UpgradeExecuted", log); err != nil {
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
func (_SequencerInbox *SequencerInboxFilterer) ParseUpgradeExecuted(log types.Log) (*SequencerInboxUpgradeExecuted, error) {
	event := new(SequencerInboxUpgradeExecuted)
	if err := _SequencerInbox.contract.UnpackLog(event, "UpgradeExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
