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
	Bin: "0x60a06040526269696973ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801561004657600080fd5b506080516109466100756000396000818160d90152818160ff015281816101eb01526102d401526109466000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063714ba40c146100515780639d456b621461006f578063d004f0f71461008b578063d6ece467146100a7575b600080fd5b6100596100d7565b60405161006691906103f5565b60405180910390f35b610089600480360381019061008491906104b5565b6100fb565b005b6100a560048036038101906100a09190610565565b6101e7565b005b6100c160048036038101906100bc91906105a5565b6102d0565b6040516100ce9190610613565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000081565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663423eb10b858533866040518563ffffffff1660e01b815260040161015c94939291906106aa565b6020604051808303816000875af115801561017b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061019f9190610722565b9050806101e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101d8906107c1565b60405180910390fd5b50505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166377f423688433856040518463ffffffff1660e01b8152600401610246939291906107e1565b6020604051808303816000875af1158015610265573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102899190610722565b9050806102cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102c29061088a565b60405180910390fd5b505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a333e57c84846040518363ffffffff1660e01b815260040161032d9291906108aa565b602060405180830381865afa15801561034a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061036e91906108e3565b905092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006103bb6103b66103b184610376565b610396565b610376565b9050919050565b60006103cd826103a0565b9050919050565b60006103df826103c2565b9050919050565b6103ef816103d4565b82525050565b600060208201905061040a60008301846103e6565b92915050565b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b60008083601f84011261043f5761043e61041a565b5b8235905067ffffffffffffffff81111561045c5761045b61041f565b5b60208301915083600182028301111561047857610477610424565b5b9250929050565b6000819050919050565b6104928161047f565b811461049d57600080fd5b50565b6000813590506104af81610489565b92915050565b6000806000604084860312156104ce576104cd610410565b5b600084013567ffffffffffffffff8111156104ec576104eb610415565b5b6104f886828701610429565b9350935050602061050b868287016104a0565b9150509250925092565b600061052082610376565b9050919050565b600061053282610515565b9050919050565b61054281610527565b811461054d57600080fd5b50565b60008135905061055f81610539565b92915050565b6000806040838503121561057c5761057b610410565b5b600061058a85828601610550565b925050602061059b858286016104a0565b9150509250929050565b600080602083850312156105bc576105bb610410565b5b600083013567ffffffffffffffff8111156105da576105d9610415565b5b6105e685828601610429565b92509250509250929050565b60006105fd826103c2565b9050919050565b61060d816105f2565b82525050565b60006020820190506106286000830184610604565b92915050565b600082825260208201905092915050565b82818337600083830152505050565b6000601f19601f8301169050919050565b600061066b838561062e565b935061067883858461063f565b6106818361064e565b840190509392505050565b61069581610515565b82525050565b6106a48161047f565b82525050565b600060608201905081810360008301526106c581868861065f565b90506106d4602083018561068c565b6106e1604083018461069b565b95945050505050565b60008115159050919050565b6106ff816106ea565b811461070a57600080fd5b50565b60008151905061071c816106f6565b92915050565b60006020828403121561073857610737610410565b5b60006107468482850161070d565b91505092915050565b7f537761707065723a20636f6e76657274436f696e546f4552433230206661696c60008201527f6564000000000000000000000000000000000000000000000000000000000000602082015250565b60006107ab60228361062e565b91506107b68261074f565b604082019050919050565b600060208201905081810360008301526107da8161079e565b9050919050565b60006060820190506107f66000830186610604565b610803602083018561068c565b610810604083018461069b565b949350505050565b7f537761707065723a20636f6e766572744552433230546f436f696e206661696c60008201527f6564000000000000000000000000000000000000000000000000000000000000602082015250565b600061087460228361062e565b915061087f82610818565b604082019050919050565b600060208201905081810360008301526108a381610867565b9050919050565b600060208201905081810360008301526108c581848661065f565b90509392505050565b6000815190506108dd81610539565b92915050565b6000602082840312156108f9576108f8610410565b5b6000610907848285016108ce565b9150509291505056fea2646970667358221220b7df7245c0721ec8ed2e48b7569b663153dbdcb97d0839728ef82379b941f08c64736f6c63430008130033",
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
