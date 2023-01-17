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
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	coretypes "github.com/berachain/stargazer/core/types"
	"github.com/berachain/stargazer/core/vm/precompile/log"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/types/abi"
)

// `LogFactory` builds Ethereum logs from Cosmos events.
type LogFactory struct {
	// `events` is a map of `EventType`s to `*types.PrecompileEvents` for all supported Cosmos
	// events.
	events map[EventType]*log.PrecompileLog
}

// ==============================================================================
// Constructor
// ==============================================================================

// `NewLogFactory` creates and returns a new, empty `LogFactory`.
func NewLogFactory() *LogFactory {
	return &LogFactory{
		events: make(map[EventType]*log.PrecompileLog),
	}
}

// ==============================================================================
// Setter
// ==============================================================================

// `RegisterEvent` registers a Cosmos module's precompile log.
func (pef *LogFactory) RegisterEvent(
	moduleEthAddress common.Address,
	abiEvent abi.Event,
	customModuleAttributes log.ValueDecoders,
) error {
	if _, found := pef.events[EventType(abiEvent.Name)]; found {
		return errors.Wrap(ErrEthEventAlreadyRegistered, abiEvent.Name)
	}

	var err error
	pef.events[EventType(abiEvent.Name)], err = log.NewPrecompileLog(
		moduleEthAddress,
		abiEvent,
		customModuleAttributes,
	)
	return err
}

// ==============================================================================
// Builder
// ==============================================================================

// `BuildLog` builds an Ethereum event log from the given Cosmos log.
func (pef *LogFactory) BuildLog(event *sdk.Event) (*coretypes.Log, error) {
	// validate incoming Cosmos event
	pe := pef.events[EventType(abi.ToCamelCase(event.Type))]
	// NOTE: the incoming Cosmos event's `Type` field, converted to CamelCase, should be equal to
	// the Ethereum event's name.
	if pe == nil {
		return nil, errors.Wrap(ErrEthEventNotRegistered, event.Type)
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
