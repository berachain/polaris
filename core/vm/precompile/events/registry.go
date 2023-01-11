package events

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/berachain/stargazer/common"
	"github.com/berachain/stargazer/core/types"
	"github.com/berachain/stargazer/types/abi"
)

// `Registry` registers and stores all Cosmos events which are supported to be converted to Eth
// event logs during stateful precompile execution.
type Registry struct {
	eventTypesToRelayers map[string]*relayer
}

// `NewRegistry` creates and returns a new, empty `Registry` instance.
func NewRegistry() *Registry {
	return &Registry{
		eventTypesToRelayers: make(map[string]*relayer),
	}
}

// `RegisterModule` registers a Cosmos module's events from the module's stateful precompile
// `contract` at the given Eth address `moduleAddr`.
func (er *Registry) RegisterModule(moduleAddr *common.Address, contract HasEvents) {
	abiEvents := contract.ABIEvents()
	for _, eventType := range contract.CosmosEventTypes() {
		eventTypeCamelCase := abi.ToCamelCase(eventType) // Eth stores events in CamelCase
		event, ok := abiEvents[eventTypeCamelCase]
		if !ok {
			// require that every Eth event has been manually mapped to a corresponding Cosmos one
			panic(fmt.Errorf("no eth event defined for cosmos event type %s", eventType))
		}
		er.eventTypesToRelayers[eventTypeCamelCase] = &relayer{
			address:                      moduleAddr,
			id:                           event.ID,
			indexedInputs:                getIndexed(event.Inputs),
			nonIndexedInputs:             event.Inputs.NonIndexed(),
			attributeKeysToValueDecoders: contract.AttributeKeysToValueDecoders(),
		}
	}
}

// `BuildEthLog` builds the Eth event metadata for a Cosmos event and returns a geth type `Log`
// with the `Address`, `Topics` and `Data` fields filled.
func (er *Registry) BuildEthLog(event *sdk.Event) (*types.Log, error) {
	relayer, err := er.getCosmosEventData(event)
	if err != nil {
		return nil, err
	}
	topics, err := relayer.makeTopics(event)
	if err != nil {
		return nil, err
	}
	data, err := relayer.generateData(event)
	if err != nil {
		return nil, err
	}
	return &types.Log{
		Address: relayer.getAddress(),
		Topics:  topics,
		Data:    data,
	}, nil
}

// `getCosmosEventData` checks that an incoming cosmos event is valid. If valid, it returns the
// Cosmos-to-Eth event relayer. If not valid, this function panics.
func (er *Registry) getCosmosEventData(event *sdk.Event) (*relayer, error) {
	eventKey := abi.ToCamelCase(event.Type)
	relayer, ok := er.eventTypesToRelayers[eventKey]
	if !ok {
		return nil, fmt.Errorf(
			"the Eth event corresponding to Cosmos event %s has not been registered",
			event.Type,
		)
	}
	if len(event.Attributes) <
		len(relayer.indexedInputs)+len(relayer.nonIndexedInputs) {
		return nil, fmt.Errorf(
			"not enough event attributes provided for event %s",
			event.Type,
		)
	}
	return relayer, nil
}
