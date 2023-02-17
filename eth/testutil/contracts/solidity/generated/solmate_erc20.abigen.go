// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package generated

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
)

// SolmateERC20MetaData contains all meta data concerning the SolmateERC20 contract.
var SolmateERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b50604051806040016040528060058152602001642a37b5b2b760d91b81525060405180604001604052806002815260200161544b60f01b815250601282600090816200005e9190620001d1565b5060016200006d8382620001d1565b5060ff81166080524660a0526200008362000090565b60c052506200031b915050565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6000604051620000c491906200029d565b6040805191829003822060208301939093528101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200015757607f821691505b6020821081036200017857634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620001cc57600081815260208120601f850160051c81016020861015620001a75750805b601f850160051c820191505b81811015620001c857828155600101620001b3565b5050505b505050565b81516001600160401b03811115620001ed57620001ed6200012c565b6200020581620001fe845462000142565b846200017e565b602080601f8311600181146200023d5760008415620002245750858301515b600019600386901b1c1916600185901b178555620001c8565b600085815260208120601f198616915b828110156200026e578886015182559484019460019091019084016200024d565b50858210156200028d5787850151600019600388901b0f8161c191681555b5050505050600190811b01905550565b6000808354620002ad8162000142565b60018281168015620002c85760018114620002de576200030f565b60ff19841687528215158302870194506200030f565b8760005260208060002060005b85811015620003065781548a820152908401908201620002eb565b50505082870194505b50929695505050505050565b60805160a05160c051610b3a6200034b60003960006104540152600061041f015260006101440152610b3a6000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c806340c10f191161008c57806395d89b411161006657806395d89b41146101d5578063a9059cbb146101dd578063d505accf146101f0578063dd62ed3e1461020357600080fd5b806340c10f191461018057806370a08231146101955780637ecebe00146101b557600080fd5b806306fdde03146100d4578063095ea7b3146100f257806318160ddd1461011557806323b872dd1461012c578063313ce5671461013f5780633644e51514610178575b600080fd5b6100dc61022e565b6040516100e99190610857565b60405180910390f35b6101056101003660046108c1565b6102bc565b60405190151581526020016100e9565b61011e60025481565b6040519081526020016100e9565b61010561013a3660046108eb565b610329565b6101667f000000000000000000000000000000000000000000000000000000000000000081565b60405160ff90911681526020016100e9565b61011e61041b565b61019361018e3660046108c1565b610476565b005b61011e6101a3366004610927565b60036020526000908152604090205481565b61011e6101c3366004610927565b60056020526000908152604090205481565b6100dc610484565b6101056101eb3660046108c1565b610491565b6101936101fe366004610949565b610509565b61011e6102113660046109bc565b600460209081526000928352604080842090915290825290205481565b6000805461023b906109ef565b80601f0160208091040260200160405190810160405280929190818152602001828054610267906109ef565b80156102b45780601f10610289576101008083540402835291602001916102b4565b820191906000526020600020905b81548152906001019060200180831161029757829003601f168201915b505050505081565b3360008181526004602090815260408083206001600160a01b038716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925906103179086815260200190565b60405180910390a35060015b92915050565b6001600160a01b03831660009081526004602090815260408083203384529091528120546000198114610385576103608382610a3f565b6001600160a01b03861660009081526004602090815260408083203384529091529020555b6001600160a01b038516600090815260036020526040812080548592906103ad908490610a3f565b90915550506001600160a01b03808516600081815260036020526040908190208054870190555190918716907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906104089087815260200190565b60405180910390a3506001949350505050565b60007f000000000000000000000000000000000000000000000000000000000000000046146104515761044c610752565b905090565b507f000000000000000000000000000000000000000000000000000000000000000090565b61048082826107ec565b5050565b6001805461023b906109ef565b336000908152600360205260408120805483919083906104b2908490610a3f565b90915550506001600160a01b038316600081815260036020526040908190208054850190555133907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906103179086815260200190565b4284101561055e5760405162461bcd60e51b815260206004820152601760248201527f5045524d49545f444541444c494e455f4558504952454400000000000000000060448201526064015b60405180910390fd5b6000600161056a61041b565b6001600160a01b038a811660008181526005602090815260409182902080546001810190915582517f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98184015280840194909452938d166060840152608083018c905260a083019390935260c08083018b90528151808403909101815260e08301909152805192019190912061190160f01b6101008301526101028201929092526101228101919091526101420160408051601f198184030181528282528051602091820120600084529083018083525260ff871690820152606081018590526080810184905260a0016020604051602081039080840390855afa158015610676573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116158015906106ac5750876001600160a01b0316816001600160a01b0316145b6106e95760405162461bcd60e51b815260206004820152600e60248201526d24a72b20a624a22fa9a4a3a722a960911b6044820152606401610555565b6001600160a01b0390811660009081526004602090815260408083208a8516808552908352928190208990555188815291928a16917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a350505050505050565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60006040516107849190610a52565b6040805191829003822060208301939093528101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc660608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b80600260008282546107fe9190610af1565b90915550506001600160a01b0382166000818152600360209081526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b600060208083528351808285015260005b8181101561088457858101830151858201604001528201610868565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b03811681146108bc57600080fd5b919050565b600080604083850312156108d457600080fd5b6108dd836108a5565b946020939093013593505050565b60008060006060848603121561090057600080fd5b610909846108a5565b9250610917602085016108a5565b9150604084013590509250925092565b60006020828403121561093957600080fd5b610942826108a5565b9392505050565b600080600080600080600060e0888a03121561096457600080fd5b61096d886108a5565b965061097b602089016108a5565b95506040880135945060608801359350608088013560ff8116811461099f57600080fd5b9699959850939692959460a0840135945060c09093013592915050565b600080604083850312156109cf57600080fd5b6109d8836108a5565b91506109e6602084016108a5565b90509250929050565b600181811c90821680610a0357607f821691505b602082108103610a2357634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b8181038181111561032357610323610a29565b600080835481600182811c915080831680610a6e57607f831692505b60208084108203610a8d57634e487b7160e01b86526022600452602486fd5b818015610aa15760018114610ab657610ae3565b60ff1986168952841515850289019650610ae3565b60008a81526020902060005b86811015610adb5781548b820152908501908301610ac2565b505084890196505b509498975050505050505050565b8082018082111561032357610323610a2956fea264697066735822122019d6dcce9fa8a99adb050e740f9966f7fdffb937df7b84292ec3753b7ae9129a64736f6c63430008110033",
}

// SolmateERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use SolmateERC20MetaData.ABI instead.
var SolmateERC20ABI = SolmateERC20MetaData.ABI

// SolmateERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SolmateERC20MetaData.Bin instead.
var SolmateERC20Bin = SolmateERC20MetaData.Bin

// DeploySolmateERC20 deploys a new Ethereum contract, binding an instance of SolmateERC20 to it.
func DeploySolmateERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SolmateERC20, error) {
	parsed, err := SolmateERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SolmateERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SolmateERC20{SolmateERC20Caller: SolmateERC20Caller{contract: contract}, SolmateERC20Transactor: SolmateERC20Transactor{contract: contract}, SolmateERC20Filterer: SolmateERC20Filterer{contract: contract}}, nil
}

// SolmateERC20 is an auto generated Go binding around an Ethereum contract.
type SolmateERC20 struct {
	SolmateERC20Caller     // Read-only binding to the contract
	SolmateERC20Transactor // Write-only binding to the contract
	SolmateERC20Filterer   // Log filterer for contract events
}

// SolmateERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SolmateERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolmateERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SolmateERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolmateERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SolmateERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolmateERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SolmateERC20Session struct {
	Contract     *SolmateERC20     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolmateERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SolmateERC20CallerSession struct {
	Contract *SolmateERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SolmateERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SolmateERC20TransactorSession struct {
	Contract     *SolmateERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SolmateERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SolmateERC20Raw struct {
	Contract *SolmateERC20 // Generic contract binding to access the raw methods on
}

// SolmateERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SolmateERC20CallerRaw struct {
	Contract *SolmateERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SolmateERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SolmateERC20TransactorRaw struct {
	Contract *SolmateERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSolmateERC20 creates a new instance of SolmateERC20, bound to a specific deployed contract.
func NewSolmateERC20(address common.Address, backend bind.ContractBackend) (*SolmateERC20, error) {
	contract, err := bindSolmateERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SolmateERC20{SolmateERC20Caller: SolmateERC20Caller{contract: contract}, SolmateERC20Transactor: SolmateERC20Transactor{contract: contract}, SolmateERC20Filterer: SolmateERC20Filterer{contract: contract}}, nil
}

// NewSolmateERC20Caller creates a new read-only instance of SolmateERC20, bound to a specific deployed contract.
func NewSolmateERC20Caller(address common.Address, caller bind.ContractCaller) (*SolmateERC20Caller, error) {
	contract, err := bindSolmateERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SolmateERC20Caller{contract: contract}, nil
}

// NewSolmateERC20Transactor creates a new write-only instance of SolmateERC20, bound to a specific deployed contract.
func NewSolmateERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SolmateERC20Transactor, error) {
	contract, err := bindSolmateERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SolmateERC20Transactor{contract: contract}, nil
}

// NewSolmateERC20Filterer creates a new log filterer instance of SolmateERC20, bound to a specific deployed contract.
func NewSolmateERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SolmateERC20Filterer, error) {
	contract, err := bindSolmateERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SolmateERC20Filterer{contract: contract}, nil
}

// bindSolmateERC20 binds a generic wrapper to an already deployed contract.
func bindSolmateERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SolmateERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolmateERC20 *SolmateERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SolmateERC20.Contract.SolmateERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolmateERC20 *SolmateERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolmateERC20.Contract.SolmateERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolmateERC20 *SolmateERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolmateERC20.Contract.SolmateERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolmateERC20 *SolmateERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SolmateERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolmateERC20 *SolmateERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolmateERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolmateERC20 *SolmateERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolmateERC20.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_SolmateERC20 *SolmateERC20Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SolmateERC20.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_SolmateERC20 *SolmateERC20Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _SolmateERC20.Contract.DOMAINSEPARATOR(&_SolmateERC20.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_SolmateERC20 *SolmateERC20CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _SolmateERC20.Contract.DOMAINSEPARATOR(&_SolmateERC20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_SolmateERC20 *SolmateERC20Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SolmateERC20.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_SolmateERC20 *SolmateERC20Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SolmateERC20.Contract.Allowance(&_SolmateERC20.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_SolmateERC20 *SolmateERC20CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SolmateERC20.Contract.Allowance(&_SolmateERC20.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_SolmateERC20 *SolmateERC20Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SolmateERC20.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_SolmateERC20 *SolmateERC20Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _SolmateERC20.Contract.BalanceOf(&_SolmateERC20.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_SolmateERC20 *SolmateERC20CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _SolmateERC20.Contract.BalanceOf(&_SolmateERC20.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SolmateERC20 *SolmateERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SolmateERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SolmateERC20 *SolmateERC20Session) Decimals() (uint8, error) {
	return _SolmateERC20.Contract.Decimals(&_SolmateERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SolmateERC20 *SolmateERC20CallerSession) Decimals() (uint8, error) {
	return _SolmateERC20.Contract.Decimals(&_SolmateERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SolmateERC20 *SolmateERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SolmateERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SolmateERC20 *SolmateERC20Session) Name() (string, error) {
	return _SolmateERC20.Contract.Name(&_SolmateERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SolmateERC20 *SolmateERC20CallerSession) Name() (string, error) {
	return _SolmateERC20.Contract.Name(&_SolmateERC20.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_SolmateERC20 *SolmateERC20Caller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SolmateERC20.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_SolmateERC20 *SolmateERC20Session) Nonces(arg0 common.Address) (*big.Int, error) {
	return _SolmateERC20.Contract.Nonces(&_SolmateERC20.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_SolmateERC20 *SolmateERC20CallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _SolmateERC20.Contract.Nonces(&_SolmateERC20.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SolmateERC20 *SolmateERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SolmateERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SolmateERC20 *SolmateERC20Session) Symbol() (string, error) {
	return _SolmateERC20.Contract.Symbol(&_SolmateERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SolmateERC20 *SolmateERC20CallerSession) Symbol() (string, error) {
	return _SolmateERC20.Contract.Symbol(&_SolmateERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SolmateERC20 *SolmateERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SolmateERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SolmateERC20 *SolmateERC20Session) TotalSupply() (*big.Int, error) {
	return _SolmateERC20.Contract.TotalSupply(&_SolmateERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SolmateERC20 *SolmateERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _SolmateERC20.Contract.TotalSupply(&_SolmateERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SolmateERC20 *SolmateERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolmateERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SolmateERC20 *SolmateERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolmateERC20.Contract.Approve(&_SolmateERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SolmateERC20 *SolmateERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolmateERC20.Contract.Approve(&_SolmateERC20.TransactOpts, spender, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_SolmateERC20 *SolmateERC20Transactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolmateERC20.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_SolmateERC20 *SolmateERC20Session) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolmateERC20.Contract.Mint(&_SolmateERC20.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_SolmateERC20 *SolmateERC20TransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolmateERC20.Contract.Mint(&_SolmateERC20.TransactOpts, to, amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_SolmateERC20 *SolmateERC20Transactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SolmateERC20.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_SolmateERC20 *SolmateERC20Session) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SolmateERC20.Contract.Permit(&_SolmateERC20.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_SolmateERC20 *SolmateERC20TransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SolmateERC20.Contract.Permit(&_SolmateERC20.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_SolmateERC20 *SolmateERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolmateERC20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_SolmateERC20 *SolmateERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolmateERC20.Contract.Transfer(&_SolmateERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_SolmateERC20 *SolmateERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolmateERC20.Contract.Transfer(&_SolmateERC20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_SolmateERC20 *SolmateERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolmateERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_SolmateERC20 *SolmateERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolmateERC20.Contract.TransferFrom(&_SolmateERC20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_SolmateERC20 *SolmateERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolmateERC20.Contract.TransferFrom(&_SolmateERC20.TransactOpts, from, to, amount)
}

// SolmateERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SolmateERC20 contract.
type SolmateERC20ApprovalIterator struct {
	Event *SolmateERC20Approval // Event containing the contract specifics and raw log

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
func (it *SolmateERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolmateERC20Approval)
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
		it.Event = new(SolmateERC20Approval)
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
func (it *SolmateERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolmateERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolmateERC20Approval represents a Approval event raised by the SolmateERC20 contract.
type SolmateERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_SolmateERC20 *SolmateERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SolmateERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SolmateERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SolmateERC20ApprovalIterator{contract: _SolmateERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_SolmateERC20 *SolmateERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SolmateERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SolmateERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolmateERC20Approval)
				if err := _SolmateERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_SolmateERC20 *SolmateERC20Filterer) ParseApproval(log types.Log) (*SolmateERC20Approval, error) {
	event := new(SolmateERC20Approval)
	if err := _SolmateERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolmateERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SolmateERC20 contract.
type SolmateERC20TransferIterator struct {
	Event *SolmateERC20Transfer // Event containing the contract specifics and raw log

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
func (it *SolmateERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolmateERC20Transfer)
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
		it.Event = new(SolmateERC20Transfer)
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
func (it *SolmateERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolmateERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolmateERC20Transfer represents a Transfer event raised by the SolmateERC20 contract.
type SolmateERC20Transfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_SolmateERC20 *SolmateERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SolmateERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SolmateERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SolmateERC20TransferIterator{contract: _SolmateERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_SolmateERC20 *SolmateERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SolmateERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SolmateERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolmateERC20Transfer)
				if err := _SolmateERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_SolmateERC20 *SolmateERC20Filterer) ParseTransfer(log types.Log) (*SolmateERC20Transfer, error) {
	event := new(SolmateERC20Transfer)
	if err := _SolmateERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
