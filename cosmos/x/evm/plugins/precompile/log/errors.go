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

import "errors"

var (
	// ErrNoAttributeKeyFound is returned when no Cosmos event attribute is provided for a
	// certain Ethereum event's argument.
	ErrNoAttributeKeyFound = errors.New(
		"this Ethereum event argument has no matching Cosmos attribute key")
	// ErrNotEnoughAttributes is returned when a Cosmos event does not have enough attributes for
	// its corresponding Ethereum event; there are less Cosmos event attributes than Ethereum event
	// arguments.
	ErrNotEnoughAttributes = errors.New(
		"not enough event attributes provided")
	// ErrNoValueDecoderFunc is returned when a Cosmos event's attribute key is not mapped to any
	// attribute value decoder function.
	ErrNoValueDecoderFunc = errors.New(
		"no value decoder function is found for event attribute key")
	// ErrNumberOfCoinsNotSupported is returned when the number of coins in a Cosmos event for the
	// "amount" attribute is not equal to 1.
	ErrNumberOfCoinsNotSupported = errors.New(
		"number of coins not supported")
)
