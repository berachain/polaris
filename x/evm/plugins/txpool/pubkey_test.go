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

package txpool

import (
	"math/big"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"pkg.berachain.dev/stargazer/crypto/keys/ethsecp256k1"
	"pkg.berachain.dev/stargazer/eth/core/types"
	ethcrypto "pkg.berachain.dev/stargazer/eth/crypto"
)

var _ = Describe("signer.PubKey", func() {
	var (
		key, _ = ethcrypto.GenerateEthKey()

		signer   = types.NewLondonSigner(new(big.Int).SetInt64(420))
		signedTx *types.Transaction
	)
	When("we are testing the code in geth", func() {
		It("should recover the public key from a valid signature", func() {
			txData := &types.LegacyTx{
				Nonce: 0,
				Gas:   10000000,
				Data:  []byte("abcdef"),
			}
			signedTx = types.MustSignNewTx(key, signer, txData)
			Expect(signer.PubKey(signedTx)).To(Equal(ethcrypto.FromECDSAPub(&key.PublicKey)))
		})

		It("should recover the public key from a valid signature", func() {
			txData := &types.DynamicFeeTx{
				Nonce: 0,
				Gas:   10000000,
				Data:  []byte("abcdef"),
			}
			signedTx = types.MustSignNewTx(key, signer, txData)
			Expect(signer.PubKey(signedTx)).To(Equal(ethcrypto.FromECDSAPub(&key.PublicKey)))
		})

		It("should recover the public key from a valid signature", func() {
			txData := &types.AccessListTx{
				Nonce: 0,
				Gas:   10000000,
				Data:  []byte("abcdef"),
			}
			signedTx = types.MustSignNewTx(key, signer, txData)
			Expect(signer.PubKey(signedTx)).To(Equal(ethcrypto.FromECDSAPub(&key.PublicKey)))
		})
	})

	When("we are testing the cosmos type key", func() {
		BeforeEach(func() {
			txData := &types.LegacyTx{
				Nonce: 0,
				Gas:   10000000,
				Data:  []byte("abcdef"),
			}
			signedTx = types.MustSignNewTx(key, signer, txData)
			Expect(signer.PubKey(signedTx)).To(Equal(ethcrypto.FromECDSAPub(&key.PublicKey)))
			bz, err := signer.PubKey(signedTx)
			Expect(err).ToNot(HaveOccurred())
			pk := &ethsecp256k1.PubKey{Key: bz}
			Expect(pk.Address()).To(Equal(ethcrypto.PubkeyToAddress(key.PublicKey)))
			pk2, err := PubkeyFromTx(signedTx, signer)
			Expect(err).ToNot(HaveOccurred())
			Expect(*pk2).To(Equal(pk))
		})
	})
})
