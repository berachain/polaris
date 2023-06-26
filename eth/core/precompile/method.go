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
	"pkg.berachain.dev/polaris/eth/accounts/abi"
)

/**
 * 	Welcome to Stateful Precompiled Contracts! To build a stateful precompile, you must implement
 *  the `StatefulImpl` interface in `interfaces.go`; below are the suggested steps to
 *  follow:
 *	  1) Define a Solidity interface with the methods that you want implemented via a precompile.
 *	  2) Build a Go precompile contract, which implements the interface's methods.
 *       A) This precompile contract should expose the ABI's `Methods`, which can be generated via
 *          Go-Ethereum's abi package. These methods are of type `abi.Method`.
 *   	 B) This precompile contract should also expose the `Method`s. A `Method` includes the
 *          `executable`, which is the direct implementation of a corresponding ABI method, the
 *          `executable`'s `RequiredGas`, and the ABI signature. Do NOT provide the `AbiMethod` as
 *          this field will be automatically populated.
 **/

// Executable is a type of function that stateful precompiled contract will implement. Each
// Executable should directly correspond to an ABI method.
type Executable func(args ...any) (ret []any, err error)

// Method is a struct that contains the required information for the EVM to execute a stateful
// precompiled contract method.
type Method struct {
	// AbiMethod is the ABI `Methods` struct corresponding to this precompile's executable. NOTE:
	// this field should be left empty (as nil) as this will automatically be populated by the
	// corresponding interface's ABI.
	AbiMethod *abi.Method

	// AbiSig returns the method's string signature according to the ABI spec.
	// e.g.		function foo(uint32 a, int b) = "foo(uint32,int256)"
	// Note that there are no spaces and variable names in the signature.
	// Also note that "int" is substitute for its canonical representation "int256".
	AbiSig string

	// Execute is the precompile's executable which will execute the logic of the implemented
	// ABI method.
	Execute Executable

	// RequiredGas is the amount of gas (as a `uint64`) used up by the execution of `Execute`.
	// This field is optional; if left empty, the precompile's executable should consume gas using
	// the native gas meter.
	RequiredGas uint64
}

// ValidateBasic returns an error if this a precompile `Method` has invalid fields.
func (m *Method) ValidateBasic() error {
	// ensure all required fields are nonempty
	if len(m.AbiSig) == 0 || m.AbiMethod != nil || m.Execute == nil {
		return ErrIncompleteMethod
	}

	return nil
}

// Methods is a type that represents a list of precompile methods. This is what a stateful
// precompiled contract implementation should expose.
type Methods []*Method
