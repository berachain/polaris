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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"abera\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denom\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20Module\",\"outputs\":[{\"internalType\":\"contractIERC20Module\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040526269696973ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801562000046575f80fd5b505f60805173ffffffffffffffffffffffffffffffffffffffff1663096b4069333363075bcd156040518463ffffffff1660e01b81526004016200008d9392919062000416565b6020604051808303815f875af1158015620000aa573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620000d09190620004b1565b90508062000115576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200010c906200052f565b60405180910390fd5b60805173ffffffffffffffffffffffffffffffffffffffff1663a333e57c6040518163ffffffff1660e01b815260040162000150906200054f565b602060405180830381865afa1580156200016c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620001929190620005b1565b5f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060805173ffffffffffffffffffffffffffffffffffffffff1663cd22a0185f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b81526004016200022c91906200063f565b5f60405180830381865afa15801562000247573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190620002719190620007ce565b6001908162000281919062000a42565b506040518060400160405280600581526020017f6162657261000000000000000000000000000000000000000000000000000000815250805190602001206001604051620002d0919062000bce565b6040518091039020146200031b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620003129062000c34565b60405180910390fd5b5062000c54565b5f82825260208201905092915050565b7f61626572610000000000000000000000000000000000000000000000000000005f82015250565b5f6200036860058362000322565b9150620003758262000332565b602082019050919050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620003ab8262000380565b9050919050565b620003bd816200039f565b82525050565b5f819050919050565b5f819050919050565b5f819050919050565b5f620003fe620003f8620003f284620003c3565b620003d5565b620003cc565b9050919050565b6200041081620003de565b82525050565b5f6080820190508181035f8301526200042f816200035a565b9050620004406020830186620003b2565b6200044f6040830185620003b2565b6200045e606083018462000405565b949350505050565b5f604051905090565b5f80fd5b5f80fd5b5f8115159050919050565b6200048d8162000477565b811462000498575f80fd5b50565b5f81519050620004ab8162000482565b92915050565b5f60208284031215620004c957620004c86200046f565b5b5f620004d8848285016200049b565b91505092915050565b7f6661696c656420746f207472616e7366657220616265726100000000000000005f82015250565b5f6200051760188362000322565b91506200052482620004e1565b602082019050919050565b5f6020820190508181035f830152620005488162000509565b9050919050565b5f6020820190508181035f83015262000568816200035a565b9050919050565b5f6200057b826200039f565b9050919050565b6200058d816200056f565b811462000598575f80fd5b50565b5f81519050620005ab8162000582565b92915050565b5f60208284031215620005c957620005c86200046f565b5b5f620005d8848285016200059b565b91505092915050565b5f62000601620005fb620005f58462000380565b620003d5565b62000380565b9050919050565b5f6200061482620005e1565b9050919050565b5f620006278262000608565b9050919050565b62000639816200061b565b82525050565b5f602082019050620006545f8301846200062e565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b620006aa8262000662565b810181811067ffffffffffffffff82111715620006cc57620006cb62000672565b5b80604052505050565b5f620006e062000466565b9050620006ee82826200069f565b919050565b5f67ffffffffffffffff82111562000710576200070f62000672565b5b6200071b8262000662565b9050602081019050919050565b5f5b83811015620007475780820151818401526020810190506200072a565b5f8484015250505050565b5f620007686200076284620006f3565b620006d5565b9050828152602081018484840111156200078757620007866200065e565b5b6200079484828562000728565b509392505050565b5f82601f830112620007b357620007b26200065a565b5b8151620007c584826020860162000752565b91505092915050565b5f60208284031215620007e657620007e56200046f565b5b5f82015167ffffffffffffffff81111562000806576200080562000473565b5b62000814848285016200079c565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806200086c57607f821691505b60208210810362000882576200088162000827565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620008e67fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620008a9565b620008f28683620008a9565b95508019841693508086168417925050509392505050565b5f6200092a620009246200091e84620003cc565b620003d5565b620003cc565b9050919050565b5f819050919050565b62000945836200090a565b6200095d620009548262000931565b848454620008b5565b825550505050565b5f90565b6200097362000965565b620009808184846200093a565b505050565b5b81811015620009a7576200099b5f8262000969565b60018101905062000986565b5050565b601f821115620009f657620009c08162000888565b620009cb846200089a565b81016020851015620009db578190505b620009f3620009ea856200089a565b83018262000985565b50505b505050565b5f82821c905092915050565b5f62000a185f1984600802620009fb565b1980831691505092915050565b5f62000a32838362000a07565b9150826002028217905092915050565b62000a4d826200081d565b67ffffffffffffffff81111562000a695762000a6862000672565b5b62000a75825462000854565b62000a82828285620009ab565b5f60209050601f83116001811462000ab8575f841562000aa3578287015190505b62000aaf858262000a25565b86555062000b1e565b601f19841662000ac88662000888565b5f5b8281101562000af15784890151825560018201915060208501945060208101905062000aca565b8683101562000b11578489015162000b0d601f89168262000a07565b8355505b6001600288020188555050505b505050505050565b5f81905092915050565b5f819050815f5260205f209050919050565b5f815462000b508162000854565b62000b5c818662000b26565b9450600182165f811462000b79576001811462000b8f5762000bc5565b60ff198316865281151582028601935062000bc5565b62000b9a8562000b30565b5f5b8381101562000bbd5781548189015260018201915060208101905062000b9c565b838801955050505b50505092915050565b5f62000bdb828462000b42565b915081905092915050565b7f72657475726e6564207468652077726f6e672064656e6f6d00000000000000005f82015250565b5f62000c1c60188362000322565b915062000c298262000be6565b602082019050919050565b5f6020820190508181035f83015262000c4d8162000c0e565b9050919050565b60805161037962000c6c5f395f60c201526103795ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c8063558f208414610043578063714ba40c14610061578063c370b0421461007f575b5f80fd5b61004b61009d565b60405161005891906101ea565b60405180910390f35b6100696100c0565b6040516100769190610223565b60405180910390f35b6100876100e4565b60405161009491906102c6565b60405180910390f35b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b600180546100f190610313565b80601f016020809104026020016040519081016040528092919081815260200182805461011d90610313565b80156101685780601f1061013f57610100808354040283529160200191610168565b820191905f5260205f20905b81548152906001019060200180831161014b57829003601f168201915b505050505081565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f6101b26101ad6101a884610170565b61018f565b610170565b9050919050565b5f6101c382610198565b9050919050565b5f6101d4826101b9565b9050919050565b6101e4816101ca565b82525050565b5f6020820190506101fd5f8301846101db565b92915050565b5f61020d826101b9565b9050919050565b61021d81610203565b82525050565b5f6020820190506102365f830184610214565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610273578082015181840152602081019050610258565b5f8484015250505050565b5f601f19601f8301169050919050565b5f6102988261023c565b6102a28185610246565b93506102b2818560208601610256565b6102bb8161027e565b840191505092915050565b5f6020820190508181035f8301526102de818461028e565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061032a57607f821691505b60208210810361033d5761033c6102e6565b5b5091905056fea26469706673582212202cc2a4438e56dfd9ecede126f6b482fefc7f3046450291a96872d6c303f3dbc264736f6c63430008140033",
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

// Abera is a free data retrieval call binding the contract method 0x558f2084.
//
// Solidity: function abera() view returns(address)
func (_PrecompileConstructor *PrecompileConstructorCaller) Abera(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrecompileConstructor.contract.Call(opts, &out, "abera")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Abera is a free data retrieval call binding the contract method 0x558f2084.
//
// Solidity: function abera() view returns(address)
func (_PrecompileConstructor *PrecompileConstructorSession) Abera() (common.Address, error) {
	return _PrecompileConstructor.Contract.Abera(&_PrecompileConstructor.CallOpts)
}

// Abera is a free data retrieval call binding the contract method 0x558f2084.
//
// Solidity: function abera() view returns(address)
func (_PrecompileConstructor *PrecompileConstructorCallerSession) Abera() (common.Address, error) {
	return _PrecompileConstructor.Contract.Abera(&_PrecompileConstructor.CallOpts)
}

// Denom is a free data retrieval call binding the contract method 0xc370b042.
//
// Solidity: function denom() view returns(string)
func (_PrecompileConstructor *PrecompileConstructorCaller) Denom(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PrecompileConstructor.contract.Call(opts, &out, "denom")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Denom is a free data retrieval call binding the contract method 0xc370b042.
//
// Solidity: function denom() view returns(string)
func (_PrecompileConstructor *PrecompileConstructorSession) Denom() (string, error) {
	return _PrecompileConstructor.Contract.Denom(&_PrecompileConstructor.CallOpts)
}

// Denom is a free data retrieval call binding the contract method 0xc370b042.
//
// Solidity: function denom() view returns(string)
func (_PrecompileConstructor *PrecompileConstructorCallerSession) Denom() (string, error) {
	return _PrecompileConstructor.Contract.Denom(&_PrecompileConstructor.CallOpts)
}

// Erc20Module is a free data retrieval call binding the contract method 0x714ba40c.
//
// Solidity: function erc20Module() view returns(address)
func (_PrecompileConstructor *PrecompileConstructorCaller) Erc20Module(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrecompileConstructor.contract.Call(opts, &out, "erc20Module")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20Module is a free data retrieval call binding the contract method 0x714ba40c.
//
// Solidity: function erc20Module() view returns(address)
func (_PrecompileConstructor *PrecompileConstructorSession) Erc20Module() (common.Address, error) {
	return _PrecompileConstructor.Contract.Erc20Module(&_PrecompileConstructor.CallOpts)
}

// Erc20Module is a free data retrieval call binding the contract method 0x714ba40c.
//
// Solidity: function erc20Module() view returns(address)
func (_PrecompileConstructor *PrecompileConstructorCallerSession) Erc20Module() (common.Address, error) {
	return _PrecompileConstructor.Contract.Erc20Module(&_PrecompileConstructor.CallOpts)
}
