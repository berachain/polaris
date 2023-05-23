// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking

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

// MintModuleMetaData contains all meta data concerning the MintModule contract.
var MintModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"annualProvisions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inflation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MintModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use MintModuleMetaData.ABI instead.
var MintModuleABI = MintModuleMetaData.ABI

// MintModule is an auto generated Go binding around an Ethereum contract.
type MintModule struct {
	MintModuleCaller     // Read-only binding to the contract
	MintModuleTransactor // Write-only binding to the contract
	MintModuleFilterer   // Log filterer for contract events
}

// MintModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type MintModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MintModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MintModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MintModuleSession struct {
	Contract     *MintModule       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MintModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MintModuleCallerSession struct {
	Contract *MintModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MintModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MintModuleTransactorSession struct {
	Contract     *MintModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MintModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type MintModuleRaw struct {
	Contract *MintModule // Generic contract binding to access the raw methods on
}

// MintModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MintModuleCallerRaw struct {
	Contract *MintModuleCaller // Generic read-only contract binding to access the raw methods on
}

// MintModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MintModuleTransactorRaw struct {
	Contract *MintModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMintModule creates a new instance of MintModule, bound to a specific deployed contract.
func NewMintModule(address common.Address, backend bind.ContractBackend) (*MintModule, error) {
	contract, err := bindMintModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MintModule{MintModuleCaller: MintModuleCaller{contract: contract}, MintModuleTransactor: MintModuleTransactor{contract: contract}, MintModuleFilterer: MintModuleFilterer{contract: contract}}, nil
}

// NewMintModuleCaller creates a new read-only instance of MintModule, bound to a specific deployed contract.
func NewMintModuleCaller(address common.Address, caller bind.ContractCaller) (*MintModuleCaller, error) {
	contract, err := bindMintModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MintModuleCaller{contract: contract}, nil
}

// NewMintModuleTransactor creates a new write-only instance of MintModule, bound to a specific deployed contract.
func NewMintModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*MintModuleTransactor, error) {
	contract, err := bindMintModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MintModuleTransactor{contract: contract}, nil
}

// NewMintModuleFilterer creates a new log filterer instance of MintModule, bound to a specific deployed contract.
func NewMintModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*MintModuleFilterer, error) {
	contract, err := bindMintModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MintModuleFilterer{contract: contract}, nil
}

// bindMintModule binds a generic wrapper to an already deployed contract.
func bindMintModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MintModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintModule *MintModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MintModule.Contract.MintModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintModule *MintModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintModule.Contract.MintModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintModule *MintModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintModule.Contract.MintModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintModule *MintModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MintModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintModule *MintModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintModule *MintModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintModule.Contract.contract.Transact(opts, method, params...)
}

// AnnualProvisions is a free data retrieval call binding the contract method 0x14aa16c7.
//
// Solidity: function annualProvisions() view returns(uint256)
func (_MintModule *MintModuleCaller) AnnualProvisions(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MintModule.contract.Call(opts, &out, "annualProvisions")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AnnualProvisions is a free data retrieval call binding the contract method 0x14aa16c7.
//
// Solidity: function annualProvisions() view returns(uint256)
func (_MintModule *MintModuleSession) AnnualProvisions() (*big.Int, error) {
	return _MintModule.Contract.AnnualProvisions(&_MintModule.CallOpts)
}

// AnnualProvisions is a free data retrieval call binding the contract method 0x14aa16c7.
//
// Solidity: function annualProvisions() view returns(uint256)
func (_MintModule *MintModuleCallerSession) AnnualProvisions() (*big.Int, error) {
	return _MintModule.Contract.AnnualProvisions(&_MintModule.CallOpts)
}

// Inflation is a free data retrieval call binding the contract method 0xbe0522e0.
//
// Solidity: function inflation() view returns(uint256)
func (_MintModule *MintModuleCaller) Inflation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MintModule.contract.Call(opts, &out, "inflation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Inflation is a free data retrieval call binding the contract method 0xbe0522e0.
//
// Solidity: function inflation() view returns(uint256)
func (_MintModule *MintModuleSession) Inflation() (*big.Int, error) {
	return _MintModule.Contract.Inflation(&_MintModule.CallOpts)
}

// Inflation is a free data retrieval call binding the contract method 0xbe0522e0.
//
// Solidity: function inflation() view returns(uint256)
func (_MintModule *MintModuleCallerSession) Inflation() (*big.Int, error) {
	return _MintModule.Contract.Inflation(&_MintModule.CallOpts)
}
