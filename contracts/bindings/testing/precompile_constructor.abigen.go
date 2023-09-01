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

// PrecompileConstructorMetaData contains all meta data concerning the PrecompileConstructor contract.
var PrecompileConstructorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"stakingModule\",\"outputs\":[{\"internalType\":\"contractIStakingModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405273d9a998cac66092748ffec7cfbd155aae1737c2ff73ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff16815250348015610056575f80fd5b5060805173ffffffffffffffffffffffffffffffffffffffff16631904bb2e5f6040518263ffffffff1660e01b81526004016100929190610119565b5f60405180830381865afa1580156100ac573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906100d491906108c4565b5061090b565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610103826100da565b9050919050565b610113816100f9565b82525050565b5f60208201905061012c5f83018461010a565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61018d82610147565b810181811067ffffffffffffffff821117156101ac576101ab610157565b5b80604052505050565b5f6101be610132565b90506101ca8282610184565b919050565b5f80fd5b6101dc816100f9565b81146101e6575f80fd5b50565b5f815190506101f7816101d3565b92915050565b5f80fd5b5f80fd5b5f67ffffffffffffffff82111561021f5761021e610157565b5b61022882610147565b9050602081019050919050565b5f5b83811015610252578082015181840152602081019050610237565b5f8484015250505050565b5f61026f61026a84610205565b6101b5565b90508281526020810184848401111561028b5761028a610201565b5b610296848285610235565b509392505050565b5f82601f8301126102b2576102b16101fd565b5b81516102c284826020860161025d565b91505092915050565b5f8115159050919050565b6102df816102cb565b81146102e9575f80fd5b50565b5f815190506102fa816102d6565b92915050565b5f67ffffffffffffffff82111561031a57610319610157565b5b61032382610147565b9050602081019050919050565b5f61034261033d84610300565b6101b5565b90508281526020810184848401111561035e5761035d610201565b5b610369848285610235565b509392505050565b5f82601f830112610385576103846101fd565b5b8151610395848260208601610330565b91505092915050565b5f819050919050565b6103b08161039e565b81146103ba575f80fd5b50565b5f815190506103cb816103a7565b92915050565b5f60a082840312156103e6576103e5610143565b5b6103f060a06101b5565b90505f82015167ffffffffffffffff81111561040f5761040e6101cf565b5b61041b84828501610371565b5f83015250602082015167ffffffffffffffff81111561043e5761043d6101cf565b5b61044a84828501610371565b602083015250604082015167ffffffffffffffff81111561046e5761046d6101cf565b5b61047a84828501610371565b604083015250606082015167ffffffffffffffff81111561049e5761049d6101cf565b5b6104aa84828501610371565b606083015250608082015167ffffffffffffffff8111156104ce576104cd6101cf565b5b6104da84828501610371565b60808301525092915050565b5f8160070b9050919050565b6104fb816104e6565b8114610505575f80fd5b50565b5f81519050610516816104f2565b92915050565b5f6060828403121561053157610530610143565b5b61053b60606101b5565b90505f61054a848285016103bd565b5f83015250602061055d848285016103bd565b6020830152506040610571848285016103bd565b60408301525092915050565b5f6080828403121561059257610591610143565b5b61059c60406101b5565b90505f6105ab8482850161051c565b5f83015250606082015167ffffffffffffffff8111156105ce576105cd6101cf565b5b6105da84828501610371565b60208301525092915050565b5f67ffffffffffffffff821115610600576105ff610157565b5b602082029050602081019050919050565b5f80fd5b5f67ffffffffffffffff82169050919050565b61063181610615565b811461063b575f80fd5b50565b5f8151905061064c81610628565b92915050565b5f61066461065f846105e6565b6101b5565b9050808382526020820190506020840283018581111561068757610686610611565b5b835b818110156106b0578061069c888261063e565b845260208401935050602081019050610689565b5050509392505050565b5f82601f8301126106ce576106cd6101fd565b5b81516106de848260208601610652565b91505092915050565b5f6101a082840312156106fd576106fc610143565b5b6107086101a06101b5565b90505f610717848285016101e9565b5f83015250602082015167ffffffffffffffff81111561073a576107396101cf565b5b6107468482850161029e565b602083015250604061075a848285016102ec565b604083015250606082015167ffffffffffffffff81111561077e5761077d6101cf565b5b61078a84828501610371565b606083015250608061079e848285016103bd565b60808301525060a06107b2848285016103bd565b60a08301525060c082015167ffffffffffffffff8111156107d6576107d56101cf565b5b6107e2848285016103d1565b60c08301525060e06107f684828501610508565b60e08301525061010082015167ffffffffffffffff81111561081b5761081a6101cf565b5b61082784828501610371565b6101008301525061012082015167ffffffffffffffff81111561084d5761084c6101cf565b5b6108598482850161057d565b6101208301525061014061086f848285016103bd565b6101408301525061016061088584828501610508565b6101608301525061018082015167ffffffffffffffff8111156108ab576108aa6101cf565b5b6108b7848285016106ba565b6101808301525092915050565b5f602082840312156108d9576108d861013b565b5b5f82015167ffffffffffffffff8111156108f6576108f561013f565b5b610902848285016106e7565b91505092915050565b6080516101236109225f395f604601526101235ff3fe6080604052348015600e575f80fd5b50600436106026575f3560e01c8063504b82bf14602a575b5f80fd5b60306044565b604051603b919060d6565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000081565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f60a460a0609c846068565b6087565b6068565b9050919050565b5f60b3826090565b9050919050565b5f60c28260ab565b9050919050565b60d08160ba565b82525050565b5f60208201905060e75f83018460c9565b9291505056fea26469706673582212209c33115484a191f43eed189dd1fe6929b366f75b03d31a4c47a12ab0bfca7b8a64736f6c63430008140033",
}

// PrecompileConstructorABI is the input ABI used to generate the binding from.
// Deprecated: Use PrecompileConstructorMetaData.ABI instead.
var PrecompileConstructorABI = PrecompileConstructorMetaData.ABI

// PrecompileConstructorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PrecompileConstructorMetaData.Bin instead.
var PrecompileConstructorBin = PrecompileConstructorMetaData.Bin

// DeployPrecompileConstructor deploys a new Ethereum contract, binding an instance of PrecompileConstructor to it.
func DeployPrecompileConstructor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PrecompileConstructor, error) {
	parsed, err := PrecompileConstructorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PrecompileConstructorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PrecompileConstructor{PrecompileConstructorCaller: PrecompileConstructorCaller{contract: contract}, PrecompileConstructorTransactor: PrecompileConstructorTransactor{contract: contract}, PrecompileConstructorFilterer: PrecompileConstructorFilterer{contract: contract}}, nil
}

// PrecompileConstructor is an auto generated Go binding around an Ethereum contract.
type PrecompileConstructor struct {
	PrecompileConstructorCaller     // Read-only binding to the contract
	PrecompileConstructorTransactor // Write-only binding to the contract
	PrecompileConstructorFilterer   // Log filterer for contract events
}

// PrecompileConstructorCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrecompileConstructorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompileConstructorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrecompileConstructorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompileConstructorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrecompileConstructorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompileConstructorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrecompileConstructorSession struct {
	Contract     *PrecompileConstructor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PrecompileConstructorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrecompileConstructorCallerSession struct {
	Contract *PrecompileConstructorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// PrecompileConstructorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrecompileConstructorTransactorSession struct {
	Contract     *PrecompileConstructorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// PrecompileConstructorRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrecompileConstructorRaw struct {
	Contract *PrecompileConstructor // Generic contract binding to access the raw methods on
}

// PrecompileConstructorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrecompileConstructorCallerRaw struct {
	Contract *PrecompileConstructorCaller // Generic read-only contract binding to access the raw methods on
}

// PrecompileConstructorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrecompileConstructorTransactorRaw struct {
	Contract *PrecompileConstructorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrecompileConstructor creates a new instance of PrecompileConstructor, bound to a specific deployed contract.
func NewPrecompileConstructor(address common.Address, backend bind.ContractBackend) (*PrecompileConstructor, error) {
	contract, err := bindPrecompileConstructor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PrecompileConstructor{PrecompileConstructorCaller: PrecompileConstructorCaller{contract: contract}, PrecompileConstructorTransactor: PrecompileConstructorTransactor{contract: contract}, PrecompileConstructorFilterer: PrecompileConstructorFilterer{contract: contract}}, nil
}

// NewPrecompileConstructorCaller creates a new read-only instance of PrecompileConstructor, bound to a specific deployed contract.
func NewPrecompileConstructorCaller(address common.Address, caller bind.ContractCaller) (*PrecompileConstructorCaller, error) {
	contract, err := bindPrecompileConstructor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrecompileConstructorCaller{contract: contract}, nil
}

// NewPrecompileConstructorTransactor creates a new write-only instance of PrecompileConstructor, bound to a specific deployed contract.
func NewPrecompileConstructorTransactor(address common.Address, transactor bind.ContractTransactor) (*PrecompileConstructorTransactor, error) {
	contract, err := bindPrecompileConstructor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrecompileConstructorTransactor{contract: contract}, nil
}

// NewPrecompileConstructorFilterer creates a new log filterer instance of PrecompileConstructor, bound to a specific deployed contract.
func NewPrecompileConstructorFilterer(address common.Address, filterer bind.ContractFilterer) (*PrecompileConstructorFilterer, error) {
	contract, err := bindPrecompileConstructor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrecompileConstructorFilterer{contract: contract}, nil
}

// bindPrecompileConstructor binds a generic wrapper to an already deployed contract.
func bindPrecompileConstructor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PrecompileConstructorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrecompileConstructor *PrecompileConstructorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrecompileConstructor.Contract.PrecompileConstructorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrecompileConstructor *PrecompileConstructorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrecompileConstructor.Contract.PrecompileConstructorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrecompileConstructor *PrecompileConstructorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrecompileConstructor.Contract.PrecompileConstructorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrecompileConstructor *PrecompileConstructorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrecompileConstructor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrecompileConstructor *PrecompileConstructorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrecompileConstructor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrecompileConstructor *PrecompileConstructorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrecompileConstructor.Contract.contract.Transact(opts, method, params...)
}

// StakingModule is a free data retrieval call binding the contract method 0x504b82bf.
//
// Solidity: function stakingModule() view returns(address)
func (_PrecompileConstructor *PrecompileConstructorCaller) StakingModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrecompileConstructor.contract.Call(opts, &out, "stakingModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingModule is a free data retrieval call binding the contract method 0x504b82bf.
//
// Solidity: function stakingModule() view returns(address)
func (_PrecompileConstructor *PrecompileConstructorSession) StakingModule() (common.Address, error) {
	return _PrecompileConstructor.Contract.StakingModule(&_PrecompileConstructor.CallOpts)
}

// StakingModule is a free data retrieval call binding the contract method 0x504b82bf.
//
// Solidity: function stakingModule() view returns(address)
func (_PrecompileConstructor *PrecompileConstructorCallerSession) StakingModule() (common.Address, error) {
	return _PrecompileConstructor.Contract.StakingModule(&_PrecompileConstructor.CallOpts)
}
