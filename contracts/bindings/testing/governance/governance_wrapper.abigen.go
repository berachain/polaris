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

// CosmosCodecAny is an auto generated low-level Go binding around an user-defined struct.
type CosmosCodecAny struct {
	TypeURL string
	Value   []uint8
}

// CosmosCoin is an auto generated low-level Go binding around an user-defined struct.
type CosmosCoin struct {
	Amount *big.Int
	Denom  string
}

// IGovernanceModuleMsgSubmitProposal is an auto generated low-level Go binding around an user-defined struct.
type IGovernanceModuleMsgSubmitProposal struct {
	Messages       []CosmosCodecAny
	InitialDeposit []CosmosCoin
	Proposer       common.Address
	Metadata       string
	Title          string
	Summary        string
	Expedited      bool
}

// IGovernanceModuleProposal is an auto generated low-level Go binding around an user-defined struct.
type IGovernanceModuleProposal struct {
	Id               uint64
	Messages         []CosmosCodecAny
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
	Proposer         common.Address
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governanceModule\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"contractIBankModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"cancelProposal\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"typeURL\",\"type\":\"string\"},{\"internalType\":\"uint8[]\",\"name\":\"value\",\"type\":\"uint8[]\"}],\"internalType\":\"structCosmos.CodecAny[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.TallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"submitTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositEndTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structCosmos.Coin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"votingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"}],\"internalType\":\"structIGovernanceModule.Proposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"proposalStatus\",\"type\":\"int32\"}],\"name\":\"getProposals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"typeURL\",\"type\":\"string\"},{\"internalType\":\"uint8[]\",\"name\":\"value\",\"type\":\"uint8[]\"}],\"internalType\":\"structCosmos.CodecAny[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.TallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"submitTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositEndTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structCosmos.Coin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"votingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"}],\"internalType\":\"structIGovernanceModule.Proposal[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceModule\",\"outputs\":[{\"internalType\":\"contractIGovernanceModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"typeURL\",\"type\":\"string\"},{\"internalType\":\"uint8[]\",\"name\":\"value\",\"type\":\"uint8[]\"}],\"internalType\":\"structCosmos.CodecAny[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structCosmos.Coin[]\",\"name\":\"initialDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"expedited\",\"type\":\"bool\"}],\"internalType\":\"structIGovernanceModule.MsgSubmitProposal\",\"name\":\"proposal\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"submit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"internalType\":\"int32\",\"name\":\"option\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a0604052734381dc2ab14285160c808659aee005d51255add773ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801562000057575f80fd5b50604051620027773803806200277783398181016040528101906200007d91906200018e565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620000e3576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050620001be565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f62000158826200012d565b9050919050565b6200016a816200014c565b811462000175575f80fd5b50565b5f8151905062000188816200015f565b92915050565b5f60208284031215620001a657620001a562000129565b5b5f620001b58482850162000178565b91505092915050565b6080516125a0620001d75f395f61036101526125a05ff3fe608060405260043610610073575f3560e01c806376cdb03b1161004d57806376cdb03b146101215780637752e8e41461014b578063b5828df21461017b578063f1610a28146101b75761007a565b806319f7a0fb1461007e5780632b0a7032146100ba57806337a9a59e146100e45761007a565b3661007a57005b5f80fd5b348015610089575f80fd5b506100a4600480360381019061009f919061095c565b6101f3565b6040516100b191906109e2565b60405180910390f35b3480156100c5575f80fd5b506100ce61029a565b6040516100db9190610a75565b60405180910390f35b3480156100ef575f80fd5b5061010a60048036038101906101059190610a8e565b6102bd565b604051610118929190610ac8565b60405180910390f35b34801561012c575f80fd5b5061013561035f565b6040516101429190610b0f565b60405180910390f35b61016560048036038101906101609190610bda565b610383565b6040516101729190610c67565b60405180910390f35b348015610186575f80fd5b506101a1600480360381019061019c9190610c80565b610508565b6040516101ae919061128f565b60405180910390f35b3480156101c2575f80fd5b506101dd60048036038101906101d89190610a8e565b6105bd565b6040516101ea91906113e8565b60405180910390f35b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166319f7a0fb8585856040518463ffffffff1660e01b81526004016102519392919061145f565b6020604051808303815f875af115801561026d573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061029191906114c5565b90509392505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f805f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166337a9a59e846040518263ffffffff1660e01b81526004016103179190610c67565b60408051808303815f875af1158015610332573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103569190611504565b91509150915091565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f80600167ffffffffffffffff8111156103a05761039f610838565b5b6040519080825280602002602001820160405280156103d957816020015b6103c6610666565b8152602001906001900390816103be5790505b50905084848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815f8151811061043357610432611542565b5b60200260200101516020018190525082815f8151811061045657610455611542565b5b60200260200101515f0181815250505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e4112572876040518263ffffffff1660e01b81526004016104bd9190611b88565b6020604051808303815f875af11580156104d9573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104fd9190611ba8565b915050949350505050565b606061051261067f565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663917c9d9285846040518363ffffffff1660e01b815260040161056e929190611c46565b5f60405180830381865afa158015610588573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906105b091906124ad565b5090508092505050919050565b6105c56106c2565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f1610a28836040518263ffffffff1660e01b815260040161061d9190610c67565b5f60405180830381865afa158015610637573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061065f9190612523565b9050919050565b60405180604001604052805f8152602001606081525090565b6040518060a00160405280606081526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f151581526020015f151581525090565b604051806101a001604052805f67ffffffffffffffff168152602001606081526020015f60030b81526020016106f6610774565b81526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff168152602001606081526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020016060815260200160608152602001606081526020015f73ffffffffffffffffffffffffffffffffffffffff1681525090565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b5f604051905090565b5f80fd5b5f80fd5b5f67ffffffffffffffff82169050919050565b6107c9816107ad565b81146107d3575f80fd5b50565b5f813590506107e4816107c0565b92915050565b5f8160030b9050919050565b6107ff816107ea565b8114610809575f80fd5b50565b5f8135905061081a816107f6565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61086e82610828565b810181811067ffffffffffffffff8211171561088d5761088c610838565b5b80604052505050565b5f61089f61079c565b90506108ab8282610865565b919050565b5f67ffffffffffffffff8211156108ca576108c9610838565b5b6108d382610828565b9050602081019050919050565b828183375f83830152505050565b5f6109006108fb846108b0565b610896565b90508281526020810184848401111561091c5761091b610824565b5b6109278482856108e0565b509392505050565b5f82601f83011261094357610942610820565b5b81356109538482602086016108ee565b91505092915050565b5f805f60608486031215610973576109726107a5565b5b5f610980868287016107d6565b93505060206109918682870161080c565b925050604084013567ffffffffffffffff8111156109b2576109b16107a9565b5b6109be8682870161092f565b9150509250925092565b5f8115159050919050565b6109dc816109c8565b82525050565b5f6020820190506109f55f8301846109d3565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f610a3d610a38610a33846109fb565b610a1a565b6109fb565b9050919050565b5f610a4e82610a23565b9050919050565b5f610a5f82610a44565b9050919050565b610a6f81610a55565b82525050565b5f602082019050610a885f830184610a66565b92915050565b5f60208284031215610aa357610aa26107a5565b5b5f610ab0848285016107d6565b91505092915050565b610ac2816107ad565b82525050565b5f604082019050610adb5f830185610ab9565b610ae86020830184610ab9565b9392505050565b5f610af982610a44565b9050919050565b610b0981610aef565b82525050565b5f602082019050610b225f830184610b00565b92915050565b5f80fd5b5f60e08284031215610b4157610b40610b28565b5b81905092915050565b5f80fd5b5f80fd5b5f8083601f840112610b6757610b66610820565b5b8235905067ffffffffffffffff811115610b8457610b83610b4a565b5b602083019150836001820283011115610ba057610b9f610b4e565b5b9250929050565b5f819050919050565b610bb981610ba7565b8114610bc3575f80fd5b50565b5f81359050610bd481610bb0565b92915050565b5f805f8060608587031215610bf257610bf16107a5565b5b5f85013567ffffffffffffffff811115610c0f57610c0e6107a9565b5b610c1b87828801610b2c565b945050602085013567ffffffffffffffff811115610c3c57610c3b6107a9565b5b610c4887828801610b52565b93509350506040610c5b87828801610bc6565b91505092959194509250565b5f602082019050610c7a5f830184610ab9565b92915050565b5f60208284031215610c9557610c946107a5565b5b5f610ca28482850161080c565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610cdd816107ad565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610d43578082015181840152602081019050610d28565b5f8484015250505050565b5f610d5882610d0c565b610d628185610d16565b9350610d72818560208601610d26565b610d7b81610828565b840191505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f60ff82169050919050565b610dc481610daf565b82525050565b5f610dd58383610dbb565b60208301905092915050565b5f602082019050919050565b5f610df782610d86565b610e018185610d90565b9350610e0c83610da0565b805f5b83811015610e3c578151610e238882610dca565b9750610e2e83610de1565b925050600181019050610e0f565b5085935050505092915050565b5f604083015f8301518482035f860152610e638282610d4e565b91505060208301518482036020860152610e7d8282610ded565b9150508091505092915050565b5f610e958383610e49565b905092915050565b5f602082019050919050565b5f610eb382610ce3565b610ebd8185610ced565b935083602082028501610ecf85610cfd565b805f5b85811015610f0a5784840389528151610eeb8582610e8a565b9450610ef683610e9d565b925060208a01995050600181019050610ed2565b50829750879550505050505092915050565b610f25816107ea565b82525050565b5f608083015f8301518482035f860152610f458282610d4e565b91505060208301518482036020860152610f5f8282610d4e565b91505060408301518482036040860152610f798282610d4e565b91505060608301518482036060860152610f938282610d4e565b9150508091505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610fd281610ba7565b82525050565b5f604083015f830151610fed5f860182610fc9565b50602083015184820360208601526110058282610d4e565b9150508091505092915050565b5f61101d8383610fd8565b905092915050565b5f602082019050919050565b5f61103b82610fa0565b6110458185610faa565b93508360208202850161105785610fba565b805f5b8581101561109257848403895281516110738582611012565b945061107e83611025565b925060208a0199505060018101905061105a565b50829750879550505050505092915050565b5f6110ae826109fb565b9050919050565b6110be816110a4565b82525050565b5f6101a083015f8301516110da5f860182610cd4565b50602083015184820360208601526110f28282610ea9565b91505060408301516111076040860182610f1c565b506060830151848203606086015261111f8282610f2b565b91505060808301516111346080860182610cd4565b5060a083015161114760a0860182610cd4565b5060c083015184820360c086015261115f8282611031565b91505060e083015161117460e0860182610cd4565b50610100830151611189610100860182610cd4565b506101208301518482036101208601526111a38282610d4e565b9150506101408301518482036101408601526111bf8282610d4e565b9150506101608301518482036101608601526111db8282610d4e565b9150506101808301516111f26101808601826110b5565b508091505092915050565b5f61120883836110c4565b905092915050565b5f602082019050919050565b5f61122682610cab565b6112308185610cb5565b93508360208202850161124285610cc5565b805f5b8581101561127d578484038952815161125e85826111fd565b945061126983611210565b925060208a01995050600181019050611245565b50829750879550505050505092915050565b5f6020820190508181035f8301526112a7818461121c565b905092915050565b5f6101a083015f8301516112c55f860182610cd4565b50602083015184820360208601526112dd8282610ea9565b91505060408301516112f26040860182610f1c565b506060830151848203606086015261130a8282610f2b565b915050608083015161131f6080860182610cd4565b5060a083015161133260a0860182610cd4565b5060c083015184820360c086015261134a8282611031565b91505060e083015161135f60e0860182610cd4565b50610100830151611374610100860182610cd4565b5061012083015184820361012086015261138e8282610d4e565b9150506101408301518482036101408601526113aa8282610d4e565b9150506101608301518482036101608601526113c68282610d4e565b9150506101808301516113dd6101808601826110b5565b508091505092915050565b5f6020820190508181035f83015261140081846112af565b905092915050565b611411816107ea565b82525050565b5f82825260208201905092915050565b5f61143182610d0c565b61143b8185611417565b935061144b818560208601610d26565b61145481610828565b840191505092915050565b5f6060820190506114725f830186610ab9565b61147f6020830185611408565b81810360408301526114918184611427565b9050949350505050565b6114a4816109c8565b81146114ae575f80fd5b50565b5f815190506114bf8161149b565b92915050565b5f602082840312156114da576114d96107a5565b5b5f6114e7848285016114b1565b91505092915050565b5f815190506114fe816107c0565b92915050565b5f806040838503121561151a576115196107a5565b5b5f611527858286016114f0565b9250506020611538858286016114f0565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f80fd5b5f80fd5b5f80fd5b5f808335600160200384360303811261159757611596611577565b5b83810192508235915060208301925067ffffffffffffffff8211156115bf576115be61156f565b5b6020820236038313156115d5576115d4611573565b5b509250929050565b5f819050919050565b5f808335600160200384360303811261160257611601611577565b5b83810192508235915060208301925067ffffffffffffffff82111561162a5761162961156f565b5b6001820236038313156116405761163f611573565b5b509250929050565b5f6116538385610d16565b93506116608385846108e0565b61166983610828565b840190509392505050565b5f80833560016020038436030381126116905761168f611577565b5b83810192508235915060208301925067ffffffffffffffff8211156116b8576116b761156f565b5b6020820236038313156116ce576116cd611573565b5b509250929050565b5f819050919050565b6116e881610daf565b81146116f2575f80fd5b50565b5f81359050611703816116df565b92915050565b5f61171760208401846116f5565b905092915050565b5f602082019050919050565b5f6117368385610d90565b9350611741826116d6565b805f5b85811015611779576117568284611709565b6117608882610dca565b975061176b8361171f565b925050600181019050611744565b5085925050509392505050565b5f604083016117975f8401846115e6565b8583035f8701526117a9838284611648565b925050506117ba6020840184611674565b85830360208701526117cd83828461172b565b925050508091505092915050565b5f6117e68383611786565b905092915050565b5f8235600160400383360303811261180957611808611577565b5b82810191505092915050565b5f602082019050919050565b5f61182c8385610ced565b93508360208402850161183e846115dd565b805f5b8781101561188157848403895261185882846117ee565b61186285826117db565b945061186d83611815565b925060208a01995050600181019050611841565b50829750879450505050509392505050565b5f80833560016020038436030381126118af576118ae611577565b5b83810192508235915060208301925067ffffffffffffffff8211156118d7576118d661156f565b5b6020820236038313156118ed576118ec611573565b5b509250929050565b5f819050919050565b5f61190c6020840184610bc6565b905092915050565b5f604083016119255f8401846118fe565b6119315f860182610fc9565b5061193f60208401846115e6565b8583036020870152611952838284611648565b925050508091505092915050565b5f61196b8383611914565b905092915050565b5f8235600160400383360303811261198e5761198d611577565b5b82810191505092915050565b5f602082019050919050565b5f6119b18385610faa565b9350836020840285016119c3846118f5565b805f5b87811015611a065784840389526119dd8284611973565b6119e78582611960565b94506119f28361199a565b925060208a019950506001810190506119c6565b50829750879450505050509392505050565b611a21816110a4565b8114611a2b575f80fd5b50565b5f81359050611a3c81611a18565b92915050565b5f611a506020840184611a2e565b905092915050565b5f81359050611a668161149b565b92915050565b5f611a7a6020840184611a58565b905092915050565b611a8b816109c8565b82525050565b5f60e08301611aa25f84018461157b565b8583035f870152611ab4838284611821565b92505050611ac56020840184611893565b8583036020870152611ad88382846119a6565b92505050611ae96040840184611a42565b611af660408601826110b5565b50611b0460608401846115e6565b8583036060870152611b17838284611648565b92505050611b2860808401846115e6565b8583036080870152611b3b838284611648565b92505050611b4c60a08401846115e6565b85830360a0870152611b5f838284611648565b92505050611b7060c0840184611a6c565b611b7d60c0860182611a82565b508091505092915050565b5f6020820190508181035f830152611ba08184611a91565b905092915050565b5f60208284031215611bbd57611bbc6107a5565b5b5f611bca848285016114f0565b91505092915050565b5f60a083015f8301518482035f860152611bed8282610d4e565b9150506020830151611c026020860182610cd4565b506040830151611c156040860182610cd4565b506060830151611c286060860182611a82565b506080830151611c3b6080860182611a82565b508091505092915050565b5f604082019050611c595f830185611408565b8181036020830152611c6b8184611bd3565b90509392505050565b5f67ffffffffffffffff821115611c8e57611c8d610838565b5b602082029050602081019050919050565b5f80fd5b5f80fd5b5f67ffffffffffffffff821115611cc157611cc0610838565b5b602082029050602081019050919050565b5f611ce4611cdf846108b0565b610896565b905082815260208101848484011115611d0057611cff610824565b5b611d0b848285610d26565b509392505050565b5f82601f830112611d2757611d26610820565b5b8151611d37848260208601611cd2565b91505092915050565b5f67ffffffffffffffff821115611d5a57611d59610838565b5b602082029050602081019050919050565b5f81519050611d79816116df565b92915050565b5f611d91611d8c84611d40565b610896565b90508083825260208201905060208402830185811115611db457611db3610b4e565b5b835b81811015611ddd5780611dc98882611d6b565b845260208401935050602081019050611db6565b5050509392505050565b5f82601f830112611dfb57611dfa610820565b5b8151611e0b848260208601611d7f565b91505092915050565b5f60408284031215611e2957611e28611c9f565b5b611e336040610896565b90505f82015167ffffffffffffffff811115611e5257611e51611ca3565b5b611e5e84828501611d13565b5f83015250602082015167ffffffffffffffff811115611e8157611e80611ca3565b5b611e8d84828501611de7565b60208301525092915050565b5f611eab611ea684611ca7565b610896565b90508083825260208201905060208402830185811115611ece57611ecd610b4e565b5b835b81811015611f1557805167ffffffffffffffff811115611ef357611ef2610820565b5b808601611f008982611e14565b85526020850194505050602081019050611ed0565b5050509392505050565b5f82601f830112611f3357611f32610820565b5b8151611f43848260208601611e99565b91505092915050565b5f81519050611f5a816107f6565b92915050565b5f60808284031215611f7557611f74611c9f565b5b611f7f6080610896565b90505f82015167ffffffffffffffff811115611f9e57611f9d611ca3565b5b611faa84828501611d13565b5f83015250602082015167ffffffffffffffff811115611fcd57611fcc611ca3565b5b611fd984828501611d13565b602083015250604082015167ffffffffffffffff811115611ffd57611ffc611ca3565b5b61200984828501611d13565b604083015250606082015167ffffffffffffffff81111561202d5761202c611ca3565b5b61203984828501611d13565b60608301525092915050565b5f67ffffffffffffffff82111561205f5761205e610838565b5b602082029050602081019050919050565b5f8151905061207e81610bb0565b92915050565b5f6040828403121561209957612098611c9f565b5b6120a36040610896565b90505f6120b284828501612070565b5f83015250602082015167ffffffffffffffff8111156120d5576120d4611ca3565b5b6120e184828501611d13565b60208301525092915050565b5f6120ff6120fa84612045565b610896565b9050808382526020820190506020840283018581111561212257612121610b4e565b5b835b8181101561216957805167ffffffffffffffff81111561214757612146610820565b5b8086016121548982612084565b85526020850194505050602081019050612124565b5050509392505050565b5f82601f83011261218757612186610820565b5b81516121978482602086016120ed565b91505092915050565b5f815190506121ae81611a18565b92915050565b5f6101a082840312156121ca576121c9611c9f565b5b6121d56101a0610896565b90505f6121e4848285016114f0565b5f83015250602082015167ffffffffffffffff81111561220757612206611ca3565b5b61221384828501611f1f565b602083015250604061222784828501611f4c565b604083015250606082015167ffffffffffffffff81111561224b5761224a611ca3565b5b61225784828501611f60565b606083015250608061226b848285016114f0565b60808301525060a061227f848285016114f0565b60a08301525060c082015167ffffffffffffffff8111156122a3576122a2611ca3565b5b6122af84828501612173565b60c08301525060e06122c3848285016114f0565b60e0830152506101006122d8848285016114f0565b6101008301525061012082015167ffffffffffffffff8111156122fe576122fd611ca3565b5b61230a84828501611d13565b6101208301525061014082015167ffffffffffffffff8111156123305761232f611ca3565b5b61233c84828501611d13565b6101408301525061016082015167ffffffffffffffff81111561236257612361611ca3565b5b61236e84828501611d13565b61016083015250610180612384848285016121a0565b6101808301525092915050565b5f6123a361239e84611c74565b610896565b905080838252602082019050602084028301858111156123c6576123c5610b4e565b5b835b8181101561240d57805167ffffffffffffffff8111156123eb576123ea610820565b5b8086016123f889826121b4565b855260208501945050506020810190506123c8565b5050509392505050565b5f82601f83011261242b5761242a610820565b5b815161243b848260208601612391565b91505092915050565b5f6040828403121561245957612458611c9f565b5b6124636040610896565b90505f82015167ffffffffffffffff81111561248257612481611ca3565b5b61248e84828501611d13565b5f8301525060206124a1848285016114f0565b60208301525092915050565b5f80604083850312156124c3576124c26107a5565b5b5f83015167ffffffffffffffff8111156124e0576124df6107a9565b5b6124ec85828601612417565b925050602083015167ffffffffffffffff81111561250d5761250c6107a9565b5b61251985828601612444565b9150509250929050565b5f60208284031215612538576125376107a5565b5b5f82015167ffffffffffffffff811115612555576125546107a9565b5b612561848285016121b4565b9150509291505056fea2646970667358221220306d2b017ebeaa83c1fae5293a08fd14264774f7f36721a01dc2fb8dd35515ac64736f6c63430008150033",
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
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,(string,uint8[])[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,address))
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
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,(string,uint8[])[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,address))
func (_GovernanceWrapper *GovernanceWrapperSession) GetProposal(proposalId uint64) (IGovernanceModuleProposal, error) {
	return _GovernanceWrapper.Contract.GetProposal(&_GovernanceWrapper.CallOpts, proposalId)
}

// GetProposal is a free data retrieval call binding the contract method 0xf1610a28.
//
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,(string,uint8[])[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,address))
func (_GovernanceWrapper *GovernanceWrapperCallerSession) GetProposal(proposalId uint64) (IGovernanceModuleProposal, error) {
	return _GovernanceWrapper.Contract.GetProposal(&_GovernanceWrapper.CallOpts, proposalId)
}

// GetProposals is a free data retrieval call binding the contract method 0xb5828df2.
//
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,(string,uint8[])[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,address)[])
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
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,(string,uint8[])[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,address)[])
func (_GovernanceWrapper *GovernanceWrapperSession) GetProposals(proposalStatus int32) ([]IGovernanceModuleProposal, error) {
	return _GovernanceWrapper.Contract.GetProposals(&_GovernanceWrapper.CallOpts, proposalStatus)
}

// GetProposals is a free data retrieval call binding the contract method 0xb5828df2.
//
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,(string,uint8[])[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,address)[])
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

// Submit is a paid mutator transaction binding the contract method 0x7752e8e4.
//
// Solidity: function submit(((string,uint8[])[],(uint256,string)[],address,string,string,string,bool) proposal, string denom, uint256 amount) payable returns(uint64)
func (_GovernanceWrapper *GovernanceWrapperTransactor) Submit(opts *bind.TransactOpts, proposal IGovernanceModuleMsgSubmitProposal, denom string, amount *big.Int) (*types.Transaction, error) {
	return _GovernanceWrapper.contract.Transact(opts, "submit", proposal, denom, amount)
}

// Submit is a paid mutator transaction binding the contract method 0x7752e8e4.
//
// Solidity: function submit(((string,uint8[])[],(uint256,string)[],address,string,string,string,bool) proposal, string denom, uint256 amount) payable returns(uint64)
func (_GovernanceWrapper *GovernanceWrapperSession) Submit(proposal IGovernanceModuleMsgSubmitProposal, denom string, amount *big.Int) (*types.Transaction, error) {
	return _GovernanceWrapper.Contract.Submit(&_GovernanceWrapper.TransactOpts, proposal, denom, amount)
}

// Submit is a paid mutator transaction binding the contract method 0x7752e8e4.
//
// Solidity: function submit(((string,uint8[])[],(uint256,string)[],address,string,string,string,bool) proposal, string denom, uint256 amount) payable returns(uint64)
func (_GovernanceWrapper *GovernanceWrapperTransactorSession) Submit(proposal IGovernanceModuleMsgSubmitProposal, denom string, amount *big.Int) (*types.Transaction, error) {
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
