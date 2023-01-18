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

	"github.com/berachain/stargazer/core/state"
	coretypes "github.com/berachain/stargazer/core/types"
	"github.com/berachain/stargazer/core/vm/precompile"
	"github.com/berachain/stargazer/core/vm/precompile/container/types"
	"github.com/berachain/stargazer/lib/common"
)

// Compile-time assertion to ensure `PrecompileHost` adheres to `precompile.Host`.
var _ precompile.Host = (*PrecompileHost)(nil)

// `PrecompileHost` is gets and executes a precompiled container at a given address.
type PrecompileHost struct {
	// `pr` is the registry from which the precompile container will be pulled.
	pr *PrecompileRegistry

	// `lr` is the registry of Cosmos events that can be emitted as Ethereum logs.
	lr *precompile.LogRegistry
}

// `NewPrecompileHost` creates and returns a new `PrecompileHost` for the given precompile
// registry `pr` and log registry `lr`.
func NewPrecompileHost(pr *PrecompileRegistry, lr *precompile.LogRegistry) *PrecompileHost {
	return &PrecompileHost{
		pr: pr,
		lr: lr,
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
	sdb state.GethStateDB,
	input []byte,
	caller common.Address,
	value *big.Int,
	suppliedGas uint64,
	readonly bool,
) ([]byte, uint64, error) {
	psdb, ok := sdb.(state.PrecompileStateDB)
	if !ok {
		return nil, 0, ErrStateDBNotSupported
	}

	// TODO: move gas calculation to precompile container using gas meter.
	gasCost := pc.RequiredGas(input)
	if suppliedGas < gasCost {
		return nil, 0, ErrOutOfGas
	}
	suppliedGas -= gasCost

	// store the number of events before container execution, to be used as index for building logs
	// for all Cosmos events emitted during execution
	ctx := psdb.GetContext()
	beforeExecutionNumEvents := len(ctx.EventManager().Events())
	ret, err := pc.Run(ctx, input, caller, value, readonly)
	if err != nil {
		return nil, suppliedGas, err
	}

	// convert all Cosmos events emitted during container execution to logs and add to sdb
	events := ctx.EventManager().Events()
	for i := beforeExecutionNumEvents; i < len(events); i++ {
		var log *coretypes.Log
		log, err = ph.buildLog(&events[i])
		if err != nil {
			return nil, suppliedGas, err
		}
		psdb.AddLog(log)
	}

	return ret, suppliedGas, nil
}

// `buildLog` builds an Ethereum event log from the given Cosmos event.
func (ph *PrecompileHost) buildLog(event *sdk.Event) (*coretypes.Log, error) {
	// validate incoming Cosmos event
	pe := ph.lr.GetPrecompileLog(event)
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
