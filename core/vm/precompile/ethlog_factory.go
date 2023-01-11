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
	"github.com/berachain/stargazer/core/vm/precompile/types"
	"github.com/berachain/stargazer/types/abi"
)

// `EventType` is the name of an Ethereum event, which is equivalent to the CamelCase version of
// its corresponding Cosmos event's `Type`.
type EventType string

// `EthereumLogFactory` builds Ethereum logs from Cosmos events.
type EthereumLogFactory struct {
	// `precompileEvents` is a map of `EventType`s to `*types.PrecompileEvents` for all supported
	// Cosmos events.
	precompileEvents map[EventType]*types.PrecompileEvent
}

// ==============================================================================
// Constructor
// ==============================================================================

// `NewEthereumLogFactory` creates and returns a new, empty `EthereumLogFactory`.
func NewEthereumLogFactory() *EthereumLogFactory {
	return &EthereumLogFactory{
		precompileEvents: make(map[EventType]*types.PrecompileEvent),
	}
}

// ==============================================================================
// Setter
// ==============================================================================

// `RegisterNewEvent` registers a precompile event at `moduleEthAddress` for the given `abiEvent`
// and `attributeDecoders`, its corresponding Cosmos attribute keys to value decoder functions.
func (pef *EthereumLogFactory) RegisterNewEvent(
	moduleEthAddress common.Address,
	abiEvent *abi.Event,
	attributeDecoders map[string]types.AttributeValueDecoder,
) {
	// NOTE: The CamelCase version of the Cosmos SDK event's `Type` string corresponding to this
	// abiEvent is assumed to be equal to the abiEvent's `Name` field.
	pef.precompileEvents[EventType(abiEvent.Name)] = types.NewPrecompileEvent(
		moduleEthAddress,
		abiEvent,
		attributeDecoders,
	)
}

// ==============================================================================
// Builder
// ==============================================================================

// `BuildLog` builds an Ethereum event log from the given Cosmos event.
func (pef *EthereumLogFactory) BuildLog(event *sdk.Event) (*coretypes.Log, error) {
	// validate incoming Cosmos event
	// NOTE: the incoming Cosmos event's `Type` field, converted to CamelCase, should be equal to
	// the Ethereum event's name.
	pe, found := pef.precompileEvents[EventType(abi.ToCamelCase(event.Type))]
	if !found {
		return nil, fmt.Errorf(
			"the Ethereum event corresponding to Cosmos event %s has not been registered",
			event.Type,
		)
	}
	if err = pe.ValidateAttributes(event); err != nil {
		return nil, err
	}

	// build Ethereum log based on valid Cosmos event
	var err error
	ethLog := &coretypes.Log{
		Address: pe.ModuleAddress(),
	}
	if ethLog.Topics, err = pe.MakeTopics(event); err != nil {
		return nil, err
	}
	if ethLog.Data, err = pe.MakeData(event); err != nil {
		return nil, err
	}
	return ethLog, nil
}
