// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bank

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

// BankModuleMetaData contains all meta data concerning the BankModule contract.
var BankModuleMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getAllBalances\",\"inputs\":[{\"name\":\"accountAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllSpendableBalances\",\"inputs\":[{\"name\":\"accountAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBalance\",\"inputs\":[{\"name\":\"accountAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSpendableBalance\",\"inputs\":[{\"name\":\"accountAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupply\",\"inputs\":[{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"send\",\"inputs\":[{\"name\":\"toAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Burn\",\"inputs\":[{\"name\":\"burner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CoinReceived\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CoinSpent\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Coinbase\",\"inputs\":[{\"name\":\"minter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Message\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false}]",
}

// BankModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use BankModuleMetaData.ABI instead.
var BankModuleABI = BankModuleMetaData.ABI

// BankModule is an auto generated Go binding around an Ethereum contract.
type BankModule struct {
	BankModuleCaller     // Read-only binding to the contract
	BankModuleTransactor // Write-only binding to the contract
	BankModuleFilterer   // Log filterer for contract events
}

// BankModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type BankModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BankModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BankModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BankModuleSession struct {
	Contract     *BankModule       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BankModuleCallerSession struct {
	Contract *BankModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BankModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BankModuleTransactorSession struct {
	Contract     *BankModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BankModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type BankModuleRaw struct {
	Contract *BankModule // Generic contract binding to access the raw methods on
}

// BankModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BankModuleCallerRaw struct {
	Contract *BankModuleCaller // Generic read-only contract binding to access the raw methods on
}

// BankModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BankModuleTransactorRaw struct {
	Contract *BankModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBankModule creates a new instance of BankModule, bound to a specific deployed contract.
func NewBankModule(address common.Address, backend bind.ContractBackend) (*BankModule, error) {
	contract, err := bindBankModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BankModule{BankModuleCaller: BankModuleCaller{contract: contract}, BankModuleTransactor: BankModuleTransactor{contract: contract}, BankModuleFilterer: BankModuleFilterer{contract: contract}}, nil
}

// NewBankModuleCaller creates a new read-only instance of BankModule, bound to a specific deployed contract.
func NewBankModuleCaller(address common.Address, caller bind.ContractCaller) (*BankModuleCaller, error) {
	contract, err := bindBankModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BankModuleCaller{contract: contract}, nil
}

// NewBankModuleTransactor creates a new write-only instance of BankModule, bound to a specific deployed contract.
func NewBankModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*BankModuleTransactor, error) {
	contract, err := bindBankModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BankModuleTransactor{contract: contract}, nil
}

// NewBankModuleFilterer creates a new log filterer instance of BankModule, bound to a specific deployed contract.
func NewBankModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*BankModuleFilterer, error) {
	contract, err := bindBankModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BankModuleFilterer{contract: contract}, nil
}

// bindBankModule binds a generic wrapper to an already deployed contract.
func bindBankModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BankModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BankModule *BankModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BankModule.Contract.BankModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BankModule *BankModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BankModule.Contract.BankModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BankModule *BankModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BankModule.Contract.BankModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BankModule *BankModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BankModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BankModule *BankModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BankModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BankModule *BankModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BankModule.Contract.contract.Transact(opts, method, params...)
}

// GetAllBalances is a free data retrieval call binding the contract method 0xc53d6ce1.
//
// Solidity: function getAllBalances(address accountAddress) view returns((uint256,string)[])
func (_BankModule *BankModuleCaller) GetAllBalances(opts *bind.CallOpts, accountAddress common.Address) ([]CosmosCoin, error) {
	var out []interface{}
	err := _BankModule.contract.Call(opts, &out, "getAllBalances", accountAddress)

	if err != nil {
		return *new([]CosmosCoin), err
	}

	out0 := *abi.ConvertType(out[0], new([]CosmosCoin)).(*[]CosmosCoin)

	return out0, err

}

// GetAllBalances is a free data retrieval call binding the contract method 0xc53d6ce1.
//
// Solidity: function getAllBalances(address accountAddress) view returns((uint256,string)[])
func (_BankModule *BankModuleSession) GetAllBalances(accountAddress common.Address) ([]CosmosCoin, error) {
	return _BankModule.Contract.GetAllBalances(&_BankModule.CallOpts, accountAddress)
}

// GetAllBalances is a free data retrieval call binding the contract method 0xc53d6ce1.
//
// Solidity: function getAllBalances(address accountAddress) view returns((uint256,string)[])
func (_BankModule *BankModuleCallerSession) GetAllBalances(accountAddress common.Address) ([]CosmosCoin, error) {
	return _BankModule.Contract.GetAllBalances(&_BankModule.CallOpts, accountAddress)
}

// GetAllSpendableBalances is a free data retrieval call binding the contract method 0x5c70e594.
//
// Solidity: function getAllSpendableBalances(address accountAddress) view returns((uint256,string)[])
func (_BankModule *BankModuleCaller) GetAllSpendableBalances(opts *bind.CallOpts, accountAddress common.Address) ([]CosmosCoin, error) {
	var out []interface{}
	err := _BankModule.contract.Call(opts, &out, "getAllSpendableBalances", accountAddress)

	if err != nil {
		return *new([]CosmosCoin), err
	}

	out0 := *abi.ConvertType(out[0], new([]CosmosCoin)).(*[]CosmosCoin)

	return out0, err

}

// GetAllSpendableBalances is a free data retrieval call binding the contract method 0x5c70e594.
//
// Solidity: function getAllSpendableBalances(address accountAddress) view returns((uint256,string)[])
func (_BankModule *BankModuleSession) GetAllSpendableBalances(accountAddress common.Address) ([]CosmosCoin, error) {
	return _BankModule.Contract.GetAllSpendableBalances(&_BankModule.CallOpts, accountAddress)
}

// GetAllSpendableBalances is a free data retrieval call binding the contract method 0x5c70e594.
//
// Solidity: function getAllSpendableBalances(address accountAddress) view returns((uint256,string)[])
func (_BankModule *BankModuleCallerSession) GetAllSpendableBalances(accountAddress common.Address) ([]CosmosCoin, error) {
	return _BankModule.Contract.GetAllSpendableBalances(&_BankModule.CallOpts, accountAddress)
}

// GetAllSupply is a free data retrieval call binding the contract method 0xf01c9474.
//
// Solidity: function getAllSupply() view returns((uint256,string)[])
func (_BankModule *BankModuleCaller) GetAllSupply(opts *bind.CallOpts) ([]CosmosCoin, error) {
	var out []interface{}
	err := _BankModule.contract.Call(opts, &out, "getAllSupply")

	if err != nil {
		return *new([]CosmosCoin), err
	}

	out0 := *abi.ConvertType(out[0], new([]CosmosCoin)).(*[]CosmosCoin)

	return out0, err

}

// GetAllSupply is a free data retrieval call binding the contract method 0xf01c9474.
//
// Solidity: function getAllSupply() view returns((uint256,string)[])
func (_BankModule *BankModuleSession) GetAllSupply() ([]CosmosCoin, error) {
	return _BankModule.Contract.GetAllSupply(&_BankModule.CallOpts)
}

// GetAllSupply is a free data retrieval call binding the contract method 0xf01c9474.
//
// Solidity: function getAllSupply() view returns((uint256,string)[])
func (_BankModule *BankModuleCallerSession) GetAllSupply() ([]CosmosCoin, error) {
	return _BankModule.Contract.GetAllSupply(&_BankModule.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x1dd7cecf.
//
// Solidity: function getBalance(address accountAddress, string denom) view returns(uint256)
func (_BankModule *BankModuleCaller) GetBalance(opts *bind.CallOpts, accountAddress common.Address, denom string) (*big.Int, error) {
	var out []interface{}
	err := _BankModule.contract.Call(opts, &out, "getBalance", accountAddress, denom)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x1dd7cecf.
//
// Solidity: function getBalance(address accountAddress, string denom) view returns(uint256)
func (_BankModule *BankModuleSession) GetBalance(accountAddress common.Address, denom string) (*big.Int, error) {
	return _BankModule.Contract.GetBalance(&_BankModule.CallOpts, accountAddress, denom)
}

// GetBalance is a free data retrieval call binding the contract method 0x1dd7cecf.
//
// Solidity: function getBalance(address accountAddress, string denom) view returns(uint256)
func (_BankModule *BankModuleCallerSession) GetBalance(accountAddress common.Address, denom string) (*big.Int, error) {
	return _BankModule.Contract.GetBalance(&_BankModule.CallOpts, accountAddress, denom)
}

// GetSpendableBalance is a free data retrieval call binding the contract method 0x34d1fdaf.
//
// Solidity: function getSpendableBalance(address accountAddress, string denom) view returns(uint256)
func (_BankModule *BankModuleCaller) GetSpendableBalance(opts *bind.CallOpts, accountAddress common.Address, denom string) (*big.Int, error) {
	var out []interface{}
	err := _BankModule.contract.Call(opts, &out, "getSpendableBalance", accountAddress, denom)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSpendableBalance is a free data retrieval call binding the contract method 0x34d1fdaf.
//
// Solidity: function getSpendableBalance(address accountAddress, string denom) view returns(uint256)
func (_BankModule *BankModuleSession) GetSpendableBalance(accountAddress common.Address, denom string) (*big.Int, error) {
	return _BankModule.Contract.GetSpendableBalance(&_BankModule.CallOpts, accountAddress, denom)
}

// GetSpendableBalance is a free data retrieval call binding the contract method 0x34d1fdaf.
//
// Solidity: function getSpendableBalance(address accountAddress, string denom) view returns(uint256)
func (_BankModule *BankModuleCallerSession) GetSpendableBalance(accountAddress common.Address, denom string) (*big.Int, error) {
	return _BankModule.Contract.GetSpendableBalance(&_BankModule.CallOpts, accountAddress, denom)
}

// GetSupply is a free data retrieval call binding the contract method 0xfe3b2b88.
//
// Solidity: function getSupply(string denom) view returns(uint256)
func (_BankModule *BankModuleCaller) GetSupply(opts *bind.CallOpts, denom string) (*big.Int, error) {
	var out []interface{}
	err := _BankModule.contract.Call(opts, &out, "getSupply", denom)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSupply is a free data retrieval call binding the contract method 0xfe3b2b88.
//
// Solidity: function getSupply(string denom) view returns(uint256)
func (_BankModule *BankModuleSession) GetSupply(denom string) (*big.Int, error) {
	return _BankModule.Contract.GetSupply(&_BankModule.CallOpts, denom)
}

// GetSupply is a free data retrieval call binding the contract method 0xfe3b2b88.
//
// Solidity: function getSupply(string denom) view returns(uint256)
func (_BankModule *BankModuleCallerSession) GetSupply(denom string) (*big.Int, error) {
	return _BankModule.Contract.GetSupply(&_BankModule.CallOpts, denom)
}

// Send is a paid mutator transaction binding the contract method 0x7e075f07.
//
// Solidity: function send(address toAddress, (uint256,string)[] amount) payable returns(bool)
func (_BankModule *BankModuleTransactor) Send(opts *bind.TransactOpts, toAddress common.Address, amount []CosmosCoin) (*types.Transaction, error) {
	return _BankModule.contract.Transact(opts, "send", toAddress, amount)
}

// Send is a paid mutator transaction binding the contract method 0x7e075f07.
//
// Solidity: function send(address toAddress, (uint256,string)[] amount) payable returns(bool)
func (_BankModule *BankModuleSession) Send(toAddress common.Address, amount []CosmosCoin) (*types.Transaction, error) {
	return _BankModule.Contract.Send(&_BankModule.TransactOpts, toAddress, amount)
}

// Send is a paid mutator transaction binding the contract method 0x7e075f07.
//
// Solidity: function send(address toAddress, (uint256,string)[] amount) payable returns(bool)
func (_BankModule *BankModuleTransactorSession) Send(toAddress common.Address, amount []CosmosCoin) (*types.Transaction, error) {
	return _BankModule.Contract.Send(&_BankModule.TransactOpts, toAddress, amount)
}

// BankModuleBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the BankModule contract.
type BankModuleBurnIterator struct {
	Event *BankModuleBurn // Event containing the contract specifics and raw log

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
func (it *BankModuleBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankModuleBurn)
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
		it.Event = new(BankModuleBurn)
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
func (it *BankModuleBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankModuleBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankModuleBurn represents a Burn event raised by the BankModule contract.
type BankModuleBurn struct {
	Burner common.Address
	Amount []CosmosCoin
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x9fa0c2fb43a81906efbb089cd76002325d71b437612a2a987c707446629d6ab0.
//
// Solidity: event Burn(address indexed burner, (uint256,string)[] amount)
func (_BankModule *BankModuleFilterer) FilterBurn(opts *bind.FilterOpts, burner []common.Address) (*BankModuleBurnIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _BankModule.contract.FilterLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return &BankModuleBurnIterator{contract: _BankModule.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x9fa0c2fb43a81906efbb089cd76002325d71b437612a2a987c707446629d6ab0.
//
// Solidity: event Burn(address indexed burner, (uint256,string)[] amount)
func (_BankModule *BankModuleFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *BankModuleBurn, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _BankModule.contract.WatchLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankModuleBurn)
				if err := _BankModule.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x9fa0c2fb43a81906efbb089cd76002325d71b437612a2a987c707446629d6ab0.
//
// Solidity: event Burn(address indexed burner, (uint256,string)[] amount)
func (_BankModule *BankModuleFilterer) ParseBurn(log types.Log) (*BankModuleBurn, error) {
	event := new(BankModuleBurn)
	if err := _BankModule.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BankModuleCoinReceivedIterator is returned from FilterCoinReceived and is used to iterate over the raw logs and unpacked data for CoinReceived events raised by the BankModule contract.
type BankModuleCoinReceivedIterator struct {
	Event *BankModuleCoinReceived // Event containing the contract specifics and raw log

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
func (it *BankModuleCoinReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankModuleCoinReceived)
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
		it.Event = new(BankModuleCoinReceived)
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
func (it *BankModuleCoinReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankModuleCoinReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankModuleCoinReceived represents a CoinReceived event raised by the BankModule contract.
type BankModuleCoinReceived struct {
	Receiver common.Address
	Amount   []CosmosCoin
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCoinReceived is a free log retrieval operation binding the contract event 0x13f9c352919df1623a08e6d6d9eac5f774573896f09916d8fbc5d083095fc3b4.
//
// Solidity: event CoinReceived(address indexed receiver, (uint256,string)[] amount)
func (_BankModule *BankModuleFilterer) FilterCoinReceived(opts *bind.FilterOpts, receiver []common.Address) (*BankModuleCoinReceivedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _BankModule.contract.FilterLogs(opts, "CoinReceived", receiverRule)
	if err != nil {
		return nil, err
	}
	return &BankModuleCoinReceivedIterator{contract: _BankModule.contract, event: "CoinReceived", logs: logs, sub: sub}, nil
}

// WatchCoinReceived is a free log subscription operation binding the contract event 0x13f9c352919df1623a08e6d6d9eac5f774573896f09916d8fbc5d083095fc3b4.
//
// Solidity: event CoinReceived(address indexed receiver, (uint256,string)[] amount)
func (_BankModule *BankModuleFilterer) WatchCoinReceived(opts *bind.WatchOpts, sink chan<- *BankModuleCoinReceived, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _BankModule.contract.WatchLogs(opts, "CoinReceived", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankModuleCoinReceived)
				if err := _BankModule.contract.UnpackLog(event, "CoinReceived", log); err != nil {
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

// ParseCoinReceived is a log parse operation binding the contract event 0x13f9c352919df1623a08e6d6d9eac5f774573896f09916d8fbc5d083095fc3b4.
//
// Solidity: event CoinReceived(address indexed receiver, (uint256,string)[] amount)
func (_BankModule *BankModuleFilterer) ParseCoinReceived(log types.Log) (*BankModuleCoinReceived, error) {
	event := new(BankModuleCoinReceived)
	if err := _BankModule.contract.UnpackLog(event, "CoinReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BankModuleCoinSpentIterator is returned from FilterCoinSpent and is used to iterate over the raw logs and unpacked data for CoinSpent events raised by the BankModule contract.
type BankModuleCoinSpentIterator struct {
	Event *BankModuleCoinSpent // Event containing the contract specifics and raw log

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
func (it *BankModuleCoinSpentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankModuleCoinSpent)
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
		it.Event = new(BankModuleCoinSpent)
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
func (it *BankModuleCoinSpentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankModuleCoinSpentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankModuleCoinSpent represents a CoinSpent event raised by the BankModule contract.
type BankModuleCoinSpent struct {
	Spender common.Address
	Amount  []CosmosCoin
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCoinSpent is a free log retrieval operation binding the contract event 0x8b8b22fea5b121b174e6cfea34ddaf187b66b43dab67679fa291a0fae2427a99.
//
// Solidity: event CoinSpent(address indexed spender, (uint256,string)[] amount)
func (_BankModule *BankModuleFilterer) FilterCoinSpent(opts *bind.FilterOpts, spender []common.Address) (*BankModuleCoinSpentIterator, error) {

	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BankModule.contract.FilterLogs(opts, "CoinSpent", spenderRule)
	if err != nil {
		return nil, err
	}
	return &BankModuleCoinSpentIterator{contract: _BankModule.contract, event: "CoinSpent", logs: logs, sub: sub}, nil
}

// WatchCoinSpent is a free log subscription operation binding the contract event 0x8b8b22fea5b121b174e6cfea34ddaf187b66b43dab67679fa291a0fae2427a99.
//
// Solidity: event CoinSpent(address indexed spender, (uint256,string)[] amount)
func (_BankModule *BankModuleFilterer) WatchCoinSpent(opts *bind.WatchOpts, sink chan<- *BankModuleCoinSpent, spender []common.Address) (event.Subscription, error) {

	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BankModule.contract.WatchLogs(opts, "CoinSpent", spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankModuleCoinSpent)
				if err := _BankModule.contract.UnpackLog(event, "CoinSpent", log); err != nil {
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

// ParseCoinSpent is a log parse operation binding the contract event 0x8b8b22fea5b121b174e6cfea34ddaf187b66b43dab67679fa291a0fae2427a99.
//
// Solidity: event CoinSpent(address indexed spender, (uint256,string)[] amount)
func (_BankModule *BankModuleFilterer) ParseCoinSpent(log types.Log) (*BankModuleCoinSpent, error) {
	event := new(BankModuleCoinSpent)
	if err := _BankModule.contract.UnpackLog(event, "CoinSpent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BankModuleCoinbaseIterator is returned from FilterCoinbase and is used to iterate over the raw logs and unpacked data for Coinbase events raised by the BankModule contract.
type BankModuleCoinbaseIterator struct {
	Event *BankModuleCoinbase // Event containing the contract specifics and raw log

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
func (it *BankModuleCoinbaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankModuleCoinbase)
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
		it.Event = new(BankModuleCoinbase)
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
func (it *BankModuleCoinbaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankModuleCoinbaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankModuleCoinbase represents a Coinbase event raised by the BankModule contract.
type BankModuleCoinbase struct {
	Minter common.Address
	Amount []CosmosCoin
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCoinbase is a free log retrieval operation binding the contract event 0xefb3f1f2a9af64b5fcc2da3c5a088d780585c674b8075fe2a1ba6b0d906cbe9f.
//
// Solidity: event Coinbase(address indexed minter, (uint256,string)[] amount)
func (_BankModule *BankModuleFilterer) FilterCoinbase(opts *bind.FilterOpts, minter []common.Address) (*BankModuleCoinbaseIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _BankModule.contract.FilterLogs(opts, "Coinbase", minterRule)
	if err != nil {
		return nil, err
	}
	return &BankModuleCoinbaseIterator{contract: _BankModule.contract, event: "Coinbase", logs: logs, sub: sub}, nil
}

// WatchCoinbase is a free log subscription operation binding the contract event 0xefb3f1f2a9af64b5fcc2da3c5a088d780585c674b8075fe2a1ba6b0d906cbe9f.
//
// Solidity: event Coinbase(address indexed minter, (uint256,string)[] amount)
func (_BankModule *BankModuleFilterer) WatchCoinbase(opts *bind.WatchOpts, sink chan<- *BankModuleCoinbase, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _BankModule.contract.WatchLogs(opts, "Coinbase", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankModuleCoinbase)
				if err := _BankModule.contract.UnpackLog(event, "Coinbase", log); err != nil {
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

// ParseCoinbase is a log parse operation binding the contract event 0xefb3f1f2a9af64b5fcc2da3c5a088d780585c674b8075fe2a1ba6b0d906cbe9f.
//
// Solidity: event Coinbase(address indexed minter, (uint256,string)[] amount)
func (_BankModule *BankModuleFilterer) ParseCoinbase(log types.Log) (*BankModuleCoinbase, error) {
	event := new(BankModuleCoinbase)
	if err := _BankModule.contract.UnpackLog(event, "Coinbase", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BankModuleMessageIterator is returned from FilterMessage and is used to iterate over the raw logs and unpacked data for Message events raised by the BankModule contract.
type BankModuleMessageIterator struct {
	Event *BankModuleMessage // Event containing the contract specifics and raw log

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
func (it *BankModuleMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankModuleMessage)
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
		it.Event = new(BankModuleMessage)
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
func (it *BankModuleMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankModuleMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankModuleMessage represents a Message event raised by the BankModule contract.
type BankModuleMessage struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMessage is a free log retrieval operation binding the contract event 0x516772d06520d23d2705f0b46a1fa6deec0ae36a2c00db049bd5f4094a123b85.
//
// Solidity: event Message(address indexed sender)
func (_BankModule *BankModuleFilterer) FilterMessage(opts *bind.FilterOpts, sender []common.Address) (*BankModuleMessageIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BankModule.contract.FilterLogs(opts, "Message", senderRule)
	if err != nil {
		return nil, err
	}
	return &BankModuleMessageIterator{contract: _BankModule.contract, event: "Message", logs: logs, sub: sub}, nil
}

// WatchMessage is a free log subscription operation binding the contract event 0x516772d06520d23d2705f0b46a1fa6deec0ae36a2c00db049bd5f4094a123b85.
//
// Solidity: event Message(address indexed sender)
func (_BankModule *BankModuleFilterer) WatchMessage(opts *bind.WatchOpts, sink chan<- *BankModuleMessage, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BankModule.contract.WatchLogs(opts, "Message", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankModuleMessage)
				if err := _BankModule.contract.UnpackLog(event, "Message", log); err != nil {
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
func (_BankModule *BankModuleFilterer) ParseMessage(log types.Log) (*BankModuleMessage, error) {
	event := new(BankModuleMessage)
	if err := _BankModule.contract.UnpackLog(event, "Message", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BankModuleTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BankModule contract.
type BankModuleTransferIterator struct {
	Event *BankModuleTransfer // Event containing the contract specifics and raw log

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
func (it *BankModuleTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankModuleTransfer)
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
		it.Event = new(BankModuleTransfer)
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
func (it *BankModuleTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankModuleTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankModuleTransfer represents a Transfer event raised by the BankModule contract.
type BankModuleTransfer struct {
	Recipient common.Address
	Amount    []CosmosCoin
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0x65ff5b103f0567c3e1783dc0b40e725544567fb6f584d9b084abea2e26d20328.
//
// Solidity: event Transfer(address indexed recipient, (uint256,string)[] amount)
func (_BankModule *BankModuleFilterer) FilterTransfer(opts *bind.FilterOpts, recipient []common.Address) (*BankModuleTransferIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BankModule.contract.FilterLogs(opts, "Transfer", recipientRule)
	if err != nil {
		return nil, err
	}
	return &BankModuleTransferIterator{contract: _BankModule.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0x65ff5b103f0567c3e1783dc0b40e725544567fb6f584d9b084abea2e26d20328.
//
// Solidity: event Transfer(address indexed recipient, (uint256,string)[] amount)
func (_BankModule *BankModuleFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BankModuleTransfer, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _BankModule.contract.WatchLogs(opts, "Transfer", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankModuleTransfer)
				if err := _BankModule.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0x65ff5b103f0567c3e1783dc0b40e725544567fb6f584d9b084abea2e26d20328.
//
// Solidity: event Transfer(address indexed recipient, (uint256,string)[] amount)
func (_BankModule *BankModuleFilterer) ParseTransfer(log types.Log) (*BankModuleTransfer, error) {
	event := new(BankModuleTransfer)
	if err := _BankModule.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
