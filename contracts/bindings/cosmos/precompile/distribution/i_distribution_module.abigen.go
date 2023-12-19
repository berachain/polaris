// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package distribution

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

// CosmosCoin is an auto generated low-level Go binding around an user-defined struct.
type CosmosCoin struct {
	Amount *big.Int
	Denom  string
}

// IDistributionModuleValidatorReward is an auto generated low-level Go binding around an user-defined struct.
type IDistributionModuleValidatorReward struct {
	Validator common.Address
	Rewards   []CosmosCoin
}

// DistributionModuleMetaData contains all meta data concerning the DistributionModule contract.
var DistributionModuleMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getAllDelegatorRewards\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIDistributionModule.ValidatorReward[]\",\"components\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rewards\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDelegatorValidatorReward\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalDelegatorReward\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWithdrawAddress\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWithdrawEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setWithdrawAddress\",\"inputs\":[{\"name\":\"withdrawAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawDelegatorReward\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"SetWithdrawAddress\",\"inputs\":[{\"name\":\"withdrawAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawRewards\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false}]",
}

// DistributionModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use DistributionModuleMetaData.ABI instead.
var DistributionModuleABI = DistributionModuleMetaData.ABI

// DistributionModule is an auto generated Go binding around an Ethereum contract.
type DistributionModule struct {
	DistributionModuleCaller     // Read-only binding to the contract
	DistributionModuleTransactor // Write-only binding to the contract
	DistributionModuleFilterer   // Log filterer for contract events
}

// DistributionModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type DistributionModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributionModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DistributionModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributionModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DistributionModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributionModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DistributionModuleSession struct {
	Contract     *DistributionModule // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DistributionModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DistributionModuleCallerSession struct {
	Contract *DistributionModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// DistributionModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DistributionModuleTransactorSession struct {
	Contract     *DistributionModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// DistributionModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type DistributionModuleRaw struct {
	Contract *DistributionModule // Generic contract binding to access the raw methods on
}

// DistributionModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DistributionModuleCallerRaw struct {
	Contract *DistributionModuleCaller // Generic read-only contract binding to access the raw methods on
}

// DistributionModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DistributionModuleTransactorRaw struct {
	Contract *DistributionModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDistributionModule creates a new instance of DistributionModule, bound to a specific deployed contract.
func NewDistributionModule(address common.Address, backend bind.ContractBackend) (*DistributionModule, error) {
	contract, err := bindDistributionModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DistributionModule{DistributionModuleCaller: DistributionModuleCaller{contract: contract}, DistributionModuleTransactor: DistributionModuleTransactor{contract: contract}, DistributionModuleFilterer: DistributionModuleFilterer{contract: contract}}, nil
}

// NewDistributionModuleCaller creates a new read-only instance of DistributionModule, bound to a specific deployed contract.
func NewDistributionModuleCaller(address common.Address, caller bind.ContractCaller) (*DistributionModuleCaller, error) {
	contract, err := bindDistributionModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DistributionModuleCaller{contract: contract}, nil
}

// NewDistributionModuleTransactor creates a new write-only instance of DistributionModule, bound to a specific deployed contract.
func NewDistributionModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*DistributionModuleTransactor, error) {
	contract, err := bindDistributionModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DistributionModuleTransactor{contract: contract}, nil
}

// NewDistributionModuleFilterer creates a new log filterer instance of DistributionModule, bound to a specific deployed contract.
func NewDistributionModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*DistributionModuleFilterer, error) {
	contract, err := bindDistributionModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DistributionModuleFilterer{contract: contract}, nil
}

// bindDistributionModule binds a generic wrapper to an already deployed contract.
func bindDistributionModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DistributionModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DistributionModule *DistributionModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DistributionModule.Contract.DistributionModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DistributionModule *DistributionModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DistributionModule.Contract.DistributionModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DistributionModule *DistributionModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DistributionModule.Contract.DistributionModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DistributionModule *DistributionModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DistributionModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DistributionModule *DistributionModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DistributionModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DistributionModule *DistributionModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DistributionModule.Contract.contract.Transact(opts, method, params...)
}

// GetAllDelegatorRewards is a free data retrieval call binding the contract method 0x36e22c98.
//
// Solidity: function getAllDelegatorRewards(address delegator) view returns((address,(uint256,string)[])[])
func (_DistributionModule *DistributionModuleCaller) GetAllDelegatorRewards(opts *bind.CallOpts, delegator common.Address) ([]IDistributionModuleValidatorReward, error) {
	var out []interface{}
	err := _DistributionModule.contract.Call(opts, &out, "getAllDelegatorRewards", delegator)

	if err != nil {
		return *new([]IDistributionModuleValidatorReward), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDistributionModuleValidatorReward)).(*[]IDistributionModuleValidatorReward)

	return out0, err

}

// GetAllDelegatorRewards is a free data retrieval call binding the contract method 0x36e22c98.
//
// Solidity: function getAllDelegatorRewards(address delegator) view returns((address,(uint256,string)[])[])
func (_DistributionModule *DistributionModuleSession) GetAllDelegatorRewards(delegator common.Address) ([]IDistributionModuleValidatorReward, error) {
	return _DistributionModule.Contract.GetAllDelegatorRewards(&_DistributionModule.CallOpts, delegator)
}

// GetAllDelegatorRewards is a free data retrieval call binding the contract method 0x36e22c98.
//
// Solidity: function getAllDelegatorRewards(address delegator) view returns((address,(uint256,string)[])[])
func (_DistributionModule *DistributionModuleCallerSession) GetAllDelegatorRewards(delegator common.Address) ([]IDistributionModuleValidatorReward, error) {
	return _DistributionModule.Contract.GetAllDelegatorRewards(&_DistributionModule.CallOpts, delegator)
}

// GetDelegatorValidatorReward is a free data retrieval call binding the contract method 0x4d33a513.
//
// Solidity: function getDelegatorValidatorReward(address delegator, address validator) view returns((uint256,string)[])
func (_DistributionModule *DistributionModuleCaller) GetDelegatorValidatorReward(opts *bind.CallOpts, delegator common.Address, validator common.Address) ([]CosmosCoin, error) {
	var out []interface{}
	err := _DistributionModule.contract.Call(opts, &out, "getDelegatorValidatorReward", delegator, validator)

	if err != nil {
		return *new([]CosmosCoin), err
	}

	out0 := *abi.ConvertType(out[0], new([]CosmosCoin)).(*[]CosmosCoin)

	return out0, err

}

// GetDelegatorValidatorReward is a free data retrieval call binding the contract method 0x4d33a513.
//
// Solidity: function getDelegatorValidatorReward(address delegator, address validator) view returns((uint256,string)[])
func (_DistributionModule *DistributionModuleSession) GetDelegatorValidatorReward(delegator common.Address, validator common.Address) ([]CosmosCoin, error) {
	return _DistributionModule.Contract.GetDelegatorValidatorReward(&_DistributionModule.CallOpts, delegator, validator)
}

// GetDelegatorValidatorReward is a free data retrieval call binding the contract method 0x4d33a513.
//
// Solidity: function getDelegatorValidatorReward(address delegator, address validator) view returns((uint256,string)[])
func (_DistributionModule *DistributionModuleCallerSession) GetDelegatorValidatorReward(delegator common.Address, validator common.Address) ([]CosmosCoin, error) {
	return _DistributionModule.Contract.GetDelegatorValidatorReward(&_DistributionModule.CallOpts, delegator, validator)
}

// GetTotalDelegatorReward is a free data retrieval call binding the contract method 0xce3341b4.
//
// Solidity: function getTotalDelegatorReward(address delegator) view returns((uint256,string)[])
func (_DistributionModule *DistributionModuleCaller) GetTotalDelegatorReward(opts *bind.CallOpts, delegator common.Address) ([]CosmosCoin, error) {
	var out []interface{}
	err := _DistributionModule.contract.Call(opts, &out, "getTotalDelegatorReward", delegator)

	if err != nil {
		return *new([]CosmosCoin), err
	}

	out0 := *abi.ConvertType(out[0], new([]CosmosCoin)).(*[]CosmosCoin)

	return out0, err

}

// GetTotalDelegatorReward is a free data retrieval call binding the contract method 0xce3341b4.
//
// Solidity: function getTotalDelegatorReward(address delegator) view returns((uint256,string)[])
func (_DistributionModule *DistributionModuleSession) GetTotalDelegatorReward(delegator common.Address) ([]CosmosCoin, error) {
	return _DistributionModule.Contract.GetTotalDelegatorReward(&_DistributionModule.CallOpts, delegator)
}

// GetTotalDelegatorReward is a free data retrieval call binding the contract method 0xce3341b4.
//
// Solidity: function getTotalDelegatorReward(address delegator) view returns((uint256,string)[])
func (_DistributionModule *DistributionModuleCallerSession) GetTotalDelegatorReward(delegator common.Address) ([]CosmosCoin, error) {
	return _DistributionModule.Contract.GetTotalDelegatorReward(&_DistributionModule.CallOpts, delegator)
}

// GetWithdrawAddress is a free data retrieval call binding the contract method 0xafe46ea2.
//
// Solidity: function getWithdrawAddress(address delegator) view returns(address)
func (_DistributionModule *DistributionModuleCaller) GetWithdrawAddress(opts *bind.CallOpts, delegator common.Address) (common.Address, error) {
	var out []interface{}
	err := _DistributionModule.contract.Call(opts, &out, "getWithdrawAddress", delegator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWithdrawAddress is a free data retrieval call binding the contract method 0xafe46ea2.
//
// Solidity: function getWithdrawAddress(address delegator) view returns(address)
func (_DistributionModule *DistributionModuleSession) GetWithdrawAddress(delegator common.Address) (common.Address, error) {
	return _DistributionModule.Contract.GetWithdrawAddress(&_DistributionModule.CallOpts, delegator)
}

// GetWithdrawAddress is a free data retrieval call binding the contract method 0xafe46ea2.
//
// Solidity: function getWithdrawAddress(address delegator) view returns(address)
func (_DistributionModule *DistributionModuleCallerSession) GetWithdrawAddress(delegator common.Address) (common.Address, error) {
	return _DistributionModule.Contract.GetWithdrawAddress(&_DistributionModule.CallOpts, delegator)
}

// GetWithdrawEnabled is a free data retrieval call binding the contract method 0x39cc4c86.
//
// Solidity: function getWithdrawEnabled() view returns(bool)
func (_DistributionModule *DistributionModuleCaller) GetWithdrawEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DistributionModule.contract.Call(opts, &out, "getWithdrawEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetWithdrawEnabled is a free data retrieval call binding the contract method 0x39cc4c86.
//
// Solidity: function getWithdrawEnabled() view returns(bool)
func (_DistributionModule *DistributionModuleSession) GetWithdrawEnabled() (bool, error) {
	return _DistributionModule.Contract.GetWithdrawEnabled(&_DistributionModule.CallOpts)
}

// GetWithdrawEnabled is a free data retrieval call binding the contract method 0x39cc4c86.
//
// Solidity: function getWithdrawEnabled() view returns(bool)
func (_DistributionModule *DistributionModuleCallerSession) GetWithdrawEnabled() (bool, error) {
	return _DistributionModule.Contract.GetWithdrawEnabled(&_DistributionModule.CallOpts)
}

// SetWithdrawAddress is a paid mutator transaction binding the contract method 0x3ab1a494.
//
// Solidity: function setWithdrawAddress(address withdrawAddress) returns(bool)
func (_DistributionModule *DistributionModuleTransactor) SetWithdrawAddress(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _DistributionModule.contract.Transact(opts, "setWithdrawAddress", withdrawAddress)
}

// SetWithdrawAddress is a paid mutator transaction binding the contract method 0x3ab1a494.
//
// Solidity: function setWithdrawAddress(address withdrawAddress) returns(bool)
func (_DistributionModule *DistributionModuleSession) SetWithdrawAddress(withdrawAddress common.Address) (*types.Transaction, error) {
	return _DistributionModule.Contract.SetWithdrawAddress(&_DistributionModule.TransactOpts, withdrawAddress)
}

// SetWithdrawAddress is a paid mutator transaction binding the contract method 0x3ab1a494.
//
// Solidity: function setWithdrawAddress(address withdrawAddress) returns(bool)
func (_DistributionModule *DistributionModuleTransactorSession) SetWithdrawAddress(withdrawAddress common.Address) (*types.Transaction, error) {
	return _DistributionModule.Contract.SetWithdrawAddress(&_DistributionModule.TransactOpts, withdrawAddress)
}

// WithdrawDelegatorReward is a paid mutator transaction binding the contract method 0x562c67a4.
//
// Solidity: function withdrawDelegatorReward(address delegator, address validator) returns((uint256,string)[])
func (_DistributionModule *DistributionModuleTransactor) WithdrawDelegatorReward(opts *bind.TransactOpts, delegator common.Address, validator common.Address) (*types.Transaction, error) {
	return _DistributionModule.contract.Transact(opts, "withdrawDelegatorReward", delegator, validator)
}

// WithdrawDelegatorReward is a paid mutator transaction binding the contract method 0x562c67a4.
//
// Solidity: function withdrawDelegatorReward(address delegator, address validator) returns((uint256,string)[])
func (_DistributionModule *DistributionModuleSession) WithdrawDelegatorReward(delegator common.Address, validator common.Address) (*types.Transaction, error) {
	return _DistributionModule.Contract.WithdrawDelegatorReward(&_DistributionModule.TransactOpts, delegator, validator)
}

// WithdrawDelegatorReward is a paid mutator transaction binding the contract method 0x562c67a4.
//
// Solidity: function withdrawDelegatorReward(address delegator, address validator) returns((uint256,string)[])
func (_DistributionModule *DistributionModuleTransactorSession) WithdrawDelegatorReward(delegator common.Address, validator common.Address) (*types.Transaction, error) {
	return _DistributionModule.Contract.WithdrawDelegatorReward(&_DistributionModule.TransactOpts, delegator, validator)
}

// DistributionModuleSetWithdrawAddressIterator is returned from FilterSetWithdrawAddress and is used to iterate over the raw logs and unpacked data for SetWithdrawAddress events raised by the DistributionModule contract.
type DistributionModuleSetWithdrawAddressIterator struct {
	Event *DistributionModuleSetWithdrawAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DistributionModuleSetWithdrawAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistributionModuleSetWithdrawAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DistributionModuleSetWithdrawAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DistributionModuleSetWithdrawAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistributionModuleSetWithdrawAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistributionModuleSetWithdrawAddress represents a SetWithdrawAddress event raised by the DistributionModule contract.
type DistributionModuleSetWithdrawAddress struct {
	WithdrawAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetWithdrawAddress is a free log retrieval operation binding the contract event 0xb13cf87e0a7b64f90565a1b68b63ae634d746fa785450bbdef7cbd281997cfb0.
//
// Solidity: event SetWithdrawAddress(address indexed withdrawAddress)
func (_DistributionModule *DistributionModuleFilterer) FilterSetWithdrawAddress(opts *bind.FilterOpts, withdrawAddress []common.Address) (*DistributionModuleSetWithdrawAddressIterator, error) {

	var withdrawAddressRule []interface{}
	for _, withdrawAddressItem := range withdrawAddress {
		withdrawAddressRule = append(withdrawAddressRule, withdrawAddressItem)
	}

	logs, sub, err := _DistributionModule.contract.FilterLogs(opts, "SetWithdrawAddress", withdrawAddressRule)
	if err != nil {
		return nil, err
	}
	return &DistributionModuleSetWithdrawAddressIterator{contract: _DistributionModule.contract, event: "SetWithdrawAddress", logs: logs, sub: sub}, nil
}

// WatchSetWithdrawAddress is a free log subscription operation binding the contract event 0xb13cf87e0a7b64f90565a1b68b63ae634d746fa785450bbdef7cbd281997cfb0.
//
// Solidity: event SetWithdrawAddress(address indexed withdrawAddress)
func (_DistributionModule *DistributionModuleFilterer) WatchSetWithdrawAddress(opts *bind.WatchOpts, sink chan<- *DistributionModuleSetWithdrawAddress, withdrawAddress []common.Address) (event.Subscription, error) {

	var withdrawAddressRule []interface{}
	for _, withdrawAddressItem := range withdrawAddress {
		withdrawAddressRule = append(withdrawAddressRule, withdrawAddressItem)
	}

	logs, sub, err := _DistributionModule.contract.WatchLogs(opts, "SetWithdrawAddress", withdrawAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistributionModuleSetWithdrawAddress)
				if err := _DistributionModule.contract.UnpackLog(event, "SetWithdrawAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetWithdrawAddress is a log parse operation binding the contract event 0xb13cf87e0a7b64f90565a1b68b63ae634d746fa785450bbdef7cbd281997cfb0.
//
// Solidity: event SetWithdrawAddress(address indexed withdrawAddress)
func (_DistributionModule *DistributionModuleFilterer) ParseSetWithdrawAddress(log types.Log) (*DistributionModuleSetWithdrawAddress, error) {
	event := new(DistributionModuleSetWithdrawAddress)
	if err := _DistributionModule.contract.UnpackLog(event, "SetWithdrawAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DistributionModuleWithdrawRewardsIterator is returned from FilterWithdrawRewards and is used to iterate over the raw logs and unpacked data for WithdrawRewards events raised by the DistributionModule contract.
type DistributionModuleWithdrawRewardsIterator struct {
	Event *DistributionModuleWithdrawRewards // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DistributionModuleWithdrawRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistributionModuleWithdrawRewards)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DistributionModuleWithdrawRewards)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DistributionModuleWithdrawRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistributionModuleWithdrawRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistributionModuleWithdrawRewards represents a WithdrawRewards event raised by the DistributionModule contract.
type DistributionModuleWithdrawRewards struct {
	Validator common.Address
	Amount    []CosmosCoin
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawRewards is a free log retrieval operation binding the contract event 0x68e5a74be48fc61c2ca4536b8819962850d1d39d7aa32ca670914f8b10aa6e5b.
//
// Solidity: event WithdrawRewards(address indexed validator, (uint256,string)[] amount)
func (_DistributionModule *DistributionModuleFilterer) FilterWithdrawRewards(opts *bind.FilterOpts, validator []common.Address) (*DistributionModuleWithdrawRewardsIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DistributionModule.contract.FilterLogs(opts, "WithdrawRewards", validatorRule)
	if err != nil {
		return nil, err
	}
	return &DistributionModuleWithdrawRewardsIterator{contract: _DistributionModule.contract, event: "WithdrawRewards", logs: logs, sub: sub}, nil
}

// WatchWithdrawRewards is a free log subscription operation binding the contract event 0x68e5a74be48fc61c2ca4536b8819962850d1d39d7aa32ca670914f8b10aa6e5b.
//
// Solidity: event WithdrawRewards(address indexed validator, (uint256,string)[] amount)
func (_DistributionModule *DistributionModuleFilterer) WatchWithdrawRewards(opts *bind.WatchOpts, sink chan<- *DistributionModuleWithdrawRewards, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _DistributionModule.contract.WatchLogs(opts, "WithdrawRewards", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistributionModuleWithdrawRewards)
				if err := _DistributionModule.contract.UnpackLog(event, "WithdrawRewards", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawRewards is a log parse operation binding the contract event 0x68e5a74be48fc61c2ca4536b8819962850d1d39d7aa32ca670914f8b10aa6e5b.
//
// Solidity: event WithdrawRewards(address indexed validator, (uint256,string)[] amount)
func (_DistributionModule *DistributionModuleFilterer) ParseWithdrawRewards(log types.Log) (*DistributionModuleWithdrawRewards, error) {
	event := new(DistributionModuleWithdrawRewards)
	if err := _DistributionModule.contract.UnpackLog(event, "WithdrawRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
