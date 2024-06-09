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

import "errors"

var (
	// ErrMethodNotFound is returned when the precompile method is not found.
	ErrMethodNotFound = errors.New(
		"precompile method not found in contract ABI")

	// ErrContainerHasNoMethods is returned when a stateful container function is invoked but no
	// precompile methods were registered.
	ErrContainerHasNoMethods = errors.New(
		"the stateful precompile has no methods to run")

	// ErrInvalidInputToPrecompile is returned when a precompile container receives invalid
	// input.
	ErrInvalidInputToPrecompile = errors.New(
		"input bytes to precompile container are invalid")

	// ErrWrongContainerFactory is returned when the wrong precompile container factory is used
	// to build a precompile contract.
	ErrWrongContainerFactory = errors.New(
		"wrong container factory for this precompile implementation")

	// ErrNoPrecompileMethodForABIMethod is returned when no precompile method is provided for a
	// corresponding ABI method.
	ErrNoPrecompileMethodForABIMethod = errors.New(
		"this ABI method does not have a corresponding precompile method")
)
