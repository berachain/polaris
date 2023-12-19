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

// ConsumeGasMetaData contains all meta data concerning the ConsumeGas contract.
var ConsumeGasMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"consumeGas\",\"inputs\":[{\"name\":\"targetGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"GasConsumed\",\"inputs\":[{\"name\":\"gasUsed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561000f575f80fd5b506101cb8061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610029575f3560e01c8063a329e8de1461002d575b5f80fd5b610047600480360381019061004291906100e2565b610049565b005b5f5a90505b818161005a919061013a565b5a1161004e575f5a8261006d919061013a565b90507f1a2dc18f5a2dabdf3809a83ec652290b81d97d915bf5561908090bad91deffc48160405161009e919061017c565b60405180910390a1505050565b5f80fd5b5f819050919050565b6100c1816100af565b81146100cb575f80fd5b50565b5f813590506100dc816100b8565b92915050565b5f602082840312156100f7576100f66100ab565b5b5f610104848285016100ce565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610144826100af565b915061014f836100af565b92508282039050818111156101675761016661010d565b5b92915050565b610176816100af565b82525050565b5f60208201905061018f5f83018461016d565b9291505056fea2646970667358221220e84926caa73d3306d918820ca91d65bee582951ae6a1ea86d9fd343eacc467f364736f6c63430008170033",
}

// ConsumeGasABI is the input ABI used to generate the binding from.
// Deprecated: Use ConsumeGasMetaData.ABI instead.
var ConsumeGasABI = ConsumeGasMetaData.ABI

// ConsumeGasBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ConsumeGasMetaData.Bin instead.
var ConsumeGasBin = ConsumeGasMetaData.Bin

// DeployConsumeGas deploys a new Ethereum contract, binding an instance of ConsumeGas to it.
func DeployConsumeGas(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ConsumeGas, error) {
	parsed, err := ConsumeGasMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ConsumeGasBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ConsumeGas{ConsumeGasCaller: ConsumeGasCaller{contract: contract}, ConsumeGasTransactor: ConsumeGasTransactor{contract: contract}, ConsumeGasFilterer: ConsumeGasFilterer{contract: contract}}, nil
}

// ConsumeGas is an auto generated Go binding around an Ethereum contract.
type ConsumeGas struct {
	ConsumeGasCaller     // Read-only binding to the contract
	ConsumeGasTransactor // Write-only binding to the contract
	ConsumeGasFilterer   // Log filterer for contract events
}

// ConsumeGasCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConsumeGasCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsumeGasTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConsumeGasTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsumeGasFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConsumeGasFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsumeGasSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConsumeGasSession struct {
	Contract     *ConsumeGas       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConsumeGasCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConsumeGasCallerSession struct {
	Contract *ConsumeGasCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ConsumeGasTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConsumeGasTransactorSession struct {
	Contract     *ConsumeGasTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ConsumeGasRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConsumeGasRaw struct {
	Contract *ConsumeGas // Generic contract binding to access the raw methods on
}

// ConsumeGasCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConsumeGasCallerRaw struct {
	Contract *ConsumeGasCaller // Generic read-only contract binding to access the raw methods on
}

// ConsumeGasTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConsumeGasTransactorRaw struct {
	Contract *ConsumeGasTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConsumeGas creates a new instance of ConsumeGas, bound to a specific deployed contract.
func NewConsumeGas(address common.Address, backend bind.ContractBackend) (*ConsumeGas, error) {
	contract, err := bindConsumeGas(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConsumeGas{ConsumeGasCaller: ConsumeGasCaller{contract: contract}, ConsumeGasTransactor: ConsumeGasTransactor{contract: contract}, ConsumeGasFilterer: ConsumeGasFilterer{contract: contract}}, nil
}

// NewConsumeGasCaller creates a new read-only instance of ConsumeGas, bound to a specific deployed contract.
func NewConsumeGasCaller(address common.Address, caller bind.ContractCaller) (*ConsumeGasCaller, error) {
	contract, err := bindConsumeGas(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConsumeGasCaller{contract: contract}, nil
}

// NewConsumeGasTransactor creates a new write-only instance of ConsumeGas, bound to a specific deployed contract.
func NewConsumeGasTransactor(address common.Address, transactor bind.ContractTransactor) (*ConsumeGasTransactor, error) {
	contract, err := bindConsumeGas(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConsumeGasTransactor{contract: contract}, nil
}

// NewConsumeGasFilterer creates a new log filterer instance of ConsumeGas, bound to a specific deployed contract.
func NewConsumeGasFilterer(address common.Address, filterer bind.ContractFilterer) (*ConsumeGasFilterer, error) {
	contract, err := bindConsumeGas(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConsumeGasFilterer{contract: contract}, nil
}

// bindConsumeGas binds a generic wrapper to an already deployed contract.
func bindConsumeGas(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ConsumeGasMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConsumeGas *ConsumeGasRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConsumeGas.Contract.ConsumeGasCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConsumeGas *ConsumeGasRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConsumeGas.Contract.ConsumeGasTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConsumeGas *ConsumeGasRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConsumeGas.Contract.ConsumeGasTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConsumeGas *ConsumeGasCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConsumeGas.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConsumeGas *ConsumeGasTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConsumeGas.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConsumeGas *ConsumeGasTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConsumeGas.Contract.contract.Transact(opts, method, params...)
}

// ConsumeGas is a paid mutator transaction binding the contract method 0xa329e8de.
//
// Solidity: function consumeGas(uint256 targetGas) returns()
func (_ConsumeGas *ConsumeGasTransactor) ConsumeGas(opts *bind.TransactOpts, targetGas *big.Int) (*types.Transaction, error) {
	return _ConsumeGas.contract.Transact(opts, "consumeGas", targetGas)
}

// ConsumeGas is a paid mutator transaction binding the contract method 0xa329e8de.
//
// Solidity: function consumeGas(uint256 targetGas) returns()
func (_ConsumeGas *ConsumeGasSession) ConsumeGas(targetGas *big.Int) (*types.Transaction, error) {
	return _ConsumeGas.Contract.ConsumeGas(&_ConsumeGas.TransactOpts, targetGas)
}

// ConsumeGas is a paid mutator transaction binding the contract method 0xa329e8de.
//
// Solidity: function consumeGas(uint256 targetGas) returns()
func (_ConsumeGas *ConsumeGasTransactorSession) ConsumeGas(targetGas *big.Int) (*types.Transaction, error) {
	return _ConsumeGas.Contract.ConsumeGas(&_ConsumeGas.TransactOpts, targetGas)
}

// ConsumeGasGasConsumedIterator is returned from FilterGasConsumed and is used to iterate over the raw logs and unpacked data for GasConsumed events raised by the ConsumeGas contract.
type ConsumeGasGasConsumedIterator struct {
	Event *ConsumeGasGasConsumed // Event containing the contract specifics and raw log

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
func (it *ConsumeGasGasConsumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConsumeGasGasConsumed)
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
		it.Event = new(ConsumeGasGasConsumed)
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
func (it *ConsumeGasGasConsumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConsumeGasGasConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConsumeGasGasConsumed represents a GasConsumed event raised by the ConsumeGas contract.
type ConsumeGasGasConsumed struct {
	GasUsed *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGasConsumed is a free log retrieval operation binding the contract event 0x1a2dc18f5a2dabdf3809a83ec652290b81d97d915bf5561908090bad91deffc4.
//
// Solidity: event GasConsumed(uint256 gasUsed)
func (_ConsumeGas *ConsumeGasFilterer) FilterGasConsumed(opts *bind.FilterOpts) (*ConsumeGasGasConsumedIterator, error) {

	logs, sub, err := _ConsumeGas.contract.FilterLogs(opts, "GasConsumed")
	if err != nil {
		return nil, err
	}
	return &ConsumeGasGasConsumedIterator{contract: _ConsumeGas.contract, event: "GasConsumed", logs: logs, sub: sub}, nil
}

// WatchGasConsumed is a free log subscription operation binding the contract event 0x1a2dc18f5a2dabdf3809a83ec652290b81d97d915bf5561908090bad91deffc4.
//
// Solidity: event GasConsumed(uint256 gasUsed)
func (_ConsumeGas *ConsumeGasFilterer) WatchGasConsumed(opts *bind.WatchOpts, sink chan<- *ConsumeGasGasConsumed) (event.Subscription, error) {

	logs, sub, err := _ConsumeGas.contract.WatchLogs(opts, "GasConsumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConsumeGasGasConsumed)
				if err := _ConsumeGas.contract.UnpackLog(event, "GasConsumed", log); err != nil {
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

// ParseGasConsumed is a log parse operation binding the contract event 0x1a2dc18f5a2dabdf3809a83ec652290b81d97d915bf5561908090bad91deffc4.
//
// Solidity: event GasConsumed(uint256 gasUsed)
func (_ConsumeGas *ConsumeGasFilterer) ParseGasConsumed(log types.Log) (*ConsumeGasGasConsumed, error) {
	event := new(ConsumeGasGasConsumed)
	if err := _ConsumeGas.contract.UnpackLog(event, "GasConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
