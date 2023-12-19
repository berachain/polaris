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

// CosmosCodecAny is an auto generated low-level Go binding around an user-defined struct.
type CosmosCodecAny struct {
	TypeURL string
	Value   []byte
}

// CosmosCoin is an auto generated low-level Go binding around an user-defined struct.
type CosmosCoin struct {
	Amount *big.Int
	Denom  string
}

// CosmosPageRequest is an auto generated low-level Go binding around an user-defined struct.
type CosmosPageRequest struct {
	Key        string
	Offset     uint64
	Limit      uint64
	CountTotal bool
	Reverse    bool
}

// CosmosPageResponse is an auto generated low-level Go binding around an user-defined struct.
type CosmosPageResponse struct {
	NextKey string
	Total   uint64
}

// CosmosTypesMetaData contains all meta data concerning the CosmosTypes contract.
var CosmosTypesMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"codecAny\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCosmos.CodecAny\",\"components\":[{\"name\":\"typeURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"coin\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCosmos.Coin\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"pageRequest\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCosmos.PageRequest\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"offset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"limit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"countTotal\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"reverse\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"pageResponse\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCosmos.PageResponse\",\"components\":[{\"name\":\"nextKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"total\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"pure\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506102a08061001d5f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c80631713cb131461004e5780631acc976f1461006a5780632ff6e5df14610086578063426ce1a9146100a2575b5f80fd5b610068600480360381019061006391906100f4565b6100be565b005b610084600480360381019061007f9190610159565b6100c1565b005b6100a0600480360381019061009b91906101be565b6100c4565b005b6100bc60048036038101906100b79190610223565b6100c7565b005b50565b50565b50565b50565b5f80fd5b5f80fd5b5f80fd5b5f604082840312156100eb576100ea6100d2565b5b81905092915050565b5f60208284031215610109576101086100ca565b5b5f82013567ffffffffffffffff811115610126576101256100ce565b5b610132848285016100d6565b91505092915050565b5f60a082840312156101505761014f6100d2565b5b81905092915050565b5f6020828403121561016e5761016d6100ca565b5b5f82013567ffffffffffffffff81111561018b5761018a6100ce565b5b6101978482850161013b565b91505092915050565b5f604082840312156101b5576101b46100d2565b5b81905092915050565b5f602082840312156101d3576101d26100ca565b5b5f82013567ffffffffffffffff8111156101f0576101ef6100ce565b5b6101fc848285016101a0565b91505092915050565b5f6040828403121561021a576102196100d2565b5b81905092915050565b5f60208284031215610238576102376100ca565b5b5f82013567ffffffffffffffff811115610255576102546100ce565b5b61026184828501610205565b9150509291505056fea2646970667358221220812e2905b13508a6c5b348d8954a7ae59e7d366f2aa90b5e10912eb299428c7864736f6c63430008170033",
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

// CodecAny is a free data retrieval call binding the contract method 0x1713cb13.
//
// Solidity: function codecAny((string,bytes) ) pure returns()
func (_CosmosTypes *CosmosTypesCaller) CodecAny(opts *bind.CallOpts, arg0 CosmosCodecAny) error {
	var out []interface{}
	err := _CosmosTypes.contract.Call(opts, &out, "codecAny", arg0)

	if err != nil {
		return err
	}

	return err

}

// CodecAny is a free data retrieval call binding the contract method 0x1713cb13.
//
// Solidity: function codecAny((string,bytes) ) pure returns()
func (_CosmosTypes *CosmosTypesSession) CodecAny(arg0 CosmosCodecAny) error {
	return _CosmosTypes.Contract.CodecAny(&_CosmosTypes.CallOpts, arg0)
}

// CodecAny is a free data retrieval call binding the contract method 0x1713cb13.
//
// Solidity: function codecAny((string,bytes) ) pure returns()
func (_CosmosTypes *CosmosTypesCallerSession) CodecAny(arg0 CosmosCodecAny) error {
	return _CosmosTypes.Contract.CodecAny(&_CosmosTypes.CallOpts, arg0)
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

// PageRequest is a free data retrieval call binding the contract method 0x1acc976f.
//
// Solidity: function pageRequest((string,uint64,uint64,bool,bool) ) pure returns()
func (_CosmosTypes *CosmosTypesCaller) PageRequest(opts *bind.CallOpts, arg0 CosmosPageRequest) error {
	var out []interface{}
	err := _CosmosTypes.contract.Call(opts, &out, "pageRequest", arg0)

	if err != nil {
		return err
	}

	return err

}

// PageRequest is a free data retrieval call binding the contract method 0x1acc976f.
//
// Solidity: function pageRequest((string,uint64,uint64,bool,bool) ) pure returns()
func (_CosmosTypes *CosmosTypesSession) PageRequest(arg0 CosmosPageRequest) error {
	return _CosmosTypes.Contract.PageRequest(&_CosmosTypes.CallOpts, arg0)
}

// PageRequest is a free data retrieval call binding the contract method 0x1acc976f.
//
// Solidity: function pageRequest((string,uint64,uint64,bool,bool) ) pure returns()
func (_CosmosTypes *CosmosTypesCallerSession) PageRequest(arg0 CosmosPageRequest) error {
	return _CosmosTypes.Contract.PageRequest(&_CosmosTypes.CallOpts, arg0)
}

// PageResponse is a free data retrieval call binding the contract method 0x426ce1a9.
//
// Solidity: function pageResponse((string,uint64) ) pure returns()
func (_CosmosTypes *CosmosTypesCaller) PageResponse(opts *bind.CallOpts, arg0 CosmosPageResponse) error {
	var out []interface{}
	err := _CosmosTypes.contract.Call(opts, &out, "pageResponse", arg0)

	if err != nil {
		return err
	}

	return err

}

// PageResponse is a free data retrieval call binding the contract method 0x426ce1a9.
//
// Solidity: function pageResponse((string,uint64) ) pure returns()
func (_CosmosTypes *CosmosTypesSession) PageResponse(arg0 CosmosPageResponse) error {
	return _CosmosTypes.Contract.PageResponse(&_CosmosTypes.CallOpts, arg0)
}

// PageResponse is a free data retrieval call binding the contract method 0x426ce1a9.
//
// Solidity: function pageResponse((string,uint64) ) pure returns()
func (_CosmosTypes *CosmosTypesCallerSession) PageResponse(arg0 CosmosPageResponse) error {
	return _CosmosTypes.Contract.PageResponse(&_CosmosTypes.CallOpts, arg0)
}
