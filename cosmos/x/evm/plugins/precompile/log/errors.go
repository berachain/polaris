// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
