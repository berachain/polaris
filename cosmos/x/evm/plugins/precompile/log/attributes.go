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

	abci "github.com/cometbft/cometbft/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/eth/accounts/abi"
	"pkg.berachain.dev/polaris/eth/core/precompile"
)

const (
	// intBase is the base `int`s are parsed in, 10.
	intBase = 10
	// int64Bits is the number of bits stored in a variabe of `int64` type.
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
	sdk.AttributeKeyAmount:                  ConvertSdkCoin,
	stakingtypes.AttributeKeyValidator:      ConvertValAddressFromBech32,
	stakingtypes.AttributeKeySrcValidator:   ConvertValAddressFromBech32,
	stakingtypes.AttributeKeyDstValidator:   ConvertValAddressFromBech32,
	stakingtypes.AttributeKeyCreationHeight: ConvertInt64,
	stakingtypes.AttributeKeyDelegator:      ConvertAccAddressFromBech32,
	banktypes.AttributeKeySender:            ConvertAccAddressFromBech32,
	banktypes.AttributeKeyRecipient:         ConvertAccAddressFromBech32,
	banktypes.AttributeKeySpender:           ConvertAccAddressFromBech32,
	banktypes.AttributeKeyReceiver:          ConvertAccAddressFromBech32,
	banktypes.AttributeKeyMinter:            ConvertAccAddressFromBech32,
	banktypes.AttributeKeyBurner:            ConvertAccAddressFromBech32,
}

// ==============================================================================
// Default Attribute Value Decoder Functions
// ==============================================================================

// Compile-time assertions to ensure that the default attribute value decoder functions are
// valueDecoders.
var (
	_ precompile.ValueDecoder = ConvertSdkCoin
	_ precompile.ValueDecoder = ConvertValAddressFromBech32
	_ precompile.ValueDecoder = ConvertAccAddressFromBech32
	_ precompile.ValueDecoder = ConvertInt64
)

// ConvertSdkCoin converts the string representation of an `sdk.Coin` to a `*big.Int`.
//
// ConvertSdkCoin is a `precompile.ValueDecoder`.
func ConvertSdkCoin(attributeValue string) (any, error) {
	// extract the sdk.Coin from string value
	coin, err := sdk.ParseCoinNormalized(attributeValue)
	if err != nil {
		return nil, err
	}
	// convert the sdk.Coin to *big.Int
	return coin.Amount.BigInt(), nil
}

// ConvertValAddressFromBech32 converts a bech32 string representing a validator address to a
// common.Address.
//
// ConvertValAddressFromBech32 is a `precompile.ValueDecoder`.
func ConvertValAddressFromBech32(attributeValue string) (any, error) {
	// extract the sdk.ValAddress from string value
	valAddress, err := sdk.ValAddressFromBech32(attributeValue)
	if err != nil {
		return nil, err
	}
	// convert the sdk.ValAddress to common.Address
	return cosmlib.ValAddressToEthAddress(valAddress), nil
}

// ConvertAccAddressFromBech32 converts a bech32 string representing an account address to a
// common.Address.
//
// ConvertAccAddressFromBech32 is a `precompile.ValueDecoder`.
func ConvertAccAddressFromBech32(attributeValue string) (any, error) {
	// extract the sdk.AccAddress from string value
	accAddress, err := sdk.AccAddressFromBech32(attributeValue)
	if err != nil {
		return nil, err
	}
	// convert the sdk.AccAddress to common.Address
	return cosmlib.AccAddressToEthAddress(accAddress), nil
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
