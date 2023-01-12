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

package event

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/berachain/stargazer/common"
	"github.com/berachain/stargazer/types/abi"
)

// `PrecompileEvent` represents an Ethereum event emitted by a Cosmos module's precompile contract.
type PrecompileEvent struct {
	// `address` is the Ethereum address which represents a Cosmos module's account address.
	moduleAddr common.Address

	// `id` is the Ethereum event ID, to be used as an Ethereum event's first topic.
	id common.Hash

	// `indexedInputs` holds an Ethereum event's indexed arguments, emitted as event topics.
	indexedInputs abi.Arguments

	// `nonIndexedInputs` holds an Ethereum event's non-indexed arguments, emitted as event data.
	nonIndexedInputs abi.Arguments

	// `customValueDecoders` is a map of Cosmos attribute keys to attribute value decoder
	// functions for custom modules.
	customValueDecoders ValueDecoders
}

// `NewPrecompileEvent` returns a new `PrecompileEvent` with the given `moduleAddress`, `abiEvent`,
// and optional `customValueDecoders`.
func NewPrecompileEvent(
	moduleAddr common.Address,
	abiEvent *abi.Event,
	customValueDecoders ...ValueDecoders,
) *PrecompileEvent {
	pe := &PrecompileEvent{
		moduleAddr:       moduleAddr,
		id:               abiEvent.ID,
		indexedInputs:    abi.GetIndexed(abiEvent.Inputs),
		nonIndexedInputs: abiEvent.Inputs.NonIndexed(),
	}
	if len(customValueDecoders) > 0 {
		pe.customValueDecoders = customValueDecoders[0]
	}
	return pe
}

// `ModuleAddress` returns the Ethereum address corresponding to a PrecompileEvent's Cosmos module.
func (pe *PrecompileEvent) ModuleAddress() common.Address {
	return pe.moduleAddr
}

// `MakeTopics` generates the Ethereum log `Topics` field for a valid cosmos event. TODO: explain
func (pe *PrecompileEvent) MakeTopics(event *sdk.Event) ([]common.Hash, error) {
	filterQuery := make([]any, len(pe.indexedInputs)+1)
	filterQuery[0] = pe.id
	for i := 0; i < len(pe.indexedInputs); i++ {
		input := pe.indexedInputs[i]
		// below iteration has insignificant complexity as length of event.Attributes <= 3
		attrIdx := 0
		for ; attrIdx < len(event.Attributes); attrIdx++ {
			if abi.ToMixedCase(event.Attributes[attrIdx].Key) == input.Name {
				break
			}
		}
		if attrIdx == len(event.Attributes) {
			// could not find attribute key in event ABI
			return nil, fmt.Errorf(
				"no attribute key found for event %s argument %s",
				event.Type,
				input.Name,
			)
		}
		// convert attr value (string) to common.Hash
		attr := &event.Attributes[attrIdx]
		valueDecoder, err := pe.getValueDecoder(attributeKey(attr.Key))
		if err != nil {
			return nil, err
		}
		val, err := valueDecoder(attr.Value)
		if err != nil {
			return nil, err
		}
		filterQuery[i+1] = val
	}

	topics, err := abi.MakeTopics(filterQuery)
	if err != nil {
		return nil, err
	}
	return topics[0], nil
}

// `MakeData` returns the Ethereum log `Data` for a valid cosmos event. TODO: explain
func (pe *PrecompileEvent) MakeData(event *sdk.Event) ([]byte, error) {
	attrVals := make([]any, len(pe.nonIndexedInputs))
	// complexity of below iteration: O(n^2), where n is the number of non-indexed args
	for idx, input := range pe.nonIndexedInputs {
		attrIdx := 0
		for ; attrIdx < len(event.Attributes); attrIdx++ {
			if abi.ToMixedCase(event.Attributes[attrIdx].Key) == input.Name {
				break
			}
		}
		if attrIdx == len(event.Attributes) {
			// could not find attribute key in event ABI
			return nil, fmt.Errorf(
				"no attribute key found for event %s argument %s",
				event.Type,
				input.Name,
			)
		}
		// convert each attribute value to geth type
		attr := event.Attributes[attrIdx]
		valueDecoder, err := pe.getValueDecoder(attributeKey(attr.Key))
		if err != nil {
			return nil, err
		}
		val, err := valueDecoder(attr.Value)
		if err != nil {
			return nil, err
		}
		attrVals[idx] = val
	}

	data, err := pe.nonIndexedInputs.PackValues(attrVals)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// `ValidateAttributes` validates an incoming Cosmos `event`. Specifically, it verifies that the
// number of attributes provided by the Cosmos `event` are adequate for it's corresponding
// Ethereum events.
func (pe *PrecompileEvent) ValidateAttributes(event *sdk.Event) error {
	if len(event.Attributes) < len(pe.indexedInputs)+len(pe.nonIndexedInputs) {
		return fmt.Errorf(
			"not enough event attributes provided for event %s",
			event.Type,
		)
	}
	return nil
}

// `getValueDecoder` TODO: explain
func (pe *PrecompileEvent) getValueDecoder(attrKey attributeKey) (attributeValueDecoder, error) {
	// try custom precompile event attributes
	if pe.customValueDecoders != nil {
		if valueDecoder, ok := pe.customValueDecoders[attrKey]; ok {
			return valueDecoder, nil
		}
	}

	// try default Cosmos SDK event attributes
	valueDecoder, ok := defaultCosmosAttributes[attrKey]
	if !ok {
		return nil, fmt.Errorf(
			"attribute for key %s is not mapped to a value decoder",
			attrKey,
		)
	}
	return valueDecoder, nil
}
