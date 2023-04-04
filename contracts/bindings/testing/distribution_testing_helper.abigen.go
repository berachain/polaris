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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_distributionprecompile\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"distribution\",\"outputs\":[{\"internalType\":\"contractIDistributionModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWithdrawEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_withdrawAddress\",\"type\":\"address\"}],\"name\":\"setWithdrawAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161058838038061058883398181016040528101906100329190610141565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610098576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505061016e565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061010e826100e3565b9050919050565b61011e81610103565b811461012957600080fd5b50565b60008151905061013b81610115565b92915050565b600060208284031215610157576101566100de565b5b60006101658482850161012c565b91505092915050565b61040b8061017d6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806339cc4c86146100465780633ab1a494146100645780635ee58efc14610094575b600080fd5b61004e6100b2565b60405161005b919061022d565b60405180910390f35b61007e600480360381019061007991906102ab565b610149565b60405161008b919061022d565b60405180910390f35b61009c6101ee565b6040516100a99190610337565b60405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166339cc4c866040518163ffffffff1660e01b8152600401602060405180830381865afa158015610120573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610144919061037e565b905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633ab1a494836040518263ffffffff1660e01b81526004016101a591906103ba565b6020604051808303816000875af11580156101c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101e8919061037e565b50919050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008115159050919050565b61022781610212565b82525050565b6000602082019050610242600083018461021e565b92915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006102788261024d565b9050919050565b6102888161026d565b811461029357600080fd5b50565b6000813590506102a58161027f565b92915050565b6000602082840312156102c1576102c0610248565b5b60006102cf84828501610296565b91505092915050565b6000819050919050565b60006102fd6102f86102f38461024d565b6102d8565b61024d565b9050919050565b600061030f826102e2565b9050919050565b600061032182610304565b9050919050565b61033181610316565b82525050565b600060208201905061034c6000830184610328565b92915050565b61035b81610212565b811461036657600080fd5b50565b60008151905061037881610352565b92915050565b60006020828403121561039457610393610248565b5b60006103a284828501610369565b91505092915050565b6103b48161026d565b82525050565b60006020820190506103cf60008301846103ab565b9291505056fea26469706673582212203ae11499f9c4fb1c7ec861e1e34a88eaf80faa7a623fe0ba9ff671df2e060f6b64736f6c63430008130033",
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
