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
	"strconv"

	cosmlib "github.com/berachain/polaris/cosmos/lib"
	"github.com/berachain/polaris/eth/accounts/abi"
	"github.com/berachain/polaris/eth/core/precompile"

	abci "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/ethereum/go-ethereum/common"
)

const (
	// intBase is the base `int`s are parsed in, 10.
	intBase = 10
	// int64Bits is the number of bits stored in a variable of `int64` type.
	int64Bits = 64
	// notFound is a default return value for searches in which an item was not found.
	notFound = -1
)

// ==============================================================================
// Default Attribute Value Decoder Getter
// ==============================================================================

// defaultCosmosValueDecoders is a map of default Cosmos event attribute value decoder functions
// for the default Cosmos SDK event `attributeKey`s. NOTE: only the event attributes of default
// Cosmos SDK modules (bank, staking) are supported by this function.
var defaultCosmosValueDecoders = precompile.ValueDecoders{
	sdk.AttributeKeyAmount:                  ConvertSdkCoins,
	stakingtypes.AttributeKeyCreationHeight: ConvertInt64,
	govtypes.AttributeKeyProposalID:         ConvertUint64,
	govtypes.AttributeKeyProposalMessages:   ReturnStringAsIs,
	govtypes.AttributeKeyOption:             ReturnStringAsIs,
}

// ==============================================================================
// Default Attribute Value Decoder Functions
// ==============================================================================

// Compile-time assertions to ensure that the default attribute value decoder functions are
// valueDecoders.
var (
	_ precompile.ValueDecoder = ConvertSdkCoins
	_ precompile.ValueDecoder = ConvertInt64
	_ precompile.ValueDecoder = ConvertUint64
	_ precompile.ValueDecoder = ReturnStringAsIs
	_ precompile.ValueDecoder = ConvertCommonHexAddress
)

// ConvertSdkCoins converts the string representation of an `sdk.Coins`
// to a `[]generated.CosmosCoin`.
// ConvertSdkCoins is a `precompile.ValueDecoder`.
func ConvertSdkCoins(attributeValue string) (any, error) {
	// extract the sdk.Coins from string value
	coins, err := sdk.ParseCoinsNormalized(attributeValue)
	if err != nil {
		return nil, err
	}
	// convert to geth compatible coins
	evmCoins := cosmlib.SdkCoinsToEvmCoins(coins)
	return evmCoins, nil
}

// ConvertInt64 converts a creation height (from the Cosmos SDK staking module) `string`
// to an `int64`.
//
// ConvertInt64 is a `precompile.ValueDecoder`.
func ConvertInt64(attributeValue string) (any, error) {
	return strconv.ParseInt(attributeValue, intBase, int64Bits)
}

// ConvertInt64 converts a `string` to an `int64`.
//
// ConvertInt64 is a `precompile.ValueDecoder`.
func ConvertUint64(attributeValue string) (any, error) {
	return strconv.ParseUint(attributeValue, intBase, int64Bits)
}

// ReturnStringAsIs converts a given attribute of type string and returns the same string (as type
// any).
//
// ReturnStringAsIs is a `precompile.ValueDecoder`.
func ReturnStringAsIs(attributeValue string) (any, error) {
	return attributeValue, nil
}

// ConvertCommonHexAddress transfers a common hex address attribute to a common.Address and returns
// it as type any.
//
// ConvertCommonHexAddress is a `precompile.ValueDecoder`.
func ConvertCommonHexAddress(attributeValue string) (any, error) {
	return common.HexToAddress(attributeValue), nil
}

// ==============================================================================
// Helpers
// ==============================================================================

// validateAttributes validates an incoming Cosmos `event`. Specifically, it verifies that the
// number of attributes provided by the Cosmos `event` are adequate for it's corresponding
// Ethereum events.
func validateAttributes(pl *precompileLog, event *sdk.Event) error {
	if len(event.Attributes) < len(pl.indexedInputs)+len(pl.nonIndexedInputs) {
		return ErrNotEnoughAttributes
	}
	return nil
}

// searchAttributesForArg does a linear search through the given slice `attributes` for any
// attribute having a key that matches an Ethereum input `argName`. This function returns the index
// where `argName` was found or -1 if `argName` was not found.
// Complexity: O(N*M), N = len(`attributes`), M = average length of attribute key strings.
func searchAttributesForArg(attributes *[]abci.EventAttribute, argName string) int {
	for i, attribute := range *attributes {
		if abi.ToMixedCase(attribute.Key) == argName {
			return i
		}
	}
	// reached end of loop, `argName` not found
	return notFound
}
