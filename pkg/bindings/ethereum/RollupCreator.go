// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereum

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

// Config is an auto generated low-level Go binding around an user-defined struct.
type Config struct {
	ConfirmPeriodBlocks            uint64
	ExtraChallengeTimeBlocks       uint64
	StakeToken                     common.Address
	BaseStake                      *big.Int
	WasmModuleRoot                 [32]byte
	Owner                          common.Address
	LoserStakeEscrow               common.Address
	ChainId                        *big.Int
	ChainConfig                    string
	GenesisBlockNum                uint64
	SequencerInboxMaxTimeVariation ISequencerInboxMaxTimeVariation
}

// ISequencerInboxMaxTimeVariation is an auto generated low-level Go binding around an user-defined struct.
type ISequencerInboxMaxTimeVariation struct {
	DelayBlocks   *big.Int
	FutureBlocks  *big.Int
	DelaySeconds  *big.Int
	FutureSeconds *big.Int
}

// RollupCreatorRollupDeploymentParams is an auto generated low-level Go binding around an user-defined struct.
type RollupCreatorRollupDeploymentParams struct {
	Config                    Config
	BatchPoster               common.Address
	Validators                []common.Address
	MaxDataSize               *big.Int
	NativeToken               common.Address
	DeployFactoriesToL2       bool
	MaxFeePerGasForRetryables *big.Int
}

// RollupCreatorMetaData contains all meta data concerning the RollupCreator contract.
var RollupCreatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nativeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inboxAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rollupEventInbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challengeManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"adminProxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencerInbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"upgradeExecutor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validatorUtils\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validatorWalletCreator\",\"type\":\"address\"}],\"name\":\"RollupCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"TemplatesUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgeCreator\",\"outputs\":[{\"internalType\":\"contractBridgeCreator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeManagerTemplate\",\"outputs\":[{\"internalType\":\"contractIChallengeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"confirmPeriodBlocks\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"extraChallengeTimeBlocks\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"stakeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"baseStake\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"loserStakeEscrow\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"chainConfig\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"genesisBlockNum\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"delayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delaySeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureSeconds\",\"type\":\"uint256\"}],\"internalType\":\"structISequencerInbox.MaxTimeVariation\",\"name\":\"sequencerInboxMaxTimeVariation\",\"type\":\"tuple\"}],\"internalType\":\"structConfig\",\"name\":\"config\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"batchPoster\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"maxDataSize\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nativeToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"deployFactoriesToL2\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGasForRetryables\",\"type\":\"uint256\"}],\"internalType\":\"structRollupCreator.RollupDeploymentParams\",\"name\":\"deployParams\",\"type\":\"tuple\"}],\"name\":\"createRollup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2FactoriesDeployer\",\"outputs\":[{\"internalType\":\"contractDeployHelper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"osp\",\"outputs\":[{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupAdminLogic\",\"outputs\":[{\"internalType\":\"contractIRollupAdmin\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupUserLogic\",\"outputs\":[{\"internalType\":\"contractIRollupUser\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBridgeCreator\",\"name\":\"_bridgeCreator\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"_osp\",\"type\":\"address\"},{\"internalType\":\"contractIChallengeManager\",\"name\":\"_challengeManagerLogic\",\"type\":\"address\"},{\"internalType\":\"contractIRollupAdmin\",\"name\":\"_rollupAdminLogic\",\"type\":\"address\"},{\"internalType\":\"contractIRollupUser\",\"name\":\"_rollupUserLogic\",\"type\":\"address\"},{\"internalType\":\"contractIUpgradeExecutor\",\"name\":\"_upgradeExecutorLogic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorUtils\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorWalletCreator\",\"type\":\"address\"},{\"internalType\":\"contractDeployHelper\",\"name\":\"_l2FactoriesDeployer\",\"type\":\"address\"}],\"name\":\"setTemplates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeExecutorLogic\",\"outputs\":[{\"internalType\":\"contractIUpgradeExecutor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorUtils\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorWalletCreator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RollupCreatorABI is the input ABI used to generate the binding from.
// Deprecated: Use RollupCreatorMetaData.ABI instead.
var RollupCreatorABI = RollupCreatorMetaData.ABI

// RollupCreator is an auto generated Go binding around an Ethereum contract.
type RollupCreator struct {
	RollupCreatorCaller     // Read-only binding to the contract
	RollupCreatorTransactor // Write-only binding to the contract
	RollupCreatorFilterer   // Log filterer for contract events
}

// RollupCreatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupCreatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCreatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupCreatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCreatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupCreatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCreatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupCreatorSession struct {
	Contract     *RollupCreator    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupCreatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupCreatorCallerSession struct {
	Contract *RollupCreatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RollupCreatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupCreatorTransactorSession struct {
	Contract     *RollupCreatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RollupCreatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupCreatorRaw struct {
	Contract *RollupCreator // Generic contract binding to access the raw methods on
}

// RollupCreatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupCreatorCallerRaw struct {
	Contract *RollupCreatorCaller // Generic read-only contract binding to access the raw methods on
}

// RollupCreatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupCreatorTransactorRaw struct {
	Contract *RollupCreatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupCreator creates a new instance of RollupCreator, bound to a specific deployed contract.
func NewRollupCreator(address common.Address, backend bind.ContractBackend) (*RollupCreator, error) {
	contract, err := bindRollupCreator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupCreator{RollupCreatorCaller: RollupCreatorCaller{contract: contract}, RollupCreatorTransactor: RollupCreatorTransactor{contract: contract}, RollupCreatorFilterer: RollupCreatorFilterer{contract: contract}}, nil
}

// NewRollupCreatorCaller creates a new read-only instance of RollupCreator, bound to a specific deployed contract.
func NewRollupCreatorCaller(address common.Address, caller bind.ContractCaller) (*RollupCreatorCaller, error) {
	contract, err := bindRollupCreator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorCaller{contract: contract}, nil
}

// NewRollupCreatorTransactor creates a new write-only instance of RollupCreator, bound to a specific deployed contract.
func NewRollupCreatorTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupCreatorTransactor, error) {
	contract, err := bindRollupCreator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorTransactor{contract: contract}, nil
}

// NewRollupCreatorFilterer creates a new log filterer instance of RollupCreator, bound to a specific deployed contract.
func NewRollupCreatorFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupCreatorFilterer, error) {
	contract, err := bindRollupCreator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorFilterer{contract: contract}, nil
}

// bindRollupCreator binds a generic wrapper to an already deployed contract.
func bindRollupCreator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RollupCreatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupCreator *RollupCreatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupCreator.Contract.RollupCreatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupCreator *RollupCreatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupCreator.Contract.RollupCreatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupCreator *RollupCreatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupCreator.Contract.RollupCreatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupCreator *RollupCreatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupCreator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupCreator *RollupCreatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupCreator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupCreator *RollupCreatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupCreator.Contract.contract.Transact(opts, method, params...)
}

// BridgeCreator is a free data retrieval call binding the contract method 0xf860cefa.
//
// Solidity: function bridgeCreator() view returns(address)
func (_RollupCreator *RollupCreatorCaller) BridgeCreator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupCreator.contract.Call(opts, &out, "bridgeCreator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeCreator is a free data retrieval call binding the contract method 0xf860cefa.
//
// Solidity: function bridgeCreator() view returns(address)
func (_RollupCreator *RollupCreatorSession) BridgeCreator() (common.Address, error) {
	return _RollupCreator.Contract.BridgeCreator(&_RollupCreator.CallOpts)
}

// BridgeCreator is a free data retrieval call binding the contract method 0xf860cefa.
//
// Solidity: function bridgeCreator() view returns(address)
func (_RollupCreator *RollupCreatorCallerSession) BridgeCreator() (common.Address, error) {
	return _RollupCreator.Contract.BridgeCreator(&_RollupCreator.CallOpts)
}

// ChallengeManagerTemplate is a free data retrieval call binding the contract method 0x9c683d10.
//
// Solidity: function challengeManagerTemplate() view returns(address)
func (_RollupCreator *RollupCreatorCaller) ChallengeManagerTemplate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupCreator.contract.Call(opts, &out, "challengeManagerTemplate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChallengeManagerTemplate is a free data retrieval call binding the contract method 0x9c683d10.
//
// Solidity: function challengeManagerTemplate() view returns(address)
func (_RollupCreator *RollupCreatorSession) ChallengeManagerTemplate() (common.Address, error) {
	return _RollupCreator.Contract.ChallengeManagerTemplate(&_RollupCreator.CallOpts)
}

// ChallengeManagerTemplate is a free data retrieval call binding the contract method 0x9c683d10.
//
// Solidity: function challengeManagerTemplate() view returns(address)
func (_RollupCreator *RollupCreatorCallerSession) ChallengeManagerTemplate() (common.Address, error) {
	return _RollupCreator.Contract.ChallengeManagerTemplate(&_RollupCreator.CallOpts)
}

// L2FactoriesDeployer is a free data retrieval call binding the contract method 0xac0425bc.
//
// Solidity: function l2FactoriesDeployer() view returns(address)
func (_RollupCreator *RollupCreatorCaller) L2FactoriesDeployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupCreator.contract.Call(opts, &out, "l2FactoriesDeployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2FactoriesDeployer is a free data retrieval call binding the contract method 0xac0425bc.
//
// Solidity: function l2FactoriesDeployer() view returns(address)
func (_RollupCreator *RollupCreatorSession) L2FactoriesDeployer() (common.Address, error) {
	return _RollupCreator.Contract.L2FactoriesDeployer(&_RollupCreator.CallOpts)
}

// L2FactoriesDeployer is a free data retrieval call binding the contract method 0xac0425bc.
//
// Solidity: function l2FactoriesDeployer() view returns(address)
func (_RollupCreator *RollupCreatorCallerSession) L2FactoriesDeployer() (common.Address, error) {
	return _RollupCreator.Contract.L2FactoriesDeployer(&_RollupCreator.CallOpts)
}

// Osp is a free data retrieval call binding the contract method 0xf26a62c6.
//
// Solidity: function osp() view returns(address)
func (_RollupCreator *RollupCreatorCaller) Osp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupCreator.contract.Call(opts, &out, "osp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Osp is a free data retrieval call binding the contract method 0xf26a62c6.
//
// Solidity: function osp() view returns(address)
func (_RollupCreator *RollupCreatorSession) Osp() (common.Address, error) {
	return _RollupCreator.Contract.Osp(&_RollupCreator.CallOpts)
}

// Osp is a free data retrieval call binding the contract method 0xf26a62c6.
//
// Solidity: function osp() view returns(address)
func (_RollupCreator *RollupCreatorCallerSession) Osp() (common.Address, error) {
	return _RollupCreator.Contract.Osp(&_RollupCreator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupCreator *RollupCreatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupCreator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupCreator *RollupCreatorSession) Owner() (common.Address, error) {
	return _RollupCreator.Contract.Owner(&_RollupCreator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupCreator *RollupCreatorCallerSession) Owner() (common.Address, error) {
	return _RollupCreator.Contract.Owner(&_RollupCreator.CallOpts)
}

// RollupAdminLogic is a free data retrieval call binding the contract method 0x9dba3241.
//
// Solidity: function rollupAdminLogic() view returns(address)
func (_RollupCreator *RollupCreatorCaller) RollupAdminLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupCreator.contract.Call(opts, &out, "rollupAdminLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupAdminLogic is a free data retrieval call binding the contract method 0x9dba3241.
//
// Solidity: function rollupAdminLogic() view returns(address)
func (_RollupCreator *RollupCreatorSession) RollupAdminLogic() (common.Address, error) {
	return _RollupCreator.Contract.RollupAdminLogic(&_RollupCreator.CallOpts)
}

// RollupAdminLogic is a free data retrieval call binding the contract method 0x9dba3241.
//
// Solidity: function rollupAdminLogic() view returns(address)
func (_RollupCreator *RollupCreatorCallerSession) RollupAdminLogic() (common.Address, error) {
	return _RollupCreator.Contract.RollupAdminLogic(&_RollupCreator.CallOpts)
}

// RollupUserLogic is a free data retrieval call binding the contract method 0x9d4798e3.
//
// Solidity: function rollupUserLogic() view returns(address)
func (_RollupCreator *RollupCreatorCaller) RollupUserLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupCreator.contract.Call(opts, &out, "rollupUserLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupUserLogic is a free data retrieval call binding the contract method 0x9d4798e3.
//
// Solidity: function rollupUserLogic() view returns(address)
func (_RollupCreator *RollupCreatorSession) RollupUserLogic() (common.Address, error) {
	return _RollupCreator.Contract.RollupUserLogic(&_RollupCreator.CallOpts)
}

// RollupUserLogic is a free data retrieval call binding the contract method 0x9d4798e3.
//
// Solidity: function rollupUserLogic() view returns(address)
func (_RollupCreator *RollupCreatorCallerSession) RollupUserLogic() (common.Address, error) {
	return _RollupCreator.Contract.RollupUserLogic(&_RollupCreator.CallOpts)
}

// UpgradeExecutorLogic is a free data retrieval call binding the contract method 0x030cb85e.
//
// Solidity: function upgradeExecutorLogic() view returns(address)
func (_RollupCreator *RollupCreatorCaller) UpgradeExecutorLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupCreator.contract.Call(opts, &out, "upgradeExecutorLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpgradeExecutorLogic is a free data retrieval call binding the contract method 0x030cb85e.
//
// Solidity: function upgradeExecutorLogic() view returns(address)
func (_RollupCreator *RollupCreatorSession) UpgradeExecutorLogic() (common.Address, error) {
	return _RollupCreator.Contract.UpgradeExecutorLogic(&_RollupCreator.CallOpts)
}

// UpgradeExecutorLogic is a free data retrieval call binding the contract method 0x030cb85e.
//
// Solidity: function upgradeExecutorLogic() view returns(address)
func (_RollupCreator *RollupCreatorCallerSession) UpgradeExecutorLogic() (common.Address, error) {
	return _RollupCreator.Contract.UpgradeExecutorLogic(&_RollupCreator.CallOpts)
}

// ValidatorUtils is a free data retrieval call binding the contract method 0x014cc92c.
//
// Solidity: function validatorUtils() view returns(address)
func (_RollupCreator *RollupCreatorCaller) ValidatorUtils(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupCreator.contract.Call(opts, &out, "validatorUtils")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorUtils is a free data retrieval call binding the contract method 0x014cc92c.
//
// Solidity: function validatorUtils() view returns(address)
func (_RollupCreator *RollupCreatorSession) ValidatorUtils() (common.Address, error) {
	return _RollupCreator.Contract.ValidatorUtils(&_RollupCreator.CallOpts)
}

// ValidatorUtils is a free data retrieval call binding the contract method 0x014cc92c.
//
// Solidity: function validatorUtils() view returns(address)
func (_RollupCreator *RollupCreatorCallerSession) ValidatorUtils() (common.Address, error) {
	return _RollupCreator.Contract.ValidatorUtils(&_RollupCreator.CallOpts)
}

// ValidatorWalletCreator is a free data retrieval call binding the contract method 0xbc45e0ae.
//
// Solidity: function validatorWalletCreator() view returns(address)
func (_RollupCreator *RollupCreatorCaller) ValidatorWalletCreator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupCreator.contract.Call(opts, &out, "validatorWalletCreator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorWalletCreator is a free data retrieval call binding the contract method 0xbc45e0ae.
//
// Solidity: function validatorWalletCreator() view returns(address)
func (_RollupCreator *RollupCreatorSession) ValidatorWalletCreator() (common.Address, error) {
	return _RollupCreator.Contract.ValidatorWalletCreator(&_RollupCreator.CallOpts)
}

// ValidatorWalletCreator is a free data retrieval call binding the contract method 0xbc45e0ae.
//
// Solidity: function validatorWalletCreator() view returns(address)
func (_RollupCreator *RollupCreatorCallerSession) ValidatorWalletCreator() (common.Address, error) {
	return _RollupCreator.Contract.ValidatorWalletCreator(&_RollupCreator.CallOpts)
}

// CreateRollup is a paid mutator transaction binding the contract method 0xcb73d6e2.
//
// Solidity: function createRollup(((uint64,uint64,address,uint256,bytes32,address,address,uint256,string,uint64,(uint256,uint256,uint256,uint256)),address,address[],uint256,address,bool,uint256) deployParams) payable returns(address)
func (_RollupCreator *RollupCreatorTransactor) CreateRollup(opts *bind.TransactOpts, deployParams RollupCreatorRollupDeploymentParams) (*types.Transaction, error) {
	return _RollupCreator.contract.Transact(opts, "createRollup", deployParams)
}

// CreateRollup is a paid mutator transaction binding the contract method 0xcb73d6e2.
//
// Solidity: function createRollup(((uint64,uint64,address,uint256,bytes32,address,address,uint256,string,uint64,(uint256,uint256,uint256,uint256)),address,address[],uint256,address,bool,uint256) deployParams) payable returns(address)
func (_RollupCreator *RollupCreatorSession) CreateRollup(deployParams RollupCreatorRollupDeploymentParams) (*types.Transaction, error) {
	return _RollupCreator.Contract.CreateRollup(&_RollupCreator.TransactOpts, deployParams)
}

// CreateRollup is a paid mutator transaction binding the contract method 0xcb73d6e2.
//
// Solidity: function createRollup(((uint64,uint64,address,uint256,bytes32,address,address,uint256,string,uint64,(uint256,uint256,uint256,uint256)),address,address[],uint256,address,bool,uint256) deployParams) payable returns(address)
func (_RollupCreator *RollupCreatorTransactorSession) CreateRollup(deployParams RollupCreatorRollupDeploymentParams) (*types.Transaction, error) {
	return _RollupCreator.Contract.CreateRollup(&_RollupCreator.TransactOpts, deployParams)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RollupCreator *RollupCreatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupCreator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RollupCreator *RollupCreatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RollupCreator.Contract.RenounceOwnership(&_RollupCreator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RollupCreator *RollupCreatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RollupCreator.Contract.RenounceOwnership(&_RollupCreator.TransactOpts)
}

// SetTemplates is a paid mutator transaction binding the contract method 0xac9a97b4.
//
// Solidity: function setTemplates(address _bridgeCreator, address _osp, address _challengeManagerLogic, address _rollupAdminLogic, address _rollupUserLogic, address _upgradeExecutorLogic, address _validatorUtils, address _validatorWalletCreator, address _l2FactoriesDeployer) returns()
func (_RollupCreator *RollupCreatorTransactor) SetTemplates(opts *bind.TransactOpts, _bridgeCreator common.Address, _osp common.Address, _challengeManagerLogic common.Address, _rollupAdminLogic common.Address, _rollupUserLogic common.Address, _upgradeExecutorLogic common.Address, _validatorUtils common.Address, _validatorWalletCreator common.Address, _l2FactoriesDeployer common.Address) (*types.Transaction, error) {
	return _RollupCreator.contract.Transact(opts, "setTemplates", _bridgeCreator, _osp, _challengeManagerLogic, _rollupAdminLogic, _rollupUserLogic, _upgradeExecutorLogic, _validatorUtils, _validatorWalletCreator, _l2FactoriesDeployer)
}

// SetTemplates is a paid mutator transaction binding the contract method 0xac9a97b4.
//
// Solidity: function setTemplates(address _bridgeCreator, address _osp, address _challengeManagerLogic, address _rollupAdminLogic, address _rollupUserLogic, address _upgradeExecutorLogic, address _validatorUtils, address _validatorWalletCreator, address _l2FactoriesDeployer) returns()
func (_RollupCreator *RollupCreatorSession) SetTemplates(_bridgeCreator common.Address, _osp common.Address, _challengeManagerLogic common.Address, _rollupAdminLogic common.Address, _rollupUserLogic common.Address, _upgradeExecutorLogic common.Address, _validatorUtils common.Address, _validatorWalletCreator common.Address, _l2FactoriesDeployer common.Address) (*types.Transaction, error) {
	return _RollupCreator.Contract.SetTemplates(&_RollupCreator.TransactOpts, _bridgeCreator, _osp, _challengeManagerLogic, _rollupAdminLogic, _rollupUserLogic, _upgradeExecutorLogic, _validatorUtils, _validatorWalletCreator, _l2FactoriesDeployer)
}

// SetTemplates is a paid mutator transaction binding the contract method 0xac9a97b4.
//
// Solidity: function setTemplates(address _bridgeCreator, address _osp, address _challengeManagerLogic, address _rollupAdminLogic, address _rollupUserLogic, address _upgradeExecutorLogic, address _validatorUtils, address _validatorWalletCreator, address _l2FactoriesDeployer) returns()
func (_RollupCreator *RollupCreatorTransactorSession) SetTemplates(_bridgeCreator common.Address, _osp common.Address, _challengeManagerLogic common.Address, _rollupAdminLogic common.Address, _rollupUserLogic common.Address, _upgradeExecutorLogic common.Address, _validatorUtils common.Address, _validatorWalletCreator common.Address, _l2FactoriesDeployer common.Address) (*types.Transaction, error) {
	return _RollupCreator.Contract.SetTemplates(&_RollupCreator.TransactOpts, _bridgeCreator, _osp, _challengeManagerLogic, _rollupAdminLogic, _rollupUserLogic, _upgradeExecutorLogic, _validatorUtils, _validatorWalletCreator, _l2FactoriesDeployer)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RollupCreator *RollupCreatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RollupCreator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RollupCreator *RollupCreatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RollupCreator.Contract.TransferOwnership(&_RollupCreator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RollupCreator *RollupCreatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RollupCreator.Contract.TransferOwnership(&_RollupCreator.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RollupCreator *RollupCreatorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupCreator.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RollupCreator *RollupCreatorSession) Receive() (*types.Transaction, error) {
	return _RollupCreator.Contract.Receive(&_RollupCreator.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RollupCreator *RollupCreatorTransactorSession) Receive() (*types.Transaction, error) {
	return _RollupCreator.Contract.Receive(&_RollupCreator.TransactOpts)
}

// RollupCreatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RollupCreator contract.
type RollupCreatorOwnershipTransferredIterator struct {
	Event *RollupCreatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RollupCreatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupCreatorOwnershipTransferred)
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
		it.Event = new(RollupCreatorOwnershipTransferred)
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
func (it *RollupCreatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupCreatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupCreatorOwnershipTransferred represents a OwnershipTransferred event raised by the RollupCreator contract.
type RollupCreatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RollupCreator *RollupCreatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RollupCreatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RollupCreator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorOwnershipTransferredIterator{contract: _RollupCreator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RollupCreator *RollupCreatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RollupCreatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RollupCreator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupCreatorOwnershipTransferred)
				if err := _RollupCreator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RollupCreator *RollupCreatorFilterer) ParseOwnershipTransferred(log types.Log) (*RollupCreatorOwnershipTransferred, error) {
	event := new(RollupCreatorOwnershipTransferred)
	if err := _RollupCreator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupCreatorRollupCreatedIterator is returned from FilterRollupCreated and is used to iterate over the raw logs and unpacked data for RollupCreated events raised by the RollupCreator contract.
type RollupCreatorRollupCreatedIterator struct {
	Event *RollupCreatorRollupCreated // Event containing the contract specifics and raw log

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
func (it *RollupCreatorRollupCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupCreatorRollupCreated)
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
		it.Event = new(RollupCreatorRollupCreated)
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
func (it *RollupCreatorRollupCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupCreatorRollupCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupCreatorRollupCreated represents a RollupCreated event raised by the RollupCreator contract.
type RollupCreatorRollupCreated struct {
	RollupAddress          common.Address
	NativeToken            common.Address
	InboxAddress           common.Address
	Outbox                 common.Address
	RollupEventInbox       common.Address
	ChallengeManager       common.Address
	AdminProxy             common.Address
	SequencerInbox         common.Address
	Bridge                 common.Address
	UpgradeExecutor        common.Address
	ValidatorUtils         common.Address
	ValidatorWalletCreator common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRollupCreated is a free log retrieval operation binding the contract event 0x481277de518d1f364b196166b90219b996fba76138a3dc84e7fe02540eb1cbdb.
//
// Solidity: event RollupCreated(address indexed rollupAddress, address indexed nativeToken, address inboxAddress, address outbox, address rollupEventInbox, address challengeManager, address adminProxy, address sequencerInbox, address bridge, address upgradeExecutor, address validatorUtils, address validatorWalletCreator)
func (_RollupCreator *RollupCreatorFilterer) FilterRollupCreated(opts *bind.FilterOpts, rollupAddress []common.Address, nativeToken []common.Address) (*RollupCreatorRollupCreatedIterator, error) {

	var rollupAddressRule []interface{}
	for _, rollupAddressItem := range rollupAddress {
		rollupAddressRule = append(rollupAddressRule, rollupAddressItem)
	}
	var nativeTokenRule []interface{}
	for _, nativeTokenItem := range nativeToken {
		nativeTokenRule = append(nativeTokenRule, nativeTokenItem)
	}

	logs, sub, err := _RollupCreator.contract.FilterLogs(opts, "RollupCreated", rollupAddressRule, nativeTokenRule)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorRollupCreatedIterator{contract: _RollupCreator.contract, event: "RollupCreated", logs: logs, sub: sub}, nil
}

// WatchRollupCreated is a free log subscription operation binding the contract event 0x481277de518d1f364b196166b90219b996fba76138a3dc84e7fe02540eb1cbdb.
//
// Solidity: event RollupCreated(address indexed rollupAddress, address indexed nativeToken, address inboxAddress, address outbox, address rollupEventInbox, address challengeManager, address adminProxy, address sequencerInbox, address bridge, address upgradeExecutor, address validatorUtils, address validatorWalletCreator)
func (_RollupCreator *RollupCreatorFilterer) WatchRollupCreated(opts *bind.WatchOpts, sink chan<- *RollupCreatorRollupCreated, rollupAddress []common.Address, nativeToken []common.Address) (event.Subscription, error) {

	var rollupAddressRule []interface{}
	for _, rollupAddressItem := range rollupAddress {
		rollupAddressRule = append(rollupAddressRule, rollupAddressItem)
	}
	var nativeTokenRule []interface{}
	for _, nativeTokenItem := range nativeToken {
		nativeTokenRule = append(nativeTokenRule, nativeTokenItem)
	}

	logs, sub, err := _RollupCreator.contract.WatchLogs(opts, "RollupCreated", rollupAddressRule, nativeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupCreatorRollupCreated)
				if err := _RollupCreator.contract.UnpackLog(event, "RollupCreated", log); err != nil {
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

// ParseRollupCreated is a log parse operation binding the contract event 0x481277de518d1f364b196166b90219b996fba76138a3dc84e7fe02540eb1cbdb.
//
// Solidity: event RollupCreated(address indexed rollupAddress, address indexed nativeToken, address inboxAddress, address outbox, address rollupEventInbox, address challengeManager, address adminProxy, address sequencerInbox, address bridge, address upgradeExecutor, address validatorUtils, address validatorWalletCreator)
func (_RollupCreator *RollupCreatorFilterer) ParseRollupCreated(log types.Log) (*RollupCreatorRollupCreated, error) {
	event := new(RollupCreatorRollupCreated)
	if err := _RollupCreator.contract.UnpackLog(event, "RollupCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupCreatorTemplatesUpdatedIterator is returned from FilterTemplatesUpdated and is used to iterate over the raw logs and unpacked data for TemplatesUpdated events raised by the RollupCreator contract.
type RollupCreatorTemplatesUpdatedIterator struct {
	Event *RollupCreatorTemplatesUpdated // Event containing the contract specifics and raw log

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
func (it *RollupCreatorTemplatesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupCreatorTemplatesUpdated)
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
		it.Event = new(RollupCreatorTemplatesUpdated)
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
func (it *RollupCreatorTemplatesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupCreatorTemplatesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupCreatorTemplatesUpdated represents a TemplatesUpdated event raised by the RollupCreator contract.
type RollupCreatorTemplatesUpdated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTemplatesUpdated is a free log retrieval operation binding the contract event 0xc9d3947d22fa124aaec4c7e8c919f79016e2d7b48eee10568375d98b86460d1b.
//
// Solidity: event TemplatesUpdated()
func (_RollupCreator *RollupCreatorFilterer) FilterTemplatesUpdated(opts *bind.FilterOpts) (*RollupCreatorTemplatesUpdatedIterator, error) {

	logs, sub, err := _RollupCreator.contract.FilterLogs(opts, "TemplatesUpdated")
	if err != nil {
		return nil, err
	}
	return &RollupCreatorTemplatesUpdatedIterator{contract: _RollupCreator.contract, event: "TemplatesUpdated", logs: logs, sub: sub}, nil
}

// WatchTemplatesUpdated is a free log subscription operation binding the contract event 0xc9d3947d22fa124aaec4c7e8c919f79016e2d7b48eee10568375d98b86460d1b.
//
// Solidity: event TemplatesUpdated()
func (_RollupCreator *RollupCreatorFilterer) WatchTemplatesUpdated(opts *bind.WatchOpts, sink chan<- *RollupCreatorTemplatesUpdated) (event.Subscription, error) {

	logs, sub, err := _RollupCreator.contract.WatchLogs(opts, "TemplatesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupCreatorTemplatesUpdated)
				if err := _RollupCreator.contract.UnpackLog(event, "TemplatesUpdated", log); err != nil {
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

// ParseTemplatesUpdated is a log parse operation binding the contract event 0xc9d3947d22fa124aaec4c7e8c919f79016e2d7b48eee10568375d98b86460d1b.
//
// Solidity: event TemplatesUpdated()
func (_RollupCreator *RollupCreatorFilterer) ParseTemplatesUpdated(log types.Log) (*RollupCreatorTemplatesUpdated, error) {
	event := new(RollupCreatorTemplatesUpdated)
	if err := _RollupCreator.contract.UnpackLog(event, "TemplatesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
