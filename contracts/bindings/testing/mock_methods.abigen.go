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

// MockMethodsmockStruct is an auto generated low-level Go binding around an user-defined struct.
type MockMethodsmockStruct struct {
	C *big.Int
}

// MockMethodsMetaData contains all meta data concerning the MockMethods contract.
var MockMethodsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"exampleFunc\",\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"b\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"c\",\"type\":\"tuple[]\",\"internalType\":\"structMockMethods.mockStruct[]\",\"components\":[{\"name\":\"c\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"zeroReturn\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// MockMethodsABI is the input ABI used to generate the binding from.
// Deprecated: Use MockMethodsMetaData.ABI instead.
var MockMethodsABI = MockMethodsMetaData.ABI

// MockMethods is an auto generated Go binding around an Ethereum contract.
type MockMethods struct {
	MockMethodsCaller     // Read-only binding to the contract
	MockMethodsTransactor // Write-only binding to the contract
	MockMethodsFilterer   // Log filterer for contract events
}

// MockMethodsCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockMethodsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockMethodsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockMethodsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockMethodsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockMethodsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockMethodsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockMethodsSession struct {
	Contract     *MockMethods      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockMethodsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockMethodsCallerSession struct {
	Contract *MockMethodsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MockMethodsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockMethodsTransactorSession struct {
	Contract     *MockMethodsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MockMethodsRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockMethodsRaw struct {
	Contract *MockMethods // Generic contract binding to access the raw methods on
}

// MockMethodsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockMethodsCallerRaw struct {
	Contract *MockMethodsCaller // Generic read-only contract binding to access the raw methods on
}

// MockMethodsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockMethodsTransactorRaw struct {
	Contract *MockMethodsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockMethods creates a new instance of MockMethods, bound to a specific deployed contract.
func NewMockMethods(address common.Address, backend bind.ContractBackend) (*MockMethods, error) {
	contract, err := bindMockMethods(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockMethods{MockMethodsCaller: MockMethodsCaller{contract: contract}, MockMethodsTransactor: MockMethodsTransactor{contract: contract}, MockMethodsFilterer: MockMethodsFilterer{contract: contract}}, nil
}

// NewMockMethodsCaller creates a new read-only instance of MockMethods, bound to a specific deployed contract.
func NewMockMethodsCaller(address common.Address, caller bind.ContractCaller) (*MockMethodsCaller, error) {
	contract, err := bindMockMethods(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockMethodsCaller{contract: contract}, nil
}

// NewMockMethodsTransactor creates a new write-only instance of MockMethods, bound to a specific deployed contract.
func NewMockMethodsTransactor(address common.Address, transactor bind.ContractTransactor) (*MockMethodsTransactor, error) {
	contract, err := bindMockMethods(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockMethodsTransactor{contract: contract}, nil
}

// NewMockMethodsFilterer creates a new log filterer instance of MockMethods, bound to a specific deployed contract.
func NewMockMethodsFilterer(address common.Address, filterer bind.ContractFilterer) (*MockMethodsFilterer, error) {
	contract, err := bindMockMethods(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockMethodsFilterer{contract: contract}, nil
}

// bindMockMethods binds a generic wrapper to an already deployed contract.
func bindMockMethods(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockMethodsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockMethods *MockMethodsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockMethods.Contract.MockMethodsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockMethods *MockMethodsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockMethods.Contract.MockMethodsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockMethods *MockMethodsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockMethods.Contract.MockMethodsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockMethods *MockMethodsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockMethods.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockMethods *MockMethodsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockMethods.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockMethods *MockMethodsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockMethods.Contract.contract.Transact(opts, method, params...)
}

// ExampleFunc is a paid mutator transaction binding the contract method 0x6eae29f5.
//
// Solidity: function exampleFunc(uint256 a, address b, (uint256)[] c) returns(bool)
func (_MockMethods *MockMethodsTransactor) ExampleFunc(opts *bind.TransactOpts, a *big.Int, b common.Address, c []MockMethodsmockStruct) (*types.Transaction, error) {
	return _MockMethods.contract.Transact(opts, "exampleFunc", a, b, c)
}

// ExampleFunc is a paid mutator transaction binding the contract method 0x6eae29f5.
//
// Solidity: function exampleFunc(uint256 a, address b, (uint256)[] c) returns(bool)
func (_MockMethods *MockMethodsSession) ExampleFunc(a *big.Int, b common.Address, c []MockMethodsmockStruct) (*types.Transaction, error) {
	return _MockMethods.Contract.ExampleFunc(&_MockMethods.TransactOpts, a, b, c)
}

// ExampleFunc is a paid mutator transaction binding the contract method 0x6eae29f5.
//
// Solidity: function exampleFunc(uint256 a, address b, (uint256)[] c) returns(bool)
func (_MockMethods *MockMethodsTransactorSession) ExampleFunc(a *big.Int, b common.Address, c []MockMethodsmockStruct) (*types.Transaction, error) {
	return _MockMethods.Contract.ExampleFunc(&_MockMethods.TransactOpts, a, b, c)
}

// ZeroReturn is a paid mutator transaction binding the contract method 0x307bcbcc.
//
// Solidity: function zeroReturn() returns()
func (_MockMethods *MockMethodsTransactor) ZeroReturn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockMethods.contract.Transact(opts, "zeroReturn")
}

// ZeroReturn is a paid mutator transaction binding the contract method 0x307bcbcc.
//
// Solidity: function zeroReturn() returns()
func (_MockMethods *MockMethodsSession) ZeroReturn() (*types.Transaction, error) {
	return _MockMethods.Contract.ZeroReturn(&_MockMethods.TransactOpts)
}

// ZeroReturn is a paid mutator transaction binding the contract method 0x307bcbcc.
//
// Solidity: function zeroReturn() returns()
func (_MockMethods *MockMethodsTransactorSession) ZeroReturn() (*types.Transaction, error) {
	return _MockMethods.Contract.ZeroReturn(&_MockMethods.TransactOpts)
}
