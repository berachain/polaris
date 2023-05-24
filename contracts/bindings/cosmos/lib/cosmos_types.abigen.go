// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lib

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

// CosmosCoin is an auto generated low-level Go binding around an user-defined struct.
type CosmosCoin struct {
	Amount *big.Int
	Denom  string
}

// CosmosTypesMetaData contains all meta data concerning the CosmosTypes contract.
var CosmosTypesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structCosmos.Coin\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"coin\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060ea8061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80632ff6e5df14602d575b600080fd5b60436004803603810190603f91906073565b6045565b005b50565b600080fd5b600080fd5b600080fd5b600060408284031215606a5760696052565b5b81905092915050565b60006020828403121560865760856048565b5b600082013567ffffffffffffffff81111560a15760a0604d565b5b60ab848285016057565b9150509291505056fea26469706673582212204aea476627915527298b66a6628d095d342cb40837d32e1c0d8191cff189589164736f6c63430008130033",
}

// CosmosTypesABI is the input ABI used to generate the binding from.
// Deprecated: Use CosmosTypesMetaData.ABI instead.
var CosmosTypesABI = CosmosTypesMetaData.ABI

// CosmosTypesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CosmosTypesMetaData.Bin instead.
var CosmosTypesBin = CosmosTypesMetaData.Bin

// DeployCosmosTypes deploys a new Ethereum contract, binding an instance of CosmosTypes to it.
func DeployCosmosTypes(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CosmosTypes, error) {
	parsed, err := CosmosTypesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CosmosTypesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CosmosTypes{CosmosTypesCaller: CosmosTypesCaller{contract: contract}, CosmosTypesTransactor: CosmosTypesTransactor{contract: contract}, CosmosTypesFilterer: CosmosTypesFilterer{contract: contract}}, nil
}

// CosmosTypes is an auto generated Go binding around an Ethereum contract.
type CosmosTypes struct {
	CosmosTypesCaller     // Read-only binding to the contract
	CosmosTypesTransactor // Write-only binding to the contract
	CosmosTypesFilterer   // Log filterer for contract events
}

// CosmosTypesCaller is an auto generated read-only Go binding around an Ethereum contract.
type CosmosTypesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmosTypesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CosmosTypesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmosTypesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CosmosTypesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmosTypesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CosmosTypesSession struct {
	Contract     *CosmosTypes      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CosmosTypesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CosmosTypesCallerSession struct {
	Contract *CosmosTypesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CosmosTypesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CosmosTypesTransactorSession struct {
	Contract     *CosmosTypesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CosmosTypesRaw is an auto generated low-level Go binding around an Ethereum contract.
type CosmosTypesRaw struct {
	Contract *CosmosTypes // Generic contract binding to access the raw methods on
}

// CosmosTypesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CosmosTypesCallerRaw struct {
	Contract *CosmosTypesCaller // Generic read-only contract binding to access the raw methods on
}

// CosmosTypesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CosmosTypesTransactorRaw struct {
	Contract *CosmosTypesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCosmosTypes creates a new instance of CosmosTypes, bound to a specific deployed contract.
func NewCosmosTypes(address common.Address, backend bind.ContractBackend) (*CosmosTypes, error) {
	contract, err := bindCosmosTypes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CosmosTypes{CosmosTypesCaller: CosmosTypesCaller{contract: contract}, CosmosTypesTransactor: CosmosTypesTransactor{contract: contract}, CosmosTypesFilterer: CosmosTypesFilterer{contract: contract}}, nil
}

// NewCosmosTypesCaller creates a new read-only instance of CosmosTypes, bound to a specific deployed contract.
func NewCosmosTypesCaller(address common.Address, caller bind.ContractCaller) (*CosmosTypesCaller, error) {
	contract, err := bindCosmosTypes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CosmosTypesCaller{contract: contract}, nil
}

// NewCosmosTypesTransactor creates a new write-only instance of CosmosTypes, bound to a specific deployed contract.
func NewCosmosTypesTransactor(address common.Address, transactor bind.ContractTransactor) (*CosmosTypesTransactor, error) {
	contract, err := bindCosmosTypes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CosmosTypesTransactor{contract: contract}, nil
}

// NewCosmosTypesFilterer creates a new log filterer instance of CosmosTypes, bound to a specific deployed contract.
func NewCosmosTypesFilterer(address common.Address, filterer bind.ContractFilterer) (*CosmosTypesFilterer, error) {
	contract, err := bindCosmosTypes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CosmosTypesFilterer{contract: contract}, nil
}

// bindCosmosTypes binds a generic wrapper to an already deployed contract.
func bindCosmosTypes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CosmosTypesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CosmosTypes *CosmosTypesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CosmosTypes.Contract.CosmosTypesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CosmosTypes *CosmosTypesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmosTypes.Contract.CosmosTypesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CosmosTypes *CosmosTypesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CosmosTypes.Contract.CosmosTypesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CosmosTypes *CosmosTypesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CosmosTypes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CosmosTypes *CosmosTypesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmosTypes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CosmosTypes *CosmosTypesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CosmosTypes.Contract.contract.Transact(opts, method, params...)
}

// Coin is a free data retrieval call binding the contract method 0x2ff6e5df.
//
// Solidity: function coin((uint256,string) ) pure returns()
func (_CosmosTypes *CosmosTypesCaller) Coin(opts *bind.CallOpts, arg0 CosmosCoin) error {
	var out []interface{}
	err := _CosmosTypes.contract.Call(opts, &out, "coin", arg0)

	if err != nil {
		return err
	}

	return err

}

// Coin is a free data retrieval call binding the contract method 0x2ff6e5df.
//
// Solidity: function coin((uint256,string) ) pure returns()
func (_CosmosTypes *CosmosTypesSession) Coin(arg0 CosmosCoin) error {
	return _CosmosTypes.Contract.Coin(&_CosmosTypes.CallOpts, arg0)
}

// Coin is a free data retrieval call binding the contract method 0x2ff6e5df.
//
// Solidity: function coin((uint256,string) ) pure returns()
func (_CosmosTypes *CosmosTypesCallerSession) Coin(arg0 CosmosCoin) error {
	return _CosmosTypes.Contract.Coin(&_CosmosTypes.CallOpts, arg0)
}
