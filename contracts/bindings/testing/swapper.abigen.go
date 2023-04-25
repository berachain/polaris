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

// SwapperMetaData contains all meta data concerning the Swapper contract.
var SwapperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"erc20Module\",\"outputs\":[{\"internalType\":\"contractIERC20Module\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"getPolarisERC20\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040526269696973ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801561004657600080fd5b506080516109666100756000396000818160d90152818160ff015281816101ed01526102d801526109666000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063714ba40c146100515780639d456b621461006f578063d004f0f71461008b578063d6ece467146100a7575b600080fd5b6100596100d7565b60405161006691906103f9565b60405180910390f35b610089600480360381019061008491906104b9565b6100fb565b005b6100a560048036038101906100a09190610569565b6101e9565b005b6100c160048036038101906100bc91906105a9565b6102d4565b6040516100ce9190610617565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000081565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166310e0f72585853333876040518663ffffffff1660e01b815260040161015e9594939291906106ae565b6020604051808303816000875af115801561017d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101a19190610734565b9050806101e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101da906107d3565b60405180910390fd5b50505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166324b8f0fe843333866040518563ffffffff1660e01b815260040161024a94939291906107f3565b6020604051808303816000875af1158015610269573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061028d9190610734565b9050806102cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102c6906108aa565b60405180910390fd5b505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a333e57c84846040518363ffffffff1660e01b81526004016103319291906108ca565b602060405180830381865afa15801561034e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103729190610903565b905092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006103bf6103ba6103b58461037a565b61039a565b61037a565b9050919050565b60006103d1826103a4565b9050919050565b60006103e3826103c6565b9050919050565b6103f3816103d8565b82525050565b600060208201905061040e60008301846103ea565b92915050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f8401126104435761044261041e565b5b8235905067ffffffffffffffff8111156104605761045f610423565b5b60208301915083600182028301111561047c5761047b610428565b5b9250929050565b6000819050919050565b61049681610483565b81146104a157600080fd5b50565b6000813590506104b38161048d565b92915050565b6000806000604084860312156104d2576104d1610414565b5b600084013567ffffffffffffffff8111156104f0576104ef610419565b5b6104fc8682870161042d565b9350935050602061050f868287016104a4565b9150509250925092565b60006105248261037a565b9050919050565b600061053682610519565b9050919050565b6105468161052b565b811461055157600080fd5b50565b6000813590506105638161053d565b92915050565b600080604083850312156105805761057f610414565b5b600061058e85828601610554565b925050602061059f858286016104a4565b9150509250929050565b600080602083850312156105c0576105bf610414565b5b600083013567ffffffffffffffff8111156105de576105dd610419565b5b6105ea8582860161042d565b92509250509250929050565b6000610601826103c6565b9050919050565b610611816105f6565b82525050565b600060208201905061062c6000830184610608565b92915050565b600082825260208201905092915050565b82818337600083830152505050565b6000601f19601f8301169050919050565b600061066f8385610632565b935061067c838584610643565b61068583610652565b840190509392505050565b61069981610519565b82525050565b6106a881610483565b82525050565b600060808201905081810360008301526106c9818789610663565b90506106d86020830186610690565b6106e56040830185610690565b6106f2606083018461069f565b9695505050505050565b60008115159050919050565b610711816106fc565b811461071c57600080fd5b50565b60008151905061072e81610708565b92915050565b60006020828403121561074a57610749610414565b5b60006107588482850161071f565b91505092915050565b7f537761707065723a20636f6e76657274436f696e546f4552433230206661696c60008201527f6564000000000000000000000000000000000000000000000000000000000000602082015250565b60006107bd602283610632565b91506107c882610761565b604082019050919050565b600060208201905081810360008301526107ec816107b0565b9050919050565b60006080820190506108086000830187610608565b6108156020830186610690565b6108226040830185610690565b61082f606083018461069f565b95945050505050565b7f537761707065723a20636f6e766572744552433230546f436f696e206661696c60008201527f6564000000000000000000000000000000000000000000000000000000000000602082015250565b6000610894602283610632565b915061089f82610838565b604082019050919050565b600060208201905081810360008301526108c381610887565b9050919050565b600060208201905081810360008301526108e5818486610663565b90509392505050565b6000815190506108fd8161053d565b92915050565b60006020828403121561091957610918610414565b5b6000610927848285016108ee565b9150509291505056fea2646970667358221220eab45f686d4fb4fedba2ad7d069ab9a82aa81a5b0cc45f26bfd14e55fa90c52964736f6c63430008130033",
}

// SwapperABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapperMetaData.ABI instead.
var SwapperABI = SwapperMetaData.ABI

// SwapperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SwapperMetaData.Bin instead.
var SwapperBin = SwapperMetaData.Bin

// DeploySwapper deploys a new Ethereum contract, binding an instance of Swapper to it.
func DeploySwapper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Swapper, error) {
	parsed, err := SwapperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SwapperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Swapper{SwapperCaller: SwapperCaller{contract: contract}, SwapperTransactor: SwapperTransactor{contract: contract}, SwapperFilterer: SwapperFilterer{contract: contract}}, nil
}

// Swapper is an auto generated Go binding around an Ethereum contract.
type Swapper struct {
	SwapperCaller     // Read-only binding to the contract
	SwapperTransactor // Write-only binding to the contract
	SwapperFilterer   // Log filterer for contract events
}

// SwapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapperSession struct {
	Contract     *Swapper          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapperCallerSession struct {
	Contract *SwapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SwapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapperTransactorSession struct {
	Contract     *SwapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SwapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapperRaw struct {
	Contract *Swapper // Generic contract binding to access the raw methods on
}

// SwapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapperCallerRaw struct {
	Contract *SwapperCaller // Generic read-only contract binding to access the raw methods on
}

// SwapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapperTransactorRaw struct {
	Contract *SwapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapper creates a new instance of Swapper, bound to a specific deployed contract.
func NewSwapper(address common.Address, backend bind.ContractBackend) (*Swapper, error) {
	contract, err := bindSwapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swapper{SwapperCaller: SwapperCaller{contract: contract}, SwapperTransactor: SwapperTransactor{contract: contract}, SwapperFilterer: SwapperFilterer{contract: contract}}, nil
}

// NewSwapperCaller creates a new read-only instance of Swapper, bound to a specific deployed contract.
func NewSwapperCaller(address common.Address, caller bind.ContractCaller) (*SwapperCaller, error) {
	contract, err := bindSwapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapperCaller{contract: contract}, nil
}

// NewSwapperTransactor creates a new write-only instance of Swapper, bound to a specific deployed contract.
func NewSwapperTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapperTransactor, error) {
	contract, err := bindSwapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapperTransactor{contract: contract}, nil
}

// NewSwapperFilterer creates a new log filterer instance of Swapper, bound to a specific deployed contract.
func NewSwapperFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapperFilterer, error) {
	contract, err := bindSwapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapperFilterer{contract: contract}, nil
}

// bindSwapper binds a generic wrapper to an already deployed contract.
func bindSwapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SwapperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swapper *SwapperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swapper.Contract.SwapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swapper *SwapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapper.Contract.SwapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swapper *SwapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swapper.Contract.SwapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swapper *SwapperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swapper *SwapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swapper *SwapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swapper.Contract.contract.Transact(opts, method, params...)
}

// Erc20Module is a free data retrieval call binding the contract method 0x714ba40c.
//
// Solidity: function erc20Module() view returns(address)
func (_Swapper *SwapperCaller) Erc20Module(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "erc20Module")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20Module is a free data retrieval call binding the contract method 0x714ba40c.
//
// Solidity: function erc20Module() view returns(address)
func (_Swapper *SwapperSession) Erc20Module() (common.Address, error) {
	return _Swapper.Contract.Erc20Module(&_Swapper.CallOpts)
}

// Erc20Module is a free data retrieval call binding the contract method 0x714ba40c.
//
// Solidity: function erc20Module() view returns(address)
func (_Swapper *SwapperCallerSession) Erc20Module() (common.Address, error) {
	return _Swapper.Contract.Erc20Module(&_Swapper.CallOpts)
}

// GetPolarisERC20 is a free data retrieval call binding the contract method 0xd6ece467.
//
// Solidity: function getPolarisERC20(string denom) view returns(address)
func (_Swapper *SwapperCaller) GetPolarisERC20(opts *bind.CallOpts, denom string) (common.Address, error) {
	var out []interface{}
	err := _Swapper.contract.Call(opts, &out, "getPolarisERC20", denom)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPolarisERC20 is a free data retrieval call binding the contract method 0xd6ece467.
//
// Solidity: function getPolarisERC20(string denom) view returns(address)
func (_Swapper *SwapperSession) GetPolarisERC20(denom string) (common.Address, error) {
	return _Swapper.Contract.GetPolarisERC20(&_Swapper.CallOpts, denom)
}

// GetPolarisERC20 is a free data retrieval call binding the contract method 0xd6ece467.
//
// Solidity: function getPolarisERC20(string denom) view returns(address)
func (_Swapper *SwapperCallerSession) GetPolarisERC20(denom string) (common.Address, error) {
	return _Swapper.Contract.GetPolarisERC20(&_Swapper.CallOpts, denom)
}

// Swap is a paid mutator transaction binding the contract method 0x9d456b62.
//
// Solidity: function swap(string denom, uint256 amount) returns()
func (_Swapper *SwapperTransactor) Swap(opts *bind.TransactOpts, denom string, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "swap", denom, amount)
}

// Swap is a paid mutator transaction binding the contract method 0x9d456b62.
//
// Solidity: function swap(string denom, uint256 amount) returns()
func (_Swapper *SwapperSession) Swap(denom string, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.Contract.Swap(&_Swapper.TransactOpts, denom, amount)
}

// Swap is a paid mutator transaction binding the contract method 0x9d456b62.
//
// Solidity: function swap(string denom, uint256 amount) returns()
func (_Swapper *SwapperTransactorSession) Swap(denom string, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.Contract.Swap(&_Swapper.TransactOpts, denom, amount)
}

// Swap0 is a paid mutator transaction binding the contract method 0xd004f0f7.
//
// Solidity: function swap(address token, uint256 amount) returns()
func (_Swapper *SwapperTransactor) Swap0(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "swap0", token, amount)
}

// Swap0 is a paid mutator transaction binding the contract method 0xd004f0f7.
//
// Solidity: function swap(address token, uint256 amount) returns()
func (_Swapper *SwapperSession) Swap0(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.Contract.Swap0(&_Swapper.TransactOpts, token, amount)
}

// Swap0 is a paid mutator transaction binding the contract method 0xd004f0f7.
//
// Solidity: function swap(address token, uint256 amount) returns()
func (_Swapper *SwapperTransactorSession) Swap0(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.Contract.Swap0(&_Swapper.TransactOpts, token, amount)
}
