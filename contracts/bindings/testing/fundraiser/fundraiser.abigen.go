// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testing_fundraiser

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

// FundraiserMetaData contains all meta data concerning the Fundraiser contract.
var FundraiserMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structCosmos.Coin[]\",\"name\":\"coins\",\"type\":\"tuple[]\"}],\"name\":\"Donate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetRaisedAmounts\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structCosmos.Coin[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"contractIBankModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawDonations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a0604052734381dc2ab14285160c808659aee005d51255add773ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff16815250348015610056575f80fd5b5033805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a350608051610ffa61011f5f395f8181610102015281816101a6015281816101ef01526103a90152610ffa5ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c80631ecc96521461006457806376cdb03b146100805780638da5cb5b1461009e578063af1d3f52146100bc578063ce1b088a146100da578063f2fde38b146100e4575b5f80fd5b61007e60048036038101906100799190610608565b610100565b005b6100886101a4565b60405161009591906106cd565b60405180910390f35b6100a66101c8565b6040516100b39190610706565b60405180910390f35b6100c46101eb565b6040516100d191906108b6565b60405180910390f35b6100e261028d565b005b6100fe60048036038101906100f99190610900565b61046d565b005b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166384404811333085856040518563ffffffff1660e01b815260040161015f9493929190610b20565b6020604051808303815f875af115801561017b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061019f9190610b93565b505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60607f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c53d6ce1306040518263ffffffff1660e01b81526004016102469190610706565b5f60405180830381865afa158015610260573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906102889190610e3b565b905090565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461031a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031190610edc565b60405180910390fd5b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039e90610f6a565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166384404811305f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1661040c6101eb565b6040518463ffffffff1660e01b815260040161042a93929190610f88565b6020604051808303815f875af1158015610446573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061046a9190610b93565b50565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104f190610edc565b60405180910390fd5b805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a350565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f8083601f8401126105c8576105c76105a7565b5b8235905067ffffffffffffffff8111156105e5576105e46105ab565b5b602083019150836020820283011115610601576106006105af565b5b9250929050565b5f806020838503121561061e5761061d61059f565b5b5f83013567ffffffffffffffff81111561063b5761063a6105a3565b5b610647858286016105b3565b92509250509250929050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f61069561069061068b84610653565b610672565b610653565b9050919050565b5f6106a68261067b565b9050919050565b5f6106b78261069c565b9050919050565b6106c7816106ad565b82525050565b5f6020820190506106e05f8301846106be565b92915050565b5f6106f082610653565b9050919050565b610700816106e6565b82525050565b5f6020820190506107195f8301846106f7565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f819050919050565b61075a81610748565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561079757808201518184015260208101905061077c565b5f8484015250505050565b5f601f19601f8301169050919050565b5f6107bc82610760565b6107c6818561076a565b93506107d681856020860161077a565b6107df816107a2565b840191505092915050565b5f604083015f8301516107ff5f860182610751565b506020830151848203602086015261081782826107b2565b9150508091505092915050565b5f61082f83836107ea565b905092915050565b5f602082019050919050565b5f61084d8261071f565b6108578185610729565b93508360208202850161086985610739565b805f5b858110156108a457848403895281516108858582610824565b945061089083610837565b925060208a0199505060018101905061086c565b50829750879550505050505092915050565b5f6020820190508181035f8301526108ce8184610843565b905092915050565b6108df816106e6565b81146108e9575f80fd5b50565b5f813590506108fa816108d6565b92915050565b5f602082840312156109155761091461059f565b5b5f610922848285016108ec565b91505092915050565b5f819050919050565b61093d81610748565b8114610947575f80fd5b50565b5f8135905061095881610934565b92915050565b5f61096c602084018461094a565b905092915050565b5f80fd5b5f80fd5b5f80fd5b5f808335600160200384360303811261099c5761099b61097c565b5b83810192508235915060208301925067ffffffffffffffff8211156109c4576109c3610974565b5b6001820236038313156109da576109d9610978565b5b509250929050565b828183375f83830152505050565b5f6109fb838561076a565b9350610a088385846109e2565b610a11836107a2565b840190509392505050565b5f60408301610a2d5f84018461095e565b610a395f860182610751565b50610a476020840184610980565b8583036020870152610a5a8382846109f0565b925050508091505092915050565b5f610a738383610a1c565b905092915050565b5f82356001604003833603038112610a9657610a9561097c565b5b82810191505092915050565b5f602082019050919050565b5f610ab98385610729565b935083602084028501610acb8461092b565b805f5b87811015610b0e578484038952610ae58284610a7b565b610aef8582610a68565b9450610afa83610aa2565b925060208a01995050600181019050610ace565b50829750879450505050509392505050565b5f606082019050610b335f8301876106f7565b610b4060208301866106f7565b8181036040830152610b53818486610aae565b905095945050505050565b5f8115159050919050565b610b7281610b5e565b8114610b7c575f80fd5b50565b5f81519050610b8d81610b69565b92915050565b5f60208284031215610ba857610ba761059f565b5b5f610bb584828501610b7f565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610bf4826107a2565b810181811067ffffffffffffffff82111715610c1357610c12610bbe565b5b80604052505050565b5f610c25610596565b9050610c318282610beb565b919050565b5f67ffffffffffffffff821115610c5057610c4f610bbe565b5b602082029050602081019050919050565b5f80fd5b5f80fd5b5f81519050610c7781610934565b92915050565b5f80fd5b5f67ffffffffffffffff821115610c9b57610c9a610bbe565b5b610ca4826107a2565b9050602081019050919050565b5f610cc3610cbe84610c81565b610c1c565b905082815260208101848484011115610cdf57610cde610c7d565b5b610cea84828561077a565b509392505050565b5f82601f830112610d0657610d056105a7565b5b8151610d16848260208601610cb1565b91505092915050565b5f60408284031215610d3457610d33610c61565b5b610d3e6040610c1c565b90505f610d4d84828501610c69565b5f83015250602082015167ffffffffffffffff811115610d7057610d6f610c65565b5b610d7c84828501610cf2565b60208301525092915050565b5f610d9a610d9584610c36565b610c1c565b90508083825260208201905060208402830185811115610dbd57610dbc6105af565b5b835b81811015610e0457805167ffffffffffffffff811115610de257610de16105a7565b5b808601610def8982610d1f565b85526020850194505050602081019050610dbf565b5050509392505050565b5f82601f830112610e2257610e216105a7565b5b8151610e32848260208601610d88565b91505092915050565b5f60208284031215610e5057610e4f61059f565b5b5f82015167ffffffffffffffff811115610e6d57610e6c6105a3565b5b610e7984828501610e0e565b91505092915050565b5f82825260208201905092915050565b7f554e415554484f52495a454400000000000000000000000000000000000000005f82015250565b5f610ec6600c83610e82565b9150610ed182610e92565b602082019050919050565b5f6020820190508181035f830152610ef381610eba565b9050919050565b7f46756e64732077696c6c206f6e6c792062652072656c656173656420746f20745f8201527f6865206f776e6572000000000000000000000000000000000000000000000000602082015250565b5f610f54602883610e82565b9150610f5f82610efa565b604082019050919050565b5f6020820190508181035f830152610f8181610f48565b9050919050565b5f606082019050610f9b5f8301866106f7565b610fa860208301856106f7565b8181036040830152610fba8184610843565b905094935050505056fea2646970667358221220ea04b625d38b40b750c4597832776b3113e12a9fabea296a5e93b1b6312c075964736f6c63430008140033",
}

// FundraiserABI is the input ABI used to generate the binding from.
// Deprecated: Use FundraiserMetaData.ABI instead.
var FundraiserABI = FundraiserMetaData.ABI

// FundraiserBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FundraiserMetaData.Bin instead.
var FundraiserBin = FundraiserMetaData.Bin

// DeployFundraiser deploys a new Ethereum contract, binding an instance of Fundraiser to it.
func DeployFundraiser(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Fundraiser, error) {
	parsed, err := FundraiserMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FundraiserBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Fundraiser{FundraiserCaller: FundraiserCaller{contract: contract}, FundraiserTransactor: FundraiserTransactor{contract: contract}, FundraiserFilterer: FundraiserFilterer{contract: contract}}, nil
}

// Fundraiser is an auto generated Go binding around an Ethereum contract.
type Fundraiser struct {
	FundraiserCaller     // Read-only binding to the contract
	FundraiserTransactor // Write-only binding to the contract
	FundraiserFilterer   // Log filterer for contract events
}

// FundraiserCaller is an auto generated read-only Go binding around an Ethereum contract.
type FundraiserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FundraiserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FundraiserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FundraiserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FundraiserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FundraiserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FundraiserSession struct {
	Contract     *Fundraiser       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FundraiserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FundraiserCallerSession struct {
	Contract *FundraiserCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FundraiserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FundraiserTransactorSession struct {
	Contract     *FundraiserTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FundraiserRaw is an auto generated low-level Go binding around an Ethereum contract.
type FundraiserRaw struct {
	Contract *Fundraiser // Generic contract binding to access the raw methods on
}

// FundraiserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FundraiserCallerRaw struct {
	Contract *FundraiserCaller // Generic read-only contract binding to access the raw methods on
}

// FundraiserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FundraiserTransactorRaw struct {
	Contract *FundraiserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFundraiser creates a new instance of Fundraiser, bound to a specific deployed contract.
func NewFundraiser(address common.Address, backend bind.ContractBackend) (*Fundraiser, error) {
	contract, err := bindFundraiser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Fundraiser{FundraiserCaller: FundraiserCaller{contract: contract}, FundraiserTransactor: FundraiserTransactor{contract: contract}, FundraiserFilterer: FundraiserFilterer{contract: contract}}, nil
}

// NewFundraiserCaller creates a new read-only instance of Fundraiser, bound to a specific deployed contract.
func NewFundraiserCaller(address common.Address, caller bind.ContractCaller) (*FundraiserCaller, error) {
	contract, err := bindFundraiser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FundraiserCaller{contract: contract}, nil
}

// NewFundraiserTransactor creates a new write-only instance of Fundraiser, bound to a specific deployed contract.
func NewFundraiserTransactor(address common.Address, transactor bind.ContractTransactor) (*FundraiserTransactor, error) {
	contract, err := bindFundraiser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FundraiserTransactor{contract: contract}, nil
}

// NewFundraiserFilterer creates a new log filterer instance of Fundraiser, bound to a specific deployed contract.
func NewFundraiserFilterer(address common.Address, filterer bind.ContractFilterer) (*FundraiserFilterer, error) {
	contract, err := bindFundraiser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FundraiserFilterer{contract: contract}, nil
}

// bindFundraiser binds a generic wrapper to an already deployed contract.
func bindFundraiser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FundraiserMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fundraiser *FundraiserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fundraiser.Contract.FundraiserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fundraiser *FundraiserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fundraiser.Contract.FundraiserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fundraiser *FundraiserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fundraiser.Contract.FundraiserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fundraiser *FundraiserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fundraiser.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fundraiser *FundraiserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fundraiser.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fundraiser *FundraiserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fundraiser.Contract.contract.Transact(opts, method, params...)
}

// GetRaisedAmounts is a free data retrieval call binding the contract method 0xaf1d3f52.
//
// Solidity: function GetRaisedAmounts() view returns((uint256,string)[])
func (_Fundraiser *FundraiserCaller) GetRaisedAmounts(opts *bind.CallOpts) ([]CosmosCoin, error) {
	var out []interface{}
	err := _Fundraiser.contract.Call(opts, &out, "GetRaisedAmounts")

	if err != nil {
		return *new([]CosmosCoin), err
	}

	out0 := *abi.ConvertType(out[0], new([]CosmosCoin)).(*[]CosmosCoin)

	return out0, err

}

// GetRaisedAmounts is a free data retrieval call binding the contract method 0xaf1d3f52.
//
// Solidity: function GetRaisedAmounts() view returns((uint256,string)[])
func (_Fundraiser *FundraiserSession) GetRaisedAmounts() ([]CosmosCoin, error) {
	return _Fundraiser.Contract.GetRaisedAmounts(&_Fundraiser.CallOpts)
}

// GetRaisedAmounts is a free data retrieval call binding the contract method 0xaf1d3f52.
//
// Solidity: function GetRaisedAmounts() view returns((uint256,string)[])
func (_Fundraiser *FundraiserCallerSession) GetRaisedAmounts() ([]CosmosCoin, error) {
	return _Fundraiser.Contract.GetRaisedAmounts(&_Fundraiser.CallOpts)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Fundraiser *FundraiserCaller) Bank(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fundraiser.contract.Call(opts, &out, "bank")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Fundraiser *FundraiserSession) Bank() (common.Address, error) {
	return _Fundraiser.Contract.Bank(&_Fundraiser.CallOpts)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Fundraiser *FundraiserCallerSession) Bank() (common.Address, error) {
	return _Fundraiser.Contract.Bank(&_Fundraiser.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Fundraiser *FundraiserCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fundraiser.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Fundraiser *FundraiserSession) Owner() (common.Address, error) {
	return _Fundraiser.Contract.Owner(&_Fundraiser.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Fundraiser *FundraiserCallerSession) Owner() (common.Address, error) {
	return _Fundraiser.Contract.Owner(&_Fundraiser.CallOpts)
}

// Donate is a paid mutator transaction binding the contract method 0x1ecc9652.
//
// Solidity: function Donate((uint256,string)[] coins) returns()
func (_Fundraiser *FundraiserTransactor) Donate(opts *bind.TransactOpts, coins []CosmosCoin) (*types.Transaction, error) {
	return _Fundraiser.contract.Transact(opts, "Donate", coins)
}

// Donate is a paid mutator transaction binding the contract method 0x1ecc9652.
//
// Solidity: function Donate((uint256,string)[] coins) returns()
func (_Fundraiser *FundraiserSession) Donate(coins []CosmosCoin) (*types.Transaction, error) {
	return _Fundraiser.Contract.Donate(&_Fundraiser.TransactOpts, coins)
}

// Donate is a paid mutator transaction binding the contract method 0x1ecc9652.
//
// Solidity: function Donate((uint256,string)[] coins) returns()
func (_Fundraiser *FundraiserTransactorSession) Donate(coins []CosmosCoin) (*types.Transaction, error) {
	return _Fundraiser.Contract.Donate(&_Fundraiser.TransactOpts, coins)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Fundraiser *FundraiserTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Fundraiser.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Fundraiser *FundraiserSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Fundraiser.Contract.TransferOwnership(&_Fundraiser.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Fundraiser *FundraiserTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Fundraiser.Contract.TransferOwnership(&_Fundraiser.TransactOpts, newOwner)
}

// WithdrawDonations is a paid mutator transaction binding the contract method 0xce1b088a.
//
// Solidity: function withdrawDonations() returns()
func (_Fundraiser *FundraiserTransactor) WithdrawDonations(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fundraiser.contract.Transact(opts, "withdrawDonations")
}

// WithdrawDonations is a paid mutator transaction binding the contract method 0xce1b088a.
//
// Solidity: function withdrawDonations() returns()
func (_Fundraiser *FundraiserSession) WithdrawDonations() (*types.Transaction, error) {
	return _Fundraiser.Contract.WithdrawDonations(&_Fundraiser.TransactOpts)
}

// WithdrawDonations is a paid mutator transaction binding the contract method 0xce1b088a.
//
// Solidity: function withdrawDonations() returns()
func (_Fundraiser *FundraiserTransactorSession) WithdrawDonations() (*types.Transaction, error) {
	return _Fundraiser.Contract.WithdrawDonations(&_Fundraiser.TransactOpts)
}

// FundraiserOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Fundraiser contract.
type FundraiserOwnershipTransferredIterator struct {
	Event *FundraiserOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FundraiserOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FundraiserOwnershipTransferred)
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
		it.Event = new(FundraiserOwnershipTransferred)
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
func (it *FundraiserOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FundraiserOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FundraiserOwnershipTransferred represents a OwnershipTransferred event raised by the Fundraiser contract.
type FundraiserOwnershipTransferred struct {
	User     common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_Fundraiser *FundraiserFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, user []common.Address, newOwner []common.Address) (*FundraiserOwnershipTransferredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Fundraiser.contract.FilterLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FundraiserOwnershipTransferredIterator{contract: _Fundraiser.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_Fundraiser *FundraiserFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FundraiserOwnershipTransferred, user []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Fundraiser.contract.WatchLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FundraiserOwnershipTransferred)
				if err := _Fundraiser.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_Fundraiser *FundraiserFilterer) ParseOwnershipTransferred(log types.Log) (*FundraiserOwnershipTransferred, error) {
	event := new(FundraiserOwnershipTransferred)
	if err := _Fundraiser.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
