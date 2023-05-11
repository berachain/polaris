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

// IBankModuleCoin is an auto generated low-level Go binding around an user-defined struct.
type IBankModuleCoin struct {
	Amount *big.Int
	Denom  string
}

// FundraiserMetaData contains all meta data concerning the Fundraiser contract.
var FundraiserMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structIBankModule.Coin[]\",\"name\":\"coins\",\"type\":\"tuple[]\"}],\"name\":\"Donate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetRaisedAmounts\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structIBankModule.Coin[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"contractIBankModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawDonations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a0604052734381dc2ab14285160c808659aee005d51255add773ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801561005757600080fd5b5033806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35060805161106d61012560003960008181610105015281816101ac015281816101f601526103b6015261106d6000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80631ecc96521461006757806376cdb03b146100835780638da5cb5b146100a1578063af1d3f52146100bf578063ce1b088a146100dd578063f2fde38b146100e7575b600080fd5b610081600480360381019061007c9190610622565b610103565b005b61008b6101aa565b60405161009891906106ee565b60405180910390f35b6100a96101ce565b6040516100b6919061072a565b60405180910390f35b6100c76101f2565b6040516100d491906108ed565b60405180910390f35b6100e5610298565b005b61010160048036038101906100fc919061093b565b61047e565b005b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166384404811333085856040518563ffffffff1660e01b81526004016101629493929190610b6f565b6020604051808303816000875af1158015610181573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101a59190610be7565b505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60607f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c53d6ce1306040518263ffffffff1660e01b815260040161024d919061072a565b600060405180830381865afa15801561026a573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906102939190610ea1565b905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610326576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031d90610f47565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ab90610fd9565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663844048113060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1661041a6101f2565b6040518463ffffffff1660e01b815260040161043893929190610ff9565b6020604051808303816000875af1158015610457573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061047b9190610be7565b50565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461050c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050390610f47565b60405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a350565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f8401126105e2576105e16105bd565b5b8235905067ffffffffffffffff8111156105ff576105fe6105c2565b5b60208301915083602082028301111561061b5761061a6105c7565b5b9250929050565b60008060208385031215610639576106386105b3565b5b600083013567ffffffffffffffff811115610657576106566105b8565b5b610663858286016105cc565b92509250509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006106b46106af6106aa8461066f565b61068f565b61066f565b9050919050565b60006106c682610699565b9050919050565b60006106d8826106bb565b9050919050565b6106e8816106cd565b82525050565b600060208201905061070360008301846106df565b92915050565b60006107148261066f565b9050919050565b61072481610709565b82525050565b600060208201905061073f600083018461071b565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000819050919050565b61078481610771565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b838110156107c45780820151818401526020810190506107a9565b60008484015250505050565b6000601f19601f8301169050919050565b60006107ec8261078a565b6107f68185610795565b93506108068185602086016107a6565b61080f816107d0565b840191505092915050565b6000604083016000830151610832600086018261077b565b506020830151848203602086015261084a82826107e1565b9150508091505092915050565b6000610863838361081a565b905092915050565b6000602082019050919050565b600061088382610745565b61088d8185610750565b93508360208202850161089f85610761565b8060005b858110156108db57848403895281516108bc8582610857565b94506108c78361086b565b925060208a019950506001810190506108a3565b50829750879550505050505092915050565b600060208201905081810360008301526109078184610878565b905092915050565b61091881610709565b811461092357600080fd5b50565b6000813590506109358161090f565b92915050565b600060208284031215610951576109506105b3565b5b600061095f84828501610926565b91505092915050565b6000819050919050565b61097b81610771565b811461098657600080fd5b50565b60008135905061099881610972565b92915050565b60006109ad6020840184610989565b905092915050565b600080fd5b600080fd5b600080fd5b600080833560016020038436030381126109e1576109e06109bf565b5b83810192508235915060208301925067ffffffffffffffff821115610a0957610a086109b5565b5b600182023603831315610a1f57610a1e6109ba565b5b509250929050565b82818337600083830152505050565b6000610a428385610795565b9350610a4f838584610a27565b610a58836107d0565b840190509392505050565b600060408301610a76600084018461099e565b610a83600086018261077b565b50610a9160208401846109c4565b8583036020870152610aa4838284610a36565b925050508091505092915050565b6000610abe8383610a63565b905092915050565b600082356001604003833603038112610ae257610ae16109bf565b5b82810191505092915050565b6000602082019050919050565b6000610b078385610750565b935083602084028501610b1984610968565b8060005b87811015610b5d578484038952610b348284610ac6565b610b3e8582610ab2565b9450610b4983610aee565b925060208a01995050600181019050610b1d565b50829750879450505050509392505050565b6000606082019050610b84600083018761071b565b610b91602083018661071b565b8181036040830152610ba4818486610afb565b905095945050505050565b60008115159050919050565b610bc481610baf565b8114610bcf57600080fd5b50565b600081519050610be181610bbb565b92915050565b600060208284031215610bfd57610bfc6105b3565b5b6000610c0b84828501610bd2565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610c4c826107d0565b810181811067ffffffffffffffff82111715610c6b57610c6a610c14565b5b80604052505050565b6000610c7e6105a9565b9050610c8a8282610c43565b919050565b600067ffffffffffffffff821115610caa57610ca9610c14565b5b602082029050602081019050919050565b600080fd5b600080fd5b600081519050610cd481610972565b92915050565b600080fd5b600067ffffffffffffffff821115610cfa57610cf9610c14565b5b610d03826107d0565b9050602081019050919050565b6000610d23610d1e84610cdf565b610c74565b905082815260208101848484011115610d3f57610d3e610cda565b5b610d4a8482856107a6565b509392505050565b600082601f830112610d6757610d666105bd565b5b8151610d77848260208601610d10565b91505092915050565b600060408284031215610d9657610d95610cbb565b5b610da06040610c74565b90506000610db084828501610cc5565b600083015250602082015167ffffffffffffffff811115610dd457610dd3610cc0565b5b610de084828501610d52565b60208301525092915050565b6000610dff610dfa84610c8f565b610c74565b90508083825260208201905060208402830185811115610e2257610e216105c7565b5b835b81811015610e6957805167ffffffffffffffff811115610e4757610e466105bd565b5b808601610e548982610d80565b85526020850194505050602081019050610e24565b5050509392505050565b600082601f830112610e8857610e876105bd565b5b8151610e98848260208601610dec565b91505092915050565b600060208284031215610eb757610eb66105b3565b5b600082015167ffffffffffffffff811115610ed557610ed46105b8565b5b610ee184828501610e73565b91505092915050565b600082825260208201905092915050565b7f554e415554484f52495a45440000000000000000000000000000000000000000600082015250565b6000610f31600c83610eea565b9150610f3c82610efb565b602082019050919050565b60006020820190508181036000830152610f6081610f24565b9050919050565b7f46756e64732077696c6c206f6e6c792062652072656c656173656420746f207460008201527f6865206f776e6572000000000000000000000000000000000000000000000000602082015250565b6000610fc3602883610eea565b9150610fce82610f67565b604082019050919050565b60006020820190508181036000830152610ff281610fb6565b9050919050565b600060608201905061100e600083018661071b565b61101b602083018561071b565b818103604083015261102d8184610878565b905094935050505056fea2646970667358221220e9d790db946e569eaf3358115f667d78de0269c52f84c93b81975c1cab91efac64736f6c63430008130033",
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
func (_Fundraiser *FundraiserCaller) GetRaisedAmounts(opts *bind.CallOpts) ([]IBankModuleCoin, error) {
	var out []interface{}
	err := _Fundraiser.contract.Call(opts, &out, "GetRaisedAmounts")

	if err != nil {
		return *new([]IBankModuleCoin), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBankModuleCoin)).(*[]IBankModuleCoin)

	return out0, err

}

// GetRaisedAmounts is a free data retrieval call binding the contract method 0xaf1d3f52.
//
// Solidity: function GetRaisedAmounts() view returns((uint256,string)[])
func (_Fundraiser *FundraiserSession) GetRaisedAmounts() ([]IBankModuleCoin, error) {
	return _Fundraiser.Contract.GetRaisedAmounts(&_Fundraiser.CallOpts)
}

// GetRaisedAmounts is a free data retrieval call binding the contract method 0xaf1d3f52.
//
// Solidity: function GetRaisedAmounts() view returns((uint256,string)[])
func (_Fundraiser *FundraiserCallerSession) GetRaisedAmounts() ([]IBankModuleCoin, error) {
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
func (_Fundraiser *FundraiserTransactor) Donate(opts *bind.TransactOpts, coins []IBankModuleCoin) (*types.Transaction, error) {
	return _Fundraiser.contract.Transact(opts, "Donate", coins)
}

// Donate is a paid mutator transaction binding the contract method 0x1ecc9652.
//
// Solidity: function Donate((uint256,string)[] coins) returns()
func (_Fundraiser *FundraiserSession) Donate(coins []IBankModuleCoin) (*types.Transaction, error) {
	return _Fundraiser.Contract.Donate(&_Fundraiser.TransactOpts, coins)
}

// Donate is a paid mutator transaction binding the contract method 0x1ecc9652.
//
// Solidity: function Donate((uint256,string)[] coins) returns()
func (_Fundraiser *FundraiserTransactorSession) Donate(coins []IBankModuleCoin) (*types.Transaction, error) {
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
