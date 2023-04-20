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

// SolmateERC20MetaData contains all meta data concerning the SolmateERC20 contract.
var SolmateERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"deployer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endowment\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"Created\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040526040518060400160405280600581526020017f546f6b656e0000000000000000000000000000000000000000000000000000008152506040518060400160405280600281526020017f544b000000000000000000000000000000000000000000000000000000000000815250601282600090816200008391906200043c565b5081600190816200009591906200043c565b508060ff1660808160ff16815250504660a08181525050620000bc6200013260201b60201c565b60c08181525050505050604051620000d4906200057e565b60405180910390203373ffffffffffffffffffffffffffffffffffffffff167f25b7b963b8a01e6cdc7edc9ccba6ec5cbbdf469ba5e61b163f9577df8dd033933460405162000124919062000607565b60405180910390a3620007be565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6000604051620001669190620006e8565b60405180910390207fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc64630604051602001620001a795949392919062000761565b60405160208183030381529060405280519060200120905090565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200024457607f821691505b6020821081036200025a5762000259620001fc565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620002c47fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000285565b620002d0868362000285565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006200031d620003176200031184620002e8565b620002f2565b620002e8565b9050919050565b6000819050919050565b6200033983620002fc565b62000351620003488262000324565b84845462000292565b825550505050565b600090565b6200036862000359565b620003758184846200032e565b505050565b5b818110156200039d57620003916000826200035e565b6001810190506200037b565b5050565b601f821115620003ec57620003b68162000260565b620003c18462000275565b81016020851015620003d1578190505b620003e9620003e08562000275565b8301826200037a565b50505b505050565b600082821c905092915050565b60006200041160001984600802620003f1565b1980831691505092915050565b60006200042c8383620003fe565b9150826002028217905092915050565b6200044782620001c2565b67ffffffffffffffff811115620004635762000462620001cd565b5b6200046f82546200022b565b6200047c828285620003a1565b600060209050601f831160018114620004b457600084156200049f578287015190505b620004ab85826200041e565b8655506200051b565b601f198416620004c48662000260565b60005b82811015620004ee57848901518255600182019150602085019450602081019050620004c7565b868310156200050e57848901516200050a601f891682620003fe565b8355505b6001600288020188555050505b505050505050565b600081905092915050565b7f546f6b656e000000000000000000000000000000000000000000000000000000600082015250565b60006200056660058362000523565b915062000573826200052e565b600582019050919050565b60006200058b8262000557565b9150819050919050565b620005a081620002e8565b82525050565b600082825260208201905092915050565b7f544b000000000000000000000000000000000000000000000000000000000000600082015250565b6000620005ef600283620005a6565b9150620005fc82620005b7565b602082019050919050565b60006040820190506200061e600083018462000595565b81810360208301526200063181620005e0565b905092915050565b600081905092915050565b60008190508160005260206000209050919050565b6000815462000668816200022b565b62000674818662000639565b94506001821660008114620006925760018114620006a857620006df565b60ff1983168652811515820286019350620006df565b620006b38562000644565b60005b83811015620006d757815481890152600182019150602081019050620006b6565b838801955050505b50505092915050565b6000620006f6828462000659565b915081905092915050565b6000819050919050565b620007168162000701565b82525050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062000749826200071c565b9050919050565b6200075b816200073c565b82525050565b600060a0820190506200077860008301886200070b565b6200078760208301876200070b565b6200079660408301866200070b565b620007a5606083018562000595565b620007b4608083018462000750565b9695505050505050565b60805160a05160c051611641620007ee60003960006106ee015260006106ba0152600061069401526116416000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c806340c10f191161008c57806395d89b411161006657806395d89b4114610228578063a9059cbb14610246578063d505accf14610276578063dd62ed3e14610292576100cf565b806340c10f19146101ac57806370a08231146101c85780637ecebe00146101f8576100cf565b806306fdde03146100d4578063095ea7b3146100f257806318160ddd1461012257806323b872dd14610140578063313ce567146101705780633644e5151461018e575b600080fd5b6100dc6102c2565b6040516100e99190610e4b565b60405180910390f35b61010c60048036038101906101079190610f06565b610350565b6040516101199190610f61565b60405180910390f35b61012a610442565b6040516101379190610f8b565b60405180910390f35b61015a60048036038101906101559190610fa6565b610448565b6040516101679190610f61565b60405180910390f35b610178610692565b6040516101859190611015565b60405180910390f35b6101966106b6565b6040516101a39190611049565b60405180910390f35b6101c660048036038101906101c19190610f06565b610713565b005b6101e260048036038101906101dd9190611064565b61076f565b6040516101ef9190610f8b565b60405180910390f35b610212600480360381019061020d9190611064565b610787565b60405161021f9190610f8b565b60405180910390f35b61023061079f565b60405161023d9190610e4b565b60405180910390f35b610260600480360381019061025b9190610f06565b61082d565b60405161026d9190610f61565b60405180910390f35b610290600480360381019061028b91906110e9565b610941565b005b6102ac60048036038101906102a7919061118b565b610c3a565b6040516102b99190610f8b565b60405180910390f35b600080546102cf906111fa565b80601f01602080910402602001604051908101604052809291908181526020018280546102fb906111fa565b80156103485780601f1061031d57610100808354040283529160200191610348565b820191906000526020600020905b81548152906001019060200180831161032b57829003601f168201915b505050505081565b600081600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516104309190610f8b565b60405180910390a36001905092915050565b60025481565b600080600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461057e5782816104fd919061125a565b600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b82600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546105cd919061125a565b9250508190555082600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8560405161067e9190610f8b565b60405180910390a360019150509392505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60007f000000000000000000000000000000000000000000000000000000000000000046146106ec576106e7610c5f565b61070e565b7f00000000000000000000000000000000000000000000000000000000000000005b905090565b61071d8282610ceb565b8173ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885826040516107639190610f8b565b60405180910390a25050565b60036020528060005260406000206000915090505481565b60056020528060005260406000206000915090505481565b600180546107ac906111fa565b80601f01602080910402602001604051908101604052809291908181526020018280546107d8906111fa565b80156108255780601f106107fa57610100808354040283529160200191610825565b820191906000526020600020905b81548152906001019060200180831161080857829003601f168201915b505050505081565b600081600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461087e919061125a565b9250508190555081600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161092f9190610f8b565b60405180910390a36001905092915050565b42841015610984576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097b906112da565b60405180910390fd5b600060016109906106b6565b7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98a8a8a600560008f73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000815480929190600101919050558b604051602001610a1896959493929190611309565b60405160208183030381529060405280519060200120604051602001610a3f9291906113e2565b6040516020818303038152906040528051906020012085858560405160008152602001604052604051610a759493929190611419565b6020604051602081039080840390855afa158015610a97573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614158015610b0b57508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b610b4a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b41906114aa565b60405180910390fd5b85600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92587604051610c299190610f8b565b60405180910390a350505050505050565b6004602052816000526040600020602052806000526040600020600091509150505481565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6000604051610c91919061156d565b60405180910390207fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc64630604051602001610cd0959493929190611584565b60405160208183030381529060405280519060200120905090565b8060026000828254610cfd91906115d7565b9250508190555080600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610daf9190610f8b565b60405180910390a35050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610df5578082015181840152602081019050610dda565b60008484015250505050565b6000601f19601f8301169050919050565b6000610e1d82610dbb565b610e278185610dc6565b9350610e37818560208601610dd7565b610e4081610e01565b840191505092915050565b60006020820190508181036000830152610e658184610e12565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610e9d82610e72565b9050919050565b610ead81610e92565b8114610eb857600080fd5b50565b600081359050610eca81610ea4565b92915050565b6000819050919050565b610ee381610ed0565b8114610eee57600080fd5b50565b600081359050610f0081610eda565b92915050565b60008060408385031215610f1d57610f1c610e6d565b5b6000610f2b85828601610ebb565b9250506020610f3c85828601610ef1565b9150509250929050565b60008115159050919050565b610f5b81610f46565b82525050565b6000602082019050610f766000830184610f52565b92915050565b610f8581610ed0565b82525050565b6000602082019050610fa06000830184610f7c565b92915050565b600080600060608486031215610fbf57610fbe610e6d565b5b6000610fcd86828701610ebb565b9350506020610fde86828701610ebb565b9250506040610fef86828701610ef1565b9150509250925092565b600060ff82169050919050565b61100f81610ff9565b82525050565b600060208201905061102a6000830184611006565b92915050565b6000819050919050565b61104381611030565b82525050565b600060208201905061105e600083018461103a565b92915050565b60006020828403121561107a57611079610e6d565b5b600061108884828501610ebb565b91505092915050565b61109a81610ff9565b81146110a557600080fd5b50565b6000813590506110b781611091565b92915050565b6110c681611030565b81146110d157600080fd5b50565b6000813590506110e3816110bd565b92915050565b600080600080600080600060e0888a03121561110857611107610e6d565b5b60006111168a828b01610ebb565b97505060206111278a828b01610ebb565b96505060406111388a828b01610ef1565b95505060606111498a828b01610ef1565b945050608061115a8a828b016110a8565b93505060a061116b8a828b016110d4565b92505060c061117c8a828b016110d4565b91505092959891949750929550565b600080604083850312156111a2576111a1610e6d565b5b60006111b085828601610ebb565b92505060206111c185828601610ebb565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061121257607f821691505b602082108103611225576112246111cb565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061126582610ed0565b915061127083610ed0565b92508282039050818111156112885761128761122b565b5b92915050565b7f5045524d49545f444541444c494e455f45585049524544000000000000000000600082015250565b60006112c4601783610dc6565b91506112cf8261128e565b602082019050919050565b600060208201905081810360008301526112f3816112b7565b9050919050565b61130381610e92565b82525050565b600060c08201905061131e600083018961103a565b61132b60208301886112fa565b61133860408301876112fa565b6113456060830186610f7c565b6113526080830185610f7c565b61135f60a0830184610f7c565b979650505050505050565b600081905092915050565b7f1901000000000000000000000000000000000000000000000000000000000000600082015250565b60006113ab60028361136a565b91506113b682611375565b600282019050919050565b6000819050919050565b6113dc6113d782611030565b6113c1565b82525050565b60006113ed8261139e565b91506113f982856113cb565b60208201915061140982846113cb565b6020820191508190509392505050565b600060808201905061142e600083018761103a565b61143b6020830186611006565b611448604083018561103a565b611455606083018461103a565b95945050505050565b7f494e56414c49445f5349474e4552000000000000000000000000000000000000600082015250565b6000611494600e83610dc6565b915061149f8261145e565b602082019050919050565b600060208201905081810360008301526114c381611487565b9050919050565b600081905092915050565b60008190508160005260206000209050919050565b600081546114f7816111fa565b61150181866114ca565b9450600182166000811461151c576001811461153157611564565b60ff1983168652811515820286019350611564565b61153a856114d5565b60005b8381101561155c5781548189015260018201915060208101905061153d565b838801955050505b50505092915050565b600061157982846114ea565b915081905092915050565b600060a082019050611599600083018861103a565b6115a6602083018761103a565b6115b3604083018661103a565b6115c06060830185610f7c565b6115cd60808301846112fa565b9695505050505050565b60006115e282610ed0565b91506115ed83610ed0565b92508282019050808211156116055761160461122b565b5b9291505056fea2646970667358221220b1164da7a78e14deee43a677426be49492783af6c175384a81a3f14b74f01d4a64736f6c63430008130033",
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
	parsed, err := SolmateERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// SolmateERC20CreatedIterator is returned from FilterCreated and is used to iterate over the raw logs and unpacked data for Created events raised by the SolmateERC20 contract.
type SolmateERC20CreatedIterator struct {
	Event *SolmateERC20Created // Event containing the contract specifics and raw log

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
func (it *SolmateERC20CreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolmateERC20Created)
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
		it.Event = new(SolmateERC20Created)
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
func (it *SolmateERC20CreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolmateERC20CreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolmateERC20Created represents a Created event raised by the SolmateERC20 contract.
type SolmateERC20Created struct {
	Deployer  common.Address
	Endowment *big.Int
	Name      common.Hash
	Symbol    string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCreated is a free log retrieval operation binding the contract event 0x25b7b963b8a01e6cdc7edc9ccba6ec5cbbdf469ba5e61b163f9577df8dd03393.
//
// Solidity: event Created(address indexed deployer, uint256 endowment, string indexed name, string symbol)
func (_SolmateERC20 *SolmateERC20Filterer) FilterCreated(opts *bind.FilterOpts, deployer []common.Address, name []string) (*SolmateERC20CreatedIterator, error) {

	var deployerRule []interface{}
	for _, deployerItem := range deployer {
		deployerRule = append(deployerRule, deployerItem)
	}

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _SolmateERC20.contract.FilterLogs(opts, "Created", deployerRule, nameRule)
	if err != nil {
		return nil, err
	}
	return &SolmateERC20CreatedIterator{contract: _SolmateERC20.contract, event: "Created", logs: logs, sub: sub}, nil
}

// WatchCreated is a free log subscription operation binding the contract event 0x25b7b963b8a01e6cdc7edc9ccba6ec5cbbdf469ba5e61b163f9577df8dd03393.
//
// Solidity: event Created(address indexed deployer, uint256 endowment, string indexed name, string symbol)
func (_SolmateERC20 *SolmateERC20Filterer) WatchCreated(opts *bind.WatchOpts, sink chan<- *SolmateERC20Created, deployer []common.Address, name []string) (event.Subscription, error) {

	var deployerRule []interface{}
	for _, deployerItem := range deployer {
		deployerRule = append(deployerRule, deployerItem)
	}

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _SolmateERC20.contract.WatchLogs(opts, "Created", deployerRule, nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolmateERC20Created)
				if err := _SolmateERC20.contract.UnpackLog(event, "Created", log); err != nil {
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

// ParseCreated is a log parse operation binding the contract event 0x25b7b963b8a01e6cdc7edc9ccba6ec5cbbdf469ba5e61b163f9577df8dd03393.
//
// Solidity: event Created(address indexed deployer, uint256 endowment, string indexed name, string symbol)
func (_SolmateERC20 *SolmateERC20Filterer) ParseCreated(log types.Log) (*SolmateERC20Created, error) {
	event := new(SolmateERC20Created)
	if err := _SolmateERC20.contract.UnpackLog(event, "Created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolmateERC20MintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the SolmateERC20 contract.
type SolmateERC20MintIterator struct {
	Event *SolmateERC20Mint // Event containing the contract specifics and raw log

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
func (it *SolmateERC20MintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolmateERC20Mint)
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
		it.Event = new(SolmateERC20Mint)
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
func (it *SolmateERC20MintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolmateERC20MintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolmateERC20Mint represents a Mint event raised by the SolmateERC20 contract.
type SolmateERC20Mint struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed to, uint256 amount)
func (_SolmateERC20 *SolmateERC20Filterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*SolmateERC20MintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SolmateERC20.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &SolmateERC20MintIterator{contract: _SolmateERC20.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed to, uint256 amount)
func (_SolmateERC20 *SolmateERC20Filterer) WatchMint(opts *bind.WatchOpts, sink chan<- *SolmateERC20Mint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SolmateERC20.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolmateERC20Mint)
				if err := _SolmateERC20.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed to, uint256 amount)
func (_SolmateERC20 *SolmateERC20Filterer) ParseMint(log types.Log) (*SolmateERC20Mint, error) {
	event := new(SolmateERC20Mint)
	if err := _SolmateERC20.contract.UnpackLog(event, "Mint", log); err != nil {
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
