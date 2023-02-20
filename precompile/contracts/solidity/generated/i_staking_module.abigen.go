// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package generated

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
)

// IStakingModuleRedelegationEntry is an auto generated low-level Go binding around an user-defined struct.
type IStakingModuleRedelegationEntry struct {
	CreationHeight int64
	CompletionTime string
	InitialBalance *big.Int
	SharesDst      *big.Int
	UnbondingId    uint64
}

// IStakingModuleUnbondingDelegationEntry is an auto generated low-level Go binding around an user-defined struct.
type IStakingModuleUnbondingDelegationEntry struct {
	CreationHeight int64
	CompletionTime string
	InitialBalance *big.Int
	Balance        *big.Int
	UnbondingId    uint64
}

// StakingModuleMetaData contains all meta data concerning the StakingModule contract.
var StakingModuleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int64\",\"name\":\"creationHeight\",\"type\":\"int64\"}],\"name\":\"CancelUnbondingDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CreateValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Delegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceValidator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationValidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Redelegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unbond\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"srcValidator\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dstValidator\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"beginRedelegate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"srcValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dstValidator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"beginRedelegate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"int64\",\"name\":\"creationHeight\",\"type\":\"int64\"}],\"name\":\"cancelUnbondingDelegation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"int64\",\"name\":\"creationHeight\",\"type\":\"int64\"}],\"name\":\"cancelUnbondingDelegation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"}],\"name\":\"getDelegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"getDelegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"srcValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dstValidator\",\"type\":\"address\"}],\"name\":\"getRedelegations\",\"outputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"creationHeight\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"completionTime\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sharesDst\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"unbondingId\",\"type\":\"uint64\"}],\"internalType\":\"structIStakingModule.RedelegationEntry[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"srcValidator\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dstValidator\",\"type\":\"string\"}],\"name\":\"getRedelegations\",\"outputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"creationHeight\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"completionTime\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sharesDst\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"unbondingId\",\"type\":\"uint64\"}],\"internalType\":\"structIStakingModule.RedelegationEntry[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"}],\"name\":\"getUnbondingDelegation\",\"outputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"creationHeight\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"completionTime\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"unbondingId\",\"type\":\"uint64\"}],\"internalType\":\"structIStakingModule.UnbondingDelegationEntry[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"name\":\"getUnbondingDelegation\",\"outputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"creationHeight\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"completionTime\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"unbondingId\",\"type\":\"uint64\"}],\"internalType\":\"structIStakingModule.UnbondingDelegationEntry[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"validatorAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// StakingModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingModuleMetaData.ABI instead.
var StakingModuleABI = StakingModuleMetaData.ABI

// StakingModule is an auto generated Go binding around an Ethereum contract.
type StakingModule struct {
	StakingModuleCaller     // Read-only binding to the contract
	StakingModuleTransactor // Write-only binding to the contract
	StakingModuleFilterer   // Log filterer for contract events
}

// StakingModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingModuleSession struct {
	Contract     *StakingModule    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingModuleCallerSession struct {
	Contract *StakingModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StakingModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingModuleTransactorSession struct {
	Contract     *StakingModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StakingModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingModuleRaw struct {
	Contract *StakingModule // Generic contract binding to access the raw methods on
}

// StakingModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingModuleCallerRaw struct {
	Contract *StakingModuleCaller // Generic read-only contract binding to access the raw methods on
}

// StakingModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingModuleTransactorRaw struct {
	Contract *StakingModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingModule creates a new instance of StakingModule, bound to a specific deployed contract.
func NewStakingModule(address common.Address, backend bind.ContractBackend) (*StakingModule, error) {
	contract, err := bindStakingModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingModule{StakingModuleCaller: StakingModuleCaller{contract: contract}, StakingModuleTransactor: StakingModuleTransactor{contract: contract}, StakingModuleFilterer: StakingModuleFilterer{contract: contract}}, nil
}

// NewStakingModuleCaller creates a new read-only instance of StakingModule, bound to a specific deployed contract.
func NewStakingModuleCaller(address common.Address, caller bind.ContractCaller) (*StakingModuleCaller, error) {
	contract, err := bindStakingModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingModuleCaller{contract: contract}, nil
}

// NewStakingModuleTransactor creates a new write-only instance of StakingModule, bound to a specific deployed contract.
func NewStakingModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingModuleTransactor, error) {
	contract, err := bindStakingModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingModuleTransactor{contract: contract}, nil
}

// NewStakingModuleFilterer creates a new log filterer instance of StakingModule, bound to a specific deployed contract.
func NewStakingModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingModuleFilterer, error) {
	contract, err := bindStakingModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingModuleFilterer{contract: contract}, nil
}

// bindStakingModule binds a generic wrapper to an already deployed contract.
func bindStakingModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingModuleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingModule *StakingModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingModule.Contract.StakingModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingModule *StakingModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingModule.Contract.StakingModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingModule *StakingModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingModule.Contract.StakingModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingModule *StakingModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingModule *StakingModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingModule *StakingModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingModule.Contract.contract.Transact(opts, method, params...)
}

// GetDelegation is a free data retrieval call binding the contract method 0x1d813f47.
//
// Solidity: function getDelegation(string validatorAddress) view returns(uint256)
func (_StakingModule *StakingModuleCaller) GetDelegation(opts *bind.CallOpts, validatorAddress string) (*big.Int, error) {
	var out []interface{}
	err := _StakingModule.contract.Call(opts, &out, "getDelegation", validatorAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDelegation is a free data retrieval call binding the contract method 0x1d813f47.
//
// Solidity: function getDelegation(string validatorAddress) view returns(uint256)
func (_StakingModule *StakingModuleSession) GetDelegation(validatorAddress string) (*big.Int, error) {
	return _StakingModule.Contract.GetDelegation(&_StakingModule.CallOpts, validatorAddress)
}

// GetDelegation is a free data retrieval call binding the contract method 0x1d813f47.
//
// Solidity: function getDelegation(string validatorAddress) view returns(uint256)
func (_StakingModule *StakingModuleCallerSession) GetDelegation(validatorAddress string) (*big.Int, error) {
	return _StakingModule.Contract.GetDelegation(&_StakingModule.CallOpts, validatorAddress)
}

// GetDelegation0 is a free data retrieval call binding the contract method 0x2b293768.
//
// Solidity: function getDelegation(address validatorAddress) view returns(uint256)
func (_StakingModule *StakingModuleCaller) GetDelegation0(opts *bind.CallOpts, validatorAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingModule.contract.Call(opts, &out, "getDelegation0", validatorAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDelegation0 is a free data retrieval call binding the contract method 0x2b293768.
//
// Solidity: function getDelegation(address validatorAddress) view returns(uint256)
func (_StakingModule *StakingModuleSession) GetDelegation0(validatorAddress common.Address) (*big.Int, error) {
	return _StakingModule.Contract.GetDelegation0(&_StakingModule.CallOpts, validatorAddress)
}

// GetDelegation0 is a free data retrieval call binding the contract method 0x2b293768.
//
// Solidity: function getDelegation(address validatorAddress) view returns(uint256)
func (_StakingModule *StakingModuleCallerSession) GetDelegation0(validatorAddress common.Address) (*big.Int, error) {
	return _StakingModule.Contract.GetDelegation0(&_StakingModule.CallOpts, validatorAddress)
}

// GetRedelegations is a free data retrieval call binding the contract method 0x50e99cb4.
//
// Solidity: function getRedelegations(address srcValidator, address dstValidator) view returns((int64,string,uint256,uint256,uint64)[])
func (_StakingModule *StakingModuleCaller) GetRedelegations(opts *bind.CallOpts, srcValidator common.Address, dstValidator common.Address) ([]IStakingModuleRedelegationEntry, error) {
	var out []interface{}
	err := _StakingModule.contract.Call(opts, &out, "getRedelegations", srcValidator, dstValidator)

	if err != nil {
		return *new([]IStakingModuleRedelegationEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStakingModuleRedelegationEntry)).(*[]IStakingModuleRedelegationEntry)

	return out0, err

}

// GetRedelegations is a free data retrieval call binding the contract method 0x50e99cb4.
//
// Solidity: function getRedelegations(address srcValidator, address dstValidator) view returns((int64,string,uint256,uint256,uint64)[])
func (_StakingModule *StakingModuleSession) GetRedelegations(srcValidator common.Address, dstValidator common.Address) ([]IStakingModuleRedelegationEntry, error) {
	return _StakingModule.Contract.GetRedelegations(&_StakingModule.CallOpts, srcValidator, dstValidator)
}

// GetRedelegations is a free data retrieval call binding the contract method 0x50e99cb4.
//
// Solidity: function getRedelegations(address srcValidator, address dstValidator) view returns((int64,string,uint256,uint256,uint64)[])
func (_StakingModule *StakingModuleCallerSession) GetRedelegations(srcValidator common.Address, dstValidator common.Address) ([]IStakingModuleRedelegationEntry, error) {
	return _StakingModule.Contract.GetRedelegations(&_StakingModule.CallOpts, srcValidator, dstValidator)
}

// GetRedelegations0 is a free data retrieval call binding the contract method 0x91f4c8df.
//
// Solidity: function getRedelegations(string srcValidator, string dstValidator) view returns((int64,string,uint256,uint256,uint64)[])
func (_StakingModule *StakingModuleCaller) GetRedelegations0(opts *bind.CallOpts, srcValidator string, dstValidator string) ([]IStakingModuleRedelegationEntry, error) {
	var out []interface{}
	err := _StakingModule.contract.Call(opts, &out, "getRedelegations0", srcValidator, dstValidator)

	if err != nil {
		return *new([]IStakingModuleRedelegationEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStakingModuleRedelegationEntry)).(*[]IStakingModuleRedelegationEntry)

	return out0, err

}

// GetRedelegations0 is a free data retrieval call binding the contract method 0x91f4c8df.
//
// Solidity: function getRedelegations(string srcValidator, string dstValidator) view returns((int64,string,uint256,uint256,uint64)[])
func (_StakingModule *StakingModuleSession) GetRedelegations0(srcValidator string, dstValidator string) ([]IStakingModuleRedelegationEntry, error) {
	return _StakingModule.Contract.GetRedelegations0(&_StakingModule.CallOpts, srcValidator, dstValidator)
}

// GetRedelegations0 is a free data retrieval call binding the contract method 0x91f4c8df.
//
// Solidity: function getRedelegations(string srcValidator, string dstValidator) view returns((int64,string,uint256,uint256,uint64)[])
func (_StakingModule *StakingModuleCallerSession) GetRedelegations0(srcValidator string, dstValidator string) ([]IStakingModuleRedelegationEntry, error) {
	return _StakingModule.Contract.GetRedelegations0(&_StakingModule.CallOpts, srcValidator, dstValidator)
}

// GetUnbondingDelegation is a free data retrieval call binding the contract method 0x65abd5bf.
//
// Solidity: function getUnbondingDelegation(string validatorAddress) view returns((int64,string,uint256,uint256,uint64)[])
func (_StakingModule *StakingModuleCaller) GetUnbondingDelegation(opts *bind.CallOpts, validatorAddress string) ([]IStakingModuleUnbondingDelegationEntry, error) {
	var out []interface{}
	err := _StakingModule.contract.Call(opts, &out, "getUnbondingDelegation", validatorAddress)

	if err != nil {
		return *new([]IStakingModuleUnbondingDelegationEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStakingModuleUnbondingDelegationEntry)).(*[]IStakingModuleUnbondingDelegationEntry)

	return out0, err

}

// GetUnbondingDelegation is a free data retrieval call binding the contract method 0x65abd5bf.
//
// Solidity: function getUnbondingDelegation(string validatorAddress) view returns((int64,string,uint256,uint256,uint64)[])
func (_StakingModule *StakingModuleSession) GetUnbondingDelegation(validatorAddress string) ([]IStakingModuleUnbondingDelegationEntry, error) {
	return _StakingModule.Contract.GetUnbondingDelegation(&_StakingModule.CallOpts, validatorAddress)
}

// GetUnbondingDelegation is a free data retrieval call binding the contract method 0x65abd5bf.
//
// Solidity: function getUnbondingDelegation(string validatorAddress) view returns((int64,string,uint256,uint256,uint64)[])
func (_StakingModule *StakingModuleCallerSession) GetUnbondingDelegation(validatorAddress string) ([]IStakingModuleUnbondingDelegationEntry, error) {
	return _StakingModule.Contract.GetUnbondingDelegation(&_StakingModule.CallOpts, validatorAddress)
}

// GetUnbondingDelegation0 is a free data retrieval call binding the contract method 0xcd48a0b7.
//
// Solidity: function getUnbondingDelegation(address validatorAddress) view returns((int64,string,uint256,uint256,uint64)[])
func (_StakingModule *StakingModuleCaller) GetUnbondingDelegation0(opts *bind.CallOpts, validatorAddress common.Address) ([]IStakingModuleUnbondingDelegationEntry, error) {
	var out []interface{}
	err := _StakingModule.contract.Call(opts, &out, "getUnbondingDelegation0", validatorAddress)

	if err != nil {
		return *new([]IStakingModuleUnbondingDelegationEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStakingModuleUnbondingDelegationEntry)).(*[]IStakingModuleUnbondingDelegationEntry)

	return out0, err

}

// GetUnbondingDelegation0 is a free data retrieval call binding the contract method 0xcd48a0b7.
//
// Solidity: function getUnbondingDelegation(address validatorAddress) view returns((int64,string,uint256,uint256,uint64)[])
func (_StakingModule *StakingModuleSession) GetUnbondingDelegation0(validatorAddress common.Address) ([]IStakingModuleUnbondingDelegationEntry, error) {
	return _StakingModule.Contract.GetUnbondingDelegation0(&_StakingModule.CallOpts, validatorAddress)
}

// GetUnbondingDelegation0 is a free data retrieval call binding the contract method 0xcd48a0b7.
//
// Solidity: function getUnbondingDelegation(address validatorAddress) view returns((int64,string,uint256,uint256,uint64)[])
func (_StakingModule *StakingModuleCallerSession) GetUnbondingDelegation0(validatorAddress common.Address) ([]IStakingModuleUnbondingDelegationEntry, error) {
	return _StakingModule.Contract.GetUnbondingDelegation0(&_StakingModule.CallOpts, validatorAddress)
}

// BeginRedelegate is a paid mutator transaction binding the contract method 0x2e436cf2.
//
// Solidity: function beginRedelegate(string srcValidator, string dstValidator, uint256 amount) payable returns()
func (_StakingModule *StakingModuleTransactor) BeginRedelegate(opts *bind.TransactOpts, srcValidator string, dstValidator string, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.contract.Transact(opts, "beginRedelegate", srcValidator, dstValidator, amount)
}

// BeginRedelegate is a paid mutator transaction binding the contract method 0x2e436cf2.
//
// Solidity: function beginRedelegate(string srcValidator, string dstValidator, uint256 amount) payable returns()
func (_StakingModule *StakingModuleSession) BeginRedelegate(srcValidator string, dstValidator string, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.Contract.BeginRedelegate(&_StakingModule.TransactOpts, srcValidator, dstValidator, amount)
}

// BeginRedelegate is a paid mutator transaction binding the contract method 0x2e436cf2.
//
// Solidity: function beginRedelegate(string srcValidator, string dstValidator, uint256 amount) payable returns()
func (_StakingModule *StakingModuleTransactorSession) BeginRedelegate(srcValidator string, dstValidator string, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.Contract.BeginRedelegate(&_StakingModule.TransactOpts, srcValidator, dstValidator, amount)
}

// BeginRedelegate0 is a paid mutator transaction binding the contract method 0xb3a8ae3b.
//
// Solidity: function beginRedelegate(address srcValidator, address dstValidator, uint256 amount) payable returns()
func (_StakingModule *StakingModuleTransactor) BeginRedelegate0(opts *bind.TransactOpts, srcValidator common.Address, dstValidator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.contract.Transact(opts, "beginRedelegate0", srcValidator, dstValidator, amount)
}

// BeginRedelegate0 is a paid mutator transaction binding the contract method 0xb3a8ae3b.
//
// Solidity: function beginRedelegate(address srcValidator, address dstValidator, uint256 amount) payable returns()
func (_StakingModule *StakingModuleSession) BeginRedelegate0(srcValidator common.Address, dstValidator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.Contract.BeginRedelegate0(&_StakingModule.TransactOpts, srcValidator, dstValidator, amount)
}

// BeginRedelegate0 is a paid mutator transaction binding the contract method 0xb3a8ae3b.
//
// Solidity: function beginRedelegate(address srcValidator, address dstValidator, uint256 amount) payable returns()
func (_StakingModule *StakingModuleTransactorSession) BeginRedelegate0(srcValidator common.Address, dstValidator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.Contract.BeginRedelegate0(&_StakingModule.TransactOpts, srcValidator, dstValidator, amount)
}

// CancelUnbondingDelegation is a paid mutator transaction binding the contract method 0x69a2f536.
//
// Solidity: function cancelUnbondingDelegation(address validatorAddress, uint256 amount, int64 creationHeight) payable returns()
func (_StakingModule *StakingModuleTransactor) CancelUnbondingDelegation(opts *bind.TransactOpts, validatorAddress common.Address, amount *big.Int, creationHeight int64) (*types.Transaction, error) {
	return _StakingModule.contract.Transact(opts, "cancelUnbondingDelegation", validatorAddress, amount, creationHeight)
}

// CancelUnbondingDelegation is a paid mutator transaction binding the contract method 0x69a2f536.
//
// Solidity: function cancelUnbondingDelegation(address validatorAddress, uint256 amount, int64 creationHeight) payable returns()
func (_StakingModule *StakingModuleSession) CancelUnbondingDelegation(validatorAddress common.Address, amount *big.Int, creationHeight int64) (*types.Transaction, error) {
	return _StakingModule.Contract.CancelUnbondingDelegation(&_StakingModule.TransactOpts, validatorAddress, amount, creationHeight)
}

// CancelUnbondingDelegation is a paid mutator transaction binding the contract method 0x69a2f536.
//
// Solidity: function cancelUnbondingDelegation(address validatorAddress, uint256 amount, int64 creationHeight) payable returns()
func (_StakingModule *StakingModuleTransactorSession) CancelUnbondingDelegation(validatorAddress common.Address, amount *big.Int, creationHeight int64) (*types.Transaction, error) {
	return _StakingModule.Contract.CancelUnbondingDelegation(&_StakingModule.TransactOpts, validatorAddress, amount, creationHeight)
}

// CancelUnbondingDelegation0 is a paid mutator transaction binding the contract method 0xab0341d3.
//
// Solidity: function cancelUnbondingDelegation(string validatorAddress, uint256 amount, int64 creationHeight) payable returns()
func (_StakingModule *StakingModuleTransactor) CancelUnbondingDelegation0(opts *bind.TransactOpts, validatorAddress string, amount *big.Int, creationHeight int64) (*types.Transaction, error) {
	return _StakingModule.contract.Transact(opts, "cancelUnbondingDelegation0", validatorAddress, amount, creationHeight)
}

// CancelUnbondingDelegation0 is a paid mutator transaction binding the contract method 0xab0341d3.
//
// Solidity: function cancelUnbondingDelegation(string validatorAddress, uint256 amount, int64 creationHeight) payable returns()
func (_StakingModule *StakingModuleSession) CancelUnbondingDelegation0(validatorAddress string, amount *big.Int, creationHeight int64) (*types.Transaction, error) {
	return _StakingModule.Contract.CancelUnbondingDelegation0(&_StakingModule.TransactOpts, validatorAddress, amount, creationHeight)
}

// CancelUnbondingDelegation0 is a paid mutator transaction binding the contract method 0xab0341d3.
//
// Solidity: function cancelUnbondingDelegation(string validatorAddress, uint256 amount, int64 creationHeight) payable returns()
func (_StakingModule *StakingModuleTransactorSession) CancelUnbondingDelegation0(validatorAddress string, amount *big.Int, creationHeight int64) (*types.Transaction, error) {
	return _StakingModule.Contract.CancelUnbondingDelegation0(&_StakingModule.TransactOpts, validatorAddress, amount, creationHeight)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address validatorAddress, uint256 amount) payable returns()
func (_StakingModule *StakingModuleTransactor) Delegate(opts *bind.TransactOpts, validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.contract.Transact(opts, "delegate", validatorAddress, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address validatorAddress, uint256 amount) payable returns()
func (_StakingModule *StakingModuleSession) Delegate(validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.Contract.Delegate(&_StakingModule.TransactOpts, validatorAddress, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address validatorAddress, uint256 amount) payable returns()
func (_StakingModule *StakingModuleTransactorSession) Delegate(validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.Contract.Delegate(&_StakingModule.TransactOpts, validatorAddress, amount)
}

// Delegate0 is a paid mutator transaction binding the contract method 0x03f24de1.
//
// Solidity: function delegate(string validatorAddress, uint256 amount) payable returns()
func (_StakingModule *StakingModuleTransactor) Delegate0(opts *bind.TransactOpts, validatorAddress string, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.contract.Transact(opts, "delegate0", validatorAddress, amount)
}

// Delegate0 is a paid mutator transaction binding the contract method 0x03f24de1.
//
// Solidity: function delegate(string validatorAddress, uint256 amount) payable returns()
func (_StakingModule *StakingModuleSession) Delegate0(validatorAddress string, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.Contract.Delegate0(&_StakingModule.TransactOpts, validatorAddress, amount)
}

// Delegate0 is a paid mutator transaction binding the contract method 0x03f24de1.
//
// Solidity: function delegate(string validatorAddress, uint256 amount) payable returns()
func (_StakingModule *StakingModuleTransactorSession) Delegate0(validatorAddress string, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.Contract.Delegate0(&_StakingModule.TransactOpts, validatorAddress, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validatorAddress, uint256 amount) payable returns()
func (_StakingModule *StakingModuleTransactor) Undelegate(opts *bind.TransactOpts, validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.contract.Transact(opts, "undelegate", validatorAddress, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validatorAddress, uint256 amount) payable returns()
func (_StakingModule *StakingModuleSession) Undelegate(validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.Contract.Undelegate(&_StakingModule.TransactOpts, validatorAddress, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validatorAddress, uint256 amount) payable returns()
func (_StakingModule *StakingModuleTransactorSession) Undelegate(validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.Contract.Undelegate(&_StakingModule.TransactOpts, validatorAddress, amount)
}

// Undelegate0 is a paid mutator transaction binding the contract method 0x8dfc8897.
//
// Solidity: function undelegate(string validatorAddress, uint256 amount) payable returns()
func (_StakingModule *StakingModuleTransactor) Undelegate0(opts *bind.TransactOpts, validatorAddress string, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.contract.Transact(opts, "undelegate0", validatorAddress, amount)
}

// Undelegate0 is a paid mutator transaction binding the contract method 0x8dfc8897.
//
// Solidity: function undelegate(string validatorAddress, uint256 amount) payable returns()
func (_StakingModule *StakingModuleSession) Undelegate0(validatorAddress string, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.Contract.Undelegate0(&_StakingModule.TransactOpts, validatorAddress, amount)
}

// Undelegate0 is a paid mutator transaction binding the contract method 0x8dfc8897.
//
// Solidity: function undelegate(string validatorAddress, uint256 amount) payable returns()
func (_StakingModule *StakingModuleTransactorSession) Undelegate0(validatorAddress string, amount *big.Int) (*types.Transaction, error) {
	return _StakingModule.Contract.Undelegate0(&_StakingModule.TransactOpts, validatorAddress, amount)
}

// StakingModuleCancelUnbondingDelegationIterator is returned from FilterCancelUnbondingDelegation and is used to iterate over the raw logs and unpacked data for CancelUnbondingDelegation events raised by the StakingModule contract.
type StakingModuleCancelUnbondingDelegationIterator struct {
	Event *StakingModuleCancelUnbondingDelegation // Event containing the contract specifics and raw log

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
func (it *StakingModuleCancelUnbondingDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingModuleCancelUnbondingDelegation)
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
		it.Event = new(StakingModuleCancelUnbondingDelegation)
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
func (it *StakingModuleCancelUnbondingDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingModuleCancelUnbondingDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingModuleCancelUnbondingDelegation represents a CancelUnbondingDelegation event raised by the StakingModule contract.
type StakingModuleCancelUnbondingDelegation struct {
	Validator      common.Address
	Delegator      common.Address
	Amount         *big.Int
	CreationHeight int64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCancelUnbondingDelegation is a free log retrieval operation binding the contract event 0x4022879e779e0c889f536ac78c7e10619fc795d7af0efabf4500309a187588d9.
//
// Solidity: event CancelUnbondingDelegation(address indexed validator, address indexed delegator, uint256 amount, int64 creationHeight)
func (_StakingModule *StakingModuleFilterer) FilterCancelUnbondingDelegation(opts *bind.FilterOpts, validator []common.Address, delegator []common.Address) (*StakingModuleCancelUnbondingDelegationIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _StakingModule.contract.FilterLogs(opts, "CancelUnbondingDelegation", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingModuleCancelUnbondingDelegationIterator{contract: _StakingModule.contract, event: "CancelUnbondingDelegation", logs: logs, sub: sub}, nil
}

// WatchCancelUnbondingDelegation is a free log subscription operation binding the contract event 0x4022879e779e0c889f536ac78c7e10619fc795d7af0efabf4500309a187588d9.
//
// Solidity: event CancelUnbondingDelegation(address indexed validator, address indexed delegator, uint256 amount, int64 creationHeight)
func (_StakingModule *StakingModuleFilterer) WatchCancelUnbondingDelegation(opts *bind.WatchOpts, sink chan<- *StakingModuleCancelUnbondingDelegation, validator []common.Address, delegator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _StakingModule.contract.WatchLogs(opts, "CancelUnbondingDelegation", validatorRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingModuleCancelUnbondingDelegation)
				if err := _StakingModule.contract.UnpackLog(event, "CancelUnbondingDelegation", log); err != nil {
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

// ParseCancelUnbondingDelegation is a log parse operation binding the contract event 0x4022879e779e0c889f536ac78c7e10619fc795d7af0efabf4500309a187588d9.
//
// Solidity: event CancelUnbondingDelegation(address indexed validator, address indexed delegator, uint256 amount, int64 creationHeight)
func (_StakingModule *StakingModuleFilterer) ParseCancelUnbondingDelegation(log types.Log) (*StakingModuleCancelUnbondingDelegation, error) {
	event := new(StakingModuleCancelUnbondingDelegation)
	if err := _StakingModule.contract.UnpackLog(event, "CancelUnbondingDelegation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingModuleCreateValidatorIterator is returned from FilterCreateValidator and is used to iterate over the raw logs and unpacked data for CreateValidator events raised by the StakingModule contract.
type StakingModuleCreateValidatorIterator struct {
	Event *StakingModuleCreateValidator // Event containing the contract specifics and raw log

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
func (it *StakingModuleCreateValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingModuleCreateValidator)
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
		it.Event = new(StakingModuleCreateValidator)
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
func (it *StakingModuleCreateValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingModuleCreateValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingModuleCreateValidator represents a CreateValidator event raised by the StakingModule contract.
type StakingModuleCreateValidator struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCreateValidator is a free log retrieval operation binding the contract event 0x9bdb560f8135cb46033a55410c14e14b1a7bc2d3f3e9973f4b49533e176468b0.
//
// Solidity: event CreateValidator(address indexed validator, uint256 amount)
func (_StakingModule *StakingModuleFilterer) FilterCreateValidator(opts *bind.FilterOpts, validator []common.Address) (*StakingModuleCreateValidatorIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _StakingModule.contract.FilterLogs(opts, "CreateValidator", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingModuleCreateValidatorIterator{contract: _StakingModule.contract, event: "CreateValidator", logs: logs, sub: sub}, nil
}

// WatchCreateValidator is a free log subscription operation binding the contract event 0x9bdb560f8135cb46033a55410c14e14b1a7bc2d3f3e9973f4b49533e176468b0.
//
// Solidity: event CreateValidator(address indexed validator, uint256 amount)
func (_StakingModule *StakingModuleFilterer) WatchCreateValidator(opts *bind.WatchOpts, sink chan<- *StakingModuleCreateValidator, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _StakingModule.contract.WatchLogs(opts, "CreateValidator", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingModuleCreateValidator)
				if err := _StakingModule.contract.UnpackLog(event, "CreateValidator", log); err != nil {
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

// ParseCreateValidator is a log parse operation binding the contract event 0x9bdb560f8135cb46033a55410c14e14b1a7bc2d3f3e9973f4b49533e176468b0.
//
// Solidity: event CreateValidator(address indexed validator, uint256 amount)
func (_StakingModule *StakingModuleFilterer) ParseCreateValidator(log types.Log) (*StakingModuleCreateValidator, error) {
	event := new(StakingModuleCreateValidator)
	if err := _StakingModule.contract.UnpackLog(event, "CreateValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingModuleDelegateIterator is returned from FilterDelegate and is used to iterate over the raw logs and unpacked data for Delegate events raised by the StakingModule contract.
type StakingModuleDelegateIterator struct {
	Event *StakingModuleDelegate // Event containing the contract specifics and raw log

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
func (it *StakingModuleDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingModuleDelegate)
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
		it.Event = new(StakingModuleDelegate)
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
func (it *StakingModuleDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingModuleDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingModuleDelegate represents a Delegate event raised by the StakingModule contract.
type StakingModuleDelegate struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegate is a free log retrieval operation binding the contract event 0xb0d234274aef7a61aa5a2eb44c23881ebf46a068cccbd413c978bcbd555fe17f.
//
// Solidity: event Delegate(address indexed validator, uint256 amount)
func (_StakingModule *StakingModuleFilterer) FilterDelegate(opts *bind.FilterOpts, validator []common.Address) (*StakingModuleDelegateIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _StakingModule.contract.FilterLogs(opts, "Delegate", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingModuleDelegateIterator{contract: _StakingModule.contract, event: "Delegate", logs: logs, sub: sub}, nil
}

// WatchDelegate is a free log subscription operation binding the contract event 0xb0d234274aef7a61aa5a2eb44c23881ebf46a068cccbd413c978bcbd555fe17f.
//
// Solidity: event Delegate(address indexed validator, uint256 amount)
func (_StakingModule *StakingModuleFilterer) WatchDelegate(opts *bind.WatchOpts, sink chan<- *StakingModuleDelegate, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _StakingModule.contract.WatchLogs(opts, "Delegate", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingModuleDelegate)
				if err := _StakingModule.contract.UnpackLog(event, "Delegate", log); err != nil {
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

// ParseDelegate is a log parse operation binding the contract event 0xb0d234274aef7a61aa5a2eb44c23881ebf46a068cccbd413c978bcbd555fe17f.
//
// Solidity: event Delegate(address indexed validator, uint256 amount)
func (_StakingModule *StakingModuleFilterer) ParseDelegate(log types.Log) (*StakingModuleDelegate, error) {
	event := new(StakingModuleDelegate)
	if err := _StakingModule.contract.UnpackLog(event, "Delegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingModuleRedelegateIterator is returned from FilterRedelegate and is used to iterate over the raw logs and unpacked data for Redelegate events raised by the StakingModule contract.
type StakingModuleRedelegateIterator struct {
	Event *StakingModuleRedelegate // Event containing the contract specifics and raw log

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
func (it *StakingModuleRedelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingModuleRedelegate)
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
		it.Event = new(StakingModuleRedelegate)
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
func (it *StakingModuleRedelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingModuleRedelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingModuleRedelegate represents a Redelegate event raised by the StakingModule contract.
type StakingModuleRedelegate struct {
	SourceValidator      common.Address
	DestinationValidator common.Address
	Amount               *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterRedelegate is a free log retrieval operation binding the contract event 0x7eb58ff7a8b2b30f490a329e171138f209b8ccd6c3c8d1d8b7bdf309252e2603.
//
// Solidity: event Redelegate(address indexed sourceValidator, address indexed destinationValidator, uint256 amount)
func (_StakingModule *StakingModuleFilterer) FilterRedelegate(opts *bind.FilterOpts, sourceValidator []common.Address, destinationValidator []common.Address) (*StakingModuleRedelegateIterator, error) {

	var sourceValidatorRule []interface{}
	for _, sourceValidatorItem := range sourceValidator {
		sourceValidatorRule = append(sourceValidatorRule, sourceValidatorItem)
	}
	var destinationValidatorRule []interface{}
	for _, destinationValidatorItem := range destinationValidator {
		destinationValidatorRule = append(destinationValidatorRule, destinationValidatorItem)
	}

	logs, sub, err := _StakingModule.contract.FilterLogs(opts, "Redelegate", sourceValidatorRule, destinationValidatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingModuleRedelegateIterator{contract: _StakingModule.contract, event: "Redelegate", logs: logs, sub: sub}, nil
}

// WatchRedelegate is a free log subscription operation binding the contract event 0x7eb58ff7a8b2b30f490a329e171138f209b8ccd6c3c8d1d8b7bdf309252e2603.
//
// Solidity: event Redelegate(address indexed sourceValidator, address indexed destinationValidator, uint256 amount)
func (_StakingModule *StakingModuleFilterer) WatchRedelegate(opts *bind.WatchOpts, sink chan<- *StakingModuleRedelegate, sourceValidator []common.Address, destinationValidator []common.Address) (event.Subscription, error) {

	var sourceValidatorRule []interface{}
	for _, sourceValidatorItem := range sourceValidator {
		sourceValidatorRule = append(sourceValidatorRule, sourceValidatorItem)
	}
	var destinationValidatorRule []interface{}
	for _, destinationValidatorItem := range destinationValidator {
		destinationValidatorRule = append(destinationValidatorRule, destinationValidatorItem)
	}

	logs, sub, err := _StakingModule.contract.WatchLogs(opts, "Redelegate", sourceValidatorRule, destinationValidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingModuleRedelegate)
				if err := _StakingModule.contract.UnpackLog(event, "Redelegate", log); err != nil {
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

// ParseRedelegate is a log parse operation binding the contract event 0x7eb58ff7a8b2b30f490a329e171138f209b8ccd6c3c8d1d8b7bdf309252e2603.
//
// Solidity: event Redelegate(address indexed sourceValidator, address indexed destinationValidator, uint256 amount)
func (_StakingModule *StakingModuleFilterer) ParseRedelegate(log types.Log) (*StakingModuleRedelegate, error) {
	event := new(StakingModuleRedelegate)
	if err := _StakingModule.contract.UnpackLog(event, "Redelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingModuleUnbondIterator is returned from FilterUnbond and is used to iterate over the raw logs and unpacked data for Unbond events raised by the StakingModule contract.
type StakingModuleUnbondIterator struct {
	Event *StakingModuleUnbond // Event containing the contract specifics and raw log

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
func (it *StakingModuleUnbondIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingModuleUnbond)
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
		it.Event = new(StakingModuleUnbond)
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
func (it *StakingModuleUnbondIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingModuleUnbondIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingModuleUnbond represents a Unbond event raised by the StakingModule contract.
type StakingModuleUnbond struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnbond is a free log retrieval operation binding the contract event 0xb735793e7250527232d5f7e07e8691c383c85e0c77673ac750944e8bd40e71c8.
//
// Solidity: event Unbond(address indexed validator, uint256 amount)
func (_StakingModule *StakingModuleFilterer) FilterUnbond(opts *bind.FilterOpts, validator []common.Address) (*StakingModuleUnbondIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _StakingModule.contract.FilterLogs(opts, "Unbond", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingModuleUnbondIterator{contract: _StakingModule.contract, event: "Unbond", logs: logs, sub: sub}, nil
}

// WatchUnbond is a free log subscription operation binding the contract event 0xb735793e7250527232d5f7e07e8691c383c85e0c77673ac750944e8bd40e71c8.
//
// Solidity: event Unbond(address indexed validator, uint256 amount)
func (_StakingModule *StakingModuleFilterer) WatchUnbond(opts *bind.WatchOpts, sink chan<- *StakingModuleUnbond, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _StakingModule.contract.WatchLogs(opts, "Unbond", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingModuleUnbond)
				if err := _StakingModule.contract.UnpackLog(event, "Unbond", log); err != nil {
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

// ParseUnbond is a log parse operation binding the contract event 0xb735793e7250527232d5f7e07e8691c383c85e0c77673ac750944e8bd40e71c8.
//
// Solidity: event Unbond(address indexed validator, uint256 amount)
func (_StakingModule *StakingModuleFilterer) ParseUnbond(log types.Log) (*StakingModuleUnbond, error) {
	event := new(StakingModuleUnbond)
	if err := _StakingModule.contract.UnpackLog(event, "Unbond", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
