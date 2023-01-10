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

package events

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/berachain/stargazer/common"
	"github.com/berachain/stargazer/core/types"
	"github.com/berachain/stargazer/types/abi"
)

// `EventName` is a Cosmos event name.
type EventName string

// `Registry` is a registry of events emitted by Cosmos modules that
// should be relayed to Ethereum.
type Registry struct {
	eventTypesToLoaders map[EventName]*ethCosmosEvent
}

// `NewRegistry` returns a new Registry instance.
func NewRegistry() *Registry {
	return &Registry{
		eventTypesToLoaders: make(map[EventName]*ethCosmosEvent),
	}
}

// `RegisterModule` registers a Cosmos module with the event registry. This
// function should be called for every Cosmos module that emits events that
// should be relayed to Ethereum.
func (er *Registry) RegisterModule(moduleAddr *common.Address, contract HasEvents) {
	abiEvents := contract.ABIEvents()
	for _, eventType := range contract.CosmosEventTypes() {
		eventName := abi.ToCamelCase(eventType)
		event, ok := abiEvents[eventName]
		if !ok {
			// require that every Eth event has been manually mapped to a corresponding Cosmos one
			panic(fmt.Errorf("no eth event defined for cosmos event type %s", eventType))
		}
		er.eventTypesToLoaders[EventName(eventName)] = &ethCosmosEvent{
			address:                      moduleAddr,
			id:                           event.ID,
			indexedInputs:                getIndexed(event.Inputs),
			nonIndexedInputs:             event.Inputs.NonIndexed(),
			attributeKeysToValueDecoders: contract.AttributeKeysToValueDecoder(),
		}
	}
}

// `BuildEthLog` serializes event metadata from an incoming Cosmos Event into an
// Ethereum formatted Log.
func (er *Registry) BuildEthLog(event *sdk.Event) (*types.Log, error) {
	ethCosmosEvent, err := er.getCosmosEventData(event)
	if err != nil {
		return nil, err
	}
	topics, err := ethCosmosEvent.makeTopics(event)
	if err != nil {
		return nil, err
	}
	data, err := ethCosmosEvent.generateData(event)
	if err != nil {
		return nil, err
	}
	return &types.Log{
		Address: ethCosmosEvent.getAddress(),
		Topics:  topics,
		Data:    data,
	}, nil
}

// `getCosmosEventData` returns the `ethCosmosEvent` corresponding to the Cosmos
// event type that was passed in.
func (er *Registry) getCosmosEventData(event *sdk.Event) (*ethCosmosEvent, error) {
	ethCosmosEvent, ok := er.eventTypesToLoaders[EventName(abi.ToCamelCase(event.Type))]
	if !ok {
		return nil, fmt.Errorf(
			"the Eth event corresponding to Cosmos event %s has not been registered",
			event.Type,
		)
	}
	if len(event.Attributes) <
		len(ethCosmosEvent.indexedInputs)+len(ethCosmosEvent.nonIndexedInputs) {
		return nil, fmt.Errorf(
			"not enough event attributes provided for event %s",
			event.Type,
		)
	}
	return ethCosmosEvent, nil
}
