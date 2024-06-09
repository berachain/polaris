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

package log

import (
	"testing"

	"github.com/berachain/polaris/eth/accounts/abi"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/x/evm/plugins/precompile/log")
}

var _ = Describe("precompileLog", func() {
	It("should properly create a new precompile log", func() {
		var pl *precompileLog
		Expect(func() {
			pl = newPrecompileLog(common.BytesToAddress([]byte{1}), mockDefaultAbiEvent())
		}).ToNot(Panic())
		Expect(pl.RegistryKey()).To(Equal("cancel_unbonding_delegation"))
		Expect(pl.id).To(Equal(crypto.Keccak256Hash(
			[]byte("CancelUnbondingDelegation(string,(uint256,string)[],int64)"),
		)))
		Expect(pl.precompileAddr).To(Equal(common.BytesToAddress([]byte{1})))
		Expect(pl.indexedInputs).To(HaveLen(1))
		Expect(pl.nonIndexedInputs).To(HaveLen(2))
	})
})

// MOCKS BELOW.

func mockDefaultAbiEvent() abi.Event {
	coinType, _ := abi.NewType("tuple[]", "structIStakingModule.Coin[]", []abi.ArgumentMarshaling{
		{
			Name:         "amount",
			Type:         "uint256",
			InternalType: "uint256",
			Components:   nil,
			Indexed:      false,
		},
		{
			Name:         "denom",
			Type:         "string",
			InternalType: "string",
			Components:   nil,
			Indexed:      false,
		},
	})
	int64Type, _ := abi.NewType("int64", "int64", nil)
	strType, _ := abi.NewType("string", "string", nil)
	return abi.NewEvent(
		"CancelUnbondingDelegation",
		"CancelUnbondingDelegation",
		false,
		abi.Arguments{
			{
				Name:    "option",
				Type:    strType,
				Indexed: true,
			},
			{
				Name:    "amount",
				Type:    coinType,
				Indexed: false,
			},
			{
				Name:    "creationHeight",
				Type:    int64Type,
				Indexed: false,
			},
		},
	)
}
