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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Data\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"Success\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structIBankModule.Coin[]\",\"name\":\"coins\",\"type\":\"tuple[]\"}],\"name\":\"Donate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetRaisedAmounts\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structIBankModule.Coin[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"contractIBankModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawDonations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a0604052734381dc2ab14285160c808659aee005d51255add773ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801561005757600080fd5b5033806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3506080516110b961012560003960008181610151015281816101f801528181610242015261040201526110b96000f3fe6080604052600436106100595760003560e01c80631ecc96521461006557806376cdb03b1461008e5780638da5cb5b146100b9578063af1d3f52146100e4578063ce1b088a1461010f578063f2fde38b1461012657610060565b3661006057005b600080fd5b34801561007157600080fd5b5061008c6004803603810190610087919061066e565b61014f565b005b34801561009a57600080fd5b506100a36101f6565b6040516100b0919061073a565b60405180910390f35b3480156100c557600080fd5b506100ce61021a565b6040516100db9190610776565b60405180910390f35b3480156100f057600080fd5b506100f961023e565b6040516101069190610939565b60405180910390f35b34801561011b57600080fd5b506101246102e4565b005b34801561013257600080fd5b5061014d60048036038101906101489190610987565b6104ca565b005b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166384404811333085856040518563ffffffff1660e01b81526004016101ae9493929190610bbb565b6020604051808303816000875af11580156101cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101f19190610c33565b505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60607f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c53d6ce1306040518263ffffffff1660e01b81526004016102999190610776565b600060405180830381865afa1580156102b6573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906102df9190610eed565b905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610372576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161036990610f93565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610400576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f790611025565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663844048113060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1661046661023e565b6040518463ffffffff1660e01b815260040161048493929190611045565b6020604051808303816000875af11580156104a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104c79190610c33565b50565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610558576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054f90610f93565b60405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a350565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f84011261062e5761062d610609565b5b8235905067ffffffffffffffff81111561064b5761064a61060e565b5b60208301915083602082028301111561066757610666610613565b5b9250929050565b60008060208385031215610685576106846105ff565b5b600083013567ffffffffffffffff8111156106a3576106a2610604565b5b6106af85828601610618565b92509250509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006107006106fb6106f6846106bb565b6106db565b6106bb565b9050919050565b6000610712826106e5565b9050919050565b600061072482610707565b9050919050565b61073481610719565b82525050565b600060208201905061074f600083018461072b565b92915050565b6000610760826106bb565b9050919050565b61077081610755565b82525050565b600060208201905061078b6000830184610767565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000819050919050565b6107d0816107bd565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b838110156108105780820151818401526020810190506107f5565b60008484015250505050565b6000601f19601f8301169050919050565b6000610838826107d6565b61084281856107e1565b93506108528185602086016107f2565b61085b8161081c565b840191505092915050565b600060408301600083015161087e60008601826107c7565b5060208301518482036020860152610896828261082d565b9150508091505092915050565b60006108af8383610866565b905092915050565b6000602082019050919050565b60006108cf82610791565b6108d9818561079c565b9350836020820285016108eb856107ad565b8060005b85811015610927578484038952815161090885826108a3565b9450610913836108b7565b925060208a019950506001810190506108ef565b50829750879550505050505092915050565b6000602082019050818103600083015261095381846108c4565b905092915050565b61096481610755565b811461096f57600080fd5b50565b6000813590506109818161095b565b92915050565b60006020828403121561099d5761099c6105ff565b5b60006109ab84828501610972565b91505092915050565b6000819050919050565b6109c7816107bd565b81146109d257600080fd5b50565b6000813590506109e4816109be565b92915050565b60006109f960208401846109d5565b905092915050565b600080fd5b600080fd5b600080fd5b60008083356001602003843603038112610a2d57610a2c610a0b565b5b83810192508235915060208301925067ffffffffffffffff821115610a5557610a54610a01565b5b600182023603831315610a6b57610a6a610a06565b5b509250929050565b82818337600083830152505050565b6000610a8e83856107e1565b9350610a9b838584610a73565b610aa48361081c565b840190509392505050565b600060408301610ac260008401846109ea565b610acf60008601826107c7565b50610add6020840184610a10565b8583036020870152610af0838284610a82565b925050508091505092915050565b6000610b0a8383610aaf565b905092915050565b600082356001604003833603038112610b2e57610b2d610a0b565b5b82810191505092915050565b6000602082019050919050565b6000610b53838561079c565b935083602084028501610b65846109b4565b8060005b87811015610ba9578484038952610b808284610b12565b610b8a8582610afe565b9450610b9583610b3a565b925060208a01995050600181019050610b69565b50829750879450505050509392505050565b6000606082019050610bd06000830187610767565b610bdd6020830186610767565b8181036040830152610bf0818486610b47565b905095945050505050565b60008115159050919050565b610c1081610bfb565b8114610c1b57600080fd5b50565b600081519050610c2d81610c07565b92915050565b600060208284031215610c4957610c486105ff565b5b6000610c5784828501610c1e565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610c988261081c565b810181811067ffffffffffffffff82111715610cb757610cb6610c60565b5b80604052505050565b6000610cca6105f5565b9050610cd68282610c8f565b919050565b600067ffffffffffffffff821115610cf657610cf5610c60565b5b602082029050602081019050919050565b600080fd5b600080fd5b600081519050610d20816109be565b92915050565b600080fd5b600067ffffffffffffffff821115610d4657610d45610c60565b5b610d4f8261081c565b9050602081019050919050565b6000610d6f610d6a84610d2b565b610cc0565b905082815260208101848484011115610d8b57610d8a610d26565b5b610d968482856107f2565b509392505050565b600082601f830112610db357610db2610609565b5b8151610dc3848260208601610d5c565b91505092915050565b600060408284031215610de257610de1610d07565b5b610dec6040610cc0565b90506000610dfc84828501610d11565b600083015250602082015167ffffffffffffffff811115610e2057610e1f610d0c565b5b610e2c84828501610d9e565b60208301525092915050565b6000610e4b610e4684610cdb565b610cc0565b90508083825260208201905060208402830185811115610e6e57610e6d610613565b5b835b81811015610eb557805167ffffffffffffffff811115610e9357610e92610609565b5b808601610ea08982610dcc565b85526020850194505050602081019050610e70565b5050509392505050565b600082601f830112610ed457610ed3610609565b5b8151610ee4848260208601610e38565b91505092915050565b600060208284031215610f0357610f026105ff565b5b600082015167ffffffffffffffff811115610f2157610f20610604565b5b610f2d84828501610ebf565b91505092915050565b600082825260208201905092915050565b7f554e415554484f52495a45440000000000000000000000000000000000000000600082015250565b6000610f7d600c83610f36565b9150610f8882610f47565b602082019050919050565b60006020820190508181036000830152610fac81610f70565b9050919050565b7f46756e64732077696c6c206f6e6c792062652072656c656173656420746f207460008201527f6865206f776e6572000000000000000000000000000000000000000000000000602082015250565b600061100f602883610f36565b915061101a82610fb3565b604082019050919050565b6000602082019050818103600083015261103e81611002565b9050919050565b600060608201905061105a6000830186610767565b6110676020830185610767565b818103604083015261107981846108c4565b905094935050505056fea26469706673582212201c40728a1a09b731a9d7918159dafcb032cd8309003587c0bb7de571b773184664736f6c63430008130033",
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

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Fundraiser *FundraiserTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fundraiser.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Fundraiser *FundraiserSession) Receive() (*types.Transaction, error) {
	return _Fundraiser.Contract.Receive(&_Fundraiser.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Fundraiser *FundraiserTransactorSession) Receive() (*types.Transaction, error) {
	return _Fundraiser.Contract.Receive(&_Fundraiser.TransactOpts)
}

// FundraiserDataIterator is returned from FilterData and is used to iterate over the raw logs and unpacked data for Data events raised by the Fundraiser contract.
type FundraiserDataIterator struct {
	Event *FundraiserData // Event containing the contract specifics and raw log

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
func (it *FundraiserDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FundraiserData)
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
		it.Event = new(FundraiserData)
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
func (it *FundraiserDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FundraiserDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FundraiserData represents a Data event raised by the Fundraiser contract.
type FundraiserData struct {
	Data []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterData is a free log retrieval operation binding the contract event 0x0b76c48be4e2908f4c9d4eabaf7538e91577fd9ae26db46693fa8d861c6a42fb.
//
// Solidity: event Data(bytes data)
func (_Fundraiser *FundraiserFilterer) FilterData(opts *bind.FilterOpts) (*FundraiserDataIterator, error) {

	logs, sub, err := _Fundraiser.contract.FilterLogs(opts, "Data")
	if err != nil {
		return nil, err
	}
	return &FundraiserDataIterator{contract: _Fundraiser.contract, event: "Data", logs: logs, sub: sub}, nil
}

// WatchData is a free log subscription operation binding the contract event 0x0b76c48be4e2908f4c9d4eabaf7538e91577fd9ae26db46693fa8d861c6a42fb.
//
// Solidity: event Data(bytes data)
func (_Fundraiser *FundraiserFilterer) WatchData(opts *bind.WatchOpts, sink chan<- *FundraiserData) (event.Subscription, error) {

	logs, sub, err := _Fundraiser.contract.WatchLogs(opts, "Data")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FundraiserData)
				if err := _Fundraiser.contract.UnpackLog(event, "Data", log); err != nil {
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

// ParseData is a log parse operation binding the contract event 0x0b76c48be4e2908f4c9d4eabaf7538e91577fd9ae26db46693fa8d861c6a42fb.
//
// Solidity: event Data(bytes data)
func (_Fundraiser *FundraiserFilterer) ParseData(log types.Log) (*FundraiserData, error) {
	event := new(FundraiserData)
	if err := _Fundraiser.contract.UnpackLog(event, "Data", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// FundraiserSuccessIterator is returned from FilterSuccess and is used to iterate over the raw logs and unpacked data for Success events raised by the Fundraiser contract.
type FundraiserSuccessIterator struct {
	Event *FundraiserSuccess // Event containing the contract specifics and raw log

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
func (it *FundraiserSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FundraiserSuccess)
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
		it.Event = new(FundraiserSuccess)
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
func (it *FundraiserSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FundraiserSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FundraiserSuccess represents a Success event raised by the Fundraiser contract.
type FundraiserSuccess struct {
	Success bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSuccess is a free log retrieval operation binding the contract event 0x3b0a8ddef325df2bfdfa6b430ae4c8421841cd135bfa8fb5e432f200787520bb.
//
// Solidity: event Success(bool indexed success)
func (_Fundraiser *FundraiserFilterer) FilterSuccess(opts *bind.FilterOpts, success []bool) (*FundraiserSuccessIterator, error) {

	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}

	logs, sub, err := _Fundraiser.contract.FilterLogs(opts, "Success", successRule)
	if err != nil {
		return nil, err
	}
	return &FundraiserSuccessIterator{contract: _Fundraiser.contract, event: "Success", logs: logs, sub: sub}, nil
}

// WatchSuccess is a free log subscription operation binding the contract event 0x3b0a8ddef325df2bfdfa6b430ae4c8421841cd135bfa8fb5e432f200787520bb.
//
// Solidity: event Success(bool indexed success)
func (_Fundraiser *FundraiserFilterer) WatchSuccess(opts *bind.WatchOpts, sink chan<- *FundraiserSuccess, success []bool) (event.Subscription, error) {

	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}

	logs, sub, err := _Fundraiser.contract.WatchLogs(opts, "Success", successRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FundraiserSuccess)
				if err := _Fundraiser.contract.UnpackLog(event, "Success", log); err != nil {
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

// ParseSuccess is a log parse operation binding the contract event 0x3b0a8ddef325df2bfdfa6b430ae4c8421841cd135bfa8fb5e432f200787520bb.
//
// Solidity: event Success(bool indexed success)
func (_Fundraiser *FundraiserFilterer) ParseSuccess(log types.Log) (*FundraiserSuccess, error) {
	event := new(FundraiserSuccess)
	if err := _Fundraiser.contract.UnpackLog(event, "Success", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
