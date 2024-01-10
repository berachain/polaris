// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package governance

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

// IGovernanceModuleDepositParams is an auto generated low-level Go binding around an user-defined struct.
type IGovernanceModuleDepositParams struct {
	MinDeposit       []CosmosCoin
	MaxDepositPeriod uint64
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

// IGovernanceModuleParams is an auto generated low-level Go binding around an user-defined struct.
type IGovernanceModuleParams struct {
	MinDeposit                 []CosmosCoin
	MaxDepositPeriod           uint64
	VotingPeriod               uint64
	Quorum                     string
	Threshold                  string
	VetoThreshold              string
	MinInitialDepositRatio     string
	ProposalCancelRatio        string
	ProposalCancelDest         string
	ExpeditedVotingPeriod      uint64
	ExpeditedThreshold         string
	ExpeditedMinDeposit        []CosmosCoin
	BurnVoteQuorum             bool
	BurnProposalDepositPrevote bool
	BurnVoteVeto               bool
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

// IGovernanceModuleTallyParams is an auto generated low-level Go binding around an user-defined struct.
type IGovernanceModuleTallyParams struct {
	Quorum        string
	Threshold     string
	VetoThreshold string
}

// IGovernanceModuleTallyResult is an auto generated low-level Go binding around an user-defined struct.
type IGovernanceModuleTallyResult struct {
	YesCount        string
	AbstainCount    string
	NoCount         string
	NoWithVetoCount string
}

// IGovernanceModuleVote is an auto generated low-level Go binding around an user-defined struct.
type IGovernanceModuleVote struct {
	ProposalId uint64
	Voter      common.Address
	Options    []IGovernanceModuleWeightedVoteOption
	Metadata   string
}

// IGovernanceModuleVotingParams is an auto generated low-level Go binding around an user-defined struct.
type IGovernanceModuleVotingParams struct {
	VotingPeriod uint64
}

// IGovernanceModuleWeightedVoteOption is an auto generated low-level Go binding around an user-defined struct.
type IGovernanceModuleWeightedVoteOption struct {
	VoteOption int32
	Weight     string
}

// GovernanceModuleMetaData contains all meta data concerning the GovernanceModule contract.
var GovernanceModuleMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"cancelProposal\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getConstitution\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDepositParams\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIGovernanceModule.DepositParams\",\"components\":[{\"name\":\"minDeposit\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"maxDepositPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getParams\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIGovernanceModule.Params\",\"components\":[{\"name\":\"minDeposit\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"maxDepositPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"votingPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"quorum\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"threshold\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"vetoThreshold\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"minInitialDepositRatio\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"proposalCancelRatio\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"proposalCancelDest\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"expeditedVotingPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expeditedThreshold\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"expeditedMinDeposit\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"burnVoteQuorum\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"burnProposalDepositPrevote\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"burnVoteVeto\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProposal\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIGovernanceModule.Proposal\",\"components\":[{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messages\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.CodecAny[]\",\"components\":[{\"name\":\"typeURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"status\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"finalTallyResult\",\"type\":\"tuple\",\"internalType\":\"structIGovernanceModule.TallyResult\",\"components\":[{\"name\":\"yesCount\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"abstainCount\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"noCount\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"noWithVetoCount\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"submitTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"depositEndTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"totalDeposit\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"votingStartTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"votingEndTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"metadata\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"title\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"summary\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProposalDeposits\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProposalDepositsByDepositor\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"depositor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProposalTallyResult\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIGovernanceModule.TallyResult\",\"components\":[{\"name\":\"yesCount\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"abstainCount\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"noCount\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"noWithVetoCount\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProposalVotes\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pagination\",\"type\":\"tuple\",\"internalType\":\"structCosmos.PageRequest\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"offset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"limit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"countTotal\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"reverse\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIGovernanceModule.Vote[]\",\"components\":[{\"name\":\"proposalId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"voter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"options\",\"type\":\"tuple[]\",\"internalType\":\"structIGovernanceModule.WeightedVoteOption[]\",\"components\":[{\"name\":\"voteOption\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"weight\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"metadata\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCosmos.PageResponse\",\"components\":[{\"name\":\"nextKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"total\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProposalVotesByVoter\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"voter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIGovernanceModule.Vote\",\"components\":[{\"name\":\"proposalId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"voter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"options\",\"type\":\"tuple[]\",\"internalType\":\"structIGovernanceModule.WeightedVoteOption[]\",\"components\":[{\"name\":\"voteOption\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"weight\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"metadata\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProposals\",\"inputs\":[{\"name\":\"proposalStatus\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"pagination\",\"type\":\"tuple\",\"internalType\":\"structCosmos.PageRequest\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"offset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"limit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"countTotal\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"reverse\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIGovernanceModule.Proposal[]\",\"components\":[{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"messages\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.CodecAny[]\",\"components\":[{\"name\":\"typeURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"status\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"finalTallyResult\",\"type\":\"tuple\",\"internalType\":\"structIGovernanceModule.TallyResult\",\"components\":[{\"name\":\"yesCount\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"abstainCount\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"noCount\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"noWithVetoCount\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"submitTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"depositEndTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"totalDeposit\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"votingStartTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"votingEndTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"metadata\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"title\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"summary\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structCosmos.PageResponse\",\"components\":[{\"name\":\"nextKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"total\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTallyParams\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIGovernanceModule.TallyParams\",\"components\":[{\"name\":\"quorum\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"threshold\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"vetoThreshold\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVotingParams\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIGovernanceModule.VotingParams\",\"components\":[{\"name\":\"votingPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitProposal\",\"inputs\":[{\"name\":\"proposal\",\"type\":\"tuple\",\"internalType\":\"structIGovernanceModule.MsgSubmitProposal\",\"components\":[{\"name\":\"messages\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.CodecAny[]\",\"components\":[{\"name\":\"typeURL\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"initialDeposit\",\"type\":\"tuple[]\",\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"title\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"summary\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"expedited\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"vote\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"option\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"metadata\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"voteWeighted\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"options\",\"type\":\"tuple[]\",\"internalType\":\"structIGovernanceModule.WeightedVoteOption[]\",\"components\":[{\"name\":\"voteOption\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"weight\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"metadata\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CancelProposal\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProposalDeposit\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"amount\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structCosmos.Coin[]\",\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"denom\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProposalSubmitted\",\"inputs\":[{\"name\":\"proposalId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"proposalSender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProposalVoted\",\"inputs\":[{\"name\":\"proposalVote\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIGovernanceModule.Vote\",\"components\":[{\"name\":\"proposalId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"voter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"options\",\"type\":\"tuple[]\",\"internalType\":\"structIGovernanceModule.WeightedVoteOption[]\",\"components\":[{\"name\":\"voteOption\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"weight\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"metadata\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false}]",
}

// GovernanceModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use GovernanceModuleMetaData.ABI instead.
var GovernanceModuleABI = GovernanceModuleMetaData.ABI

// GovernanceModule is an auto generated Go binding around an Ethereum contract.
type GovernanceModule struct {
	GovernanceModuleCaller     // Read-only binding to the contract
	GovernanceModuleTransactor // Write-only binding to the contract
	GovernanceModuleFilterer   // Log filterer for contract events
}

// GovernanceModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernanceModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernanceModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernanceModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernanceModuleSession struct {
	Contract     *GovernanceModule // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernanceModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernanceModuleCallerSession struct {
	Contract *GovernanceModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// GovernanceModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernanceModuleTransactorSession struct {
	Contract     *GovernanceModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// GovernanceModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernanceModuleRaw struct {
	Contract *GovernanceModule // Generic contract binding to access the raw methods on
}

// GovernanceModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernanceModuleCallerRaw struct {
	Contract *GovernanceModuleCaller // Generic read-only contract binding to access the raw methods on
}

// GovernanceModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernanceModuleTransactorRaw struct {
	Contract *GovernanceModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernanceModule creates a new instance of GovernanceModule, bound to a specific deployed contract.
func NewGovernanceModule(address common.Address, backend bind.ContractBackend) (*GovernanceModule, error) {
	contract, err := bindGovernanceModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GovernanceModule{GovernanceModuleCaller: GovernanceModuleCaller{contract: contract}, GovernanceModuleTransactor: GovernanceModuleTransactor{contract: contract}, GovernanceModuleFilterer: GovernanceModuleFilterer{contract: contract}}, nil
}

// NewGovernanceModuleCaller creates a new read-only instance of GovernanceModule, bound to a specific deployed contract.
func NewGovernanceModuleCaller(address common.Address, caller bind.ContractCaller) (*GovernanceModuleCaller, error) {
	contract, err := bindGovernanceModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceModuleCaller{contract: contract}, nil
}

// NewGovernanceModuleTransactor creates a new write-only instance of GovernanceModule, bound to a specific deployed contract.
func NewGovernanceModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernanceModuleTransactor, error) {
	contract, err := bindGovernanceModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceModuleTransactor{contract: contract}, nil
}

// NewGovernanceModuleFilterer creates a new log filterer instance of GovernanceModule, bound to a specific deployed contract.
func NewGovernanceModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernanceModuleFilterer, error) {
	contract, err := bindGovernanceModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernanceModuleFilterer{contract: contract}, nil
}

// bindGovernanceModule binds a generic wrapper to an already deployed contract.
func bindGovernanceModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GovernanceModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernanceModule *GovernanceModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GovernanceModule.Contract.GovernanceModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernanceModule *GovernanceModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceModule.Contract.GovernanceModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernanceModule *GovernanceModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernanceModule.Contract.GovernanceModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernanceModule *GovernanceModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GovernanceModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernanceModule *GovernanceModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernanceModule *GovernanceModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernanceModule.Contract.contract.Transact(opts, method, params...)
}

// GetConstitution is a free data retrieval call binding the contract method 0xee05ad82.
//
// Solidity: function getConstitution() view returns(string)
func (_GovernanceModule *GovernanceModuleCaller) GetConstitution(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GovernanceModule.contract.Call(opts, &out, "getConstitution")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetConstitution is a free data retrieval call binding the contract method 0xee05ad82.
//
// Solidity: function getConstitution() view returns(string)
func (_GovernanceModule *GovernanceModuleSession) GetConstitution() (string, error) {
	return _GovernanceModule.Contract.GetConstitution(&_GovernanceModule.CallOpts)
}

// GetConstitution is a free data retrieval call binding the contract method 0xee05ad82.
//
// Solidity: function getConstitution() view returns(string)
func (_GovernanceModule *GovernanceModuleCallerSession) GetConstitution() (string, error) {
	return _GovernanceModule.Contract.GetConstitution(&_GovernanceModule.CallOpts)
}

// GetDepositParams is a free data retrieval call binding the contract method 0x8e1e4829.
//
// Solidity: function getDepositParams() view returns(((uint256,string)[],uint64))
func (_GovernanceModule *GovernanceModuleCaller) GetDepositParams(opts *bind.CallOpts) (IGovernanceModuleDepositParams, error) {
	var out []interface{}
	err := _GovernanceModule.contract.Call(opts, &out, "getDepositParams")

	if err != nil {
		return *new(IGovernanceModuleDepositParams), err
	}

	out0 := *abi.ConvertType(out[0], new(IGovernanceModuleDepositParams)).(*IGovernanceModuleDepositParams)

	return out0, err

}

// GetDepositParams is a free data retrieval call binding the contract method 0x8e1e4829.
//
// Solidity: function getDepositParams() view returns(((uint256,string)[],uint64))
func (_GovernanceModule *GovernanceModuleSession) GetDepositParams() (IGovernanceModuleDepositParams, error) {
	return _GovernanceModule.Contract.GetDepositParams(&_GovernanceModule.CallOpts)
}

// GetDepositParams is a free data retrieval call binding the contract method 0x8e1e4829.
//
// Solidity: function getDepositParams() view returns(((uint256,string)[],uint64))
func (_GovernanceModule *GovernanceModuleCallerSession) GetDepositParams() (IGovernanceModuleDepositParams, error) {
	return _GovernanceModule.Contract.GetDepositParams(&_GovernanceModule.CallOpts)
}

// GetParams is a free data retrieval call binding the contract method 0x5e615a6b.
//
// Solidity: function getParams() view returns(((uint256,string)[],uint64,uint64,string,string,string,string,string,string,uint64,string,(uint256,string)[],bool,bool,bool))
func (_GovernanceModule *GovernanceModuleCaller) GetParams(opts *bind.CallOpts) (IGovernanceModuleParams, error) {
	var out []interface{}
	err := _GovernanceModule.contract.Call(opts, &out, "getParams")

	if err != nil {
		return *new(IGovernanceModuleParams), err
	}

	out0 := *abi.ConvertType(out[0], new(IGovernanceModuleParams)).(*IGovernanceModuleParams)

	return out0, err

}

// GetParams is a free data retrieval call binding the contract method 0x5e615a6b.
//
// Solidity: function getParams() view returns(((uint256,string)[],uint64,uint64,string,string,string,string,string,string,uint64,string,(uint256,string)[],bool,bool,bool))
func (_GovernanceModule *GovernanceModuleSession) GetParams() (IGovernanceModuleParams, error) {
	return _GovernanceModule.Contract.GetParams(&_GovernanceModule.CallOpts)
}

// GetParams is a free data retrieval call binding the contract method 0x5e615a6b.
//
// Solidity: function getParams() view returns(((uint256,string)[],uint64,uint64,string,string,string,string,string,string,uint64,string,(uint256,string)[],bool,bool,bool))
func (_GovernanceModule *GovernanceModuleCallerSession) GetParams() (IGovernanceModuleParams, error) {
	return _GovernanceModule.Contract.GetParams(&_GovernanceModule.CallOpts)
}

// GetProposal is a free data retrieval call binding the contract method 0xf1610a28.
//
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,(string,bytes)[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,address))
func (_GovernanceModule *GovernanceModuleCaller) GetProposal(opts *bind.CallOpts, proposalId uint64) (IGovernanceModuleProposal, error) {
	var out []interface{}
	err := _GovernanceModule.contract.Call(opts, &out, "getProposal", proposalId)

	if err != nil {
		return *new(IGovernanceModuleProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(IGovernanceModuleProposal)).(*IGovernanceModuleProposal)

	return out0, err

}

// GetProposal is a free data retrieval call binding the contract method 0xf1610a28.
//
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,(string,bytes)[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,address))
func (_GovernanceModule *GovernanceModuleSession) GetProposal(proposalId uint64) (IGovernanceModuleProposal, error) {
	return _GovernanceModule.Contract.GetProposal(&_GovernanceModule.CallOpts, proposalId)
}

// GetProposal is a free data retrieval call binding the contract method 0xf1610a28.
//
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,(string,bytes)[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,address))
func (_GovernanceModule *GovernanceModuleCallerSession) GetProposal(proposalId uint64) (IGovernanceModuleProposal, error) {
	return _GovernanceModule.Contract.GetProposal(&_GovernanceModule.CallOpts, proposalId)
}

// GetProposalDeposits is a free data retrieval call binding the contract method 0x1bea3dc5.
//
// Solidity: function getProposalDeposits(uint64 proposalId) view returns((uint256,string)[])
func (_GovernanceModule *GovernanceModuleCaller) GetProposalDeposits(opts *bind.CallOpts, proposalId uint64) ([]CosmosCoin, error) {
	var out []interface{}
	err := _GovernanceModule.contract.Call(opts, &out, "getProposalDeposits", proposalId)

	if err != nil {
		return *new([]CosmosCoin), err
	}

	out0 := *abi.ConvertType(out[0], new([]CosmosCoin)).(*[]CosmosCoin)

	return out0, err

}

// GetProposalDeposits is a free data retrieval call binding the contract method 0x1bea3dc5.
//
// Solidity: function getProposalDeposits(uint64 proposalId) view returns((uint256,string)[])
func (_GovernanceModule *GovernanceModuleSession) GetProposalDeposits(proposalId uint64) ([]CosmosCoin, error) {
	return _GovernanceModule.Contract.GetProposalDeposits(&_GovernanceModule.CallOpts, proposalId)
}

// GetProposalDeposits is a free data retrieval call binding the contract method 0x1bea3dc5.
//
// Solidity: function getProposalDeposits(uint64 proposalId) view returns((uint256,string)[])
func (_GovernanceModule *GovernanceModuleCallerSession) GetProposalDeposits(proposalId uint64) ([]CosmosCoin, error) {
	return _GovernanceModule.Contract.GetProposalDeposits(&_GovernanceModule.CallOpts, proposalId)
}

// GetProposalDepositsByDepositor is a free data retrieval call binding the contract method 0x6d18e2e2.
//
// Solidity: function getProposalDepositsByDepositor(uint64 proposalId, address depositor) view returns((uint256,string)[])
func (_GovernanceModule *GovernanceModuleCaller) GetProposalDepositsByDepositor(opts *bind.CallOpts, proposalId uint64, depositor common.Address) ([]CosmosCoin, error) {
	var out []interface{}
	err := _GovernanceModule.contract.Call(opts, &out, "getProposalDepositsByDepositor", proposalId, depositor)

	if err != nil {
		return *new([]CosmosCoin), err
	}

	out0 := *abi.ConvertType(out[0], new([]CosmosCoin)).(*[]CosmosCoin)

	return out0, err

}

// GetProposalDepositsByDepositor is a free data retrieval call binding the contract method 0x6d18e2e2.
//
// Solidity: function getProposalDepositsByDepositor(uint64 proposalId, address depositor) view returns((uint256,string)[])
func (_GovernanceModule *GovernanceModuleSession) GetProposalDepositsByDepositor(proposalId uint64, depositor common.Address) ([]CosmosCoin, error) {
	return _GovernanceModule.Contract.GetProposalDepositsByDepositor(&_GovernanceModule.CallOpts, proposalId, depositor)
}

// GetProposalDepositsByDepositor is a free data retrieval call binding the contract method 0x6d18e2e2.
//
// Solidity: function getProposalDepositsByDepositor(uint64 proposalId, address depositor) view returns((uint256,string)[])
func (_GovernanceModule *GovernanceModuleCallerSession) GetProposalDepositsByDepositor(proposalId uint64, depositor common.Address) ([]CosmosCoin, error) {
	return _GovernanceModule.Contract.GetProposalDepositsByDepositor(&_GovernanceModule.CallOpts, proposalId, depositor)
}

// GetProposalTallyResult is a free data retrieval call binding the contract method 0xefdc5825.
//
// Solidity: function getProposalTallyResult(uint64 proposalId) view returns((string,string,string,string))
func (_GovernanceModule *GovernanceModuleCaller) GetProposalTallyResult(opts *bind.CallOpts, proposalId uint64) (IGovernanceModuleTallyResult, error) {
	var out []interface{}
	err := _GovernanceModule.contract.Call(opts, &out, "getProposalTallyResult", proposalId)

	if err != nil {
		return *new(IGovernanceModuleTallyResult), err
	}

	out0 := *abi.ConvertType(out[0], new(IGovernanceModuleTallyResult)).(*IGovernanceModuleTallyResult)

	return out0, err

}

// GetProposalTallyResult is a free data retrieval call binding the contract method 0xefdc5825.
//
// Solidity: function getProposalTallyResult(uint64 proposalId) view returns((string,string,string,string))
func (_GovernanceModule *GovernanceModuleSession) GetProposalTallyResult(proposalId uint64) (IGovernanceModuleTallyResult, error) {
	return _GovernanceModule.Contract.GetProposalTallyResult(&_GovernanceModule.CallOpts, proposalId)
}

// GetProposalTallyResult is a free data retrieval call binding the contract method 0xefdc5825.
//
// Solidity: function getProposalTallyResult(uint64 proposalId) view returns((string,string,string,string))
func (_GovernanceModule *GovernanceModuleCallerSession) GetProposalTallyResult(proposalId uint64) (IGovernanceModuleTallyResult, error) {
	return _GovernanceModule.Contract.GetProposalTallyResult(&_GovernanceModule.CallOpts, proposalId)
}

// GetProposalVotes is a free data retrieval call binding the contract method 0x0a6d4ae5.
//
// Solidity: function getProposalVotes(uint64 proposalId, (string,uint64,uint64,bool,bool) pagination) view returns((uint64,address,(int32,string)[],string)[], (string,uint64))
func (_GovernanceModule *GovernanceModuleCaller) GetProposalVotes(opts *bind.CallOpts, proposalId uint64, pagination CosmosPageRequest) ([]IGovernanceModuleVote, CosmosPageResponse, error) {
	var out []interface{}
	err := _GovernanceModule.contract.Call(opts, &out, "getProposalVotes", proposalId, pagination)

	if err != nil {
		return *new([]IGovernanceModuleVote), *new(CosmosPageResponse), err
	}

	out0 := *abi.ConvertType(out[0], new([]IGovernanceModuleVote)).(*[]IGovernanceModuleVote)
	out1 := *abi.ConvertType(out[1], new(CosmosPageResponse)).(*CosmosPageResponse)

	return out0, out1, err

}

// GetProposalVotes is a free data retrieval call binding the contract method 0x0a6d4ae5.
//
// Solidity: function getProposalVotes(uint64 proposalId, (string,uint64,uint64,bool,bool) pagination) view returns((uint64,address,(int32,string)[],string)[], (string,uint64))
func (_GovernanceModule *GovernanceModuleSession) GetProposalVotes(proposalId uint64, pagination CosmosPageRequest) ([]IGovernanceModuleVote, CosmosPageResponse, error) {
	return _GovernanceModule.Contract.GetProposalVotes(&_GovernanceModule.CallOpts, proposalId, pagination)
}

// GetProposalVotes is a free data retrieval call binding the contract method 0x0a6d4ae5.
//
// Solidity: function getProposalVotes(uint64 proposalId, (string,uint64,uint64,bool,bool) pagination) view returns((uint64,address,(int32,string)[],string)[], (string,uint64))
func (_GovernanceModule *GovernanceModuleCallerSession) GetProposalVotes(proposalId uint64, pagination CosmosPageRequest) ([]IGovernanceModuleVote, CosmosPageResponse, error) {
	return _GovernanceModule.Contract.GetProposalVotes(&_GovernanceModule.CallOpts, proposalId, pagination)
}

// GetProposalVotesByVoter is a free data retrieval call binding the contract method 0x5a274e33.
//
// Solidity: function getProposalVotesByVoter(uint64 proposalId, address voter) view returns((uint64,address,(int32,string)[],string))
func (_GovernanceModule *GovernanceModuleCaller) GetProposalVotesByVoter(opts *bind.CallOpts, proposalId uint64, voter common.Address) (IGovernanceModuleVote, error) {
	var out []interface{}
	err := _GovernanceModule.contract.Call(opts, &out, "getProposalVotesByVoter", proposalId, voter)

	if err != nil {
		return *new(IGovernanceModuleVote), err
	}

	out0 := *abi.ConvertType(out[0], new(IGovernanceModuleVote)).(*IGovernanceModuleVote)

	return out0, err

}

// GetProposalVotesByVoter is a free data retrieval call binding the contract method 0x5a274e33.
//
// Solidity: function getProposalVotesByVoter(uint64 proposalId, address voter) view returns((uint64,address,(int32,string)[],string))
func (_GovernanceModule *GovernanceModuleSession) GetProposalVotesByVoter(proposalId uint64, voter common.Address) (IGovernanceModuleVote, error) {
	return _GovernanceModule.Contract.GetProposalVotesByVoter(&_GovernanceModule.CallOpts, proposalId, voter)
}

// GetProposalVotesByVoter is a free data retrieval call binding the contract method 0x5a274e33.
//
// Solidity: function getProposalVotesByVoter(uint64 proposalId, address voter) view returns((uint64,address,(int32,string)[],string))
func (_GovernanceModule *GovernanceModuleCallerSession) GetProposalVotesByVoter(proposalId uint64, voter common.Address) (IGovernanceModuleVote, error) {
	return _GovernanceModule.Contract.GetProposalVotesByVoter(&_GovernanceModule.CallOpts, proposalId, voter)
}

// GetProposals is a free data retrieval call binding the contract method 0x917c9d92.
//
// Solidity: function getProposals(int32 proposalStatus, (string,uint64,uint64,bool,bool) pagination) view returns((uint64,(string,bytes)[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,address)[], (string,uint64))
func (_GovernanceModule *GovernanceModuleCaller) GetProposals(opts *bind.CallOpts, proposalStatus int32, pagination CosmosPageRequest) ([]IGovernanceModuleProposal, CosmosPageResponse, error) {
	var out []interface{}
	err := _GovernanceModule.contract.Call(opts, &out, "getProposals", proposalStatus, pagination)

	if err != nil {
		return *new([]IGovernanceModuleProposal), *new(CosmosPageResponse), err
	}

	out0 := *abi.ConvertType(out[0], new([]IGovernanceModuleProposal)).(*[]IGovernanceModuleProposal)
	out1 := *abi.ConvertType(out[1], new(CosmosPageResponse)).(*CosmosPageResponse)

	return out0, out1, err

}

// GetProposals is a free data retrieval call binding the contract method 0x917c9d92.
//
// Solidity: function getProposals(int32 proposalStatus, (string,uint64,uint64,bool,bool) pagination) view returns((uint64,(string,bytes)[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,address)[], (string,uint64))
func (_GovernanceModule *GovernanceModuleSession) GetProposals(proposalStatus int32, pagination CosmosPageRequest) ([]IGovernanceModuleProposal, CosmosPageResponse, error) {
	return _GovernanceModule.Contract.GetProposals(&_GovernanceModule.CallOpts, proposalStatus, pagination)
}

// GetProposals is a free data retrieval call binding the contract method 0x917c9d92.
//
// Solidity: function getProposals(int32 proposalStatus, (string,uint64,uint64,bool,bool) pagination) view returns((uint64,(string,bytes)[],int32,(string,string,string,string),uint64,uint64,(uint256,string)[],uint64,uint64,string,string,string,address)[], (string,uint64))
func (_GovernanceModule *GovernanceModuleCallerSession) GetProposals(proposalStatus int32, pagination CosmosPageRequest) ([]IGovernanceModuleProposal, CosmosPageResponse, error) {
	return _GovernanceModule.Contract.GetProposals(&_GovernanceModule.CallOpts, proposalStatus, pagination)
}

// GetTallyParams is a free data retrieval call binding the contract method 0x2f07b4a4.
//
// Solidity: function getTallyParams() view returns((string,string,string))
func (_GovernanceModule *GovernanceModuleCaller) GetTallyParams(opts *bind.CallOpts) (IGovernanceModuleTallyParams, error) {
	var out []interface{}
	err := _GovernanceModule.contract.Call(opts, &out, "getTallyParams")

	if err != nil {
		return *new(IGovernanceModuleTallyParams), err
	}

	out0 := *abi.ConvertType(out[0], new(IGovernanceModuleTallyParams)).(*IGovernanceModuleTallyParams)

	return out0, err

}

// GetTallyParams is a free data retrieval call binding the contract method 0x2f07b4a4.
//
// Solidity: function getTallyParams() view returns((string,string,string))
func (_GovernanceModule *GovernanceModuleSession) GetTallyParams() (IGovernanceModuleTallyParams, error) {
	return _GovernanceModule.Contract.GetTallyParams(&_GovernanceModule.CallOpts)
}

// GetTallyParams is a free data retrieval call binding the contract method 0x2f07b4a4.
//
// Solidity: function getTallyParams() view returns((string,string,string))
func (_GovernanceModule *GovernanceModuleCallerSession) GetTallyParams() (IGovernanceModuleTallyParams, error) {
	return _GovernanceModule.Contract.GetTallyParams(&_GovernanceModule.CallOpts)
}

// GetVotingParams is a free data retrieval call binding the contract method 0xa6c8210e.
//
// Solidity: function getVotingParams() view returns((uint64))
func (_GovernanceModule *GovernanceModuleCaller) GetVotingParams(opts *bind.CallOpts) (IGovernanceModuleVotingParams, error) {
	var out []interface{}
	err := _GovernanceModule.contract.Call(opts, &out, "getVotingParams")

	if err != nil {
		return *new(IGovernanceModuleVotingParams), err
	}

	out0 := *abi.ConvertType(out[0], new(IGovernanceModuleVotingParams)).(*IGovernanceModuleVotingParams)

	return out0, err

}

// GetVotingParams is a free data retrieval call binding the contract method 0xa6c8210e.
//
// Solidity: function getVotingParams() view returns((uint64))
func (_GovernanceModule *GovernanceModuleSession) GetVotingParams() (IGovernanceModuleVotingParams, error) {
	return _GovernanceModule.Contract.GetVotingParams(&_GovernanceModule.CallOpts)
}

// GetVotingParams is a free data retrieval call binding the contract method 0xa6c8210e.
//
// Solidity: function getVotingParams() view returns((uint64))
func (_GovernanceModule *GovernanceModuleCallerSession) GetVotingParams() (IGovernanceModuleVotingParams, error) {
	return _GovernanceModule.Contract.GetVotingParams(&_GovernanceModule.CallOpts)
}

// CancelProposal is a paid mutator transaction binding the contract method 0x37a9a59e.
//
// Solidity: function cancelProposal(uint64 proposalId) returns(uint64, uint64)
func (_GovernanceModule *GovernanceModuleTransactor) CancelProposal(opts *bind.TransactOpts, proposalId uint64) (*types.Transaction, error) {
	return _GovernanceModule.contract.Transact(opts, "cancelProposal", proposalId)
}

// CancelProposal is a paid mutator transaction binding the contract method 0x37a9a59e.
//
// Solidity: function cancelProposal(uint64 proposalId) returns(uint64, uint64)
func (_GovernanceModule *GovernanceModuleSession) CancelProposal(proposalId uint64) (*types.Transaction, error) {
	return _GovernanceModule.Contract.CancelProposal(&_GovernanceModule.TransactOpts, proposalId)
}

// CancelProposal is a paid mutator transaction binding the contract method 0x37a9a59e.
//
// Solidity: function cancelProposal(uint64 proposalId) returns(uint64, uint64)
func (_GovernanceModule *GovernanceModuleTransactorSession) CancelProposal(proposalId uint64) (*types.Transaction, error) {
	return _GovernanceModule.Contract.CancelProposal(&_GovernanceModule.TransactOpts, proposalId)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0x8ed6982d.
//
// Solidity: function submitProposal(((string,bytes)[],(uint256,string)[],address,string,string,string,bool) proposal) returns(uint64)
func (_GovernanceModule *GovernanceModuleTransactor) SubmitProposal(opts *bind.TransactOpts, proposal IGovernanceModuleMsgSubmitProposal) (*types.Transaction, error) {
	return _GovernanceModule.contract.Transact(opts, "submitProposal", proposal)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0x8ed6982d.
//
// Solidity: function submitProposal(((string,bytes)[],(uint256,string)[],address,string,string,string,bool) proposal) returns(uint64)
func (_GovernanceModule *GovernanceModuleSession) SubmitProposal(proposal IGovernanceModuleMsgSubmitProposal) (*types.Transaction, error) {
	return _GovernanceModule.Contract.SubmitProposal(&_GovernanceModule.TransactOpts, proposal)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0x8ed6982d.
//
// Solidity: function submitProposal(((string,bytes)[],(uint256,string)[],address,string,string,string,bool) proposal) returns(uint64)
func (_GovernanceModule *GovernanceModuleTransactorSession) SubmitProposal(proposal IGovernanceModuleMsgSubmitProposal) (*types.Transaction, error) {
	return _GovernanceModule.Contract.SubmitProposal(&_GovernanceModule.TransactOpts, proposal)
}

// Vote is a paid mutator transaction binding the contract method 0x19f7a0fb.
//
// Solidity: function vote(uint64 proposalId, int32 option, string metadata) returns(bool)
func (_GovernanceModule *GovernanceModuleTransactor) Vote(opts *bind.TransactOpts, proposalId uint64, option int32, metadata string) (*types.Transaction, error) {
	return _GovernanceModule.contract.Transact(opts, "vote", proposalId, option, metadata)
}

// Vote is a paid mutator transaction binding the contract method 0x19f7a0fb.
//
// Solidity: function vote(uint64 proposalId, int32 option, string metadata) returns(bool)
func (_GovernanceModule *GovernanceModuleSession) Vote(proposalId uint64, option int32, metadata string) (*types.Transaction, error) {
	return _GovernanceModule.Contract.Vote(&_GovernanceModule.TransactOpts, proposalId, option, metadata)
}

// Vote is a paid mutator transaction binding the contract method 0x19f7a0fb.
//
// Solidity: function vote(uint64 proposalId, int32 option, string metadata) returns(bool)
func (_GovernanceModule *GovernanceModuleTransactorSession) Vote(proposalId uint64, option int32, metadata string) (*types.Transaction, error) {
	return _GovernanceModule.Contract.Vote(&_GovernanceModule.TransactOpts, proposalId, option, metadata)
}

// VoteWeighted is a paid mutator transaction binding the contract method 0xf028295e.
//
// Solidity: function voteWeighted(uint64 proposalId, (int32,string)[] options, string metadata) returns(bool)
func (_GovernanceModule *GovernanceModuleTransactor) VoteWeighted(opts *bind.TransactOpts, proposalId uint64, options []IGovernanceModuleWeightedVoteOption, metadata string) (*types.Transaction, error) {
	return _GovernanceModule.contract.Transact(opts, "voteWeighted", proposalId, options, metadata)
}

// VoteWeighted is a paid mutator transaction binding the contract method 0xf028295e.
//
// Solidity: function voteWeighted(uint64 proposalId, (int32,string)[] options, string metadata) returns(bool)
func (_GovernanceModule *GovernanceModuleSession) VoteWeighted(proposalId uint64, options []IGovernanceModuleWeightedVoteOption, metadata string) (*types.Transaction, error) {
	return _GovernanceModule.Contract.VoteWeighted(&_GovernanceModule.TransactOpts, proposalId, options, metadata)
}

// VoteWeighted is a paid mutator transaction binding the contract method 0xf028295e.
//
// Solidity: function voteWeighted(uint64 proposalId, (int32,string)[] options, string metadata) returns(bool)
func (_GovernanceModule *GovernanceModuleTransactorSession) VoteWeighted(proposalId uint64, options []IGovernanceModuleWeightedVoteOption, metadata string) (*types.Transaction, error) {
	return _GovernanceModule.Contract.VoteWeighted(&_GovernanceModule.TransactOpts, proposalId, options, metadata)
}

// GovernanceModuleCancelProposalIterator is returned from FilterCancelProposal and is used to iterate over the raw logs and unpacked data for CancelProposal events raised by the GovernanceModule contract.
type GovernanceModuleCancelProposalIterator struct {
	Event *GovernanceModuleCancelProposal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceModuleCancelProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceModuleCancelProposal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceModuleCancelProposal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceModuleCancelProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceModuleCancelProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceModuleCancelProposal represents a CancelProposal event raised by the GovernanceModule contract.
type GovernanceModuleCancelProposal struct {
	ProposalId uint64
	Sender     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCancelProposal is a free log retrieval operation binding the contract event 0xa6503d2a0de5ae1ea468cd5b57a9b85d8dc0d79fb1fea0be143a8333b95328fc.
//
// Solidity: event CancelProposal(uint64 indexed proposalId, address indexed sender)
func (_GovernanceModule *GovernanceModuleFilterer) FilterCancelProposal(opts *bind.FilterOpts, proposalId []uint64, sender []common.Address) (*GovernanceModuleCancelProposalIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GovernanceModule.contract.FilterLogs(opts, "CancelProposal", proposalIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceModuleCancelProposalIterator{contract: _GovernanceModule.contract, event: "CancelProposal", logs: logs, sub: sub}, nil
}

// WatchCancelProposal is a free log subscription operation binding the contract event 0xa6503d2a0de5ae1ea468cd5b57a9b85d8dc0d79fb1fea0be143a8333b95328fc.
//
// Solidity: event CancelProposal(uint64 indexed proposalId, address indexed sender)
func (_GovernanceModule *GovernanceModuleFilterer) WatchCancelProposal(opts *bind.WatchOpts, sink chan<- *GovernanceModuleCancelProposal, proposalId []uint64, sender []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GovernanceModule.contract.WatchLogs(opts, "CancelProposal", proposalIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceModuleCancelProposal)
				if err := _GovernanceModule.contract.UnpackLog(event, "CancelProposal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCancelProposal is a log parse operation binding the contract event 0xa6503d2a0de5ae1ea468cd5b57a9b85d8dc0d79fb1fea0be143a8333b95328fc.
//
// Solidity: event CancelProposal(uint64 indexed proposalId, address indexed sender)
func (_GovernanceModule *GovernanceModuleFilterer) ParseCancelProposal(log types.Log) (*GovernanceModuleCancelProposal, error) {
	event := new(GovernanceModuleCancelProposal)
	if err := _GovernanceModule.contract.UnpackLog(event, "CancelProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceModuleProposalDepositIterator is returned from FilterProposalDeposit and is used to iterate over the raw logs and unpacked data for ProposalDeposit events raised by the GovernanceModule contract.
type GovernanceModuleProposalDepositIterator struct {
	Event *GovernanceModuleProposalDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceModuleProposalDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceModuleProposalDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceModuleProposalDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceModuleProposalDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceModuleProposalDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceModuleProposalDeposit represents a ProposalDeposit event raised by the GovernanceModule contract.
type GovernanceModuleProposalDeposit struct {
	ProposalId uint64
	Amount     []CosmosCoin
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalDeposit is a free log retrieval operation binding the contract event 0x0b8153af883fcde0ae58cdf61d0344f4e2a7ed7c15d89542956ffebd34fc3e65.
//
// Solidity: event ProposalDeposit(uint64 indexed proposalId, (uint256,string)[] amount)
func (_GovernanceModule *GovernanceModuleFilterer) FilterProposalDeposit(opts *bind.FilterOpts, proposalId []uint64) (*GovernanceModuleProposalDepositIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _GovernanceModule.contract.FilterLogs(opts, "ProposalDeposit", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceModuleProposalDepositIterator{contract: _GovernanceModule.contract, event: "ProposalDeposit", logs: logs, sub: sub}, nil
}

// WatchProposalDeposit is a free log subscription operation binding the contract event 0x0b8153af883fcde0ae58cdf61d0344f4e2a7ed7c15d89542956ffebd34fc3e65.
//
// Solidity: event ProposalDeposit(uint64 indexed proposalId, (uint256,string)[] amount)
func (_GovernanceModule *GovernanceModuleFilterer) WatchProposalDeposit(opts *bind.WatchOpts, sink chan<- *GovernanceModuleProposalDeposit, proposalId []uint64) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _GovernanceModule.contract.WatchLogs(opts, "ProposalDeposit", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceModuleProposalDeposit)
				if err := _GovernanceModule.contract.UnpackLog(event, "ProposalDeposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalDeposit is a log parse operation binding the contract event 0x0b8153af883fcde0ae58cdf61d0344f4e2a7ed7c15d89542956ffebd34fc3e65.
//
// Solidity: event ProposalDeposit(uint64 indexed proposalId, (uint256,string)[] amount)
func (_GovernanceModule *GovernanceModuleFilterer) ParseProposalDeposit(log types.Log) (*GovernanceModuleProposalDeposit, error) {
	event := new(GovernanceModuleProposalDeposit)
	if err := _GovernanceModule.contract.UnpackLog(event, "ProposalDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceModuleProposalSubmittedIterator is returned from FilterProposalSubmitted and is used to iterate over the raw logs and unpacked data for ProposalSubmitted events raised by the GovernanceModule contract.
type GovernanceModuleProposalSubmittedIterator struct {
	Event *GovernanceModuleProposalSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceModuleProposalSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceModuleProposalSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceModuleProposalSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceModuleProposalSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceModuleProposalSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceModuleProposalSubmitted represents a ProposalSubmitted event raised by the GovernanceModule contract.
type GovernanceModuleProposalSubmitted struct {
	ProposalId     uint64
	ProposalSender common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterProposalSubmitted is a free log retrieval operation binding the contract event 0xbee1516ed28c1813e21a96532fa36a6e3399ec32b15f3cd7c8e0b4d928a88b84.
//
// Solidity: event ProposalSubmitted(uint64 indexed proposalId, address indexed proposalSender)
func (_GovernanceModule *GovernanceModuleFilterer) FilterProposalSubmitted(opts *bind.FilterOpts, proposalId []uint64, proposalSender []common.Address) (*GovernanceModuleProposalSubmittedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var proposalSenderRule []interface{}
	for _, proposalSenderItem := range proposalSender {
		proposalSenderRule = append(proposalSenderRule, proposalSenderItem)
	}

	logs, sub, err := _GovernanceModule.contract.FilterLogs(opts, "ProposalSubmitted", proposalIdRule, proposalSenderRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceModuleProposalSubmittedIterator{contract: _GovernanceModule.contract, event: "ProposalSubmitted", logs: logs, sub: sub}, nil
}

// WatchProposalSubmitted is a free log subscription operation binding the contract event 0xbee1516ed28c1813e21a96532fa36a6e3399ec32b15f3cd7c8e0b4d928a88b84.
//
// Solidity: event ProposalSubmitted(uint64 indexed proposalId, address indexed proposalSender)
func (_GovernanceModule *GovernanceModuleFilterer) WatchProposalSubmitted(opts *bind.WatchOpts, sink chan<- *GovernanceModuleProposalSubmitted, proposalId []uint64, proposalSender []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var proposalSenderRule []interface{}
	for _, proposalSenderItem := range proposalSender {
		proposalSenderRule = append(proposalSenderRule, proposalSenderItem)
	}

	logs, sub, err := _GovernanceModule.contract.WatchLogs(opts, "ProposalSubmitted", proposalIdRule, proposalSenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceModuleProposalSubmitted)
				if err := _GovernanceModule.contract.UnpackLog(event, "ProposalSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalSubmitted is a log parse operation binding the contract event 0xbee1516ed28c1813e21a96532fa36a6e3399ec32b15f3cd7c8e0b4d928a88b84.
//
// Solidity: event ProposalSubmitted(uint64 indexed proposalId, address indexed proposalSender)
func (_GovernanceModule *GovernanceModuleFilterer) ParseProposalSubmitted(log types.Log) (*GovernanceModuleProposalSubmitted, error) {
	event := new(GovernanceModuleProposalSubmitted)
	if err := _GovernanceModule.contract.UnpackLog(event, "ProposalSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceModuleProposalVotedIterator is returned from FilterProposalVoted and is used to iterate over the raw logs and unpacked data for ProposalVoted events raised by the GovernanceModule contract.
type GovernanceModuleProposalVotedIterator struct {
	Event *GovernanceModuleProposalVoted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceModuleProposalVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceModuleProposalVoted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceModuleProposalVoted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceModuleProposalVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceModuleProposalVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceModuleProposalVoted represents a ProposalVoted event raised by the GovernanceModule contract.
type GovernanceModuleProposalVoted struct {
	ProposalVote IGovernanceModuleVote
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProposalVoted is a free log retrieval operation binding the contract event 0xbea88e2fb5ab72eba769e6ac6b62f35f8ffd2e85facdf45d068bc0e88d3b74e7.
//
// Solidity: event ProposalVoted((uint64,address,(int32,string)[],string) proposalVote)
func (_GovernanceModule *GovernanceModuleFilterer) FilterProposalVoted(opts *bind.FilterOpts) (*GovernanceModuleProposalVotedIterator, error) {

	logs, sub, err := _GovernanceModule.contract.FilterLogs(opts, "ProposalVoted")
	if err != nil {
		return nil, err
	}
	return &GovernanceModuleProposalVotedIterator{contract: _GovernanceModule.contract, event: "ProposalVoted", logs: logs, sub: sub}, nil
}

// WatchProposalVoted is a free log subscription operation binding the contract event 0xbea88e2fb5ab72eba769e6ac6b62f35f8ffd2e85facdf45d068bc0e88d3b74e7.
//
// Solidity: event ProposalVoted((uint64,address,(int32,string)[],string) proposalVote)
func (_GovernanceModule *GovernanceModuleFilterer) WatchProposalVoted(opts *bind.WatchOpts, sink chan<- *GovernanceModuleProposalVoted) (event.Subscription, error) {

	logs, sub, err := _GovernanceModule.contract.WatchLogs(opts, "ProposalVoted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceModuleProposalVoted)
				if err := _GovernanceModule.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalVoted is a log parse operation binding the contract event 0xbea88e2fb5ab72eba769e6ac6b62f35f8ffd2e85facdf45d068bc0e88d3b74e7.
//
// Solidity: event ProposalVoted((uint64,address,(int32,string)[],string) proposalVote)
func (_GovernanceModule *GovernanceModuleFilterer) ParseProposalVoted(log types.Log) (*GovernanceModuleProposalVoted, error) {
	event := new(GovernanceModuleProposalVoted)
	if err := _GovernanceModule.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
