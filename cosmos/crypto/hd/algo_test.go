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

package hd

import (
	"strings"

	ethsecp256k1 "github.com/berachain/polaris/cosmos/crypto/keys/ethsecp256k1"
	"github.com/berachain/polaris/eth/accounts"
	"github.com/berachain/polaris/lib/utils"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	mnemonic = "absurd surge gather author blanket acquire proof struggle runway attract " +
		"cereal quiz tattoo shed almost sudden survey boring film memory picnic favorite " +
		"verb tank"
)

var _ = Describe("HD", func() {
	It("should derive the correct key", func() {
		EthSecp256k1 := EthSecp256k1

		// Derive the 0'th key from the mnemonic.
		bz, err := EthSecp256k1.Derive()(mnemonic, keyring.DefaultBIP39Passphrase,
			accounts.BIP44HDPath)
		Expect(err).NotTo(HaveOccurred())
		Expect(bz).NotTo(BeEmpty())

		badBz, err := EthSecp256k1.Derive()(mnemonic, keyring.DefaultBIP39Passphrase,
			"44'/118'/0'/0/0")
		Expect(err).NotTo(HaveOccurred())
		Expect(badBz).NotTo(BeEmpty())

		Expect(bz).NotTo(Equal(badBz))

		privkey := EthSecp256k1.Generate()(bz)
		badPrivKey := EthSecp256k1.Generate()(badBz)

		Expect(privkey.Equals(badPrivKey)).To(BeFalse())

		pk, err := utils.MustGetAs[*ethsecp256k1.PrivKey](privkey).ToECDSA()
		Expect(err).NotTo(HaveOccurred())

		wallet, path, err := GenerateWallet(mnemonic)
		Expect(err).NotTo(HaveOccurred())
		*path = strings.Replace(*path, "H", "'", 3) // TODO: figure out why this is needed.
		Expect(*path).To(Equal(accounts.BIP44HDPath))
		Expect(crypto.FromECDSA(wallet)).To(Equal(privkey.Bytes()))

		// Check to verify that the address is correct.
		// Also verified manually with metamask: https://imgur.com/a/Bz2jLaP
		Expect(crypto.PubkeyToAddress(pk.PublicKey).String()).
			To(Equal("0x20f33CE90A13a4b5E7697E3544c3083B8F8A51D4"))
		Expect(crypto.PubkeyToAddress(wallet.PublicKey).String()).
			To(Equal("0x20f33CE90A13a4b5E7697E3544c3083B8F8A51D4"))
		Expect(common.BytesToAddress(privkey.PubKey().Address().Bytes()).String()).
			To(Equal("0x20f33CE90A13a4b5E7697E3544c3083B8F8A51D4"))
		Expect(common.BytesToAddress(privkey.PubKey().Address()).String()).
			To(Equal(crypto.PubkeyToAddress(wallet.PublicKey).String()))
	})
})

var _ = Describe("Prove EDSCAify isn't needed", func() {
	It("should round trip", func() {
		// Generate a random private key.
		key, err := ethsecp256k1.GenPrivKey()
		Expect(err).NotTo(HaveOccurred())

		// Convert the private key to an ECDSA private key.
		x, err := ethsecp256k1.PrivKey{Key: key.Key}.ToECDSA()
		Expect(err).NotTo(HaveOccurred())
		Expect(key.Key).To(Equal(crypto.FromECDSA(x)))
	})
})
