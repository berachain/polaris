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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"abera\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20Module\",\"outputs\":[{\"internalType\":\"contractIERC20Module\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040526269696973ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801561004657600080fd5b50600060805173ffffffffffffffffffffffffffffffffffffffff1663423eb10b3363075bcd156040518363ffffffff1660e01b815260040161008a9291906103ed565b6020604051808303816000875af11580156100a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100cd9190610475565b90508061010f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610106906104ee565b60405180910390fd5b60805173ffffffffffffffffffffffffffffffffffffffff1663a333e57c6040518163ffffffff1660e01b81526004016101489061050e565b602060405180830381865afa158015610165573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610189919061056c565b6000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600060805173ffffffffffffffffffffffffffffffffffffffff1663cd22a01860008054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b815260040161022591906105ee565b600060405180830381865afa158015610242573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061026b919061076a565b905060405160200161027c906107e1565b60405160208183030381529060405280519060200120816040516020016102a39190610832565b60405160208183030381529060405280519060200120146102f9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f090610895565b60405180910390fd5b50506108b5565b600082825260208201905092915050565b7f6162657261000000000000000000000000000000000000000000000000000000600082015250565b6000610347600583610300565b915061035282610311565b602082019050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006103888261035d565b9050919050565b6103988161037d565b82525050565b6000819050919050565b6000819050919050565b6000819050919050565b60006103d76103d26103cd8461039e565b6103b2565b6103a8565b9050919050565b6103e7816103bc565b82525050565b600060608201905081810360008301526104068161033a565b9050610415602083018561038f565b61042260408301846103de565b9392505050565b6000604051905090565b600080fd5b600080fd5b60008115159050919050565b6104528161043d565b811461045d57600080fd5b50565b60008151905061046f81610449565b92915050565b60006020828403121561048b5761048a610433565b5b600061049984828501610460565b91505092915050565b7f6661696c656420746f20636f6e76657274206162657261000000000000000000600082015250565b60006104d8601783610300565b91506104e3826104a2565b602082019050919050565b60006020820190508181036000830152610507816104cb565b9050919050565b600060208201905081810360008301526105278161033a565b9050919050565b60006105398261037d565b9050919050565b6105498161052e565b811461055457600080fd5b50565b60008151905061056681610540565b92915050565b60006020828403121561058257610581610433565b5b600061059084828501610557565b91505092915050565b60006105b46105af6105aa8461035d565b6103b2565b61035d565b9050919050565b60006105c682610599565b9050919050565b60006105d8826105bb565b9050919050565b6105e8816105cd565b82525050565b600060208201905061060360008301846105df565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61065c82610613565b810181811067ffffffffffffffff8211171561067b5761067a610624565b5b80604052505050565b600061068e610429565b905061069a8282610653565b919050565b600067ffffffffffffffff8211156106ba576106b9610624565b5b6106c382610613565b9050602081019050919050565b60005b838110156106ee5780820151818401526020810190506106d3565b60008484015250505050565b600061070d6107088461069f565b610684565b9050828152602081018484840111156107295761072861060e565b5b6107348482856106d0565b509392505050565b600082601f83011261075157610750610609565b5b81516107618482602086016106fa565b91505092915050565b6000602082840312156107805761077f610433565b5b600082015167ffffffffffffffff81111561079e5761079d610438565b5b6107aa8482850161073c565b91505092915050565b600081905092915050565b60006107cb6005836107b3565b91506107d682610311565b600582019050919050565b60006107ec826107be565b9150819050919050565b600081519050919050565b600061080c826107f6565b61081681856107b3565b93506108268185602086016106d0565b80840191505092915050565b600061083e8284610801565b915081905092915050565b7f72657475726e6564207468652077726f6e672064656e6f6d0000000000000000600082015250565b600061087f601883610300565b915061088a82610849565b602082019050919050565b600060208201905081810360008301526108ae81610872565b9050919050565b6080516101cb6108cf6000396000609d01526101cb6000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063558f20841461003b578063714ba40c14610059575b600080fd5b610043610077565b604051610050919061013e565b60405180910390f35b61006161009b565b60405161006e919061017a565b60405180910390f35b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006101046100ff6100fa846100bf565b6100df565b6100bf565b9050919050565b6000610116826100e9565b9050919050565b60006101288261010b565b9050919050565b6101388161011d565b82525050565b6000602082019050610153600083018461012f565b92915050565b60006101648261010b565b9050919050565b61017481610159565b82525050565b600060208201905061018f600083018461016b565b9291505056fea26469706673582212202a4e3065b6275283dc3a9ee2337177a4050cdc51bee91472779f532230cefba564736f6c63430008130033",
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
