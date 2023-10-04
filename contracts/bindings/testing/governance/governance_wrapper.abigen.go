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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governanceModule\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"contractIBankModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"cancelProposal\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"typeURL\",\"type\":\"string\"},{\"internalType\":\"uint8[]\",\"name\":\"value\",\"type\":\"uint8[]\"}],\"internalType\":\"structCosmos.CodecAny[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.TallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"submitTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositEndTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structCosmos.Coin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"votingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposer\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Proposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"proposalStatus\",\"type\":\"int32\"}],\"name\":\"getProposals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"typeURL\",\"type\":\"string\"},{\"internalType\":\"uint8[]\",\"name\":\"value\",\"type\":\"uint8[]\"}],\"internalType\":\"structCosmos.CodecAny[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.TallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"submitTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositEndTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structCosmos.Coin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"votingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposer\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Proposal[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceModule\",\"outputs\":[{\"internalType\":\"contractIGovernanceModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"typeURL\",\"type\":\"string\"},{\"internalType\":\"uint8[]\",\"name\":\"value\",\"type\":\"uint8[]\"}],\"internalType\":\"structCosmos.CodecAny[]\",\"name\":\"messages\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structCosmos.Coin[]\",\"name\":\"initialDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"proposer\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"expedited\",\"type\":\"bool\"}],\"internalType\":\"structIGovernanceModule.MsgSubmitProposal\",\"name\":\"proposal\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"submit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"internalType\":\"int32\",\"name\":\"option\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a0604052734381dc2ab14285160c808659aee005d51255add773ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801562000057575f80fd5b50604051620027213803806200272183398181016040528101906200007d91906200018e565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620000e3576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050620001be565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f62000158826200012d565b9050919050565b6200016a816200014c565b811462000175575f80fd5b50565b5f8151905062000188816200015f565b92915050565b5f60208284031215620001a657620001a562000129565b5b5f620001b58482850162000178565b91505092915050565b60805161254a620001d75f395f610361015261254a5ff3fe608060405260043610610073575f3560e01c806376cdb03b1161004d57806376cdb03b14610121578063773fb15e1461014b578063b5828df21461017b578063f1610a28146101b75761007a565b806319f7a0fb1461007e5780632b0a7032146100ba57806337a9a59e146100e45761007a565b3661007a57005b5f80fd5b348015610089575f80fd5b506100a4600480360381019061009f9190610947565b6101f3565b6040516100b191906109cd565b60405180910390f35b3480156100c5575f80fd5b506100ce61029a565b6040516100db9190610a60565b60405180910390f35b3480156100ef575f80fd5b5061010a60048036038101906101059190610a79565b6102bd565b604051610118929190610ab3565b60405180910390f35b34801561012c575f80fd5b5061013561035f565b6040516101429190610afa565b60405180910390f35b61016560048036038101906101609190610bc5565b610383565b6040516101729190610c52565b60405180910390f35b348015610186575f80fd5b506101a1600480360381019061019c9190610c6b565b610508565b6040516101ae9190611261565b60405180910390f35b3480156101c2575f80fd5b506101dd60048036038101906101d89190610a79565b6105bd565b6040516101ea91906113c1565b60405180910390f35b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166319f7a0fb8585856040518463ffffffff1660e01b815260040161025193929190611438565b6020604051808303815f875af115801561026d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610291919061149e565b90509392505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f805f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166337a9a59e846040518263ffffffff1660e01b81526004016103179190610c52565b60408051808303815f875af1158015610332573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061035691906114dd565b91509150915091565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f80600167ffffffffffffffff8111156103a05761039f610823565b5b6040519080825280602002602001820160405280156103d957816020015b6103c6610666565b8152602001906001900390816103be5790505b50905084848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815f815181106104335761043261151b565b5b60200260200101516020018190525082815f815181106104565761045561151b565b5b60200260200101515f0181815250505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d6ff05e7876040518263ffffffff1660e01b81526004016104bd9190611b2a565b6020604051808303815f875af11580156104d9573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104fd9190611b4a565b915050949350505050565b606061051261067f565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663917c9d9285846040518363ffffffff1660e01b815260040161056e929190611be8565b5f60405180830381865afa158015610588573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906105b09190612457565b5090508092505050919050565b6105c56106c2565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f1610a28836040518263ffffffff1660e01b815260040161061d9190610c52565b5f60405180830381865afa158015610637573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061065f91906124cd565b9050919050565b60405180604001604052805f8152602001606081525090565b6040518060a00160405280606081526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f151581526020015f151581525090565b604051806101a001604052805f67ffffffffffffffff168152602001606081526020015f60030b81526020016106f661075f565b81526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff168152602001606081526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff168152602001606081526020016060815260200160608152602001606081525090565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b5f604051905090565b5f80fd5b5f80fd5b5f67ffffffffffffffff82169050919050565b6107b481610798565b81146107be575f80fd5b50565b5f813590506107cf816107ab565b92915050565b5f8160030b9050919050565b6107ea816107d5565b81146107f4575f80fd5b50565b5f81359050610805816107e1565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61085982610813565b810181811067ffffffffffffffff8211171561087857610877610823565b5b80604052505050565b5f61088a610787565b90506108968282610850565b919050565b5f67ffffffffffffffff8211156108b5576108b4610823565b5b6108be82610813565b9050602081019050919050565b828183375f83830152505050565b5f6108eb6108e68461089b565b610881565b9050828152602081018484840111156109075761090661080f565b5b6109128482856108cb565b509392505050565b5f82601f83011261092e5761092d61080b565b5b813561093e8482602086016108d9565b91505092915050565b5f805f6060848603121561095e5761095d610790565b5b5f61096b868287016107c1565b935050602061097c868287016107f7565b925050604084013567ffffffffffffffff81111561099d5761099c610794565b5b6109a98682870161091a565b9150509250925092565b5f8115159050919050565b6109c7816109b3565b82525050565b5f6020820190506109e05f8301846109be565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f610a28610a23610a1e846109e6565b610a05565b6109e6565b9050919050565b5f610a3982610a0e565b9050919050565b5f610a4a82610a2f565b9050919050565b610a5a81610a40565b82525050565b5f602082019050610a735f830184610a51565b92915050565b5f60208284031215610a8e57610a8d610790565b5b5f610a9b848285016107c1565b91505092915050565b610aad81610798565b82525050565b5f604082019050610ac65f830185610aa4565b610ad36020830184610aa4565b9392505050565b5f610ae482610a2f565b9050919050565b610af481610ada565b82525050565b5f602082019050610b0d5f830184610aeb565b92915050565b5f80fd5b5f60e08284031215610b2c57610b2b610b13565b5b81905092915050565b5f80fd5b5f80fd5b5f8083601f840112610b5257610b5161080b565b5b8235905067ffffffffffffffff811115610b6f57610b6e610b35565b5b602083019150836001820283011115610b8b57610b8a610b39565b5b9250929050565b5f819050919050565b610ba481610b92565b8114610bae575f80fd5b50565b5f81359050610bbf81610b9b565b92915050565b5f805f8060608587031215610bdd57610bdc610790565b5b5f85013567ffffffffffffffff811115610bfa57610bf9610794565b5b610c0687828801610b17565b945050602085013567ffffffffffffffff811115610c2757610c26610794565b5b610c3387828801610b3d565b93509350506040610c4687828801610bb1565b91505092959194509250565b5f602082019050610c655f830184610aa4565b92915050565b5f60208284031215610c8057610c7f610790565b5b5f610c8d848285016107f7565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610cc881610798565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610d2e578082015181840152602081019050610d13565b5f8484015250505050565b5f610d4382610cf7565b610d4d8185610d01565b9350610d5d818560208601610d11565b610d6681610813565b840191505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f60ff82169050919050565b610daf81610d9a565b82525050565b5f610dc08383610da6565b60208301905092915050565b5f602082019050919050565b5f610de282610d71565b610dec8185610d7b565b9350610df783610d8b565b805f5b83811015610e27578151610e0e8882610db5565b9750610e1983610dcc565b925050600181019050610dfa565b5085935050505092915050565b5f604083015f8301518482035f860152610e4e8282610d39565b91505060208301518482036020860152610e688282610dd8565b9150508091505092915050565b5f610e808383610e34565b905092915050565b5f602082019050919050565b5f610e9e82610cce565b610ea88185610cd8565b935083602082028501610eba85610ce8565b805f5b85811015610ef55784840389528151610ed68582610e75565b9450610ee183610e88565b925060208a01995050600181019050610ebd565b50829750879550505050505092915050565b610f10816107d5565b82525050565b5f608083015f8301518482035f860152610f308282610d39565b91505060208301518482036020860152610f4a8282610d39565b91505060408301518482036040860152610f648282610d39565b91505060608301518482036060860152610f7e8282610d39565b9150508091505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610fbd81610b92565b82525050565b5f604083015f830151610fd85f860182610fb4565b5060208301518482036020860152610ff08282610d39565b9150508091505092915050565b5f6110088383610fc3565b905092915050565b5f602082019050919050565b5f61102682610f8b565b6110308185610f95565b93508360208202850161104285610fa5565b805f5b8581101561107d578484038952815161105e8582610ffd565b945061106983611010565b925060208a01995050600181019050611045565b50829750879550505050505092915050565b5f6101a083015f8301516110a55f860182610cbf565b50602083015184820360208601526110bd8282610e94565b91505060408301516110d26040860182610f07565b50606083015184820360608601526110ea8282610f16565b91505060808301516110ff6080860182610cbf565b5060a083015161111260a0860182610cbf565b5060c083015184820360c086015261112a828261101c565b91505060e083015161113f60e0860182610cbf565b50610100830151611154610100860182610cbf565b5061012083015184820361012086015261116e8282610d39565b91505061014083015184820361014086015261118a8282610d39565b9150506101608301518482036101608601526111a68282610d39565b9150506101808301518482036101808601526111c28282610d39565b9150508091505092915050565b5f6111da838361108f565b905092915050565b5f602082019050919050565b5f6111f882610c96565b6112028185610ca0565b93508360208202850161121485610cb0565b805f5b8581101561124f578484038952815161123085826111cf565b945061123b836111e2565b925060208a01995050600181019050611217565b50829750879550505050505092915050565b5f6020820190508181035f83015261127981846111ee565b905092915050565b5f6101a083015f8301516112975f860182610cbf565b50602083015184820360208601526112af8282610e94565b91505060408301516112c46040860182610f07565b50606083015184820360608601526112dc8282610f16565b91505060808301516112f16080860182610cbf565b5060a083015161130460a0860182610cbf565b5060c083015184820360c086015261131c828261101c565b91505060e083015161133160e0860182610cbf565b50610100830151611346610100860182610cbf565b506101208301518482036101208601526113608282610d39565b91505061014083015184820361014086015261137c8282610d39565b9150506101608301518482036101608601526113988282610d39565b9150506101808301518482036101808601526113b48282610d39565b9150508091505092915050565b5f6020820190508181035f8301526113d98184611281565b905092915050565b6113ea816107d5565b82525050565b5f82825260208201905092915050565b5f61140a82610cf7565b61141481856113f0565b9350611424818560208601610d11565b61142d81610813565b840191505092915050565b5f60608201905061144b5f830186610aa4565b61145860208301856113e1565b818103604083015261146a8184611400565b9050949350505050565b61147d816109b3565b8114611487575f80fd5b50565b5f8151905061149881611474565b92915050565b5f602082840312156114b3576114b2610790565b5b5f6114c08482850161148a565b91505092915050565b5f815190506114d7816107ab565b92915050565b5f80604083850312156114f3576114f2610790565b5b5f611500858286016114c9565b9250506020611511858286016114c9565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f80fd5b5f80fd5b5f80fd5b5f80833560016020038436030381126115705761156f611550565b5b83810192508235915060208301925067ffffffffffffffff82111561159857611597611548565b5b6020820236038313156115ae576115ad61154c565b5b509250929050565b5f819050919050565b5f80833560016020038436030381126115db576115da611550565b5b83810192508235915060208301925067ffffffffffffffff82111561160357611602611548565b5b6001820236038313156116195761161861154c565b5b509250929050565b5f61162c8385610d01565b93506116398385846108cb565b61164283610813565b840190509392505050565b5f808335600160200384360303811261166957611668611550565b5b83810192508235915060208301925067ffffffffffffffff82111561169157611690611548565b5b6020820236038313156116a7576116a661154c565b5b509250929050565b5f819050919050565b6116c181610d9a565b81146116cb575f80fd5b50565b5f813590506116dc816116b8565b92915050565b5f6116f060208401846116ce565b905092915050565b5f602082019050919050565b5f61170f8385610d7b565b935061171a826116af565b805f5b858110156117525761172f82846116e2565b6117398882610db5565b9750611744836116f8565b92505060018101905061171d565b5085925050509392505050565b5f604083016117705f8401846115bf565b8583035f870152611782838284611621565b92505050611793602084018461164d565b85830360208701526117a6838284611704565b925050508091505092915050565b5f6117bf838361175f565b905092915050565b5f823560016040038336030381126117e2576117e1611550565b5b82810191505092915050565b5f602082019050919050565b5f6118058385610cd8565b935083602084028501611817846115b6565b805f5b8781101561185a57848403895261183182846117c7565b61183b85826117b4565b9450611846836117ee565b925060208a0199505060018101905061181a565b50829750879450505050509392505050565b5f808335600160200384360303811261188857611887611550565b5b83810192508235915060208301925067ffffffffffffffff8211156118b0576118af611548565b5b6020820236038313156118c6576118c561154c565b5b509250929050565b5f819050919050565b5f6118e56020840184610bb1565b905092915050565b5f604083016118fe5f8401846118d7565b61190a5f860182610fb4565b5061191860208401846115bf565b858303602087015261192b838284611621565b925050508091505092915050565b5f61194483836118ed565b905092915050565b5f8235600160400383360303811261196757611966611550565b5b82810191505092915050565b5f602082019050919050565b5f61198a8385610f95565b93508360208402850161199c846118ce565b805f5b878110156119df5784840389526119b6828461194c565b6119c08582611939565b94506119cb83611973565b925060208a0199505060018101905061199f565b50829750879450505050509392505050565b5f813590506119ff81611474565b92915050565b5f611a1360208401846119f1565b905092915050565b611a24816109b3565b82525050565b5f60e08301611a3b5f840184611554565b8583035f870152611a4d8382846117fa565b92505050611a5e602084018461186c565b8583036020870152611a7183828461197f565b92505050611a8260408401846115bf565b8583036040870152611a95838284611621565b92505050611aa660608401846115bf565b8583036060870152611ab9838284611621565b92505050611aca60808401846115bf565b8583036080870152611add838284611621565b92505050611aee60a08401846115bf565b85830360a0870152611b01838284611621565b92505050611b1260c0840184611a05565b611b1f60c0860182611a1b565b508091505092915050565b5f6020820190508181035f830152611b428184611a2a565b905092915050565b5f60208284031215611b5f57611b5e610790565b5b5f611b6c848285016114c9565b91505092915050565b5f60a083015f8301518482035f860152611b8f8282610d39565b9150506020830151611ba46020860182610cbf565b506040830151611bb76040860182610cbf565b506060830151611bca6060860182611a1b565b506080830151611bdd6080860182611a1b565b508091505092915050565b5f604082019050611bfb5f8301856113e1565b8181036020830152611c0d8184611b75565b90509392505050565b5f67ffffffffffffffff821115611c3057611c2f610823565b5b602082029050602081019050919050565b5f80fd5b5f80fd5b5f67ffffffffffffffff821115611c6357611c62610823565b5b602082029050602081019050919050565b5f611c86611c818461089b565b610881565b905082815260208101848484011115611ca257611ca161080f565b5b611cad848285610d11565b509392505050565b5f82601f830112611cc957611cc861080b565b5b8151611cd9848260208601611c74565b91505092915050565b5f67ffffffffffffffff821115611cfc57611cfb610823565b5b602082029050602081019050919050565b5f81519050611d1b816116b8565b92915050565b5f611d33611d2e84611ce2565b610881565b90508083825260208201905060208402830185811115611d5657611d55610b39565b5b835b81811015611d7f5780611d6b8882611d0d565b845260208401935050602081019050611d58565b5050509392505050565b5f82601f830112611d9d57611d9c61080b565b5b8151611dad848260208601611d21565b91505092915050565b5f60408284031215611dcb57611dca611c41565b5b611dd56040610881565b90505f82015167ffffffffffffffff811115611df457611df3611c45565b5b611e0084828501611cb5565b5f83015250602082015167ffffffffffffffff811115611e2357611e22611c45565b5b611e2f84828501611d89565b60208301525092915050565b5f611e4d611e4884611c49565b610881565b90508083825260208201905060208402830185811115611e7057611e6f610b39565b5b835b81811015611eb757805167ffffffffffffffff811115611e9557611e9461080b565b5b808601611ea28982611db6565b85526020850194505050602081019050611e72565b5050509392505050565b5f82601f830112611ed557611ed461080b565b5b8151611ee5848260208601611e3b565b91505092915050565b5f81519050611efc816107e1565b92915050565b5f60808284031215611f1757611f16611c41565b5b611f216080610881565b90505f82015167ffffffffffffffff811115611f4057611f3f611c45565b5b611f4c84828501611cb5565b5f83015250602082015167ffffffffffffffff811115611f6f57611f6e611c45565b5b611f7b84828501611cb5565b602083015250604082015167ffffffffffffffff811115611f9f57611f9e611c45565b5b611fab84828501611cb5565b604083015250606082015167ffffffffffffffff811115611fcf57611fce611c45565b5b611fdb84828501611cb5565b60608301525092915050565b5f67ffffffffffffffff82111561200157612000610823565b5b602082029050602081019050919050565b5f8151905061202081610b9b565b92915050565b5f6040828403121561203b5761203a611c41565b5b6120456040610881565b90505f61205484828501612012565b5f83015250602082015167ffffffffffffffff81111561207757612076611c45565b5b61208384828501611cb5565b60208301525092915050565b5f6120a161209c84611fe7565b610881565b905080838252602082019050602084028301858111156120c4576120c3610b39565b5b835b8181101561210b57805167ffffffffffffffff8111156120e9576120e861080b565b5b8086016120f68982612026565b855260208501945050506020810190506120c6565b5050509392505050565b5f82601f8301126121295761212861080b565b5b815161213984826020860161208f565b91505092915050565b5f6101a0828403121561215857612157611c41565b5b6121636101a0610881565b90505f612172848285016114c9565b5f83015250602082015167ffffffffffffffff81111561219557612194611c45565b5b6121a184828501611ec1565b60208301525060406121b584828501611eee565b604083015250606082015167ffffffffffffffff8111156121d9576121d8611c45565b5b6121e584828501611f02565b60608301525060806121f9848285016114c9565b60808301525060a061220d848285016114c9565b60a08301525060c082015167ffffffffffffffff81111561223157612230611c45565b5b61223d84828501612115565b60c08301525060e0612251848285016114c9565b60e083015250610100612266848285016114c9565b6101008301525061012082015167ffffffffffffffff81111561228c5761228b611c45565b5b61229884828501611cb5565b6101208301525061014082015167ffffffffffffffff8111156122be576122bd611c45565b5b6122ca84828501611cb5565b6101408301525061016082015167ffffffffffffffff8111156122f0576122ef611c45565b5b6122fc84828501611cb5565b6101608301525061018082015167ffffffffffffffff81111561232257612321611c45565b5b61232e84828501611cb5565b6101808301525092915050565b5f61234d61234884611c16565b610881565b905080838252602082019050602084028301858111156123705761236f610b39565b5b835b818110156123b757805167ffffffffffffffff8111156123955761239461080b565b5b8086016123a28982612142565b85526020850194505050602081019050612372565b5050509392505050565b5f82601f8301126123d5576123d461080b565b5b81516123e584826020860161233b565b91505092915050565b5f6040828403121561240357612402611c41565b5b61240d6040610881565b90505f82015167ffffffffffffffff81111561242c5761242b611c45565b5b61243884828501611cb5565b5f83015250602061244b848285016114c9565b60208301525092915050565b5f806040838503121561246d5761246c610790565b5b5f83015167ffffffffffffffff81111561248a57612489610794565b5b612496858286016123c1565b925050602083015167ffffffffffffffff8111156124b7576124b6610794565b5b6124c3858286016123ee565b9150509250929050565b5f602082840312156124e2576124e1610790565b5b5f82015167ffffffffffffffff8111156124ff576124fe610794565b5b61250b84828501612142565b9150509291505056fea264697066735822122060e7fa4f1ecbfce40bec94de51b2058f0dd2bd8559093bbc62026b5ecc9d4fa664736f6c63430008150033",
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
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,(string,uint8[])[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string))
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
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,(string,uint8[])[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string))
func (_GovernanceWrapper *GovernanceWrapperSession) GetProposal(proposalId uint64) (IGovernanceModuleProposal, error) {
	return _GovernanceWrapper.Contract.GetProposal(&_GovernanceWrapper.CallOpts, proposalId)
}

// GetProposal is a free data retrieval call binding the contract method 0xf1610a28.
//
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,(string,uint8[])[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string))
func (_GovernanceWrapper *GovernanceWrapperCallerSession) GetProposal(proposalId uint64) (IGovernanceModuleProposal, error) {
	return _GovernanceWrapper.Contract.GetProposal(&_GovernanceWrapper.CallOpts, proposalId)
}

// GetProposals is a free data retrieval call binding the contract method 0xb5828df2.
//
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,(string,uint8[])[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string)[])
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
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,(string,uint8[])[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string)[])
func (_GovernanceWrapper *GovernanceWrapperSession) GetProposals(proposalStatus int32) ([]IGovernanceModuleProposal, error) {
	return _GovernanceWrapper.Contract.GetProposals(&_GovernanceWrapper.CallOpts, proposalStatus)
}

// GetProposals is a free data retrieval call binding the contract method 0xb5828df2.
//
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,(string,uint8[])[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,string)[])
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

// Submit is a paid mutator transaction binding the contract method 0x773fb15e.
//
// Solidity: function submit(((string,uint8[])[],(uint256,string)[],string,string,string,string,bool) proposal, string denom, uint256 amount) payable returns(uint64)
func (_GovernanceWrapper *GovernanceWrapperTransactor) Submit(opts *bind.TransactOpts, proposal IGovernanceModuleMsgSubmitProposal, denom string, amount *big.Int) (*types.Transaction, error) {
	return _GovernanceWrapper.contract.Transact(opts, "submit", proposal, denom, amount)
}

// Submit is a paid mutator transaction binding the contract method 0x773fb15e.
//
// Solidity: function submit(((string,uint8[])[],(uint256,string)[],string,string,string,string,bool) proposal, string denom, uint256 amount) payable returns(uint64)
func (_GovernanceWrapper *GovernanceWrapperSession) Submit(proposal IGovernanceModuleMsgSubmitProposal, denom string, amount *big.Int) (*types.Transaction, error) {
	return _GovernanceWrapper.Contract.Submit(&_GovernanceWrapper.TransactOpts, proposal, denom, amount)
}

// Submit is a paid mutator transaction binding the contract method 0x773fb15e.
//
// Solidity: function submit(((string,uint8[])[],(uint256,string)[],string,string,string,string,bool) proposal, string denom, uint256 amount) payable returns(uint64)
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
