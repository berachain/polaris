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

// IAuthModuleCoin is an auto generated low-level Go binding around an user-defined struct.
type IAuthModuleCoin struct {
	Amount *big.Int
	Denom  string
}

// AuthModuleMetaData contains all meta data concerning the AuthModule contract.
var AuthModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"account\",\"type\":\"string\"}],\"name\":\"convertBech32ToHexAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"convertHexToBech32\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"granter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structIAuthModule.Coin[]\",\"name\":\"limit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"sendGrant\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AuthModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use AuthModuleMetaData.ABI instead.
var AuthModuleABI = AuthModuleMetaData.ABI

// AuthModule is an auto generated Go binding around an Ethereum contract.
type AuthModule struct {
	AuthModuleCaller     // Read-only binding to the contract
	AuthModuleTransactor // Write-only binding to the contract
	AuthModuleFilterer   // Log filterer for contract events
}

// AuthModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuthModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuthModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuthModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuthModuleSession struct {
	Contract     *AuthModule       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuthModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuthModuleCallerSession struct {
	Contract *AuthModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AuthModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuthModuleTransactorSession struct {
	Contract     *AuthModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AuthModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuthModuleRaw struct {
	Contract *AuthModule // Generic contract binding to access the raw methods on
}

// AuthModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuthModuleCallerRaw struct {
	Contract *AuthModuleCaller // Generic read-only contract binding to access the raw methods on
}

// AuthModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuthModuleTransactorRaw struct {
	Contract *AuthModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuthModule creates a new instance of AuthModule, bound to a specific deployed contract.
func NewAuthModule(address common.Address, backend bind.ContractBackend) (*AuthModule, error) {
	contract, err := bindAuthModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuthModule{AuthModuleCaller: AuthModuleCaller{contract: contract}, AuthModuleTransactor: AuthModuleTransactor{contract: contract}, AuthModuleFilterer: AuthModuleFilterer{contract: contract}}, nil
}

// NewAuthModuleCaller creates a new read-only instance of AuthModule, bound to a specific deployed contract.
func NewAuthModuleCaller(address common.Address, caller bind.ContractCaller) (*AuthModuleCaller, error) {
	contract, err := bindAuthModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuthModuleCaller{contract: contract}, nil
}

// NewAuthModuleTransactor creates a new write-only instance of AuthModule, bound to a specific deployed contract.
func NewAuthModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*AuthModuleTransactor, error) {
	contract, err := bindAuthModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuthModuleTransactor{contract: contract}, nil
}

// NewAuthModuleFilterer creates a new log filterer instance of AuthModule, bound to a specific deployed contract.
func NewAuthModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*AuthModuleFilterer, error) {
	contract, err := bindAuthModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuthModuleFilterer{contract: contract}, nil
}

// bindAuthModule binds a generic wrapper to an already deployed contract.
func bindAuthModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AuthModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuthModule *AuthModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuthModule.Contract.AuthModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuthModule *AuthModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuthModule.Contract.AuthModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuthModule *AuthModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuthModule.Contract.AuthModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuthModule *AuthModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuthModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuthModule *AuthModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuthModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuthModule *AuthModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuthModule.Contract.contract.Transact(opts, method, params...)
}

// ConvertBech32ToHexAddress is a free data retrieval call binding the contract method 0xc769a484.
//
// Solidity: function convertBech32ToHexAddress(string account) view returns(address)
func (_AuthModule *AuthModuleCaller) ConvertBech32ToHexAddress(opts *bind.CallOpts, account string) (common.Address, error) {
	var out []interface{}
	err := _AuthModule.contract.Call(opts, &out, "convertBech32ToHexAddress", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConvertBech32ToHexAddress is a free data retrieval call binding the contract method 0xc769a484.
//
// Solidity: function convertBech32ToHexAddress(string account) view returns(address)
func (_AuthModule *AuthModuleSession) ConvertBech32ToHexAddress(account string) (common.Address, error) {
	return _AuthModule.Contract.ConvertBech32ToHexAddress(&_AuthModule.CallOpts, account)
}

// ConvertBech32ToHexAddress is a free data retrieval call binding the contract method 0xc769a484.
//
// Solidity: function convertBech32ToHexAddress(string account) view returns(address)
func (_AuthModule *AuthModuleCallerSession) ConvertBech32ToHexAddress(account string) (common.Address, error) {
	return _AuthModule.Contract.ConvertBech32ToHexAddress(&_AuthModule.CallOpts, account)
}

// ConvertHexToBech32 is a free data retrieval call binding the contract method 0x25435c5d.
//
// Solidity: function convertHexToBech32(address account) view returns(string)
func (_AuthModule *AuthModuleCaller) ConvertHexToBech32(opts *bind.CallOpts, account common.Address) (string, error) {
	var out []interface{}
	err := _AuthModule.contract.Call(opts, &out, "convertHexToBech32", account)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ConvertHexToBech32 is a free data retrieval call binding the contract method 0x25435c5d.
//
// Solidity: function convertHexToBech32(address account) view returns(string)
func (_AuthModule *AuthModuleSession) ConvertHexToBech32(account common.Address) (string, error) {
	return _AuthModule.Contract.ConvertHexToBech32(&_AuthModule.CallOpts, account)
}

// ConvertHexToBech32 is a free data retrieval call binding the contract method 0x25435c5d.
//
// Solidity: function convertHexToBech32(address account) view returns(string)
func (_AuthModule *AuthModuleCallerSession) ConvertHexToBech32(account common.Address) (string, error) {
	return _AuthModule.Contract.ConvertHexToBech32(&_AuthModule.CallOpts, account)
}

// SendGrant is a paid mutator transaction binding the contract method 0x9d66e40e.
//
// Solidity: function sendGrant(address granter, address grantee, (uint256,string)[] limit, uint256 expiration) returns(bool)
func (_AuthModule *AuthModuleTransactor) SendGrant(opts *bind.TransactOpts, granter common.Address, grantee common.Address, limit []IAuthModuleCoin, expiration *big.Int) (*types.Transaction, error) {
	return _AuthModule.contract.Transact(opts, "sendGrant", granter, grantee, limit, expiration)
}

// SendGrant is a paid mutator transaction binding the contract method 0x9d66e40e.
//
// Solidity: function sendGrant(address granter, address grantee, (uint256,string)[] limit, uint256 expiration) returns(bool)
func (_AuthModule *AuthModuleSession) SendGrant(granter common.Address, grantee common.Address, limit []IAuthModuleCoin, expiration *big.Int) (*types.Transaction, error) {
	return _AuthModule.Contract.SendGrant(&_AuthModule.TransactOpts, granter, grantee, limit, expiration)
}

// SendGrant is a paid mutator transaction binding the contract method 0x9d66e40e.
//
// Solidity: function sendGrant(address granter, address grantee, (uint256,string)[] limit, uint256 expiration) returns(bool)
func (_AuthModule *AuthModuleTransactorSession) SendGrant(granter common.Address, grantee common.Address, limit []IAuthModuleCoin, expiration *big.Int) (*types.Transaction, error) {
	return _AuthModule.Contract.SendGrant(&_AuthModule.TransactOpts, granter, grantee, limit, expiration)
}
