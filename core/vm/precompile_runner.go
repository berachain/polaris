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

package vm

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"

	coretypes "github.com/berachain/stargazer/core/types"
	"github.com/berachain/stargazer/core/vm/precompile"
	"github.com/berachain/stargazer/core/vm/precompile/container/types"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/errors"
)

// Compile-time assertion to ensure `PrecompileRunner` adheres to `precompile.Runner`.
var _ precompile.Runner = (*PrecompileRunner)(nil)

// `PrecompileRunner` is the execution environment of a precompiled container at a given address.
// The runner manages the execution of the container and emission of Cosmos events to Ethereum
// logs.
type PrecompileRunner struct {
	// `registry` is the registry from which the precompile container will be pulled and precompile
	// logs can be built.
	registry *PrecompileRegistry

	// `psdb` is the precompile StateDB to support state changes during precompile execution.
	psdb PrecompileStateDB
}

// `NewPrecompileRunner` creates and returns a new `PrecompileRunner` for the given precompile
// registry `registry` and precompile StateDB `psdb`.
func NewPrecompileRunner(registry *PrecompileRegistry, psdb PrecompileStateDB) *PrecompileRunner {
	return &PrecompileRunner{
		registry: registry,
		psdb:     psdb,
	}
}

// `Exists` gets a precompile container at the given `addr` from the precompile registry.
//
// `Exists` implements `precompile.Runner`.
func (pr *PrecompileRunner) Exists(addr common.Address) (types.PrecompileContainer, bool) {
	return pr.registry.Get(addr)
}

// `Run` runs the given precompile container and returns the remaining gas after execution. This
// function returns an error if the given statedb is not compatible with precompiles, insufficient
// gas is provided, or the precompile execution returns an error.
//
// `Run` implements `precompile.Runner`.
func (pr *PrecompileRunner) Run(
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

	// TODO: generalize adding logs
	// store the number of events before precompile container execution, to be used as index for
	// building logs for all Cosmos events emitted during execution
	ctx := pr.psdb.GetContext()
	beforeExecutionNumEvents := len(ctx.EventManager().Events())

	ret, err := pc.Run(ctx, input, caller, value, readonly)
	if err != nil {
		return nil, suppliedGas, err
	}

	// We add logs after the precompile container execution to ensure that if the precompile
	// reverts, the logs are not added. This is a design choice.
	// TODO: generalize adding logs, maybe `Execute` should return logs to append. The goal here is
	// to make it so precompile runner does not need to know about Cosmos events.

	// convert all Cosmos events emitted during precompile container execution to logs and add to
	// StateDB
	events := ctx.EventManager().Events()
	for i := beforeExecutionNumEvents; i < len(events); i++ {
		var log *coretypes.Log
		log, err = pr.buildLog(&events[i])
		if err != nil {
			return nil, suppliedGas, err
		}
		pr.psdb.AddLog(log)
	}

	return ret, suppliedGas, nil
}

// `buildLog` builds an Ethereum event log from the given Cosmos event.
func (pr *PrecompileRunner) buildLog(event *sdk.Event) (*coretypes.Log, error) {
	// NOTE: the incoming Cosmos event's `Type` field, converted to CamelCase, should be equal to
	// the Ethereum event's name.
	log := pr.registry.logRegistry.GetPrecompileLog(event.Type)
	if log == nil {
		return nil, errors.Wrapf(precompile.ErrEthEventNotRegistered, "cosmos event %s", event.Type)
	}
	var err error
	if err = log.ValidateAttributes(event); err != nil {
		return nil, errors.Wrapf(precompile.ErrEventHasIssues, "cosmos event %s", event.Type)
	}

	// build Ethereum log based on valid Cosmos event
	eventLog := &coretypes.Log{
		Address: log.GetPrecompileAddress(),
	}
	if eventLog.Topics, err = log.MakeTopics(event); err != nil {
		return nil, errors.Wrapf(precompile.ErrEventHasIssues, "cosmos event %s", event.Type)
	}
	if eventLog.Data, err = log.MakeData(event); err != nil {
		return nil, errors.Wrapf(precompile.ErrEventHasIssues, "cosmos event %s", event.Type)
	}
	return eventLog, nil
}
