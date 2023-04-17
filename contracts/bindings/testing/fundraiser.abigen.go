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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Data\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"Success\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structIBankModule.Coin[]\",\"name\":\"coins\",\"type\":\"tuple[]\"}],\"name\":\"Donate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetRaisedAmounts\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structIBankModule.Coin[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"contractIBankModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawDonations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a0604052734381dc2ab14285160c808659aee005d51255add773ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801561005757600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550608051610e076100c86000396000818161011d015281816101c40152818161020e01526103400152610e076000f3fe60806040526004361061004e5760003560e01c80631ecc96521461005a57806376cdb03b146100835780638da5cb5b146100ae578063af1d3f52146100d9578063ce1b088a1461010457610055565b3661005557005b600080fd5b34801561006657600080fd5b50610081600480360381019061007c9190610481565b61011b565b005b34801561008f57600080fd5b506100986101c2565b6040516100a5919061054d565b60405180910390f35b3480156100ba57600080fd5b506100c36101e6565b6040516100d09190610589565b60405180910390f35b3480156100e557600080fd5b506100ee61020a565b6040516100fb919061074c565b60405180910390f35b34801561011057600080fd5b506101196102b0565b005b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166384404811333085856040518563ffffffff1660e01b815260040161017a9493929190610975565b6020604051808303816000875af1158015610199573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101bd91906109ed565b505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60607f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c53d6ce1306040518263ffffffff1660e01b81526004016102659190610589565b600060405180830381865afa158015610282573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906102ab9190610ca7565b905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461033e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033590610d73565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663844048113060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff166103a461020a565b6040518463ffffffff1660e01b81526004016103c293929190610d93565b6020604051808303816000875af11580156103e1573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061040591906109ed565b50565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f8401126104415761044061041c565b5b8235905067ffffffffffffffff81111561045e5761045d610421565b5b60208301915083602082028301111561047a57610479610426565b5b9250929050565b6000806020838503121561049857610497610412565b5b600083013567ffffffffffffffff8111156104b6576104b5610417565b5b6104c28582860161042b565b92509250509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061051361050e610509846104ce565b6104ee565b6104ce565b9050919050565b6000610525826104f8565b9050919050565b60006105378261051a565b9050919050565b6105478161052c565b82525050565b6000602082019050610562600083018461053e565b92915050565b6000610573826104ce565b9050919050565b61058381610568565b82525050565b600060208201905061059e600083018461057a565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000819050919050565b6105e3816105d0565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610623578082015181840152602081019050610608565b60008484015250505050565b6000601f19601f8301169050919050565b600061064b826105e9565b61065581856105f4565b9350610665818560208601610605565b61066e8161062f565b840191505092915050565b600060408301600083015161069160008601826105da565b50602083015184820360208601526106a98282610640565b9150508091505092915050565b60006106c28383610679565b905092915050565b6000602082019050919050565b60006106e2826105a4565b6106ec81856105af565b9350836020820285016106fe856105c0565b8060005b8581101561073a578484038952815161071b85826106b6565b9450610726836106ca565b925060208a01995050600181019050610702565b50829750879550505050505092915050565b6000602082019050818103600083015261076681846106d7565b905092915050565b6000819050919050565b610781816105d0565b811461078c57600080fd5b50565b60008135905061079e81610778565b92915050565b60006107b3602084018461078f565b905092915050565b600080fd5b600080fd5b600080fd5b600080833560016020038436030381126107e7576107e66107c5565b5b83810192508235915060208301925067ffffffffffffffff82111561080f5761080e6107bb565b5b600182023603831315610825576108246107c0565b5b509250929050565b82818337600083830152505050565b600061084883856105f4565b935061085583858461082d565b61085e8361062f565b840190509392505050565b60006040830161087c60008401846107a4565b61088960008601826105da565b5061089760208401846107ca565b85830360208701526108aa83828461083c565b925050508091505092915050565b60006108c48383610869565b905092915050565b6000823560016040038336030381126108e8576108e76107c5565b5b82810191505092915050565b6000602082019050919050565b600061090d83856105af565b93508360208402850161091f8461076e565b8060005b8781101561096357848403895261093a82846108cc565b61094485826108b8565b945061094f836108f4565b925060208a01995050600181019050610923565b50829750879450505050509392505050565b600060608201905061098a600083018761057a565b610997602083018661057a565b81810360408301526109aa818486610901565b905095945050505050565b60008115159050919050565b6109ca816109b5565b81146109d557600080fd5b50565b6000815190506109e7816109c1565b92915050565b600060208284031215610a0357610a02610412565b5b6000610a11848285016109d8565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610a528261062f565b810181811067ffffffffffffffff82111715610a7157610a70610a1a565b5b80604052505050565b6000610a84610408565b9050610a908282610a49565b919050565b600067ffffffffffffffff821115610ab057610aaf610a1a565b5b602082029050602081019050919050565b600080fd5b600080fd5b600081519050610ada81610778565b92915050565b600080fd5b600067ffffffffffffffff821115610b0057610aff610a1a565b5b610b098261062f565b9050602081019050919050565b6000610b29610b2484610ae5565b610a7a565b905082815260208101848484011115610b4557610b44610ae0565b5b610b50848285610605565b509392505050565b600082601f830112610b6d57610b6c61041c565b5b8151610b7d848260208601610b16565b91505092915050565b600060408284031215610b9c57610b9b610ac1565b5b610ba66040610a7a565b90506000610bb684828501610acb565b600083015250602082015167ffffffffffffffff811115610bda57610bd9610ac6565b5b610be684828501610b58565b60208301525092915050565b6000610c05610c0084610a95565b610a7a565b90508083825260208201905060208402830185811115610c2857610c27610426565b5b835b81811015610c6f57805167ffffffffffffffff811115610c4d57610c4c61041c565b5b808601610c5a8982610b86565b85526020850194505050602081019050610c2a565b5050509392505050565b600082601f830112610c8e57610c8d61041c565b5b8151610c9e848260208601610bf2565b91505092915050565b600060208284031215610cbd57610cbc610412565b5b600082015167ffffffffffffffff811115610cdb57610cda610417565b5b610ce784828501610c79565b91505092915050565b600082825260208201905092915050565b7f46756e64732077696c6c206f6e6c792062652072656c656173656420746f207460008201527f6865206f776e6572000000000000000000000000000000000000000000000000602082015250565b6000610d5d602883610cf0565b9150610d6882610d01565b604082019050919050565b60006020820190508181036000830152610d8c81610d50565b9050919050565b6000606082019050610da8600083018661057a565b610db5602083018561057a565b8181036040830152610dc781846106d7565b905094935050505056fea26469706673582212207587ae464e4d7053fe712e6c72c8af1ebf1df4f3300d982c14616d3293342b7864736f6c63430008130033",
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
