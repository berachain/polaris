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
	Bin: "0x608060405234801561000f575f80fd5b50604051610bf7380380610bf7833981810160405281019061003191906101b2565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614801561009757505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b156100ce576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050506101f0565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61018182610158565b9050919050565b61019181610177565b811461019b575f80fd5b50565b5f815190506101ac81610188565b92915050565b5f80604083850312156101c8576101c7610154565b5b5f6101d58582860161019e565b92505060206101e68582860161019e565b9150509250929050565b6109fa806101fd5f395ff3fe608060405260043610610054575f3560e01c806339cc4c86146100585780633ab1a494146100825780634cf088d9146100be5780635c19a95c146100e85780635ee58efc14610104578063e20981ca1461012e575b5f80fd5b348015610063575f80fd5b5061006c610156565b604051610079919061042d565b60405180910390f35b34801561008d575f80fd5b506100a860048036038101906100a391906104b1565b6101e9565b6040516100b5919061042d565b60405180910390f35b3480156100c9575f80fd5b506100d2610289565b6040516100df9190610537565b60405180910390f35b61010260048036038101906100fd91906104b1565b6102ae565b005b34801561010f575f80fd5b5061011861034e565b6040516101259190610570565b60405180910390f35b348015610139575f80fd5b50610154600480360381019061014f9190610589565b610371565b005b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166339cc4c866040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101c0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101e491906105f1565b905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633ab1a494836040518263ffffffff1660e01b8152600401610243919061062b565b6020604051808303815f875af115801561025f573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061028391906105f1565b50919050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663026e402b82346040518363ffffffff1660e01b815260040161030a92919061065c565b6020604051808303815f875af1158015610326573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061034a91906105f1565b5050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663562c67a483836040518363ffffffff1660e01b81526004016103cb929190610683565b5f604051808303815f875af11580156103e6573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061040e919061097d565b505050565b5f8115159050919050565b61042781610413565b82525050565b5f6020820190506104405f83018461041e565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61048082610457565b9050919050565b61049081610476565b811461049a575f80fd5b50565b5f813590506104ab81610487565b92915050565b5f602082840312156104c6576104c561044f565b5b5f6104d38482850161049d565b91505092915050565b5f819050919050565b5f6104ff6104fa6104f584610457565b6104dc565b610457565b9050919050565b5f610510826104e5565b9050919050565b5f61052182610506565b9050919050565b61053181610517565b82525050565b5f60208201905061054a5f830184610528565b92915050565b5f61055a82610506565b9050919050565b61056a81610550565b82525050565b5f6020820190506105835f830184610561565b92915050565b5f806040838503121561059f5761059e61044f565b5b5f6105ac8582860161049d565b92505060206105bd8582860161049d565b9150509250929050565b6105d081610413565b81146105da575f80fd5b50565b5f815190506105eb816105c7565b92915050565b5f602082840312156106065761060561044f565b5b5f610613848285016105dd565b91505092915050565b61062581610476565b82525050565b5f60208201905061063e5f83018461061c565b92915050565b5f819050919050565b61065681610644565b82525050565b5f60408201905061066f5f83018561061c565b61067c602083018461064d565b9392505050565b5f6040820190506106965f83018561061c565b6106a3602083018461061c565b9392505050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6106f4826106ae565b810181811067ffffffffffffffff82111715610713576107126106be565b5b80604052505050565b5f610725610446565b905061073182826106eb565b919050565b5f67ffffffffffffffff8211156107505761074f6106be565b5b602082029050602081019050919050565b5f80fd5b5f80fd5b5f80fd5b61077681610644565b8114610780575f80fd5b50565b5f815190506107918161076d565b92915050565b5f80fd5b5f67ffffffffffffffff8211156107b5576107b46106be565b5b6107be826106ae565b9050602081019050919050565b5f5b838110156107e85780820151818401526020810190506107cd565b5f8484015250505050565b5f6108056108008461079b565b61071c565b90508281526020810184848401111561082157610820610797565b5b61082c8482856107cb565b509392505050565b5f82601f830112610848576108476106aa565b5b81516108588482602086016107f3565b91505092915050565b5f6040828403121561087657610875610765565b5b610880604061071c565b90505f61088f84828501610783565b5f83015250602082015167ffffffffffffffff8111156108b2576108b1610769565b5b6108be84828501610834565b60208301525092915050565b5f6108dc6108d784610736565b61071c565b905080838252602082019050602084028301858111156108ff576108fe610761565b5b835b8181101561094657805167ffffffffffffffff811115610924576109236106aa565b5b8086016109318982610861565b85526020850194505050602081019050610901565b5050509392505050565b5f82601f830112610964576109636106aa565b5b81516109748482602086016108ca565b91505092915050565b5f602082840312156109925761099161044f565b5b5f82015167ffffffffffffffff8111156109af576109ae610453565b5b6109bb84828501610950565b9150509291505056fea2646970667358221220c549335c2aef93ae2f225c272220c8fc53de203579bad9bf7128d89a8280abe664736f6c63430008140033",
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
