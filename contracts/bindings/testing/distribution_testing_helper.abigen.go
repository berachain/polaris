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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_distributionprecompile\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingprecompile\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distribution\",\"outputs\":[{\"internalType\":\"contractIDistributionModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWithdrawEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_withdrawAddress\",\"type\":\"address\"}],\"name\":\"setWithdrawAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staking\",\"outputs\":[{\"internalType\":\"contractIStakingModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"}],\"name\":\"withdrawRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610c42380380610c42833981810160405281019061003291906101bc565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614801561009a5750600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b156100d1576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050506101fc565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006101898261015e565b9050919050565b6101998161017e565b81146101a457600080fd5b50565b6000815190506101b681610190565b92915050565b600080604083850312156101d3576101d2610159565b5b60006101e1858286016101a7565b92505060206101f2858286016101a7565b9150509250929050565b610a378061020b6000396000f3fe6080604052600436106100555760003560e01c806339cc4c861461005a5780633ab1a494146100855780634cf088d9146100ae5780635c19a95c146100d95780635ee58efc146100f5578063e20981ca14610120575b600080fd5b34801561006657600080fd5b5061006f610149565b60405161007c9190610431565b60405180910390f35b34801561009157600080fd5b506100ac60048036038101906100a791906104be565b6101e0565b005b3480156100ba57600080fd5b506100c3610280565b6040516100d0919061054a565b60405180910390f35b6100f360048036038101906100ee91906104be565b6102a6565b005b34801561010157600080fd5b5061010a61034a565b6040516101179190610586565b60405180910390f35b34801561012c57600080fd5b50610147600480360381019061014291906105a1565b61036e565b005b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166339cc4c866040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101db919061060d565b905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633ab1a494826040518263ffffffff1660e01b81526004016102399190610649565b6020604051808303816000875af1158015610258573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061027c919061060d565b5050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663026e402b82346040518363ffffffff1660e01b815260040161030392919061067d565b6020604051808303816000875af1158015610322573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610346919061060d565b5050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663562c67a483836040518363ffffffff1660e01b81526004016103c99291906106a6565b6000604051808303816000875af11580156103e8573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061041191906109b8565b505050565b60008115159050919050565b61042b81610416565b82525050565b60006020820190506104466000830184610422565b92915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061048b82610460565b9050919050565b61049b81610480565b81146104a657600080fd5b50565b6000813590506104b881610492565b92915050565b6000602082840312156104d4576104d3610456565b5b60006104e2848285016104a9565b91505092915050565b6000819050919050565b600061051061050b61050684610460565b6104eb565b610460565b9050919050565b6000610522826104f5565b9050919050565b600061053482610517565b9050919050565b61054481610529565b82525050565b600060208201905061055f600083018461053b565b92915050565b600061057082610517565b9050919050565b61058081610565565b82525050565b600060208201905061059b6000830184610577565b92915050565b600080604083850312156105b8576105b7610456565b5b60006105c6858286016104a9565b92505060206105d7858286016104a9565b9150509250929050565b6105ea81610416565b81146105f557600080fd5b50565b600081519050610607816105e1565b92915050565b60006020828403121561062357610622610456565b5b6000610631848285016105f8565b91505092915050565b61064381610480565b82525050565b600060208201905061065e600083018461063a565b92915050565b6000819050919050565b61067781610664565b82525050565b6000604082019050610692600083018561063a565b61069f602083018461066e565b9392505050565b60006040820190506106bb600083018561063a565b6106c8602083018461063a565b9392505050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61071d826106d4565b810181811067ffffffffffffffff8211171561073c5761073b6106e5565b5b80604052505050565b600061074f61044c565b905061075b8282610714565b919050565b600067ffffffffffffffff82111561077b5761077a6106e5565b5b602082029050602081019050919050565b600080fd5b600080fd5b600080fd5b6107a481610664565b81146107af57600080fd5b50565b6000815190506107c18161079b565b92915050565b600080fd5b600067ffffffffffffffff8211156107e7576107e66106e5565b5b6107f0826106d4565b9050602081019050919050565b60005b8381101561081b578082015181840152602081019050610800565b60008484015250505050565b600061083a610835846107cc565b610745565b905082815260208101848484011115610856576108556107c7565b5b6108618482856107fd565b509392505050565b600082601f83011261087e5761087d6106cf565b5b815161088e848260208601610827565b91505092915050565b6000604082840312156108ad576108ac610791565b5b6108b76040610745565b905060006108c7848285016107b2565b600083015250602082015167ffffffffffffffff8111156108eb576108ea610796565b5b6108f784828501610869565b60208301525092915050565b600061091661091184610760565b610745565b905080838252602082019050602084028301858111156109395761093861078c565b5b835b8181101561098057805167ffffffffffffffff81111561095e5761095d6106cf565b5b80860161096b8982610897565b8552602085019450505060208101905061093b565b5050509392505050565b600082601f83011261099f5761099e6106cf565b5b81516109af848260208601610903565b91505092915050565b6000602082840312156109ce576109cd610456565b5b600082015167ffffffffffffffff8111156109ec576109eb61045b565b5b6109f88482850161098a565b9150509291505056fea264697066735822122030eef51d63926b4656b438a27a7bdf1fd92ab8404548e57dc536983dd516f3b864736f6c63430008130033",
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
// Solidity: function setWithdrawAddress(address _withdrawAddress) returns()
func (_DistributionWrapper *DistributionWrapperTransactor) SetWithdrawAddress(opts *bind.TransactOpts, _withdrawAddress common.Address) (*types.Transaction, error) {
	return _DistributionWrapper.contract.Transact(opts, "setWithdrawAddress", _withdrawAddress)
}

// SetWithdrawAddress is a paid mutator transaction binding the contract method 0x3ab1a494.
//
// Solidity: function setWithdrawAddress(address _withdrawAddress) returns()
func (_DistributionWrapper *DistributionWrapperSession) SetWithdrawAddress(_withdrawAddress common.Address) (*types.Transaction, error) {
	return _DistributionWrapper.Contract.SetWithdrawAddress(&_DistributionWrapper.TransactOpts, _withdrawAddress)
}

// SetWithdrawAddress is a paid mutator transaction binding the contract method 0x3ab1a494.
//
// Solidity: function setWithdrawAddress(address _withdrawAddress) returns()
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
