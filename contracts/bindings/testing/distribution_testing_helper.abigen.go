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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_distributionprecompile\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_stakingprecompile\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegate\",\"inputs\":[{\"name\":\"_validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"distribution\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDistributionModule\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWithdrawEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setWithdrawAddress\",\"inputs\":[{\"name\":\"_withdrawAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"staking\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakingModule\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawRewards\",\"inputs\":[{\"name\":\"_delegatorAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_validatorAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051610bf8380380610bf8833981810160405281019061003191906101b2565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614801561009757505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b156100ce576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050506101f0565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61018182610158565b9050919050565b61019181610177565b811461019b575f80fd5b50565b5f815190506101ac81610188565b92915050565b5f80604083850312156101c8576101c7610154565b5b5f6101d58582860161019e565b92505060206101e68582860161019e565b9150509250929050565b6109fb806101fd5f395ff3fe608060405260043610610054575f3560e01c806339cc4c86146100585780633ab1a494146100825780634cf088d9146100be5780635c19a95c146100e85780635ee58efc14610104578063e20981ca1461012e575b5f80fd5b348015610063575f80fd5b5061006c610156565b604051610079919061042e565b60405180910390f35b34801561008d575f80fd5b506100a860048036038101906100a391906104b2565b6101e9565b6040516100b5919061042e565b60405180910390f35b3480156100c9575f80fd5b506100d261028a565b6040516100df9190610538565b60405180910390f35b61010260048036038101906100fd91906104b2565b6102af565b005b34801561010f575f80fd5b5061011861034f565b6040516101259190610571565b60405180910390f35b348015610139575f80fd5b50610154600480360381019061014f919061058a565b610372565b005b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166339cc4c866040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101c0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101e491906105f2565b905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633ab1a494836040518263ffffffff1660e01b8152600401610243919061062c565b6020604051808303815f875af115801561025f573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061028391906105f2565b9050919050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663026e402b82346040518363ffffffff1660e01b815260040161030b92919061065d565b6020604051808303815f875af1158015610327573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061034b91906105f2565b5050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663562c67a483836040518363ffffffff1660e01b81526004016103cc929190610684565b5f604051808303815f875af11580156103e7573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061040f919061097e565b505050565b5f8115159050919050565b61042881610414565b82525050565b5f6020820190506104415f83018461041f565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61048182610458565b9050919050565b61049181610477565b811461049b575f80fd5b50565b5f813590506104ac81610488565b92915050565b5f602082840312156104c7576104c6610450565b5b5f6104d48482850161049e565b91505092915050565b5f819050919050565b5f6105006104fb6104f684610458565b6104dd565b610458565b9050919050565b5f610511826104e6565b9050919050565b5f61052282610507565b9050919050565b61053281610518565b82525050565b5f60208201905061054b5f830184610529565b92915050565b5f61055b82610507565b9050919050565b61056b81610551565b82525050565b5f6020820190506105845f830184610562565b92915050565b5f80604083850312156105a05761059f610450565b5b5f6105ad8582860161049e565b92505060206105be8582860161049e565b9150509250929050565b6105d181610414565b81146105db575f80fd5b50565b5f815190506105ec816105c8565b92915050565b5f6020828403121561060757610606610450565b5b5f610614848285016105de565b91505092915050565b61062681610477565b82525050565b5f60208201905061063f5f83018461061d565b92915050565b5f819050919050565b61065781610645565b82525050565b5f6040820190506106705f83018561061d565b61067d602083018461064e565b9392505050565b5f6040820190506106975f83018561061d565b6106a4602083018461061d565b9392505050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6106f5826106af565b810181811067ffffffffffffffff82111715610714576107136106bf565b5b80604052505050565b5f610726610447565b905061073282826106ec565b919050565b5f67ffffffffffffffff821115610751576107506106bf565b5b602082029050602081019050919050565b5f80fd5b5f80fd5b5f80fd5b61077781610645565b8114610781575f80fd5b50565b5f815190506107928161076e565b92915050565b5f80fd5b5f67ffffffffffffffff8211156107b6576107b56106bf565b5b6107bf826106af565b9050602081019050919050565b5f5b838110156107e95780820151818401526020810190506107ce565b5f8484015250505050565b5f6108066108018461079c565b61071d565b90508281526020810184848401111561082257610821610798565b5b61082d8482856107cc565b509392505050565b5f82601f830112610849576108486106ab565b5b81516108598482602086016107f4565b91505092915050565b5f6040828403121561087757610876610766565b5b610881604061071d565b90505f61089084828501610784565b5f83015250602082015167ffffffffffffffff8111156108b3576108b261076a565b5b6108bf84828501610835565b60208301525092915050565b5f6108dd6108d884610737565b61071d565b90508083825260208201905060208402830185811115610900576108ff610762565b5b835b8181101561094757805167ffffffffffffffff811115610925576109246106ab565b5b8086016109328982610862565b85526020850194505050602081019050610902565b5050509392505050565b5f82601f830112610965576109646106ab565b5b81516109758482602086016108cb565b91505092915050565b5f6020828403121561099357610992610450565b5b5f82015167ffffffffffffffff8111156109b0576109af610454565b5b6109bc84828501610951565b9150509291505056fea2646970667358221220d470139c4b771c066c702ac9cc8ef914f3867f9779239e96fcc44723dcd6f0d964736f6c63430008170033",
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
