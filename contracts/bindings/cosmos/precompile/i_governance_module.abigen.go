// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package precompile

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

// IGovernanceModuleWeightedVoteOption is an auto generated low-level Go binding around an user-defined struct.
type IGovernanceModuleWeightedVoteOption struct {
	VoteOption int32
	Weight     string
}

// GovernanceModuleMetaData contains all meta data concerning the GovernanceModule contract.
var GovernanceModuleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"option\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"ProposalVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"proposalMessage\",\"type\":\"string\"}],\"name\":\"SubmitProposal\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"cancelProposal\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.TallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"submitTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositEndTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Coin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"votingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposer\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Proposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"proposalStatus\",\"type\":\"int32\"}],\"name\":\"getProposals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.TallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"submitTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"depositEndTime\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Coin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"votingStartTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"votingEndTime\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposer\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.Proposal[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proposal\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"submitProposal\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"internalType\":\"int32\",\"name\":\"option\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"int32\",\"name\":\"voteOption\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"weight\",\"type\":\"string\"}],\"internalType\":\"structIGovernanceModule.WeightedVoteOption[]\",\"name\":\"options\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"voteWeighted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// GetProposal is a free data retrieval call binding the contract method 0xf1610a28.
//
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint64,string)[],uint64,uint64,string,string,string,string))
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
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint64,string)[],uint64,uint64,string,string,string,string))
func (_GovernanceModule *GovernanceModuleSession) GetProposal(proposalId uint64) (IGovernanceModuleProposal, error) {
	return _GovernanceModule.Contract.GetProposal(&_GovernanceModule.CallOpts, proposalId)
}

// GetProposal is a free data retrieval call binding the contract method 0xf1610a28.
//
// Solidity: function getProposal(uint64 proposalId) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint64,string)[],uint64,uint64,string,string,string,string))
func (_GovernanceModule *GovernanceModuleCallerSession) GetProposal(proposalId uint64) (IGovernanceModuleProposal, error) {
	return _GovernanceModule.Contract.GetProposal(&_GovernanceModule.CallOpts, proposalId)
}

// GetProposals is a free data retrieval call binding the contract method 0xb5828df2.
//
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint64,string)[],uint64,uint64,string,string,string,string)[])
func (_GovernanceModule *GovernanceModuleCaller) GetProposals(opts *bind.CallOpts, proposalStatus int32) ([]IGovernanceModuleProposal, error) {
	var out []interface{}
	err := _GovernanceModule.contract.Call(opts, &out, "getProposals", proposalStatus)

	if err != nil {
		return *new([]IGovernanceModuleProposal), err
	}

	out0 := *abi.ConvertType(out[0], new([]IGovernanceModuleProposal)).(*[]IGovernanceModuleProposal)

	return out0, err

}

// GetProposals is a free data retrieval call binding the contract method 0xb5828df2.
//
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint64,string)[],uint64,uint64,string,string,string,string)[])
func (_GovernanceModule *GovernanceModuleSession) GetProposals(proposalStatus int32) ([]IGovernanceModuleProposal, error) {
	return _GovernanceModule.Contract.GetProposals(&_GovernanceModule.CallOpts, proposalStatus)
}

// GetProposals is a free data retrieval call binding the contract method 0xb5828df2.
//
// Solidity: function getProposals(int32 proposalStatus) view returns((uint64,bytes,int32,(string,string,string,string),uint64,uint64,(uint64,string)[],uint64,uint64,string,string,string,string)[])
func (_GovernanceModule *GovernanceModuleCallerSession) GetProposals(proposalStatus int32) ([]IGovernanceModuleProposal, error) {
	return _GovernanceModule.Contract.GetProposals(&_GovernanceModule.CallOpts, proposalStatus)
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

// SubmitProposal is a paid mutator transaction binding the contract method 0x474d7f35.
//
// Solidity: function submitProposal(bytes proposal, bytes message) returns(uint64)
func (_GovernanceModule *GovernanceModuleTransactor) SubmitProposal(opts *bind.TransactOpts, proposal []byte, message []byte) (*types.Transaction, error) {
	return _GovernanceModule.contract.Transact(opts, "submitProposal", proposal, message)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0x474d7f35.
//
// Solidity: function submitProposal(bytes proposal, bytes message) returns(uint64)
func (_GovernanceModule *GovernanceModuleSession) SubmitProposal(proposal []byte, message []byte) (*types.Transaction, error) {
	return _GovernanceModule.Contract.SubmitProposal(&_GovernanceModule.TransactOpts, proposal, message)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0x474d7f35.
//
// Solidity: function submitProposal(bytes proposal, bytes message) returns(uint64)
func (_GovernanceModule *GovernanceModuleTransactorSession) SubmitProposal(proposal []byte, message []byte) (*types.Transaction, error) {
	return _GovernanceModule.Contract.SubmitProposal(&_GovernanceModule.TransactOpts, proposal, message)
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

// GovernanceModuleProposalVoteIterator is returned from FilterProposalVote and is used to iterate over the raw logs and unpacked data for ProposalVote events raised by the GovernanceModule contract.
type GovernanceModuleProposalVoteIterator struct {
	Event *GovernanceModuleProposalVote // Event containing the contract specifics and raw log

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
func (it *GovernanceModuleProposalVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceModuleProposalVote)
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
		it.Event = new(GovernanceModuleProposalVote)
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
func (it *GovernanceModuleProposalVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceModuleProposalVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceModuleProposalVote represents a ProposalVote event raised by the GovernanceModule contract.
type GovernanceModuleProposalVote struct {
	Option     common.Hash
	ProposalId uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalVote is a free log retrieval operation binding the contract event 0xfa2e2ddc78fcaa00a2df7fb51ad4637527240242794177715d2e014467da4730.
//
// Solidity: event ProposalVote(string indexed option, uint64 indexed proposalId)
func (_GovernanceModule *GovernanceModuleFilterer) FilterProposalVote(opts *bind.FilterOpts, option []string, proposalId []uint64) (*GovernanceModuleProposalVoteIterator, error) {

	var optionRule []interface{}
	for _, optionItem := range option {
		optionRule = append(optionRule, optionItem)
	}
	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _GovernanceModule.contract.FilterLogs(opts, "ProposalVote", optionRule, proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceModuleProposalVoteIterator{contract: _GovernanceModule.contract, event: "ProposalVote", logs: logs, sub: sub}, nil
}

// WatchProposalVote is a free log subscription operation binding the contract event 0xfa2e2ddc78fcaa00a2df7fb51ad4637527240242794177715d2e014467da4730.
//
// Solidity: event ProposalVote(string indexed option, uint64 indexed proposalId)
func (_GovernanceModule *GovernanceModuleFilterer) WatchProposalVote(opts *bind.WatchOpts, sink chan<- *GovernanceModuleProposalVote, option []string, proposalId []uint64) (event.Subscription, error) {

	var optionRule []interface{}
	for _, optionItem := range option {
		optionRule = append(optionRule, optionItem)
	}
	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _GovernanceModule.contract.WatchLogs(opts, "ProposalVote", optionRule, proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceModuleProposalVote)
				if err := _GovernanceModule.contract.UnpackLog(event, "ProposalVote", log); err != nil {
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

// ParseProposalVote is a log parse operation binding the contract event 0xfa2e2ddc78fcaa00a2df7fb51ad4637527240242794177715d2e014467da4730.
//
// Solidity: event ProposalVote(string indexed option, uint64 indexed proposalId)
func (_GovernanceModule *GovernanceModuleFilterer) ParseProposalVote(log types.Log) (*GovernanceModuleProposalVote, error) {
	event := new(GovernanceModuleProposalVote)
	if err := _GovernanceModule.contract.UnpackLog(event, "ProposalVote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceModuleSubmitProposalIterator is returned from FilterSubmitProposal and is used to iterate over the raw logs and unpacked data for SubmitProposal events raised by the GovernanceModule contract.
type GovernanceModuleSubmitProposalIterator struct {
	Event *GovernanceModuleSubmitProposal // Event containing the contract specifics and raw log

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
func (it *GovernanceModuleSubmitProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceModuleSubmitProposal)
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
		it.Event = new(GovernanceModuleSubmitProposal)
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
func (it *GovernanceModuleSubmitProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceModuleSubmitProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceModuleSubmitProposal represents a SubmitProposal event raised by the GovernanceModule contract.
type GovernanceModuleSubmitProposal struct {
	ProposalId      uint64
	ProposalMessage common.Hash
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSubmitProposal is a free log retrieval operation binding the contract event 0x49f3ba3f7499ad49b71e9f07e9b953519567c10f84a87cfeb5b56f83f1590887.
//
// Solidity: event SubmitProposal(uint64 indexed proposalId, string indexed proposalMessage)
func (_GovernanceModule *GovernanceModuleFilterer) FilterSubmitProposal(opts *bind.FilterOpts, proposalId []uint64, proposalMessage []string) (*GovernanceModuleSubmitProposalIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var proposalMessageRule []interface{}
	for _, proposalMessageItem := range proposalMessage {
		proposalMessageRule = append(proposalMessageRule, proposalMessageItem)
	}

	logs, sub, err := _GovernanceModule.contract.FilterLogs(opts, "SubmitProposal", proposalIdRule, proposalMessageRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceModuleSubmitProposalIterator{contract: _GovernanceModule.contract, event: "SubmitProposal", logs: logs, sub: sub}, nil
}

// WatchSubmitProposal is a free log subscription operation binding the contract event 0x49f3ba3f7499ad49b71e9f07e9b953519567c10f84a87cfeb5b56f83f1590887.
//
// Solidity: event SubmitProposal(uint64 indexed proposalId, string indexed proposalMessage)
func (_GovernanceModule *GovernanceModuleFilterer) WatchSubmitProposal(opts *bind.WatchOpts, sink chan<- *GovernanceModuleSubmitProposal, proposalId []uint64, proposalMessage []string) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var proposalMessageRule []interface{}
	for _, proposalMessageItem := range proposalMessage {
		proposalMessageRule = append(proposalMessageRule, proposalMessageItem)
	}

	logs, sub, err := _GovernanceModule.contract.WatchLogs(opts, "SubmitProposal", proposalIdRule, proposalMessageRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceModuleSubmitProposal)
				if err := _GovernanceModule.contract.UnpackLog(event, "SubmitProposal", log); err != nil {
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

// ParseSubmitProposal is a log parse operation binding the contract event 0x49f3ba3f7499ad49b71e9f07e9b953519567c10f84a87cfeb5b56f83f1590887.
//
// Solidity: event SubmitProposal(uint64 indexed proposalId, string indexed proposalMessage)
func (_GovernanceModule *GovernanceModuleFilterer) ParseSubmitProposal(log types.Log) (*GovernanceModuleSubmitProposal, error) {
	event := new(GovernanceModuleSubmitProposal)
	if err := _GovernanceModule.contract.UnpackLog(event, "SubmitProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
