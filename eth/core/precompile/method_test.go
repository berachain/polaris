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
	"reflect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"pkg.berachain.dev/polaris/eth/accounts/abi"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/precompile"
)

var _ = Describe("Method", func() {
	Context("Calling the method", func() {
		It("should be able to call the Method's executable", func() {
			method := &precompile.Method{
				AbiMethod:   &abi.Method{},
				AbiSig:      "mockExecutable()",
				Execute:     reflect.ValueOf(mockExecutable),
				RequiredGas: 10,
			}
			res := method.Execute.Call(
				[]reflect.Value{
					reflect.ValueOf(context.Background()),
					reflect.ValueOf(mockEVM{}),
					reflect.ValueOf(common.Address{}),
					reflect.ValueOf(big.NewInt(0)),
					reflect.ValueOf(false),
					reflect.ValueOf([]byte{}),
				})
			Expect(res[0].IsNil()).To(BeTrue())
		})
	})
})

// MOCKS BELOW.

func mockExecutable(
	_ context.Context,
	_ precompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	_ ...any,
) ([]any, error) {
	return nil, nil
}
