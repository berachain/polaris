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
	Value   []byte
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_governanceModule\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"bank\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBankModule\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelProposal\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getProposal\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIGovernanceModule.Proposal\",\"components\":[{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messages\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.CodecAny[]\",\"components\":[{\"name\":\"typeURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"status\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"finalTallyResult\",\"type\":\"tuple\",\"internalType\":\"structIGovernanceModule.TallyResult\",\"components\":[{\"name\":\"yesCount\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"abstainCount\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"noCount\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"noWithVetoCount\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"submitTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"depositEndTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"totalDeposit\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"votingStartTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"votingEndTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"metadata\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"title\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"summary\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProposals\",\"inputs\":[{\"name\":\"proposalStatus\",\"type\":\"int32\",\"internalType\":\"int32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIGovernanceModule.Proposal[]\",\"components\":[{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messages\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.CodecAny[]\",\"components\":[{\"name\":\"typeURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"status\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"finalTallyResult\",\"type\":\"tuple\",\"internalType\":\"structIGovernanceModule.TallyResult\",\"components\":[{\"name\":\"yesCount\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"abstainCount\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"noCount\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"noWithVetoCount\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"submitTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"depositEndTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"totalDeposit\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"votingStartTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"votingEndTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"metadata\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"title\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"summary\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"governanceModule\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIGovernanceModule\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submit\",\"inputs\":[{\"name\":\"proposal\",\"type\":\"tuple\",\"internalType\":\"structIGovernanceModule.MsgSubmitProposal\",\"components\":[{\"name\":\"messages\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.CodecAny[]\",\"components\":[{\"name\":\"typeURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"initialDeposit\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"title\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"summary\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"expedited\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"vote\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"option\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"metadata\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a0604052734381dc2ab14285160c808659aee005d51255add773ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801562000057575f80fd5b506040516200264b3803806200264b83398181016040528101906200007d91906200018e565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620000e3576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050620001be565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f62000158826200012d565b9050919050565b6200016a816200014c565b811462000175575f80fd5b50565b5f8151905062000188816200015f565b92915050565b5f60208284031215620001a657620001a562000129565b5b5f620001b58482850162000178565b91505092915050565b608051612474620001d75f395f6104e501526124745ff3fe608060405260043610610072575f3560e01c806337a9a59e1161004d57806337a9a59e1461011357806376cdb03b14610150578063b5828df21461017a578063f1610a28146101b657610079565b8062e66c9a1461007d57806319f7a0fb146100ad5780632b0a7032146100e957610079565b3661007957005b5f80fd5b61009760048036038101906100929190610862565b6101f2565b6040516100a49190610911565b60405180910390f35b3480156100b8575f80fd5b506100d360048036038101906100ce9190610ac2565b610377565b6040516100e09190610b48565b60405180910390f35b3480156100f4575f80fd5b506100fd61041e565b60405161010a9190610bdb565b60405180910390f35b34801561011e575f80fd5b5061013960048036038101906101349190610bf4565b610441565b604051610147929190610c1f565b60405180910390f35b34801561015b575f80fd5b506101646104e3565b6040516101719190610c66565b60405180910390f35b348015610185575f80fd5b506101a0600480360381019061019b9190610c7f565b610507565b6040516101ad919061121d565b60405180910390f35b3480156101c1575f80fd5b506101dc60048036038101906101d79190610bf4565b6105bc565b6040516101e99190611376565b60405180910390f35b5f80600167ffffffffffffffff81111561020f5761020e61099e565b5b60405190808252806020026020018201604052801561024857816020015b610235610665565b81526020019060019003908161022d5790505b50905084848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815f815181106102a2576102a1611396565b5b60200260200101516020018190525082815f815181106102c5576102c4611396565b5b60200260200101515f0181815250505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638ed6982d876040518263ffffffff1660e01b815260040161032c919061196e565b6020604051808303815f875af1158015610348573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061036c91906119a2565b915050949350505050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166319f7a0fb8585856040518463ffffffff1660e01b81526004016103d593929190611a24565b6020604051808303815f875af11580156103f1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104159190611a74565b90509392505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f805f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166337a9a59e846040518263ffffffff1660e01b815260040161049b9190610911565b60408051808303815f875af11580156104b6573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104da9190611a9f565b91509150915091565b7f000000000000000000000000000000000000000000000000000000000000000081565b606061051161067e565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663917c9d9285846040518363ffffffff1660e01b815260040161056d929190611b50565b5f60405180830381865afa158015610587573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906105af9190612381565b5090508092505050919050565b6105c46106c1565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f1610a28836040518263ffffffff1660e01b815260040161061c9190610911565b5f60405180830381865afa158015610636573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061065e91906123f7565b9050919050565b60405180604001604052805f8152602001606081525090565b6040518060a00160405280606081526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f151581526020015f151581525090565b604051806101a001604052805f67ffffffffffffffff168152602001606081526020015f60030b81526020016106f5610773565b81526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff168152602001606081526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020016060815260200160608152602001606081526020015f73ffffffffffffffffffffffffffffffffffffffff1681525090565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f60e082840312156107c5576107c46107ac565b5b81905092915050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f8401126107ef576107ee6107ce565b5b8235905067ffffffffffffffff81111561080c5761080b6107d2565b5b602083019150836001820283011115610828576108276107d6565b5b9250929050565b5f819050919050565b6108418161082f565b811461084b575f80fd5b50565b5f8135905061085c81610838565b92915050565b5f805f806060858703121561087a576108796107a4565b5b5f85013567ffffffffffffffff811115610897576108966107a8565b5b6108a3878288016107b0565b945050602085013567ffffffffffffffff8111156108c4576108c36107a8565b5b6108d0878288016107da565b935093505060406108e38782880161084e565b91505092959194509250565b5f67ffffffffffffffff82169050919050565b61090b816108ef565b82525050565b5f6020820190506109245f830184610902565b92915050565b610933816108ef565b811461093d575f80fd5b50565b5f8135905061094e8161092a565b92915050565b5f8160030b9050919050565b61096981610954565b8114610973575f80fd5b50565b5f8135905061098481610960565b92915050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6109d48261098e565b810181811067ffffffffffffffff821117156109f3576109f261099e565b5b80604052505050565b5f610a0561079b565b9050610a1182826109cb565b919050565b5f67ffffffffffffffff821115610a3057610a2f61099e565b5b610a398261098e565b9050602081019050919050565b828183375f83830152505050565b5f610a66610a6184610a16565b6109fc565b905082815260208101848484011115610a8257610a8161098a565b5b610a8d848285610a46565b509392505050565b5f82601f830112610aa957610aa86107ce565b5b8135610ab9848260208601610a54565b91505092915050565b5f805f60608486031215610ad957610ad86107a4565b5b5f610ae686828701610940565b9350506020610af786828701610976565b925050604084013567ffffffffffffffff811115610b1857610b176107a8565b5b610b2486828701610a95565b9150509250925092565b5f8115159050919050565b610b4281610b2e565b82525050565b5f602082019050610b5b5f830184610b39565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f610ba3610b9e610b9984610b61565b610b80565b610b61565b9050919050565b5f610bb482610b89565b9050919050565b5f610bc582610baa565b9050919050565b610bd581610bbb565b82525050565b5f602082019050610bee5f830184610bcc565b92915050565b5f60208284031215610c0957610c086107a4565b5b5f610c1684828501610940565b91505092915050565b5f604082019050610c325f830185610902565b610c3f6020830184610902565b9392505050565b5f610c5082610baa565b9050919050565b610c6081610c46565b82525050565b5f602082019050610c795f830184610c57565b92915050565b5f60208284031215610c9457610c936107a4565b5b5f610ca184828501610976565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610cdc816108ef565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610d42578082015181840152602081019050610d27565b5f8484015250505050565b5f610d5782610d0b565b610d618185610d15565b9350610d71818560208601610d25565b610d7a8161098e565b840191505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f610da982610d85565b610db38185610d8f565b9350610dc3818560208601610d25565b610dcc8161098e565b840191505092915050565b5f604083015f8301518482035f860152610df18282610d4d565b91505060208301518482036020860152610e0b8282610d9f565b9150508091505092915050565b5f610e238383610dd7565b905092915050565b5f602082019050919050565b5f610e4182610ce2565b610e4b8185610cec565b935083602082028501610e5d85610cfc565b805f5b85811015610e985784840389528151610e798582610e18565b9450610e8483610e2b565b925060208a01995050600181019050610e60565b50829750879550505050505092915050565b610eb381610954565b82525050565b5f608083015f8301518482035f860152610ed38282610d4d565b91505060208301518482036020860152610eed8282610d4d565b91505060408301518482036040860152610f078282610d4d565b91505060608301518482036060860152610f218282610d4d565b9150508091505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610f608161082f565b82525050565b5f604083015f830151610f7b5f860182610f57565b5060208301518482036020860152610f938282610d4d565b9150508091505092915050565b5f610fab8383610f66565b905092915050565b5f602082019050919050565b5f610fc982610f2e565b610fd38185610f38565b935083602082028501610fe585610f48565b805f5b8581101561102057848403895281516110018582610fa0565b945061100c83610fb3565b925060208a01995050600181019050610fe8565b50829750879550505050505092915050565b5f61103c82610b61565b9050919050565b61104c81611032565b82525050565b5f6101a083015f8301516110685f860182610cd3565b50602083015184820360208601526110808282610e37565b91505060408301516110956040860182610eaa565b50606083015184820360608601526110ad8282610eb9565b91505060808301516110c26080860182610cd3565b5060a08301516110d560a0860182610cd3565b5060c083015184820360c08601526110ed8282610fbf565b91505060e083015161110260e0860182610cd3565b50610100830151611117610100860182610cd3565b506101208301518482036101208601526111318282610d4d565b91505061014083015184820361014086015261114d8282610d4d565b9150506101608301518482036101608601526111698282610d4d565b915050610180830151611180610180860182611043565b508091505092915050565b5f6111968383611052565b905092915050565b5f602082019050919050565b5f6111b482610caa565b6111be8185610cb4565b9350836020820285016111d085610cc4565b805f5b8581101561120b57848403895281516111ec858261118b565b94506111f78361119e565b925060208a019950506001810190506111d3565b50829750879550505050505092915050565b5f6020820190508181035f83015261123581846111aa565b905092915050565b5f6101a083015f8301516112535f860182610cd3565b506020830151848203602086015261126b8282610e37565b91505060408301516112806040860182610eaa565b50606083015184820360608601526112988282610eb9565b91505060808301516112ad6080860182610cd3565b5060a08301516112c060a0860182610cd3565b5060c083015184820360c08601526112d88282610fbf565b91505060e08301516112ed60e0860182610cd3565b50610100830151611302610100860182610cd3565b5061012083015184820361012086015261131c8282610d4d565b9150506101408301518482036101408601526113388282610d4d565b9150506101608301518482036101608601526113548282610d4d565b91505061018083015161136b610180860182611043565b508091505092915050565b5f6020820190508181035f83015261138e818461123d565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f80fd5b5f80fd5b5f80fd5b5f80833560016020038436030381126113eb576113ea6113cb565b5b83810192508235915060208301925067ffffffffffffffff821115611413576114126113c3565b5b602082023603831315611429576114286113c7565b5b509250929050565b5f819050919050565b5f8083356001602003843603038112611456576114556113cb565b5b83810192508235915060208301925067ffffffffffffffff82111561147e5761147d6113c3565b5b600182023603831315611494576114936113c7565b5b509250929050565b5f6114a78385610d15565b93506114b4838584610a46565b6114bd8361098e565b840190509392505050565b5f80833560016020038436030381126114e4576114e36113cb565b5b83810192508235915060208301925067ffffffffffffffff82111561150c5761150b6113c3565b5b600182023603831315611522576115216113c7565b5b509250929050565b5f6115358385610d8f565b9350611542838584610a46565b61154b8361098e565b840190509392505050565b5f604083016115675f84018461143a565b8583035f87015261157983828461149c565b9250505061158a60208401846114c8565b858303602087015261159d83828461152a565b925050508091505092915050565b5f6115b68383611556565b905092915050565b5f823560016040038336030381126115d9576115d86113cb565b5b82810191505092915050565b5f602082019050919050565b5f6115fc8385610cec565b93508360208402850161160e84611431565b805f5b8781101561165157848403895261162882846115be565b61163285826115ab565b945061163d836115e5565b925060208a01995050600181019050611611565b50829750879450505050509392505050565b5f808335600160200384360303811261167f5761167e6113cb565b5b83810192508235915060208301925067ffffffffffffffff8211156116a7576116a66113c3565b5b6020820236038313156116bd576116bc6113c7565b5b509250929050565b5f819050919050565b5f6116dc602084018461084e565b905092915050565b5f604083016116f55f8401846116ce565b6117015f860182610f57565b5061170f602084018461143a565b858303602087015261172283828461149c565b925050508091505092915050565b5f61173b83836116e4565b905092915050565b5f8235600160400383360303811261175e5761175d6113cb565b5b82810191505092915050565b5f602082019050919050565b5f6117818385610f38565b935083602084028501611793846116c5565b805f5b878110156117d65784840389526117ad8284611743565b6117b78582611730565b94506117c28361176a565b925060208a01995050600181019050611796565b50829750879450505050509392505050565b6117f181611032565b81146117fb575f80fd5b50565b5f8135905061180c816117e8565b92915050565b5f61182060208401846117fe565b905092915050565b61183181610b2e565b811461183b575f80fd5b50565b5f8135905061184c81611828565b92915050565b5f611860602084018461183e565b905092915050565b61187181610b2e565b82525050565b5f60e083016118885f8401846113cf565b8583035f87015261189a8382846115f1565b925050506118ab6020840184611663565b85830360208701526118be838284611776565b925050506118cf6040840184611812565b6118dc6040860182611043565b506118ea606084018461143a565b85830360608701526118fd83828461149c565b9250505061190e608084018461143a565b858303608087015261192183828461149c565b9250505061193260a084018461143a565b85830360a087015261194583828461149c565b9250505061195660c0840184611852565b61196360c0860182611868565b508091505092915050565b5f6020820190508181035f8301526119868184611877565b905092915050565b5f8151905061199c8161092a565b92915050565b5f602082840312156119b7576119b66107a4565b5b5f6119c48482850161198e565b91505092915050565b6119d681610954565b82525050565b5f82825260208201905092915050565b5f6119f682610d0b565b611a0081856119dc565b9350611a10818560208601610d25565b611a198161098e565b840191505092915050565b5f606082019050611a375f830186610902565b611a4460208301856119cd565b8181036040830152611a5681846119ec565b9050949350505050565b5f81519050611a6e81611828565b92915050565b5f60208284031215611a8957611a886107a4565b5b5f611a9684828501611a60565b91505092915050565b5f8060408385031215611ab557611ab46107a4565b5b5f611ac28582860161198e565b9250506020611ad38582860161198e565b9150509250929050565b5f60a083015f8301518482035f860152611af78282610d4d565b9150506020830151611b0c6020860182610cd3565b506040830151611b1f6040860182610cd3565b506060830151611b326060860182611868565b506080830151611b456080860182611868565b508091505092915050565b5f604082019050611b635f8301856119cd565b8181036020830152611b758184611add565b90509392505050565b5f67ffffffffffffffff821115611b9857611b9761099e565b5b602082029050602081019050919050565b5f80fd5b5f80fd5b5f67ffffffffffffffff821115611bcb57611bca61099e565b5b602082029050602081019050919050565b5f611bee611be984610a16565b6109fc565b905082815260208101848484011115611c0a57611c0961098a565b5b611c15848285610d25565b509392505050565b5f82601f830112611c3157611c306107ce565b5b8151611c41848260208601611bdc565b91505092915050565b5f67ffffffffffffffff821115611c6457611c6361099e565b5b611c6d8261098e565b9050602081019050919050565b5f611c8c611c8784611c4a565b6109fc565b905082815260208101848484011115611ca857611ca761098a565b5b611cb3848285610d25565b509392505050565b5f82601f830112611ccf57611cce6107ce565b5b8151611cdf848260208601611c7a565b91505092915050565b5f60408284031215611cfd57611cfc611ba9565b5b611d0760406109fc565b90505f82015167ffffffffffffffff811115611d2657611d25611bad565b5b611d3284828501611c1d565b5f83015250602082015167ffffffffffffffff811115611d5557611d54611bad565b5b611d6184828501611cbb565b60208301525092915050565b5f611d7f611d7a84611bb1565b6109fc565b90508083825260208201905060208402830185811115611da257611da16107d6565b5b835b81811015611de957805167ffffffffffffffff811115611dc757611dc66107ce565b5b808601611dd48982611ce8565b85526020850194505050602081019050611da4565b5050509392505050565b5f82601f830112611e0757611e066107ce565b5b8151611e17848260208601611d6d565b91505092915050565b5f81519050611e2e81610960565b92915050565b5f60808284031215611e4957611e48611ba9565b5b611e5360806109fc565b90505f82015167ffffffffffffffff811115611e7257611e71611bad565b5b611e7e84828501611c1d565b5f83015250602082015167ffffffffffffffff811115611ea157611ea0611bad565b5b611ead84828501611c1d565b602083015250604082015167ffffffffffffffff811115611ed157611ed0611bad565b5b611edd84828501611c1d565b604083015250606082015167ffffffffffffffff811115611f0157611f00611bad565b5b611f0d84828501611c1d565b60608301525092915050565b5f67ffffffffffffffff821115611f3357611f3261099e565b5b602082029050602081019050919050565b5f81519050611f5281610838565b92915050565b5f60408284031215611f6d57611f6c611ba9565b5b611f7760406109fc565b90505f611f8684828501611f44565b5f83015250602082015167ffffffffffffffff811115611fa957611fa8611bad565b5b611fb584828501611c1d565b60208301525092915050565b5f611fd3611fce84611f19565b6109fc565b90508083825260208201905060208402830185811115611ff657611ff56107d6565b5b835b8181101561203d57805167ffffffffffffffff81111561201b5761201a6107ce565b5b8086016120288982611f58565b85526020850194505050602081019050611ff8565b5050509392505050565b5f82601f83011261205b5761205a6107ce565b5b815161206b848260208601611fc1565b91505092915050565b5f81519050612082816117e8565b92915050565b5f6101a0828403121561209e5761209d611ba9565b5b6120a96101a06109fc565b90505f6120b88482850161198e565b5f83015250602082015167ffffffffffffffff8111156120db576120da611bad565b5b6120e784828501611df3565b60208301525060406120fb84828501611e20565b604083015250606082015167ffffffffffffffff81111561211f5761211e611bad565b5b61212b84828501611e34565b606083015250608061213f8482850161198e565b60808301525060a06121538482850161198e565b60a08301525060c082015167ffffffffffffffff81111561217757612176611bad565b5b61218384828501612047565b60c08301525060e06121978482850161198e565b60e0830152506101006121ac8482850161198e565b6101008301525061012082015167ffffffffffffffff8111156121d2576121d1611bad565b5b6121de84828501611c1d565b6101208301525061014082015167ffffffffffffffff81111561220457612203611bad565b5b61221084828501611c1d565b6101408301525061016082015167ffffffffffffffff81111561223657612235611bad565b5b61224284828501611c1d565b6101608301525061018061225884828501612074565b6101808301525092915050565b5f61227761227284611b7e565b6109fc565b9050808382526020820190506020840283018581111561229a576122996107d6565b5b835b818110156122e157805167ffffffffffffffff8111156122bf576122be6107ce565b5b8086016122cc8982612088565b8552602085019450505060208101905061229c565b5050509392505050565b5f82601f8301126122ff576122fe6107ce565b5b815161230f848260208601612265565b91505092915050565b5f6040828403121561232d5761232c611ba9565b5b61233760406109fc565b90505f82015167ffffffffffffffff81111561235657612355611bad565b5b61236284828501611c1d565b5f8301525060206123758482850161198e565b60208301525092915050565b5f8060408385031215612397576123966107a4565b5b5f83015167ffffffffffffffff8111156123b4576123b36107a8565b5b6123c0858286016122eb565b925050602083015167ffffffffffffffff8111156123e1576123e06107a8565b5b6123ed85828601612318565b9150509250929050565b5f6020828403121561240c5761240b6107a4565b5b5f82015167ffffffffffffffff811115612429576124286107a8565b5b61243584828501612088565b9150509291505056fea2646970667358221220f86c11bbafad99c494796136cccc24a01940ebe62573abc6f290ab431a019e7264736f6c63430008170033",
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
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,(string,bytes)[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,address))
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
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,(string,bytes)[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,address))
func (_GovernanceWrapper *GovernanceWrapperSession) GetProposal(proposalId uint64) (IGovernanceModuleProposal, error) {
	return _GovernanceWrapper.Contract.GetProposal(&_GovernanceWrapper.CallOpts, proposalId)
}

// GetProposal is a free data retrieval call binding the contract method 0xf1610a28.
//
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,(string,bytes)[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,address))
func (_GovernanceWrapper *GovernanceWrapperCallerSession) GetProposal(proposalId uint64) (IGovernanceModuleProposal, error) {
	return _GovernanceWrapper.Contract.GetProposal(&_GovernanceWrapper.CallOpts, proposalId)
}

// GetProposals is a free data retrieval call binding the contract method 0xb5828df2.
//
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,(string,bytes)[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,address)[])
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
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,(string,bytes)[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,address)[])
func (_GovernanceWrapper *GovernanceWrapperSession) GetProposals(proposalStatus int32) ([]IGovernanceModuleProposal, error) {
	return _GovernanceWrapper.Contract.GetProposals(&_GovernanceWrapper.CallOpts, proposalStatus)
}

// GetProposals is a free data retrieval call binding the contract method 0xb5828df2.
//
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,(string,bytes)[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,address)[])
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

// Submit is a paid mutator transaction binding the contract method 0x00e66c9a.
//
// Solidity: function submit(((string,bytes)[],(uint256,string)[],address,string,string,string,bool) proposal, string denom, uint256 amount) payable returns(uint64)
func (_GovernanceWrapper *GovernanceWrapperTransactor) Submit(opts *bind.TransactOpts, proposal IGovernanceModuleMsgSubmitProposal, denom string, amount *big.Int) (*types.Transaction, error) {
	return _GovernanceWrapper.contract.Transact(opts, "submit", proposal, denom, amount)
}

// Submit is a paid mutator transaction binding the contract method 0x00e66c9a.
//
// Solidity: function submit(((string,bytes)[],(uint256,string)[],address,string,string,string,bool) proposal, string denom, uint256 amount) payable returns(uint64)
func (_GovernanceWrapper *GovernanceWrapperSession) Submit(proposal IGovernanceModuleMsgSubmitProposal, denom string, amount *big.Int) (*types.Transaction, error) {
	return _GovernanceWrapper.Contract.Submit(&_GovernanceWrapper.TransactOpts, proposal, denom, amount)
}

// Submit is a paid mutator transaction binding the contract method 0x00e66c9a.
//
// Solidity: function submit(((string,bytes)[],(uint256,string)[],address,string,string,string,bool) proposal, string denom, uint256 amount) payable returns(uint64)
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
