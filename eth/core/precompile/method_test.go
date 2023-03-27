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

package precompile_test

import (
	"context"
	"math/big"

	"pkg.berachain.dev/polaris/eth/accounts/abi"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/precompile"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Method", func() {
	Context("Basic - ValidateBasic Tests", func() {
		It("should error on missing Abi function signature", func() {
			methodMissingSig := &precompile.Method{
				Execute:     mockExecutable,
				RequiredGas: 10,
			}
			err := methodMissingSig.ValidateBasic()
			Expect(err).To(HaveOccurred())
		})

		It("should error on missing precompile executable", func() {
			methodMissingFunc := &precompile.Method{
				AbiSig:      "contractFunc(address)",
				RequiredGas: 10,
			}
			err := methodMissingFunc.ValidateBasic()
			Expect(err).To(HaveOccurred())
		})

		It("should error on given abi method", func() {
			methodMissingFunc := &precompile.Method{
				AbiSig:      "contractFunc(address)",
				RequiredGas: 10,
				Execute:     mockExecutable,
				AbiMethod:   &abi.Method{},
			}
			err := methodMissingFunc.ValidateBasic()
			Expect(err).To(HaveOccurred())
		})
	})
})

// MOCKS BELOW.

func mockExecutable(
	ctx context.Context,
	evm precompile.EVM,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	return nil, nil
}
