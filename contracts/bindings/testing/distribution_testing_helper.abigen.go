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

// DistributionTestHelperMetaData contains all meta data concerning the DistributionTestHelper contract.
var DistributionTestHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_distributionprecompile\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"distribution\",\"outputs\":[{\"internalType\":\"contractIDistributionModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWithdrawEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_withdrawAddress\",\"type\":\"address\"}],\"name\":\"setWithdrawAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"}],\"name\":\"withdrawRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610a15380380610a1583398181016040528101906100329190610141565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610098576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505061016e565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061010e826100e3565b9050919050565b61011e81610103565b811461012957600080fd5b50565b60008151905061013b81610115565b92915050565b600060208284031215610157576101566100de565b5b60006101658482850161012c565b91505092915050565b6108988061017d6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806339cc4c86146100515780633ab1a4941461006f5780635ee58efc1461009f578063e20981ca146100bd575b600080fd5b6100596100d9565b60405161006691906102fc565b60405180910390f35b61008960048036038101906100849190610389565b610170565b60405161009691906102fc565b60405180910390f35b6100a7610215565b6040516100b49190610415565b60405180910390f35b6100d760048036038101906100d29190610430565b610239565b005b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166339cc4c866040518163ffffffff1660e01b8152600401602060405180830381865afa158015610147573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061016b919061049c565b905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633ab1a494836040518263ffffffff1660e01b81526004016101cc91906104d8565b6020604051808303816000875af11580156101eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061020f919061049c565b50919050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663562c67a483836040518363ffffffff1660e01b81526004016102949291906104f3565b6000604051808303816000875af11580156102b3573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906102dc9190610819565b505050565b60008115159050919050565b6102f6816102e1565b82525050565b600060208201905061031160008301846102ed565b92915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006103568261032b565b9050919050565b6103668161034b565b811461037157600080fd5b50565b6000813590506103838161035d565b92915050565b60006020828403121561039f5761039e610321565b5b60006103ad84828501610374565b91505092915050565b6000819050919050565b60006103db6103d66103d18461032b565b6103b6565b61032b565b9050919050565b60006103ed826103c0565b9050919050565b60006103ff826103e2565b9050919050565b61040f816103f4565b82525050565b600060208201905061042a6000830184610406565b92915050565b6000806040838503121561044757610446610321565b5b600061045585828601610374565b925050602061046685828601610374565b9150509250929050565b610479816102e1565b811461048457600080fd5b50565b60008151905061049681610470565b92915050565b6000602082840312156104b2576104b1610321565b5b60006104c084828501610487565b91505092915050565b6104d28161034b565b82525050565b60006020820190506104ed60008301846104c9565b92915050565b600060408201905061050860008301856104c9565b61051560208301846104c9565b9392505050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61056a82610521565b810181811067ffffffffffffffff8211171561058957610588610532565b5b80604052505050565b600061059c610317565b90506105a88282610561565b919050565b600067ffffffffffffffff8211156105c8576105c7610532565b5b602082029050602081019050919050565b600080fd5b600080fd5b600080fd5b600067ffffffffffffffff82169050919050565b610605816105e8565b811461061057600080fd5b50565b600081519050610622816105fc565b92915050565b600080fd5b600067ffffffffffffffff82111561064857610647610532565b5b61065182610521565b9050602081019050919050565b60005b8381101561067c578082015181840152602081019050610661565b60008484015250505050565b600061069b6106968461062d565b610592565b9050828152602081018484840111156106b7576106b6610628565b5b6106c284828561065e565b509392505050565b600082601f8301126106df576106de61051c565b5b81516106ef848260208601610688565b91505092915050565b60006040828403121561070e5761070d6105de565b5b6107186040610592565b9050600061072884828501610613565b600083015250602082015167ffffffffffffffff81111561074c5761074b6105e3565b5b610758848285016106ca565b60208301525092915050565b6000610777610772846105ad565b610592565b9050808382526020820190506020840283018581111561079a576107996105d9565b5b835b818110156107e157805167ffffffffffffffff8111156107bf576107be61051c565b5b8086016107cc89826106f8565b8552602085019450505060208101905061079c565b5050509392505050565b600082601f830112610800576107ff61051c565b5b8151610810848260208601610764565b91505092915050565b60006020828403121561082f5761082e610321565b5b600082015167ffffffffffffffff81111561084d5761084c610326565b5b610859848285016107eb565b9150509291505056fea26469706673582212209bdfdbb7f7d5dfd06bfef1c80dad5173f3241d2574619720ce33c118dfc0862564736f6c63430008130033",
}

// DistributionTestHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use DistributionTestHelperMetaData.ABI instead.
var DistributionTestHelperABI = DistributionTestHelperMetaData.ABI

// DistributionTestHelperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DistributionTestHelperMetaData.Bin instead.
var DistributionTestHelperBin = DistributionTestHelperMetaData.Bin

// DeployDistributionTestHelper deploys a new Ethereum contract, binding an instance of DistributionTestHelper to it.
func DeployDistributionTestHelper(auth *bind.TransactOpts, backend bind.ContractBackend, _distributionprecompile common.Address) (common.Address, *types.Transaction, *DistributionTestHelper, error) {
	parsed, err := DistributionTestHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DistributionTestHelperBin), backend, _distributionprecompile)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DistributionTestHelper{DistributionTestHelperCaller: DistributionTestHelperCaller{contract: contract}, DistributionTestHelperTransactor: DistributionTestHelperTransactor{contract: contract}, DistributionTestHelperFilterer: DistributionTestHelperFilterer{contract: contract}}, nil
}

// DistributionTestHelper is an auto generated Go binding around an Ethereum contract.
type DistributionTestHelper struct {
	DistributionTestHelperCaller     // Read-only binding to the contract
	DistributionTestHelperTransactor // Write-only binding to the contract
	DistributionTestHelperFilterer   // Log filterer for contract events
}

// DistributionTestHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type DistributionTestHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributionTestHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DistributionTestHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributionTestHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DistributionTestHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributionTestHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DistributionTestHelperSession struct {
	Contract     *DistributionTestHelper // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DistributionTestHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DistributionTestHelperCallerSession struct {
	Contract *DistributionTestHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// DistributionTestHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DistributionTestHelperTransactorSession struct {
	Contract     *DistributionTestHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// DistributionTestHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type DistributionTestHelperRaw struct {
	Contract *DistributionTestHelper // Generic contract binding to access the raw methods on
}

// DistributionTestHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DistributionTestHelperCallerRaw struct {
	Contract *DistributionTestHelperCaller // Generic read-only contract binding to access the raw methods on
}

// DistributionTestHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DistributionTestHelperTransactorRaw struct {
	Contract *DistributionTestHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDistributionTestHelper creates a new instance of DistributionTestHelper, bound to a specific deployed contract.
func NewDistributionTestHelper(address common.Address, backend bind.ContractBackend) (*DistributionTestHelper, error) {
	contract, err := bindDistributionTestHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DistributionTestHelper{DistributionTestHelperCaller: DistributionTestHelperCaller{contract: contract}, DistributionTestHelperTransactor: DistributionTestHelperTransactor{contract: contract}, DistributionTestHelperFilterer: DistributionTestHelperFilterer{contract: contract}}, nil
}

// NewDistributionTestHelperCaller creates a new read-only instance of DistributionTestHelper, bound to a specific deployed contract.
func NewDistributionTestHelperCaller(address common.Address, caller bind.ContractCaller) (*DistributionTestHelperCaller, error) {
	contract, err := bindDistributionTestHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DistributionTestHelperCaller{contract: contract}, nil
}

// NewDistributionTestHelperTransactor creates a new write-only instance of DistributionTestHelper, bound to a specific deployed contract.
func NewDistributionTestHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*DistributionTestHelperTransactor, error) {
	contract, err := bindDistributionTestHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DistributionTestHelperTransactor{contract: contract}, nil
}

// NewDistributionTestHelperFilterer creates a new log filterer instance of DistributionTestHelper, bound to a specific deployed contract.
func NewDistributionTestHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*DistributionTestHelperFilterer, error) {
	contract, err := bindDistributionTestHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DistributionTestHelperFilterer{contract: contract}, nil
}

// bindDistributionTestHelper binds a generic wrapper to an already deployed contract.
func bindDistributionTestHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DistributionTestHelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DistributionTestHelper *DistributionTestHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DistributionTestHelper.Contract.DistributionTestHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DistributionTestHelper *DistributionTestHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DistributionTestHelper.Contract.DistributionTestHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DistributionTestHelper *DistributionTestHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DistributionTestHelper.Contract.DistributionTestHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DistributionTestHelper *DistributionTestHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DistributionTestHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DistributionTestHelper *DistributionTestHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DistributionTestHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DistributionTestHelper *DistributionTestHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DistributionTestHelper.Contract.contract.Transact(opts, method, params...)
}

// Distribution is a free data retrieval call binding the contract method 0x5ee58efc.
//
// Solidity: function distribution() view returns(address)
func (_DistributionTestHelper *DistributionTestHelperCaller) Distribution(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DistributionTestHelper.contract.Call(opts, &out, "distribution")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Distribution is a free data retrieval call binding the contract method 0x5ee58efc.
//
// Solidity: function distribution() view returns(address)
func (_DistributionTestHelper *DistributionTestHelperSession) Distribution() (common.Address, error) {
	return _DistributionTestHelper.Contract.Distribution(&_DistributionTestHelper.CallOpts)
}

// Distribution is a free data retrieval call binding the contract method 0x5ee58efc.
//
// Solidity: function distribution() view returns(address)
func (_DistributionTestHelper *DistributionTestHelperCallerSession) Distribution() (common.Address, error) {
	return _DistributionTestHelper.Contract.Distribution(&_DistributionTestHelper.CallOpts)
}

// GetWithdrawEnabled is a free data retrieval call binding the contract method 0x39cc4c86.
//
// Solidity: function getWithdrawEnabled() view returns(bool)
func (_DistributionTestHelper *DistributionTestHelperCaller) GetWithdrawEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DistributionTestHelper.contract.Call(opts, &out, "getWithdrawEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetWithdrawEnabled is a free data retrieval call binding the contract method 0x39cc4c86.
//
// Solidity: function getWithdrawEnabled() view returns(bool)
func (_DistributionTestHelper *DistributionTestHelperSession) GetWithdrawEnabled() (bool, error) {
	return _DistributionTestHelper.Contract.GetWithdrawEnabled(&_DistributionTestHelper.CallOpts)
}

// GetWithdrawEnabled is a free data retrieval call binding the contract method 0x39cc4c86.
//
// Solidity: function getWithdrawEnabled() view returns(bool)
func (_DistributionTestHelper *DistributionTestHelperCallerSession) GetWithdrawEnabled() (bool, error) {
	return _DistributionTestHelper.Contract.GetWithdrawEnabled(&_DistributionTestHelper.CallOpts)
}

// SetWithdrawAddress is a paid mutator transaction binding the contract method 0x3ab1a494.
//
// Solidity: function setWithdrawAddress(address _withdrawAddress) returns(bool)
func (_DistributionTestHelper *DistributionTestHelperTransactor) SetWithdrawAddress(opts *bind.TransactOpts, _withdrawAddress common.Address) (*types.Transaction, error) {
	return _DistributionTestHelper.contract.Transact(opts, "setWithdrawAddress", _withdrawAddress)
}

// SetWithdrawAddress is a paid mutator transaction binding the contract method 0x3ab1a494.
//
// Solidity: function setWithdrawAddress(address _withdrawAddress) returns(bool)
func (_DistributionTestHelper *DistributionTestHelperSession) SetWithdrawAddress(_withdrawAddress common.Address) (*types.Transaction, error) {
	return _DistributionTestHelper.Contract.SetWithdrawAddress(&_DistributionTestHelper.TransactOpts, _withdrawAddress)
}

// SetWithdrawAddress is a paid mutator transaction binding the contract method 0x3ab1a494.
//
// Solidity: function setWithdrawAddress(address _withdrawAddress) returns(bool)
func (_DistributionTestHelper *DistributionTestHelperTransactorSession) SetWithdrawAddress(_withdrawAddress common.Address) (*types.Transaction, error) {
	return _DistributionTestHelper.Contract.SetWithdrawAddress(&_DistributionTestHelper.TransactOpts, _withdrawAddress)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xe20981ca.
//
// Solidity: function withdrawRewards(address _delegatorAddress, address _validatorAddress) returns()
func (_DistributionTestHelper *DistributionTestHelperTransactor) WithdrawRewards(opts *bind.TransactOpts, _delegatorAddress common.Address, _validatorAddress common.Address) (*types.Transaction, error) {
	return _DistributionTestHelper.contract.Transact(opts, "withdrawRewards", _delegatorAddress, _validatorAddress)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xe20981ca.
//
// Solidity: function withdrawRewards(address _delegatorAddress, address _validatorAddress) returns()
func (_DistributionTestHelper *DistributionTestHelperSession) WithdrawRewards(_delegatorAddress common.Address, _validatorAddress common.Address) (*types.Transaction, error) {
	return _DistributionTestHelper.Contract.WithdrawRewards(&_DistributionTestHelper.TransactOpts, _delegatorAddress, _validatorAddress)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xe20981ca.
//
// Solidity: function withdrawRewards(address _delegatorAddress, address _validatorAddress) returns()
func (_DistributionTestHelper *DistributionTestHelperTransactorSession) WithdrawRewards(_delegatorAddress common.Address, _validatorAddress common.Address) (*types.Transaction, error) {
	return _DistributionTestHelper.Contract.WithdrawRewards(&_DistributionTestHelper.TransactOpts, _delegatorAddress, _validatorAddress)
}
