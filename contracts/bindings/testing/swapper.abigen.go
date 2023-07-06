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

// SwapperMetaData contains all meta data concerning the Swapper contract.
var SwapperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20Module\",\"outputs\":[{\"internalType\":\"contractIERC20Module\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"getPolarisERC20\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040526269696973ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff16815250348015610045575f80fd5b50608051610b116100735f395f81816101c2015281816101e7015281816102d101526103b80152610b115ff3fe608060405234801561000f575f80fd5b5060043610610055575f3560e01c806347e7ef2414610059578063714ba40c146100755780639d456b6214610093578063d004f0f7146100af578063d6ece467146100cb575b5f80fd5b610073600480360381019061006e91906104ed565b6100fb565b005b61007d6101c0565b60405161008a9190610586565b60405180910390f35b6100ad60048036038101906100a89190610600565b6101e4565b005b6100c960048036038101906100c49190610698565b6102ce565b005b6100e560048036038101906100e091906106d6565b6103b5565b6040516100f29190610741565b60405180910390f35b5f8273ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b815260040161013993929190610778565b6020604051808303815f875af1158015610155573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061017991906107e2565b9050806101bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101b290610867565b60405180910390fd5b505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663096b406985853333876040518663ffffffff1660e01b81526004016102469594939291906108cf565b6020604051808303815f875af1158015610262573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061028691906107e2565b9050806102c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102bf9061098b565b60405180910390fd5b50505050565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663b96d8bec843333866040518563ffffffff1660e01b815260040161032e94939291906109a9565b6020604051808303815f875af115801561034a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061036e91906107e2565b9050806103b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103a790610a5c565b60405180910390fd5b505050565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a333e57c84846040518363ffffffff1660e01b8152600401610411929190610a7a565b602060405180830381865afa15801561042c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104509190610ab0565b905092915050565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61048982610460565b9050919050565b6104998161047f565b81146104a3575f80fd5b50565b5f813590506104b481610490565b92915050565b5f819050919050565b6104cc816104ba565b81146104d6575f80fd5b50565b5f813590506104e7816104c3565b92915050565b5f806040838503121561050357610502610458565b5b5f610510858286016104a6565b9250506020610521858286016104d9565b9150509250929050565b5f819050919050565b5f61054e61054961054484610460565b61052b565b610460565b9050919050565b5f61055f82610534565b9050919050565b5f61057082610555565b9050919050565b61058081610566565b82525050565b5f6020820190506105995f830184610577565b92915050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f8401126105c0576105bf61059f565b5b8235905067ffffffffffffffff8111156105dd576105dc6105a3565b5b6020830191508360018202830111156105f9576105f86105a7565b5b9250929050565b5f805f6040848603121561061757610616610458565b5b5f84013567ffffffffffffffff8111156106345761063361045c565b5b610640868287016105ab565b93509350506020610653868287016104d9565b9150509250925092565b5f6106678261047f565b9050919050565b6106778161065d565b8114610681575f80fd5b50565b5f813590506106928161066e565b92915050565b5f80604083850312156106ae576106ad610458565b5b5f6106bb85828601610684565b92505060206106cc858286016104d9565b9150509250929050565b5f80602083850312156106ec576106eb610458565b5b5f83013567ffffffffffffffff8111156107095761070861045c565b5b610715858286016105ab565b92509250509250929050565b5f61072b82610555565b9050919050565b61073b81610721565b82525050565b5f6020820190506107545f830184610732565b92915050565b6107638161047f565b82525050565b610772816104ba565b82525050565b5f60608201905061078b5f83018661075a565b610798602083018561075a565b6107a56040830184610769565b949350505050565b5f8115159050919050565b6107c1816107ad565b81146107cb575f80fd5b50565b5f815190506107dc816107b8565b92915050565b5f602082840312156107f7576107f6610458565b5b5f610804848285016107ce565b91505092915050565b5f82825260208201905092915050565b7f537761707065723a207472616e7366657246726f6d206661696c6564000000005f82015250565b5f610851601c8361080d565b915061085c8261081d565b602082019050919050565b5f6020820190508181035f83015261087e81610845565b9050919050565b828183375f83830152505050565b5f601f19601f8301169050919050565b5f6108ae838561080d565b93506108bb838584610885565b6108c483610893565b840190509392505050565b5f6080820190508181035f8301526108e88187896108a3565b90506108f7602083018661075a565b610904604083018561075a565b6109116060830184610769565b9695505050505050565b7f537761707065723a207472616e73666572436f696e546f4552433230206661695f8201527f6c65640000000000000000000000000000000000000000000000000000000000602082015250565b5f61097560238361080d565b91506109808261091b565b604082019050919050565b5f6020820190508181035f8301526109a281610969565b9050919050565b5f6080820190506109bc5f830187610732565b6109c9602083018661075a565b6109d6604083018561075a565b6109e36060830184610769565b95945050505050565b7f537761707065723a207472616e736665724552433230546f436f696e206661695f8201527f6c65640000000000000000000000000000000000000000000000000000000000602082015250565b5f610a4660238361080d565b9150610a51826109ec565b604082019050919050565b5f6020820190508181035f830152610a7381610a3a565b9050919050565b5f6020820190508181035f830152610a938184866108a3565b90509392505050565b5f81519050610aaa8161066e565b92915050565b5f60208284031215610ac557610ac4610458565b5b5f610ad284828501610a9c565b9150509291505056fea264697066735822122049ffe89211407e46386baaa6daacd01f1bcdd9ef2d60eaa15051e9758d4e951164736f6c63430008140033",
}

// SwapperABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapperMetaData.ABI instead.
var SwapperABI = SwapperMetaData.ABI

// SwapperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SwapperMetaData.Bin instead.
var SwapperBin = SwapperMetaData.Bin

// DeploySwapper deploys a new Ethereum contract, binding an instance of Swapper to it.
func DeploySwapper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Swapper, error) {
	parsed, err := SwapperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SwapperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Swapper{SwapperCaller: SwapperCaller{contract: contract}, SwapperTransactor: SwapperTransactor{contract: contract}, SwapperFilterer: SwapperFilterer{contract: contract}}, nil
}

// Swapper is an auto generated Go binding around an Ethereum contract.
type Swapper struct {
	SwapperCaller     // Read-only binding to the contract
	SwapperTransactor // Write-only binding to the contract
	SwapperFilterer   // Log filterer for contract events
}

// SwapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapperSession struct {
	Contract     *Swapper          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapperCallerSession struct {
	Contract *SwapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SwapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapperTransactorSession struct {
	Contract     *SwapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SwapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapperRaw struct {
	Contract *Swapper // Generic contract binding to access the raw methods on
}

// SwapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapperCallerRaw struct {
	Contract *SwapperCaller // Generic read-only contract binding to access the raw methods on
}

// SwapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapperTransactorRaw struct {
	Contract *SwapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapper creates a new instance of Swapper, bound to a specific deployed contract.
func NewSwapper(address common.Address, backend bind.ContractBackend) (*Swapper, error) {
	contract, err := bindSwapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swapper{SwapperCaller: SwapperCaller{contract: contract}, SwapperTransactor: SwapperTransactor{contract: contract}, SwapperFilterer: SwapperFilterer{contract: contract}}, nil
}

// NewSwapperCaller creates a new read-only instance of Swapper, bound to a specific deployed contract.
func NewSwapperCaller(address common.Address, caller bind.ContractCaller) (*SwapperCaller, error) {
	contract, err := bindSwapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapperCaller{contract: contract}, nil
}

// NewSwapperTransactor creates a new write-only instance of Swapper, bound to a specific deployed contract.
func NewSwapperTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapperTransactor, error) {
	contract, err := bindSwapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapperTransactor{contract: contract}, nil
}

// NewSwapperFilterer creates a new log filterer instance of Swapper, bound to a specific deployed contract.
func NewSwapperFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapperFilterer, error) {
	contract, err := bindSwapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapperFilterer{contract: contract}, nil
}

// bindSwapper binds a generic wrapper to an already deployed contract.
func bindSwapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SwapperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swapper *SwapperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swapper.Contract.SwapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swapper *SwapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapper.Contract.SwapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swapper *SwapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swapper.Contract.SwapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swapper *SwapperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swapper *SwapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swapper *SwapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swapper.Contract.contract.Transact(opts, method, params...)
}

// Erc20Module is a free data retrieval call binding the contract method 0x714ba40c.
//
// Solidity: function erc20Module() view returns(address)
func (_Swapper *SwapperCaller) Erc20Module(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "erc20Module")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20Module is a free data retrieval call binding the contract method 0x714ba40c.
//
// Solidity: function erc20Module() view returns(address)
func (_Swapper *SwapperSession) Erc20Module() (common.Address, error) {
	return _Swapper.Contract.Erc20Module(&_Swapper.CallOpts)
}

// Erc20Module is a free data retrieval call binding the contract method 0x714ba40c.
//
// Solidity: function erc20Module() view returns(address)
func (_Swapper *SwapperCallerSession) Erc20Module() (common.Address, error) {
	return _Swapper.Contract.Erc20Module(&_Swapper.CallOpts)
}

// GetPolarisERC20 is a free data retrieval call binding the contract method 0xd6ece467.
//
// Solidity: function getPolarisERC20(string denom) view returns(address)
func (_Swapper *SwapperCaller) GetPolarisERC20(opts *bind.CallOpts, denom string) (common.Address, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "getPolarisERC20", denom)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPolarisERC20 is a free data retrieval call binding the contract method 0xd6ece467.
//
// Solidity: function getPolarisERC20(string denom) view returns(address)
func (_Swapper *SwapperSession) GetPolarisERC20(denom string) (common.Address, error) {
	return _Swapper.Contract.GetPolarisERC20(&_Swapper.CallOpts, denom)
}

// GetPolarisERC20 is a free data retrieval call binding the contract method 0xd6ece467.
//
// Solidity: function getPolarisERC20(string denom) view returns(address)
func (_Swapper *SwapperCallerSession) GetPolarisERC20(denom string) (common.Address, error) {
	return _Swapper.Contract.GetPolarisERC20(&_Swapper.CallOpts, denom)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_Swapper *SwapperTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_Swapper *SwapperSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.Contract.Deposit(&_Swapper.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_Swapper *SwapperTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.Contract.Deposit(&_Swapper.TransactOpts, token, amount)
}

// Swap is a paid mutator transaction binding the contract method 0x9d456b62.
//
// Solidity: function swap(string denom, uint256 amount) returns()
func (_Swapper *SwapperTransactor) Swap(opts *bind.TransactOpts, denom string, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "swap", denom, amount)
}

// Swap is a paid mutator transaction binding the contract method 0x9d456b62.
//
// Solidity: function swap(string denom, uint256 amount) returns()
func (_Swapper *SwapperSession) Swap(denom string, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.Contract.Swap(&_Swapper.TransactOpts, denom, amount)
}

// Swap is a paid mutator transaction binding the contract method 0x9d456b62.
//
// Solidity: function swap(string denom, uint256 amount) returns()
func (_Swapper *SwapperTransactorSession) Swap(denom string, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.Contract.Swap(&_Swapper.TransactOpts, denom, amount)
}

// Swap0 is a paid mutator transaction binding the contract method 0xd004f0f7.
//
// Solidity: function swap(address token, uint256 amount) returns()
func (_Swapper *SwapperTransactor) Swap0(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "swap0", token, amount)
}

// Swap0 is a paid mutator transaction binding the contract method 0xd004f0f7.
//
// Solidity: function swap(address token, uint256 amount) returns()
func (_Swapper *SwapperSession) Swap0(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.Contract.Swap0(&_Swapper.TransactOpts, token, amount)
}

// Swap0 is a paid mutator transaction binding the contract method 0xd004f0f7.
//
// Solidity: function swap(address token, uint256 amount) returns()
func (_Swapper *SwapperTransactorSession) Swap0(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.Contract.Swap0(&_Swapper.TransactOpts, token, amount)
}
