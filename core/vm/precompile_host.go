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

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	coretypes "github.com/berachain/stargazer/core/types"
	"github.com/berachain/stargazer/core/vm/precompile"
	"github.com/berachain/stargazer/core/vm/precompile/container/types"
	"github.com/berachain/stargazer/lib/common"
)

// Compile-time assertion to ensure `PrecompileHost` adheres to `precompile.Host`.
var _ precompile.Host = (*PrecompileHost)(nil)

// `PrecompileHost` is the execution environment of a precompiled container at a given address.
// The host manages the execution of the container and emission of Cosmos events to Ethereum logs.
type PrecompileHost struct {
	// `pr` is the registry from which the precompile container will be pulled and precompile logs
	// can be built.
	pr *PrecompileRegistry

	// `psdb` is the precompile StateDB to support state changes during precompile execution.
	psdb PrecompileStateDB
}

// `NewPrecompileHost` creates and returns a new `PrecompileHost` for the given precompile
// registry `pr` and precompile StateDB `psdb`.
func NewPrecompileHost(pr *PrecompileRegistry, psdb PrecompileStateDB) *PrecompileHost {
	return &PrecompileHost{
		pr:   pr,
		psdb: psdb,
	}
}

// `Exists` gets a precompile container at the given `addr` from the precompile registry.
//
// `Exists` implements `precompile.Host`.
func (ph *PrecompileHost) Exists(addr common.Address) (types.PrecompileContainer, bool) {
	return ph.pr.Get(addr)
}

// `Run` runs the given precompile container and returns the remaining gas after execution. This
// function returns an error if the given statedb is not compatible with precompiles, insufficient
// gas is provided, or the precompile execution returns an error.
//
// `Run` implements `precompile.Host`.
func (ph *PrecompileHost) Run(
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

	// store the number of events before precompile container execution, to be used as index for
	// building logs for all Cosmos events emitted during execution
	ctx := ph.psdb.GetContext()
	beforeExecutionNumEvents := len(ctx.EventManager().Events())
	ret, err := pc.Run(ctx, input, caller, value, readonly)
	if err != nil {
		return nil, suppliedGas, err
	}

	// convert all Cosmos events emitted during precompile container execution to logs and add to
	// StateDB
	events := ctx.EventManager().Events()
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
func (ph *PrecompileHost) buildLog(event *sdk.Event) (*coretypes.Log, error) {
	// validate incoming Cosmos event
	pe := ph.pr.logRegistry.GetPrecompileLog(event)
	if pe == nil {
		return nil, errors.Wrap(precompile.ErrEthEventNotRegistered, event.Type)
	}
	var err error
	if err = pe.ValidateAttributes(event); err != nil {
		return nil, errors.Wrapf(err, "Cosmos event %s has issue", event.Type)
	}

	// build Ethereum log based on valid Cosmos event
	eventLog := &coretypes.Log{
		Address: pe.ModuleAddress(),
	}
	if eventLog.Topics, err = pe.MakeTopics(event); err != nil {
		return nil, errors.Wrapf(err, "Cosmos event %s has issue", event.Type)
	}
	if eventLog.Data, err = pe.MakeData(event); err != nil {
		return nil, errors.Wrapf(err, "Cosmos event %s has issue", event.Type)
	}
	return eventLog, nil
}
