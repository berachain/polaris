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

// IERC20ModuleCoin is an auto generated low-level Go binding around an user-defined struct.
type IERC20ModuleCoin struct {
	Amount *big.Int
	Denom  string
}

// ERC20ModuleMetaData contains all meta data concerning the ERC20Module contract.
var ERC20ModuleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIERC20Module.Coin[]\",\"name\":\"amount\",\"type\":\"tuple[]\"}],\"name\":\"TransferCoinToErc20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIERC20Module.Coin[]\",\"name\":\"amount\",\"type\":\"tuple[]\"}],\"name\":\"TransferErc20ToCoin\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"coinDenomForERC20Address\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"token\",\"type\":\"string\"}],\"name\":\"coinDenomForERC20Address\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"erc20AddressForCoinDenom\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferCoinToERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferCoinToERC20From\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"owner\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"recipient\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferCoinToERC20From\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"recipient\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferCoinToERC20To\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferCoinToERC20To\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferERC20ToCoin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"owner\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"recipient\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferERC20ToCoinFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferERC20ToCoinFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferERC20ToCoinTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"recipient\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferERC20ToCoinTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC20ModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20ModuleMetaData.ABI instead.
var ERC20ModuleABI = ERC20ModuleMetaData.ABI

// ERC20Module is an auto generated Go binding around an Ethereum contract.
type ERC20Module struct {
	ERC20ModuleCaller     // Read-only binding to the contract
	ERC20ModuleTransactor // Write-only binding to the contract
	ERC20ModuleFilterer   // Log filterer for contract events
}

// ERC20ModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20ModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20ModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20ModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20ModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20ModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20ModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20ModuleSession struct {
	Contract     *ERC20Module      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20ModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20ModuleCallerSession struct {
	Contract *ERC20ModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ERC20ModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20ModuleTransactorSession struct {
	Contract     *ERC20ModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ERC20ModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20ModuleRaw struct {
	Contract *ERC20Module // Generic contract binding to access the raw methods on
}

// ERC20ModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20ModuleCallerRaw struct {
	Contract *ERC20ModuleCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20ModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20ModuleTransactorRaw struct {
	Contract *ERC20ModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Module creates a new instance of ERC20Module, bound to a specific deployed contract.
func NewERC20Module(address common.Address, backend bind.ContractBackend) (*ERC20Module, error) {
	contract, err := bindERC20Module(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Module{ERC20ModuleCaller: ERC20ModuleCaller{contract: contract}, ERC20ModuleTransactor: ERC20ModuleTransactor{contract: contract}, ERC20ModuleFilterer: ERC20ModuleFilterer{contract: contract}}, nil
}

// NewERC20ModuleCaller creates a new read-only instance of ERC20Module, bound to a specific deployed contract.
func NewERC20ModuleCaller(address common.Address, caller bind.ContractCaller) (*ERC20ModuleCaller, error) {
	contract, err := bindERC20Module(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20ModuleCaller{contract: contract}, nil
}

// NewERC20ModuleTransactor creates a new write-only instance of ERC20Module, bound to a specific deployed contract.
func NewERC20ModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20ModuleTransactor, error) {
	contract, err := bindERC20Module(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20ModuleTransactor{contract: contract}, nil
}

// NewERC20ModuleFilterer creates a new log filterer instance of ERC20Module, bound to a specific deployed contract.
func NewERC20ModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20ModuleFilterer, error) {
	contract, err := bindERC20Module(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20ModuleFilterer{contract: contract}, nil
}

// bindERC20Module binds a generic wrapper to an already deployed contract.
func bindERC20Module(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20ModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Module *ERC20ModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Module.Contract.ERC20ModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Module *ERC20ModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Module.Contract.ERC20ModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Module *ERC20ModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Module.Contract.ERC20ModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Module *ERC20ModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Module.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Module *ERC20ModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Module.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Module *ERC20ModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Module.Contract.contract.Transact(opts, method, params...)
}

// CoinDenomForERC20Address is a free data retrieval call binding the contract method 0xcd22a018.
//
// Solidity: function coinDenomForERC20Address(address token) view returns(string)
func (_ERC20Module *ERC20ModuleCaller) CoinDenomForERC20Address(opts *bind.CallOpts, token common.Address) (string, error) {
	var out []interface{}
	err := _ERC20Module.contract.Call(opts, &out, "coinDenomForERC20Address", token)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CoinDenomForERC20Address is a free data retrieval call binding the contract method 0xcd22a018.
//
// Solidity: function coinDenomForERC20Address(address token) view returns(string)
func (_ERC20Module *ERC20ModuleSession) CoinDenomForERC20Address(token common.Address) (string, error) {
	return _ERC20Module.Contract.CoinDenomForERC20Address(&_ERC20Module.CallOpts, token)
}

// CoinDenomForERC20Address is a free data retrieval call binding the contract method 0xcd22a018.
//
// Solidity: function coinDenomForERC20Address(address token) view returns(string)
func (_ERC20Module *ERC20ModuleCallerSession) CoinDenomForERC20Address(token common.Address) (string, error) {
	return _ERC20Module.Contract.CoinDenomForERC20Address(&_ERC20Module.CallOpts, token)
}

// CoinDenomForERC20Address0 is a free data retrieval call binding the contract method 0xe2bea1fe.
//
// Solidity: function coinDenomForERC20Address(string token) view returns(string)
func (_ERC20Module *ERC20ModuleCaller) CoinDenomForERC20Address0(opts *bind.CallOpts, token string) (string, error) {
	var out []interface{}
	err := _ERC20Module.contract.Call(opts, &out, "coinDenomForERC20Address0", token)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CoinDenomForERC20Address0 is a free data retrieval call binding the contract method 0xe2bea1fe.
//
// Solidity: function coinDenomForERC20Address(string token) view returns(string)
func (_ERC20Module *ERC20ModuleSession) CoinDenomForERC20Address0(token string) (string, error) {
	return _ERC20Module.Contract.CoinDenomForERC20Address0(&_ERC20Module.CallOpts, token)
}

// CoinDenomForERC20Address0 is a free data retrieval call binding the contract method 0xe2bea1fe.
//
// Solidity: function coinDenomForERC20Address(string token) view returns(string)
func (_ERC20Module *ERC20ModuleCallerSession) CoinDenomForERC20Address0(token string) (string, error) {
	return _ERC20Module.Contract.CoinDenomForERC20Address0(&_ERC20Module.CallOpts, token)
}

// Erc20AddressForCoinDenom is a free data retrieval call binding the contract method 0xa333e57c.
//
// Solidity: function erc20AddressForCoinDenom(string denom) view returns(address)
func (_ERC20Module *ERC20ModuleCaller) Erc20AddressForCoinDenom(opts *bind.CallOpts, denom string) (common.Address, error) {
	var out []interface{}
	err := _ERC20Module.contract.Call(opts, &out, "erc20AddressForCoinDenom", denom)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20AddressForCoinDenom is a free data retrieval call binding the contract method 0xa333e57c.
//
// Solidity: function erc20AddressForCoinDenom(string denom) view returns(address)
func (_ERC20Module *ERC20ModuleSession) Erc20AddressForCoinDenom(denom string) (common.Address, error) {
	return _ERC20Module.Contract.Erc20AddressForCoinDenom(&_ERC20Module.CallOpts, denom)
}

// Erc20AddressForCoinDenom is a free data retrieval call binding the contract method 0xa333e57c.
//
// Solidity: function erc20AddressForCoinDenom(string denom) view returns(address)
func (_ERC20Module *ERC20ModuleCallerSession) Erc20AddressForCoinDenom(denom string) (common.Address, error) {
	return _ERC20Module.Contract.Erc20AddressForCoinDenom(&_ERC20Module.CallOpts, denom)
}

// TransferCoinToERC20 is a paid mutator transaction binding the contract method 0x416daf89.
//
// Solidity: function transferCoinToERC20(string denom, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactor) TransferCoinToERC20(opts *bind.TransactOpts, denom string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.contract.Transact(opts, "transferCoinToERC20", denom, amount)
}

// TransferCoinToERC20 is a paid mutator transaction binding the contract method 0x416daf89.
//
// Solidity: function transferCoinToERC20(string denom, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleSession) TransferCoinToERC20(denom string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.TransferCoinToERC20(&_ERC20Module.TransactOpts, denom, amount)
}

// TransferCoinToERC20 is a paid mutator transaction binding the contract method 0x416daf89.
//
// Solidity: function transferCoinToERC20(string denom, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactorSession) TransferCoinToERC20(denom string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.TransferCoinToERC20(&_ERC20Module.TransactOpts, denom, amount)
}

// TransferCoinToERC20From is a paid mutator transaction binding the contract method 0x096b4069.
//
// Solidity: function transferCoinToERC20From(string denom, address owner, address recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactor) TransferCoinToERC20From(opts *bind.TransactOpts, denom string, owner common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.contract.Transact(opts, "transferCoinToERC20From", denom, owner, recipient, amount)
}

// TransferCoinToERC20From is a paid mutator transaction binding the contract method 0x096b4069.
//
// Solidity: function transferCoinToERC20From(string denom, address owner, address recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleSession) TransferCoinToERC20From(denom string, owner common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.TransferCoinToERC20From(&_ERC20Module.TransactOpts, denom, owner, recipient, amount)
}

// TransferCoinToERC20From is a paid mutator transaction binding the contract method 0x096b4069.
//
// Solidity: function transferCoinToERC20From(string denom, address owner, address recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactorSession) TransferCoinToERC20From(denom string, owner common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.TransferCoinToERC20From(&_ERC20Module.TransactOpts, denom, owner, recipient, amount)
}

// TransferCoinToERC20From0 is a paid mutator transaction binding the contract method 0x83228263.
//
// Solidity: function transferCoinToERC20From(string denom, string owner, string recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactor) TransferCoinToERC20From0(opts *bind.TransactOpts, denom string, owner string, recipient string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.contract.Transact(opts, "transferCoinToERC20From0", denom, owner, recipient, amount)
}

// TransferCoinToERC20From0 is a paid mutator transaction binding the contract method 0x83228263.
//
// Solidity: function transferCoinToERC20From(string denom, string owner, string recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleSession) TransferCoinToERC20From0(denom string, owner string, recipient string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.TransferCoinToERC20From0(&_ERC20Module.TransactOpts, denom, owner, recipient, amount)
}

// TransferCoinToERC20From0 is a paid mutator transaction binding the contract method 0x83228263.
//
// Solidity: function transferCoinToERC20From(string denom, string owner, string recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactorSession) TransferCoinToERC20From0(denom string, owner string, recipient string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.TransferCoinToERC20From0(&_ERC20Module.TransactOpts, denom, owner, recipient, amount)
}

// TransferCoinToERC20To is a paid mutator transaction binding the contract method 0x404ffd2f.
//
// Solidity: function transferCoinToERC20To(string denom, string recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactor) TransferCoinToERC20To(opts *bind.TransactOpts, denom string, recipient string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.contract.Transact(opts, "transferCoinToERC20To", denom, recipient, amount)
}

// TransferCoinToERC20To is a paid mutator transaction binding the contract method 0x404ffd2f.
//
// Solidity: function transferCoinToERC20To(string denom, string recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleSession) TransferCoinToERC20To(denom string, recipient string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.TransferCoinToERC20To(&_ERC20Module.TransactOpts, denom, recipient, amount)
}

// TransferCoinToERC20To is a paid mutator transaction binding the contract method 0x404ffd2f.
//
// Solidity: function transferCoinToERC20To(string denom, string recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactorSession) TransferCoinToERC20To(denom string, recipient string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.TransferCoinToERC20To(&_ERC20Module.TransactOpts, denom, recipient, amount)
}

// TransferCoinToERC20To0 is a paid mutator transaction binding the contract method 0xe33b064d.
//
// Solidity: function transferCoinToERC20To(string denom, address recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactor) TransferCoinToERC20To0(opts *bind.TransactOpts, denom string, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.contract.Transact(opts, "transferCoinToERC20To0", denom, recipient, amount)
}

// TransferCoinToERC20To0 is a paid mutator transaction binding the contract method 0xe33b064d.
//
// Solidity: function transferCoinToERC20To(string denom, address recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleSession) TransferCoinToERC20To0(denom string, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.TransferCoinToERC20To0(&_ERC20Module.TransactOpts, denom, recipient, amount)
}

// TransferCoinToERC20To0 is a paid mutator transaction binding the contract method 0xe33b064d.
//
// Solidity: function transferCoinToERC20To(string denom, address recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactorSession) TransferCoinToERC20To0(denom string, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.TransferCoinToERC20To0(&_ERC20Module.TransactOpts, denom, recipient, amount)
}

// TransferERC20ToCoin is a paid mutator transaction binding the contract method 0x4c2b7543.
//
// Solidity: function transferERC20ToCoin(address token, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactor) TransferERC20ToCoin(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.contract.Transact(opts, "transferERC20ToCoin", token, amount)
}

// TransferERC20ToCoin is a paid mutator transaction binding the contract method 0x4c2b7543.
//
// Solidity: function transferERC20ToCoin(address token, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleSession) TransferERC20ToCoin(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.TransferERC20ToCoin(&_ERC20Module.TransactOpts, token, amount)
}

// TransferERC20ToCoin is a paid mutator transaction binding the contract method 0x4c2b7543.
//
// Solidity: function transferERC20ToCoin(address token, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactorSession) TransferERC20ToCoin(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.TransferERC20ToCoin(&_ERC20Module.TransactOpts, token, amount)
}

// TransferERC20ToCoinFrom is a paid mutator transaction binding the contract method 0x02ecf60f.
//
// Solidity: function transferERC20ToCoinFrom(address token, string owner, string recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactor) TransferERC20ToCoinFrom(opts *bind.TransactOpts, token common.Address, owner string, recipient string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.contract.Transact(opts, "transferERC20ToCoinFrom", token, owner, recipient, amount)
}

// TransferERC20ToCoinFrom is a paid mutator transaction binding the contract method 0x02ecf60f.
//
// Solidity: function transferERC20ToCoinFrom(address token, string owner, string recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleSession) TransferERC20ToCoinFrom(token common.Address, owner string, recipient string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.TransferERC20ToCoinFrom(&_ERC20Module.TransactOpts, token, owner, recipient, amount)
}

// TransferERC20ToCoinFrom is a paid mutator transaction binding the contract method 0x02ecf60f.
//
// Solidity: function transferERC20ToCoinFrom(address token, string owner, string recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactorSession) TransferERC20ToCoinFrom(token common.Address, owner string, recipient string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.TransferERC20ToCoinFrom(&_ERC20Module.TransactOpts, token, owner, recipient, amount)
}

// TransferERC20ToCoinFrom0 is a paid mutator transaction binding the contract method 0xb96d8bec.
//
// Solidity: function transferERC20ToCoinFrom(address token, address owner, address recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactor) TransferERC20ToCoinFrom0(opts *bind.TransactOpts, token common.Address, owner common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.contract.Transact(opts, "transferERC20ToCoinFrom0", token, owner, recipient, amount)
}

// TransferERC20ToCoinFrom0 is a paid mutator transaction binding the contract method 0xb96d8bec.
//
// Solidity: function transferERC20ToCoinFrom(address token, address owner, address recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleSession) TransferERC20ToCoinFrom0(token common.Address, owner common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.TransferERC20ToCoinFrom0(&_ERC20Module.TransactOpts, token, owner, recipient, amount)
}

// TransferERC20ToCoinFrom0 is a paid mutator transaction binding the contract method 0xb96d8bec.
//
// Solidity: function transferERC20ToCoinFrom(address token, address owner, address recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactorSession) TransferERC20ToCoinFrom0(token common.Address, owner common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.TransferERC20ToCoinFrom0(&_ERC20Module.TransactOpts, token, owner, recipient, amount)
}

// TransferERC20ToCoinTo is a paid mutator transaction binding the contract method 0x226eb41b.
//
// Solidity: function transferERC20ToCoinTo(address token, address recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactor) TransferERC20ToCoinTo(opts *bind.TransactOpts, token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.contract.Transact(opts, "transferERC20ToCoinTo", token, recipient, amount)
}

// TransferERC20ToCoinTo is a paid mutator transaction binding the contract method 0x226eb41b.
//
// Solidity: function transferERC20ToCoinTo(address token, address recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleSession) TransferERC20ToCoinTo(token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.TransferERC20ToCoinTo(&_ERC20Module.TransactOpts, token, recipient, amount)
}

// TransferERC20ToCoinTo is a paid mutator transaction binding the contract method 0x226eb41b.
//
// Solidity: function transferERC20ToCoinTo(address token, address recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactorSession) TransferERC20ToCoinTo(token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.TransferERC20ToCoinTo(&_ERC20Module.TransactOpts, token, recipient, amount)
}

// TransferERC20ToCoinTo0 is a paid mutator transaction binding the contract method 0x6ba2a25a.
//
// Solidity: function transferERC20ToCoinTo(address token, string recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactor) TransferERC20ToCoinTo0(opts *bind.TransactOpts, token common.Address, recipient string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.contract.Transact(opts, "transferERC20ToCoinTo0", token, recipient, amount)
}

// TransferERC20ToCoinTo0 is a paid mutator transaction binding the contract method 0x6ba2a25a.
//
// Solidity: function transferERC20ToCoinTo(address token, string recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleSession) TransferERC20ToCoinTo0(token common.Address, recipient string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.TransferERC20ToCoinTo0(&_ERC20Module.TransactOpts, token, recipient, amount)
}

// TransferERC20ToCoinTo0 is a paid mutator transaction binding the contract method 0x6ba2a25a.
//
// Solidity: function transferERC20ToCoinTo(address token, string recipient, uint256 amount) returns(bool)
func (_ERC20Module *ERC20ModuleTransactorSession) TransferERC20ToCoinTo0(token common.Address, recipient string, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Module.Contract.TransferERC20ToCoinTo0(&_ERC20Module.TransactOpts, token, recipient, amount)
}

// ERC20ModuleTransferCoinToErc20Iterator is returned from FilterTransferCoinToErc20 and is used to iterate over the raw logs and unpacked data for TransferCoinToErc20 events raised by the ERC20Module contract.
type ERC20ModuleTransferCoinToErc20Iterator struct {
	Event *ERC20ModuleTransferCoinToErc20 // Event containing the contract specifics and raw log

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
func (it *ERC20ModuleTransferCoinToErc20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20ModuleTransferCoinToErc20)
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
		it.Event = new(ERC20ModuleTransferCoinToErc20)
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
func (it *ERC20ModuleTransferCoinToErc20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ModuleTransferCoinToErc20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20ModuleTransferCoinToErc20 represents a TransferCoinToErc20 event raised by the ERC20Module contract.
type ERC20ModuleTransferCoinToErc20 struct {
	Denom     common.Hash
	Owner     common.Address
	Recipient common.Address
	Amount    []IERC20ModuleCoin
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransferCoinToErc20 is a free log retrieval operation binding the contract event 0xaa6b0040cf4e38b34df3fd01e51cc906a049cfcf3d5a6b81a9ac19dedcf533f7.
//
// Solidity: event TransferCoinToErc20(string indexed denom, address indexed owner, address indexed recipient, (uint256,string)[] amount)
func (_ERC20Module *ERC20ModuleFilterer) FilterTransferCoinToErc20(opts *bind.FilterOpts, denom []string, owner []common.Address, recipient []common.Address) (*ERC20ModuleTransferCoinToErc20Iterator, error) {

	var denomRule []interface{}
	for _, denomItem := range denom {
		denomRule = append(denomRule, denomItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20Module.contract.FilterLogs(opts, "TransferCoinToErc20", denomRule, ownerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ModuleTransferCoinToErc20Iterator{contract: _ERC20Module.contract, event: "TransferCoinToErc20", logs: logs, sub: sub}, nil
}

// WatchTransferCoinToErc20 is a free log subscription operation binding the contract event 0xaa6b0040cf4e38b34df3fd01e51cc906a049cfcf3d5a6b81a9ac19dedcf533f7.
//
// Solidity: event TransferCoinToErc20(string indexed denom, address indexed owner, address indexed recipient, (uint256,string)[] amount)
func (_ERC20Module *ERC20ModuleFilterer) WatchTransferCoinToErc20(opts *bind.WatchOpts, sink chan<- *ERC20ModuleTransferCoinToErc20, denom []string, owner []common.Address, recipient []common.Address) (event.Subscription, error) {

	var denomRule []interface{}
	for _, denomItem := range denom {
		denomRule = append(denomRule, denomItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20Module.contract.WatchLogs(opts, "TransferCoinToErc20", denomRule, ownerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20ModuleTransferCoinToErc20)
				if err := _ERC20Module.contract.UnpackLog(event, "TransferCoinToErc20", log); err != nil {
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

// ParseTransferCoinToErc20 is a log parse operation binding the contract event 0xaa6b0040cf4e38b34df3fd01e51cc906a049cfcf3d5a6b81a9ac19dedcf533f7.
//
// Solidity: event TransferCoinToErc20(string indexed denom, address indexed owner, address indexed recipient, (uint256,string)[] amount)
func (_ERC20Module *ERC20ModuleFilterer) ParseTransferCoinToErc20(log types.Log) (*ERC20ModuleTransferCoinToErc20, error) {
	event := new(ERC20ModuleTransferCoinToErc20)
	if err := _ERC20Module.contract.UnpackLog(event, "TransferCoinToErc20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20ModuleTransferErc20ToCoinIterator is returned from FilterTransferErc20ToCoin and is used to iterate over the raw logs and unpacked data for TransferErc20ToCoin events raised by the ERC20Module contract.
type ERC20ModuleTransferErc20ToCoinIterator struct {
	Event *ERC20ModuleTransferErc20ToCoin // Event containing the contract specifics and raw log

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
func (it *ERC20ModuleTransferErc20ToCoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20ModuleTransferErc20ToCoin)
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
		it.Event = new(ERC20ModuleTransferErc20ToCoin)
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
func (it *ERC20ModuleTransferErc20ToCoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ModuleTransferErc20ToCoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20ModuleTransferErc20ToCoin represents a TransferErc20ToCoin event raised by the ERC20Module contract.
type ERC20ModuleTransferErc20ToCoin struct {
	Token     common.Address
	Owner     common.Address
	Recipient common.Address
	Amount    []IERC20ModuleCoin
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransferErc20ToCoin is a free log retrieval operation binding the contract event 0x32c4702a0530d1dd750fcee1239d68494583ac2b6084a1d42ad10060f3f53eb0.
//
// Solidity: event TransferErc20ToCoin(address indexed token, address indexed owner, address indexed recipient, (uint256,string)[] amount)
func (_ERC20Module *ERC20ModuleFilterer) FilterTransferErc20ToCoin(opts *bind.FilterOpts, token []common.Address, owner []common.Address, recipient []common.Address) (*ERC20ModuleTransferErc20ToCoinIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20Module.contract.FilterLogs(opts, "TransferErc20ToCoin", tokenRule, ownerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ModuleTransferErc20ToCoinIterator{contract: _ERC20Module.contract, event: "TransferErc20ToCoin", logs: logs, sub: sub}, nil
}

// WatchTransferErc20ToCoin is a free log subscription operation binding the contract event 0x32c4702a0530d1dd750fcee1239d68494583ac2b6084a1d42ad10060f3f53eb0.
//
// Solidity: event TransferErc20ToCoin(address indexed token, address indexed owner, address indexed recipient, (uint256,string)[] amount)
func (_ERC20Module *ERC20ModuleFilterer) WatchTransferErc20ToCoin(opts *bind.WatchOpts, sink chan<- *ERC20ModuleTransferErc20ToCoin, token []common.Address, owner []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20Module.contract.WatchLogs(opts, "TransferErc20ToCoin", tokenRule, ownerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20ModuleTransferErc20ToCoin)
				if err := _ERC20Module.contract.UnpackLog(event, "TransferErc20ToCoin", log); err != nil {
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

// ParseTransferErc20ToCoin is a log parse operation binding the contract event 0x32c4702a0530d1dd750fcee1239d68494583ac2b6084a1d42ad10060f3f53eb0.
//
// Solidity: event TransferErc20ToCoin(address indexed token, address indexed owner, address indexed recipient, (uint256,string)[] amount)
func (_ERC20Module *ERC20ModuleFilterer) ParseTransferErc20ToCoin(log types.Log) (*ERC20ModuleTransferErc20ToCoin, error) {
	event := new(ERC20ModuleTransferErc20ToCoin)
	if err := _ERC20Module.contract.UnpackLog(event, "TransferErc20ToCoin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
