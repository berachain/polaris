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

package precompile

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"

	coretypes "github.com/berachain/stargazer/core/types"
	"github.com/berachain/stargazer/core/vm/precompile/container/types"
	"github.com/berachain/stargazer/core/vm/precompile/log"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/errors"
)

// Compile-time assertion to ensure `Runner` adheres to `Runner`.
var _ RunnerI = (*Runner)(nil)

// `Runner` is the execution environment of a precompiled container at a given address.
// The host manages the execution of the container and emission of Cosmos events to Ethereum logs.
type Runner struct {
	// `pr` is the registry from which the precompile container will be pulled and precompile logs
	// can be built.
	pr *Registry

	// `psdb` is the precompile StateDB to support state changes during precompile execution.
	psdb PrecompileStateDB
}

// `NewRunner` creates and returns a new `Runner` for the given precompile
// registry `pr` and precompile StateDB `psdb`.
func NewRunner(pr *Registry, psdb PrecompileStateDB) *Runner {
	return &Runner{
		pr:   pr,
		psdb: psdb,
	}
}

// `Exists` gets a precompile container at the given `addr` from the precompile registry.
//
// `Exists` implements `Runner`.
func (ph *Runner) Exists(addr common.Address) (types.PrecompileContainer, bool) {
	return ph.pr.Get(addr)
}

// `Run` runs the given precompile container and returns the remaining gas after execution. This
// function returns an error if the given statedb is not compatible with precompiles, insufficient
// gas is provided, or the precompile execution returns an error.
//
// `Run` implements `Runner`.
func (ph *Runner) Run(
	pc types.PrecompileContainer,
	input []byte,
	caller common.Address,
	value *big.Int,
	suppliedGas uint64,
	readonly bool,
) ([]byte, uint64, error) {
	// TODO: move gas calculation to precompile container using gas meter.
	gasCost := pc.RequiredGas(input)
	if suppliedGas < gasCost {
		return nil, 0, ErrOutOfGas
	}
	suppliedGas -= gasCost

	// todo: generalize adding logs
	// store the number of events before precompile container execution, to be used as index for
	// building logs for all Cosmos events emitted during execution
	ctx := ph.psdb.GetContext()
	beforeExecutionNumEvents := len(sdk.UnwrapSDKContext(ctx).EventManager().Events())

	ret, err := pc.Run(ctx, input, caller, value, readonly)
	if err != nil {
		return nil, suppliedGas, err
	}

	// We add logs after the precompile container execution to ensure that if the precompile reverts,
	// the logs are not added. This is a design choice.

	// todo: generalize adding logs, maybe `Execute` should return logs to append.
	// The goal here is to make it so precompile runner does not need to know about Cosmos events
	// convert all Cosmos events emitted during precompile container execution to logs and add to
	// StateDB
	events := sdk.UnwrapSDKContext(ctx).EventManager().Events()
	for i := beforeExecutionNumEvents; i < len(events); i++ {
		var log *coretypes.Log
		log, err = ph.buildLog(&events[i])
		if err != nil {
			return nil, suppliedGas, err
		}
		ph.psdb.AddLog(log)
	}

	return ret, suppliedGas, nil
}

// `buildLog` builds an Ethereum event log from the given Cosmos event.
func (ph *Runner) buildLog(event *sdk.Event) (*coretypes.Log, error) {
	// NOTE: the incoming Cosmos event's `Type` field, converted to CamelCase, should be equal to
	// the Ethereum event's name.
	_log := ph.pr.Registry.GetPrecompileLog(event.Type)
	if _log == nil {
		return nil, errors.Wrapf(log.ErrEthEventNotRegistered, "cosmos event %s", event.Type)
	}

	var i any = event
	return ph.pr.Registry.Translator.BuildLog(_log, i)
}
