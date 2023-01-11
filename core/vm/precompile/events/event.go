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
	"fmt"

	"github.com/berachain/stargazer/common"
	"github.com/berachain/stargazer/core/types"
	"github.com/berachain/stargazer/types/abi"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// `PrecompileEvent` represents an Eth event emitted by a Cosmos module's precompile contract.
type PrecompileEvent struct {
	// `address` is the Eth address which represents a Cosmos module's account address.
	moduleAddress common.Address

	// `id` is the Eth event ID, to be used as an Eth event's first topic
	id common.Hash

	// `indexedInputs` holds an Eth event's indexed arguments, emitted as event topics
	indexedInputs abi.Arguments

	// `nonIndexedInputs` holds an Eth event's non-indexed arguments, emitted as event data
	nonIndexedInputs abi.Arguments

	// `attributeKeysToValueDecoders` is a map of Cosmos attribute keys to value decoder functions
	valueDecoders map[string]AttributeValueDecoder
}

// `NewPrecompileEvent` returns a new `PrecompileEvent` with the given `moduleAddress` and `abiEvent`.
func NewPrecompileEvent(moduleAddress common.Address, abiEvent *abi.Event) *PrecompileEvent {
	return &PrecompileEvent{
		moduleAddress:    moduleAddress,
		id:               abiEvent.ID,
		indexedInputs:    abi.GetIndexed(abiEvent.Inputs),
		nonIndexedInputs: abiEvent.Inputs.NonIndexed(),
	}
}

// `BuildEthLog` builds the Eth event metadata for a Cosmos event and returns a geth type `Log`
// with the `Address`, `Topics` and `Data` fields filled.
func (pe *PrecompileEvent) BuildEthLog(event *sdk.Event) (*types.Log, error) {
	if len(event.Attributes) <
		len(pe.indexedInputs)+len(pe.nonIndexedInputs) {
		return nil, fmt.Errorf(
			"not enough event attributes provided for event %s",
			event.Type,
		)
	}

	topics, err := pe.makeTopicsField(event)
	if err != nil {
		return nil, err
	}
	data, err := pe.makeDataField(event)
	if err != nil {
		return nil, err
	}
	return &types.Log{
		Address: pe.moduleAddress,
		Topics:  topics,
		Data:    data,
	}, nil
}

// `makeTopics` generates the Eth log `Topics` field for a valid cosmos event.
func (pe *PrecompileEvent) makeTopicsField(event *sdk.Event) ([]common.Hash, error) {
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
		// convert attribute value (string) to common.Hash
		attribute := &event.Attributes[attrIdx]
		valueDecoder, ok := pe.valueDecoders[attribute.Key]
		if !ok {
			return nil, fmt.Errorf(
				"attribute for key %s is not mapped to a value decoder",
				attribute.Key,
			)
		}
		val, err := valueDecoder(attribute.Value)
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

// `generateData` returns the Eth log `Data` for a valid cosmos event.
func (pe *PrecompileEvent) makeDataField(event *sdk.Event) ([]byte, error) {
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
		attribute := event.Attributes[attrIdx]
		valueDecoder, ok := pe.valueDecoders[attribute.Key]
		if !ok {
			return nil, fmt.Errorf(
				"attribute for key %s is not mapped to a value decoder",
				attribute.Key,
			)
		}
		val, err := valueDecoder(attribute.Value)
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
