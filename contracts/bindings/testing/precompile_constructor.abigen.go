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
	Bin: "0x60a06040526269696973ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff168152503480156200004757600080fd5b50600060805173ffffffffffffffffffffffffffffffffffffffff166310e0f725333363075bcd156040518463ffffffff1660e01b81526004016200008f9392919062000427565b6020604051808303816000875af1158015620000af573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000d59190620004ca565b9050806200011a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000111906200054c565b60405180910390fd5b60805173ffffffffffffffffffffffffffffffffffffffff1663a333e57c6040518163ffffffff1660e01b815260040162000155906200056e565b602060405180830381865afa15801562000173573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001999190620005d5565b6000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060805173ffffffffffffffffffffffffffffffffffffffff1663cd22a01860008054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b815260040162000235919062000668565b600060405180830381865afa15801562000253573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906200027e919062000804565b600190816200028e919062000a8c565b50604051602001620002a09062000ba5565b604051602081830303815290604052805190602001206001604051602001620002ca919062000c4b565b604051602081830303815290604052805190602001201462000323576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200031a9062000cb4565b60405180910390fd5b5062000cd6565b600082825260208201905092915050565b7f6162657261000000000000000000000000000000000000000000000000000000600082015250565b6000620003736005836200032a565b915062000380826200033b565b602082019050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620003b8826200038b565b9050919050565b620003ca81620003ab565b82525050565b6000819050919050565b6000819050919050565b6000819050919050565b60006200040f620004096200040384620003d0565b620003e4565b620003da565b9050919050565b6200042181620003ee565b82525050565b60006080820190508181036000830152620004428162000364565b9050620004536020830186620003bf565b620004626040830185620003bf565b62000471606083018462000416565b949350505050565b6000604051905090565b600080fd5b600080fd5b60008115159050919050565b620004a4816200048d565b8114620004b057600080fd5b50565b600081519050620004c48162000499565b92915050565b600060208284031215620004e357620004e262000483565b5b6000620004f384828501620004b3565b91505092915050565b7f6661696c656420746f20636f6e76657274206162657261000000000000000000600082015250565b6000620005346017836200032a565b91506200054182620004fc565b602082019050919050565b60006020820190508181036000830152620005678162000525565b9050919050565b60006020820190508181036000830152620005898162000364565b9050919050565b60006200059d82620003ab565b9050919050565b620005af8162000590565b8114620005bb57600080fd5b50565b600081519050620005cf81620005a4565b92915050565b600060208284031215620005ee57620005ed62000483565b5b6000620005fe84828501620005be565b91505092915050565b600062000628620006226200061c846200038b565b620003e4565b6200038b565b9050919050565b60006200063c8262000607565b9050919050565b600062000650826200062f565b9050919050565b620006628162000643565b82525050565b60006020820190506200067f600083018462000657565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620006da826200068f565b810181811067ffffffffffffffff82111715620006fc57620006fb620006a0565b5b80604052505050565b60006200071162000479565b90506200071f8282620006cf565b919050565b600067ffffffffffffffff821115620007425762000741620006a0565b5b6200074d826200068f565b9050602081019050919050565b60005b838110156200077a5780820151818401526020810190506200075d565b60008484015250505050565b60006200079d620007978462000724565b62000705565b905082815260208101848484011115620007bc57620007bb6200068a565b5b620007c98482856200075a565b509392505050565b600082601f830112620007e957620007e862000685565b5b8151620007fb84826020860162000786565b91505092915050565b6000602082840312156200081d576200081c62000483565b5b600082015167ffffffffffffffff8111156200083e576200083d62000488565b5b6200084c84828501620007d1565b91505092915050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620008a857607f821691505b602082108103620008be57620008bd62000860565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620009287fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620008e9565b620009348683620008e9565b95508019841693508086168417925050509392505050565b60006200096d620009676200096184620003da565b620003e4565b620003da565b9050919050565b6000819050919050565b62000989836200094c565b620009a1620009988262000974565b848454620008f6565b825550505050565b600090565b620009b8620009a9565b620009c58184846200097e565b505050565b5b81811015620009ed57620009e1600082620009ae565b600181019050620009cb565b5050565b601f82111562000a3c5762000a0681620008c4565b62000a1184620008d9565b8101602085101562000a21578190505b62000a3962000a3085620008d9565b830182620009ca565b50505b505050565b600082821c905092915050565b600062000a616000198460080262000a41565b1980831691505092915050565b600062000a7c838362000a4e565b9150826002028217905092915050565b62000a978262000855565b67ffffffffffffffff81111562000ab35762000ab2620006a0565b5b62000abf82546200088f565b62000acc828285620009f1565b600060209050601f83116001811462000b04576000841562000aef578287015190505b62000afb858262000a6e565b86555062000b6b565b601f19841662000b1486620008c4565b60005b8281101562000b3e5784890151825560018201915060208501945060208101905062000b17565b8683101562000b5e578489015162000b5a601f89168262000a4e565b8355505b6001600288020188555050505b505050505050565b600081905092915050565b600062000b8d60058362000b73565b915062000b9a826200033b565b600582019050919050565b600062000bb28262000b7e565b9150819050919050565b6000815462000bcb816200088f565b62000bd7818662000b73565b9450600182166000811462000bf5576001811462000c0b5762000c42565b60ff198316865281151582028601935062000c42565b62000c1685620008c4565b60005b8381101562000c3a5781548189015260018201915060208101905062000c19565b838801955050505b50505092915050565b600062000c59828462000bbc565b915081905092915050565b7f72657475726e6564207468652077726f6e672064656e6f6d0000000000000000600082015250565b600062000c9c6018836200032a565b915062000ca98262000c64565b602082019050919050565b6000602082019050818103600083015262000ccf8162000c8d565b9050919050565b60805161039462000cf1600039600060c601526103946000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063558f208414610046578063714ba40c14610064578063c370b04214610082575b600080fd5b61004e6100a0565b60405161005b91906101f5565b60405180910390f35b61006c6100c4565b6040516100799190610231565b60405180910390f35b61008a6100e8565b60405161009791906102dc565b60405180910390f35b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b600180546100f59061032d565b80601f01602080910402602001604051908101604052809291908181526020018280546101219061032d565b801561016e5780601f106101435761010080835404028352916020019161016e565b820191906000526020600020905b81548152906001019060200180831161015157829003601f168201915b505050505081565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006101bb6101b66101b184610176565b610196565b610176565b9050919050565b60006101cd826101a0565b9050919050565b60006101df826101c2565b9050919050565b6101ef816101d4565b82525050565b600060208201905061020a60008301846101e6565b92915050565b600061021b826101c2565b9050919050565b61022b81610210565b82525050565b60006020820190506102466000830184610222565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561028657808201518184015260208101905061026b565b60008484015250505050565b6000601f19601f8301169050919050565b60006102ae8261024c565b6102b88185610257565b93506102c8818560208601610268565b6102d181610292565b840191505092915050565b600060208201905081810360008301526102f681846102a3565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061034557607f821691505b602082108103610358576103576102fe565b5b5091905056fea264697066735822122007c88ce5919c151ffca8d116bf267464437d6cfc1164eb061221c68f1772b83264736f6c63430008130033",
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
