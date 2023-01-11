package events

import (
	"github.com/berachain/stargazer/types/abi"
)

// `HasEvents` defines the minimum set of required functions for a stateful precompile to implement
// if it wants to emit each of its Cosmos emitted events as an Eth event log.
type HasEvents interface {
	// `CosmosEventTypes` should return a slice of the Cosmos event type strings (under_score)
	// formatting for all Cosmos events the stateful precompile wishes to emit as Eth events.
	CosmosEventTypes() []string

	// `ABIEvents` should return a map of Eth event names (should be CamelCase formatted) to
	// geth abi `Event` structs. Note: this can be directly loaded from the `Events` field of a
	// geth ABI struct, which is built from a solidity library, interface or contract.
	ABIEvents() map[string]abi.Event

	// `AttributeKeysToValueDecoders` should return a map of all Cosmos event attribute keys to
	// `AttributeValueDecoder`s for all supported Cosmos events.
	AttributeKeysToValueDecoders() map[string]AttributeValueDecoder
}
