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

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/berachain/stargazer/common"
	"github.com/berachain/stargazer/types/abi"
)

// `maxTopicsLen` is the maximum number of topics hashes allowed in an Eth log.
const maxTopicsLen = 4

// `cosmosEventRelayer` holds an event's Cosmos and Eth metadata that is used to convert an incoming Cosmos
// event to its corresponding Eth event log.
type cosmosEventRelayer struct {
	// `address` is the Eth address which represents a Cosmos module's account address.
	address *common.Address

	// `id` is the Eth event ID, to be used as an Eth event's first topic
	id common.Hash

	// `indexedInputs` holds an Eth event's indexed arguments, emitted as event topics
	indexedInputs abi.Arguments

	// `nonIndexedInputs` holds an Eth event's non-indexed arguments, emitted as event data
	nonIndexedInputs abi.Arguments

	// `attributeKeysToValueDecoders` is a map of Cosmos attribute keys to value decoder functions
	attributeKeysToValueDecoders map[string]AttributeValueDecoder
}

// `getAddress` returns the Eth address (which represents account address of the event's
// corresponding Cosmos module) for an event.
func (r *cosmosEventRelayer) getAddress() common.Address {
	return *r.address
}

// `makeTopics` generates the Eth log `Topics` field for a valid cosmos event.
func (r *cosmosEventRelayer) makeTopics(event *sdk.Event) ([]common.Hash, error) {
	filterQuery := make([]any, len(r.indexedInputs)+1)
	filterQuery[0] = r.id
	for i := 0; i < len(r.indexedInputs); i++ {
		input := r.indexedInputs[i]
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
		valueDecoder, ok := r.attributeKeysToValueDecoders[attribute.Key]
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
func (r *cosmosEventRelayer) generateData(event *sdk.Event) ([]byte, error) {
	attrVals := make([]any, len(r.nonIndexedInputs))
	// complexity of below iteration: O(n^2), where n is the number of non-indexed args
	for idx, input := range r.nonIndexedInputs {
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
		valueDecoder, ok := r.attributeKeysToValueDecoders[attribute.Key]
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

	data, err := r.nonIndexedInputs.PackValues(attrVals)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// `getIndexed` filters and returns indexed arguments from an Eth event's arguments. This function
// panics if more than 3 indexed arguments are provided in `args`.
func getIndexed(args abi.Arguments) abi.Arguments {
	var indexed abi.Arguments
	numIndexed := 0
	for _, arg := range args {
		if arg.Indexed {
			if numIndexed == maxTopicsLen {
				panic("number of indexed arguments is more than allowed by Eth event log")
			}
			indexed = append(indexed, arg)
			numIndexed++
		}
	}
	return indexed
}
