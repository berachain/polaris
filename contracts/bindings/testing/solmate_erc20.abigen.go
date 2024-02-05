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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60e060405234801562000010575f80fd5b506040518060400160405280600581526020017f546f6b656e0000000000000000000000000000000000000000000000000000008152506040518060400160405280600281526020017f544b0000000000000000000000000000000000000000000000000000000000008152506012825f90816200008f9190620003ca565b508160019081620000a19190620003ca565b508060ff1660808160ff16815250504660a08181525050620000c8620000d860201b60201c565b60c0818152505050505062000637565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f6040516200010a919062000556565b60405180910390207fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc646306040516020016200014b959493929190620005dc565b60405160208183030381529060405280519060200120905090565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680620001e257607f821691505b602082108103620001f857620001f76200019d565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026200025c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200021f565b6200026886836200021f565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f620002b2620002ac620002a68462000280565b62000289565b62000280565b9050919050565b5f819050919050565b620002cd8362000292565b620002e5620002dc82620002b9565b8484546200022b565b825550505050565b5f90565b620002fb620002ed565b62000308818484620002c2565b505050565b5b818110156200032f57620003235f82620002f1565b6001810190506200030e565b5050565b601f8211156200037e576200034881620001fe565b620003538462000210565b8101602085101562000363578190505b6200037b620003728562000210565b8301826200030d565b50505b505050565b5f82821c905092915050565b5f620003a05f198460080262000383565b1980831691505092915050565b5f620003ba83836200038f565b9150826002028217905092915050565b620003d58262000166565b67ffffffffffffffff811115620003f157620003f062000170565b5b620003fd8254620001ca565b6200040a82828562000333565b5f60209050601f83116001811462000440575f84156200042b578287015190505b620004378582620003ad565b865550620004a6565b601f1984166200045086620001fe565b5f5b82811015620004795784890151825560018201915060208501945060208101905062000452565b8683101562000499578489015162000495601f8916826200038f565b8355505b6001600288020188555050505b505050505050565b5f81905092915050565b5f819050815f5260205f209050919050565b5f8154620004d881620001ca565b620004e48186620004ae565b9450600182165f811462000501576001811462000517576200054d565b60ff19831686528115158202860193506200054d565b6200052285620004b8565b5f5b83811015620005455781548189015260018201915060208101905062000524565b838801955050505b50505092915050565b5f620005638284620004ca565b915081905092915050565b5f819050919050565b62000582816200056e565b82525050565b620005938162000280565b82525050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620005c48262000599565b9050919050565b620005d681620005b8565b82525050565b5f60a082019050620005f15f83018862000577565b62000600602083018762000577565b6200060f604083018662000577565b6200061e606083018562000588565b6200062d6080830184620005cb565b9695505050505050565b60805160a05160c0516115b0620006625f395f6106d301525f61069f01525f61067a01526115b05ff3fe608060405234801561000f575f80fd5b50600436106100cd575f3560e01c806340c10f191161008a57806395d89b411161006457806395d89b4114610225578063a9059cbb14610243578063d505accf14610273578063dd62ed3e1461028f576100cd565b806340c10f19146101a957806370a08231146101c55780637ecebe00146101f5576100cd565b806306fdde03146100d1578063095ea7b3146100ef57806318160ddd1461011f57806323b872dd1461013d578063313ce5671461016d5780633644e5151461018b575b5f80fd5b6100d96102bf565b6040516100e69190610e03565b60405180910390f35b61010960048036038101906101049190610eb4565b61034a565b6040516101169190610f0c565b60405180910390f35b610127610437565b6040516101349190610f34565b60405180910390f35b61015760048036038101906101529190610f4d565b61043d565b6040516101649190610f0c565b60405180910390f35b610175610678565b6040516101829190610fb8565b60405180910390f35b61019361069c565b6040516101a09190610fe9565b60405180910390f35b6101c360048036038101906101be9190610eb4565b6106f8565b005b6101df60048036038101906101da9190611002565b610754565b6040516101ec9190610f34565b60405180910390f35b61020f600480360381019061020a9190611002565b610769565b60405161021c9190610f34565b60405180910390f35b61022d61077e565b60405161023a9190610e03565b60405180910390f35b61025d60048036038101906102589190610eb4565b61080a565b60405161026a9190610f0c565b60405180910390f35b61028d60048036038101906102889190611081565b610917565b005b6102a960048036038101906102a4919061111e565b610c04565b6040516102b69190610f34565b60405180910390f35b5f80546102cb90611189565b80601f01602080910402602001604051908101604052809291908181526020018280546102f790611189565b80156103425780601f1061031957610100808354040283529160200191610342565b820191905f5260205f20905b81548152906001019060200180831161032557829003601f168201915b505050505081565b5f8160045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516104259190610f34565b60405180910390a36001905092915050565b60025481565b5f8060045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461056a5782816104ed91906111e6565b60045f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505b8260035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546105b691906111e6565b925050819055508260035f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040516106649190610f34565b60405180910390a360019150509392505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f7f000000000000000000000000000000000000000000000000000000000000000046146106d1576106cc610c24565b6106f3565b7f00000000000000000000000000000000000000000000000000000000000000005b905090565b6107028282610cae565b8173ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885826040516107489190610f34565b60405180910390a25050565b6003602052805f5260405f205f915090505481565b6005602052805f5260405f205f915090505481565b6001805461078b90611189565b80601f01602080910402602001604051908101604052809291908181526020018280546107b790611189565b80156108025780601f106107d957610100808354040283529160200191610802565b820191905f5260205f20905b8154815290600101906020018083116107e557829003601f168201915b505050505081565b5f8160035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461085791906111e6565b925050819055508160035f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516109059190610f34565b60405180910390a36001905092915050565b4284101561095a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161095190611263565b60405180910390fd5b5f600161096561069c565b7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98a8a8a60055f8f73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f815480929190600101919050558b6040516020016109ea96959493929190611290565b60405160208183030381529060405280519060200120604051602001610a11929190611363565b604051602081830303815290604052805190602001208585856040515f8152602001604052604051610a469493929190611399565b6020604051602081039080840390855afa158015610a66573d5f803e3d5ffd5b5050506020604051035190505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614158015610ad957508773ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b610b18576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b0f90611426565b60405180910390fd5b8560045f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550508573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92587604051610bf39190610f34565b60405180910390a350505050505050565b6004602052815f5260405f20602052805f5260405f205f91509150505481565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f5f604051610c5491906114e0565b60405180910390207fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc64630604051602001610c939594939291906114f6565b60405160208183030381529060405280519060200120905090565b8060025f828254610cbf9190611547565b925050819055508060035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055508173ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610d6d9190610f34565b60405180910390a35050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610db0578082015181840152602081019050610d95565b5f8484015250505050565b5f601f19601f8301169050919050565b5f610dd582610d79565b610ddf8185610d83565b9350610def818560208601610d93565b610df881610dbb565b840191505092915050565b5f6020820190508181035f830152610e1b8184610dcb565b905092915050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610e5082610e27565b9050919050565b610e6081610e46565b8114610e6a575f80fd5b50565b5f81359050610e7b81610e57565b92915050565b5f819050919050565b610e9381610e81565b8114610e9d575f80fd5b50565b5f81359050610eae81610e8a565b92915050565b5f8060408385031215610eca57610ec9610e23565b5b5f610ed785828601610e6d565b9250506020610ee885828601610ea0565b9150509250929050565b5f8115159050919050565b610f0681610ef2565b82525050565b5f602082019050610f1f5f830184610efd565b92915050565b610f2e81610e81565b82525050565b5f602082019050610f475f830184610f25565b92915050565b5f805f60608486031215610f6457610f63610e23565b5b5f610f7186828701610e6d565b9350506020610f8286828701610e6d565b9250506040610f9386828701610ea0565b9150509250925092565b5f60ff82169050919050565b610fb281610f9d565b82525050565b5f602082019050610fcb5f830184610fa9565b92915050565b5f819050919050565b610fe381610fd1565b82525050565b5f602082019050610ffc5f830184610fda565b92915050565b5f6020828403121561101757611016610e23565b5b5f61102484828501610e6d565b91505092915050565b61103681610f9d565b8114611040575f80fd5b50565b5f813590506110518161102d565b92915050565b61106081610fd1565b811461106a575f80fd5b50565b5f8135905061107b81611057565b92915050565b5f805f805f805f60e0888a03121561109c5761109b610e23565b5b5f6110a98a828b01610e6d565b97505060206110ba8a828b01610e6d565b96505060406110cb8a828b01610ea0565b95505060606110dc8a828b01610ea0565b94505060806110ed8a828b01611043565b93505060a06110fe8a828b0161106d565b92505060c061110f8a828b0161106d565b91505092959891949750929550565b5f806040838503121561113457611133610e23565b5b5f61114185828601610e6d565b925050602061115285828601610e6d565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806111a057607f821691505b6020821081036111b3576111b261115c565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6111f082610e81565b91506111fb83610e81565b9250828203905081811115611213576112126111b9565b5b92915050565b7f5045524d49545f444541444c494e455f455850495245440000000000000000005f82015250565b5f61124d601783610d83565b915061125882611219565b602082019050919050565b5f6020820190508181035f83015261127a81611241565b9050919050565b61128a81610e46565b82525050565b5f60c0820190506112a35f830189610fda565b6112b06020830188611281565b6112bd6040830187611281565b6112ca6060830186610f25565b6112d76080830185610f25565b6112e460a0830184610f25565b979650505050505050565b5f81905092915050565b7f19010000000000000000000000000000000000000000000000000000000000005f82015250565b5f61132d6002836112ef565b9150611338826112f9565b600282019050919050565b5f819050919050565b61135d61135882610fd1565b611343565b82525050565b5f61136d82611321565b9150611379828561134c565b602082019150611389828461134c565b6020820191508190509392505050565b5f6080820190506113ac5f830187610fda565b6113b96020830186610fa9565b6113c66040830185610fda565b6113d36060830184610fda565b95945050505050565b7f494e56414c49445f5349474e45520000000000000000000000000000000000005f82015250565b5f611410600e83610d83565b915061141b826113dc565b602082019050919050565b5f6020820190508181035f83015261143d81611404565b9050919050565b5f81905092915050565b5f819050815f5260205f209050919050565b5f815461146c81611189565b6114768186611444565b9450600182165f811461149057600181146114a5576114d7565b60ff19831686528115158202860193506114d7565b6114ae8561144e565b5f5b838110156114cf578154818901526001820191506020810190506114b0565b838801955050505b50505092915050565b5f6114eb8284611460565b915081905092915050565b5f60a0820190506115095f830188610fda565b6115166020830187610fda565b6115236040830186610fda565b6115306060830185610f25565b61153d6080830184611281565b9695505050505050565b5f61155182610e81565b915061155c83610e81565b9250828201905080821115611574576115736111b9565b5b9291505056fea264697066735822122055bf768bf36436d5a69a17c7d948e07a55cbe33209947755d748cbcca08a2c5e64736f6c63430008170033",
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
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SolmateERC20 *SolmateERC20Filterer) ParseApproval(log types.Log) (*SolmateERC20Approval, error) {
	event := new(SolmateERC20Approval)
	if err := _SolmateERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SolmateERC20 *SolmateERC20Filterer) ParseTransfer(log types.Log) (*SolmateERC20Transfer, error) {
	event := new(SolmateERC20Transfer)
	if err := _SolmateERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
