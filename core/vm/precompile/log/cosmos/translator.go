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

package cosmos

import (
	coretypes "github.com/berachain/stargazer/core/types"
	"github.com/berachain/stargazer/core/vm/precompile/log"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/errors"
	"github.com/berachain/stargazer/types/abi"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ log.Translator[sdk.Event] = (*Translator)(nil)

type Translator struct {
	// `customValueDecoders` is a map of Cosmos attribute keys to attribute value decoder
	// functions for custom modules.
	customValueDecoders ValueDecoders
}

func NewTranslator(customValueDecoders ValueDecoders) *Translator {
	return &Translator{
		customValueDecoders: customValueDecoders,
	}
}

// `BuildLog` builds an Ethereum log from a valid Cosmos event. The Ethereum log is built by
// converting the Cosmos event's attributes into the Ethereum event's indexed and non-indexed
// arguments. The Ethereum log's `Topics` field is built by converting the Cosmos event's
// attributes into the Ethereum event's indexed arguments. The Ethereum log's `Data` field is
// built by converting the Cosmos event's attributes into the Ethereum event's non-indexed
// arguments. The Ethereum log's `Address` field is set to the precompile address of the
// Ethereum event.
func (clf *Translator) BuildLog(log *log.PrecompileLog, event *sdk.Event) (*coretypes.Log, error) {
	var err error
	if err = validateAttributes(log, event); err != nil {
		return nil, errors.Wrapf(ErrEventHasIssues, "cosmos event %s", event.Type)
	}

	// build Ethereum log based on valid Cosmos event
	eventLog := &coretypes.Log{
		Address: log.GetPrecompileAddress(),
	}
	if eventLog.Topics, err = clf.makeTopics(log, event); err != nil {
		return nil, errors.Wrapf(ErrEventHasIssues, "cosmos event %s", event.Type)
	}
	if eventLog.Data, err = clf.makeData(log, event); err != nil {
		return nil, errors.Wrapf(ErrEventHasIssues, "cosmos event %s", event.Type)
	}
	return eventLog, nil
}

// `MakeTopics` generates the Ethereum log `Topics` field for a valid cosmos event. `Topics` is a
// slice of at most 4 hashes, in which the first topic is the Ethereum event's ID. The optional and
// following 3 topics are hashes of the Ethereum event's indexed arguments. This function builds
// this slice of `Topics` by building a filter query of all the corresponding arguments:
// [eventID, indexed_arg1, ...]. Then this query is converted to topics using geth's
// `abi.MakeTopics` function, which outputs hashes of all arguments in the query. The slice of
// hashes is returned.
func (clf *Translator) makeTopics(pl *log.PrecompileLog, event *sdk.Event) ([]common.Hash, error) {
	indexedInputs := pl.IndexedInputs()
	filterQuery := make([]any, len(indexedInputs)+1)
	filterQuery[0] = pl.ID()

	// for each Ethereum indexed argument, get the corresponding Cosmos event attribute and
	// convert to a geth compatible type. NOTE: this iteration has total complexity O(M), where
	// M = average length of atrribute key strings, as length of `indexedInputs` <= 3.
	for i, arg := range pl.IndexedInputs() {
		attrIdx := searchAttributesForArg(&event.Attributes, arg.Name)
		if attrIdx == notFound {
			return nil, errors.Wrap(ErrNoAttributeKeyFound, arg.Name)
		}

		// convert attribute value (string) to geth compatible type
		attr := &event.Attributes[attrIdx]
		decode, err := clf.getValueDecoder(attr.Key)
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

// `MakeData` returns the Ethereum log `Data` field for a valid cosmos event. `Data` is a slice of
// bytes which store an Ethereum event's non-indexed arguments, packed into bytes. This function
// packs the values of the incoming Cosmos event's attributes, which correspond to the
// Ethereum event's non-indexed arguments, into bytes and returns a byte slice.
func (clf *Translator) makeData(pl *log.PrecompileLog, event *sdk.Event) ([]byte, error) {
	nonIndexedInputs := pl.NonIndexedInputs()
	attrVals := make([]any, len(nonIndexedInputs))

	// for each Ethereum non-indexed argument, get the corresponding Cosmos event attribute and
	// convert to a geth compatible type. NOTE: the total complexity of this iteration: O(M*N^2),
	// where N is the # of non-indexed args, M = average length of atrribute key strings.
	for i, arg := range nonIndexedInputs {
		attrIdx := searchAttributesForArg(&event.Attributes, arg.Name)
		if attrIdx == notFound {
			return nil, errors.Wrap(ErrNoAttributeKeyFound, arg.Name)
		}

		// convert attribute value (string) to geth compatible type
		attr := event.Attributes[attrIdx]
		decode, err := clf.getValueDecoder(attr.Key)
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
	data, err := nonIndexedInputs.PackValues(attrVals)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// `getValueDecoder` returns an attribute value decoder function for a certain Cosmos event
// attribute key.
func (clf *Translator) getValueDecoder(attrKey string) (valueDecoder, error) {
	// try custom precompile event attributes
	if clf.customValueDecoders != nil {
		if customDecoder, found := clf.customValueDecoders[attrKey]; found {
			return customDecoder, nil
		}
	}

	// try default Cosmos SDK event attributes
	if defaultDecoder, found := defaultCosmosValueDecoders[attrKey]; found {
		return defaultDecoder, nil
	}

	// no value decoder function was found for attribute key
	return nil, errors.Wrap(ErrNoValueDecoderFunc, attrKey)
}

// `validateAttributes` validates an incoming Cosmos `event`. Specifically, it verifies that the
// number of attributes provided by the Cosmos `event` are adequate for it's corresponding
// Ethereum events.
func validateAttributes(log *log.PrecompileLog, event *sdk.Event) error {
	if len(event.Attributes) < len(log.IndexedInputs())+len(log.NonIndexedInputs()) {
		return ErrNotEnoughAttributes
	}
	return nil
}
