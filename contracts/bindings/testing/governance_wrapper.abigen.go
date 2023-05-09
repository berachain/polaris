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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governanceModule\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"contractIBankModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"cancelProposal\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.TallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"submitTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositEndTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Coin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"votingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposer\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Proposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"proposalStatus\",\"type\":\"int32\"}],\"name\":\"getProposals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.TallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"submitTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositEndTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Coin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"votingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposer\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Proposal[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceModule\",\"outputs\":[{\"internalType\":\"contractIGovernanceModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proposal\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"submit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"internalType\":\"int32\",\"name\":\"option\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a0604052734381dc2ab14285160c808659aee005d51255add773ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff168152503480156200005857600080fd5b506040516200205a3803806200205a83398181016040528101906200007e919062000196565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620000e5576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050620001c8565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200015e8262000131565b9050919050565b620001708162000151565b81146200017c57600080fd5b50565b600081519050620001908162000165565b92915050565b600060208284031215620001af57620001ae6200012c565b5b6000620001bf848285016200017f565b91505092915050565b608051611e6f620001eb6000396000818161037401526105d50152611e6f6000f3fe6080604052600436106100745760003560e01c806376cdb03b1161004e57806376cdb03b14610126578063b5828df214610151578063f1610a281461018e578063fbab7815146101cb5761007b565b806319f7a0fb146100805780632b0a7032146100bd57806337a9a59e146100e85761007b565b3661007b57005b600080fd5b34801561008c57600080fd5b506100a760048036038101906100a291906109dc565b6101fb565b6040516100b49190610a66565b60405180910390f35b3480156100c957600080fd5b506100d26102a7565b6040516100df9190610b00565b60405180910390f35b3480156100f457600080fd5b5061010f600480360381019061010a9190610b1b565b6102cb565b60405161011d929190610b57565b60405180910390f35b34801561013257600080fd5b5061013b610372565b6040516101489190610ba1565b60405180910390f35b34801561015d57600080fd5b5061017860048036038101906101739190610bbc565b610396565b6040516101859190611070565b60405180910390f35b34801561019a57600080fd5b506101b560048036038101906101b09190610b1b565b61043e565b6040516101c291906111d5565b60405180910390f35b6101e560048036038101906101e091906112d9565b6104ec565b6040516101f291906113a2565b60405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166319f7a0fb8585856040518463ffffffff1660e01b815260040161025b93929190611416565b6020604051808303816000875af115801561027a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061029e9190611480565b90509392505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166337a9a59e846040518263ffffffff1660e01b815260040161032791906113a2565b60408051808303816000875af1158015610345573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061036991906114c2565b91509150915091565b7f000000000000000000000000000000000000000000000000000000000000000081565b606060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b5828df2836040518263ffffffff1660e01b81526004016103f19190611502565b600060405180830381865afa15801561040e573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906104379190611b74565b9050919050565b610446610724565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f1610a28836040518263ffffffff1660e01b815260040161049f91906113a2565b600060405180830381865afa1580156104bc573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906104e59190611bbd565b9050919050565b600080600167ffffffffffffffff81111561050a576105096108b1565b5b60405190808252806020026020018201604052801561054357816020015b6105306107c7565b8152602001906001900390816105285790505b50905084848080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508160008151811061059f5761059e611c06565b5b60200260200101516020018190525082816000815181106105c3576105c2611c06565b5b602002602001015160000181815250507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663844048113330846040518463ffffffff1660e01b815260040161063093929190611d55565b6020604051808303816000875af115801561064f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106739190611480565b5060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663474d7f358a8a8a8a6040518563ffffffff1660e01b81526004016106d39493929190611dd1565b6020604051808303816000875af11580156106f2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107169190611e0c565b915050979650505050505050565b604051806101a00160405280600067ffffffffffffffff16815260200160608152602001600060030b815260200161075a6107e1565b8152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001606081526020016060815260200160608152602001606081525090565b604051806040016040528060008152602001606081525090565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000604051905090565b600080fd5b600080fd5b600067ffffffffffffffff82169050919050565b61083a8161081d565b811461084557600080fd5b50565b60008135905061085781610831565b92915050565b60008160030b9050919050565b6108738161085d565b811461087e57600080fd5b50565b6000813590506108908161086a565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6108e9826108a0565b810181811067ffffffffffffffff82111715610908576109076108b1565b5b80604052505050565b600061091b610809565b905061092782826108e0565b919050565b600067ffffffffffffffff821115610947576109466108b1565b5b610950826108a0565b9050602081019050919050565b82818337600083830152505050565b600061097f61097a8461092c565b610911565b90508281526020810184848401111561099b5761099a61089b565b5b6109a684828561095d565b509392505050565b600082601f8301126109c3576109c2610896565b5b81356109d384826020860161096c565b91505092915050565b6000806000606084860312156109f5576109f4610813565b5b6000610a0386828701610848565b9350506020610a1486828701610881565b925050604084013567ffffffffffffffff811115610a3557610a34610818565b5b610a41868287016109ae565b9150509250925092565b60008115159050919050565b610a6081610a4b565b82525050565b6000602082019050610a7b6000830184610a57565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610ac6610ac1610abc84610a81565b610aa1565b610a81565b9050919050565b6000610ad882610aab565b9050919050565b6000610aea82610acd565b9050919050565b610afa81610adf565b82525050565b6000602082019050610b156000830184610af1565b92915050565b600060208284031215610b3157610b30610813565b5b6000610b3f84828501610848565b91505092915050565b610b518161081d565b82525050565b6000604082019050610b6c6000830185610b48565b610b796020830184610b48565b9392505050565b6000610b8b82610acd565b9050919050565b610b9b81610b80565b82525050565b6000602082019050610bb66000830184610b92565b92915050565b600060208284031215610bd257610bd1610813565b5b6000610be084828501610881565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610c1e8161081d565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610c5e578082015181840152602081019050610c43565b60008484015250505050565b6000610c7582610c24565b610c7f8185610c2f565b9350610c8f818560208601610c40565b610c98816108a0565b840191505092915050565b610cac8161085d565b82525050565b600081519050919050565b600082825260208201905092915050565b6000610cd982610cb2565b610ce38185610cbd565b9350610cf3818560208601610c40565b610cfc816108a0565b840191505092915050565b60006080830160008301518482036000860152610d248282610cce565b91505060208301518482036020860152610d3e8282610cce565b91505060408301518482036040860152610d588282610cce565b91505060608301518482036060860152610d728282610cce565b9150508091505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000819050919050565b610dbe81610dab565b82525050565b6000604083016000830151610ddc6000860182610db5565b5060208301518482036020860152610df48282610cce565b9150508091505092915050565b6000610e0d8383610dc4565b905092915050565b6000602082019050919050565b6000610e2d82610d7f565b610e378185610d8a565b935083602082028501610e4985610d9b565b8060005b85811015610e855784840389528151610e668582610e01565b9450610e7183610e15565b925060208a01995050600181019050610e4d565b50829750879550505050505092915050565b60006101a083016000830151610eb06000860182610c15565b5060208301518482036020860152610ec88282610c6a565b9150506040830151610edd6040860182610ca3565b5060608301518482036060860152610ef58282610d07565b9150506080830151610f0a6080860182610c15565b5060a0830151610f1d60a0860182610c15565b5060c083015184820360c0860152610f358282610e22565b91505060e0830151610f4a60e0860182610c15565b50610100830151610f5f610100860182610c15565b50610120830151848203610120860152610f798282610cce565b915050610140830151848203610140860152610f958282610cce565b915050610160830151848203610160860152610fb18282610cce565b915050610180830151848203610180860152610fcd8282610cce565b9150508091505092915050565b6000610fe68383610e97565b905092915050565b6000602082019050919050565b600061100682610be9565b6110108185610bf4565b93508360208202850161102285610c05565b8060005b8581101561105e578484038952815161103f8582610fda565b945061104a83610fee565b925060208a01995050600181019050611026565b50829750879550505050505092915050565b6000602082019050818103600083015261108a8184610ffb565b905092915050565b60006101a0830160008301516110ab6000860182610c15565b50602083015184820360208601526110c38282610c6a565b91505060408301516110d86040860182610ca3565b50606083015184820360608601526110f08282610d07565b91505060808301516111056080860182610c15565b5060a083015161111860a0860182610c15565b5060c083015184820360c08601526111308282610e22565b91505060e083015161114560e0860182610c15565b5061010083015161115a610100860182610c15565b506101208301518482036101208601526111748282610cce565b9150506101408301518482036101408601526111908282610cce565b9150506101608301518482036101608601526111ac8282610cce565b9150506101808301518482036101808601526111c88282610cce565b9150508091505092915050565b600060208201905081810360008301526111ef8184611092565b905092915050565b600080fd5b600080fd5b60008083601f84011261121757611216610896565b5b8235905067ffffffffffffffff811115611234576112336111f7565b5b6020830191508360018202830111156112505761124f6111fc565b5b9250929050565b60008083601f84011261126d5761126c610896565b5b8235905067ffffffffffffffff81111561128a576112896111f7565b5b6020830191508360018202830111156112a6576112a56111fc565b5b9250929050565b6112b681610dab565b81146112c157600080fd5b50565b6000813590506112d3816112ad565b92915050565b60008060008060008060006080888a0312156112f8576112f7610813565b5b600088013567ffffffffffffffff81111561131657611315610818565b5b6113228a828b01611201565b9750975050602088013567ffffffffffffffff81111561134557611344610818565b5b6113518a828b01611201565b9550955050604088013567ffffffffffffffff81111561137457611373610818565b5b6113808a828b01611257565b935093505060606113938a828b016112c4565b91505092959891949750929550565b60006020820190506113b76000830184610b48565b92915050565b6113c68161085d565b82525050565b600082825260208201905092915050565b60006113e882610cb2565b6113f281856113cc565b9350611402818560208601610c40565b61140b816108a0565b840191505092915050565b600060608201905061142b6000830186610b48565b61143860208301856113bd565b818103604083015261144a81846113dd565b9050949350505050565b61145d81610a4b565b811461146857600080fd5b50565b60008151905061147a81611454565b92915050565b60006020828403121561149657611495610813565b5b60006114a48482850161146b565b91505092915050565b6000815190506114bc81610831565b92915050565b600080604083850312156114d9576114d8610813565b5b60006114e7858286016114ad565b92505060206114f8858286016114ad565b9150509250929050565b600060208201905061151760008301846113bd565b92915050565b600067ffffffffffffffff821115611538576115376108b1565b5b602082029050602081019050919050565b600080fd5b600080fd5b600067ffffffffffffffff82111561156e5761156d6108b1565b5b611577826108a0565b9050602081019050919050565b600061159761159284611553565b610911565b9050828152602081018484840111156115b3576115b261089b565b5b6115be848285610c40565b509392505050565b600082601f8301126115db576115da610896565b5b81516115eb848260208601611584565b91505092915050565b6000815190506116038161086a565b92915050565b600061161c6116178461092c565b610911565b9050828152602081018484840111156116385761163761089b565b5b611643848285610c40565b509392505050565b600082601f8301126116605761165f610896565b5b8151611670848260208601611609565b91505092915050565b60006080828403121561168f5761168e611549565b5b6116996080610911565b9050600082015167ffffffffffffffff8111156116b9576116b861154e565b5b6116c58482850161164b565b600083015250602082015167ffffffffffffffff8111156116e9576116e861154e565b5b6116f58482850161164b565b602083015250604082015167ffffffffffffffff8111156117195761171861154e565b5b6117258482850161164b565b604083015250606082015167ffffffffffffffff8111156117495761174861154e565b5b6117558482850161164b565b60608301525092915050565b600067ffffffffffffffff82111561177c5761177b6108b1565b5b602082029050602081019050919050565b60008151905061179c816112ad565b92915050565b6000604082840312156117b8576117b7611549565b5b6117c26040610911565b905060006117d28482850161178d565b600083015250602082015167ffffffffffffffff8111156117f6576117f561154e565b5b6118028482850161164b565b60208301525092915050565b600061182161181c84611761565b610911565b90508083825260208201905060208402830185811115611844576118436111fc565b5b835b8181101561188b57805167ffffffffffffffff81111561186957611868610896565b5b80860161187689826117a2565b85526020850194505050602081019050611846565b5050509392505050565b600082601f8301126118aa576118a9610896565b5b81516118ba84826020860161180e565b91505092915050565b60006101a082840312156118da576118d9611549565b5b6118e56101a0610911565b905060006118f5848285016114ad565b600083015250602082015167ffffffffffffffff8111156119195761191861154e565b5b611925848285016115c6565b6020830152506040611939848285016115f4565b604083015250606082015167ffffffffffffffff81111561195d5761195c61154e565b5b61196984828501611679565b606083015250608061197d848285016114ad565b60808301525060a0611991848285016114ad565b60a08301525060c082015167ffffffffffffffff8111156119b5576119b461154e565b5b6119c184828501611895565b60c08301525060e06119d5848285016114ad565b60e0830152506101006119ea848285016114ad565b6101008301525061012082015167ffffffffffffffff811115611a1057611a0f61154e565b5b611a1c8482850161164b565b6101208301525061014082015167ffffffffffffffff811115611a4257611a4161154e565b5b611a4e8482850161164b565b6101408301525061016082015167ffffffffffffffff811115611a7457611a7361154e565b5b611a808482850161164b565b6101608301525061018082015167ffffffffffffffff811115611aa657611aa561154e565b5b611ab28482850161164b565b6101808301525092915050565b6000611ad2611acd8461151d565b610911565b90508083825260208201905060208402830185811115611af557611af46111fc565b5b835b81811015611b3c57805167ffffffffffffffff811115611b1a57611b19610896565b5b808601611b2789826118c3565b85526020850194505050602081019050611af7565b5050509392505050565b600082601f830112611b5b57611b5a610896565b5b8151611b6b848260208601611abf565b91505092915050565b600060208284031215611b8a57611b89610813565b5b600082015167ffffffffffffffff811115611ba857611ba7610818565b5b611bb484828501611b46565b91505092915050565b600060208284031215611bd357611bd2610813565b5b600082015167ffffffffffffffff811115611bf157611bf0610818565b5b611bfd848285016118c3565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000611c4082610a81565b9050919050565b611c5081611c35565b82525050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000604083016000830151611c9a6000860182610db5565b5060208301518482036020860152611cb28282610cce565b9150508091505092915050565b6000611ccb8383611c82565b905092915050565b6000602082019050919050565b6000611ceb82611c56565b611cf58185611c61565b935083602082028501611d0785611c72565b8060005b85811015611d435784840389528151611d248582611cbf565b9450611d2f83611cd3565b925060208a01995050600181019050611d0b565b50829750879550505050505092915050565b6000606082019050611d6a6000830186611c47565b611d776020830185611c47565b8181036040830152611d898184611ce0565b9050949350505050565b600082825260208201905092915050565b6000611db08385611d93565b9350611dbd83858461095d565b611dc6836108a0565b840190509392505050565b60006040820190508181036000830152611dec818688611da4565b90508181036020830152611e01818486611da4565b905095945050505050565b600060208284031215611e2257611e21610813565b5b6000611e30848285016114ad565b9150509291505056fea26469706673582212201e8b159211b44805438dc8fd96e77e7dd3c55388584fa50dea1a0f6d18e25c7264736f6c63430008130033",
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

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_GovernanceWrapper *GovernanceWrapperCaller) Bank(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GovernanceWrapper.contract.Call(opts, &out, "bank")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_GovernanceWrapper *GovernanceWrapperSession) Bank() (common.Address, error) {
	return _GovernanceWrapper.Contract.Bank(&_GovernanceWrapper.CallOpts)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_GovernanceWrapper *GovernanceWrapperCallerSession) Bank() (common.Address, error) {
	return _GovernanceWrapper.Contract.Bank(&_GovernanceWrapper.CallOpts)
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

// Submit is a paid mutator transaction binding the contract method 0xfbab7815.
//
// Solidity: function submit(bytes proposal, bytes message, string denom, uint256 amount) payable returns(uint64)
func (_GovernanceWrapper *GovernanceWrapperTransactor) Submit(opts *bind.TransactOpts, proposal []byte, message []byte, denom string, amount *big.Int) (*types.Transaction, error) {
	return _GovernanceWrapper.contract.Transact(opts, "submit", proposal, message, denom, amount)
}

// Submit is a paid mutator transaction binding the contract method 0xfbab7815.
//
// Solidity: function submit(bytes proposal, bytes message, string denom, uint256 amount) payable returns(uint64)
func (_GovernanceWrapper *GovernanceWrapperSession) Submit(proposal []byte, message []byte, denom string, amount *big.Int) (*types.Transaction, error) {
	return _GovernanceWrapper.Contract.Submit(&_GovernanceWrapper.TransactOpts, proposal, message, denom, amount)
}

// Submit is a paid mutator transaction binding the contract method 0xfbab7815.
//
// Solidity: function submit(bytes proposal, bytes message, string denom, uint256 amount) payable returns(uint64)
func (_GovernanceWrapper *GovernanceWrapperTransactorSession) Submit(proposal []byte, message []byte, denom string, amount *big.Int) (*types.Transaction, error) {
	return _GovernanceWrapper.Contract.Submit(&_GovernanceWrapper.TransactOpts, proposal, message, denom, amount)
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
