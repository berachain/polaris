// Copyright (C) 2022, Berachain Foundation. All rights reserved.
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
	"fmt"

	"github.com/berachain/stargazer/core/vm"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/params"
)

type StateTransition struct {
	// An instance of the  Virtual Machine
	evm *vm.StargazerEVM

	// The message to deliver to the EVM
	msg Message

	// Gas consumption tracking
	gas        uint64
	initialGas uint64
}

// =============================================================================
// Constructor
// =============================================================================

// NewStateTransition creates a new state transition object.
func NewStateTransition(evm *vm.StargazerEVM, msg Message) *StateTransition {
	// Configure transaction config from the message.
	evm.TxContext.Origin = msg.From()
	evm.TxContext.GasPrice = msg.GasPrice()

	return &StateTransition{
		evm:        evm,
		msg:        msg,
		gas:        msg.Gas(),
		initialGas: msg.Gas(),
	}
}

// =============================================================================
// Low Level Transition w/State Machine
// =============================================================================

// `transitionDB` executes the configured message in the Ethereum Virtual Machine (EVM) and
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
func (st *StateTransition) transitionDB() (*ExecutionResult, error) {
	var (
		msgFrom  = st.msg.From()
		msgValue = st.msg.Value()
		ctx      = st.evm.Context
		msgData  = st.msg.Data()
		sender   = vm.AccountRef(msgFrom)
		rules    = st.evm.ChainConfig().Rules(
			ctx.BlockNumber,
			ctx.Random != nil,
		)
		sdb              = st.evm.StateDB()
		contractCreation = st.msg.To() == nil
	)

	gas, err := EthIntrinsicGas(msgData, st.msg.AccessList(),
		contractCreation, rules.IsHomestead, rules.IsIstanbul)
	if err != nil {
		return nil, err
	}
	if st.gas < gas {
		return nil, fmt.Errorf("%w: have %d, want %d", ErrIntrinsicGas, st.gas, gas)
	}
	st.gas -= gas

	// Check to ensure the sender has the funds to cover the value being sent.
	if msgValue.Sign() > 0 && !ctx.CanTransfer(sdb, msgFrom, msgValue) {
		return nil, fmt.Errorf("%w: address %v", ErrInsufficientFundsForTransfer, msgFrom.Hex())
	}

	// if rules.IsBerlin {
	// 	sdb.PrepareAccessList(
	// 		msgFrom,
	// 		st.msg.To(),
	// 		st.evm.ActivePrecompiles(rules),
	// 		st.msg.AccessList(),
	// 	)
	// }

	var (
		ret   []byte // return bytes from evm execution
		vmErr error  // vm errors don't effect consensus and are therefore not passed to err
	)

	// take over the nonce management from evm:
	// - reset sender's nonce to msg.Nonce() before calling evm.
	// - increase sender's nonce by one no matter the result.
	// - this is probably not required, but adds a safety measure,
	//   to ensure that the nonce is getting updated correctly.
	//
	if contractCreation {
		// take over from the evm, as to not
		// increment the nonce twice.
		// TODO: Review nonce accounting.
		sdb.SetNonce(sender.Address(), st.msg.Nonce())
		ret, _, st.gas, vmErr = st.evm.Create(sender,
			msgData, st.gas, msgValue)
		sdb.SetNonce(sender.Address(), st.msg.Nonce()+1)
	} else {
		ret, st.gas, vmErr = st.evm.Call(sender, st.to(),
			msgData, st.gas, msgValue)
	}

	if !rules.IsLondon {
		// Before EIP-3529: refunds were capped to gasUsed / 2
		st.refundGas(params.RefundQuotient)
	} else {
		// After EIP-3529: refunds are capped to gasUsed / 5
		st.refundGas(params.RefundQuotientEIP3529)
	}

	return &ExecutionResult{
		UsedGas:    st.gasUsed(),
		Err:        vmErr,
		ReturnData: ret,
		// Logs:       st.evm.StateDB.Logs(),
	}, nil
}

// `traceTransitionDB` is wrapper around `transitionDB` that adds a tracer to the EVM
// and switches it to debug mode. The tracer is used to capture the execution trace
// of the message in the EVM. After execution it captures the gas remaining and
// returns the execution result, while also setting the EVM back to non-debug mode.
func (st *StateTransition) traceTransitionDB(tracer vm.EVMLogger) (*ExecutionResult, error) {
	// Add a safety check to ensure that the tracer is not nil, as this will cause
	// a panic in the EVM.
	if tracer == nil {
		return nil, fmt.Errorf("invalid tracer")
	}

	// Apply the supplied tracer to the EVM as well as switch it to debug mode.
	st.evm.Config.Tracer = tracer
	st.evm.Config.Debug = true

	// Capture the starting gas for the tracer, we can skip the check for debug mode that is
	// present in geth, as we already know that the EVM is in debug mode from the lines above.
	st.evm.Config.Tracer.CaptureTxStart(st.initialGas)
	defer func() {
		// After execution is completed we need to capture gas remaining.
		st.evm.Config.Tracer.CaptureTxEnd(st.gas)
		// We also take the EVM out of debug mode as this allows us to optimize the normal
		// execution mode by being able to skip setting debug to false in that code path.
		st.evm.Config.Debug = false
	}()

	// Perform the state machine execution
	return st.transitionDB()
}

func (st *StateTransition) gasUsed() uint64 {
	return st.initialGas - st.gas
}

func (st *StateTransition) refundGas(refundQuotient uint64) {
	sdb := st.evm.StateDB()
	// Apply refund counter, capped to a refund quotient
	refund := st.gasUsed() / refundQuotient
	if refund > sdb.GetRefund() {
		refund = sdb.GetRefund()
	}
	st.gas += refund

	// In Geth, we would have refunded the cost of the unused gas to the sender here.
	// However, in <NAME> we do this in the StateProcessor, since currently, gas fees
	// are deducted in the AnteHandler and not in TransitionDB.

	// TODO: we could potentially add the gas cost refund here, since we do have access to a
	// bank keeper from the statedb. Though it really doesn't matter since unless we are calling
	// this in a block, none of the state is persisted anyways.
	// Ante Handler + Refund in StateProcessor, does sort of make more sense, since we only do
	// the coin math during a block and not on queries.

	// Moving buyGas and refundGas to here however... would open the door to potentially using
	// the Geth/Erigon state transition code, which would be nice. We would then just do no
	// gas fee deduction in the AnteHandler, as the native state transition does that.
}

// to returns the recipient of the message.
func (st *StateTransition) to() common.Address {
	if st.msg == nil || st.msg.To() == nil /* contract creation */ {
		return common.Address{}
	}
	return *st.msg.To()
}
