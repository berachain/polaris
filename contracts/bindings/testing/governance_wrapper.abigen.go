// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testing

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

// IGovernanceModuleCoin is an auto generated low-level Go binding around an user-defined struct.
type IGovernanceModuleCoin struct {
	Amount uint64
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
	TotalDeposit     []IGovernanceModuleCoin
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governanceModule\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.TallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"submitTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositEndTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Coin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"votingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposer\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Proposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"proposalStatus\",\"type\":\"int32\"}],\"name\":\"getProposals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.TallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"submitTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositEndTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Coin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"votingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposer\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Proposal[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceModule\",\"outputs\":[{\"internalType\":\"contractIGovernanceModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proposal\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"submitProposalWrapepr\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620017353803806200173583398181016040528101906200003791906200014f565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036200009e576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000181565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200011782620000ea565b9050919050565b62000129816200010a565b81146200013557600080fd5b50565b60008151905062000149816200011e565b92915050565b600060208284031215620001685762000167620000e5565b5b6000620001788482850162000138565b91505092915050565b6115a480620001916000396000f3fe6080604052600436106100435760003560e01c80632b0a70321461004f578063b5828df21461007a578063f1610a28146100b7578063fa4204cb146100f45761004a565b3661004a57005b600080fd5b34801561005b57600080fd5b50610064610124565b60405161007191906104dd565b60405180910390f35b34801561008657600080fd5b506100a1600480360381019061009c9190610545565b610148565b6040516100ae9190610a05565b60405180910390f35b3480156100c357600080fd5b506100de60048036038101906100d99190610a53565b6101f0565b6040516100eb9190610bc3565b60405180910390f35b61010e60048036038101906101099190610c4a565b61029e565b60405161011b9190610cda565b60405180910390f35b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b606060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b5828df2836040518263ffffffff1660e01b81526004016101a39190610d04565b600060405180830381865afa1580156101c0573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906101e99190611427565b9050919050565b6101f8610393565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f1610a28836040518263ffffffff1660e01b81526004016102519190610cda565b600060405180830381865afa15801561026e573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906102979190611470565b9050919050565b60003073ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f193505050501580156102e6573d6000803e3d6000fd5b5060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663474d7f35868686866040518563ffffffff1660e01b81526004016103469493929190611506565b6020604051808303816000875af1158015610365573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103899190611541565b9050949350505050565b604051806101a00160405280600067ffffffffffffffff16815260200160608152602001600060030b81526020016103c9610436565b8152602001600067ffffffffffffffff168152602001600067ffffffffffffffff16815260200160608152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001606081526020016060815260200160608152602001606081525090565b6040518060800160405280606081526020016060815260200160608152602001606081525090565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006104a361049e6104998461045e565b61047e565b61045e565b9050919050565b60006104b582610488565b9050919050565b60006104c7826104aa565b9050919050565b6104d7816104bc565b82525050565b60006020820190506104f260008301846104ce565b92915050565b6000604051905090565b600080fd5b600080fd5b60008160030b9050919050565b6105228161050c565b811461052d57600080fd5b50565b60008135905061053f81610519565b92915050565b60006020828403121561055b5761055a610502565b5b600061056984828501610530565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600067ffffffffffffffff82169050919050565b6105bb8161059e565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b838110156105fb5780820151818401526020810190506105e0565b60008484015250505050565b6000601f19601f8301169050919050565b6000610623826105c1565b61062d81856105cc565b935061063d8185602086016105dd565b61064681610607565b840191505092915050565b61065a8161050c565b82525050565b600081519050919050565b600082825260208201905092915050565b600061068782610660565b610691818561066b565b93506106a18185602086016105dd565b6106aa81610607565b840191505092915050565b600060808301600083015184820360008601526106d2828261067c565b915050602083015184820360208601526106ec828261067c565b91505060408301518482036040860152610706828261067c565b91505060608301518482036060860152610720828261067c565b9150508091505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600060408301600083015161077160008601826105b2565b5060208301518482036020860152610789828261067c565b9150508091505092915050565b60006107a28383610759565b905092915050565b6000602082019050919050565b60006107c28261072d565b6107cc8185610738565b9350836020820285016107de85610749565b8060005b8581101561081a57848403895281516107fb8582610796565b9450610806836107aa565b925060208a019950506001810190506107e2565b50829750879550505050505092915050565b60006101a08301600083015161084560008601826105b2565b506020830151848203602086015261085d8282610618565b91505060408301516108726040860182610651565b506060830151848203606086015261088a82826106b5565b915050608083015161089f60808601826105b2565b5060a08301516108b260a08601826105b2565b5060c083015184820360c08601526108ca82826107b7565b91505060e08301516108df60e08601826105b2565b506101008301516108f46101008601826105b2565b5061012083015184820361012086015261090e828261067c565b91505061014083015184820361014086015261092a828261067c565b915050610160830151848203610160860152610946828261067c565b915050610180830151848203610180860152610962828261067c565b9150508091505092915050565b600061097b838361082c565b905092915050565b6000602082019050919050565b600061099b82610572565b6109a5818561057d565b9350836020820285016109b78561058e565b8060005b858110156109f357848403895281516109d4858261096f565b94506109df83610983565b925060208a019950506001810190506109bb565b50829750879550505050505092915050565b60006020820190508181036000830152610a1f8184610990565b905092915050565b610a308161059e565b8114610a3b57600080fd5b50565b600081359050610a4d81610a27565b92915050565b600060208284031215610a6957610a68610502565b5b6000610a7784828501610a3e565b91505092915050565b60006101a083016000830151610a9960008601826105b2565b5060208301518482036020860152610ab18282610618565b9150506040830151610ac66040860182610651565b5060608301518482036060860152610ade82826106b5565b9150506080830151610af360808601826105b2565b5060a0830151610b0660a08601826105b2565b5060c083015184820360c0860152610b1e82826107b7565b91505060e0830151610b3360e08601826105b2565b50610100830151610b486101008601826105b2565b50610120830151848203610120860152610b62828261067c565b915050610140830151848203610140860152610b7e828261067c565b915050610160830151848203610160860152610b9a828261067c565b915050610180830151848203610180860152610bb6828261067c565b9150508091505092915050565b60006020820190508181036000830152610bdd8184610a80565b905092915050565b600080fd5b600080fd5b600080fd5b60008083601f840112610c0a57610c09610be5565b5b8235905067ffffffffffffffff811115610c2757610c26610bea565b5b602083019150836001820283011115610c4357610c42610bef565b5b9250929050565b60008060008060408587031215610c6457610c63610502565b5b600085013567ffffffffffffffff811115610c8257610c81610507565b5b610c8e87828801610bf4565b9450945050602085013567ffffffffffffffff811115610cb157610cb0610507565b5b610cbd87828801610bf4565b925092505092959194509250565b610cd48161059e565b82525050565b6000602082019050610cef6000830184610ccb565b92915050565b610cfe8161050c565b82525050565b6000602082019050610d196000830184610cf5565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610d5782610607565b810181811067ffffffffffffffff82111715610d7657610d75610d1f565b5b80604052505050565b6000610d896104f8565b9050610d958282610d4e565b919050565b600067ffffffffffffffff821115610db557610db4610d1f565b5b602082029050602081019050919050565b600080fd5b600080fd5b600081519050610ddf81610a27565b92915050565b600080fd5b600067ffffffffffffffff821115610e0557610e04610d1f565b5b610e0e82610607565b9050602081019050919050565b6000610e2e610e2984610dea565b610d7f565b905082815260208101848484011115610e4a57610e49610de5565b5b610e558482856105dd565b509392505050565b600082601f830112610e7257610e71610be5565b5b8151610e82848260208601610e1b565b91505092915050565b600081519050610e9a81610519565b92915050565b600067ffffffffffffffff821115610ebb57610eba610d1f565b5b610ec482610607565b9050602081019050919050565b6000610ee4610edf84610ea0565b610d7f565b905082815260208101848484011115610f0057610eff610de5565b5b610f0b8482856105dd565b509392505050565b600082601f830112610f2857610f27610be5565b5b8151610f38848260208601610ed1565b91505092915050565b600060808284031215610f5757610f56610dc6565b5b610f616080610d7f565b9050600082015167ffffffffffffffff811115610f8157610f80610dcb565b5b610f8d84828501610f13565b600083015250602082015167ffffffffffffffff811115610fb157610fb0610dcb565b5b610fbd84828501610f13565b602083015250604082015167ffffffffffffffff811115610fe157610fe0610dcb565b5b610fed84828501610f13565b604083015250606082015167ffffffffffffffff81111561101157611010610dcb565b5b61101d84828501610f13565b60608301525092915050565b600067ffffffffffffffff82111561104457611043610d1f565b5b602082029050602081019050919050565b60006040828403121561106b5761106a610dc6565b5b6110756040610d7f565b9050600061108584828501610dd0565b600083015250602082015167ffffffffffffffff8111156110a9576110a8610dcb565b5b6110b584828501610f13565b60208301525092915050565b60006110d46110cf84611029565b610d7f565b905080838252602082019050602084028301858111156110f7576110f6610bef565b5b835b8181101561113e57805167ffffffffffffffff81111561111c5761111b610be5565b5b8086016111298982611055565b855260208501945050506020810190506110f9565b5050509392505050565b600082601f83011261115d5761115c610be5565b5b815161116d8482602086016110c1565b91505092915050565b60006101a0828403121561118d5761118c610dc6565b5b6111986101a0610d7f565b905060006111a884828501610dd0565b600083015250602082015167ffffffffffffffff8111156111cc576111cb610dcb565b5b6111d884828501610e5d565b60208301525060406111ec84828501610e8b565b604083015250606082015167ffffffffffffffff8111156112105761120f610dcb565b5b61121c84828501610f41565b606083015250608061123084828501610dd0565b60808301525060a061124484828501610dd0565b60a08301525060c082015167ffffffffffffffff81111561126857611267610dcb565b5b61127484828501611148565b60c08301525060e061128884828501610dd0565b60e08301525061010061129d84828501610dd0565b6101008301525061012082015167ffffffffffffffff8111156112c3576112c2610dcb565b5b6112cf84828501610f13565b6101208301525061014082015167ffffffffffffffff8111156112f5576112f4610dcb565b5b61130184828501610f13565b6101408301525061016082015167ffffffffffffffff81111561132757611326610dcb565b5b61133384828501610f13565b6101608301525061018082015167ffffffffffffffff81111561135957611358610dcb565b5b61136584828501610f13565b6101808301525092915050565b600061138561138084610d9a565b610d7f565b905080838252602082019050602084028301858111156113a8576113a7610bef565b5b835b818110156113ef57805167ffffffffffffffff8111156113cd576113cc610be5565b5b8086016113da8982611176565b855260208501945050506020810190506113aa565b5050509392505050565b600082601f83011261140e5761140d610be5565b5b815161141e848260208601611372565b91505092915050565b60006020828403121561143d5761143c610502565b5b600082015167ffffffffffffffff81111561145b5761145a610507565b5b611467848285016113f9565b91505092915050565b60006020828403121561148657611485610502565b5b600082015167ffffffffffffffff8111156114a4576114a3610507565b5b6114b084828501611176565b91505092915050565b600082825260208201905092915050565b82818337600083830152505050565b60006114e583856114b9565b93506114f28385846114ca565b6114fb83610607565b840190509392505050565b600060408201905081810360008301526115218186886114d9565b905081810360208301526115368184866114d9565b905095945050505050565b60006020828403121561155757611556610502565b5b600061156584828501610dd0565b9150509291505056fea2646970667358221220d9e69412bebc671990b0f2f14fffccb3057e6adf9c40035f81dcc0d6c821c9fa64736f6c63430008130033",
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

// GetProposal is a free data retrieval call binding the contract method 0xf1610a28.
//
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint64,string)[],uint64,uint64,string,string,string,string))
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
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint64,string)[],uint64,uint64,string,string,string,string))
func (_GovernanceWrapper *GovernanceWrapperSession) GetProposal(proposalId uint64) (IGovernanceModuleProposal, error) {
	return _GovernanceWrapper.Contract.GetProposal(&_GovernanceWrapper.CallOpts, proposalId)
}

// GetProposal is a free data retrieval call binding the contract method 0xf1610a28.
//
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint64,string)[],uint64,uint64,string,string,string,string))
func (_GovernanceWrapper *GovernanceWrapperCallerSession) GetProposal(proposalId uint64) (IGovernanceModuleProposal, error) {
	return _GovernanceWrapper.Contract.GetProposal(&_GovernanceWrapper.CallOpts, proposalId)
}

// GetProposals is a free data retrieval call binding the contract method 0xb5828df2.
//
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint64,string)[],uint64,uint64,string,string,string,string)[])
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
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint64,string)[],uint64,uint64,string,string,string,string)[])
func (_GovernanceWrapper *GovernanceWrapperSession) GetProposals(proposalStatus int32) ([]IGovernanceModuleProposal, error) {
	return _GovernanceWrapper.Contract.GetProposals(&_GovernanceWrapper.CallOpts, proposalStatus)
}

// GetProposals is a free data retrieval call binding the contract method 0xb5828df2.
//
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint64,string)[],uint64,uint64,string,string,string,string)[])
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

// SubmitProposalWrapepr is a paid mutator transaction binding the contract method 0xfa4204cb.
//
// Solidity: function submitProposalWrapepr(bytes proposal, bytes message) payable returns(uint64)
func (_GovernanceWrapper *GovernanceWrapperTransactor) SubmitProposalWrapepr(opts *bind.TransactOpts, proposal []byte, message []byte) (*types.Transaction, error) {
	return _GovernanceWrapper.contract.Transact(opts, "submitProposalWrapepr", proposal, message)
}

// SubmitProposalWrapepr is a paid mutator transaction binding the contract method 0xfa4204cb.
//
// Solidity: function submitProposalWrapepr(bytes proposal, bytes message) payable returns(uint64)
func (_GovernanceWrapper *GovernanceWrapperSession) SubmitProposalWrapepr(proposal []byte, message []byte) (*types.Transaction, error) {
	return _GovernanceWrapper.Contract.SubmitProposalWrapepr(&_GovernanceWrapper.TransactOpts, proposal, message)
}

// SubmitProposalWrapepr is a paid mutator transaction binding the contract method 0xfa4204cb.
//
// Solidity: function submitProposalWrapepr(bytes proposal, bytes message) payable returns(uint64)
func (_GovernanceWrapper *GovernanceWrapperTransactorSession) SubmitProposalWrapepr(proposal []byte, message []byte) (*types.Transaction, error) {
	return _GovernanceWrapper.Contract.SubmitProposalWrapepr(&_GovernanceWrapper.TransactOpts, proposal, message)
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
