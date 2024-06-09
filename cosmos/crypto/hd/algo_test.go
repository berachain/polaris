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
