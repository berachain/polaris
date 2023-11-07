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

package precompile

import (
	"reflect"
	"unicode"

	"github.com/berachain/polaris/eth/accounts/abi"
)

// this function finds the ABI method that matches the given impl method. It returns the key in the
// ABI methods map that matches the impl method.
func findMatchingABIMethod(
	implMethod reflect.Method, precompileABI map[string]abi.Method,
) (string, error) {
	implMethodName := formatName(implMethod.Name)

	for i := len(implMethodName) - 1; i >= 1; i-- {
		// try to match the substring of the impl method name to a ABI method name
		var matchedAbiMethod *abi.Method
		for name := range precompileABI {
			abiMethod := precompileABI[name]
			if implMethodName == abiMethod.RawName {
				if tryMatchInputs(implMethod, &abiMethod) {
					matchedAbiMethod = &abiMethod // we have a match
					break
				}
			}
		}

		// no match found, try again with a smaller substring
		if matchedAbiMethod == nil {
			implMethodName = implMethodName[:i]
			continue
		}

		// we found a matching impl method for the ABI method based on the inputs, now validate
		// that the outputs match
		if err := validateOutputs(implMethod, matchedAbiMethod); err != nil {
			return "", err
		}
		return matchedAbiMethod.Name, nil
	}

	return "", nil
}

// tryMatchInputs returns true iff the argument types match between the Go implementation and the
// ABI method.
func tryMatchInputs(implMethod reflect.Method, abiMethod *abi.Method) bool {
	abiMethodNumIn := len(abiMethod.Inputs)

	// First two args of Go precompile implementation are the receiver contract and the Context, so
	// verify that the ABI method has exactly 2 fewer inputs than the implementation method.
	if implMethod.Type.NumIn()-2 != abiMethodNumIn {
		return false
	}

	// If the function does not take any inputs, no need to check.
	if abiMethodNumIn > 0 {
		// Validate that the precompile input args types match ABI input arg types, excluding the
		// first two args (receiver contract and Context).
		for i := 2; i < implMethod.Type.NumIn(); i++ {
			implMethodParamType := implMethod.Type.In(i)
			abiMethodParamType := abiMethod.Inputs[i-2].Type.GetType()
			if validateArg(
				reflect.New(implMethodParamType).Elem(), reflect.New(abiMethodParamType).Elem(),
			) != nil {
				return false
			}
		}
	}

	return true
}

// formatName converts to first character of name to lowercase.
func formatName(name string) string {
	if len(name) == 0 {
		return name
	}

	ret := []rune(name)
	ret[0] = unicode.ToLower(ret[0])

	return string(ret)
}
