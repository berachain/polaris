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

package log

import (
	"github.com/berachain/stargazer/eth/types/abi"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// `makeTopics` generates the Ethereum log `Topics` field for a valid cosmos event. `Topics` is a
// slice of at most 4 hashes, in which the first topic is the Ethereum event's ID. The optional and
// following 3 topics are hashes of the Ethereum event's indexed arguments. This function builds
// this slice of `Topics` by building a filter query of all the corresponding arguments:
// [eventID, indexed_arg1, ...]. Then this query is converted to topics using geth's
// `abi.MakeTopics` function, which outputs hashes of all arguments in the query. The slice of
// hashes is returned.
func (f *Factory) makeTopics(pl *precompileLog, event *sdk.Event) ([]common.Hash, error) {
	filterQuery := make([]any, len(pl.indexedInputs)+1)
	filterQuery[0] = pl.id

	// for each Ethereum indexed argument, get the corresponding Cosmos event attribute and
	// convert to a geth compatible type. NOTE: this iteration has total complexity O(M), where
	// M = average length of atrribute key strings, as length of `indexedInputs` <= 3.
	for i, arg := range pl.indexedInputs {
		attrIdx := searchAttributesForArg(&event.Attributes, arg.Name)
		if attrIdx == notFound {
			return nil, errors.Wrap(ErrNoAttributeKeyFound, arg.Name)
		}

		// convert attribute value (string) to geth compatible type
		attr := &event.Attributes[attrIdx]
		decode, err := f.getValueDecoder(attr.Key)
		if err != nil {
			return nil, err
		}
		value, err := decode(attr.Value)
		if err != nil {
			return nil, err
		}
		filterQuery[i+1] = value
	}

	// convert the filter query to a slice of `Topics` hashes
	topics, err := abi.MakeTopics(filterQuery)
	if err != nil {
		return nil, err
	}
	return topics[0], nil
}

// `makeData` returns the Ethereum log `Data` field for a valid cosmos event. `Data` is a slice of
// bytes which store an Ethereum event's non-indexed arguments, packed into bytes. This function
// packs the values of the incoming Cosmos event's attributes, which correspond to the
// Ethereum event's non-indexed arguments, into bytes and returns a byte slice.
func (f *Factory) makeData(pl *precompileLog, event *sdk.Event) ([]byte, error) {
	attrVals := make([]any, len(pl.nonIndexedInputs))

	// for each Ethereum non-indexed argument, get the corresponding Cosmos event attribute and
	// convert to a geth compatible type. NOTE: the total complexity of this iteration: O(M*N^2),
	// where N is the # of non-indexed args, M = average length of atrribute key strings.
	for i, arg := range pl.nonIndexedInputs {
		attrIdx := searchAttributesForArg(&event.Attributes, arg.Name)
		if attrIdx == notFound {
			return nil, errors.Wrap(ErrNoAttributeKeyFound, arg.Name)
		}

		// convert attribute value (string) to geth compatible type
		attr := event.Attributes[attrIdx]
		decode, err := f.getValueDecoder(attr.Key)
		if err != nil {
			return nil, err
		}
		value, err := decode(attr.Value)
		if err != nil {
			return nil, err
		}
		attrVals[i] = value
	}

	// pack the Cosmos event's attribute values into bytes
	data, err := pl.nonIndexedInputs.PackValues(attrVals)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// `getValueDecoder` returns an attribute value decoder function for a certain Cosmos event
// attribute key.
func (f *Factory) getValueDecoder(attrKey string) (valueDecoder, error) {
	// try custom precompile event attributes
	if customDecoder, found := f.customValueDecoders[attrKey]; found {
		return customDecoder, nil
	}

	// try default Cosmos SDK event attributes
	if defaultDecoder, found := defaultCosmosValueDecoders[attrKey]; found {
		return defaultDecoder, nil
	}

	// no value decoder function was found for attribute key
	return nil, errors.Wrap(ErrNoValueDecoderFunc, attrKey)
}
