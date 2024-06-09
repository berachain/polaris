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
