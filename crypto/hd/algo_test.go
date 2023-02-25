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
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	ethsecp256k1 "pkg.berachain.dev/stargazer/crypto/keys/ethsecp256k1"
	"pkg.berachain.dev/stargazer/eth/accounts"
	"pkg.berachain.dev/stargazer/eth/common"
	crypto "pkg.berachain.dev/stargazer/eth/crypto"
)

const (
	mnemonic = "picnic rent average infant boat squirrel federal assault mercy purity very" +
		"motor fossil wheel verify upset box fresh horse vivid copy predict square regret"
)

func TestHD(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "crypto/hd")
}

var _ = Describe("HD", func() {
	It("should derive the correct key", func() {
		EthSecp256k1 := EthSecp256k1

		bz, err := EthSecp256k1.Derive()(mnemonic, keyring.DefaultBIP39Passphrase,
			accounts.BIP44HDPath)
		Expect(err).NotTo(HaveOccurred())
		Expect(bz).NotTo(BeEmpty())

		badBz, err := EthSecp256k1.Derive()(mnemonic, keyring.DefaultBIP39Passphrase,
			"44'/60'/0'/0/0")
		Expect(err).NotTo(HaveOccurred())
		Expect(badBz).NotTo(BeEmpty())

		Expect(bz).NotTo(Equal(badBz))

		privkey := EthSecp256k1.Generate()(bz)
		badPrivKey := EthSecp256k1.Generate()(badBz)

		Expect(privkey.Equals(badPrivKey)).To(BeFalse())

		pk, err := privkey.(*ethsecp256k1.PrivKey).ToECDSA()
		Expect(err).NotTo(HaveOccurred())

		wallet, path, err := GenerateWallet(mnemonic)
		Expect(err).NotTo(HaveOccurred())
		*path = strings.Replace(*path, "H", "'", 3) // TODO: figure out why this is needed.
		Expect(*path).To(Equal(accounts.BIP44HDPath))
		Expect(crypto.FromECDSA(wallet)).To(Equal(privkey.Bytes()))

		// Equality of Addresses BIP44
		Expect(crypto.PubkeyToAddress(pk.PublicKey).String()).
			To(Equal("0xA588C66983a81e800Db4dF74564F09f91c026351"))
		Expect(crypto.PubkeyToAddress(wallet.PublicKey).String()).
			To(Equal("0xA588C66983a81e800Db4dF74564F09f91c026351"))
		Expect(common.BytesToAddress(privkey.PubKey().Address().Bytes()).String()).
			To(Equal("0xA588C66983a81e800Db4dF74564F09f91c026351"))
		Expect(common.BytesToAddress(privkey.PubKey().Address()).String()).
			To(Equal(crypto.PubkeyToAddress(wallet.PublicKey).String()))
	})
})

// func HDWallet(mnemonic string) (*secp256k1.PrivKey, *string, error) {
// 	hdPath := accounts.BIP44HDPath

// 	bz, err := hd.EthSecp256k1.Derive()(mnemonic, keyring.DefaultBIP39Passphrase, hdPath)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	privkey := hd.EthSecp256k1.Generate()(bz)
// 	path := accounts.StringPath(&bz)

// 	return privkey, path, nil
// }
