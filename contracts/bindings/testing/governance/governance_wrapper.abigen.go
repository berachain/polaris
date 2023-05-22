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
	Bin: "0x60a0604052734381dc2ab14285160c808659aee005d51255add773ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801562000057575f80fd5b5060405162001f1038038062001f1083398181016040528101906200007d91906200018e565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620000e3576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050620001be565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f62000158826200012d565b9050919050565b6200016a816200014c565b811462000175575f80fd5b50565b5f8151905062000188816200015f565b92915050565b5f60208284031215620001a657620001a562000129565b5b5f620001b58482850162000178565b91505092915050565b608051611d32620001de5f395f818161036101526105b30152611d325ff3fe608060405260043610610073575f3560e01c806376cdb03b1161004d57806376cdb03b14610121578063b5828df21461014b578063f1610a2814610187578063fbab7815146101c35761007a565b806319f7a0fb1461007e5780632b0a7032146100ba57806337a9a59e146100e45761007a565b3661007a57005b5f80fd5b348015610089575f80fd5b506100a4600480360381019061009f9190610999565b6101f3565b6040516100b19190610a1f565b60405180910390f35b3480156100c5575f80fd5b506100ce61029a565b6040516100db9190610ab2565b60405180910390f35b3480156100ef575f80fd5b5061010a60048036038101906101059190610acb565b6102bd565b604051610118929190610b05565b60405180910390f35b34801561012c575f80fd5b5061013561035f565b6040516101429190610b4c565b60405180910390f35b348015610156575f80fd5b50610171600480360381019061016c9190610b65565b610383565b60405161017e9190610ff7565b60405180910390f35b348015610192575f80fd5b506101ad60048036038101906101a89190610acb565b610426565b6040516101ba9190611157565b60405180910390f35b6101dd60048036038101906101d89190611253565b6104cf565b6040516101ea9190611317565b60405180910390f35b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166319f7a0fb8585856040518463ffffffff1660e01b815260040161025193929190611387565b6020604051808303815f875af115801561026d573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061029191906113ed565b90509392505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f805f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166337a9a59e846040518263ffffffff1660e01b81526004016103179190611317565b60408051808303815f875af1158015610332573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610356919061142c565b91509150915091565b7f000000000000000000000000000000000000000000000000000000000000000081565b60605f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b5828df2836040518263ffffffff1660e01b81526004016103dd919061146a565b5f60405180830381865afa1580156103f7573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061041f9190611ac2565b9050919050565b61042e6106fb565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f1610a28836040518263ffffffff1660e01b81526004016104869190611317565b5f60405180830381865afa1580156104a0573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906104c89190611b09565b9050919050565b5f80600167ffffffffffffffff8111156104ec576104eb610875565b5b60405190808252806020026020018201604052801561052557816020015b610512610798565b81526020019060019003908161050a5790505b50905084848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815f8151811061057f5761057e611b50565b5b60200260200101516020018190525082815f815181106105a2576105a1611b50565b5b60200260200101515f0181815250507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663844048113330846040518463ffffffff1660e01b815260040161060e93929190611c20565b6020604051808303815f875af115801561062a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061064e91906113ed565b505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663474d7f358a8a8a8a6040518563ffffffff1660e01b81526004016106ad9493929190611c98565b6020604051808303815f875af11580156106c9573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106ed9190611cd1565b915050979650505050505050565b604051806101a001604052805f67ffffffffffffffff168152602001606081526020015f60030b815260200161072f6107b1565b81526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff168152602001606081526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff168152602001606081526020016060815260200160608152602001606081525090565b60405180604001604052805f8152602001606081525090565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b5f604051905090565b5f80fd5b5f80fd5b5f67ffffffffffffffff82169050919050565b610806816107ea565b8114610810575f80fd5b50565b5f81359050610821816107fd565b92915050565b5f8160030b9050919050565b61083c81610827565b8114610846575f80fd5b50565b5f8135905061085781610833565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6108ab82610865565b810181811067ffffffffffffffff821117156108ca576108c9610875565b5b80604052505050565b5f6108dc6107d9565b90506108e882826108a2565b919050565b5f67ffffffffffffffff82111561090757610906610875565b5b61091082610865565b9050602081019050919050565b828183375f83830152505050565b5f61093d610938846108ed565b6108d3565b90508281526020810184848401111561095957610958610861565b5b61096484828561091d565b509392505050565b5f82601f8301126109805761097f61085d565b5b813561099084826020860161092b565b91505092915050565b5f805f606084860312156109b0576109af6107e2565b5b5f6109bd86828701610813565b93505060206109ce86828701610849565b925050604084013567ffffffffffffffff8111156109ef576109ee6107e6565b5b6109fb8682870161096c565b9150509250925092565b5f8115159050919050565b610a1981610a05565b82525050565b5f602082019050610a325f830184610a10565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f610a7a610a75610a7084610a38565b610a57565b610a38565b9050919050565b5f610a8b82610a60565b9050919050565b5f610a9c82610a81565b9050919050565b610aac81610a92565b82525050565b5f602082019050610ac55f830184610aa3565b92915050565b5f60208284031215610ae057610adf6107e2565b5b5f610aed84828501610813565b91505092915050565b610aff816107ea565b82525050565b5f604082019050610b185f830185610af6565b610b256020830184610af6565b9392505050565b5f610b3682610a81565b9050919050565b610b4681610b2c565b82525050565b5f602082019050610b5f5f830184610b3d565b92915050565b5f60208284031215610b7a57610b796107e2565b5b5f610b8784828501610849565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610bc2816107ea565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610bff578082015181840152602081019050610be4565b5f8484015250505050565b5f610c1482610bc8565b610c1e8185610bd2565b9350610c2e818560208601610be2565b610c3781610865565b840191505092915050565b610c4b81610827565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f610c7582610c51565b610c7f8185610c5b565b9350610c8f818560208601610be2565b610c9881610865565b840191505092915050565b5f608083015f8301518482035f860152610cbd8282610c6b565b91505060208301518482036020860152610cd78282610c6b565b91505060408301518482036040860152610cf18282610c6b565b91505060608301518482036060860152610d0b8282610c6b565b9150508091505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f819050919050565b610d5381610d41565b82525050565b5f604083015f830151610d6e5f860182610d4a565b5060208301518482036020860152610d868282610c6b565b9150508091505092915050565b5f610d9e8383610d59565b905092915050565b5f602082019050919050565b5f610dbc82610d18565b610dc68185610d22565b935083602082028501610dd885610d32565b805f5b85811015610e135784840389528151610df48582610d93565b9450610dff83610da6565b925060208a01995050600181019050610ddb565b50829750879550505050505092915050565b5f6101a083015f830151610e3b5f860182610bb9565b5060208301518482036020860152610e538282610c0a565b9150506040830151610e686040860182610c42565b5060608301518482036060860152610e808282610ca3565b9150506080830151610e956080860182610bb9565b5060a0830151610ea860a0860182610bb9565b5060c083015184820360c0860152610ec08282610db2565b91505060e0830151610ed560e0860182610bb9565b50610100830151610eea610100860182610bb9565b50610120830151848203610120860152610f048282610c6b565b915050610140830151848203610140860152610f208282610c6b565b915050610160830151848203610160860152610f3c8282610c6b565b915050610180830151848203610180860152610f588282610c6b565b9150508091505092915050565b5f610f708383610e25565b905092915050565b5f602082019050919050565b5f610f8e82610b90565b610f988185610b9a565b935083602082028501610faa85610baa565b805f5b85811015610fe55784840389528151610fc68582610f65565b9450610fd183610f78565b925060208a01995050600181019050610fad565b50829750879550505050505092915050565b5f6020820190508181035f83015261100f8184610f84565b905092915050565b5f6101a083015f83015161102d5f860182610bb9565b50602083015184820360208601526110458282610c0a565b915050604083015161105a6040860182610c42565b50606083015184820360608601526110728282610ca3565b91505060808301516110876080860182610bb9565b5060a083015161109a60a0860182610bb9565b5060c083015184820360c08601526110b28282610db2565b91505060e08301516110c760e0860182610bb9565b506101008301516110dc610100860182610bb9565b506101208301518482036101208601526110f68282610c6b565b9150506101408301518482036101408601526111128282610c6b565b91505061016083015184820361016086015261112e8282610c6b565b91505061018083015184820361018086015261114a8282610c6b565b9150508091505092915050565b5f6020820190508181035f83015261116f8184611017565b905092915050565b5f80fd5b5f80fd5b5f8083601f8401126111945761119361085d565b5b8235905067ffffffffffffffff8111156111b1576111b0611177565b5b6020830191508360018202830111156111cd576111cc61117b565b5b9250929050565b5f8083601f8401126111e9576111e861085d565b5b8235905067ffffffffffffffff81111561120657611205611177565b5b6020830191508360018202830111156112225761122161117b565b5b9250929050565b61123281610d41565b811461123c575f80fd5b50565b5f8135905061124d81611229565b92915050565b5f805f805f805f6080888a03121561126e5761126d6107e2565b5b5f88013567ffffffffffffffff81111561128b5761128a6107e6565b5b6112978a828b0161117f565b9750975050602088013567ffffffffffffffff8111156112ba576112b96107e6565b5b6112c68a828b0161117f565b9550955050604088013567ffffffffffffffff8111156112e9576112e86107e6565b5b6112f58a828b016111d4565b935093505060606113088a828b0161123f565b91505092959891949750929550565b5f60208201905061132a5f830184610af6565b92915050565b61133981610827565b82525050565b5f82825260208201905092915050565b5f61135982610c51565b611363818561133f565b9350611373818560208601610be2565b61137c81610865565b840191505092915050565b5f60608201905061139a5f830186610af6565b6113a76020830185611330565b81810360408301526113b9818461134f565b9050949350505050565b6113cc81610a05565b81146113d6575f80fd5b50565b5f815190506113e7816113c3565b92915050565b5f60208284031215611402576114016107e2565b5b5f61140f848285016113d9565b91505092915050565b5f81519050611426816107fd565b92915050565b5f8060408385031215611442576114416107e2565b5b5f61144f85828601611418565b925050602061146085828601611418565b9150509250929050565b5f60208201905061147d5f830184611330565b92915050565b5f67ffffffffffffffff82111561149d5761149c610875565b5b602082029050602081019050919050565b5f80fd5b5f80fd5b5f67ffffffffffffffff8211156114d0576114cf610875565b5b6114d982610865565b9050602081019050919050565b5f6114f86114f3846114b6565b6108d3565b90508281526020810184848401111561151457611513610861565b5b61151f848285610be2565b509392505050565b5f82601f83011261153b5761153a61085d565b5b815161154b8482602086016114e6565b91505092915050565b5f8151905061156281610833565b92915050565b5f61157a611575846108ed565b6108d3565b90508281526020810184848401111561159657611595610861565b5b6115a1848285610be2565b509392505050565b5f82601f8301126115bd576115bc61085d565b5b81516115cd848260208601611568565b91505092915050565b5f608082840312156115eb576115ea6114ae565b5b6115f560806108d3565b90505f82015167ffffffffffffffff811115611614576116136114b2565b5b611620848285016115a9565b5f83015250602082015167ffffffffffffffff811115611643576116426114b2565b5b61164f848285016115a9565b602083015250604082015167ffffffffffffffff811115611673576116726114b2565b5b61167f848285016115a9565b604083015250606082015167ffffffffffffffff8111156116a3576116a26114b2565b5b6116af848285016115a9565b60608301525092915050565b5f67ffffffffffffffff8211156116d5576116d4610875565b5b602082029050602081019050919050565b5f815190506116f481611229565b92915050565b5f6040828403121561170f5761170e6114ae565b5b61171960406108d3565b90505f611728848285016116e6565b5f83015250602082015167ffffffffffffffff81111561174b5761174a6114b2565b5b611757848285016115a9565b60208301525092915050565b5f611775611770846116bb565b6108d3565b905080838252602082019050602084028301858111156117985761179761117b565b5b835b818110156117df57805167ffffffffffffffff8111156117bd576117bc61085d565b5b8086016117ca89826116fa565b8552602085019450505060208101905061179a565b5050509392505050565b5f82601f8301126117fd576117fc61085d565b5b815161180d848260208601611763565b91505092915050565b5f6101a0828403121561182c5761182b6114ae565b5b6118376101a06108d3565b90505f61184684828501611418565b5f83015250602082015167ffffffffffffffff811115611869576118686114b2565b5b61187584828501611527565b602083015250604061188984828501611554565b604083015250606082015167ffffffffffffffff8111156118ad576118ac6114b2565b5b6118b9848285016115d6565b60608301525060806118cd84828501611418565b60808301525060a06118e184828501611418565b60a08301525060c082015167ffffffffffffffff811115611905576119046114b2565b5b611911848285016117e9565b60c08301525060e061192584828501611418565b60e08301525061010061193a84828501611418565b6101008301525061012082015167ffffffffffffffff8111156119605761195f6114b2565b5b61196c848285016115a9565b6101208301525061014082015167ffffffffffffffff811115611992576119916114b2565b5b61199e848285016115a9565b6101408301525061016082015167ffffffffffffffff8111156119c4576119c36114b2565b5b6119d0848285016115a9565b6101608301525061018082015167ffffffffffffffff8111156119f6576119f56114b2565b5b611a02848285016115a9565b6101808301525092915050565b5f611a21611a1c84611483565b6108d3565b90508083825260208201905060208402830185811115611a4457611a4361117b565b5b835b81811015611a8b57805167ffffffffffffffff811115611a6957611a6861085d565b5b808601611a768982611816565b85526020850194505050602081019050611a46565b5050509392505050565b5f82601f830112611aa957611aa861085d565b5b8151611ab9848260208601611a0f565b91505092915050565b5f60208284031215611ad757611ad66107e2565b5b5f82015167ffffffffffffffff811115611af457611af36107e6565b5b611b0084828501611a95565b91505092915050565b5f60208284031215611b1e57611b1d6107e2565b5b5f82015167ffffffffffffffff811115611b3b57611b3a6107e6565b5b611b4784828501611816565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f611b8782610a38565b9050919050565b611b9781611b7d565b82525050565b5f82825260208201905092915050565b5f611bb782610d18565b611bc18185611b9d565b935083602082028501611bd385610d32565b805f5b85811015611c0e5784840389528151611bef8582610d93565b9450611bfa83610da6565b925060208a01995050600181019050611bd6565b50829750879550505050505092915050565b5f606082019050611c335f830186611b8e565b611c406020830185611b8e565b8181036040830152611c528184611bad565b9050949350505050565b5f82825260208201905092915050565b5f611c778385611c5c565b9350611c8483858461091d565b611c8d83610865565b840190509392505050565b5f6040820190508181035f830152611cb1818688611c6c565b90508181036020830152611cc6818486611c6c565b905095945050505050565b5f60208284031215611ce657611ce56107e2565b5b5f611cf384828501611418565b9150509291505056fea2646970667358221220edf29f0cd7e5959303b1f68c6f94e27d443b0b3e45b01246f2d9a5a24b861de464736f6c63430008140033",
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
