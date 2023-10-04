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
	TypeUrl string
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
	Proposer       string
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governanceModule\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"contractIBankModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"cancelProposal\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"typeUrl\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structCosmos.CodecAny[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.TallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"submitTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositEndTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structCosmos.Coin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"votingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposer\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Proposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"proposalStatus\",\"type\":\"int32\"}],\"name\":\"getProposals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"typeUrl\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structCosmos.CodecAny[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.TallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"submitTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositEndTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structCosmos.Coin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"votingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposer\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Proposal[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceModule\",\"outputs\":[{\"internalType\":\"contractIGovernanceModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"typeUrl\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structCosmos.CodecAny[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structCosmos.Coin[]\",\"name\":\"initialDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"proposer\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"expedited\",\"type\":\"bool\"}],\"internalType\":\"structIGovernanceModule.MsgSubmitProposal\",\"name\":\"proposal\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"submit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"internalType\":\"int32\",\"name\":\"option\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a0604052734381dc2ab14285160c808659aee005d51255add773ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801562000057575f80fd5b50604051620025f6380380620025f683398181016040528101906200007d91906200018e565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620000e3576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050620001be565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f62000158826200012d565b9050919050565b6200016a816200014c565b811462000175575f80fd5b50565b5f8151905062000188816200015f565b92915050565b5f60208284031215620001a657620001a562000129565b5b5f620001b58482850162000178565b91505092915050565b60805161241f620001d75f395f610361015261241f5ff3fe608060405260043610610073575f3560e01c806376cdb03b1161004d57806376cdb03b14610121578063b5828df21461014b578063ebce138514610187578063f1610a28146101b75761007a565b806319f7a0fb1461007e5780632b0a7032146100ba57806337a9a59e146100e45761007a565b3661007a57005b5f80fd5b348015610089575f80fd5b506100a4600480360381019061009f9190610947565b6101f3565b6040516100b191906109cd565b60405180910390f35b3480156100c5575f80fd5b506100ce61029a565b6040516100db9190610a60565b60405180910390f35b3480156100ef575f80fd5b5061010a60048036038101906101059190610a79565b6102bd565b604051610118929190610ab3565b60405180910390f35b34801561012c575f80fd5b5061013561035f565b6040516101429190610afa565b60405180910390f35b348015610156575f80fd5b50610171600480360381019061016c9190610b13565b610383565b60405161017e91906110a1565b60405180910390f35b6101a1600480360381019061019c919061116a565b610438565b6040516101ae91906111f7565b60405180910390f35b3480156101c2575f80fd5b506101dd60048036038101906101d89190610a79565b6105bd565b6040516101ea9190611350565b60405180910390f35b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166319f7a0fb8585856040518463ffffffff1660e01b8152600401610251939291906113c7565b6020604051808303815f875af115801561026d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610291919061142d565b90509392505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f805f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166337a9a59e846040518263ffffffff1660e01b815260040161031791906111f7565b60408051808303815f875af1158015610332573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610356919061146c565b91509150915091565b7f000000000000000000000000000000000000000000000000000000000000000081565b606061038d610666565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663917c9d9285846040518363ffffffff1660e01b81526004016103e992919061152c565b5f60405180830381865afa158015610403573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061042b9190611d65565b5090508092505050919050565b5f80600167ffffffffffffffff81111561045557610454610823565b5b60405190808252806020026020018201604052801561048e57816020015b61047b6106a9565b8152602001906001900390816104735790505b50905084848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815f815181106104e8576104e7611ddb565b5b60200260200101516020018190525082815f8151811061050b5761050a611ddb565b5b60200260200101515f0181815250505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639e4916bc876040518263ffffffff1660e01b81526004016105729190612357565b6020604051808303815f875af115801561058e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105b29190612377565b915050949350505050565b6105c56106c2565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f1610a28836040518263ffffffff1660e01b815260040161061d91906111f7565b5f60405180830381865afa158015610637573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061065f91906123a2565b9050919050565b6040518060a00160405280606081526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f151581526020015f151581525090565b60405180604001604052805f8152602001606081525090565b604051806101a001604052805f67ffffffffffffffff168152602001606081526020015f60030b81526020016106f661075f565b81526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff168152602001606081526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff168152602001606081526020016060815260200160608152602001606081525090565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b5f604051905090565b5f80fd5b5f80fd5b5f67ffffffffffffffff82169050919050565b6107b481610798565b81146107be575f80fd5b50565b5f813590506107cf816107ab565b92915050565b5f8160030b9050919050565b6107ea816107d5565b81146107f4575f80fd5b50565b5f81359050610805816107e1565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61085982610813565b810181811067ffffffffffffffff8211171561087857610877610823565b5b80604052505050565b5f61088a610787565b90506108968282610850565b919050565b5f67ffffffffffffffff8211156108b5576108b4610823565b5b6108be82610813565b9050602081019050919050565b828183375f83830152505050565b5f6108eb6108e68461089b565b610881565b9050828152602081018484840111156109075761090661080f565b5b6109128482856108cb565b509392505050565b5f82601f83011261092e5761092d61080b565b5b813561093e8482602086016108d9565b91505092915050565b5f805f6060848603121561095e5761095d610790565b5b5f61096b868287016107c1565b935050602061097c868287016107f7565b925050604084013567ffffffffffffffff81111561099d5761099c610794565b5b6109a98682870161091a565b9150509250925092565b5f8115159050919050565b6109c7816109b3565b82525050565b5f6020820190506109e05f8301846109be565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f610a28610a23610a1e846109e6565b610a05565b6109e6565b9050919050565b5f610a3982610a0e565b9050919050565b5f610a4a82610a2f565b9050919050565b610a5a81610a40565b82525050565b5f602082019050610a735f830184610a51565b92915050565b5f60208284031215610a8e57610a8d610790565b5b5f610a9b848285016107c1565b91505092915050565b610aad81610798565b82525050565b5f604082019050610ac65f830185610aa4565b610ad36020830184610aa4565b9392505050565b5f610ae482610a2f565b9050919050565b610af481610ada565b82525050565b5f602082019050610b0d5f830184610aeb565b92915050565b5f60208284031215610b2857610b27610790565b5b5f610b35848285016107f7565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610b7081610798565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610bd6578082015181840152602081019050610bbb565b5f8484015250505050565b5f610beb82610b9f565b610bf58185610ba9565b9350610c05818560208601610bb9565b610c0e81610813565b840191505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f610c3d82610c19565b610c478185610c23565b9350610c57818560208601610bb9565b610c6081610813565b840191505092915050565b5f604083015f8301518482035f860152610c858282610be1565b91505060208301518482036020860152610c9f8282610c33565b9150508091505092915050565b5f610cb78383610c6b565b905092915050565b5f602082019050919050565b5f610cd582610b76565b610cdf8185610b80565b935083602082028501610cf185610b90565b805f5b85811015610d2c5784840389528151610d0d8582610cac565b9450610d1883610cbf565b925060208a01995050600181019050610cf4565b50829750879550505050505092915050565b610d47816107d5565b82525050565b5f608083015f8301518482035f860152610d678282610be1565b91505060208301518482036020860152610d818282610be1565b91505060408301518482036040860152610d9b8282610be1565b91505060608301518482036060860152610db58282610be1565b9150508091505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f819050919050565b610dfd81610deb565b82525050565b5f604083015f830151610e185f860182610df4565b5060208301518482036020860152610e308282610be1565b9150508091505092915050565b5f610e488383610e03565b905092915050565b5f602082019050919050565b5f610e6682610dc2565b610e708185610dcc565b935083602082028501610e8285610ddc565b805f5b85811015610ebd5784840389528151610e9e8582610e3d565b9450610ea983610e50565b925060208a01995050600181019050610e85565b50829750879550505050505092915050565b5f6101a083015f830151610ee55f860182610b67565b5060208301518482036020860152610efd8282610ccb565b9150506040830151610f126040860182610d3e565b5060608301518482036060860152610f2a8282610d4d565b9150506080830151610f3f6080860182610b67565b5060a0830151610f5260a0860182610b67565b5060c083015184820360c0860152610f6a8282610e5c565b91505060e0830151610f7f60e0860182610b67565b50610100830151610f94610100860182610b67565b50610120830151848203610120860152610fae8282610be1565b915050610140830151848203610140860152610fca8282610be1565b915050610160830151848203610160860152610fe68282610be1565b9150506101808301518482036101808601526110028282610be1565b9150508091505092915050565b5f61101a8383610ecf565b905092915050565b5f602082019050919050565b5f61103882610b3e565b6110428185610b48565b93508360208202850161105485610b58565b805f5b8581101561108f5784840389528151611070858261100f565b945061107b83611022565b925060208a01995050600181019050611057565b50829750879550505050505092915050565b5f6020820190508181035f8301526110b9818461102e565b905092915050565b5f80fd5b5f60e082840312156110da576110d96110c1565b5b81905092915050565b5f80fd5b5f80fd5b5f8083601f840112611100576110ff61080b565b5b8235905067ffffffffffffffff81111561111d5761111c6110e3565b5b602083019150836001820283011115611139576111386110e7565b5b9250929050565b61114981610deb565b8114611153575f80fd5b50565b5f8135905061116481611140565b92915050565b5f805f806060858703121561118257611181610790565b5b5f85013567ffffffffffffffff81111561119f5761119e610794565b5b6111ab878288016110c5565b945050602085013567ffffffffffffffff8111156111cc576111cb610794565b5b6111d8878288016110eb565b935093505060406111eb87828801611156565b91505092959194509250565b5f60208201905061120a5f830184610aa4565b92915050565b5f6101a083015f8301516112265f860182610b67565b506020830151848203602086015261123e8282610ccb565b91505060408301516112536040860182610d3e565b506060830151848203606086015261126b8282610d4d565b91505060808301516112806080860182610b67565b5060a083015161129360a0860182610b67565b5060c083015184820360c08601526112ab8282610e5c565b91505060e08301516112c060e0860182610b67565b506101008301516112d5610100860182610b67565b506101208301518482036101208601526112ef8282610be1565b91505061014083015184820361014086015261130b8282610be1565b9150506101608301518482036101608601526113278282610be1565b9150506101808301518482036101808601526113438282610be1565b9150508091505092915050565b5f6020820190508181035f8301526113688184611210565b905092915050565b611379816107d5565b82525050565b5f82825260208201905092915050565b5f61139982610b9f565b6113a3818561137f565b93506113b3818560208601610bb9565b6113bc81610813565b840191505092915050565b5f6060820190506113da5f830186610aa4565b6113e76020830185611370565b81810360408301526113f9818461138f565b9050949350505050565b61140c816109b3565b8114611416575f80fd5b50565b5f8151905061142781611403565b92915050565b5f6020828403121561144257611441610790565b5b5f61144f84828501611419565b91505092915050565b5f81519050611466816107ab565b92915050565b5f806040838503121561148257611481610790565b5b5f61148f85828601611458565b92505060206114a085828601611458565b9150509250929050565b6114b3816109b3565b82525050565b5f60a083015f8301518482035f8601526114d38282610be1565b91505060208301516114e86020860182610b67565b5060408301516114fb6040860182610b67565b50606083015161150e60608601826114aa565b50608083015161152160808601826114aa565b508091505092915050565b5f60408201905061153f5f830185611370565b818103602083015261155181846114b9565b90509392505050565b5f67ffffffffffffffff82111561157457611573610823565b5b602082029050602081019050919050565b5f80fd5b5f80fd5b5f67ffffffffffffffff8211156115a7576115a6610823565b5b602082029050602081019050919050565b5f6115ca6115c58461089b565b610881565b9050828152602081018484840111156115e6576115e561080f565b5b6115f1848285610bb9565b509392505050565b5f82601f83011261160d5761160c61080b565b5b815161161d8482602086016115b8565b91505092915050565b5f67ffffffffffffffff8211156116405761163f610823565b5b61164982610813565b9050602081019050919050565b5f61166861166384611626565b610881565b9050828152602081018484840111156116845761168361080f565b5b61168f848285610bb9565b509392505050565b5f82601f8301126116ab576116aa61080b565b5b81516116bb848260208601611656565b91505092915050565b5f604082840312156116d9576116d8611585565b5b6116e36040610881565b90505f82015167ffffffffffffffff81111561170257611701611589565b5b61170e848285016115f9565b5f83015250602082015167ffffffffffffffff81111561173157611730611589565b5b61173d84828501611697565b60208301525092915050565b5f61175b6117568461158d565b610881565b9050808382526020820190506020840283018581111561177e5761177d6110e7565b5b835b818110156117c557805167ffffffffffffffff8111156117a3576117a261080b565b5b8086016117b089826116c4565b85526020850194505050602081019050611780565b5050509392505050565b5f82601f8301126117e3576117e261080b565b5b81516117f3848260208601611749565b91505092915050565b5f8151905061180a816107e1565b92915050565b5f6080828403121561182557611824611585565b5b61182f6080610881565b90505f82015167ffffffffffffffff81111561184e5761184d611589565b5b61185a848285016115f9565b5f83015250602082015167ffffffffffffffff81111561187d5761187c611589565b5b611889848285016115f9565b602083015250604082015167ffffffffffffffff8111156118ad576118ac611589565b5b6118b9848285016115f9565b604083015250606082015167ffffffffffffffff8111156118dd576118dc611589565b5b6118e9848285016115f9565b60608301525092915050565b5f67ffffffffffffffff82111561190f5761190e610823565b5b602082029050602081019050919050565b5f8151905061192e81611140565b92915050565b5f6040828403121561194957611948611585565b5b6119536040610881565b90505f61196284828501611920565b5f83015250602082015167ffffffffffffffff81111561198557611984611589565b5b611991848285016115f9565b60208301525092915050565b5f6119af6119aa846118f5565b610881565b905080838252602082019050602084028301858111156119d2576119d16110e7565b5b835b81811015611a1957805167ffffffffffffffff8111156119f7576119f661080b565b5b808601611a048982611934565b855260208501945050506020810190506119d4565b5050509392505050565b5f82601f830112611a3757611a3661080b565b5b8151611a4784826020860161199d565b91505092915050565b5f6101a08284031215611a6657611a65611585565b5b611a716101a0610881565b90505f611a8084828501611458565b5f83015250602082015167ffffffffffffffff811115611aa357611aa2611589565b5b611aaf848285016117cf565b6020830152506040611ac3848285016117fc565b604083015250606082015167ffffffffffffffff811115611ae757611ae6611589565b5b611af384828501611810565b6060830152506080611b0784828501611458565b60808301525060a0611b1b84828501611458565b60a08301525060c082015167ffffffffffffffff811115611b3f57611b3e611589565b5b611b4b84828501611a23565b60c08301525060e0611b5f84828501611458565b60e083015250610100611b7484828501611458565b6101008301525061012082015167ffffffffffffffff811115611b9a57611b99611589565b5b611ba6848285016115f9565b6101208301525061014082015167ffffffffffffffff811115611bcc57611bcb611589565b5b611bd8848285016115f9565b6101408301525061016082015167ffffffffffffffff811115611bfe57611bfd611589565b5b611c0a848285016115f9565b6101608301525061018082015167ffffffffffffffff811115611c3057611c2f611589565b5b611c3c848285016115f9565b6101808301525092915050565b5f611c5b611c568461155a565b610881565b90508083825260208201905060208402830185811115611c7e57611c7d6110e7565b5b835b81811015611cc557805167ffffffffffffffff811115611ca357611ca261080b565b5b808601611cb08982611a50565b85526020850194505050602081019050611c80565b5050509392505050565b5f82601f830112611ce357611ce261080b565b5b8151611cf3848260208601611c49565b91505092915050565b5f60408284031215611d1157611d10611585565b5b611d1b6040610881565b90505f82015167ffffffffffffffff811115611d3a57611d39611589565b5b611d46848285016115f9565b5f830152506020611d5984828501611458565b60208301525092915050565b5f8060408385031215611d7b57611d7a610790565b5b5f83015167ffffffffffffffff811115611d9857611d97610794565b5b611da485828601611ccf565b925050602083015167ffffffffffffffff811115611dc557611dc4610794565b5b611dd185828601611cfc565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f80fd5b5f80fd5b5f80fd5b5f8083356001602003843603038112611e3057611e2f611e10565b5b83810192508235915060208301925067ffffffffffffffff821115611e5857611e57611e08565b5b602082023603831315611e6e57611e6d611e0c565b5b509250929050565b5f819050919050565b5f8083356001602003843603038112611e9b57611e9a611e10565b5b83810192508235915060208301925067ffffffffffffffff821115611ec357611ec2611e08565b5b600182023603831315611ed957611ed8611e0c565b5b509250929050565b5f611eec8385610ba9565b9350611ef98385846108cb565b611f0283610813565b840190509392505050565b5f8083356001602003843603038112611f2957611f28611e10565b5b83810192508235915060208301925067ffffffffffffffff821115611f5157611f50611e08565b5b600182023603831315611f6757611f66611e0c565b5b509250929050565b5f611f7a8385610c23565b9350611f878385846108cb565b611f9083610813565b840190509392505050565b5f60408301611fac5f840184611e7f565b8583035f870152611fbe838284611ee1565b92505050611fcf6020840184611f0d565b8583036020870152611fe2838284611f6f565b925050508091505092915050565b5f611ffb8383611f9b565b905092915050565b5f8235600160400383360303811261201e5761201d611e10565b5b82810191505092915050565b5f602082019050919050565b5f6120418385610b80565b93508360208402850161205384611e76565b805f5b8781101561209657848403895261206d8284612003565b6120778582611ff0565b94506120828361202a565b925060208a01995050600181019050612056565b50829750879450505050509392505050565b5f80833560016020038436030381126120c4576120c3611e10565b5b83810192508235915060208301925067ffffffffffffffff8211156120ec576120eb611e08565b5b60208202360383131561210257612101611e0c565b5b509250929050565b5f819050919050565b5f6121216020840184611156565b905092915050565b5f6040830161213a5f840184612113565b6121465f860182610df4565b506121546020840184611e7f565b8583036020870152612167838284611ee1565b925050508091505092915050565b5f6121808383612129565b905092915050565b5f823560016040038336030381126121a3576121a2611e10565b5b82810191505092915050565b5f602082019050919050565b5f6121c68385610dcc565b9350836020840285016121d88461210a565b805f5b8781101561221b5784840389526121f28284612188565b6121fc8582612175565b9450612207836121af565b925060208a019950506001810190506121db565b50829750879450505050509392505050565b5f8135905061223b81611403565b92915050565b5f61224f602084018461222d565b905092915050565b5f60e083016122685f840184611e14565b8583035f87015261227a838284612036565b9250505061228b60208401846120a8565b858303602087015261229e8382846121bb565b925050506122af6040840184611e7f565b85830360408701526122c2838284611ee1565b925050506122d36060840184611e7f565b85830360608701526122e6838284611ee1565b925050506122f76080840184611e7f565b858303608087015261230a838284611ee1565b9250505061231b60a0840184611e7f565b85830360a087015261232e838284611ee1565b9250505061233f60c0840184612241565b61234c60c08601826114aa565b508091505092915050565b5f6020820190508181035f83015261236f8184612257565b905092915050565b5f6020828403121561238c5761238b610790565b5b5f61239984828501611458565b91505092915050565b5f602082840312156123b7576123b6610790565b5b5f82015167ffffffffffffffff8111156123d4576123d3610794565b5b6123e084828501611a50565b9150509291505056fea2646970667358221220aa3d67bd636022b8e1329a0d80e80158692768ef5bd204a49a3c2ab6d598c73264736f6c63430008150033",
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
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,(string,bytes)[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string))
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
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,(string,bytes)[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string))
func (_GovernanceWrapper *GovernanceWrapperSession) GetProposal(proposalId uint64) (IGovernanceModuleProposal, error) {
	return _GovernanceWrapper.Contract.GetProposal(&_GovernanceWrapper.CallOpts, proposalId)
}

// GetProposal is a free data retrieval call binding the contract method 0xf1610a28.
//
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,(string,bytes)[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string))
func (_GovernanceWrapper *GovernanceWrapperCallerSession) GetProposal(proposalId uint64) (IGovernanceModuleProposal, error) {
	return _GovernanceWrapper.Contract.GetProposal(&_GovernanceWrapper.CallOpts, proposalId)
}

// GetProposals is a free data retrieval call binding the contract method 0xb5828df2.
//
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,(string,bytes)[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string)[])
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
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,(string,bytes)[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string)[])
func (_GovernanceWrapper *GovernanceWrapperSession) GetProposals(proposalStatus int32) ([]IGovernanceModuleProposal, error) {
	return _GovernanceWrapper.Contract.GetProposals(&_GovernanceWrapper.CallOpts, proposalStatus)
}

// GetProposals is a free data retrieval call binding the contract method 0xb5828df2.
//
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,(string,bytes)[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string)[])
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

// Submit is a paid mutator transaction binding the contract method 0xebce1385.
//
// Solidity: function submit(((string,bytes)[],(uint256,string)[],string,string,string,string,bool) proposal, string denom, uint256 amount) payable returns(uint64)
func (_GovernanceWrapper *GovernanceWrapperTransactor) Submit(opts *bind.TransactOpts, proposal IGovernanceModuleMsgSubmitProposal, denom string, amount *big.Int) (*types.Transaction, error) {
	return _GovernanceWrapper.contract.Transact(opts, "submit", proposal, denom, amount)
}

// Submit is a paid mutator transaction binding the contract method 0xebce1385.
//
// Solidity: function submit(((string,bytes)[],(uint256,string)[],string,string,string,string,bool) proposal, string denom, uint256 amount) payable returns(uint64)
func (_GovernanceWrapper *GovernanceWrapperSession) Submit(proposal IGovernanceModuleMsgSubmitProposal, denom string, amount *big.Int) (*types.Transaction, error) {
	return _GovernanceWrapper.Contract.Submit(&_GovernanceWrapper.TransactOpts, proposal, denom, amount)
}

// Submit is a paid mutator transaction binding the contract method 0xebce1385.
//
// Solidity: function submit(((string,bytes)[],(uint256,string)[],string,string,string,string,bool) proposal, string denom, uint256 amount) payable returns(uint64)
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
