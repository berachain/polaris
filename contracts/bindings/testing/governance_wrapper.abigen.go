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

// IGovernanceModuleCoin is an auto generated low-level Go binding around an user-defined struct.
type IGovernanceModuleCoin struct {
	Amount uint64
	Denom  string
}

// IGovernanceModuleProposal is an auto generated low-level Go binding around an user-defined struct.
type IGovernanceModuleProposal struct {
	Id               uint64
	Message          []byte
	Status           int32
	FinalTallyResult IGovernanceModuleTallyResult
	SubmitTime       uint64
	DepositEndTime   uint64
	TotalDeposit     []IGovernanceModuleCoin
	VotingStartTime  uint64
	VotingEndTime    uint64
	Metadata         string
	Title            string
	Summary          string
	Proposer         string
}

// IGovernanceModuleTallyResult is an auto generated low-level Go binding around an user-defined struct.
type IGovernanceModuleTallyResult struct {
	YesCount        string
	AbstainCount    string
	NoCount         string
	NoWithVetoCount string
}

// GovernanceWrapperMetaData contains all meta data concerning the GovernanceWrapper contract.
var GovernanceWrapperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governanceModule\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.TallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"submitTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositEndTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Coin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"votingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposer\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Proposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"proposalStatus\",\"type\":\"int32\"}],\"name\":\"getProposals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.TallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"submitTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositEndTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Coin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"votingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposer\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Proposal[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceModule\",\"outputs\":[{\"internalType\":\"contractIGovernanceModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proposal\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"submitProposalWrapepr\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"internalType\":\"int32\",\"name\":\"option\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001a1f38038062001a1f83398181016040528101906200003791906200014f565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036200009e576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000181565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200011782620000ea565b9050919050565b62000129816200010a565b81146200013557600080fd5b50565b60008151905062000149816200011e565b92915050565b600060208284031215620001685762000167620000e5565b5b6000620001788482850162000138565b91505092915050565b61188e80620001916000396000f3fe60806040526004361061004e5760003560e01c806319f7a0fb1461005a5780632b0a703214610097578063b5828df2146100c2578063f1610a28146100ff578063fa4204cb1461013c57610055565b3661005557005b600080fd5b34801561006657600080fd5b50610081600480360381019061007c9190610725565b61016c565b60405161008e91906107af565b60405180910390f35b3480156100a357600080fd5b506100ac610218565b6040516100b99190610849565b60405180910390f35b3480156100ce57600080fd5b506100e960048036038101906100e49190610864565b61023c565b6040516100f69190610cff565b60405180910390f35b34801561010b57600080fd5b5061012660048036038101906101219190610d21565b6102e4565b6040516101339190610e91565b60405180910390f35b61015660048036038101906101519190610f13565b610392565b6040516101639190610fa3565b60405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166319f7a0fb8585856040518463ffffffff1660e01b81526004016101cc93929190611017565b6020604051808303816000875af11580156101eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061020f9190611081565b90509392505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b606060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b5828df2836040518263ffffffff1660e01b815260040161029791906110ae565b600060405180830381865afa1580156102b4573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906102dd9190611720565b9050919050565b6102ec610487565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f1610a28836040518263ffffffff1660e01b81526004016103459190610fa3565b600060405180830381865afa158015610362573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061038b9190611769565b9050919050565b60003073ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f193505050501580156103da573d6000803e3d6000fd5b5060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663474d7f35868686866040518563ffffffff1660e01b815260040161043a94939291906117f0565b6020604051808303816000875af1158015610459573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061047d919061182b565b9050949350505050565b604051806101a00160405280600067ffffffffffffffff16815260200160608152602001600060030b81526020016104bd61052a565b8152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001606081526020016060815260200160608152602001606081525090565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000604051905090565b600080fd5b600080fd5b600067ffffffffffffffff82169050919050565b61058381610566565b811461058e57600080fd5b50565b6000813590506105a08161057a565b92915050565b60008160030b9050919050565b6105bc816105a6565b81146105c757600080fd5b50565b6000813590506105d9816105b3565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610632826105e9565b810181811067ffffffffffffffff82111715610651576106506105fa565b5b80604052505050565b6000610664610552565b90506106708282610629565b919050565b600067ffffffffffffffff8211156106905761068f6105fa565b5b610699826105e9565b9050602081019050919050565b82818337600083830152505050565b60006106c86106c384610675565b61065a565b9050828152602081018484840111156106e4576106e36105e4565b5b6106ef8482856106a6565b509392505050565b600082601f83011261070c5761070b6105df565b5b813561071c8482602086016106b5565b91505092915050565b60008060006060848603121561073e5761073d61055c565b5b600061074c86828701610591565b935050602061075d868287016105ca565b925050604084013567ffffffffffffffff81111561077e5761077d610561565b5b61078a868287016106f7565b9150509250925092565b60008115159050919050565b6107a981610794565b82525050565b60006020820190506107c460008301846107a0565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061080f61080a610805846107ca565b6107ea565b6107ca565b9050919050565b6000610821826107f4565b9050919050565b600061083382610816565b9050919050565b61084381610828565b82525050565b600060208201905061085e600083018461083a565b92915050565b60006020828403121561087a5761087961055c565b5b6000610888848285016105ca565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6108c681610566565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b838110156109065780820151818401526020810190506108eb565b60008484015250505050565b600061091d826108cc565b61092781856108d7565b93506109378185602086016108e8565b610940816105e9565b840191505092915050565b610954816105a6565b82525050565b600081519050919050565b600082825260208201905092915050565b60006109818261095a565b61098b8185610965565b935061099b8185602086016108e8565b6109a4816105e9565b840191505092915050565b600060808301600083015184820360008601526109cc8282610976565b915050602083015184820360208601526109e68282610976565b91505060408301518482036040860152610a008282610976565b91505060608301518482036060860152610a1a8282610976565b9150508091505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000604083016000830151610a6b60008601826108bd565b5060208301518482036020860152610a838282610976565b9150508091505092915050565b6000610a9c8383610a53565b905092915050565b6000602082019050919050565b6000610abc82610a27565b610ac68185610a32565b935083602082028501610ad885610a43565b8060005b85811015610b145784840389528151610af58582610a90565b9450610b0083610aa4565b925060208a01995050600181019050610adc565b50829750879550505050505092915050565b60006101a083016000830151610b3f60008601826108bd565b5060208301518482036020860152610b578282610912565b9150506040830151610b6c604086018261094b565b5060608301518482036060860152610b8482826109af565b9150506080830151610b9960808601826108bd565b5060a0830151610bac60a08601826108bd565b5060c083015184820360c0860152610bc48282610ab1565b91505060e0830151610bd960e08601826108bd565b50610100830151610bee6101008601826108bd565b50610120830151848203610120860152610c088282610976565b915050610140830151848203610140860152610c248282610976565b915050610160830151848203610160860152610c408282610976565b915050610180830151848203610180860152610c5c8282610976565b9150508091505092915050565b6000610c758383610b26565b905092915050565b6000602082019050919050565b6000610c9582610891565b610c9f818561089c565b935083602082028501610cb1856108ad565b8060005b85811015610ced5784840389528151610cce8582610c69565b9450610cd983610c7d565b925060208a01995050600181019050610cb5565b50829750879550505050505092915050565b60006020820190508181036000830152610d198184610c8a565b905092915050565b600060208284031215610d3757610d3661055c565b5b6000610d4584828501610591565b91505092915050565b60006101a083016000830151610d6760008601826108bd565b5060208301518482036020860152610d7f8282610912565b9150506040830151610d94604086018261094b565b5060608301518482036060860152610dac82826109af565b9150506080830151610dc160808601826108bd565b5060a0830151610dd460a08601826108bd565b5060c083015184820360c0860152610dec8282610ab1565b91505060e0830151610e0160e08601826108bd565b50610100830151610e166101008601826108bd565b50610120830151848203610120860152610e308282610976565b915050610140830151848203610140860152610e4c8282610976565b915050610160830151848203610160860152610e688282610976565b915050610180830151848203610180860152610e848282610976565b9150508091505092915050565b60006020820190508181036000830152610eab8184610d4e565b905092915050565b600080fd5b600080fd5b60008083601f840112610ed357610ed26105df565b5b8235905067ffffffffffffffff811115610ef057610eef610eb3565b5b602083019150836001820283011115610f0c57610f0b610eb8565b5b9250929050565b60008060008060408587031215610f2d57610f2c61055c565b5b600085013567ffffffffffffffff811115610f4b57610f4a610561565b5b610f5787828801610ebd565b9450945050602085013567ffffffffffffffff811115610f7a57610f79610561565b5b610f8687828801610ebd565b925092505092959194509250565b610f9d81610566565b82525050565b6000602082019050610fb86000830184610f94565b92915050565b610fc7816105a6565b82525050565b600082825260208201905092915050565b6000610fe98261095a565b610ff38185610fcd565b93506110038185602086016108e8565b61100c816105e9565b840191505092915050565b600060608201905061102c6000830186610f94565b6110396020830185610fbe565b818103604083015261104b8184610fde565b9050949350505050565b61105e81610794565b811461106957600080fd5b50565b60008151905061107b81611055565b92915050565b6000602082840312156110975761109661055c565b5b60006110a58482850161106c565b91505092915050565b60006020820190506110c36000830184610fbe565b92915050565b600067ffffffffffffffff8211156110e4576110e36105fa565b5b602082029050602081019050919050565b600080fd5b600080fd5b60008151905061110e8161057a565b92915050565b600067ffffffffffffffff82111561112f5761112e6105fa565b5b611138826105e9565b9050602081019050919050565b600061115861115384611114565b61065a565b905082815260208101848484011115611174576111736105e4565b5b61117f8482856108e8565b509392505050565b600082601f83011261119c5761119b6105df565b5b81516111ac848260208601611145565b91505092915050565b6000815190506111c4816105b3565b92915050565b60006111dd6111d884610675565b61065a565b9050828152602081018484840111156111f9576111f86105e4565b5b6112048482856108e8565b509392505050565b600082601f830112611221576112206105df565b5b81516112318482602086016111ca565b91505092915050565b6000608082840312156112505761124f6110f5565b5b61125a608061065a565b9050600082015167ffffffffffffffff81111561127a576112796110fa565b5b6112868482850161120c565b600083015250602082015167ffffffffffffffff8111156112aa576112a96110fa565b5b6112b68482850161120c565b602083015250604082015167ffffffffffffffff8111156112da576112d96110fa565b5b6112e68482850161120c565b604083015250606082015167ffffffffffffffff81111561130a576113096110fa565b5b6113168482850161120c565b60608301525092915050565b600067ffffffffffffffff82111561133d5761133c6105fa565b5b602082029050602081019050919050565b600060408284031215611364576113636110f5565b5b61136e604061065a565b9050600061137e848285016110ff565b600083015250602082015167ffffffffffffffff8111156113a2576113a16110fa565b5b6113ae8482850161120c565b60208301525092915050565b60006113cd6113c884611322565b61065a565b905080838252602082019050602084028301858111156113f0576113ef610eb8565b5b835b8181101561143757805167ffffffffffffffff811115611415576114146105df565b5b808601611422898261134e565b855260208501945050506020810190506113f2565b5050509392505050565b600082601f830112611456576114556105df565b5b81516114668482602086016113ba565b91505092915050565b60006101a08284031215611486576114856110f5565b5b6114916101a061065a565b905060006114a1848285016110ff565b600083015250602082015167ffffffffffffffff8111156114c5576114c46110fa565b5b6114d184828501611187565b60208301525060406114e5848285016111b5565b604083015250606082015167ffffffffffffffff811115611509576115086110fa565b5b6115158482850161123a565b6060830152506080611529848285016110ff565b60808301525060a061153d848285016110ff565b60a08301525060c082015167ffffffffffffffff811115611561576115606110fa565b5b61156d84828501611441565b60c08301525060e0611581848285016110ff565b60e083015250610100611596848285016110ff565b6101008301525061012082015167ffffffffffffffff8111156115bc576115bb6110fa565b5b6115c88482850161120c565b6101208301525061014082015167ffffffffffffffff8111156115ee576115ed6110fa565b5b6115fa8482850161120c565b6101408301525061016082015167ffffffffffffffff8111156116205761161f6110fa565b5b61162c8482850161120c565b6101608301525061018082015167ffffffffffffffff811115611652576116516110fa565b5b61165e8482850161120c565b6101808301525092915050565b600061167e611679846110c9565b61065a565b905080838252602082019050602084028301858111156116a1576116a0610eb8565b5b835b818110156116e857805167ffffffffffffffff8111156116c6576116c56105df565b5b8086016116d3898261146f565b855260208501945050506020810190506116a3565b5050509392505050565b600082601f830112611707576117066105df565b5b815161171784826020860161166b565b91505092915050565b6000602082840312156117365761173561055c565b5b600082015167ffffffffffffffff81111561175457611753610561565b5b611760848285016116f2565b91505092915050565b60006020828403121561177f5761177e61055c565b5b600082015167ffffffffffffffff81111561179d5761179c610561565b5b6117a98482850161146f565b91505092915050565b600082825260208201905092915050565b60006117cf83856117b2565b93506117dc8385846106a6565b6117e5836105e9565b840190509392505050565b6000604082019050818103600083015261180b8186886117c3565b905081810360208301526118208184866117c3565b905095945050505050565b6000602082840312156118415761184061055c565b5b600061184f848285016110ff565b9150509291505056fea26469706673582212203c0343ba3da1783d08d15716dd47f36fb144bcc75bc7917ed16313c2e44ae9d264736f6c63430008130033",
}

// GovernanceWrapperABI is the input ABI used to generate the binding from.
// Deprecated: Use GovernanceWrapperMetaData.ABI instead.
var GovernanceWrapperABI = GovernanceWrapperMetaData.ABI

// GovernanceWrapperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GovernanceWrapperMetaData.Bin instead.
var GovernanceWrapperBin = GovernanceWrapperMetaData.Bin

// DeployGovernanceWrapper deploys a new Ethereum contract, binding an instance of GovernanceWrapper to it.
func DeployGovernanceWrapper(auth *bind.TransactOpts, backend bind.ContractBackend, _governanceModule common.Address) (common.Address, *types.Transaction, *GovernanceWrapper, error) {
	parsed, err := GovernanceWrapperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GovernanceWrapperBin), backend, _governanceModule)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GovernanceWrapper{GovernanceWrapperCaller: GovernanceWrapperCaller{contract: contract}, GovernanceWrapperTransactor: GovernanceWrapperTransactor{contract: contract}, GovernanceWrapperFilterer: GovernanceWrapperFilterer{contract: contract}}, nil
}

// GovernanceWrapper is an auto generated Go binding around an Ethereum contract.
type GovernanceWrapper struct {
	GovernanceWrapperCaller     // Read-only binding to the contract
	GovernanceWrapperTransactor // Write-only binding to the contract
	GovernanceWrapperFilterer   // Log filterer for contract events
}

// GovernanceWrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernanceWrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceWrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernanceWrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceWrapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernanceWrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceWrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernanceWrapperSession struct {
	Contract     *GovernanceWrapper // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GovernanceWrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernanceWrapperCallerSession struct {
	Contract *GovernanceWrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// GovernanceWrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernanceWrapperTransactorSession struct {
	Contract     *GovernanceWrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// GovernanceWrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernanceWrapperRaw struct {
	Contract *GovernanceWrapper // Generic contract binding to access the raw methods on
}

// GovernanceWrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernanceWrapperCallerRaw struct {
	Contract *GovernanceWrapperCaller // Generic read-only contract binding to access the raw methods on
}

// GovernanceWrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernanceWrapperTransactorRaw struct {
	Contract *GovernanceWrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernanceWrapper creates a new instance of GovernanceWrapper, bound to a specific deployed contract.
func NewGovernanceWrapper(address common.Address, backend bind.ContractBackend) (*GovernanceWrapper, error) {
	contract, err := bindGovernanceWrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GovernanceWrapper{GovernanceWrapperCaller: GovernanceWrapperCaller{contract: contract}, GovernanceWrapperTransactor: GovernanceWrapperTransactor{contract: contract}, GovernanceWrapperFilterer: GovernanceWrapperFilterer{contract: contract}}, nil
}

// NewGovernanceWrapperCaller creates a new read-only instance of GovernanceWrapper, bound to a specific deployed contract.
func NewGovernanceWrapperCaller(address common.Address, caller bind.ContractCaller) (*GovernanceWrapperCaller, error) {
	contract, err := bindGovernanceWrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceWrapperCaller{contract: contract}, nil
}

// NewGovernanceWrapperTransactor creates a new write-only instance of GovernanceWrapper, bound to a specific deployed contract.
func NewGovernanceWrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernanceWrapperTransactor, error) {
	contract, err := bindGovernanceWrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceWrapperTransactor{contract: contract}, nil
}

// NewGovernanceWrapperFilterer creates a new log filterer instance of GovernanceWrapper, bound to a specific deployed contract.
func NewGovernanceWrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernanceWrapperFilterer, error) {
	contract, err := bindGovernanceWrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernanceWrapperFilterer{contract: contract}, nil
}

// bindGovernanceWrapper binds a generic wrapper to an already deployed contract.
func bindGovernanceWrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GovernanceWrapperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernanceWrapper *GovernanceWrapperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GovernanceWrapper.Contract.GovernanceWrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernanceWrapper *GovernanceWrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceWrapper.Contract.GovernanceWrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernanceWrapper *GovernanceWrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernanceWrapper.Contract.GovernanceWrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernanceWrapper *GovernanceWrapperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GovernanceWrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernanceWrapper *GovernanceWrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceWrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernanceWrapper *GovernanceWrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernanceWrapper.Contract.contract.Transact(opts, method, params...)
}

// GetProposal is a free data retrieval call binding the contract method 0xf1610a28.
//
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint64,string)[],uint64,uint64,string,string,string,string))
func (_GovernanceWrapper *GovernanceWrapperCaller) GetProposal(opts *bind.CallOpts, proposalId uint64) (IGovernanceModuleProposal, error) {
	var out []interface{}
	err := _GovernanceWrapper.contract.Call(opts, &out, "getProposal", proposalId)

	if err != nil {
		return *new(IGovernanceModuleProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(IGovernanceModuleProposal)).(*IGovernanceModuleProposal)

	return out0, err

}

// GetProposal is a free data retrieval call binding the contract method 0xf1610a28.
//
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint64,string)[],uint64,uint64,string,string,string,string))
func (_GovernanceWrapper *GovernanceWrapperSession) GetProposal(proposalId uint64) (IGovernanceModuleProposal, error) {
	return _GovernanceWrapper.Contract.GetProposal(&_GovernanceWrapper.CallOpts, proposalId)
}

// GetProposal is a free data retrieval call binding the contract method 0xf1610a28.
//
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint64,string)[],uint64,uint64,string,string,string,string))
func (_GovernanceWrapper *GovernanceWrapperCallerSession) GetProposal(proposalId uint64) (IGovernanceModuleProposal, error) {
	return _GovernanceWrapper.Contract.GetProposal(&_GovernanceWrapper.CallOpts, proposalId)
}

// GetProposals is a free data retrieval call binding the contract method 0xb5828df2.
//
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint64,string)[],uint64,uint64,string,string,string,string)[])
func (_GovernanceWrapper *GovernanceWrapperCaller) GetProposals(opts *bind.CallOpts, proposalStatus int32) ([]IGovernanceModuleProposal, error) {
	var out []interface{}
	err := _GovernanceWrapper.contract.Call(opts, &out, "getProposals", proposalStatus)

	if err != nil {
		return *new([]IGovernanceModuleProposal), err
	}

	out0 := *abi.ConvertType(out[0], new([]IGovernanceModuleProposal)).(*[]IGovernanceModuleProposal)

	return out0, err

}

// GetProposals is a free data retrieval call binding the contract method 0xb5828df2.
//
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint64,string)[],uint64,uint64,string,string,string,string)[])
func (_GovernanceWrapper *GovernanceWrapperSession) GetProposals(proposalStatus int32) ([]IGovernanceModuleProposal, error) {
	return _GovernanceWrapper.Contract.GetProposals(&_GovernanceWrapper.CallOpts, proposalStatus)
}

// GetProposals is a free data retrieval call binding the contract method 0xb5828df2.
//
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint64,string)[],uint64,uint64,string,string,string,string)[])
func (_GovernanceWrapper *GovernanceWrapperCallerSession) GetProposals(proposalStatus int32) ([]IGovernanceModuleProposal, error) {
	return _GovernanceWrapper.Contract.GetProposals(&_GovernanceWrapper.CallOpts, proposalStatus)
}

// GovernanceModule is a free data retrieval call binding the contract method 0x2b0a7032.
//
// Solidity: function governanceModule() view returns(address)
func (_GovernanceWrapper *GovernanceWrapperCaller) GovernanceModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GovernanceWrapper.contract.Call(opts, &out, "governanceModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceModule is a free data retrieval call binding the contract method 0x2b0a7032.
//
// Solidity: function governanceModule() view returns(address)
func (_GovernanceWrapper *GovernanceWrapperSession) GovernanceModule() (common.Address, error) {
	return _GovernanceWrapper.Contract.GovernanceModule(&_GovernanceWrapper.CallOpts)
}

// GovernanceModule is a free data retrieval call binding the contract method 0x2b0a7032.
//
// Solidity: function governanceModule() view returns(address)
func (_GovernanceWrapper *GovernanceWrapperCallerSession) GovernanceModule() (common.Address, error) {
	return _GovernanceWrapper.Contract.GovernanceModule(&_GovernanceWrapper.CallOpts)
}

// SubmitProposalWrapepr is a paid mutator transaction binding the contract method 0xfa4204cb.
//
// Solidity: function submitProposalWrapepr(bytes proposal, bytes message) payable returns(uint64)
func (_GovernanceWrapper *GovernanceWrapperTransactor) SubmitProposalWrapepr(opts *bind.TransactOpts, proposal []byte, message []byte) (*types.Transaction, error) {
	return _GovernanceWrapper.contract.Transact(opts, "submitProposalWrapepr", proposal, message)
}

// SubmitProposalWrapepr is a paid mutator transaction binding the contract method 0xfa4204cb.
//
// Solidity: function submitProposalWrapepr(bytes proposal, bytes message) payable returns(uint64)
func (_GovernanceWrapper *GovernanceWrapperSession) SubmitProposalWrapepr(proposal []byte, message []byte) (*types.Transaction, error) {
	return _GovernanceWrapper.Contract.SubmitProposalWrapepr(&_GovernanceWrapper.TransactOpts, proposal, message)
}

// SubmitProposalWrapepr is a paid mutator transaction binding the contract method 0xfa4204cb.
//
// Solidity: function submitProposalWrapepr(bytes proposal, bytes message) payable returns(uint64)
func (_GovernanceWrapper *GovernanceWrapperTransactorSession) SubmitProposalWrapepr(proposal []byte, message []byte) (*types.Transaction, error) {
	return _GovernanceWrapper.Contract.SubmitProposalWrapepr(&_GovernanceWrapper.TransactOpts, proposal, message)
}

// Vote is a paid mutator transaction binding the contract method 0x19f7a0fb.
//
// Solidity: function vote(uint64 proposalId, int32 option, string metadata) returns(bool)
func (_GovernanceWrapper *GovernanceWrapperTransactor) Vote(opts *bind.TransactOpts, proposalId uint64, option int32, metadata string) (*types.Transaction, error) {
	return _GovernanceWrapper.contract.Transact(opts, "vote", proposalId, option, metadata)
}

// Vote is a paid mutator transaction binding the contract method 0x19f7a0fb.
//
// Solidity: function vote(uint64 proposalId, int32 option, string metadata) returns(bool)
func (_GovernanceWrapper *GovernanceWrapperSession) Vote(proposalId uint64, option int32, metadata string) (*types.Transaction, error) {
	return _GovernanceWrapper.Contract.Vote(&_GovernanceWrapper.TransactOpts, proposalId, option, metadata)
}

// Vote is a paid mutator transaction binding the contract method 0x19f7a0fb.
//
// Solidity: function vote(uint64 proposalId, int32 option, string metadata) returns(bool)
func (_GovernanceWrapper *GovernanceWrapperTransactorSession) Vote(proposalId uint64, option int32, metadata string) (*types.Transaction, error) {
	return _GovernanceWrapper.Contract.Vote(&_GovernanceWrapper.TransactOpts, proposalId, option, metadata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GovernanceWrapper *GovernanceWrapperTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceWrapper.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GovernanceWrapper *GovernanceWrapperSession) Receive() (*types.Transaction, error) {
	return _GovernanceWrapper.Contract.Receive(&_GovernanceWrapper.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GovernanceWrapper *GovernanceWrapperTransactorSession) Receive() (*types.Transaction, error) {
	return _GovernanceWrapper.Contract.Receive(&_GovernanceWrapper.TransactOpts)
}
