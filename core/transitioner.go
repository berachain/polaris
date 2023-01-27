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
	"github.com/berachain/stargazer/core/vm"
	"github.com/berachain/stargazer/lib/errors"
)

// =============================================================================
// Transition Execution
// =============================================================================

// Create a new state transitioner to process a single transaction.
func ApplyMessage(
	evm vm.StargazerEVM,
	msg Message,
) (*ExecutionResult, error) {
	return NewStateTransition(evm, msg).TransitionDB()
}

// Create a new state transitioner to process a single transaction.
func ApplyMessageAndCommit(
	evm vm.StargazerEVM,
	msg Message,
) (*ExecutionResult, error) {
	res, err := NewStateTransition(evm, msg).TransitionDB()
	if err != nil {
		return nil, errors.Wrap(err, "failed to TransitionDB")
	}

	// Persist state.
	if err = evm.StateDB().FinalizeTx(); err != nil {
		return nil, errors.Wrap(err, "failed to commit stateDB")
	}

	return res, nil
}

// Create a new state transitioner to process a single transaction.
func ApplyMessageWithTracer(
	evm vm.StargazerEVM,
	msg Message,
	tracer vm.EVMLogger,
) (*ExecutionResult, error) {
	return NewStateTransition(evm, msg).traceTransitionDB(tracer)
}

// Create a new state transitioner to process a single transaction.
func ApplyMessageWithTracerAndCommit(
	evm vm.StargazerEVM,
	msg Message,
	tracer vm.EVMLogger,
) (*ExecutionResult, error) {
	res, err := ApplyMessageWithTracer(evm, msg, tracer)
	if err != nil {
		return nil, errors.Wrap(err, "failed to TransitionDB")
	}

	// Persist state.
	if err = evm.StateDB().FinalizeTx(); err != nil {
		return nil, errors.Wrap(err, "failed to commit stateDB")
	}
	return res, nil
}
