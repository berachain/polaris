// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package core

import (
	"errors"
	"fmt"
	"math"

	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/eth/params"
	errorslib "github.com/berachain/stargazer/lib/errors"
)

// `StateTransition` is the main object which takes care of applying a
// transaction to the current state.
type StateTransition struct {
	// An instance of the  Virtual Machine
	evm vm.StargazerEVM

	// The message to deliver to the EVM
	msg Message

	// Gas consumption tracking
	gp GasPlugin
}

// =============================================================================
// Transition Execution
// =============================================================================

// `ApplyMessage` transitions the state by applying the given message to the chain state
// using the given EVM. It also finalizes the change.
func ApplyMessage(
	evm vm.StargazerEVM,
	gp GasPlugin,
	msg Message,
	commit bool,
) (*ExecutionResult, error) {
	res, err := NewStateTransition(evm, gp, msg).transitionDB()
	if err != nil {
		return nil, errorslib.Wrap(err, "failed to TransitionDB")
	}

	if commit && !res.Failed() {
		evm.StateDB().Finalize()
	}
	return res, nil
}

// =============================================================================
// Constructor
// =============================================================================

// NewStateTransition creates a new state transition object.
func NewStateTransition(evm vm.StargazerEVM, gp GasPlugin, msg Message) *StateTransition {
	// Configure transaction config from the message.
	evm.SetTxContext(vm.TxContext{
		Origin:   msg.From(),
		GasPrice: msg.GasPrice(),
	})

	// Setup the gas plugin with the message gas limit.
	// TODO handle error?
	_ = gp.SetGasLimit(msg.Gas())

	return &StateTransition{
		evm: evm,
		msg: msg,
		gp:  gp,
	}
}

// =============================================================================
// Low Level Transition w/State Machine
// =============================================================================

// `TransitionDB` executes the configured message in the Ethereum Virtual Machine (EVM) and
// returns the execution result. The function does a number of checks and operations
// before and after executing the message in the EVM, including:
//
//  1. Computing the intrinsic gas needed for the message based on its data,
//     access list, and other parameters
//  2. Checking that the sender has sufficient funds to send value in the message, if applicable
//  3. Preparing the access list for the message, if applicable
//  4. Executing the message in the EVM, either through a contract creation or a call
//  5. Checking that the EVM did not use more gas than was supplied
//  6. Calculating and applying any gas refunds, if applicable
//  7. Updating the sender's nonce in the state database (sdb)
//
//nolint:funlen // all this code is logically contagious.
func (st *StateTransition) transitionDB() (*ExecutionResult, error) {
	var (
		msgFrom  = st.msg.From()
		msgValue = st.msg.Value()
		ctx      = st.evm.Context()
		sender   = vm.AccountRef(msgFrom)
		rules    = st.evm.ChainConfig().Rules(
			ctx.BlockNumber,
			ctx.Random != nil,
		)
		sdb              = st.evm.StateDB()
		contractCreation = st.msg.To() == nil
		tracer           = st.evm.Config().Tracer
	)

	if tracer != nil && st.evm.Config().Debug {
		// Capture the starting gas for the tracer, we can skip the check for debug mode that is
		// present in geth, as we already know that the EVM is in debug mode from the lines above.
		tracer.CaptureTxStart(st.gp.GasRemaining())
		defer func() {
			// After execution is completed we need to capture gas remaining.
			tracer.CaptureTxEnd(st.gp.GasRemaining())
		}()
	}

	// Ensure that the intrinsic gas is consumed.
	// TODO: Handle updated gas requirements for Shanghai.
	if err := st.ConsumeEthIntrinsicGas(contractCreation, rules.IsHomestead, rules.IsIstanbul); err != nil {
		return nil, err
	}

	// Check to ensure the sender has the funds to cover the value being sent.
	if msgValue.Sign() > 0 && !ctx.CanTransfer(sdb, msgFrom, msgValue) {
		return nil, fmt.Errorf("%w: address %v", ErrInsufficientFundsForTransfer, msgFrom.Hex())
	}

	// Stargazer does not support access lists.
	// if rules.IsBerlin {
	// 	sdb.PrepareAccessList(
	// 		msgFrom,
	// 		st.msg.To(),
	// 		st.evm.ActivePrecompiles(rules),
	// 		st.msg.AccessList(),
	// 	)
	// }

	var (
		ret              []byte // return bytes from evm execution
		vmErr            error  // vm errors don't effect consensus and are therefore not passed to err
		postExecutionGas uint64
	)

	// take over the nonce management from evm:
	// - reset sender's nonce to msg.Nonce() before calling evm.
	// - increase sender's nonce by one no matter the result.
	// - this is probably not required, but adds a safety measure,
	//   to ensure that the nonce is getting updated correctly.
	//
	if contractCreation {
		// TODO: Review nonce accounting. Leaving the management of the nonce
		// up to the implementing chain?
		ret, _, postExecutionGas, vmErr = st.evm.Create(sender,
			st.msg.Data(), st.gp.GasRemaining(), msgValue)
	} else {
		// TODO: Review nonce accounting. Leaving the management of the nonce
		// up to the implementing chain?
		sdb.SetNonce(sender.Address(), st.msg.Nonce()+1)
		// It is to deference st.msg.To() here, as it is checked
		// to be non-nil higher up in this function.
		ret, postExecutionGas, vmErr = st.evm.Call(sender, *st.msg.To(),
			st.msg.Data(), st.gp.GasRemaining(), msgValue)
	}

	// Consume the gas used by the EVM execution.
	if err := st.gp.ConsumeGas(st.gp.GasRemaining() - postExecutionGas); err != nil {
		vmErr = vm.ErrOutOfGas
		if errors.Is(err, ErrBlockOutOfGas) {
			// If consuming the amount of gas would exceed the block limit, we should
			// consume up to the limit here.
			// Cumulative gas used should be equal to the gas consumed in the block thus far,
			// INCLUDING the gas consumed as part of the Intrinsic gas calculation above.
			if err = st.gp.ConsumeGas(st.gp.BlockGasLimit() - st.gp.CumulativeGasUsed()); err != nil {
				return nil, err
			}
		} else {
			// If we error here for any other reason, we should return a consensus breaking error.
			return nil, err
		}
	}

	if !rules.IsLondon {
		// Before EIP-3529: refunds were capped to gasUsed / 2
		st.refundGas(params.RefundQuotient)
	} else {
		// After EIP-3529: refunds are capped to gasUsed / 5
		st.refundGas(params.RefundQuotientEIP3529)
	}

	return &ExecutionResult{
		UsedGas:    st.gp.GasUsed(),
		Err:        vmErr,
		ReturnData: ret,
	}, nil
}

// `refundGas` is a helper function that refunds the gas to the sender. It is used
// to refund unused gas after a transaction has been executed. The refund is capped
// to a refund quotient.
func (st *StateTransition) refundGas(refundQuotient uint64) {
	sdb := st.evm.StateDB()
	// Apply refund counter, capped to a refund quotient
	refund := st.gp.GasUsed() / refundQuotient
	if refund > sdb.GetRefund() {
		refund = sdb.GetRefund()
	}
	st.gp.RefundGas(refund)

	// Stargazer does not handle the actual token refund, just the gas refund.
}

// `consumeEthIntrinsicGas` is a helper function that calculates the intrinsic gas for the message with
// its given data.
func (st *StateTransition) ConsumeEthIntrinsicGas(
	isContractCreation bool, isHomestead, isEIP2028 bool,
) error {
	var gas uint64
	// Consume the starting gas for the raw transaction.
	gasUsed := st.gp.GasUsed()
	if isContractCreation && isHomestead {
		// If the meter has not yet consumed 53000 gas, we
		// want to make the gasPlugin consumes the delta.
		if gasUsed <= params.TxGasContractCreation {
			gas = params.TxGasContractCreation - gasUsed
		}
	} else {
		// If the meter has not yet consumed 21000 gas, we
		// want to make the gasPlugin consumes the delta.
		if gasUsed <= params.TxGas {
			gas = params.TxGas - gasUsed
		}
	}

	// Bump the required gas by the amount of transactional data
	if data := st.msg.Data(); len(data) > 0 {
		// Zero and non-zero bytes are priced differently
		var nz uint64
		for _, byt := range data {
			if byt != 0 {
				nz++
			}
		}
		// Make sure we don't exceed uint64 for all data combinations
		nonZeroGas := params.TxDataNonZeroGasFrontier
		if isEIP2028 {
			nonZeroGas = params.TxDataNonZeroGasEIP2028
		}
		if (math.MaxUint64-gas)/nonZeroGas < nz {
			return ErrGasUintOverflow
		}
		gas += nz * nonZeroGas

		z := uint64(len(data)) - nz
		if (math.MaxUint64-gas)/params.TxDataZeroGas < z {
			return ErrGasUintOverflow
		}
		gas += z * params.TxDataZeroGas

		// TODO: EIP-3860 dynamic transaction pricing
		// https://github.com/ethereum/go-ethereum/commit/793f0f9ec860f6f51e0cec943a268c10863097c7
	}
	if accessList := st.msg.AccessList(); accessList != nil {
		gas += uint64(len(accessList)) * params.TxAccessListAddressGas
		gas += uint64(accessList.StorageKeys()) * params.TxAccessListStorageKeyGas
	}

	// Now that we have calculated the intrinsic gas, we can consume it using the gas plugin.
	if err := st.gp.ConsumeGas(gas); err != nil {
		return fmt.Errorf("%w: have %d, want %d", ErrIntrinsicGas, st.gp.GasRemaining(), gas)
	}

	return nil
}
