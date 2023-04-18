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

// DistributionWrapperMetaData contains all meta data concerning the DistributionWrapper contract.
var DistributionWrapperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_distributionprecompile\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingprecompile\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distribution\",\"outputs\":[{\"internalType\":\"contractIDistributionModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWithdrawEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_withdrawAddress\",\"type\":\"address\"}],\"name\":\"setWithdrawAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staking\",\"outputs\":[{\"internalType\":\"contractIStakingModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"}],\"name\":\"withdrawRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610c5c380380610c5c833981810160405281019061003291906101bc565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614801561009a5750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b156100d1576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050506101fc565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006101898261015e565b9050919050565b6101998161017e565b81146101a457600080fd5b50565b6000815190506101b681610190565b92915050565b600080604083850312156101d3576101d2610159565b5b60006101e1858286016101a7565b92505060206101f2858286016101a7565b9150509250929050565b610a518061020b6000396000f3fe6080604052600436106100555760003560e01c806339cc4c861461005a5780633ab1a494146100855780634cf088d9146100c25780635c19a95c146100ed5780635ee58efc14610109578063e20981ca14610134575b600080fd5b34801561006657600080fd5b5061006f61015d565b60405161007c919061044b565b60405180910390f35b34801561009157600080fd5b506100ac60048036038101906100a791906104d8565b6101f4565b6040516100b9919061044b565b60405180910390f35b3480156100ce57600080fd5b506100d761029a565b6040516100e49190610564565b60405180910390f35b610107600480360381019061010291906104d8565b6102c0565b005b34801561011557600080fd5b5061011e610364565b60405161012b91906105a0565b60405180910390f35b34801561014057600080fd5b5061015b600480360381019061015691906105bb565b610388565b005b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166339cc4c866040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101cb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101ef9190610627565b905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633ab1a494836040518263ffffffff1660e01b81526004016102509190610663565b6020604051808303816000875af115801561026f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102939190610627565b9050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663026e402b82346040518363ffffffff1660e01b815260040161031d929190610697565b6020604051808303816000875af115801561033c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103609190610627565b5050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663562c67a483836040518363ffffffff1660e01b81526004016103e39291906106c0565b6000604051808303816000875af1158015610402573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061042b91906109d2565b505050565b60008115159050919050565b61044581610430565b82525050565b6000602082019050610460600083018461043c565b92915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006104a58261047a565b9050919050565b6104b58161049a565b81146104c057600080fd5b50565b6000813590506104d2816104ac565b92915050565b6000602082840312156104ee576104ed610470565b5b60006104fc848285016104c3565b91505092915050565b6000819050919050565b600061052a6105256105208461047a565b610505565b61047a565b9050919050565b600061053c8261050f565b9050919050565b600061054e82610531565b9050919050565b61055e81610543565b82525050565b60006020820190506105796000830184610555565b92915050565b600061058a82610531565b9050919050565b61059a8161057f565b82525050565b60006020820190506105b56000830184610591565b92915050565b600080604083850312156105d2576105d1610470565b5b60006105e0858286016104c3565b92505060206105f1858286016104c3565b9150509250929050565b61060481610430565b811461060f57600080fd5b50565b600081519050610621816105fb565b92915050565b60006020828403121561063d5761063c610470565b5b600061064b84828501610612565b91505092915050565b61065d8161049a565b82525050565b60006020820190506106786000830184610654565b92915050565b6000819050919050565b6106918161067e565b82525050565b60006040820190506106ac6000830185610654565b6106b96020830184610688565b9392505050565b60006040820190506106d56000830185610654565b6106e26020830184610654565b9392505050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610737826106ee565b810181811067ffffffffffffffff82111715610756576107556106ff565b5b80604052505050565b6000610769610466565b9050610775828261072e565b919050565b600067ffffffffffffffff821115610795576107946106ff565b5b602082029050602081019050919050565b600080fd5b600080fd5b600080fd5b6107be8161067e565b81146107c957600080fd5b50565b6000815190506107db816107b5565b92915050565b600080fd5b600067ffffffffffffffff821115610801576108006106ff565b5b61080a826106ee565b9050602081019050919050565b60005b8381101561083557808201518184015260208101905061081a565b60008484015250505050565b600061085461084f846107e6565b61075f565b9050828152602081018484840111156108705761086f6107e1565b5b61087b848285610817565b509392505050565b600082601f830112610898576108976106e9565b5b81516108a8848260208601610841565b91505092915050565b6000604082840312156108c7576108c66107ab565b5b6108d1604061075f565b905060006108e1848285016107cc565b600083015250602082015167ffffffffffffffff811115610905576109046107b0565b5b61091184828501610883565b60208301525092915050565b600061093061092b8461077a565b61075f565b90508083825260208201905060208402830185811115610953576109526107a6565b5b835b8181101561099a57805167ffffffffffffffff811115610978576109776106e9565b5b80860161098589826108b1565b85526020850194505050602081019050610955565b5050509392505050565b600082601f8301126109b9576109b86106e9565b5b81516109c984826020860161091d565b91505092915050565b6000602082840312156109e8576109e7610470565b5b600082015167ffffffffffffffff811115610a0657610a05610475565b5b610a12848285016109a4565b9150509291505056fea2646970667358221220a90c297c19bcdee914cb4565b17354c3474c36d7b82f0a37292847037ca3136f64736f6c63430008130033",
}

// DistributionWrapperABI is the input ABI used to generate the binding from.
// Deprecated: Use DistributionWrapperMetaData.ABI instead.
var DistributionWrapperABI = DistributionWrapperMetaData.ABI

// DistributionWrapperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DistributionWrapperMetaData.Bin instead.
var DistributionWrapperBin = DistributionWrapperMetaData.Bin

// DeployDistributionWrapper deploys a new Ethereum contract, binding an instance of DistributionWrapper to it.
func DeployDistributionWrapper(auth *bind.TransactOpts, backend bind.ContractBackend, _distributionprecompile common.Address, _stakingprecompile common.Address) (common.Address, *types.Transaction, *DistributionWrapper, error) {
	parsed, err := DistributionWrapperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DistributionWrapperBin), backend, _distributionprecompile, _stakingprecompile)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DistributionWrapper{DistributionWrapperCaller: DistributionWrapperCaller{contract: contract}, DistributionWrapperTransactor: DistributionWrapperTransactor{contract: contract}, DistributionWrapperFilterer: DistributionWrapperFilterer{contract: contract}}, nil
}

// DistributionWrapper is an auto generated Go binding around an Ethereum contract.
type DistributionWrapper struct {
	DistributionWrapperCaller     // Read-only binding to the contract
	DistributionWrapperTransactor // Write-only binding to the contract
	DistributionWrapperFilterer   // Log filterer for contract events
}

// DistributionWrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type DistributionWrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributionWrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DistributionWrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributionWrapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DistributionWrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributionWrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DistributionWrapperSession struct {
	Contract     *DistributionWrapper // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DistributionWrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DistributionWrapperCallerSession struct {
	Contract *DistributionWrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// DistributionWrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DistributionWrapperTransactorSession struct {
	Contract     *DistributionWrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// DistributionWrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type DistributionWrapperRaw struct {
	Contract *DistributionWrapper // Generic contract binding to access the raw methods on
}

// DistributionWrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DistributionWrapperCallerRaw struct {
	Contract *DistributionWrapperCaller // Generic read-only contract binding to access the raw methods on
}

// DistributionWrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DistributionWrapperTransactorRaw struct {
	Contract *DistributionWrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDistributionWrapper creates a new instance of DistributionWrapper, bound to a specific deployed contract.
func NewDistributionWrapper(address common.Address, backend bind.ContractBackend) (*DistributionWrapper, error) {
	contract, err := bindDistributionWrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DistributionWrapper{DistributionWrapperCaller: DistributionWrapperCaller{contract: contract}, DistributionWrapperTransactor: DistributionWrapperTransactor{contract: contract}, DistributionWrapperFilterer: DistributionWrapperFilterer{contract: contract}}, nil
}

// NewDistributionWrapperCaller creates a new read-only instance of DistributionWrapper, bound to a specific deployed contract.
func NewDistributionWrapperCaller(address common.Address, caller bind.ContractCaller) (*DistributionWrapperCaller, error) {
	contract, err := bindDistributionWrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DistributionWrapperCaller{contract: contract}, nil
}

// NewDistributionWrapperTransactor creates a new write-only instance of DistributionWrapper, bound to a specific deployed contract.
func NewDistributionWrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*DistributionWrapperTransactor, error) {
	contract, err := bindDistributionWrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DistributionWrapperTransactor{contract: contract}, nil
}

// NewDistributionWrapperFilterer creates a new log filterer instance of DistributionWrapper, bound to a specific deployed contract.
func NewDistributionWrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*DistributionWrapperFilterer, error) {
	contract, err := bindDistributionWrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DistributionWrapperFilterer{contract: contract}, nil
}

// bindDistributionWrapper binds a generic wrapper to an already deployed contract.
func bindDistributionWrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DistributionWrapperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DistributionWrapper *DistributionWrapperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DistributionWrapper.Contract.DistributionWrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DistributionWrapper *DistributionWrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DistributionWrapper.Contract.DistributionWrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DistributionWrapper *DistributionWrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DistributionWrapper.Contract.DistributionWrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DistributionWrapper *DistributionWrapperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DistributionWrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DistributionWrapper *DistributionWrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DistributionWrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DistributionWrapper *DistributionWrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DistributionWrapper.Contract.contract.Transact(opts, method, params...)
}

// Distribution is a free data retrieval call binding the contract method 0x5ee58efc.
//
// Solidity: function distribution() view returns(address)
func (_DistributionWrapper *DistributionWrapperCaller) Distribution(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DistributionWrapper.contract.Call(opts, &out, "distribution")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Distribution is a free data retrieval call binding the contract method 0x5ee58efc.
//
// Solidity: function distribution() view returns(address)
func (_DistributionWrapper *DistributionWrapperSession) Distribution() (common.Address, error) {
	return _DistributionWrapper.Contract.Distribution(&_DistributionWrapper.CallOpts)
}

// Distribution is a free data retrieval call binding the contract method 0x5ee58efc.
//
// Solidity: function distribution() view returns(address)
func (_DistributionWrapper *DistributionWrapperCallerSession) Distribution() (common.Address, error) {
	return _DistributionWrapper.Contract.Distribution(&_DistributionWrapper.CallOpts)
}

// GetWithdrawEnabled is a free data retrieval call binding the contract method 0x39cc4c86.
//
// Solidity: function getWithdrawEnabled() view returns(bool)
func (_DistributionWrapper *DistributionWrapperCaller) GetWithdrawEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DistributionWrapper.contract.Call(opts, &out, "getWithdrawEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetWithdrawEnabled is a free data retrieval call binding the contract method 0x39cc4c86.
//
// Solidity: function getWithdrawEnabled() view returns(bool)
func (_DistributionWrapper *DistributionWrapperSession) GetWithdrawEnabled() (bool, error) {
	return _DistributionWrapper.Contract.GetWithdrawEnabled(&_DistributionWrapper.CallOpts)
}

// GetWithdrawEnabled is a free data retrieval call binding the contract method 0x39cc4c86.
//
// Solidity: function getWithdrawEnabled() view returns(bool)
func (_DistributionWrapper *DistributionWrapperCallerSession) GetWithdrawEnabled() (bool, error) {
	return _DistributionWrapper.Contract.GetWithdrawEnabled(&_DistributionWrapper.CallOpts)
}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_DistributionWrapper *DistributionWrapperCaller) Staking(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DistributionWrapper.contract.Call(opts, &out, "staking")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_DistributionWrapper *DistributionWrapperSession) Staking() (common.Address, error) {
	return _DistributionWrapper.Contract.Staking(&_DistributionWrapper.CallOpts)
}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_DistributionWrapper *DistributionWrapperCallerSession) Staking() (common.Address, error) {
	return _DistributionWrapper.Contract.Staking(&_DistributionWrapper.CallOpts)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address _validator) payable returns()
func (_DistributionWrapper *DistributionWrapperTransactor) Delegate(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _DistributionWrapper.contract.Transact(opts, "delegate", _validator)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address _validator) payable returns()
func (_DistributionWrapper *DistributionWrapperSession) Delegate(_validator common.Address) (*types.Transaction, error) {
	return _DistributionWrapper.Contract.Delegate(&_DistributionWrapper.TransactOpts, _validator)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address _validator) payable returns()
func (_DistributionWrapper *DistributionWrapperTransactorSession) Delegate(_validator common.Address) (*types.Transaction, error) {
	return _DistributionWrapper.Contract.Delegate(&_DistributionWrapper.TransactOpts, _validator)
}

// SetWithdrawAddress is a paid mutator transaction binding the contract method 0x3ab1a494.
//
// Solidity: function setWithdrawAddress(address _withdrawAddress) returns(bool)
func (_DistributionWrapper *DistributionWrapperTransactor) SetWithdrawAddress(opts *bind.TransactOpts, _withdrawAddress common.Address) (*types.Transaction, error) {
	return _DistributionWrapper.contract.Transact(opts, "setWithdrawAddress", _withdrawAddress)
}

// SetWithdrawAddress is a paid mutator transaction binding the contract method 0x3ab1a494.
//
// Solidity: function setWithdrawAddress(address _withdrawAddress) returns(bool)
func (_DistributionWrapper *DistributionWrapperSession) SetWithdrawAddress(_withdrawAddress common.Address) (*types.Transaction, error) {
	return _DistributionWrapper.Contract.SetWithdrawAddress(&_DistributionWrapper.TransactOpts, _withdrawAddress)
}

// SetWithdrawAddress is a paid mutator transaction binding the contract method 0x3ab1a494.
//
// Solidity: function setWithdrawAddress(address _withdrawAddress) returns(bool)
func (_DistributionWrapper *DistributionWrapperTransactorSession) SetWithdrawAddress(_withdrawAddress common.Address) (*types.Transaction, error) {
	return _DistributionWrapper.Contract.SetWithdrawAddress(&_DistributionWrapper.TransactOpts, _withdrawAddress)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xe20981ca.
//
// Solidity: function withdrawRewards(address _delegatorAddress, address _validatorAddress) returns()
func (_DistributionWrapper *DistributionWrapperTransactor) WithdrawRewards(opts *bind.TransactOpts, _delegatorAddress common.Address, _validatorAddress common.Address) (*types.Transaction, error) {
	return _DistributionWrapper.contract.Transact(opts, "withdrawRewards", _delegatorAddress, _validatorAddress)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xe20981ca.
//
// Solidity: function withdrawRewards(address _delegatorAddress, address _validatorAddress) returns()
func (_DistributionWrapper *DistributionWrapperSession) WithdrawRewards(_delegatorAddress common.Address, _validatorAddress common.Address) (*types.Transaction, error) {
	return _DistributionWrapper.Contract.WithdrawRewards(&_DistributionWrapper.TransactOpts, _delegatorAddress, _validatorAddress)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xe20981ca.
//
// Solidity: function withdrawRewards(address _delegatorAddress, address _validatorAddress) returns()
func (_DistributionWrapper *DistributionWrapperTransactorSession) WithdrawRewards(_delegatorAddress common.Address, _validatorAddress common.Address) (*types.Transaction, error) {
	return _DistributionWrapper.Contract.WithdrawRewards(&_DistributionWrapper.TransactOpts, _delegatorAddress, _validatorAddress)
}
