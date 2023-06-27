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

import "errors"

var (
	// ErrIncompleteMethod is returned when a `FnAndGas` has missing, or nil, required fields.
	ErrIncompleteMethod = errors.New("incomplete precompile Method")

	// ErrAbiSigInvalid is returned when a user-provided ABI signature (`FnAndGas.AbiSig`) does
	// not match the Go-Ethereum style function signatures. Please check
	// core/vm/precompile/container/types.go for more information.
	ErrAbiSigInvalid = errors.New("user-provided ABI signature invalid: ")

	// ErrMethodNotFound is returned when the precompile method is not found.
	ErrMethodNotFound = errors.New("precompile method not found in contract ABI")

	// ErrContainerHasNoMethods is returned when a stateful container function is invoked but no
	// precompile methods were registered.
	ErrContainerHasNoMethods = errors.New("the stateful precompile has no methods to run")

	// ErrInvalidInputToPrecompile is returned when a precompile container receives invalid
	// input.
	ErrInvalidInputToPrecompile = errors.New("input bytes to precompile container are invalid")

	// ErrWrongContainerFactory is returned when the wrong precompile container factory is used
	// to build a precompile contract.
	ErrWrongContainerFactory = errors.New("this precompile contract implementation is not implemented")

	// ErrNoPrecompileMethodForABIMethod is returned when no precompile method is provided for a
	// corresponding ABI method.
	ErrNoPrecompileMethodForABIMethod = errors.New("this ABI method does not have a corresponding precompile method")

	ErrNoContext = errors.New("this precompile method implementation requires a PolarContext but none was provided")
)
