// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testing_governance

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

// CosmosCoin is an auto generated low-level Go binding around an user-defined struct.
type CosmosCoin struct {
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
	TotalDeposit     []CosmosCoin
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governanceModule\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"contractIBankModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"cancelProposal\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.TallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"submitTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositEndTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structCosmos.Coin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"votingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposer\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Proposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"proposalStatus\",\"type\":\"int32\"}],\"name\":\"getProposals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.TallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"submitTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositEndTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structCosmos.Coin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"votingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposer\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Proposal[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceModule\",\"outputs\":[{\"internalType\":\"contractIGovernanceModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proposal\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"submit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"internalType\":\"int32\",\"name\":\"option\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a0604052734381dc2ab14285160c808659aee005d51255add773ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801562000057575f80fd5b5060405162001f1038038062001f1083398181016040528101906200007d91906200018e565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620000e3576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050620001be565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f62000158826200012d565b9050919050565b6200016a816200014c565b811462000175575f80fd5b50565b5f8151905062000188816200015f565b92915050565b5f60208284031215620001a657620001a562000129565b5b5f620001b58482850162000178565b91505092915050565b608051611d39620001d75f395f6103610152611d395ff3fe608060405260043610610073575f3560e01c806376cdb03b1161004d57806376cdb03b14610121578063b5828df21461014b578063f1610a2814610187578063fbab7815146101c35761007a565b806319f7a0fb1461007e5780632b0a7032146100ba57806337a9a59e146100e45761007a565b3661007a57005b5f80fd5b348015610089575f80fd5b506100a4600480360381019061009f9190610950565b6101f3565b6040516100b191906109d6565b60405180910390f35b3480156100c5575f80fd5b506100ce61029a565b6040516100db9190610a69565b60405180910390f35b3480156100ef575f80fd5b5061010a60048036038101906101059190610a82565b6102bd565b604051610118929190610abc565b60405180910390f35b34801561012c575f80fd5b5061013561035f565b6040516101429190610b03565b60405180910390f35b348015610156575f80fd5b50610171600480360381019061016c9190610b1c565b610383565b60405161017e9190610fae565b60405180910390f35b348015610192575f80fd5b506101ad60048036038101906101a89190610a82565b610438565b6040516101ba919061110e565b60405180910390f35b6101dd60048036038101906101d8919061120a565b6104e1565b6040516101ea91906112ce565b60405180910390f35b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166319f7a0fb8585856040518463ffffffff1660e01b81526004016102519392919061133e565b6020604051808303815f875af115801561026d573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061029191906113a4565b90509392505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f805f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166337a9a59e846040518263ffffffff1660e01b815260040161031791906112ce565b60408051808303815f875af1158015610332573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061035691906113e3565b91509150915091565b7f000000000000000000000000000000000000000000000000000000000000000081565b606061038d61066f565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663917c9d9285846040518363ffffffff1660e01b81526004016103e99291906114a3565b5f60405180830381865afa158015610403573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061042b9190611b79565b5090508092505050919050565b6104406106b2565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f1610a28836040518263ffffffff1660e01b815260040161049891906112ce565b5f60405180830381865afa1580156104b2573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906104da9190611bef565b9050919050565b5f80600167ffffffffffffffff8111156104fe576104fd61082c565b5b60405190808252806020026020018201604052801561053757816020015b61052461074f565b81526020019060019003908161051c5790505b50905084848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815f8151811061059157610590611c36565b5b60200260200101516020018190525082815f815181106105b4576105b3611c36565b5b60200260200101515f0181815250505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663474d7f358a8a8a8a6040518563ffffffff1660e01b81526004016106219493929190611c9f565b6020604051808303815f875af115801561063d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106619190611cd8565b915050979650505050505050565b6040518060a00160405280606081526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f151581526020015f151581525090565b604051806101a001604052805f67ffffffffffffffff168152602001606081526020015f60030b81526020016106e6610768565b81526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff168152602001606081526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff168152602001606081526020016060815260200160608152602001606081525090565b60405180604001604052805f8152602001606081525090565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b5f604051905090565b5f80fd5b5f80fd5b5f67ffffffffffffffff82169050919050565b6107bd816107a1565b81146107c7575f80fd5b50565b5f813590506107d8816107b4565b92915050565b5f8160030b9050919050565b6107f3816107de565b81146107fd575f80fd5b50565b5f8135905061080e816107ea565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6108628261081c565b810181811067ffffffffffffffff821117156108815761088061082c565b5b80604052505050565b5f610893610790565b905061089f8282610859565b919050565b5f67ffffffffffffffff8211156108be576108bd61082c565b5b6108c78261081c565b9050602081019050919050565b828183375f83830152505050565b5f6108f46108ef846108a4565b61088a565b9050828152602081018484840111156109105761090f610818565b5b61091b8482856108d4565b509392505050565b5f82601f83011261093757610936610814565b5b81356109478482602086016108e2565b91505092915050565b5f805f6060848603121561096757610966610799565b5b5f610974868287016107ca565b935050602061098586828701610800565b925050604084013567ffffffffffffffff8111156109a6576109a561079d565b5b6109b286828701610923565b9150509250925092565b5f8115159050919050565b6109d0816109bc565b82525050565b5f6020820190506109e95f8301846109c7565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f610a31610a2c610a27846109ef565b610a0e565b6109ef565b9050919050565b5f610a4282610a17565b9050919050565b5f610a5382610a38565b9050919050565b610a6381610a49565b82525050565b5f602082019050610a7c5f830184610a5a565b92915050565b5f60208284031215610a9757610a96610799565b5b5f610aa4848285016107ca565b91505092915050565b610ab6816107a1565b82525050565b5f604082019050610acf5f830185610aad565b610adc6020830184610aad565b9392505050565b5f610aed82610a38565b9050919050565b610afd81610ae3565b82525050565b5f602082019050610b165f830184610af4565b92915050565b5f60208284031215610b3157610b30610799565b5b5f610b3e84828501610800565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610b79816107a1565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610bb6578082015181840152602081019050610b9b565b5f8484015250505050565b5f610bcb82610b7f565b610bd58185610b89565b9350610be5818560208601610b99565b610bee8161081c565b840191505092915050565b610c02816107de565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f610c2c82610c08565b610c368185610c12565b9350610c46818560208601610b99565b610c4f8161081c565b840191505092915050565b5f608083015f8301518482035f860152610c748282610c22565b91505060208301518482036020860152610c8e8282610c22565b91505060408301518482036040860152610ca88282610c22565b91505060608301518482036060860152610cc28282610c22565b9150508091505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f819050919050565b610d0a81610cf8565b82525050565b5f604083015f830151610d255f860182610d01565b5060208301518482036020860152610d3d8282610c22565b9150508091505092915050565b5f610d558383610d10565b905092915050565b5f602082019050919050565b5f610d7382610ccf565b610d7d8185610cd9565b935083602082028501610d8f85610ce9565b805f5b85811015610dca5784840389528151610dab8582610d4a565b9450610db683610d5d565b925060208a01995050600181019050610d92565b50829750879550505050505092915050565b5f6101a083015f830151610df25f860182610b70565b5060208301518482036020860152610e0a8282610bc1565b9150506040830151610e1f6040860182610bf9565b5060608301518482036060860152610e378282610c5a565b9150506080830151610e4c6080860182610b70565b5060a0830151610e5f60a0860182610b70565b5060c083015184820360c0860152610e778282610d69565b91505060e0830151610e8c60e0860182610b70565b50610100830151610ea1610100860182610b70565b50610120830151848203610120860152610ebb8282610c22565b915050610140830151848203610140860152610ed78282610c22565b915050610160830151848203610160860152610ef38282610c22565b915050610180830151848203610180860152610f0f8282610c22565b9150508091505092915050565b5f610f278383610ddc565b905092915050565b5f602082019050919050565b5f610f4582610b47565b610f4f8185610b51565b935083602082028501610f6185610b61565b805f5b85811015610f9c5784840389528151610f7d8582610f1c565b9450610f8883610f2f565b925060208a01995050600181019050610f64565b50829750879550505050505092915050565b5f6020820190508181035f830152610fc68184610f3b565b905092915050565b5f6101a083015f830151610fe45f860182610b70565b5060208301518482036020860152610ffc8282610bc1565b91505060408301516110116040860182610bf9565b50606083015184820360608601526110298282610c5a565b915050608083015161103e6080860182610b70565b5060a083015161105160a0860182610b70565b5060c083015184820360c08601526110698282610d69565b91505060e083015161107e60e0860182610b70565b50610100830151611093610100860182610b70565b506101208301518482036101208601526110ad8282610c22565b9150506101408301518482036101408601526110c98282610c22565b9150506101608301518482036101608601526110e58282610c22565b9150506101808301518482036101808601526111018282610c22565b9150508091505092915050565b5f6020820190508181035f8301526111268184610fce565b905092915050565b5f80fd5b5f80fd5b5f8083601f84011261114b5761114a610814565b5b8235905067ffffffffffffffff8111156111685761116761112e565b5b60208301915083600182028301111561118457611183611132565b5b9250929050565b5f8083601f8401126111a05761119f610814565b5b8235905067ffffffffffffffff8111156111bd576111bc61112e565b5b6020830191508360018202830111156111d9576111d8611132565b5b9250929050565b6111e981610cf8565b81146111f3575f80fd5b50565b5f81359050611204816111e0565b92915050565b5f805f805f805f6080888a03121561122557611224610799565b5b5f88013567ffffffffffffffff8111156112425761124161079d565b5b61124e8a828b01611136565b9750975050602088013567ffffffffffffffff8111156112715761127061079d565b5b61127d8a828b01611136565b9550955050604088013567ffffffffffffffff8111156112a05761129f61079d565b5b6112ac8a828b0161118b565b935093505060606112bf8a828b016111f6565b91505092959891949750929550565b5f6020820190506112e15f830184610aad565b92915050565b6112f0816107de565b82525050565b5f82825260208201905092915050565b5f61131082610c08565b61131a81856112f6565b935061132a818560208601610b99565b6113338161081c565b840191505092915050565b5f6060820190506113515f830186610aad565b61135e60208301856112e7565b81810360408301526113708184611306565b9050949350505050565b611383816109bc565b811461138d575f80fd5b50565b5f8151905061139e8161137a565b92915050565b5f602082840312156113b9576113b8610799565b5b5f6113c684828501611390565b91505092915050565b5f815190506113dd816107b4565b92915050565b5f80604083850312156113f9576113f8610799565b5b5f611406858286016113cf565b9250506020611417858286016113cf565b9150509250929050565b61142a816109bc565b82525050565b5f60a083015f8301518482035f86015261144a8282610c22565b915050602083015161145f6020860182610b70565b5060408301516114726040860182610b70565b5060608301516114856060860182611421565b5060808301516114986080860182611421565b508091505092915050565b5f6040820190506114b65f8301856112e7565b81810360208301526114c88184611430565b90509392505050565b5f67ffffffffffffffff8211156114eb576114ea61082c565b5b602082029050602081019050919050565b5f80fd5b5f80fd5b5f67ffffffffffffffff82111561151e5761151d61082c565b5b6115278261081c565b9050602081019050919050565b5f61154661154184611504565b61088a565b90508281526020810184848401111561156257611561610818565b5b61156d848285610b99565b509392505050565b5f82601f83011261158957611588610814565b5b8151611599848260208601611534565b91505092915050565b5f815190506115b0816107ea565b92915050565b5f6115c86115c3846108a4565b61088a565b9050828152602081018484840111156115e4576115e3610818565b5b6115ef848285610b99565b509392505050565b5f82601f83011261160b5761160a610814565b5b815161161b8482602086016115b6565b91505092915050565b5f60808284031215611639576116386114fc565b5b611643608061088a565b90505f82015167ffffffffffffffff81111561166257611661611500565b5b61166e848285016115f7565b5f83015250602082015167ffffffffffffffff81111561169157611690611500565b5b61169d848285016115f7565b602083015250604082015167ffffffffffffffff8111156116c1576116c0611500565b5b6116cd848285016115f7565b604083015250606082015167ffffffffffffffff8111156116f1576116f0611500565b5b6116fd848285016115f7565b60608301525092915050565b5f67ffffffffffffffff8211156117235761172261082c565b5b602082029050602081019050919050565b5f81519050611742816111e0565b92915050565b5f6040828403121561175d5761175c6114fc565b5b611767604061088a565b90505f61177684828501611734565b5f83015250602082015167ffffffffffffffff81111561179957611798611500565b5b6117a5848285016115f7565b60208301525092915050565b5f6117c36117be84611709565b61088a565b905080838252602082019050602084028301858111156117e6576117e5611132565b5b835b8181101561182d57805167ffffffffffffffff81111561180b5761180a610814565b5b8086016118188982611748565b855260208501945050506020810190506117e8565b5050509392505050565b5f82601f83011261184b5761184a610814565b5b815161185b8482602086016117b1565b91505092915050565b5f6101a0828403121561187a576118796114fc565b5b6118856101a061088a565b90505f611894848285016113cf565b5f83015250602082015167ffffffffffffffff8111156118b7576118b6611500565b5b6118c384828501611575565b60208301525060406118d7848285016115a2565b604083015250606082015167ffffffffffffffff8111156118fb576118fa611500565b5b61190784828501611624565b606083015250608061191b848285016113cf565b60808301525060a061192f848285016113cf565b60a08301525060c082015167ffffffffffffffff81111561195357611952611500565b5b61195f84828501611837565b60c08301525060e0611973848285016113cf565b60e083015250610100611988848285016113cf565b6101008301525061012082015167ffffffffffffffff8111156119ae576119ad611500565b5b6119ba848285016115f7565b6101208301525061014082015167ffffffffffffffff8111156119e0576119df611500565b5b6119ec848285016115f7565b6101408301525061016082015167ffffffffffffffff811115611a1257611a11611500565b5b611a1e848285016115f7565b6101608301525061018082015167ffffffffffffffff811115611a4457611a43611500565b5b611a50848285016115f7565b6101808301525092915050565b5f611a6f611a6a846114d1565b61088a565b90508083825260208201905060208402830185811115611a9257611a91611132565b5b835b81811015611ad957805167ffffffffffffffff811115611ab757611ab6610814565b5b808601611ac48982611864565b85526020850194505050602081019050611a94565b5050509392505050565b5f82601f830112611af757611af6610814565b5b8151611b07848260208601611a5d565b91505092915050565b5f60408284031215611b2557611b246114fc565b5b611b2f604061088a565b90505f82015167ffffffffffffffff811115611b4e57611b4d611500565b5b611b5a848285016115f7565b5f830152506020611b6d848285016113cf565b60208301525092915050565b5f8060408385031215611b8f57611b8e610799565b5b5f83015167ffffffffffffffff811115611bac57611bab61079d565b5b611bb885828601611ae3565b925050602083015167ffffffffffffffff811115611bd957611bd861079d565b5b611be585828601611b10565b9150509250929050565b5f60208284031215611c0457611c03610799565b5b5f82015167ffffffffffffffff811115611c2157611c2061079d565b5b611c2d84828501611864565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f82825260208201905092915050565b5f611c7e8385611c63565b9350611c8b8385846108d4565b611c948361081c565b840190509392505050565b5f6040820190508181035f830152611cb8818688611c73565b90508181036020830152611ccd818486611c73565b905095945050505050565b5f60208284031215611ced57611cec610799565b5b5f611cfa848285016113cf565b9150509291505056fea264697066735822122041ddfee5fe0e1208450e47ffbf9a5a4d292cbf1eb09d21d73ffa61cbd7019bfa64736f6c63430008140033",
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
