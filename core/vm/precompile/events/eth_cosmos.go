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
	"github.com/berachain/stargazer/types/abi"
)

const maxTopicsLen = 4

type ethCosmosEvent struct {
	address                      *common.Address
	id                           common.Hash
	indexedInputs                abi.Arguments
	nonIndexedInputs             abi.Arguments
	attributeKeysToValueDecoders map[string]AttributeValueDecoder
}

// get the eth Address for an eth event.
func (ece *ethCosmosEvent) getAddress() common.Address {
	return *ece.address
}

// generate Eth Topics for valid cosmos event.
func (ece *ethCosmosEvent) makeTopics(event *sdk.Event) ([]common.Hash, error) {
	filterQuery := make([]any, len(ece.indexedInputs)+1)
	filterQuery[0] = ece.id
	for i := 0; i < len(ece.indexedInputs); i++ {
		input := ece.indexedInputs[i]
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
		valueDecoder, ok := ece.attributeKeysToValueDecoders[attribute.Key]
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

// generate Eth data for valid cosmos event.
func (ece *ethCosmosEvent) generateData(event *sdk.Event) ([]byte, error) {
	attrVals := make([]any, len(ece.nonIndexedInputs))
	// complexity of below iteration: O(n^2), where n is the number of non-indexed args
	for idx, input := range ece.nonIndexedInputs {
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
		valueDecoder, ok := ece.attributeKeysToValueDecoders[attribute.Key]
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
	data, err := ece.nonIndexedInputs.PackValues(attrVals)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Extracts indexed arguments from an Events' inputs. Will panic if more than 3 indexed arguments
// are provided by the inputs ABI.
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
