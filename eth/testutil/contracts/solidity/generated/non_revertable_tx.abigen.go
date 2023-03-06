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

// NonRevertableTxMetaData contains all meta data concerning the NonRevertableTx contract.
var NonRevertableTxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea26469706673582212205206a54f3e5dc5af32b3095ceba73837e3bd884281d28a81fc73db328196cbf164736f6c63430008110033",
}

// NonRevertableTxABI is the input ABI used to generate the binding from.
// Deprecated: Use NonRevertableTxMetaData.ABI instead.
var NonRevertableTxABI = NonRevertableTxMetaData.ABI

// NonRevertableTxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NonRevertableTxMetaData.Bin instead.
var NonRevertableTxBin = NonRevertableTxMetaData.Bin

// DeployNonRevertableTx deploys a new Ethereum contract, binding an instance of NonRevertableTx to it.
func DeployNonRevertableTx(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NonRevertableTx, error) {
	parsed, err := NonRevertableTxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NonRevertableTxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NonRevertableTx{NonRevertableTxCaller: NonRevertableTxCaller{contract: contract}, NonRevertableTxTransactor: NonRevertableTxTransactor{contract: contract}, NonRevertableTxFilterer: NonRevertableTxFilterer{contract: contract}}, nil
}

// NonRevertableTx is an auto generated Go binding around an Ethereum contract.
type NonRevertableTx struct {
	NonRevertableTxCaller     // Read-only binding to the contract
	NonRevertableTxTransactor // Write-only binding to the contract
	NonRevertableTxFilterer   // Log filterer for contract events
}

// NonRevertableTxCaller is an auto generated read-only Go binding around an Ethereum contract.
type NonRevertableTxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonRevertableTxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NonRevertableTxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonRevertableTxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NonRevertableTxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonRevertableTxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NonRevertableTxSession struct {
	Contract     *NonRevertableTx  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NonRevertableTxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NonRevertableTxCallerSession struct {
	Contract *NonRevertableTxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// NonRevertableTxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NonRevertableTxTransactorSession struct {
	Contract     *NonRevertableTxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// NonRevertableTxRaw is an auto generated low-level Go binding around an Ethereum contract.
type NonRevertableTxRaw struct {
	Contract *NonRevertableTx // Generic contract binding to access the raw methods on
}

// NonRevertableTxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NonRevertableTxCallerRaw struct {
	Contract *NonRevertableTxCaller // Generic read-only contract binding to access the raw methods on
}

// NonRevertableTxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NonRevertableTxTransactorRaw struct {
	Contract *NonRevertableTxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNonRevertableTx creates a new instance of NonRevertableTx, bound to a specific deployed contract.
func NewNonRevertableTx(address common.Address, backend bind.ContractBackend) (*NonRevertableTx, error) {
	contract, err := bindNonRevertableTx(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NonRevertableTx{NonRevertableTxCaller: NonRevertableTxCaller{contract: contract}, NonRevertableTxTransactor: NonRevertableTxTransactor{contract: contract}, NonRevertableTxFilterer: NonRevertableTxFilterer{contract: contract}}, nil
}

// NewNonRevertableTxCaller creates a new read-only instance of NonRevertableTx, bound to a specific deployed contract.
func NewNonRevertableTxCaller(address common.Address, caller bind.ContractCaller) (*NonRevertableTxCaller, error) {
	contract, err := bindNonRevertableTx(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NonRevertableTxCaller{contract: contract}, nil
}

// NewNonRevertableTxTransactor creates a new write-only instance of NonRevertableTx, bound to a specific deployed contract.
func NewNonRevertableTxTransactor(address common.Address, transactor bind.ContractTransactor) (*NonRevertableTxTransactor, error) {
	contract, err := bindNonRevertableTx(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NonRevertableTxTransactor{contract: contract}, nil
}

// NewNonRevertableTxFilterer creates a new log filterer instance of NonRevertableTx, bound to a specific deployed contract.
func NewNonRevertableTxFilterer(address common.Address, filterer bind.ContractFilterer) (*NonRevertableTxFilterer, error) {
	contract, err := bindNonRevertableTx(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NonRevertableTxFilterer{contract: contract}, nil
}

// bindNonRevertableTx binds a generic wrapper to an already deployed contract.
func bindNonRevertableTx(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NonRevertableTxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NonRevertableTx *NonRevertableTxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NonRevertableTx.Contract.NonRevertableTxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NonRevertableTx *NonRevertableTxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NonRevertableTx.Contract.NonRevertableTxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NonRevertableTx *NonRevertableTxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NonRevertableTx.Contract.NonRevertableTxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NonRevertableTx *NonRevertableTxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NonRevertableTx.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NonRevertableTx *NonRevertableTxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NonRevertableTx.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NonRevertableTx *NonRevertableTxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NonRevertableTx.Contract.contract.Transact(opts, method, params...)
}
