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

package lib_test

import (
	cosmlib "github.com/berachain/polaris/cosmos/lib"

	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Address", func() {
	var addr common.Address
	var bech32 string

	BeforeEach(func() {
		addr = common.HexToAddress("0xCd8c4Cb0C7f93a2B74B3e522a1C7BE35bE1Fbc73")
		bech32 = "cosmos1ekxyevx8lyazka9nu532r3a7xklpl0rnjrc2a9"
	})

	It("should convert directly from eth to acc bech32 and vice versa", func() {
		accCodec := addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix())

		Expect(cosmlib.MustEthAddressFromString(accCodec, bech32)).To(Equal(addr))
		Expect(cosmlib.MustStringFromEthAddress(accCodec, addr)).To(Equal(bech32))

		acc, err := sdk.AccAddressFromBech32(bech32)
		Expect(err).NotTo(HaveOccurred())

		addr2 := cosmlib.MustEthAddressFromString(accCodec, acc.String())
		Expect(addr.String()).To(Equal(addr2.String()))

		bech32Str := sdk.MustBech32ifyAddressBytes(
			sdk.GetConfig().GetBech32AccountAddrPrefix(), addr.Bytes(),
		)
		Expect(bech32Str).To(Equal(bech32))
	})

	It("should return the correct val/cons address", func() {
		valCodec := addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32ValidatorAddrPrefix())

		valBech32 := cosmlib.MustStringFromEthAddress(valCodec, addr)
		valAddr, err := sdk.ValAddressFromBech32(valBech32)
		Expect(err).ToNot(HaveOccurred())
		valEthAddr := cosmlib.MustEthAddressFromString(valCodec, valAddr.String())
		Expect(addr.String()).To(Equal(valEthAddr.String()))

		bech32Str := sdk.MustBech32ifyAddressBytes(
			sdk.GetConfig().GetBech32ValidatorAddrPrefix(), valEthAddr.Bytes(),
		)
		Expect(bech32Str).To(Equal("cosmosvaloper1ekxyevx8lyazka9nu532r3a7xklpl0rnhhvl3k"))
	})
})
