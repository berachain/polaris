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
