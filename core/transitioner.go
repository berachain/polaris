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
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/berachain/stargazer/core/vm"
)

// The StateTransitioner is responsible for executing state transtions.
// It also caches EVM config / param variables to prevent having to pull the parameters from the
// store on every transaction.
type StateTransitioner struct{}

// =============================================================================
// Transition Execution
// =============================================================================

// Create a new state transitioner to process a single transaction.
func (str *StateTransitioner) ApplyMessage(
	evm *vm.StargazerEVM,
	msg Message,
) (*ExecutionResult, error) {
	return str.buildStateTransition(evm, msg).transitionDB()
}

// Create a new state transitioner to process a single transaction.
func (str *StateTransitioner) ApplyMessageAndCommit(
	evm *vm.StargazerEVM,
	msg Message,
) (*ExecutionResult, error) {
	res, err := str.buildStateTransition(evm, msg).transitionDB()
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to transitionDB")
	}

	// Persist state.
	if err = evm.StateDB().FinalizeTx(); err != nil {
		return nil, sdkerrors.Wrap(err, "failed to commit stateDB")
	}

	return res, nil
}

// Create a new state transitioner to process a single transaction.
func (str *StateTransitioner) ApplyMessageWithTracer(
	evm *vm.StargazerEVM,
	msg Message,
	tracer vm.EVMLogger,
) (*ExecutionResult, error) {
	return str.buildStateTransition(evm, msg).traceTransitionDB(tracer)
}

// Create a new state transitioner to process a single transaction.
func (str *StateTransitioner) ApplyMessageWithTracerAndCommit(
	evm *vm.StargazerEVM,
	msg Message,
	tracer vm.EVMLogger,
) (*ExecutionResult, error) {
	res, err := str.ApplyMessageWithTracer(evm, msg, tracer)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to transitionDB")
	}

	// Persist state.
	if err = evm.StateDB().FinalizeTx(); err != nil {
		return nil, sdkerrors.Wrap(err, "failed to commit stateDB")
	}
	return res, nil
}

// Create a new state transitioner to process a single transaction.
func (str *StateTransitioner) buildStateTransition(
	evm *vm.StargazerEVM,
	msg Message,
) *StateTransition {
	return NewStateTransition(evm, msg)
}

// //=============================================================================
// // Expose Params Provider
// //=============================================================================

// // GetParamsProvider returns the params provider.
// func (str *StateTransitioner) GetParamsProvider() *params.Provider {
// 	return str.pp
// }

// func (str *StateTransitioner) GetPetrolStation() *PetrolStation {
// 	return &str.ps
// }

// func (str *StateTransitioner) GetStateDBFactory() *state.StateDBFactory {
// 	return str.sdf
// }
