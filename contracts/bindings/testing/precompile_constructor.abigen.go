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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"erc20Module\",\"outputs\":[{\"internalType\":\"contractIERC20Module\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040526269696973ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801561004657600080fd5b50600060805173ffffffffffffffffffffffffffffffffffffffff1663423eb10b3363075bcd156040518363ffffffff1660e01b815260040161008a929190610393565b6020604051808303816000875af11580156100a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100cd919061041b565b90508061010f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161010690610494565b60405180910390fd5b600060805173ffffffffffffffffffffffffffffffffffffffff1663a333e57c6040518163ffffffff1660e01b815260040161014a906104b4565b602060405180830381865afa158015610167573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061018b9190610512565b9050600060805173ffffffffffffffffffffffffffffffffffffffff1663cd22a018836040518263ffffffff1660e01b81526004016101ca9190610594565b600060405180830381865afa1580156101e7573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906102109190610710565b905060405160200161022190610787565b604051602081830303815290604052805190602001208160405160200161024891906107d8565b604051602081830303815290604052805190602001201461029e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102959061083b565b60405180910390fd5b50505061085b565b600082825260208201905092915050565b7f6162657261000000000000000000000000000000000000000000000000000000600082015250565b60006102ed6005836102a6565b91506102f8826102b7565b602082019050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061032e82610303565b9050919050565b61033e81610323565b82525050565b6000819050919050565b6000819050919050565b6000819050919050565b600061037d61037861037384610344565b610358565b61034e565b9050919050565b61038d81610362565b82525050565b600060608201905081810360008301526103ac816102e0565b90506103bb6020830185610335565b6103c86040830184610384565b9392505050565b6000604051905090565b600080fd5b600080fd5b60008115159050919050565b6103f8816103e3565b811461040357600080fd5b50565b600081519050610415816103ef565b92915050565b600060208284031215610431576104306103d9565b5b600061043f84828501610406565b91505092915050565b7f6661696c656420746f20636f6e76657274206162657261000000000000000000600082015250565b600061047e6017836102a6565b915061048982610448565b602082019050919050565b600060208201905081810360008301526104ad81610471565b9050919050565b600060208201905081810360008301526104cd816102e0565b9050919050565b60006104df82610323565b9050919050565b6104ef816104d4565b81146104fa57600080fd5b50565b60008151905061050c816104e6565b92915050565b600060208284031215610528576105276103d9565b5b6000610536848285016104fd565b91505092915050565b600061055a61055561055084610303565b610358565b610303565b9050919050565b600061056c8261053f565b9050919050565b600061057e82610561565b9050919050565b61058e81610573565b82525050565b60006020820190506105a96000830184610585565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610602826105b9565b810181811067ffffffffffffffff82111715610621576106206105ca565b5b80604052505050565b60006106346103cf565b905061064082826105f9565b919050565b600067ffffffffffffffff8211156106605761065f6105ca565b5b610669826105b9565b9050602081019050919050565b60005b83811015610694578082015181840152602081019050610679565b60008484015250505050565b60006106b36106ae84610645565b61062a565b9050828152602081018484840111156106cf576106ce6105b4565b5b6106da848285610676565b509392505050565b600082601f8301126106f7576106f66105af565b5b81516107078482602086016106a0565b91505092915050565b600060208284031215610726576107256103d9565b5b600082015167ffffffffffffffff811115610744576107436103de565b5b610750848285016106e2565b91505092915050565b600081905092915050565b6000610771600583610759565b915061077c826102b7565b600582019050919050565b600061079282610764565b9150819050919050565b600081519050919050565b60006107b28261079c565b6107bc8185610759565b93506107cc818560208601610676565b80840191505092915050565b60006107e482846107a7565b915081905092915050565b7f72657475726e6564207468652077726f6e672064656e6f6d0000000000000000600082015250565b60006108256018836102a6565b9150610830826107ef565b602082019050919050565b6000602082019050818103600083015261085481610818565b9050919050565b60805161012d61087560003960006049015261012d6000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063714ba40c14602d575b600080fd5b60336047565b604051603e919060de565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000081565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060aa60a660a284606b565b608b565b606b565b9050919050565b600060ba826095565b9050919050565b600060ca8260b1565b9050919050565b60d88160c1565b82525050565b600060208201905060f1600083018460d1565b9291505056fea2646970667358221220b1a9058a88107b3945ef7df11d800909553dda575686bea22a558ba771ed0cf764736f6c63430008130033",
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
