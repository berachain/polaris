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

	"pkg.berachain.dev/polaris/eth/accounts/abi"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/core/vm"
	vmmock "pkg.berachain.dev/polaris/eth/core/vm/mock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Method", func() {
	Context("Calling the method", func() {
		It("should be able to call the Method's executable", func() {
			method := precompile.NewMethod(
				&abi.Method{},
				"mockExecutable()",
				reflect.ValueOf(mockExecutable),
			)
			ctx := vm.NewPolarContext(
				context.Background(),
				vmmock.NewEVM(),
				common.Address{1},
				big.NewInt(0),
			)

			// due to how the go "reflect" package works, we need to pass in the `stateful` in the
			// method call as the first parameter to thef function. this is taken care of for the
			// caller of the precompile under the hood, and users dont have to worry when
			// implementing their own precompiles.
			sc := &mockStateful{&mockBase{}}
			res, err := method.Call(sc, ctx, []byte{0, 0, 0, 0})
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(BeNil())
		})
	})
})

// MOCKS BELOW.

//nolint:revive // needed for go "reflect" package.
func mockExecutable(
	_ precompile.Registrable,
	_ context.Context,
) any {
	return nil
}
