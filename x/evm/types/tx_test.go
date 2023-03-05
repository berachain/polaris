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

package types_test

import (
	fmt "fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"

	coretypes "pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/eth/crypto"
	"pkg.berachain.dev/stargazer/eth/params"
	"pkg.berachain.dev/stargazer/x/evm/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("EthTransactionRequest", func() {
	var (
		key, _  = crypto.GenerateEthKey()
		address = crypto.PubkeyToAddress(key.PublicKey)
	)
	When("it is a legacy tx", func() {
		var etr *types.EthTransactionRequest
		BeforeEach(func() {
			signer := coretypes.NewEIP2930Signer(params.DefaultChainConfig.ChainID)
			ltxData := &coretypes.LegacyTx{
				Nonce:    0,
				GasPrice: big.NewInt(2),
				Data:     []byte("abcdef"),
				To:       nil,
				Value:    new(big.Int),
			}
			tx, err := coretypes.SignNewTx(key, signer, ltxData)
			Expect(err).ToNot(HaveOccurred())
			etr = types.NewFromTransaction(tx)
		})

		It("should return the correct signer", func() {
			Expect(etr.GetSender()).To(Equal(address))
			Expect(etr.GetSigners()).To(Equal([]sdk.AccAddress{address.Bytes()}))
			fmt.Println(etr.AsTransaction().RawSignatureValues())
			_, err := etr.GetSignature()
			Expect(err).ToNot(HaveOccurred())
		})
	})

	When("it is a dynamic fee tx", func() {
		var etr *types.EthTransactionRequest
		BeforeEach(func() {
			signer := coretypes.LatestSignerForChainID(params.DefaultChainConfig.ChainID)
			dtxData := &coretypes.DynamicFeeTx{
				ChainID:   params.DefaultChainConfig.ChainID,
				Nonce:     0,
				Gas:       10,
				GasTipCap: new(big.Int),
				GasFeeCap: new(big.Int),
				To:        nil,
				Value:     new(big.Int),
				Data:      nil,
			}
			etr = types.NewFromTransaction(coretypes.MustSignNewTx(key, signer, dtxData))
		})

		It("should return the correct signer", func() {
			Expect(etr.GetSender()).To(Equal(address))
			Expect(etr.GetSigners()).To(Equal([]sdk.AccAddress{address.Bytes()}))
			_, err := etr.GetSignature()
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
