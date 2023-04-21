// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package precompile

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

// BuilderModuleMetaData contains all meta data concerning the BuilderModule contract.
var BuilderModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactions\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"timeout\",\"type\":\"uint64\"}],\"name\":\"auctionBid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// BuilderModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use BuilderModuleMetaData.ABI instead.
var BuilderModuleABI = BuilderModuleMetaData.ABI

// BuilderModule is an auto generated Go binding around an Ethereum contract.
type BuilderModule struct {
	BuilderModuleCaller     // Read-only binding to the contract
	BuilderModuleTransactor // Write-only binding to the contract
	BuilderModuleFilterer   // Log filterer for contract events
}

// BuilderModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type BuilderModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuilderModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BuilderModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuilderModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BuilderModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuilderModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BuilderModuleSession struct {
	Contract     *BuilderModule    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BuilderModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BuilderModuleCallerSession struct {
	Contract *BuilderModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BuilderModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BuilderModuleTransactorSession struct {
	Contract     *BuilderModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BuilderModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type BuilderModuleRaw struct {
	Contract *BuilderModule // Generic contract binding to access the raw methods on
}

// BuilderModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BuilderModuleCallerRaw struct {
	Contract *BuilderModuleCaller // Generic read-only contract binding to access the raw methods on
}

// BuilderModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BuilderModuleTransactorRaw struct {
	Contract *BuilderModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBuilderModule creates a new instance of BuilderModule, bound to a specific deployed contract.
func NewBuilderModule(address common.Address, backend bind.ContractBackend) (*BuilderModule, error) {
	contract, err := bindBuilderModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BuilderModule{BuilderModuleCaller: BuilderModuleCaller{contract: contract}, BuilderModuleTransactor: BuilderModuleTransactor{contract: contract}, BuilderModuleFilterer: BuilderModuleFilterer{contract: contract}}, nil
}

// NewBuilderModuleCaller creates a new read-only instance of BuilderModule, bound to a specific deployed contract.
func NewBuilderModuleCaller(address common.Address, caller bind.ContractCaller) (*BuilderModuleCaller, error) {
	contract, err := bindBuilderModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BuilderModuleCaller{contract: contract}, nil
}

// NewBuilderModuleTransactor creates a new write-only instance of BuilderModule, bound to a specific deployed contract.
func NewBuilderModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*BuilderModuleTransactor, error) {
	contract, err := bindBuilderModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BuilderModuleTransactor{contract: contract}, nil
}

// NewBuilderModuleFilterer creates a new log filterer instance of BuilderModule, bound to a specific deployed contract.
func NewBuilderModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*BuilderModuleFilterer, error) {
	contract, err := bindBuilderModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BuilderModuleFilterer{contract: contract}, nil
}

// bindBuilderModule binds a generic wrapper to an already deployed contract.
func bindBuilderModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BuilderModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BuilderModule *BuilderModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BuilderModule.Contract.BuilderModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BuilderModule *BuilderModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuilderModule.Contract.BuilderModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BuilderModule *BuilderModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BuilderModule.Contract.BuilderModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BuilderModule *BuilderModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BuilderModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BuilderModule *BuilderModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BuilderModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BuilderModule *BuilderModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BuilderModule.Contract.contract.Transact(opts, method, params...)
}

// AuctionBid is a paid mutator transaction binding the contract method 0x98c4be1e.
//
// Solidity: function auctionBid(uint256 bid, bytes[] transactions, uint64 timeout) payable returns(bool)
func (_BuilderModule *BuilderModuleTransactor) AuctionBid(opts *bind.TransactOpts, bid *big.Int, transactions [][]byte, timeout uint64) (*types.Transaction, error) {
	return _BuilderModule.contract.Transact(opts, "auctionBid", bid, transactions, timeout)
}

// AuctionBid is a paid mutator transaction binding the contract method 0x98c4be1e.
//
// Solidity: function auctionBid(uint256 bid, bytes[] transactions, uint64 timeout) payable returns(bool)
func (_BuilderModule *BuilderModuleSession) AuctionBid(bid *big.Int, transactions [][]byte, timeout uint64) (*types.Transaction, error) {
	return _BuilderModule.Contract.AuctionBid(&_BuilderModule.TransactOpts, bid, transactions, timeout)
}

// AuctionBid is a paid mutator transaction binding the contract method 0x98c4be1e.
//
// Solidity: function auctionBid(uint256 bid, bytes[] transactions, uint64 timeout) payable returns(bool)
func (_BuilderModule *BuilderModuleTransactorSession) AuctionBid(bid *big.Int, transactions [][]byte, timeout uint64) (*types.Transaction, error) {
	return _BuilderModule.Contract.AuctionBid(&_BuilderModule.TransactOpts, bid, transactions, timeout)
}
