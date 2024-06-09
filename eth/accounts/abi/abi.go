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
