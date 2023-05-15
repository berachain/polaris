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

package abi

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// maxIndexedArgs is the maximum number of indexed arguments allowed in an Ethereum event log.
const maxIndexedArgs = 3

type (
	ABI                = abi.ABI
	Argument           = abi.Argument
	ArgumentMarshaling = abi.ArgumentMarshaling
	Arguments          = abi.Arguments
	Event              = abi.Event
	Method             = abi.Method
)

var (
	MakeTopics = abi.MakeTopics
	NewEvent   = abi.NewEvent
	NewType    = abi.NewType
)

// ToMixedCase converts a under_score formatted string to mixedCase format (camelCase with the
// first letter lowercase). This function is inspired by the geth `abi.ToCamelCase“ function.
func ToMixedCase(input string) string {
	parts := strings.Split(input, "_")
	for i, s := range parts {
		if i > 0 && len(s) > 0 {
			parts[i] = strings.ToUpper(s[:1]) + s[1:]
		}
	}
	return strings.Join(parts, "")
}

// ToUnderScore converts a mixedCase formatted string to under_score format. This function is
// inspired by the geth `abi.ToCamelCase` function, but has the opposite behavior.
func ToUnderScore(input string) string {
	var output string
	for i, s := range input {
		if i > 0 && s >= 'A' && s <= 'Z' {
			output += "_"
		}
		output += string(s)
	}
	return strings.ToLower(output)
}

// GetIndexed extracts indexed arguments from a set of arguments. Will panic if more than 3
// indexed arguments are provided by the inputs ABI.
func GetIndexed(args abi.Arguments) abi.Arguments {
	var indexed abi.Arguments
	for _, arg := range args {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}

	if len(indexed) > maxIndexedArgs {
		panic(ErrTooManyIndexedArgs)
	}

	return indexed
}
