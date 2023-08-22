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
	Messages         [][]byte
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governanceModule\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"contractIBankModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"cancelProposal\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"bytes[]\",\"name\":\"messages\",\"type\":\"bytes[]\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.TallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"submitTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositEndTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structCosmos.Coin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"votingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposer\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Proposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"proposalStatus\",\"type\":\"int32\"}],\"name\":\"getProposals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"bytes[]\",\"name\":\"messages\",\"type\":\"bytes[]\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.TallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"submitTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositEndTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structCosmos.Coin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"votingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposer\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Proposal[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceModule\",\"outputs\":[{\"internalType\":\"contractIGovernanceModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proposal\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"submit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"internalType\":\"int32\",\"name\":\"option\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a0604052734381dc2ab14285160c808659aee005d51255add773ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801562000057575f80fd5b5060405162001ed538038062001ed583398181016040528101906200007d91906200018e565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620000e3576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050620001be565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f62000158826200012d565b9050919050565b6200016a816200014c565b811462000175575f80fd5b50565b5f8151905062000188816200015f565b92915050565b5f60208284031215620001a657620001a562000129565b5b5f620001b58482850162000178565b91505092915050565b608051611cfe620001d75f395f6104e90152611cfe5ff3fe608060405260043610610073575f3560e01c8063566fbd001161004d578063566fbd001461012157806376cdb03b14610151578063b5828df21461017b578063f1610a28146101b75761007a565b806319f7a0fb1461007e5780632b0a7032146100ba57806337a9a59e146100e45761007a565b3661007a57005b5f80fd5b348015610089575f80fd5b506100a4600480360381019061009f91906108f5565b6101f3565b6040516100b1919061097b565b60405180910390f35b3480156100c5575f80fd5b506100ce61029a565b6040516100db9190610a0e565b60405180910390f35b3480156100ef575f80fd5b5061010a60048036038101906101059190610a27565b6102bd565b604051610118929190610a61565b60405180910390f35b61013b60048036038101906101369190610b6d565b61035f565b6040516101489190610bfe565b60405180910390f35b34801561015c575f80fd5b506101656104e7565b6040516101729190610c37565b60405180910390f35b348015610186575f80fd5b506101a1600480360381019061019c9190610c50565b61050b565b6040516101ae9190611194565b60405180910390f35b3480156101c2575f80fd5b506101dd60048036038101906101d89190610a27565b6105ae565b6040516101ea91906112f4565b60405180910390f35b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166319f7a0fb8585856040518463ffffffff1660e01b81526004016102519392919061136b565b6020604051808303815f875af115801561026d573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061029191906113d1565b90509392505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f805f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166337a9a59e846040518263ffffffff1660e01b81526004016103179190610bfe565b60408051808303815f875af1158015610332573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103569190611410565b91509150915091565b5f80600167ffffffffffffffff81111561037c5761037b6107d1565b5b6040519080825280602002602001820160405280156103b557816020015b6103a2610657565b81526020019060019003908161039a5790505b50905084848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815f8151811061040f5761040e61144e565b5b60200260200101516020018190525082815f815181106104325761043161144e565b5b60200260200101515f0181815250505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d238313688886040518363ffffffff1660e01b815260040161049b9291906114b7565b6020604051808303815f875af11580156104b7573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104db91906114d9565b91505095945050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60605f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b5828df2836040518263ffffffff1660e01b81526004016105659190611504565b5f60405180830381865afa15801561057f573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906105a79190611c3a565b9050919050565b6105b6610670565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f1610a28836040518263ffffffff1660e01b815260040161060e9190610bfe565b5f60405180830381865afa158015610628573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906106509190611c81565b9050919050565b60405180604001604052805f8152602001606081525090565b604051806101a001604052805f67ffffffffffffffff168152602001606081526020015f60030b81526020016106a461070d565b81526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff168152602001606081526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff168152602001606081526020016060815260200160608152602001606081525090565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b5f604051905090565b5f80fd5b5f80fd5b5f67ffffffffffffffff82169050919050565b61076281610746565b811461076c575f80fd5b50565b5f8135905061077d81610759565b92915050565b5f8160030b9050919050565b61079881610783565b81146107a2575f80fd5b50565b5f813590506107b38161078f565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610807826107c1565b810181811067ffffffffffffffff82111715610826576108256107d1565b5b80604052505050565b5f610838610735565b905061084482826107fe565b919050565b5f67ffffffffffffffff821115610863576108626107d1565b5b61086c826107c1565b9050602081019050919050565b828183375f83830152505050565b5f61089961089484610849565b61082f565b9050828152602081018484840111156108b5576108b46107bd565b5b6108c0848285610879565b509392505050565b5f82601f8301126108dc576108db6107b9565b5b81356108ec848260208601610887565b91505092915050565b5f805f6060848603121561090c5761090b61073e565b5b5f6109198682870161076f565b935050602061092a868287016107a5565b925050604084013567ffffffffffffffff81111561094b5761094a610742565b5b610957868287016108c8565b9150509250925092565b5f8115159050919050565b61097581610961565b82525050565b5f60208201905061098e5f83018461096c565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f6109d66109d16109cc84610994565b6109b3565b610994565b9050919050565b5f6109e7826109bc565b9050919050565b5f6109f8826109dd565b9050919050565b610a08816109ee565b82525050565b5f602082019050610a215f8301846109ff565b92915050565b5f60208284031215610a3c57610a3b61073e565b5b5f610a498482850161076f565b91505092915050565b610a5b81610746565b82525050565b5f604082019050610a745f830185610a52565b610a816020830184610a52565b9392505050565b5f80fd5b5f80fd5b5f8083601f840112610aa557610aa46107b9565b5b8235905067ffffffffffffffff811115610ac257610ac1610a88565b5b602083019150836001820283011115610ade57610add610a8c565b5b9250929050565b5f8083601f840112610afa57610af96107b9565b5b8235905067ffffffffffffffff811115610b1757610b16610a88565b5b602083019150836001820283011115610b3357610b32610a8c565b5b9250929050565b5f819050919050565b610b4c81610b3a565b8114610b56575f80fd5b50565b5f81359050610b6781610b43565b92915050565b5f805f805f60608688031215610b8657610b8561073e565b5b5f86013567ffffffffffffffff811115610ba357610ba2610742565b5b610baf88828901610a90565b9550955050602086013567ffffffffffffffff811115610bd257610bd1610742565b5b610bde88828901610ae5565b93509350506040610bf188828901610b59565b9150509295509295909350565b5f602082019050610c115f830184610a52565b92915050565b5f610c21826109dd565b9050919050565b610c3181610c17565b82525050565b5f602082019050610c4a5f830184610c28565b92915050565b5f60208284031215610c6557610c6461073e565b5b5f610c72848285016107a5565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610cad81610746565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610d13578082015181840152602081019050610cf8565b5f8484015250505050565b5f610d2882610cdc565b610d328185610ce6565b9350610d42818560208601610cf6565b610d4b816107c1565b840191505092915050565b5f610d618383610d1e565b905092915050565b5f602082019050919050565b5f610d7f82610cb3565b610d898185610cbd565b935083602082028501610d9b85610ccd565b805f5b85811015610dd65784840389528151610db78582610d56565b9450610dc283610d69565b925060208a01995050600181019050610d9e565b50829750879550505050505092915050565b610df181610783565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f610e1b82610df7565b610e258185610e01565b9350610e35818560208601610cf6565b610e3e816107c1565b840191505092915050565b5f608083015f8301518482035f860152610e638282610e11565b91505060208301518482036020860152610e7d8282610e11565b91505060408301518482036040860152610e978282610e11565b91505060608301518482036060860152610eb18282610e11565b9150508091505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610ef081610b3a565b82525050565b5f604083015f830151610f0b5f860182610ee7565b5060208301518482036020860152610f238282610e11565b9150508091505092915050565b5f610f3b8383610ef6565b905092915050565b5f602082019050919050565b5f610f5982610ebe565b610f638185610ec8565b935083602082028501610f7585610ed8565b805f5b85811015610fb05784840389528151610f918582610f30565b9450610f9c83610f43565b925060208a01995050600181019050610f78565b50829750879550505050505092915050565b5f6101a083015f830151610fd85f860182610ca4565b5060208301518482036020860152610ff08282610d75565b91505060408301516110056040860182610de8565b506060830151848203606086015261101d8282610e49565b91505060808301516110326080860182610ca4565b5060a083015161104560a0860182610ca4565b5060c083015184820360c086015261105d8282610f4f565b91505060e083015161107260e0860182610ca4565b50610100830151611087610100860182610ca4565b506101208301518482036101208601526110a18282610e11565b9150506101408301518482036101408601526110bd8282610e11565b9150506101608301518482036101608601526110d98282610e11565b9150506101808301518482036101808601526110f58282610e11565b9150508091505092915050565b5f61110d8383610fc2565b905092915050565b5f602082019050919050565b5f61112b82610c7b565b6111358185610c85565b93508360208202850161114785610c95565b805f5b8581101561118257848403895281516111638582611102565b945061116e83611115565b925060208a0199505060018101905061114a565b50829750879550505050505092915050565b5f6020820190508181035f8301526111ac8184611121565b905092915050565b5f6101a083015f8301516111ca5f860182610ca4565b50602083015184820360208601526111e28282610d75565b91505060408301516111f76040860182610de8565b506060830151848203606086015261120f8282610e49565b91505060808301516112246080860182610ca4565b5060a083015161123760a0860182610ca4565b5060c083015184820360c086015261124f8282610f4f565b91505060e083015161126460e0860182610ca4565b50610100830151611279610100860182610ca4565b506101208301518482036101208601526112938282610e11565b9150506101408301518482036101408601526112af8282610e11565b9150506101608301518482036101608601526112cb8282610e11565b9150506101808301518482036101808601526112e78282610e11565b9150508091505092915050565b5f6020820190508181035f83015261130c81846111b4565b905092915050565b61131d81610783565b82525050565b5f82825260208201905092915050565b5f61133d82610df7565b6113478185611323565b9350611357818560208601610cf6565b611360816107c1565b840191505092915050565b5f60608201905061137e5f830186610a52565b61138b6020830185611314565b818103604083015261139d8184611333565b9050949350505050565b6113b081610961565b81146113ba575f80fd5b50565b5f815190506113cb816113a7565b92915050565b5f602082840312156113e6576113e561073e565b5b5f6113f3848285016113bd565b91505092915050565b5f8151905061140a81610759565b92915050565b5f80604083850312156114265761142561073e565b5b5f611433858286016113fc565b9250506020611444858286016113fc565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f82825260208201905092915050565b5f611496838561147b565b93506114a3838584610879565b6114ac836107c1565b840190509392505050565b5f6020820190508181035f8301526114d081848661148b565b90509392505050565b5f602082840312156114ee576114ed61073e565b5b5f6114fb848285016113fc565b91505092915050565b5f6020820190506115175f830184611314565b92915050565b5f67ffffffffffffffff821115611537576115366107d1565b5b602082029050602081019050919050565b5f80fd5b5f80fd5b5f67ffffffffffffffff82111561156a576115696107d1565b5b602082029050602081019050919050565b5f67ffffffffffffffff821115611595576115946107d1565b5b61159e826107c1565b9050602081019050919050565b5f6115bd6115b88461157b565b61082f565b9050828152602081018484840111156115d9576115d86107bd565b5b6115e4848285610cf6565b509392505050565b5f82601f830112611600576115ff6107b9565b5b81516116108482602086016115ab565b91505092915050565b5f61162b61162684611550565b61082f565b9050808382526020820190506020840283018581111561164e5761164d610a8c565b5b835b8181101561169557805167ffffffffffffffff811115611673576116726107b9565b5b80860161168089826115ec565b85526020850194505050602081019050611650565b5050509392505050565b5f82601f8301126116b3576116b26107b9565b5b81516116c3848260208601611619565b91505092915050565b5f815190506116da8161078f565b92915050565b5f6116f26116ed84610849565b61082f565b90508281526020810184848401111561170e5761170d6107bd565b5b611719848285610cf6565b509392505050565b5f82601f830112611735576117346107b9565b5b81516117458482602086016116e0565b91505092915050565b5f6080828403121561176357611762611548565b5b61176d608061082f565b90505f82015167ffffffffffffffff81111561178c5761178b61154c565b5b61179884828501611721565b5f83015250602082015167ffffffffffffffff8111156117bb576117ba61154c565b5b6117c784828501611721565b602083015250604082015167ffffffffffffffff8111156117eb576117ea61154c565b5b6117f784828501611721565b604083015250606082015167ffffffffffffffff81111561181b5761181a61154c565b5b61182784828501611721565b60608301525092915050565b5f67ffffffffffffffff82111561184d5761184c6107d1565b5b602082029050602081019050919050565b5f8151905061186c81610b43565b92915050565b5f6040828403121561188757611886611548565b5b611891604061082f565b90505f6118a08482850161185e565b5f83015250602082015167ffffffffffffffff8111156118c3576118c261154c565b5b6118cf84828501611721565b60208301525092915050565b5f6118ed6118e884611833565b61082f565b905080838252602082019050602084028301858111156119105761190f610a8c565b5b835b8181101561195757805167ffffffffffffffff811115611935576119346107b9565b5b8086016119428982611872565b85526020850194505050602081019050611912565b5050509392505050565b5f82601f830112611975576119746107b9565b5b81516119858482602086016118db565b91505092915050565b5f6101a082840312156119a4576119a3611548565b5b6119af6101a061082f565b90505f6119be848285016113fc565b5f83015250602082015167ffffffffffffffff8111156119e1576119e061154c565b5b6119ed8482850161169f565b6020830152506040611a01848285016116cc565b604083015250606082015167ffffffffffffffff811115611a2557611a2461154c565b5b611a318482850161174e565b6060830152506080611a45848285016113fc565b60808301525060a0611a59848285016113fc565b60a08301525060c082015167ffffffffffffffff811115611a7d57611a7c61154c565b5b611a8984828501611961565b60c08301525060e0611a9d848285016113fc565b60e083015250610100611ab2848285016113fc565b6101008301525061012082015167ffffffffffffffff811115611ad857611ad761154c565b5b611ae484828501611721565b6101208301525061014082015167ffffffffffffffff811115611b0a57611b0961154c565b5b611b1684828501611721565b6101408301525061016082015167ffffffffffffffff811115611b3c57611b3b61154c565b5b611b4884828501611721565b6101608301525061018082015167ffffffffffffffff811115611b6e57611b6d61154c565b5b611b7a84828501611721565b6101808301525092915050565b5f611b99611b948461151d565b61082f565b90508083825260208201905060208402830185811115611bbc57611bbb610a8c565b5b835b81811015611c0357805167ffffffffffffffff811115611be157611be06107b9565b5b808601611bee898261198e565b85526020850194505050602081019050611bbe565b5050509392505050565b5f82601f830112611c2157611c206107b9565b5b8151611c31848260208601611b87565b91505092915050565b5f60208284031215611c4f57611c4e61073e565b5b5f82015167ffffffffffffffff811115611c6c57611c6b610742565b5b611c7884828501611c0d565b91505092915050565b5f60208284031215611c9657611c9561073e565b5b5f82015167ffffffffffffffff811115611cb357611cb2610742565b5b611cbf8482850161198e565b9150509291505056fea264697066735822122061d0eddf0eebde54dc90ccf9ef4c08de6d4b93a21ece6bbb2b168590b145344564736f6c63430008140033",
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
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,bytes[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string))
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
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,bytes[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string))
func (_GovernanceWrapper *GovernanceWrapperSession) GetProposal(proposalId uint64) (IGovernanceModuleProposal, error) {
	return _GovernanceWrapper.Contract.GetProposal(&_GovernanceWrapper.CallOpts, proposalId)
}

// GetProposal is a free data retrieval call binding the contract method 0xf1610a28.
//
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,bytes[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string))
func (_GovernanceWrapper *GovernanceWrapperCallerSession) GetProposal(proposalId uint64) (IGovernanceModuleProposal, error) {
	return _GovernanceWrapper.Contract.GetProposal(&_GovernanceWrapper.CallOpts, proposalId)
}

// GetProposals is a free data retrieval call binding the contract method 0xb5828df2.
//
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,bytes[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string)[])
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
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,bytes[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string)[])
func (_GovernanceWrapper *GovernanceWrapperSession) GetProposals(proposalStatus int32) ([]IGovernanceModuleProposal, error) {
	return _GovernanceWrapper.Contract.GetProposals(&_GovernanceWrapper.CallOpts, proposalStatus)
}

// GetProposals is a free data retrieval call binding the contract method 0xb5828df2.
//
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,bytes[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string)[])
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

// Submit is a paid mutator transaction binding the contract method 0x566fbd00.
//
// Solidity: function submit(bytes proposal, string denom, uint256 amount) payable returns(uint64)
func (_GovernanceWrapper *GovernanceWrapperTransactor) Submit(opts *bind.TransactOpts, proposal []byte, denom string, amount *big.Int) (*types.Transaction, error) {
	return _GovernanceWrapper.contract.Transact(opts, "submit", proposal, denom, amount)
}

// Submit is a paid mutator transaction binding the contract method 0x566fbd00.
//
// Solidity: function submit(bytes proposal, string denom, uint256 amount) payable returns(uint64)
func (_GovernanceWrapper *GovernanceWrapperSession) Submit(proposal []byte, denom string, amount *big.Int) (*types.Transaction, error) {
	return _GovernanceWrapper.Contract.Submit(&_GovernanceWrapper.TransactOpts, proposal, denom, amount)
}

// Submit is a paid mutator transaction binding the contract method 0x566fbd00.
//
// Solidity: function submit(bytes proposal, string denom, uint256 amount) payable returns(uint64)
func (_GovernanceWrapper *GovernanceWrapperTransactorSession) Submit(proposal []byte, denom string, amount *big.Int) (*types.Transaction, error) {
	return _GovernanceWrapper.Contract.Submit(&_GovernanceWrapper.TransactOpts, proposal, denom, amount)
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
