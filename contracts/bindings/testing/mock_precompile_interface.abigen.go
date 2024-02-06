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

// MockPrecompileInterfaceObject is an auto generated low-level Go binding around an user-defined struct.
type MockPrecompileInterfaceObject struct {
	CreationHeight *big.Int
	TimeStamp      string
}

// MockPrecompileMetaData contains all meta data concerning the MockPrecompile contract.
var MockPrecompileMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"contractFunc\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"ans\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"contractFuncStr\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOutput\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structMockPrecompileInterface.Object[]\",\"components\":[{\"name\":\"creationHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeStamp\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOutputPartial\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structMockPrecompileInterface.Object\",\"components\":[{\"name\":\"creationHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeStamp\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"overloadedFunc\",\"inputs\":[],\"outputs\":[{\"name\":\"ans\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"overloadedFunc\",\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"ans\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"}]",
}

// MockPrecompileABI is the input ABI used to generate the binding from.
// Deprecated: Use MockPrecompileMetaData.ABI instead.
var MockPrecompileABI = MockPrecompileMetaData.ABI

// MockPrecompile is an auto generated Go binding around an Ethereum contract.
type MockPrecompile struct {
	MockPrecompileCaller     // Read-only binding to the contract
	MockPrecompileTransactor // Write-only binding to the contract
	MockPrecompileFilterer   // Log filterer for contract events
}

// MockPrecompileCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockPrecompileCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockPrecompileTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockPrecompileTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockPrecompileFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockPrecompileFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockPrecompileSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockPrecompileSession struct {
	Contract     *MockPrecompile   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockPrecompileCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockPrecompileCallerSession struct {
	Contract *MockPrecompileCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MockPrecompileTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockPrecompileTransactorSession struct {
	Contract     *MockPrecompileTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MockPrecompileRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockPrecompileRaw struct {
	Contract *MockPrecompile // Generic contract binding to access the raw methods on
}

// MockPrecompileCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockPrecompileCallerRaw struct {
	Contract *MockPrecompileCaller // Generic read-only contract binding to access the raw methods on
}

// MockPrecompileTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockPrecompileTransactorRaw struct {
	Contract *MockPrecompileTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockPrecompile creates a new instance of MockPrecompile, bound to a specific deployed contract.
func NewMockPrecompile(address common.Address, backend bind.ContractBackend) (*MockPrecompile, error) {
	contract, err := bindMockPrecompile(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockPrecompile{MockPrecompileCaller: MockPrecompileCaller{contract: contract}, MockPrecompileTransactor: MockPrecompileTransactor{contract: contract}, MockPrecompileFilterer: MockPrecompileFilterer{contract: contract}}, nil
}

// NewMockPrecompileCaller creates a new read-only instance of MockPrecompile, bound to a specific deployed contract.
func NewMockPrecompileCaller(address common.Address, caller bind.ContractCaller) (*MockPrecompileCaller, error) {
	contract, err := bindMockPrecompile(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockPrecompileCaller{contract: contract}, nil
}

// NewMockPrecompileTransactor creates a new write-only instance of MockPrecompile, bound to a specific deployed contract.
func NewMockPrecompileTransactor(address common.Address, transactor bind.ContractTransactor) (*MockPrecompileTransactor, error) {
	contract, err := bindMockPrecompile(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockPrecompileTransactor{contract: contract}, nil
}

// NewMockPrecompileFilterer creates a new log filterer instance of MockPrecompile, bound to a specific deployed contract.
func NewMockPrecompileFilterer(address common.Address, filterer bind.ContractFilterer) (*MockPrecompileFilterer, error) {
	contract, err := bindMockPrecompile(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockPrecompileFilterer{contract: contract}, nil
}

// bindMockPrecompile binds a generic wrapper to an already deployed contract.
func bindMockPrecompile(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockPrecompileMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockPrecompile *MockPrecompileRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockPrecompile.Contract.MockPrecompileCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockPrecompile *MockPrecompileRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockPrecompile.Contract.MockPrecompileTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockPrecompile *MockPrecompileRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockPrecompile.Contract.MockPrecompileTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockPrecompile *MockPrecompileCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockPrecompile.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockPrecompile *MockPrecompileTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockPrecompile.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockPrecompile *MockPrecompileTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockPrecompile.Contract.contract.Transact(opts, method, params...)
}

// ContractFunc is a paid mutator transaction binding the contract method 0xc7dda0b9.
//
// Solidity: function contractFunc(address addr) returns(uint256 ans)
func (_MockPrecompile *MockPrecompileTransactor) ContractFunc(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _MockPrecompile.contract.Transact(opts, "contractFunc", addr)
}

// ContractFunc is a paid mutator transaction binding the contract method 0xc7dda0b9.
//
// Solidity: function contractFunc(address addr) returns(uint256 ans)
func (_MockPrecompile *MockPrecompileSession) ContractFunc(addr common.Address) (*types.Transaction, error) {
	return _MockPrecompile.Contract.ContractFunc(&_MockPrecompile.TransactOpts, addr)
}

// ContractFunc is a paid mutator transaction binding the contract method 0xc7dda0b9.
//
// Solidity: function contractFunc(address addr) returns(uint256 ans)
func (_MockPrecompile *MockPrecompileTransactorSession) ContractFunc(addr common.Address) (*types.Transaction, error) {
	return _MockPrecompile.Contract.ContractFunc(&_MockPrecompile.TransactOpts, addr)
}

// ContractFuncStr is a paid mutator transaction binding the contract method 0x04bb5393.
//
// Solidity: function contractFuncStr(string str) returns(bool)
func (_MockPrecompile *MockPrecompileTransactor) ContractFuncStr(opts *bind.TransactOpts, str string) (*types.Transaction, error) {
	return _MockPrecompile.contract.Transact(opts, "contractFuncStr", str)
}

// ContractFuncStr is a paid mutator transaction binding the contract method 0x04bb5393.
//
// Solidity: function contractFuncStr(string str) returns(bool)
func (_MockPrecompile *MockPrecompileSession) ContractFuncStr(str string) (*types.Transaction, error) {
	return _MockPrecompile.Contract.ContractFuncStr(&_MockPrecompile.TransactOpts, str)
}

// ContractFuncStr is a paid mutator transaction binding the contract method 0x04bb5393.
//
// Solidity: function contractFuncStr(string str) returns(bool)
func (_MockPrecompile *MockPrecompileTransactorSession) ContractFuncStr(str string) (*types.Transaction, error) {
	return _MockPrecompile.Contract.ContractFuncStr(&_MockPrecompile.TransactOpts, str)
}

// GetOutput is a paid mutator transaction binding the contract method 0xb5c11fc2.
//
// Solidity: function getOutput(string str) returns((uint256,string)[])
func (_MockPrecompile *MockPrecompileTransactor) GetOutput(opts *bind.TransactOpts, str string) (*types.Transaction, error) {
	return _MockPrecompile.contract.Transact(opts, "getOutput", str)
}

// GetOutput is a paid mutator transaction binding the contract method 0xb5c11fc2.
//
// Solidity: function getOutput(string str) returns((uint256,string)[])
func (_MockPrecompile *MockPrecompileSession) GetOutput(str string) (*types.Transaction, error) {
	return _MockPrecompile.Contract.GetOutput(&_MockPrecompile.TransactOpts, str)
}

// GetOutput is a paid mutator transaction binding the contract method 0xb5c11fc2.
//
// Solidity: function getOutput(string str) returns((uint256,string)[])
func (_MockPrecompile *MockPrecompileTransactorSession) GetOutput(str string) (*types.Transaction, error) {
	return _MockPrecompile.Contract.GetOutput(&_MockPrecompile.TransactOpts, str)
}

// GetOutputPartial is a paid mutator transaction binding the contract method 0x7acaaeb9.
//
// Solidity: function getOutputPartial() returns((uint256,string))
func (_MockPrecompile *MockPrecompileTransactor) GetOutputPartial(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockPrecompile.contract.Transact(opts, "getOutputPartial")
}

// GetOutputPartial is a paid mutator transaction binding the contract method 0x7acaaeb9.
//
// Solidity: function getOutputPartial() returns((uint256,string))
func (_MockPrecompile *MockPrecompileSession) GetOutputPartial() (*types.Transaction, error) {
	return _MockPrecompile.Contract.GetOutputPartial(&_MockPrecompile.TransactOpts)
}

// GetOutputPartial is a paid mutator transaction binding the contract method 0x7acaaeb9.
//
// Solidity: function getOutputPartial() returns((uint256,string))
func (_MockPrecompile *MockPrecompileTransactorSession) GetOutputPartial() (*types.Transaction, error) {
	return _MockPrecompile.Contract.GetOutputPartial(&_MockPrecompile.TransactOpts)
}

// OverloadedFunc is a paid mutator transaction binding the contract method 0x1e61d5aa.
//
// Solidity: function overloadedFunc() returns(uint256 ans)
func (_MockPrecompile *MockPrecompileTransactor) OverloadedFunc(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockPrecompile.contract.Transact(opts, "overloadedFunc")
}

// OverloadedFunc is a paid mutator transaction binding the contract method 0x1e61d5aa.
//
// Solidity: function overloadedFunc() returns(uint256 ans)
func (_MockPrecompile *MockPrecompileSession) OverloadedFunc() (*types.Transaction, error) {
	return _MockPrecompile.Contract.OverloadedFunc(&_MockPrecompile.TransactOpts)
}

// OverloadedFunc is a paid mutator transaction binding the contract method 0x1e61d5aa.
//
// Solidity: function overloadedFunc() returns(uint256 ans)
func (_MockPrecompile *MockPrecompileTransactorSession) OverloadedFunc() (*types.Transaction, error) {
	return _MockPrecompile.Contract.OverloadedFunc(&_MockPrecompile.TransactOpts)
}

// OverloadedFunc0 is a paid mutator transaction binding the contract method 0x5482a42b.
//
// Solidity: function overloadedFunc(uint256 a) returns(uint256 ans)
func (_MockPrecompile *MockPrecompileTransactor) OverloadedFunc0(opts *bind.TransactOpts, a *big.Int) (*types.Transaction, error) {
	return _MockPrecompile.contract.Transact(opts, "overloadedFunc0", a)
}

// OverloadedFunc0 is a paid mutator transaction binding the contract method 0x5482a42b.
//
// Solidity: function overloadedFunc(uint256 a) returns(uint256 ans)
func (_MockPrecompile *MockPrecompileSession) OverloadedFunc0(a *big.Int) (*types.Transaction, error) {
	return _MockPrecompile.Contract.OverloadedFunc0(&_MockPrecompile.TransactOpts, a)
}

// OverloadedFunc0 is a paid mutator transaction binding the contract method 0x5482a42b.
//
// Solidity: function overloadedFunc(uint256 a) returns(uint256 ans)
func (_MockPrecompile *MockPrecompileTransactorSession) OverloadedFunc0(a *big.Int) (*types.Transaction, error) {
	return _MockPrecompile.Contract.OverloadedFunc0(&_MockPrecompile.TransactOpts, a)
}
