// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testing

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

// LoadTestOpsMetaData contains all meta data concerning the LoadTestOps contract.
var LoadTestOpsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"loadData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260015f553480156012575f80fd5b5060788061001f5f395ff3fe6080604052348015600e575f80fd5b50600436106026575f3560e01c8063100b815d14602a575b5f80fd5b60306032565b005b60405160608152602081205f555056fea2646970667358221220d97583e2ab82a155b3ab643b1f9c4de9c067816ecf83449a8504450a69a6bd9964736f6c63430008160033",
}

// LoadTestOpsABI is the input ABI used to generate the binding from.
// Deprecated: Use LoadTestOpsMetaData.ABI instead.
var LoadTestOpsABI = LoadTestOpsMetaData.ABI

// LoadTestOpsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LoadTestOpsMetaData.Bin instead.
var LoadTestOpsBin = LoadTestOpsMetaData.Bin

// DeployLoadTestOps deploys a new Ethereum contract, binding an instance of LoadTestOps to it.
func DeployLoadTestOps(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LoadTestOps, error) {
	parsed, err := LoadTestOpsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LoadTestOpsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LoadTestOps{LoadTestOpsCaller: LoadTestOpsCaller{contract: contract}, LoadTestOpsTransactor: LoadTestOpsTransactor{contract: contract}, LoadTestOpsFilterer: LoadTestOpsFilterer{contract: contract}}, nil
}

// LoadTestOps is an auto generated Go binding around an Ethereum contract.
type LoadTestOps struct {
	LoadTestOpsCaller     // Read-only binding to the contract
	LoadTestOpsTransactor // Write-only binding to the contract
	LoadTestOpsFilterer   // Log filterer for contract events
}

// LoadTestOpsCaller is an auto generated read-only Go binding around an Ethereum contract.
type LoadTestOpsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoadTestOpsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LoadTestOpsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoadTestOpsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LoadTestOpsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoadTestOpsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LoadTestOpsSession struct {
	Contract     *LoadTestOps      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoadTestOpsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LoadTestOpsCallerSession struct {
	Contract *LoadTestOpsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// LoadTestOpsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LoadTestOpsTransactorSession struct {
	Contract     *LoadTestOpsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// LoadTestOpsRaw is an auto generated low-level Go binding around an Ethereum contract.
type LoadTestOpsRaw struct {
	Contract *LoadTestOps // Generic contract binding to access the raw methods on
}

// LoadTestOpsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LoadTestOpsCallerRaw struct {
	Contract *LoadTestOpsCaller // Generic read-only contract binding to access the raw methods on
}

// LoadTestOpsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LoadTestOpsTransactorRaw struct {
	Contract *LoadTestOpsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLoadTestOps creates a new instance of LoadTestOps, bound to a specific deployed contract.
func NewLoadTestOps(address common.Address, backend bind.ContractBackend) (*LoadTestOps, error) {
	contract, err := bindLoadTestOps(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LoadTestOps{LoadTestOpsCaller: LoadTestOpsCaller{contract: contract}, LoadTestOpsTransactor: LoadTestOpsTransactor{contract: contract}, LoadTestOpsFilterer: LoadTestOpsFilterer{contract: contract}}, nil
}

// NewLoadTestOpsCaller creates a new read-only instance of LoadTestOps, bound to a specific deployed contract.
func NewLoadTestOpsCaller(address common.Address, caller bind.ContractCaller) (*LoadTestOpsCaller, error) {
	contract, err := bindLoadTestOps(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LoadTestOpsCaller{contract: contract}, nil
}

// NewLoadTestOpsTransactor creates a new write-only instance of LoadTestOps, bound to a specific deployed contract.
func NewLoadTestOpsTransactor(address common.Address, transactor bind.ContractTransactor) (*LoadTestOpsTransactor, error) {
	contract, err := bindLoadTestOps(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LoadTestOpsTransactor{contract: contract}, nil
}

// NewLoadTestOpsFilterer creates a new log filterer instance of LoadTestOps, bound to a specific deployed contract.
func NewLoadTestOpsFilterer(address common.Address, filterer bind.ContractFilterer) (*LoadTestOpsFilterer, error) {
	contract, err := bindLoadTestOps(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LoadTestOpsFilterer{contract: contract}, nil
}

// bindLoadTestOps binds a generic wrapper to an already deployed contract.
func bindLoadTestOps(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LoadTestOpsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LoadTestOps *LoadTestOpsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LoadTestOps.Contract.LoadTestOpsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LoadTestOps *LoadTestOpsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoadTestOps.Contract.LoadTestOpsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LoadTestOps *LoadTestOpsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LoadTestOps.Contract.LoadTestOpsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LoadTestOps *LoadTestOpsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LoadTestOps.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LoadTestOps *LoadTestOpsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoadTestOps.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LoadTestOps *LoadTestOpsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LoadTestOps.Contract.contract.Transact(opts, method, params...)
}

// LoadData is a paid mutator transaction binding the contract method 0x100b815d.
//
// Solidity: function loadData() returns()
func (_LoadTestOps *LoadTestOpsTransactor) LoadData(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoadTestOps.contract.Transact(opts, "loadData")
}

// LoadData is a paid mutator transaction binding the contract method 0x100b815d.
//
// Solidity: function loadData() returns()
func (_LoadTestOps *LoadTestOpsSession) LoadData() (*types.Transaction, error) {
	return _LoadTestOps.Contract.LoadData(&_LoadTestOps.TransactOpts)
}

// LoadData is a paid mutator transaction binding the contract method 0x100b815d.
//
// Solidity: function loadData() returns()
func (_LoadTestOps *LoadTestOpsTransactorSession) LoadData() (*types.Transaction, error) {
	return _LoadTestOps.Contract.LoadData(&_LoadTestOps.TransactOpts)
}
