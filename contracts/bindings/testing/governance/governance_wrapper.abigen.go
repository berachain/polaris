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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governanceModule\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"contractIBankModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"cancelProposal\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"bytes[]\",\"name\":\"messages\",\"type\":\"bytes[]\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.TallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"submitTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositEndTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structCosmos.Coin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"votingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposer\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Proposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"proposalStatus\",\"type\":\"int32\"}],\"name\":\"getProposals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"bytes[]\",\"name\":\"messages\",\"type\":\"bytes[]\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.TallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"submitTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositEndTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structCosmos.Coin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"votingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposer\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Proposal[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceModule\",\"outputs\":[{\"internalType\":\"contractIGovernanceModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proposal\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"submit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"internalType\":\"int32\",\"name\":\"option\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a0604052734381dc2ab14285160c808659aee005d51255add773ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801562000057575f80fd5b5060405162001f2538038062001f2583398181016040528101906200007d91906200018e565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620000e3576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050620001be565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f62000158826200012d565b9050919050565b6200016a816200014c565b811462000175575f80fd5b50565b5f8151905062000188816200015f565b92915050565b5f60208284031215620001a657620001a562000129565b5b5f620001b58482850162000178565b91505092915050565b608051611d4e620001d75f395f6103610152611d4e5ff3fe608060405260043610610073575f3560e01c806376cdb03b1161004d57806376cdb03b14610121578063b5828df21461014b578063f1610a2814610187578063fbab7815146101c35761007a565b806319f7a0fb1461007e5780632b0a7032146100ba57806337a9a59e146100e45761007a565b3661007a57005b5f80fd5b348015610089575f80fd5b506100a4600480360381019061009f91906108fb565b6101f3565b6040516100b19190610981565b60405180910390f35b3480156100c5575f80fd5b506100ce61029a565b6040516100db9190610a14565b60405180910390f35b3480156100ef575f80fd5b5061010a60048036038101906101059190610a2d565b6102bd565b604051610118929190610a67565b60405180910390f35b34801561012c575f80fd5b5061013561035f565b6040516101429190610aae565b60405180910390f35b348015610156575f80fd5b50610171600480360381019061016c9190610ac7565b610383565b60405161017e9190611014565b60405180910390f35b348015610192575f80fd5b506101ad60048036038101906101a89190610a2d565b610426565b6040516101ba9190611174565b60405180910390f35b6101dd60048036038101906101d89190611270565b6104cf565b6040516101ea9190611334565b60405180910390f35b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166319f7a0fb8585856040518463ffffffff1660e01b8152600401610251939291906113a4565b6020604051808303815f875af115801561026d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610291919061140a565b90509392505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f805f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166337a9a59e846040518263ffffffff1660e01b81526004016103179190611334565b60408051808303815f875af1158015610332573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103569190611449565b91509150915091565b7f000000000000000000000000000000000000000000000000000000000000000081565b60605f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b5828df2836040518263ffffffff1660e01b81526004016103dd9190611487565b5f60405180830381865afa1580156103f7573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f8201168201806040525081019061041f9190611bbd565b9050919050565b61042e61065d565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f1610a28836040518263ffffffff1660e01b81526004016104869190611334565b5f60405180830381865afa1580156104a0573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906104c89190611c04565b9050919050565b5f80600167ffffffffffffffff8111156104ec576104eb6107d7565b5b60405190808252806020026020018201604052801561052557816020015b6105126106fa565b81526020019060019003908161050a5790505b50905084848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815f8151811061057f5761057e611c4b565b5b60200260200101516020018190525082815f815181106105a2576105a1611c4b565b5b60200260200101515f0181815250505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663474d7f358a8a8a8a6040518563ffffffff1660e01b815260040161060f9493929190611cb4565b6020604051808303815f875af115801561062b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061064f9190611ced565b915050979650505050505050565b604051806101a001604052805f67ffffffffffffffff168152602001606081526020015f60030b8152602001610691610713565b81526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff168152602001606081526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff168152602001606081526020016060815260200160608152602001606081525090565b60405180604001604052805f8152602001606081525090565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b5f604051905090565b5f80fd5b5f80fd5b5f67ffffffffffffffff82169050919050565b6107688161074c565b8114610772575f80fd5b50565b5f813590506107838161075f565b92915050565b5f8160030b9050919050565b61079e81610789565b81146107a8575f80fd5b50565b5f813590506107b981610795565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61080d826107c7565b810181811067ffffffffffffffff8211171561082c5761082b6107d7565b5b80604052505050565b5f61083e61073b565b905061084a8282610804565b919050565b5f67ffffffffffffffff821115610869576108686107d7565b5b610872826107c7565b9050602081019050919050565b828183375f83830152505050565b5f61089f61089a8461084f565b610835565b9050828152602081018484840111156108bb576108ba6107c3565b5b6108c684828561087f565b509392505050565b5f82601f8301126108e2576108e16107bf565b5b81356108f284826020860161088d565b91505092915050565b5f805f6060848603121561091257610911610744565b5b5f61091f86828701610775565b9350506020610930868287016107ab565b925050604084013567ffffffffffffffff81111561095157610950610748565b5b61095d868287016108ce565b9150509250925092565b5f8115159050919050565b61097b81610967565b82525050565b5f6020820190506109945f830184610972565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f6109dc6109d76109d28461099a565b6109b9565b61099a565b9050919050565b5f6109ed826109c2565b9050919050565b5f6109fe826109e3565b9050919050565b610a0e816109f4565b82525050565b5f602082019050610a275f830184610a05565b92915050565b5f60208284031215610a4257610a41610744565b5b5f610a4f84828501610775565b91505092915050565b610a618161074c565b82525050565b5f604082019050610a7a5f830185610a58565b610a876020830184610a58565b9392505050565b5f610a98826109e3565b9050919050565b610aa881610a8e565b82525050565b5f602082019050610ac15f830184610a9f565b92915050565b5f60208284031215610adc57610adb610744565b5b5f610ae9848285016107ab565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610b248161074c565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610b8a578082015181840152602081019050610b6f565b5f8484015250505050565b5f610b9f82610b53565b610ba98185610b5d565b9350610bb9818560208601610b6d565b610bc2816107c7565b840191505092915050565b5f610bd88383610b95565b905092915050565b5f602082019050919050565b5f610bf682610b2a565b610c008185610b34565b935083602082028501610c1285610b44565b805f5b85811015610c4d5784840389528151610c2e8582610bcd565b9450610c3983610be0565b925060208a01995050600181019050610c15565b50829750879550505050505092915050565b610c6881610789565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f610c9282610c6e565b610c9c8185610c78565b9350610cac818560208601610b6d565b610cb5816107c7565b840191505092915050565b5f608083015f8301518482035f860152610cda8282610c88565b91505060208301518482036020860152610cf48282610c88565b91505060408301518482036040860152610d0e8282610c88565b91505060608301518482036060860152610d288282610c88565b9150508091505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f819050919050565b610d7081610d5e565b82525050565b5f604083015f830151610d8b5f860182610d67565b5060208301518482036020860152610da38282610c88565b9150508091505092915050565b5f610dbb8383610d76565b905092915050565b5f602082019050919050565b5f610dd982610d35565b610de38185610d3f565b935083602082028501610df585610d4f565b805f5b85811015610e305784840389528151610e118582610db0565b9450610e1c83610dc3565b925060208a01995050600181019050610df8565b50829750879550505050505092915050565b5f6101a083015f830151610e585f860182610b1b565b5060208301518482036020860152610e708282610bec565b9150506040830151610e856040860182610c5f565b5060608301518482036060860152610e9d8282610cc0565b9150506080830151610eb26080860182610b1b565b5060a0830151610ec560a0860182610b1b565b5060c083015184820360c0860152610edd8282610dcf565b91505060e0830151610ef260e0860182610b1b565b50610100830151610f07610100860182610b1b565b50610120830151848203610120860152610f218282610c88565b915050610140830151848203610140860152610f3d8282610c88565b915050610160830151848203610160860152610f598282610c88565b915050610180830151848203610180860152610f758282610c88565b9150508091505092915050565b5f610f8d8383610e42565b905092915050565b5f602082019050919050565b5f610fab82610af2565b610fb58185610afc565b935083602082028501610fc785610b0c565b805f5b858110156110025784840389528151610fe38582610f82565b9450610fee83610f95565b925060208a01995050600181019050610fca565b50829750879550505050505092915050565b5f6020820190508181035f83015261102c8184610fa1565b905092915050565b5f6101a083015f83015161104a5f860182610b1b565b50602083015184820360208601526110628282610bec565b91505060408301516110776040860182610c5f565b506060830151848203606086015261108f8282610cc0565b91505060808301516110a46080860182610b1b565b5060a08301516110b760a0860182610b1b565b5060c083015184820360c08601526110cf8282610dcf565b91505060e08301516110e460e0860182610b1b565b506101008301516110f9610100860182610b1b565b506101208301518482036101208601526111138282610c88565b91505061014083015184820361014086015261112f8282610c88565b91505061016083015184820361016086015261114b8282610c88565b9150506101808301518482036101808601526111678282610c88565b9150508091505092915050565b5f6020820190508181035f83015261118c8184611034565b905092915050565b5f80fd5b5f80fd5b5f8083601f8401126111b1576111b06107bf565b5b8235905067ffffffffffffffff8111156111ce576111cd611194565b5b6020830191508360018202830111156111ea576111e9611198565b5b9250929050565b5f8083601f840112611206576112056107bf565b5b8235905067ffffffffffffffff81111561122357611222611194565b5b60208301915083600182028301111561123f5761123e611198565b5b9250929050565b61124f81610d5e565b8114611259575f80fd5b50565b5f8135905061126a81611246565b92915050565b5f805f805f805f6080888a03121561128b5761128a610744565b5b5f88013567ffffffffffffffff8111156112a8576112a7610748565b5b6112b48a828b0161119c565b9750975050602088013567ffffffffffffffff8111156112d7576112d6610748565b5b6112e38a828b0161119c565b9550955050604088013567ffffffffffffffff81111561130657611305610748565b5b6113128a828b016111f1565b935093505060606113258a828b0161125c565b91505092959891949750929550565b5f6020820190506113475f830184610a58565b92915050565b61135681610789565b82525050565b5f82825260208201905092915050565b5f61137682610c6e565b611380818561135c565b9350611390818560208601610b6d565b611399816107c7565b840191505092915050565b5f6060820190506113b75f830186610a58565b6113c4602083018561134d565b81810360408301526113d6818461136c565b9050949350505050565b6113e981610967565b81146113f3575f80fd5b50565b5f81519050611404816113e0565b92915050565b5f6020828403121561141f5761141e610744565b5b5f61142c848285016113f6565b91505092915050565b5f815190506114438161075f565b92915050565b5f806040838503121561145f5761145e610744565b5b5f61146c85828601611435565b925050602061147d85828601611435565b9150509250929050565b5f60208201905061149a5f83018461134d565b92915050565b5f67ffffffffffffffff8211156114ba576114b96107d7565b5b602082029050602081019050919050565b5f80fd5b5f80fd5b5f67ffffffffffffffff8211156114ed576114ec6107d7565b5b602082029050602081019050919050565b5f67ffffffffffffffff821115611518576115176107d7565b5b611521826107c7565b9050602081019050919050565b5f61154061153b846114fe565b610835565b90508281526020810184848401111561155c5761155b6107c3565b5b611567848285610b6d565b509392505050565b5f82601f830112611583576115826107bf565b5b815161159384826020860161152e565b91505092915050565b5f6115ae6115a9846114d3565b610835565b905080838252602082019050602084028301858111156115d1576115d0611198565b5b835b8181101561161857805167ffffffffffffffff8111156115f6576115f56107bf565b5b808601611603898261156f565b855260208501945050506020810190506115d3565b5050509392505050565b5f82601f830112611636576116356107bf565b5b815161164684826020860161159c565b91505092915050565b5f8151905061165d81610795565b92915050565b5f6116756116708461084f565b610835565b905082815260208101848484011115611691576116906107c3565b5b61169c848285610b6d565b509392505050565b5f82601f8301126116b8576116b76107bf565b5b81516116c8848260208601611663565b91505092915050565b5f608082840312156116e6576116e56114cb565b5b6116f06080610835565b90505f82015167ffffffffffffffff81111561170f5761170e6114cf565b5b61171b848285016116a4565b5f83015250602082015167ffffffffffffffff81111561173e5761173d6114cf565b5b61174a848285016116a4565b602083015250604082015167ffffffffffffffff81111561176e5761176d6114cf565b5b61177a848285016116a4565b604083015250606082015167ffffffffffffffff81111561179e5761179d6114cf565b5b6117aa848285016116a4565b60608301525092915050565b5f67ffffffffffffffff8211156117d0576117cf6107d7565b5b602082029050602081019050919050565b5f815190506117ef81611246565b92915050565b5f6040828403121561180a576118096114cb565b5b6118146040610835565b90505f611823848285016117e1565b5f83015250602082015167ffffffffffffffff811115611846576118456114cf565b5b611852848285016116a4565b60208301525092915050565b5f61187061186b846117b6565b610835565b9050808382526020820190506020840283018581111561189357611892611198565b5b835b818110156118da57805167ffffffffffffffff8111156118b8576118b76107bf565b5b8086016118c589826117f5565b85526020850194505050602081019050611895565b5050509392505050565b5f82601f8301126118f8576118f76107bf565b5b815161190884826020860161185e565b91505092915050565b5f6101a08284031215611927576119266114cb565b5b6119326101a0610835565b90505f61194184828501611435565b5f83015250602082015167ffffffffffffffff811115611964576119636114cf565b5b61197084828501611622565b60208301525060406119848482850161164f565b604083015250606082015167ffffffffffffffff8111156119a8576119a76114cf565b5b6119b4848285016116d1565b60608301525060806119c884828501611435565b60808301525060a06119dc84828501611435565b60a08301525060c082015167ffffffffffffffff811115611a00576119ff6114cf565b5b611a0c848285016118e4565b60c08301525060e0611a2084828501611435565b60e083015250610100611a3584828501611435565b6101008301525061012082015167ffffffffffffffff811115611a5b57611a5a6114cf565b5b611a67848285016116a4565b6101208301525061014082015167ffffffffffffffff811115611a8d57611a8c6114cf565b5b611a99848285016116a4565b6101408301525061016082015167ffffffffffffffff811115611abf57611abe6114cf565b5b611acb848285016116a4565b6101608301525061018082015167ffffffffffffffff811115611af157611af06114cf565b5b611afd848285016116a4565b6101808301525092915050565b5f611b1c611b17846114a0565b610835565b90508083825260208201905060208402830185811115611b3f57611b3e611198565b5b835b81811015611b8657805167ffffffffffffffff811115611b6457611b636107bf565b5b808601611b718982611911565b85526020850194505050602081019050611b41565b5050509392505050565b5f82601f830112611ba457611ba36107bf565b5b8151611bb4848260208601611b0a565b91505092915050565b5f60208284031215611bd257611bd1610744565b5b5f82015167ffffffffffffffff811115611bef57611bee610748565b5b611bfb84828501611b90565b91505092915050565b5f60208284031215611c1957611c18610744565b5b5f82015167ffffffffffffffff811115611c3657611c35610748565b5b611c4284828501611911565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f82825260208201905092915050565b5f611c938385611c78565b9350611ca083858461087f565b611ca9836107c7565b840190509392505050565b5f6040820190508181035f830152611ccd818688611c88565b90508181036020830152611ce2818486611c88565b905095945050505050565b5f60208284031215611d0257611d01610744565b5b5f611d0f84828501611435565b9150509291505056fea2646970667358221220ea6eca4cc72d313396323d5de61b215083554e3ac135afa5504b95780999298964736f6c63430008140033",
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
