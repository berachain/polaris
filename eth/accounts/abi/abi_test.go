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

package abi_test

import (
	"testing"

	"github.com/berachain/polaris/eth/accounts/abi"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestABI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "eth/accounts/abi")
}

var _ = Describe("ABI Test Suite", func() {
	Describe("Test ToMixedCase", func() {
		It("should correctly convert under_score strings to mixedCase", func() {
			Expect(abi.ToMixedCase("creation_height")).To(Equal("creationHeight"))
			Expect(abi.ToMixedCase("creation_height_arg")).To(Equal("creationHeightArg"))
		})
	})

	Describe("Test GetIndexed", func() {
		var allArgs abi.Arguments
		BeforeEach(func() {
			allArgs = abi.Arguments{
				abi.Argument{},
				abi.Argument{
					Name:    "1",
					Indexed: true,
				},
				abi.Argument{
					Name:    "2",
					Indexed: true,
				},
				abi.Argument{},
				abi.Argument{},
				abi.Argument{},
				abi.Argument{
					Name:    "3",
					Indexed: true,
				},
			}
		})

		It("should correctly filter out indexed arguments", func() {
			indexedArgs := abi.Arguments{
				abi.Argument{
					Name:    "1",
					Indexed: true,
				},
				abi.Argument{
					Name:    "2",
					Indexed: true,
				},
				abi.Argument{
					Name:    "3",
					Indexed: true,
				},
			}
			args := abi.GetIndexed(allArgs)
			Expect(args).To(Equal(indexedArgs))
		})

		It("should panic if more than 3 indexed args are given", func() {
			Expect(func() { abi.GetIndexed(append(allArgs, abi.Argument{Indexed: true})) }).To(Panic())
		})
	})

	Describe("Test ToUnderScore", func() {
		It("should correctly convert mixedCase strings to under_score", func() {
			Expect(abi.ToUnderScore("Creation4Height")).To(Equal("creation4_height"))
			Expect(abi.ToUnderScore("creationHeight")).To(Equal("creation_height"))
			Expect(abi.ToUnderScore("creationHeightArg")).To(Equal("creation_height_arg"))
			Expect(abi.ToUnderScore("creation")).To(Equal("creation"))
			Expect(abi.ToUnderScore("creation_height")).To(Equal("creation_height"))
		})
	})
})
