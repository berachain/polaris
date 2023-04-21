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

// PrecompileConstructorMetaData contains all meta data concerning the PrecompileConstructor contract.
var PrecompileConstructorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"abera\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denom\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20Module\",\"outputs\":[{\"internalType\":\"contractIERC20Module\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040526269696973ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff168152503480156200004757600080fd5b50600060805173ffffffffffffffffffffffffffffffffffffffff1663dbeeeb5c63075bcd156040518263ffffffff1660e01b81526004016200008b9190620003de565b6020604051808303816000875af1158015620000ab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000d1919062000461565b90508062000116576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200010d90620004e3565b60405180910390fd5b60805173ffffffffffffffffffffffffffffffffffffffff1663a333e57c6040518163ffffffff1660e01b8152600401620001519062000505565b602060405180830381865afa1580156200016f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001959190620005a0565b6000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060805173ffffffffffffffffffffffffffffffffffffffff1663cd22a01860008054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b815260040162000231919062000633565b600060405180830381865afa1580156200024f573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906200027a9190620007cf565b600190816200028a919062000a57565b506040516020016200029c9062000b70565b604051602081830303815290604052805190602001206001604051602001620002c6919062000c16565b60405160208183030381529060405280519060200120146200031f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620003169062000c7f565b60405180910390fd5b5062000ca1565b600082825260208201905092915050565b7f6162657261000000000000000000000000000000000000000000000000000000600082015250565b60006200036f60058362000326565b91506200037c8262000337565b602082019050919050565b6000819050919050565b6000819050919050565b6000819050919050565b6000620003c6620003c0620003ba8462000387565b6200039b565b62000391565b9050919050565b620003d881620003a5565b82525050565b60006040820190508181036000830152620003f98162000360565b90506200040a6020830184620003cd565b92915050565b6000604051905090565b600080fd5b600080fd5b60008115159050919050565b6200043b8162000424565b81146200044757600080fd5b50565b6000815190506200045b8162000430565b92915050565b6000602082840312156200047a57620004796200041a565b5b60006200048a848285016200044a565b91505092915050565b7f6661696c656420746f20636f6e76657274206162657261000000000000000000600082015250565b6000620004cb60178362000326565b9150620004d88262000493565b602082019050919050565b60006020820190508181036000830152620004fe81620004bc565b9050919050565b60006020820190508181036000830152620005208162000360565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620005548262000527565b9050919050565b6000620005688262000547565b9050919050565b6200057a816200055b565b81146200058657600080fd5b50565b6000815190506200059a816200056f565b92915050565b600060208284031215620005b957620005b86200041a565b5b6000620005c98482850162000589565b91505092915050565b6000620005f3620005ed620005e78462000527565b6200039b565b62000527565b9050919050565b60006200060782620005d2565b9050919050565b60006200061b82620005fa565b9050919050565b6200062d816200060e565b82525050565b60006020820190506200064a600083018462000622565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620006a5826200065a565b810181811067ffffffffffffffff82111715620006c757620006c66200066b565b5b80604052505050565b6000620006dc62000410565b9050620006ea82826200069a565b919050565b600067ffffffffffffffff8211156200070d576200070c6200066b565b5b62000718826200065a565b9050602081019050919050565b60005b838110156200074557808201518184015260208101905062000728565b60008484015250505050565b6000620007686200076284620006ef565b620006d0565b90508281526020810184848401111562000787576200078662000655565b5b6200079484828562000725565b509392505050565b600082601f830112620007b457620007b362000650565b5b8151620007c684826020860162000751565b91505092915050565b600060208284031215620007e857620007e76200041a565b5b600082015167ffffffffffffffff8111156200080957620008086200041f565b5b62000817848285016200079c565b91505092915050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200087357607f821691505b6020821081036200088957620008886200082b565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620008f37fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620008b4565b620008ff8683620008b4565b95508019841693508086168417925050509392505050565b600062000938620009326200092c8462000391565b6200039b565b62000391565b9050919050565b6000819050919050565b620009548362000917565b6200096c62000963826200093f565b848454620008c1565b825550505050565b600090565b6200098362000974565b6200099081848462000949565b505050565b5b81811015620009b857620009ac60008262000979565b60018101905062000996565b5050565b601f82111562000a0757620009d1816200088f565b620009dc84620008a4565b81016020851015620009ec578190505b62000a04620009fb85620008a4565b83018262000995565b50505b505050565b600082821c905092915050565b600062000a2c6000198460080262000a0c565b1980831691505092915050565b600062000a47838362000a19565b9150826002028217905092915050565b62000a628262000820565b67ffffffffffffffff81111562000a7e5762000a7d6200066b565b5b62000a8a82546200085a565b62000a97828285620009bc565b600060209050601f83116001811462000acf576000841562000aba578287015190505b62000ac6858262000a39565b86555062000b36565b601f19841662000adf866200088f565b60005b8281101562000b095784890151825560018201915060208501945060208101905062000ae2565b8683101562000b29578489015162000b25601f89168262000a19565b8355505b6001600288020188555050505b505050505050565b600081905092915050565b600062000b5860058362000b3e565b915062000b658262000337565b600582019050919050565b600062000b7d8262000b49565b9150819050919050565b6000815462000b96816200085a565b62000ba2818662000b3e565b9450600182166000811462000bc0576001811462000bd65762000c0d565b60ff198316865281151582028601935062000c0d565b62000be1856200088f565b60005b8381101562000c055781548189015260018201915060208101905062000be4565b838801955050505b50505092915050565b600062000c24828462000b87565b915081905092915050565b7f72657475726e6564207468652077726f6e672064656e6f6d0000000000000000600082015250565b600062000c6760188362000326565b915062000c748262000c2f565b602082019050919050565b6000602082019050818103600083015262000c9a8162000c58565b9050919050565b60805161039462000cbc600039600060c601526103946000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063558f208414610046578063714ba40c14610064578063c370b04214610082575b600080fd5b61004e6100a0565b60405161005b91906101f5565b60405180910390f35b61006c6100c4565b6040516100799190610231565b60405180910390f35b61008a6100e8565b60405161009791906102dc565b60405180910390f35b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b600180546100f59061032d565b80601f01602080910402602001604051908101604052809291908181526020018280546101219061032d565b801561016e5780601f106101435761010080835404028352916020019161016e565b820191906000526020600020905b81548152906001019060200180831161015157829003601f168201915b505050505081565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006101bb6101b66101b184610176565b610196565b610176565b9050919050565b60006101cd826101a0565b9050919050565b60006101df826101c2565b9050919050565b6101ef816101d4565b82525050565b600060208201905061020a60008301846101e6565b92915050565b600061021b826101c2565b9050919050565b61022b81610210565b82525050565b60006020820190506102466000830184610222565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561028657808201518184015260208101905061026b565b60008484015250505050565b6000601f19601f8301169050919050565b60006102ae8261024c565b6102b88185610257565b93506102c8818560208601610268565b6102d181610292565b840191505092915050565b600060208201905081810360008301526102f681846102a3565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061034557607f821691505b602082108103610358576103576102fe565b5b5091905056fea2646970667358221220a2fbfce11484d1dfdbaf53f023ba63f8b0445096dcd9700e2cbefd8d05c4e6e864736f6c63430008130033",
}

// PrecompileConstructorABI is the input ABI used to generate the binding from.
// Deprecated: Use PrecompileConstructorMetaData.ABI instead.
var PrecompileConstructorABI = PrecompileConstructorMetaData.ABI

// PrecompileConstructorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PrecompileConstructorMetaData.Bin instead.
var PrecompileConstructorBin = PrecompileConstructorMetaData.Bin

// DeployPrecompileConstructor deploys a new Ethereum contract, binding an instance of PrecompileConstructor to it.
func DeployPrecompileConstructor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PrecompileConstructor, error) {
	parsed, err := PrecompileConstructorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PrecompileConstructorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PrecompileConstructor{PrecompileConstructorCaller: PrecompileConstructorCaller{contract: contract}, PrecompileConstructorTransactor: PrecompileConstructorTransactor{contract: contract}, PrecompileConstructorFilterer: PrecompileConstructorFilterer{contract: contract}}, nil
}

// PrecompileConstructor is an auto generated Go binding around an Ethereum contract.
type PrecompileConstructor struct {
	PrecompileConstructorCaller     // Read-only binding to the contract
	PrecompileConstructorTransactor // Write-only binding to the contract
	PrecompileConstructorFilterer   // Log filterer for contract events
}

// PrecompileConstructorCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrecompileConstructorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompileConstructorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrecompileConstructorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompileConstructorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrecompileConstructorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompileConstructorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrecompileConstructorSession struct {
	Contract     *PrecompileConstructor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PrecompileConstructorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrecompileConstructorCallerSession struct {
	Contract *PrecompileConstructorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// PrecompileConstructorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrecompileConstructorTransactorSession struct {
	Contract     *PrecompileConstructorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// PrecompileConstructorRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrecompileConstructorRaw struct {
	Contract *PrecompileConstructor // Generic contract binding to access the raw methods on
}

// PrecompileConstructorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrecompileConstructorCallerRaw struct {
	Contract *PrecompileConstructorCaller // Generic read-only contract binding to access the raw methods on
}

// PrecompileConstructorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrecompileConstructorTransactorRaw struct {
	Contract *PrecompileConstructorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrecompileConstructor creates a new instance of PrecompileConstructor, bound to a specific deployed contract.
func NewPrecompileConstructor(address common.Address, backend bind.ContractBackend) (*PrecompileConstructor, error) {
	contract, err := bindPrecompileConstructor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PrecompileConstructor{PrecompileConstructorCaller: PrecompileConstructorCaller{contract: contract}, PrecompileConstructorTransactor: PrecompileConstructorTransactor{contract: contract}, PrecompileConstructorFilterer: PrecompileConstructorFilterer{contract: contract}}, nil
}

// NewPrecompileConstructorCaller creates a new read-only instance of PrecompileConstructor, bound to a specific deployed contract.
func NewPrecompileConstructorCaller(address common.Address, caller bind.ContractCaller) (*PrecompileConstructorCaller, error) {
	contract, err := bindPrecompileConstructor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrecompileConstructorCaller{contract: contract}, nil
}

// NewPrecompileConstructorTransactor creates a new write-only instance of PrecompileConstructor, bound to a specific deployed contract.
func NewPrecompileConstructorTransactor(address common.Address, transactor bind.ContractTransactor) (*PrecompileConstructorTransactor, error) {
	contract, err := bindPrecompileConstructor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrecompileConstructorTransactor{contract: contract}, nil
}

// NewPrecompileConstructorFilterer creates a new log filterer instance of PrecompileConstructor, bound to a specific deployed contract.
func NewPrecompileConstructorFilterer(address common.Address, filterer bind.ContractFilterer) (*PrecompileConstructorFilterer, error) {
	contract, err := bindPrecompileConstructor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrecompileConstructorFilterer{contract: contract}, nil
}

// bindPrecompileConstructor binds a generic wrapper to an already deployed contract.
func bindPrecompileConstructor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PrecompileConstructorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrecompileConstructor *PrecompileConstructorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrecompileConstructor.Contract.PrecompileConstructorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrecompileConstructor *PrecompileConstructorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrecompileConstructor.Contract.PrecompileConstructorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrecompileConstructor *PrecompileConstructorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrecompileConstructor.Contract.PrecompileConstructorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrecompileConstructor *PrecompileConstructorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrecompileConstructor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrecompileConstructor *PrecompileConstructorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrecompileConstructor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrecompileConstructor *PrecompileConstructorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrecompileConstructor.Contract.contract.Transact(opts, method, params...)
}

// Abera is a free data retrieval call binding the contract method 0x558f2084.
//
// Solidity: function abera() view returns(address)
func (_PrecompileConstructor *PrecompileConstructorCaller) Abera(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrecompileConstructor.contract.Call(opts, &out, "abera")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Abera is a free data retrieval call binding the contract method 0x558f2084.
//
// Solidity: function abera() view returns(address)
func (_PrecompileConstructor *PrecompileConstructorSession) Abera() (common.Address, error) {
	return _PrecompileConstructor.Contract.Abera(&_PrecompileConstructor.CallOpts)
}

// Abera is a free data retrieval call binding the contract method 0x558f2084.
//
// Solidity: function abera() view returns(address)
func (_PrecompileConstructor *PrecompileConstructorCallerSession) Abera() (common.Address, error) {
	return _PrecompileConstructor.Contract.Abera(&_PrecompileConstructor.CallOpts)
}

// Denom is a free data retrieval call binding the contract method 0xc370b042.
//
// Solidity: function denom() view returns(string)
func (_PrecompileConstructor *PrecompileConstructorCaller) Denom(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PrecompileConstructor.contract.Call(opts, &out, "denom")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Denom is a free data retrieval call binding the contract method 0xc370b042.
//
// Solidity: function denom() view returns(string)
func (_PrecompileConstructor *PrecompileConstructorSession) Denom() (string, error) {
	return _PrecompileConstructor.Contract.Denom(&_PrecompileConstructor.CallOpts)
}

// Denom is a free data retrieval call binding the contract method 0xc370b042.
//
// Solidity: function denom() view returns(string)
func (_PrecompileConstructor *PrecompileConstructorCallerSession) Denom() (string, error) {
	return _PrecompileConstructor.Contract.Denom(&_PrecompileConstructor.CallOpts)
}

// Erc20Module is a free data retrieval call binding the contract method 0x714ba40c.
//
// Solidity: function erc20Module() view returns(address)
func (_PrecompileConstructor *PrecompileConstructorCaller) Erc20Module(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrecompileConstructor.contract.Call(opts, &out, "erc20Module")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20Module is a free data retrieval call binding the contract method 0x714ba40c.
//
// Solidity: function erc20Module() view returns(address)
func (_PrecompileConstructor *PrecompileConstructorSession) Erc20Module() (common.Address, error) {
	return _PrecompileConstructor.Contract.Erc20Module(&_PrecompileConstructor.CallOpts)
}

// Erc20Module is a free data retrieval call binding the contract method 0x714ba40c.
//
// Solidity: function erc20Module() view returns(address)
func (_PrecompileConstructor *PrecompileConstructorCallerSession) Erc20Module() (common.Address, error) {
	return _PrecompileConstructor.Contract.Erc20Module(&_PrecompileConstructor.CallOpts)
}
