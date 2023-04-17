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
	Amount *big.Int
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governanceModule\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"cancelProposal\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.TallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"submitTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositEndTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Coin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"votingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposer\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Proposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"proposalStatus\",\"type\":\"int32\"}],\"name\":\"getProposals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.TallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"submitTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositEndTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Coin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"votingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposer\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Proposal[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceModule\",\"outputs\":[{\"internalType\":\"contractIGovernanceModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proposal\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"submitProposalWrapepr\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"internalType\":\"int32\",\"name\":\"option\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001bbd38038062001bbd83398181016040528101906200003791906200014f565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036200009e576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000181565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200011782620000ea565b9050919050565b62000129816200010a565b81146200013557600080fd5b50565b60008151905062000149816200011e565b92915050565b600060208284031215620001685762000167620000e5565b5b6000620001788482850162000138565b91505092915050565b611a2c80620001916000396000f3fe6080604052600436106100595760003560e01c806319f7a0fb146100655780632b0a7032146100a257806337a9a59e146100cd578063b5828df21461010b578063f1610a2814610148578063fa4204cb1461018557610060565b3661006057005b600080fd5b34801561007157600080fd5b5061008c60048036038101906100879190610815565b6101b5565b604051610099919061089f565b60405180910390f35b3480156100ae57600080fd5b506100b7610261565b6040516100c49190610939565b60405180910390f35b3480156100d957600080fd5b506100f460048036038101906100ef9190610954565b610285565b604051610102929190610990565b60405180910390f35b34801561011757600080fd5b50610132600480360381019061012d91906109b9565b61032c565b60405161013f9190610e6d565b60405180910390f35b34801561015457600080fd5b5061016f600480360381019061016a9190610954565b6103d4565b60405161017c9190610fd2565b60405180910390f35b61019f600480360381019061019a9190611054565b610482565b6040516101ac91906110d5565b60405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166319f7a0fb8585856040518463ffffffff1660e01b815260040161021593929190611149565b6020604051808303816000875af1158015610234573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061025891906111b3565b90509392505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166337a9a59e846040518263ffffffff1660e01b81526004016102e191906110d5565b60408051808303816000875af11580156102ff573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061032391906111f5565b91509150915091565b606060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b5828df2836040518263ffffffff1660e01b81526004016103879190611235565b600060405180830381865afa1580156103a4573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906103cd91906118be565b9050919050565b6103dc610577565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f1610a28836040518263ffffffff1660e01b815260040161043591906110d5565b600060405180830381865afa158015610452573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061047b9190611907565b9050919050565b60003073ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f193505050501580156104ca573d6000803e3d6000fd5b5060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663474d7f35868686866040518563ffffffff1660e01b815260040161052a949392919061198e565b6020604051808303816000875af1158015610549573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061056d91906119c9565b9050949350505050565b604051806101a00160405280600067ffffffffffffffff16815260200160608152602001600060030b81526020016105ad61061a565b8152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001606081526020016060815260200160608152602001606081525090565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000604051905090565b600080fd5b600080fd5b600067ffffffffffffffff82169050919050565b61067381610656565b811461067e57600080fd5b50565b6000813590506106908161066a565b92915050565b60008160030b9050919050565b6106ac81610696565b81146106b757600080fd5b50565b6000813590506106c9816106a3565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610722826106d9565b810181811067ffffffffffffffff82111715610741576107406106ea565b5b80604052505050565b6000610754610642565b90506107608282610719565b919050565b600067ffffffffffffffff8211156107805761077f6106ea565b5b610789826106d9565b9050602081019050919050565b82818337600083830152505050565b60006107b86107b384610765565b61074a565b9050828152602081018484840111156107d4576107d36106d4565b5b6107df848285610796565b509392505050565b600082601f8301126107fc576107fb6106cf565b5b813561080c8482602086016107a5565b91505092915050565b60008060006060848603121561082e5761082d61064c565b5b600061083c86828701610681565b935050602061084d868287016106ba565b925050604084013567ffffffffffffffff81111561086e5761086d610651565b5b61087a868287016107e7565b9150509250925092565b60008115159050919050565b61089981610884565b82525050565b60006020820190506108b46000830184610890565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006108ff6108fa6108f5846108ba565b6108da565b6108ba565b9050919050565b6000610911826108e4565b9050919050565b600061092382610906565b9050919050565b61093381610918565b82525050565b600060208201905061094e600083018461092a565b92915050565b60006020828403121561096a5761096961064c565b5b600061097884828501610681565b91505092915050565b61098a81610656565b82525050565b60006040820190506109a56000830185610981565b6109b26020830184610981565b9392505050565b6000602082840312156109cf576109ce61064c565b5b60006109dd848285016106ba565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610a1b81610656565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610a5b578082015181840152602081019050610a40565b60008484015250505050565b6000610a7282610a21565b610a7c8185610a2c565b9350610a8c818560208601610a3d565b610a95816106d9565b840191505092915050565b610aa981610696565b82525050565b600081519050919050565b600082825260208201905092915050565b6000610ad682610aaf565b610ae08185610aba565b9350610af0818560208601610a3d565b610af9816106d9565b840191505092915050565b60006080830160008301518482036000860152610b218282610acb565b91505060208301518482036020860152610b3b8282610acb565b91505060408301518482036040860152610b558282610acb565b91505060608301518482036060860152610b6f8282610acb565b9150508091505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000819050919050565b610bbb81610ba8565b82525050565b6000604083016000830151610bd96000860182610bb2565b5060208301518482036020860152610bf18282610acb565b9150508091505092915050565b6000610c0a8383610bc1565b905092915050565b6000602082019050919050565b6000610c2a82610b7c565b610c348185610b87565b935083602082028501610c4685610b98565b8060005b85811015610c825784840389528151610c638582610bfe565b9450610c6e83610c12565b925060208a01995050600181019050610c4a565b50829750879550505050505092915050565b60006101a083016000830151610cad6000860182610a12565b5060208301518482036020860152610cc58282610a67565b9150506040830151610cda6040860182610aa0565b5060608301518482036060860152610cf28282610b04565b9150506080830151610d076080860182610a12565b5060a0830151610d1a60a0860182610a12565b5060c083015184820360c0860152610d328282610c1f565b91505060e0830151610d4760e0860182610a12565b50610100830151610d5c610100860182610a12565b50610120830151848203610120860152610d768282610acb565b915050610140830151848203610140860152610d928282610acb565b915050610160830151848203610160860152610dae8282610acb565b915050610180830151848203610180860152610dca8282610acb565b9150508091505092915050565b6000610de38383610c94565b905092915050565b6000602082019050919050565b6000610e03826109e6565b610e0d81856109f1565b935083602082028501610e1f85610a02565b8060005b85811015610e5b5784840389528151610e3c8582610dd7565b9450610e4783610deb565b925060208a01995050600181019050610e23565b50829750879550505050505092915050565b60006020820190508181036000830152610e878184610df8565b905092915050565b60006101a083016000830151610ea86000860182610a12565b5060208301518482036020860152610ec08282610a67565b9150506040830151610ed56040860182610aa0565b5060608301518482036060860152610eed8282610b04565b9150506080830151610f026080860182610a12565b5060a0830151610f1560a0860182610a12565b5060c083015184820360c0860152610f2d8282610c1f565b91505060e0830151610f4260e0860182610a12565b50610100830151610f57610100860182610a12565b50610120830151848203610120860152610f718282610acb565b915050610140830151848203610140860152610f8d8282610acb565b915050610160830151848203610160860152610fa98282610acb565b915050610180830151848203610180860152610fc58282610acb565b9150508091505092915050565b60006020820190508181036000830152610fec8184610e8f565b905092915050565b600080fd5b600080fd5b60008083601f840112611014576110136106cf565b5b8235905067ffffffffffffffff81111561103157611030610ff4565b5b60208301915083600182028301111561104d5761104c610ff9565b5b9250929050565b6000806000806040858703121561106e5761106d61064c565b5b600085013567ffffffffffffffff81111561108c5761108b610651565b5b61109887828801610ffe565b9450945050602085013567ffffffffffffffff8111156110bb576110ba610651565b5b6110c787828801610ffe565b925092505092959194509250565b60006020820190506110ea6000830184610981565b92915050565b6110f981610696565b82525050565b600082825260208201905092915050565b600061111b82610aaf565b61112581856110ff565b9350611135818560208601610a3d565b61113e816106d9565b840191505092915050565b600060608201905061115e6000830186610981565b61116b60208301856110f0565b818103604083015261117d8184611110565b9050949350505050565b61119081610884565b811461119b57600080fd5b50565b6000815190506111ad81611187565b92915050565b6000602082840312156111c9576111c861064c565b5b60006111d78482850161119e565b91505092915050565b6000815190506111ef8161066a565b92915050565b6000806040838503121561120c5761120b61064c565b5b600061121a858286016111e0565b925050602061122b858286016111e0565b9150509250929050565b600060208201905061124a60008301846110f0565b92915050565b600067ffffffffffffffff82111561126b5761126a6106ea565b5b602082029050602081019050919050565b600080fd5b600080fd5b600067ffffffffffffffff8211156112a1576112a06106ea565b5b6112aa826106d9565b9050602081019050919050565b60006112ca6112c584611286565b61074a565b9050828152602081018484840111156112e6576112e56106d4565b5b6112f1848285610a3d565b509392505050565b600082601f83011261130e5761130d6106cf565b5b815161131e8482602086016112b7565b91505092915050565b600081519050611336816106a3565b92915050565b600061134f61134a84610765565b61074a565b90508281526020810184848401111561136b5761136a6106d4565b5b611376848285610a3d565b509392505050565b600082601f830112611393576113926106cf565b5b81516113a384826020860161133c565b91505092915050565b6000608082840312156113c2576113c161127c565b5b6113cc608061074a565b9050600082015167ffffffffffffffff8111156113ec576113eb611281565b5b6113f88482850161137e565b600083015250602082015167ffffffffffffffff81111561141c5761141b611281565b5b6114288482850161137e565b602083015250604082015167ffffffffffffffff81111561144c5761144b611281565b5b6114588482850161137e565b604083015250606082015167ffffffffffffffff81111561147c5761147b611281565b5b6114888482850161137e565b60608301525092915050565b600067ffffffffffffffff8211156114af576114ae6106ea565b5b602082029050602081019050919050565b6114c981610ba8565b81146114d457600080fd5b50565b6000815190506114e6816114c0565b92915050565b6000604082840312156115025761150161127c565b5b61150c604061074a565b9050600061151c848285016114d7565b600083015250602082015167ffffffffffffffff8111156115405761153f611281565b5b61154c8482850161137e565b60208301525092915050565b600061156b61156684611494565b61074a565b9050808382526020820190506020840283018581111561158e5761158d610ff9565b5b835b818110156115d557805167ffffffffffffffff8111156115b3576115b26106cf565b5b8086016115c089826114ec565b85526020850194505050602081019050611590565b5050509392505050565b600082601f8301126115f4576115f36106cf565b5b8151611604848260208601611558565b91505092915050565b60006101a082840312156116245761162361127c565b5b61162f6101a061074a565b9050600061163f848285016111e0565b600083015250602082015167ffffffffffffffff81111561166357611662611281565b5b61166f848285016112f9565b602083015250604061168384828501611327565b604083015250606082015167ffffffffffffffff8111156116a7576116a6611281565b5b6116b3848285016113ac565b60608301525060806116c7848285016111e0565b60808301525060a06116db848285016111e0565b60a08301525060c082015167ffffffffffffffff8111156116ff576116fe611281565b5b61170b848285016115df565b60c08301525060e061171f848285016111e0565b60e083015250610100611734848285016111e0565b6101008301525061012082015167ffffffffffffffff81111561175a57611759611281565b5b6117668482850161137e565b6101208301525061014082015167ffffffffffffffff81111561178c5761178b611281565b5b6117988482850161137e565b6101408301525061016082015167ffffffffffffffff8111156117be576117bd611281565b5b6117ca8482850161137e565b6101608301525061018082015167ffffffffffffffff8111156117f0576117ef611281565b5b6117fc8482850161137e565b6101808301525092915050565b600061181c61181784611250565b61074a565b9050808382526020820190506020840283018581111561183f5761183e610ff9565b5b835b8181101561188657805167ffffffffffffffff811115611864576118636106cf565b5b808601611871898261160d565b85526020850194505050602081019050611841565b5050509392505050565b600082601f8301126118a5576118a46106cf565b5b81516118b5848260208601611809565b91505092915050565b6000602082840312156118d4576118d361064c565b5b600082015167ffffffffffffffff8111156118f2576118f1610651565b5b6118fe84828501611890565b91505092915050565b60006020828403121561191d5761191c61064c565b5b600082015167ffffffffffffffff81111561193b5761193a610651565b5b6119478482850161160d565b91505092915050565b600082825260208201905092915050565b600061196d8385611950565b935061197a838584610796565b611983836106d9565b840190509392505050565b600060408201905081810360008301526119a9818688611961565b905081810360208301526119be818486611961565b905095945050505050565b6000602082840312156119df576119de61064c565b5b60006119ed848285016111e0565b9150509291505056fea2646970667358221220626df970ac3f13d40d728f94ea8c8e22bd0458e60b08d406c081f4836daa2d6564736f6c63430008130033",
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
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string))
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
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string))
func (_GovernanceWrapper *GovernanceWrapperSession) GetProposal(proposalId uint64) (IGovernanceModuleProposal, error) {
	return _GovernanceWrapper.Contract.GetProposal(&_GovernanceWrapper.CallOpts, proposalId)
}

// GetProposal is a free data retrieval call binding the contract method 0xf1610a28.
//
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string))
func (_GovernanceWrapper *GovernanceWrapperCallerSession) GetProposal(proposalId uint64) (IGovernanceModuleProposal, error) {
	return _GovernanceWrapper.Contract.GetProposal(&_GovernanceWrapper.CallOpts, proposalId)
}

// GetProposals is a free data retrieval call binding the contract method 0xb5828df2.
//
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string)[])
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
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string)[])
func (_GovernanceWrapper *GovernanceWrapperSession) GetProposals(proposalStatus int32) ([]IGovernanceModuleProposal, error) {
	return _GovernanceWrapper.Contract.GetProposals(&_GovernanceWrapper.CallOpts, proposalStatus)
}

// GetProposals is a free data retrieval call binding the contract method 0xb5828df2.
//
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string)[])
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

// CancelProposal is a paid mutator transaction binding the contract method 0x37a9a59e.
//
// Solidity: function cancelProposal(uint64 proposalId) returns(uint64, uint64)
func (_GovernanceWrapper *GovernanceWrapperTransactor) CancelProposal(opts *bind.TransactOpts, proposalId uint64) (*types.Transaction, error) {
	return _GovernanceWrapper.contract.Transact(opts, "cancelProposal", proposalId)
}

// CancelProposal is a paid mutator transaction binding the contract method 0x37a9a59e.
//
// Solidity: function cancelProposal(uint64 proposalId) returns(uint64, uint64)
func (_GovernanceWrapper *GovernanceWrapperSession) CancelProposal(proposalId uint64) (*types.Transaction, error) {
	return _GovernanceWrapper.Contract.CancelProposal(&_GovernanceWrapper.TransactOpts, proposalId)
}

// CancelProposal is a paid mutator transaction binding the contract method 0x37a9a59e.
//
// Solidity: function cancelProposal(uint64 proposalId) returns(uint64, uint64)
func (_GovernanceWrapper *GovernanceWrapperTransactorSession) CancelProposal(proposalId uint64) (*types.Transaction, error) {
	return _GovernanceWrapper.Contract.CancelProposal(&_GovernanceWrapper.TransactOpts, proposalId)
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
