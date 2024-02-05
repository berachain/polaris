// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package log

import (
	"github.com/berachain/polaris/eth/accounts/abi"
	"github.com/berachain/polaris/eth/core/precompile"
	"github.com/berachain/polaris/lib/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
)

// makeTopics generates the Ethereum log `Topics` field for a valid cosmos event. `Topics` is a
// slice of at most 4 hashes, in which the first topic is the Ethereum event's ID. The optional and
// following 3 topics are hashes of the Ethereum event's indexed arguments. This function builds
// this slice of `Topics` by building a filter query of all the corresponding arguments:
// [eventID, indexed_arg1, ...]. Then this query is converted to topics using geth's
// abi.MakeTopics function, which outputs hashes of all arguments in the query. The slice of
// hashes is returned.
func (f *Factory) makeTopics(pl *precompileLog, event *sdk.Event) ([]common.Hash, error) {
	filterQuery := make([]any, len(pl.indexedInputs)+1)
	filterQuery[0] = pl.id

	// for each Ethereum indexed argument, get the corresponding Cosmos event attribute and
	// convert to a geth compatible type. NOTE: this iteration has total complexity O(M), where
	// M = average length of attribute key strings, as length of `indexedInputs` <= 3.
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

// makeData returns the Ethereum log `Data` field for a valid cosmos event. `Data` is a slice of
// bytes which store an Ethereum event's non-indexed arguments, packed into bytes. This function
// packs the values of the incoming Cosmos event's attributes, which correspond to the
// Ethereum event's non-indexed arguments, into bytes and returns a byte slice.
func (f *Factory) makeData(pl *precompileLog, event *sdk.Event) ([]byte, error) {
	attrVals := make([]any, len(pl.nonIndexedInputs))

	// for each Ethereum non-indexed argument, get the corresponding Cosmos event attribute and
	// convert to a geth compatible type. NOTE: the total complexity of this iteration: O(M*N^2),
	// where N is the # of non-indexed args, M = average length of attribute key strings.
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

// getValueDecoder returns an attribute value decoder function for a certain Cosmos event
// attribute key.
func (f *Factory) getValueDecoder(attrKey string) (precompile.ValueDecoder, error) {
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
