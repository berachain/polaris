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

// CosmosPageRequest is an auto generated low-level Go binding around an user-defined struct.
type CosmosPageRequest struct {
	Key        string
	Offset     uint64
	Limit      uint64
	CountTotal bool
	Reverse    bool
}

// CosmosPageResponse is an auto generated low-level Go binding around an user-defined struct.
type CosmosPageResponse struct {
	NextKey string
	Total   uint64
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governanceModule\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"contractIBankModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"cancelProposal\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.TallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"submitTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositEndTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structCosmos.Coin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"votingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposer\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Proposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"proposalStatus\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structCosmos.PageRequest\",\"name\":\"pageRequest\",\"type\":\"tuple\"}],\"name\":\"getProposals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.TallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"submitTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositEndTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structCosmos.Coin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"votingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposer\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Proposal[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"nextKey\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structCosmos.PageResponse\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceModule\",\"outputs\":[{\"internalType\":\"contractIGovernanceModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proposal\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"submit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"internalType\":\"int32\",\"name\":\"option\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a0604052734381dc2ab14285160c808659aee005d51255add773ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801562000057575f80fd5b50604051620020913803806200209183398181016040528101906200007d91906200018e565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620000e3576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050620001be565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f62000158826200012d565b9050919050565b6200016a816200014c565b811462000175575f80fd5b50565b5f8151905062000188816200015f565b92915050565b5f60208284031215620001a657620001a562000129565b5b5f620001b58482850162000178565b91505092915050565b608051611eba620001d75f395f6103620152611eba5ff3fe608060405260043610610073575f3560e01c806376cdb03b1161004d57806376cdb03b14610121578063917c9d921461014b578063f1610a2814610188578063fbab7815146101c45761007a565b806319f7a0fb1461007e5780632b0a7032146100ba57806337a9a59e146100e45761007a565b3661007a57005b5f80fd5b348015610089575f80fd5b506100a4600480360381019061009f919061092d565b6101f4565b6040516100b191906109b3565b60405180910390f35b3480156100c5575f80fd5b506100ce61029b565b6040516100db9190610a46565b60405180910390f35b3480156100ef575f80fd5b5061010a60048036038101906101059190610a5f565b6102be565b604051610118929190610a99565b60405180910390f35b34801561012c575f80fd5b50610135610360565b6040516101429190610ae0565b60405180910390f35b348015610156575f80fd5b50610171600480360381019061016c9190610b1b565b610384565b60405161017f929190611016565b60405180910390f35b348015610193575f80fd5b506101ae60048036038101906101a99190610a5f565b610435565b6040516101bb919061118b565b60405180910390f35b6101de60048036038101906101d99190611287565b6104de565b6040516101eb919061134b565b60405180910390f35b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166319f7a0fb8585856040518463ffffffff1660e01b8152600401610252939291906113bb565b6020604051808303815f875af115801561026e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102929190611421565b90509392505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f805f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166337a9a59e846040518263ffffffff1660e01b8152600401610318919061134b565b60408051808303815f875af1158015610333573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103579190611460565b91509150915091565b7f000000000000000000000000000000000000000000000000000000000000000081565b606061038e61066c565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663917c9d9285856040518363ffffffff1660e01b81526004016103e8929190611624565b5f60405180830381865afa158015610402573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061042a9190611cfa565b915091509250929050565b61043d61068f565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f1610a28836040518263ffffffff1660e01b8152600401610495919061134b565b5f60405180830381865afa1580156104af573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906104d79190611d70565b9050919050565b5f80600167ffffffffffffffff8111156104fb576104fa610809565b5b60405190808252806020026020018201604052801561053457816020015b61052161072c565b8152602001906001900390816105195790505b50905084848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815f8151811061058e5761058d611db7565b5b60200260200101516020018190525082815f815181106105b1576105b0611db7565b5b60200260200101515f0181815250505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663474d7f358a8a8a8a6040518563ffffffff1660e01b815260040161061e9493929190611e20565b6020604051808303815f875af115801561063a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061065e9190611e59565b915050979650505050505050565b6040518060400160405280606081526020015f67ffffffffffffffff1681525090565b604051806101a001604052805f67ffffffffffffffff168152602001606081526020015f60030b81526020016106c3610745565b81526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff168152602001606081526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff168152602001606081526020016060815260200160608152602001606081525090565b60405180604001604052805f8152602001606081525090565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b5f604051905090565b5f80fd5b5f80fd5b5f67ffffffffffffffff82169050919050565b61079a8161077e565b81146107a4575f80fd5b50565b5f813590506107b581610791565b92915050565b5f8160030b9050919050565b6107d0816107bb565b81146107da575f80fd5b50565b5f813590506107eb816107c7565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61083f826107f9565b810181811067ffffffffffffffff8211171561085e5761085d610809565b5b80604052505050565b5f61087061076d565b905061087c8282610836565b919050565b5f67ffffffffffffffff82111561089b5761089a610809565b5b6108a4826107f9565b9050602081019050919050565b828183375f83830152505050565b5f6108d16108cc84610881565b610867565b9050828152602081018484840111156108ed576108ec6107f5565b5b6108f88482856108b1565b509392505050565b5f82601f830112610914576109136107f1565b5b81356109248482602086016108bf565b91505092915050565b5f805f6060848603121561094457610943610776565b5b5f610951868287016107a7565b9350506020610962868287016107dd565b925050604084013567ffffffffffffffff8111156109835761098261077a565b5b61098f86828701610900565b9150509250925092565b5f8115159050919050565b6109ad81610999565b82525050565b5f6020820190506109c65f8301846109a4565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f610a0e610a09610a04846109cc565b6109eb565b6109cc565b9050919050565b5f610a1f826109f4565b9050919050565b5f610a3082610a15565b9050919050565b610a4081610a26565b82525050565b5f602082019050610a595f830184610a37565b92915050565b5f60208284031215610a7457610a73610776565b5b5f610a81848285016107a7565b91505092915050565b610a938161077e565b82525050565b5f604082019050610aac5f830185610a8a565b610ab96020830184610a8a565b9392505050565b5f610aca82610a15565b9050919050565b610ada81610ac0565b82525050565b5f602082019050610af35f830184610ad1565b92915050565b5f80fd5b5f60a08284031215610b1257610b11610af9565b5b81905092915050565b5f8060408385031215610b3157610b30610776565b5b5f610b3e858286016107dd565b925050602083013567ffffffffffffffff811115610b5f57610b5e61077a565b5b610b6b85828601610afd565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610ba78161077e565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610be4578082015181840152602081019050610bc9565b5f8484015250505050565b5f610bf982610bad565b610c038185610bb7565b9350610c13818560208601610bc7565b610c1c816107f9565b840191505092915050565b610c30816107bb565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f610c5a82610c36565b610c648185610c40565b9350610c74818560208601610bc7565b610c7d816107f9565b840191505092915050565b5f608083015f8301518482035f860152610ca28282610c50565b91505060208301518482036020860152610cbc8282610c50565b91505060408301518482036040860152610cd68282610c50565b91505060608301518482036060860152610cf08282610c50565b9150508091505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f819050919050565b610d3881610d26565b82525050565b5f604083015f830151610d535f860182610d2f565b5060208301518482036020860152610d6b8282610c50565b9150508091505092915050565b5f610d838383610d3e565b905092915050565b5f602082019050919050565b5f610da182610cfd565b610dab8185610d07565b935083602082028501610dbd85610d17565b805f5b85811015610df85784840389528151610dd98582610d78565b9450610de483610d8b565b925060208a01995050600181019050610dc0565b50829750879550505050505092915050565b5f6101a083015f830151610e205f860182610b9e565b5060208301518482036020860152610e388282610bef565b9150506040830151610e4d6040860182610c27565b5060608301518482036060860152610e658282610c88565b9150506080830151610e7a6080860182610b9e565b5060a0830151610e8d60a0860182610b9e565b5060c083015184820360c0860152610ea58282610d97565b91505060e0830151610eba60e0860182610b9e565b50610100830151610ecf610100860182610b9e565b50610120830151848203610120860152610ee98282610c50565b915050610140830151848203610140860152610f058282610c50565b915050610160830151848203610160860152610f218282610c50565b915050610180830151848203610180860152610f3d8282610c50565b9150508091505092915050565b5f610f558383610e0a565b905092915050565b5f602082019050919050565b5f610f7382610b75565b610f7d8185610b7f565b935083602082028501610f8f85610b8f565b805f5b85811015610fca5784840389528151610fab8582610f4a565b9450610fb683610f5d565b925060208a01995050600181019050610f92565b50829750879550505050505092915050565b5f604083015f8301518482035f860152610ff68282610c50565b915050602083015161100b6020860182610b9e565b508091505092915050565b5f6040820190508181035f83015261102e8185610f69565b905081810360208301526110428184610fdc565b90509392505050565b5f6101a083015f8301516110615f860182610b9e565b50602083015184820360208601526110798282610bef565b915050604083015161108e6040860182610c27565b50606083015184820360608601526110a68282610c88565b91505060808301516110bb6080860182610b9e565b5060a08301516110ce60a0860182610b9e565b5060c083015184820360c08601526110e68282610d97565b91505060e08301516110fb60e0860182610b9e565b50610100830151611110610100860182610b9e565b5061012083015184820361012086015261112a8282610c50565b9150506101408301518482036101408601526111468282610c50565b9150506101608301518482036101608601526111628282610c50565b91505061018083015184820361018086015261117e8282610c50565b9150508091505092915050565b5f6020820190508181035f8301526111a3818461104b565b905092915050565b5f80fd5b5f80fd5b5f8083601f8401126111c8576111c76107f1565b5b8235905067ffffffffffffffff8111156111e5576111e46111ab565b5b602083019150836001820283011115611201576112006111af565b5b9250929050565b5f8083601f84011261121d5761121c6107f1565b5b8235905067ffffffffffffffff81111561123a576112396111ab565b5b602083019150836001820283011115611256576112556111af565b5b9250929050565b61126681610d26565b8114611270575f80fd5b50565b5f813590506112818161125d565b92915050565b5f805f805f805f6080888a0312156112a2576112a1610776565b5b5f88013567ffffffffffffffff8111156112bf576112be61077a565b5b6112cb8a828b016111b3565b9750975050602088013567ffffffffffffffff8111156112ee576112ed61077a565b5b6112fa8a828b016111b3565b9550955050604088013567ffffffffffffffff81111561131d5761131c61077a565b5b6113298a828b01611208565b9350935050606061133c8a828b01611273565b91505092959891949750929550565b5f60208201905061135e5f830184610a8a565b92915050565b61136d816107bb565b82525050565b5f82825260208201905092915050565b5f61138d82610c36565b6113978185611373565b93506113a7818560208601610bc7565b6113b0816107f9565b840191505092915050565b5f6060820190506113ce5f830186610a8a565b6113db6020830185611364565b81810360408301526113ed8184611383565b9050949350505050565b61140081610999565b811461140a575f80fd5b50565b5f8151905061141b816113f7565b92915050565b5f6020828403121561143657611435610776565b5b5f6114438482850161140d565b91505092915050565b5f8151905061145a81610791565b92915050565b5f806040838503121561147657611475610776565b5b5f6114838582860161144c565b92505060206114948582860161144c565b9150509250929050565b5f80fd5b5f80fd5b5f80fd5b5f80833560016020038436030381126114c6576114c56114a6565b5b83810192508235915060208301925067ffffffffffffffff8211156114ee576114ed61149e565b5b600182023603831315611504576115036114a2565b5b509250929050565b5f6115178385610c40565b93506115248385846108b1565b61152d836107f9565b840190509392505050565b5f61154660208401846107a7565b905092915050565b5f8135905061155c816113f7565b92915050565b5f611570602084018461154e565b905092915050565b61158181610999565b82525050565b5f60a083016115985f8401846114aa565b8583035f8701526115aa83828461150c565b925050506115bb6020840184611538565b6115c86020860182610b9e565b506115d66040840184611538565b6115e36040860182610b9e565b506115f16060840184611562565b6115fe6060860182611578565b5061160c6080840184611562565b6116196080860182611578565b508091505092915050565b5f6040820190506116375f830185611364565b81810360208301526116498184611587565b90509392505050565b5f67ffffffffffffffff82111561166c5761166b610809565b5b602082029050602081019050919050565b5f80fd5b5f80fd5b5f67ffffffffffffffff82111561169f5761169e610809565b5b6116a8826107f9565b9050602081019050919050565b5f6116c76116c284611685565b610867565b9050828152602081018484840111156116e3576116e26107f5565b5b6116ee848285610bc7565b509392505050565b5f82601f83011261170a576117096107f1565b5b815161171a8482602086016116b5565b91505092915050565b5f81519050611731816107c7565b92915050565b5f61174961174484610881565b610867565b905082815260208101848484011115611765576117646107f5565b5b611770848285610bc7565b509392505050565b5f82601f83011261178c5761178b6107f1565b5b815161179c848260208601611737565b91505092915050565b5f608082840312156117ba576117b961167d565b5b6117c46080610867565b90505f82015167ffffffffffffffff8111156117e3576117e2611681565b5b6117ef84828501611778565b5f83015250602082015167ffffffffffffffff81111561181257611811611681565b5b61181e84828501611778565b602083015250604082015167ffffffffffffffff81111561184257611841611681565b5b61184e84828501611778565b604083015250606082015167ffffffffffffffff81111561187257611871611681565b5b61187e84828501611778565b60608301525092915050565b5f67ffffffffffffffff8211156118a4576118a3610809565b5b602082029050602081019050919050565b5f815190506118c38161125d565b92915050565b5f604082840312156118de576118dd61167d565b5b6118e86040610867565b90505f6118f7848285016118b5565b5f83015250602082015167ffffffffffffffff81111561191a57611919611681565b5b61192684828501611778565b60208301525092915050565b5f61194461193f8461188a565b610867565b90508083825260208201905060208402830185811115611967576119666111af565b5b835b818110156119ae57805167ffffffffffffffff81111561198c5761198b6107f1565b5b80860161199989826118c9565b85526020850194505050602081019050611969565b5050509392505050565b5f82601f8301126119cc576119cb6107f1565b5b81516119dc848260208601611932565b91505092915050565b5f6101a082840312156119fb576119fa61167d565b5b611a066101a0610867565b90505f611a158482850161144c565b5f83015250602082015167ffffffffffffffff811115611a3857611a37611681565b5b611a44848285016116f6565b6020830152506040611a5884828501611723565b604083015250606082015167ffffffffffffffff811115611a7c57611a7b611681565b5b611a88848285016117a5565b6060830152506080611a9c8482850161144c565b60808301525060a0611ab08482850161144c565b60a08301525060c082015167ffffffffffffffff811115611ad457611ad3611681565b5b611ae0848285016119b8565b60c08301525060e0611af48482850161144c565b60e083015250610100611b098482850161144c565b6101008301525061012082015167ffffffffffffffff811115611b2f57611b2e611681565b5b611b3b84828501611778565b6101208301525061014082015167ffffffffffffffff811115611b6157611b60611681565b5b611b6d84828501611778565b6101408301525061016082015167ffffffffffffffff811115611b9357611b92611681565b5b611b9f84828501611778565b6101608301525061018082015167ffffffffffffffff811115611bc557611bc4611681565b5b611bd184828501611778565b6101808301525092915050565b5f611bf0611beb84611652565b610867565b90508083825260208201905060208402830185811115611c1357611c126111af565b5b835b81811015611c5a57805167ffffffffffffffff811115611c3857611c376107f1565b5b808601611c4589826119e5565b85526020850194505050602081019050611c15565b5050509392505050565b5f82601f830112611c7857611c776107f1565b5b8151611c88848260208601611bde565b91505092915050565b5f60408284031215611ca657611ca561167d565b5b611cb06040610867565b90505f82015167ffffffffffffffff811115611ccf57611cce611681565b5b611cdb84828501611778565b5f830152506020611cee8482850161144c565b60208301525092915050565b5f8060408385031215611d1057611d0f610776565b5b5f83015167ffffffffffffffff811115611d2d57611d2c61077a565b5b611d3985828601611c64565b925050602083015167ffffffffffffffff811115611d5a57611d5961077a565b5b611d6685828601611c91565b9150509250929050565b5f60208284031215611d8557611d84610776565b5b5f82015167ffffffffffffffff811115611da257611da161077a565b5b611dae848285016119e5565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f82825260208201905092915050565b5f611dff8385611de4565b9350611e0c8385846108b1565b611e15836107f9565b840190509392505050565b5f6040820190508181035f830152611e39818688611df4565b90508181036020830152611e4e818486611df4565b905095945050505050565b5f60208284031215611e6e57611e6d610776565b5b5f611e7b8482850161144c565b9150509291505056fea264697066735822122079b4a291d0f859439095b91d1201920ac6a3f1903d124508fee1cfa51c74642864736f6c63430008140033",
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

// GetProposals is a free data retrieval call binding the contract method 0x917c9d92.
//
// Solidity: function getProposals(int32 proposalStatus, (string,uint64,uint64,bool,bool) pageRequest) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string)[], (string,uint64))
func (_GovernanceWrapper *GovernanceWrapperCaller) GetProposals(opts *bind.CallOpts, proposalStatus int32, pageRequest CosmosPageRequest) ([]IGovernanceModuleProposal, CosmosPageResponse, error) {
	var out []interface{}
	err := _GovernanceWrapper.contract.Call(opts, &out, "getProposals", proposalStatus, pageRequest)

	if err != nil {
		return *new([]IGovernanceModuleProposal), *new(CosmosPageResponse), err
	}

	out0 := *abi.ConvertType(out[0], new([]IGovernanceModuleProposal)).(*[]IGovernanceModuleProposal)
	out1 := *abi.ConvertType(out[1], new(CosmosPageResponse)).(*CosmosPageResponse)

	return out0, out1, err

}

// GetProposals is a free data retrieval call binding the contract method 0x917c9d92.
//
// Solidity: function getProposals(int32 proposalStatus, (string,uint64,uint64,bool,bool) pageRequest) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string)[], (string,uint64))
func (_GovernanceWrapper *GovernanceWrapperSession) GetProposals(proposalStatus int32, pageRequest CosmosPageRequest) ([]IGovernanceModuleProposal, CosmosPageResponse, error) {
	return _GovernanceWrapper.Contract.GetProposals(&_GovernanceWrapper.CallOpts, proposalStatus, pageRequest)
}

// GetProposals is a free data retrieval call binding the contract method 0x917c9d92.
//
// Solidity: function getProposals(int32 proposalStatus, (string,uint64,uint64,bool,bool) pageRequest) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string)[], (string,uint64))
func (_GovernanceWrapper *GovernanceWrapperCallerSession) GetProposals(proposalStatus int32, pageRequest CosmosPageRequest) ([]IGovernanceModuleProposal, CosmosPageResponse, error) {
	return _GovernanceWrapper.Contract.GetProposals(&_GovernanceWrapper.CallOpts, proposalStatus, pageRequest)
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
