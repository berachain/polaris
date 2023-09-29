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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governanceModule\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"contractIBankModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"cancelProposal\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.TallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"submitTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositEndTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structCosmos.Coin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"votingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposer\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Proposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"proposalStatus\",\"type\":\"int32\"}],\"name\":\"getProposals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.TallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"submitTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositEndTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structCosmos.Coin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"votingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposer\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Proposal[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceModule\",\"outputs\":[{\"internalType\":\"contractIGovernanceModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proposal\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"submit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"internalType\":\"int32\",\"name\":\"option\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a0604052734381dc2ab14285160c808659aee005d51255add773ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801562000057575f80fd5b5060405162001ec038038062001ec083398181016040528101906200007d91906200018e565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620000e3576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050620001be565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f62000158826200012d565b9050919050565b6200016a816200014c565b811462000175575f80fd5b50565b5f8151905062000188816200015f565b92915050565b5f60208284031215620001a657620001a562000129565b5b5f620001b58482850162000178565b91505092915050565b608051611ce9620001d75f395f6104e90152611ce95ff3fe608060405260043610610073575f3560e01c8063566fbd001161004d578063566fbd001461012157806376cdb03b14610151578063b5828df21461017b578063f1610a28146101b75761007a565b806319f7a0fb1461007e5780632b0a7032146100ba57806337a9a59e146100e45761007a565b3661007a57005b5f80fd5b348015610089575f80fd5b506100a4600480360381019061009f919061094a565b6101f3565b6040516100b191906109d0565b60405180910390f35b3480156100c5575f80fd5b506100ce61029a565b6040516100db9190610a63565b60405180910390f35b3480156100ef575f80fd5b5061010a60048036038101906101059190610a7c565b6102bd565b604051610118929190610ab6565b60405180910390f35b61013b60048036038101906101369190610bc2565b61035f565b6040516101489190610c53565b60405180910390f35b34801561015c575f80fd5b506101656104e7565b6040516101729190610c8c565b60405180910390f35b348015610186575f80fd5b506101a1600480360381019061019c9190610ca5565b61050b565b6040516101ae919061112e565b60405180910390f35b3480156101c2575f80fd5b506101dd60048036038101906101d89190610a7c565b6105c0565b6040516101ea919061128e565b60405180910390f35b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166319f7a0fb8585856040518463ffffffff1660e01b815260040161025193929190611305565b6020604051808303815f875af115801561026d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610291919061136b565b90509392505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f805f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166337a9a59e846040518263ffffffff1660e01b81526004016103179190610c53565b60408051808303815f875af1158015610332573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061035691906113aa565b91509150915091565b5f80600167ffffffffffffffff81111561037c5761037b610826565b5b6040519080825280602002602001820160405280156103b557816020015b6103a2610669565b81526020019060019003908161039a5790505b50905084848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815f8151811061040f5761040e6113e8565b5b60200260200101516020018190525082815f81518110610432576104316113e8565b5b60200260200101515f0181815250505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d238313688886040518363ffffffff1660e01b815260040161049b929190611451565b6020604051808303815f875af11580156104b7573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104db9190611473565b91505095945050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6060610515610682565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663917c9d9285846040518363ffffffff1660e01b8152600401610571929190611520565b5f60405180830381865afa15801561058b573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906105b39190611bf6565b5090508092505050919050565b6105c86106c5565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f1610a28836040518263ffffffff1660e01b81526004016106209190610c53565b5f60405180830381865afa15801561063a573d5f803e3d5ffd5b505050506040513d5f823e3d601f19601f820116820180604052508101906106629190611c6c565b9050919050565b60405180604001604052805f8152602001606081525090565b6040518060a00160405280606081526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff1681526020015f151581526020015f151581525090565b604051806101a001604052805f67ffffffffffffffff168152602001606081526020015f60030b81526020016106f9610762565b81526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff168152602001606081526020015f67ffffffffffffffff1681526020015f67ffffffffffffffff168152602001606081526020016060815260200160608152602001606081525090565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b5f604051905090565b5f80fd5b5f80fd5b5f67ffffffffffffffff82169050919050565b6107b78161079b565b81146107c1575f80fd5b50565b5f813590506107d2816107ae565b92915050565b5f8160030b9050919050565b6107ed816107d8565b81146107f7575f80fd5b50565b5f81359050610808816107e4565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61085c82610816565b810181811067ffffffffffffffff8211171561087b5761087a610826565b5b80604052505050565b5f61088d61078a565b90506108998282610853565b919050565b5f67ffffffffffffffff8211156108b8576108b7610826565b5b6108c182610816565b9050602081019050919050565b828183375f83830152505050565b5f6108ee6108e98461089e565b610884565b90508281526020810184848401111561090a57610909610812565b5b6109158482856108ce565b509392505050565b5f82601f8301126109315761093061080e565b5b81356109418482602086016108dc565b91505092915050565b5f805f6060848603121561096157610960610793565b5b5f61096e868287016107c4565b935050602061097f868287016107fa565b925050604084013567ffffffffffffffff8111156109a05761099f610797565b5b6109ac8682870161091d565b9150509250925092565b5f8115159050919050565b6109ca816109b6565b82525050565b5f6020820190506109e35f8301846109c1565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f610a2b610a26610a21846109e9565b610a08565b6109e9565b9050919050565b5f610a3c82610a11565b9050919050565b5f610a4d82610a32565b9050919050565b610a5d81610a43565b82525050565b5f602082019050610a765f830184610a54565b92915050565b5f60208284031215610a9157610a90610793565b5b5f610a9e848285016107c4565b91505092915050565b610ab08161079b565b82525050565b5f604082019050610ac95f830185610aa7565b610ad66020830184610aa7565b9392505050565b5f80fd5b5f80fd5b5f8083601f840112610afa57610af961080e565b5b8235905067ffffffffffffffff811115610b1757610b16610add565b5b602083019150836001820283011115610b3357610b32610ae1565b5b9250929050565b5f8083601f840112610b4f57610b4e61080e565b5b8235905067ffffffffffffffff811115610b6c57610b6b610add565b5b602083019150836001820283011115610b8857610b87610ae1565b5b9250929050565b5f819050919050565b610ba181610b8f565b8114610bab575f80fd5b50565b5f81359050610bbc81610b98565b92915050565b5f805f805f60608688031215610bdb57610bda610793565b5b5f86013567ffffffffffffffff811115610bf857610bf7610797565b5b610c0488828901610ae5565b9550955050602086013567ffffffffffffffff811115610c2757610c26610797565b5b610c3388828901610b3a565b93509350506040610c4688828901610bae565b9150509295509295909350565b5f602082019050610c665f830184610aa7565b92915050565b5f610c7682610a32565b9050919050565b610c8681610c6c565b82525050565b5f602082019050610c9f5f830184610c7d565b92915050565b5f60208284031215610cba57610cb9610793565b5b5f610cc7848285016107fa565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610d028161079b565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610d3f578082015181840152602081019050610d24565b5f8484015250505050565b5f610d5482610d08565b610d5e8185610d12565b9350610d6e818560208601610d22565b610d7781610816565b840191505092915050565b610d8b816107d8565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f610db582610d91565b610dbf8185610d9b565b9350610dcf818560208601610d22565b610dd881610816565b840191505092915050565b5f608083015f8301518482035f860152610dfd8282610dab565b91505060208301518482036020860152610e178282610dab565b91505060408301518482036040860152610e318282610dab565b91505060608301518482036060860152610e4b8282610dab565b9150508091505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610e8a81610b8f565b82525050565b5f604083015f830151610ea55f860182610e81565b5060208301518482036020860152610ebd8282610dab565b9150508091505092915050565b5f610ed58383610e90565b905092915050565b5f602082019050919050565b5f610ef382610e58565b610efd8185610e62565b935083602082028501610f0f85610e72565b805f5b85811015610f4a5784840389528151610f2b8582610eca565b9450610f3683610edd565b925060208a01995050600181019050610f12565b50829750879550505050505092915050565b5f6101a083015f830151610f725f860182610cf9565b5060208301518482036020860152610f8a8282610d4a565b9150506040830151610f9f6040860182610d82565b5060608301518482036060860152610fb78282610de3565b9150506080830151610fcc6080860182610cf9565b5060a0830151610fdf60a0860182610cf9565b5060c083015184820360c0860152610ff78282610ee9565b91505060e083015161100c60e0860182610cf9565b50610100830151611021610100860182610cf9565b5061012083015184820361012086015261103b8282610dab565b9150506101408301518482036101408601526110578282610dab565b9150506101608301518482036101608601526110738282610dab565b91505061018083015184820361018086015261108f8282610dab565b9150508091505092915050565b5f6110a78383610f5c565b905092915050565b5f602082019050919050565b5f6110c582610cd0565b6110cf8185610cda565b9350836020820285016110e185610cea565b805f5b8581101561111c57848403895281516110fd858261109c565b9450611108836110af565b925060208a019950506001810190506110e4565b50829750879550505050505092915050565b5f6020820190508181035f83015261114681846110bb565b905092915050565b5f6101a083015f8301516111645f860182610cf9565b506020830151848203602086015261117c8282610d4a565b91505060408301516111916040860182610d82565b50606083015184820360608601526111a98282610de3565b91505060808301516111be6080860182610cf9565b5060a08301516111d160a0860182610cf9565b5060c083015184820360c08601526111e98282610ee9565b91505060e08301516111fe60e0860182610cf9565b50610100830151611213610100860182610cf9565b5061012083015184820361012086015261122d8282610dab565b9150506101408301518482036101408601526112498282610dab565b9150506101608301518482036101608601526112658282610dab565b9150506101808301518482036101808601526112818282610dab565b9150508091505092915050565b5f6020820190508181035f8301526112a6818461114e565b905092915050565b6112b7816107d8565b82525050565b5f82825260208201905092915050565b5f6112d782610d91565b6112e181856112bd565b93506112f1818560208601610d22565b6112fa81610816565b840191505092915050565b5f6060820190506113185f830186610aa7565b61132560208301856112ae565b818103604083015261133781846112cd565b9050949350505050565b61134a816109b6565b8114611354575f80fd5b50565b5f8151905061136581611341565b92915050565b5f602082840312156113805761137f610793565b5b5f61138d84828501611357565b91505092915050565b5f815190506113a4816107ae565b92915050565b5f80604083850312156113c0576113bf610793565b5b5f6113cd85828601611396565b92505060206113de85828601611396565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f82825260208201905092915050565b5f6114308385611415565b935061143d8385846108ce565b61144683610816565b840190509392505050565b5f6020820190508181035f83015261146a818486611425565b90509392505050565b5f6020828403121561148857611487610793565b5b5f61149584828501611396565b91505092915050565b6114a7816109b6565b82525050565b5f60a083015f8301518482035f8601526114c78282610dab565b91505060208301516114dc6020860182610cf9565b5060408301516114ef6040860182610cf9565b506060830151611502606086018261149e565b506080830151611515608086018261149e565b508091505092915050565b5f6040820190506115335f8301856112ae565b818103602083015261154581846114ad565b90509392505050565b5f67ffffffffffffffff82111561156857611567610826565b5b602082029050602081019050919050565b5f80fd5b5f80fd5b5f67ffffffffffffffff82111561159b5761159a610826565b5b6115a482610816565b9050602081019050919050565b5f6115c36115be84611581565b610884565b9050828152602081018484840111156115df576115de610812565b5b6115ea848285610d22565b509392505050565b5f82601f8301126116065761160561080e565b5b81516116168482602086016115b1565b91505092915050565b5f8151905061162d816107e4565b92915050565b5f6116456116408461089e565b610884565b90508281526020810184848401111561166157611660610812565b5b61166c848285610d22565b509392505050565b5f82601f8301126116885761168761080e565b5b8151611698848260208601611633565b91505092915050565b5f608082840312156116b6576116b5611579565b5b6116c06080610884565b90505f82015167ffffffffffffffff8111156116df576116de61157d565b5b6116eb84828501611674565b5f83015250602082015167ffffffffffffffff81111561170e5761170d61157d565b5b61171a84828501611674565b602083015250604082015167ffffffffffffffff81111561173e5761173d61157d565b5b61174a84828501611674565b604083015250606082015167ffffffffffffffff81111561176e5761176d61157d565b5b61177a84828501611674565b60608301525092915050565b5f67ffffffffffffffff8211156117a05761179f610826565b5b602082029050602081019050919050565b5f815190506117bf81610b98565b92915050565b5f604082840312156117da576117d9611579565b5b6117e46040610884565b90505f6117f3848285016117b1565b5f83015250602082015167ffffffffffffffff8111156118165761181561157d565b5b61182284828501611674565b60208301525092915050565b5f61184061183b84611786565b610884565b9050808382526020820190506020840283018581111561186357611862610ae1565b5b835b818110156118aa57805167ffffffffffffffff8111156118885761188761080e565b5b80860161189589826117c5565b85526020850194505050602081019050611865565b5050509392505050565b5f82601f8301126118c8576118c761080e565b5b81516118d884826020860161182e565b91505092915050565b5f6101a082840312156118f7576118f6611579565b5b6119026101a0610884565b90505f61191184828501611396565b5f83015250602082015167ffffffffffffffff8111156119345761193361157d565b5b611940848285016115f2565b60208301525060406119548482850161161f565b604083015250606082015167ffffffffffffffff8111156119785761197761157d565b5b611984848285016116a1565b606083015250608061199884828501611396565b60808301525060a06119ac84828501611396565b60a08301525060c082015167ffffffffffffffff8111156119d0576119cf61157d565b5b6119dc848285016118b4565b60c08301525060e06119f084828501611396565b60e083015250610100611a0584828501611396565b6101008301525061012082015167ffffffffffffffff811115611a2b57611a2a61157d565b5b611a3784828501611674565b6101208301525061014082015167ffffffffffffffff811115611a5d57611a5c61157d565b5b611a6984828501611674565b6101408301525061016082015167ffffffffffffffff811115611a8f57611a8e61157d565b5b611a9b84828501611674565b6101608301525061018082015167ffffffffffffffff811115611ac157611ac061157d565b5b611acd84828501611674565b6101808301525092915050565b5f611aec611ae78461154e565b610884565b90508083825260208201905060208402830185811115611b0f57611b0e610ae1565b5b835b81811015611b5657805167ffffffffffffffff811115611b3457611b3361080e565b5b808601611b4189826118e1565b85526020850194505050602081019050611b11565b5050509392505050565b5f82601f830112611b7457611b7361080e565b5b8151611b84848260208601611ada565b91505092915050565b5f60408284031215611ba257611ba1611579565b5b611bac6040610884565b90505f82015167ffffffffffffffff811115611bcb57611bca61157d565b5b611bd784828501611674565b5f830152506020611bea84828501611396565b60208301525092915050565b5f8060408385031215611c0c57611c0b610793565b5b5f83015167ffffffffffffffff811115611c2957611c28610797565b5b611c3585828601611b60565b925050602083015167ffffffffffffffff811115611c5657611c55610797565b5b611c6285828601611b8d565b9150509250929050565b5f60208284031215611c8157611c80610793565b5b5f82015167ffffffffffffffff811115611c9e57611c9d610797565b5b611caa848285016118e1565b9150509291505056fea2646970667358221220da8f16972d0cd9d19a50f85d4b99f4471a098bbe3528987ff5e29119a7e43d1064736f6c63430008150033",
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
