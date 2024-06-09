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
	"context"
	"math/big"
	"reflect"

	"github.com/berachain/polaris/eth/accounts/abi"
	"github.com/berachain/polaris/eth/core/vm"
	vmmock "github.com/berachain/polaris/eth/core/vm/mock"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Method", func() {
	Context("Calling the method", func() {
		It("should be able to call the Method's executable", func() {
			sc := &mockStatefulWithMethod{&mockBase{}, false}
			execute, found := reflect.TypeOf(sc).MethodByName("MockExecutable")
			Expect(found).To(BeTrue())
			method := newMethod(
				sc,
				abi.Method{},
				execute,
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
			res, err := method.Call(ctx, []byte{0, 0, 0, 0})
			Expect(err).ToNot(HaveOccurred())
			Expect(res).To(BeNil())
			Expect(sc.executableCalled).To(BeTrue())
		})
	})
})

var _ = Describe("Test MethoID", func() {
	It("should work", func() {
		x := make([]byte, 0)
		x = append(x, 0x12)
		x = append(x, 0x34)
		x = append(x, 0x56)
		x = append(x, 0x78)
		Expect(methodID(x)).To(Equal(methodID{0x12, 0x34, 0x56, 0x78}))
	})
})

// MOCKS BELOW.

type mockStatefulWithMethod struct {
	*mockBase
	executableCalled bool
}

func (ms *mockStatefulWithMethod) MockExecutable(
	_ context.Context,
) any {
	ms.executableCalled = true
	return nil
}
