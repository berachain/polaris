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
	_ = abi.ConvertType
)

// RevertableTxMetaData contains all meta data concerning the RevertableTx contract.
var RevertableTxMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052348015600f57600080fd5b5060bc8061001e6000396000f3fe60806040523660445760405162461bcd60e51b815260206004820152600c60248201526b0a4caeccae4e8c2c4d8caa8f60a31b60448201526064015b60405180910390fd5b348015604f57600080fd5b5060405162461bcd60e51b815260206004820152600c60248201526b0a4caeccae4e8c2c4d8caa8f60a31b6044820152606401603b56fea2646970667358221220ddc4c09355bddb486aadad0f12baa38abd7d589ed8f9d490a32a92adf2db088964736f6c63430008110033",
}

// RevertableTxABI is the input ABI used to generate the binding from.
// Deprecated: Use RevertableTxMetaData.ABI instead.
var RevertableTxABI = RevertableTxMetaData.ABI

// RevertableTxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RevertableTxMetaData.Bin instead.
var RevertableTxBin = RevertableTxMetaData.Bin

// DeployRevertableTx deploys a new Ethereum contract, binding an instance of RevertableTx to it.
func DeployRevertableTx(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RevertableTx, error) {
	parsed, err := RevertableTxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RevertableTxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RevertableTx{RevertableTxCaller: RevertableTxCaller{contract: contract}, RevertableTxTransactor: RevertableTxTransactor{contract: contract}, RevertableTxFilterer: RevertableTxFilterer{contract: contract}}, nil
}

// RevertableTx is an auto generated Go binding around an Ethereum contract.
type RevertableTx struct {
	RevertableTxCaller     // Read-only binding to the contract
	RevertableTxTransactor // Write-only binding to the contract
	RevertableTxFilterer   // Log filterer for contract events
}

// RevertableTxCaller is an auto generated read-only Go binding around an Ethereum contract.
type RevertableTxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevertableTxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RevertableTxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevertableTxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RevertableTxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevertableTxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RevertableTxSession struct {
	Contract     *RevertableTx     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RevertableTxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RevertableTxCallerSession struct {
	Contract *RevertableTxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RevertableTxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RevertableTxTransactorSession struct {
	Contract     *RevertableTxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RevertableTxRaw is an auto generated low-level Go binding around an Ethereum contract.
type RevertableTxRaw struct {
	Contract *RevertableTx // Generic contract binding to access the raw methods on
}

// RevertableTxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RevertableTxCallerRaw struct {
	Contract *RevertableTxCaller // Generic read-only contract binding to access the raw methods on
}

// RevertableTxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RevertableTxTransactorRaw struct {
	Contract *RevertableTxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRevertableTx creates a new instance of RevertableTx, bound to a specific deployed contract.
func NewRevertableTx(address common.Address, backend bind.ContractBackend) (*RevertableTx, error) {
	contract, err := bindRevertableTx(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RevertableTx{RevertableTxCaller: RevertableTxCaller{contract: contract}, RevertableTxTransactor: RevertableTxTransactor{contract: contract}, RevertableTxFilterer: RevertableTxFilterer{contract: contract}}, nil
}

// NewRevertableTxCaller creates a new read-only instance of RevertableTx, bound to a specific deployed contract.
func NewRevertableTxCaller(address common.Address, caller bind.ContractCaller) (*RevertableTxCaller, error) {
	contract, err := bindRevertableTx(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RevertableTxCaller{contract: contract}, nil
}

// NewRevertableTxTransactor creates a new write-only instance of RevertableTx, bound to a specific deployed contract.
func NewRevertableTxTransactor(address common.Address, transactor bind.ContractTransactor) (*RevertableTxTransactor, error) {
	contract, err := bindRevertableTx(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RevertableTxTransactor{contract: contract}, nil
}

// NewRevertableTxFilterer creates a new log filterer instance of RevertableTx, bound to a specific deployed contract.
func NewRevertableTxFilterer(address common.Address, filterer bind.ContractFilterer) (*RevertableTxFilterer, error) {
	contract, err := bindRevertableTx(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RevertableTxFilterer{contract: contract}, nil
}

// bindRevertableTx binds a generic wrapper to an already deployed contract.
func bindRevertableTx(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RevertableTxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RevertableTx *RevertableTxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RevertableTx.Contract.RevertableTxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RevertableTx *RevertableTxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevertableTx.Contract.RevertableTxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RevertableTx *RevertableTxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RevertableTx.Contract.RevertableTxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RevertableTx *RevertableTxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RevertableTx.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RevertableTx *RevertableTxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevertableTx.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RevertableTx *RevertableTxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RevertableTx.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_RevertableTx *RevertableTxTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _RevertableTx.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_RevertableTx *RevertableTxSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RevertableTx.Contract.Fallback(&_RevertableTx.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_RevertableTx *RevertableTxTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _RevertableTx.Contract.Fallback(&_RevertableTx.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RevertableTx *RevertableTxTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevertableTx.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RevertableTx *RevertableTxSession) Receive() (*types.Transaction, error) {
	return _RevertableTx.Contract.Receive(&_RevertableTx.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RevertableTx *RevertableTxTransactorSession) Receive() (*types.Transaction, error) {
	return _RevertableTx.Contract.Receive(&_RevertableTx.TransactOpts)
}
