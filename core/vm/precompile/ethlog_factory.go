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
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/berachain/stargazer/common"
	coretypes "github.com/berachain/stargazer/core/types"
	"github.com/berachain/stargazer/core/vm/precompile/event"
	"github.com/berachain/stargazer/types/abi"
)

// `EthereumLogFactory` builds Ethereum logs from Cosmos events.
type EthereumLogFactory struct {
	// `precompileEvents` is a map of `EventType`s to `*types.PrecompileEvents` for all supported
	// Cosmos events.
	precompileEvents map[EventType]*event.PrecompileEvent
}

// ==============================================================================
// Constructor
// ==============================================================================

// `NewEthereumLogFactory` creates and returns a new, empty `EthereumLogFactory`.
func NewEthereumLogFactory() *EthereumLogFactory {
	return &EthereumLogFactory{
		precompileEvents: make(map[EventType]*event.PrecompileEvent),
	}
}

// ==============================================================================
// Setter
// ==============================================================================

// `RegisterEvent` registers a Cosmos module's precompile event.
func (pef *EthereumLogFactory) RegisterEvent(
	moduleEthAddress common.Address,
	abiEvent abi.Event,
	customModuleAttributes event.ValueDecoders,
) {
	pef.precompileEvents[EventType(abiEvent.Name)] = event.NewPrecompileEvent(
		moduleEthAddress,
		abiEvent,
		customModuleAttributes,
	)
}

// ==============================================================================
// Builder
// ==============================================================================

// `BuildLog` builds an Ethereum event log from the given Cosmos event.
func (pef *EthereumLogFactory) BuildLog(event *sdk.Event) (*coretypes.Log, error) {
	// validate incoming Cosmos event
	pe := pef.precompileEvents[EventType(abi.ToCamelCase(event.Type))]
	// NOTE: the incoming Cosmos event's `Type` field, converted to CamelCase, should be equal to
	// the Ethereum event's name.
	if pe == nil {
		return nil, fmt.Errorf(
			"the Ethereum event corresponding to Cosmos event %s was not registered",
			event.Type,
		)
	}
	var err error
	if err = pe.ValidateAttributes(event); err != nil {
		return nil, fmt.Errorf("%s for event %s", err.Error(), event.Type)
	}

	// build Ethereum log based on valid Cosmos event
	log := &coretypes.Log{
		Address: pe.ModuleAddress(),
	}
	if log.Topics, err = pe.MakeTopics(event); err != nil {
		return nil, fmt.Errorf("%s for event %s", err.Error(), event.Type)
	}
	if log.Data, err = pe.MakeData(event); err != nil {
		return nil, fmt.Errorf("%s for event %s", err.Error(), event.Type)
	}
	return log, nil
}
