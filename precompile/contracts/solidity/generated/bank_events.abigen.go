// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package generated

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

// BankEventsMetaData contains all meta data concerning the BankEvents contract.
var BankEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CoinReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CoinSpent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Coinbase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Message\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawRewards\",\"type\":\"event\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b386acc738527d2f310c26980656407dd31879eeb6e810c64878cdf86995512364736f6c63430008110033",
}

// BankEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use BankEventsMetaData.ABI instead.
var BankEventsABI = BankEventsMetaData.ABI

// BankEventsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BankEventsMetaData.Bin instead.
var BankEventsBin = BankEventsMetaData.Bin

// DeployBankEvents deploys a new Ethereum contract, binding an instance of BankEvents to it.
func DeployBankEvents(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BankEvents, error) {
	parsed, err := BankEventsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BankEventsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BankEvents{BankEventsCaller: BankEventsCaller{contract: contract}, BankEventsTransactor: BankEventsTransactor{contract: contract}, BankEventsFilterer: BankEventsFilterer{contract: contract}}, nil
}

// BankEvents is an auto generated Go binding around an Ethereum contract.
type BankEvents struct {
	BankEventsCaller     // Read-only binding to the contract
	BankEventsTransactor // Write-only binding to the contract
	BankEventsFilterer   // Log filterer for contract events
}

// BankEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BankEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BankEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BankEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BankEventsSession struct {
	Contract     *BankEvents       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BankEventsCallerSession struct {
	Contract *BankEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BankEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BankEventsTransactorSession struct {
	Contract     *BankEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BankEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BankEventsRaw struct {
	Contract *BankEvents // Generic contract binding to access the raw methods on
}

// BankEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BankEventsCallerRaw struct {
	Contract *BankEventsCaller // Generic read-only contract binding to access the raw methods on
}

// BankEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BankEventsTransactorRaw struct {
	Contract *BankEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBankEvents creates a new instance of BankEvents, bound to a specific deployed contract.
func NewBankEvents(address common.Address, backend bind.ContractBackend) (*BankEvents, error) {
	contract, err := bindBankEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BankEvents{BankEventsCaller: BankEventsCaller{contract: contract}, BankEventsTransactor: BankEventsTransactor{contract: contract}, BankEventsFilterer: BankEventsFilterer{contract: contract}}, nil
}

// NewBankEventsCaller creates a new read-only instance of BankEvents, bound to a specific deployed contract.
func NewBankEventsCaller(address common.Address, caller bind.ContractCaller) (*BankEventsCaller, error) {
	contract, err := bindBankEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BankEventsCaller{contract: contract}, nil
}

// NewBankEventsTransactor creates a new write-only instance of BankEvents, bound to a specific deployed contract.
func NewBankEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*BankEventsTransactor, error) {
	contract, err := bindBankEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BankEventsTransactor{contract: contract}, nil
}

// NewBankEventsFilterer creates a new log filterer instance of BankEvents, bound to a specific deployed contract.
func NewBankEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*BankEventsFilterer, error) {
	contract, err := bindBankEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BankEventsFilterer{contract: contract}, nil
}

// bindBankEvents binds a generic wrapper to an already deployed contract.
func bindBankEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BankEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BankEvents *BankEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BankEvents.Contract.BankEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BankEvents *BankEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BankEvents.Contract.BankEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BankEvents *BankEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BankEvents.Contract.BankEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BankEvents *BankEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BankEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BankEvents *BankEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BankEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BankEvents *BankEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BankEvents.Contract.contract.Transact(opts, method, params...)
}

// BankEventsBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the BankEvents contract.
type BankEventsBurnIterator struct {
	Event *BankEventsBurn // Event containing the contract specifics and raw log

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
func (it *BankEventsBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankEventsBurn)
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
		it.Event = new(BankEventsBurn)
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
func (it *BankEventsBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankEventsBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankEventsBurn represents a Burn event raised by the BankEvents contract.
type BankEventsBurn struct {
	Burner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed burner, uint256 amount)
func (_BankEvents *BankEventsFilterer) FilterBurn(opts *bind.FilterOpts, burner []common.Address) (*BankEventsBurnIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _BankEvents.contract.FilterLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return &BankEventsBurnIterator{contract: _BankEvents.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed burner, uint256 amount)
func (_BankEvents *BankEventsFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *BankEventsBurn, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _BankEvents.contract.WatchLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankEventsBurn)
				if err := _BankEvents.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed burner, uint256 amount)
func (_BankEvents *BankEventsFilterer) ParseBurn(log types.Log) (*BankEventsBurn, error) {
	event := new(BankEventsBurn)
	if err := _BankEvents.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BankEventsCoinReceivedIterator is returned from FilterCoinReceived and is used to iterate over the raw logs and unpacked data for CoinReceived events raised by the BankEvents contract.
type BankEventsCoinReceivedIterator struct {
	Event *BankEventsCoinReceived // Event containing the contract specifics and raw log

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
func (it *BankEventsCoinReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankEventsCoinReceived)
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
		it.Event = new(BankEventsCoinReceived)
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
func (it *BankEventsCoinReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankEventsCoinReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankEventsCoinReceived represents a CoinReceived event raised by the BankEvents contract.
type BankEventsCoinReceived struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCoinReceived is a free log retrieval operation binding the contract event 0xb2217585c246e507e3fcadd4d32d0177479d19c83452fabed3d7650cac3b0420.
//
// Solidity: event CoinReceived(address indexed receiver, uint256 amount)
func (_BankEvents *BankEventsFilterer) FilterCoinReceived(opts *bind.FilterOpts, receiver []common.Address) (*BankEventsCoinReceivedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _BankEvents.contract.FilterLogs(opts, "CoinReceived", receiverRule)
	if err != nil {
		return nil, err
	}
	return &BankEventsCoinReceivedIterator{contract: _BankEvents.contract, event: "CoinReceived", logs: logs, sub: sub}, nil
}

// WatchCoinReceived is a free log subscription operation binding the contract event 0xb2217585c246e507e3fcadd4d32d0177479d19c83452fabed3d7650cac3b0420.
//
// Solidity: event CoinReceived(address indexed receiver, uint256 amount)
func (_BankEvents *BankEventsFilterer) WatchCoinReceived(opts *bind.WatchOpts, sink chan<- *BankEventsCoinReceived, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _BankEvents.contract.WatchLogs(opts, "CoinReceived", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankEventsCoinReceived)
				if err := _BankEvents.contract.UnpackLog(event, "CoinReceived", log); err != nil {
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

// ParseCoinReceived is a log parse operation binding the contract event 0xb2217585c246e507e3fcadd4d32d0177479d19c83452fabed3d7650cac3b0420.
//
// Solidity: event CoinReceived(address indexed receiver, uint256 amount)
func (_BankEvents *BankEventsFilterer) ParseCoinReceived(log types.Log) (*BankEventsCoinReceived, error) {
	event := new(BankEventsCoinReceived)
	if err := _BankEvents.contract.UnpackLog(event, "CoinReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BankEventsCoinSpentIterator is returned from FilterCoinSpent and is used to iterate over the raw logs and unpacked data for CoinSpent events raised by the BankEvents contract.
type BankEventsCoinSpentIterator struct {
	Event *BankEventsCoinSpent // Event containing the contract specifics and raw log

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
func (it *BankEventsCoinSpentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankEventsCoinSpent)
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
		it.Event = new(BankEventsCoinSpent)
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
func (it *BankEventsCoinSpentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankEventsCoinSpentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankEventsCoinSpent represents a CoinSpent event raised by the BankEvents contract.
type BankEventsCoinSpent struct {
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCoinSpent is a free log retrieval operation binding the contract event 0xcd91156b4607a66e03194df5423537108085ebd28cca92794de7e9c53bd4c1c7.
//
// Solidity: event CoinSpent(address indexed spender, uint256 amount)
func (_BankEvents *BankEventsFilterer) FilterCoinSpent(opts *bind.FilterOpts, spender []common.Address) (*BankEventsCoinSpentIterator, error) {

	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BankEvents.contract.FilterLogs(opts, "CoinSpent", spenderRule)
	if err != nil {
		return nil, err
	}
	return &BankEventsCoinSpentIterator{contract: _BankEvents.contract, event: "CoinSpent", logs: logs, sub: sub}, nil
}

// WatchCoinSpent is a free log subscription operation binding the contract event 0xcd91156b4607a66e03194df5423537108085ebd28cca92794de7e9c53bd4c1c7.
//
// Solidity: event CoinSpent(address indexed spender, uint256 amount)
func (_BankEvents *BankEventsFilterer) WatchCoinSpent(opts *bind.WatchOpts, sink chan<- *BankEventsCoinSpent, spender []common.Address) (event.Subscription, error) {

	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BankEvents.contract.WatchLogs(opts, "CoinSpent", spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankEventsCoinSpent)
				if err := _BankEvents.contract.UnpackLog(event, "CoinSpent", log); err != nil {
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

// ParseCoinSpent is a log parse operation binding the contract event 0xcd91156b4607a66e03194df5423537108085ebd28cca92794de7e9c53bd4c1c7.
//
// Solidity: event CoinSpent(address indexed spender, uint256 amount)
func (_BankEvents *BankEventsFilterer) ParseCoinSpent(log types.Log) (*BankEventsCoinSpent, error) {
	event := new(BankEventsCoinSpent)
	if err := _BankEvents.contract.UnpackLog(event, "CoinSpent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BankEventsCoinbaseIterator is returned from FilterCoinbase and is used to iterate over the raw logs and unpacked data for Coinbase events raised by the BankEvents contract.
type BankEventsCoinbaseIterator struct {
	Event *BankEventsCoinbase // Event containing the contract specifics and raw log

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
func (it *BankEventsCoinbaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankEventsCoinbase)
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
		it.Event = new(BankEventsCoinbase)
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
func (it *BankEventsCoinbaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankEventsCoinbaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankEventsCoinbase represents a Coinbase event raised by the BankEvents contract.
type BankEventsCoinbase struct {
	Minter common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCoinbase is a free log retrieval operation binding the contract event 0x0bebe1d83c6be8a77c1af17badeea8b39108f5d721cfa73933bda78266af64e2.
//
// Solidity: event Coinbase(address indexed minter, uint256 amount)
func (_BankEvents *BankEventsFilterer) FilterCoinbase(opts *bind.FilterOpts, minter []common.Address) (*BankEventsCoinbaseIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _BankEvents.contract.FilterLogs(opts, "Coinbase", minterRule)
	if err != nil {
		return nil, err
	}
	return &BankEventsCoinbaseIterator{contract: _BankEvents.contract, event: "Coinbase", logs: logs, sub: sub}, nil
}

// WatchCoinbase is a free log subscription operation binding the contract event 0x0bebe1d83c6be8a77c1af17badeea8b39108f5d721cfa73933bda78266af64e2.
//
// Solidity: event Coinbase(address indexed minter, uint256 amount)
func (_BankEvents *BankEventsFilterer) WatchCoinbase(opts *bind.WatchOpts, sink chan<- *BankEventsCoinbase, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _BankEvents.contract.WatchLogs(opts, "Coinbase", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankEventsCoinbase)
				if err := _BankEvents.contract.UnpackLog(event, "Coinbase", log); err != nil {
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

// ParseCoinbase is a log parse operation binding the contract event 0x0bebe1d83c6be8a77c1af17badeea8b39108f5d721cfa73933bda78266af64e2.
//
// Solidity: event Coinbase(address indexed minter, uint256 amount)
func (_BankEvents *BankEventsFilterer) ParseCoinbase(log types.Log) (*BankEventsCoinbase, error) {
	event := new(BankEventsCoinbase)
	if err := _BankEvents.contract.UnpackLog(event, "Coinbase", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BankEventsMessageIterator is returned from FilterMessage and is used to iterate over the raw logs and unpacked data for Message events raised by the BankEvents contract.
type BankEventsMessageIterator struct {
	Event *BankEventsMessage // Event containing the contract specifics and raw log

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
func (it *BankEventsMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankEventsMessage)
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
		it.Event = new(BankEventsMessage)
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
func (it *BankEventsMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankEventsMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankEventsMessage represents a Message event raised by the BankEvents contract.
type BankEventsMessage struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMessage is a free log retrieval operation binding the contract event 0x516772d06520d23d2705f0b46a1fa6deec0ae36a2c00db049bd5f4094a123b85.
//
// Solidity: event Message(address indexed sender)
func (_BankEvents *BankEventsFilterer) FilterMessage(opts *bind.FilterOpts, sender []common.Address) (*BankEventsMessageIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BankEvents.contract.FilterLogs(opts, "Message", senderRule)
	if err != nil {
		return nil, err
	}
	return &BankEventsMessageIterator{contract: _BankEvents.contract, event: "Message", logs: logs, sub: sub}, nil
}

// WatchMessage is a free log subscription operation binding the contract event 0x516772d06520d23d2705f0b46a1fa6deec0ae36a2c00db049bd5f4094a123b85.
//
// Solidity: event Message(address indexed sender)
func (_BankEvents *BankEventsFilterer) WatchMessage(opts *bind.WatchOpts, sink chan<- *BankEventsMessage, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BankEvents.contract.WatchLogs(opts, "Message", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankEventsMessage)
				if err := _BankEvents.contract.UnpackLog(event, "Message", log); err != nil {
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

// ParseMessage is a log parse operation binding the contract event 0x516772d06520d23d2705f0b46a1fa6deec0ae36a2c00db049bd5f4094a123b85.
//
// Solidity: event Message(address indexed sender)
func (_BankEvents *BankEventsFilterer) ParseMessage(log types.Log) (*BankEventsMessage, error) {
	event := new(BankEventsMessage)
	if err := _BankEvents.contract.UnpackLog(event, "Message", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BankEventsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BankEvents contract.
type BankEventsTransferIterator struct {
	Event *BankEventsTransfer // Event containing the contract specifics and raw log

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
func (it *BankEventsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankEventsTransfer)
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
		it.Event = new(BankEventsTransfer)
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
func (it *BankEventsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankEventsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankEventsTransfer represents a Transfer event raised by the BankEvents contract.
type BankEventsTransfer struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address indexed recipient, uint256 amount)
func (_BankEvents *BankEventsFilterer) FilterTransfer(opts *bind.FilterOpts, recipient []common.Address) (*BankEventsTransferIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BankEvents.contract.FilterLogs(opts, "Transfer", recipientRule)
	if err != nil {
		return nil, err
	}
	return &BankEventsTransferIterator{contract: _BankEvents.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address indexed recipient, uint256 amount)
func (_BankEvents *BankEventsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BankEventsTransfer, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BankEvents.contract.WatchLogs(opts, "Transfer", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankEventsTransfer)
				if err := _BankEvents.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address indexed recipient, uint256 amount)
func (_BankEvents *BankEventsFilterer) ParseTransfer(log types.Log) (*BankEventsTransfer, error) {
	event := new(BankEventsTransfer)
	if err := _BankEvents.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BankEventsWithdrawRewardsIterator is returned from FilterWithdrawRewards and is used to iterate over the raw logs and unpacked data for WithdrawRewards events raised by the BankEvents contract.
type BankEventsWithdrawRewardsIterator struct {
	Event *BankEventsWithdrawRewards // Event containing the contract specifics and raw log

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
func (it *BankEventsWithdrawRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankEventsWithdrawRewards)
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
		it.Event = new(BankEventsWithdrawRewards)
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
func (it *BankEventsWithdrawRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankEventsWithdrawRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankEventsWithdrawRewards represents a WithdrawRewards event raised by the BankEvents contract.
type BankEventsWithdrawRewards struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawRewards is a free log retrieval operation binding the contract event 0xaa1377f7ec93c239e959efa811f7b8554c036fd7a706c23e58024626a8f3db96.
//
// Solidity: event WithdrawRewards(address indexed validator, uint256 amount)
func (_BankEvents *BankEventsFilterer) FilterWithdrawRewards(opts *bind.FilterOpts, validator []common.Address) (*BankEventsWithdrawRewardsIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BankEvents.contract.FilterLogs(opts, "WithdrawRewards", validatorRule)
	if err != nil {
		return nil, err
	}
	return &BankEventsWithdrawRewardsIterator{contract: _BankEvents.contract, event: "WithdrawRewards", logs: logs, sub: sub}, nil
}

// WatchWithdrawRewards is a free log subscription operation binding the contract event 0xaa1377f7ec93c239e959efa811f7b8554c036fd7a706c23e58024626a8f3db96.
//
// Solidity: event WithdrawRewards(address indexed validator, uint256 amount)
func (_BankEvents *BankEventsFilterer) WatchWithdrawRewards(opts *bind.WatchOpts, sink chan<- *BankEventsWithdrawRewards, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _BankEvents.contract.WatchLogs(opts, "WithdrawRewards", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankEventsWithdrawRewards)
				if err := _BankEvents.contract.UnpackLog(event, "WithdrawRewards", log); err != nil {
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

// ParseWithdrawRewards is a log parse operation binding the contract event 0xaa1377f7ec93c239e959efa811f7b8554c036fd7a706c23e58024626a8f3db96.
//
// Solidity: event WithdrawRewards(address indexed validator, uint256 amount)
func (_BankEvents *BankEventsFilterer) ParseWithdrawRewards(log types.Log) (*BankEventsWithdrawRewards, error) {
	event := new(BankEventsWithdrawRewards)
	if err := _BankEvents.contract.UnpackLog(event, "WithdrawRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
