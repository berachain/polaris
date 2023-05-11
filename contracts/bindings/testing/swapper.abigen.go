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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20Module\",\"outputs\":[{\"internalType\":\"contractIERC20Module\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"getPolarisERC20\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040526269696973ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801561004657600080fd5b50608051610b65610077600039600081816101c9015281816101ef015281816102dd01526103c80152610b656000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806347e7ef241461005c578063714ba40c146100785780639d456b6214610096578063d004f0f7146100b2578063d6ece467146100ce575b600080fd5b61007660048036038101906100719190610508565b6100fe565b005b6100806101c7565b60405161008d91906105a7565b60405180910390f35b6100b060048036038101906100ab9190610627565b6101eb565b005b6100cc60048036038101906100c791906106c5565b6102d9565b005b6100e860048036038101906100e39190610705565b6103c4565b6040516100f59190610773565b60405180910390f35b60008273ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b815260040161013d939291906107ac565b6020604051808303816000875af115801561015c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610180919061081b565b9050806101c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101b9906108a5565b60405180910390fd5b505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663096b406985853333876040518663ffffffff1660e01b815260040161024e959493929190610912565b6020604051808303816000875af115801561026d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610291919061081b565b9050806102d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102ca906109d2565b60405180910390fd5b50505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663b96d8bec843333866040518563ffffffff1660e01b815260040161033a94939291906109f2565b6020604051808303816000875af1158015610359573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061037d919061081b565b9050806103bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103b690610aa9565b60405180910390fd5b505050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a333e57c84846040518363ffffffff1660e01b8152600401610421929190610ac9565b602060405180830381865afa15801561043e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104629190610b02565b905092915050565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061049f82610474565b9050919050565b6104af81610494565b81146104ba57600080fd5b50565b6000813590506104cc816104a6565b92915050565b6000819050919050565b6104e5816104d2565b81146104f057600080fd5b50565b600081359050610502816104dc565b92915050565b6000806040838503121561051f5761051e61046a565b5b600061052d858286016104bd565b925050602061053e858286016104f3565b9150509250929050565b6000819050919050565b600061056d61056861056384610474565b610548565b610474565b9050919050565b600061057f82610552565b9050919050565b600061059182610574565b9050919050565b6105a181610586565b82525050565b60006020820190506105bc6000830184610598565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f8401126105e7576105e66105c2565b5b8235905067ffffffffffffffff811115610604576106036105c7565b5b6020830191508360018202830111156106205761061f6105cc565b5b9250929050565b6000806000604084860312156106405761063f61046a565b5b600084013567ffffffffffffffff81111561065e5761065d61046f565b5b61066a868287016105d1565b9350935050602061067d868287016104f3565b9150509250925092565b600061069282610494565b9050919050565b6106a281610687565b81146106ad57600080fd5b50565b6000813590506106bf81610699565b92915050565b600080604083850312156106dc576106db61046a565b5b60006106ea858286016106b0565b92505060206106fb858286016104f3565b9150509250929050565b6000806020838503121561071c5761071b61046a565b5b600083013567ffffffffffffffff81111561073a5761073961046f565b5b610746858286016105d1565b92509250509250929050565b600061075d82610574565b9050919050565b61076d81610752565b82525050565b60006020820190506107886000830184610764565b92915050565b61079781610494565b82525050565b6107a6816104d2565b82525050565b60006060820190506107c1600083018661078e565b6107ce602083018561078e565b6107db604083018461079d565b949350505050565b60008115159050919050565b6107f8816107e3565b811461080357600080fd5b50565b600081519050610815816107ef565b92915050565b6000602082840312156108315761083061046a565b5b600061083f84828501610806565b91505092915050565b600082825260208201905092915050565b7f537761707065723a207472616e7366657246726f6d206661696c656400000000600082015250565b600061088f601c83610848565b915061089a82610859565b602082019050919050565b600060208201905081810360008301526108be81610882565b9050919050565b82818337600083830152505050565b6000601f19601f8301169050919050565b60006108f18385610848565b93506108fe8385846108c5565b610907836108d4565b840190509392505050565b6000608082019050818103600083015261092d8187896108e5565b905061093c602083018661078e565b610949604083018561078e565b610956606083018461079d565b9695505050505050565b7f537761707065723a207472616e73666572436f696e546f45524332302066616960008201527f6c65640000000000000000000000000000000000000000000000000000000000602082015250565b60006109bc602383610848565b91506109c782610960565b604082019050919050565b600060208201905081810360008301526109eb816109af565b9050919050565b6000608082019050610a076000830187610764565b610a14602083018661078e565b610a21604083018561078e565b610a2e606083018461079d565b95945050505050565b7f537761707065723a207472616e736665724552433230546f436f696e2066616960008201527f6c65640000000000000000000000000000000000000000000000000000000000602082015250565b6000610a93602383610848565b9150610a9e82610a37565b604082019050919050565b60006020820190508181036000830152610ac281610a86565b9050919050565b60006020820190508181036000830152610ae48184866108e5565b90509392505050565b600081519050610afc81610699565b92915050565b600060208284031215610b1857610b1761046a565b5b6000610b2684828501610aed565b9150509291505056fea2646970667358221220849cafc9e3543c3ba7113d7385382be522eafd4027418045de215be5e543238a64736f6c63430008130033",
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

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_Swapper *SwapperTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_Swapper *SwapperSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.Contract.Deposit(&_Swapper.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns()
func (_Swapper *SwapperTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Swapper.Contract.Deposit(&_Swapper.TransactOpts, token, amount)
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
