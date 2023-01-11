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

package events

import (
	"errors"

	"github.com/berachain/stargazer/common"
	coretypes "github.com/berachain/stargazer/core/types"
	"github.com/berachain/stargazer/core/vm/precompile/types"
	"github.com/berachain/stargazer/types/abi"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type EventType string

// `EthereumLogFactory` builds Ethereum logs from Cosmos events.
type EthereumLogFactory struct {
	// address to signature to event
	precompileEvents map[EventType]*types.PrecompileEvent

	// attributeDecoders is a map of attribute key to decoder function
	attributeDecoders map[string]types.AttributeValueDecoder
}

// ==============================================================================
// Constructor
// ==============================================================================

// `NewEthereumLogFactory` creates and returns a new, empty `EthereumLogFactory` instance.
func NewEthereumLogFactory(
	attributeDecoders map[string]types.AttributeValueDecoder,
) *EthereumLogFactory {
	return &EthereumLogFactory{
		precompileEvents:  make(map[EventType]*types.PrecompileEvent),
		attributeDecoders: attributeDecoders,
	}
}

// ==============================================================================
// Builder
// ==============================================================================

// `BuildEthLog` builds an Eth event log from the given Cosmos event.
func (pef *EthereumLogFactory) BuildEthLog(event *sdk.Event) (*coretypes.Log, error) {
	pe, found := pef.precompileEvents[EventType(abi.ToCamelCase(event.Type))]
	if !found {
		return nil, errors.New("event not found")
	}

	var err error
	ethLog := &coretypes.Log{
		Address: pe.ModuleAddress(),
	}

	if err = pe.ValidateAttributes(event); err != nil {
		return nil, err
	}

	if ethLog.Topics, err = pe.MakeTopicsField(event); err != nil {
		return nil, err
	}

	if ethLog.Data, err = pe.MakeDataField(event); err != nil {
		return nil, err
	}
	return ethLog, nil
}

// ==============================================================================
// EthereumLogFactory - Setter
// ==============================================================================

// `RegisterNewEvent` registers a precompile event for the given event type.
func (pef *EthereumLogFactory) RegisterNewCosmosEvent(
	moduleEthAddress common.Address,
	abiEvent *abi.Event,
) {
	// NOTE: The` abi.CamelCase()` of the Cosmos SDK event's `EventType`
	// corresponding to this abiEvent is assumed to be equal to the abiEvent's `Name()`.
	eventType := EventType(abiEvent.Name)
	pef.precompileEvents[eventType] = types.NewPrecompileEvent(
		moduleEthAddress,
		abiEvent, pef.attributeDecoders,
	)
}
