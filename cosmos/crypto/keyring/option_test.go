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

package keyring

import (
	"os"
	"strings"
	"testing"

	cryptocodec "github.com/berachain/polaris/cosmos/crypto/codec"
	"github.com/berachain/polaris/cosmos/crypto/hd"
	accounts "github.com/berachain/polaris/eth/accounts"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/std"

	"github.com/ethereum/go-ethereum/common"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var cdc *codec.ProtoCodec

func registerCodec() {
	interfaceRegistry := types.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
	cdc = codec.NewProtoCodec(interfaceRegistry)
}

func TestKeyring(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/crypto/keyring")
}

var _ = Describe("Keyring", func() {
	var (
		dir    string
		mockIn *strings.Reader
		kr     keyring.Keyring
	)

	BeforeEach(func() {
		var err error
		dir, err = os.MkdirTemp("", "keyring_test")
		Expect(err).NotTo(HaveOccurred())
		registerCodec()

		mockIn = strings.NewReader("")

		interfaceRegistry := types.NewInterfaceRegistry()
		std.RegisterInterfaces(interfaceRegistry)
		cryptocodec.RegisterInterfaces(interfaceRegistry)
		cdc = codec.NewProtoCodec(interfaceRegistry)

		kr, err = keyring.New(
			"accounts", keyring.BackendTest, dir, mockIn, cdc, OnlyEthSecp256k1Option(),
		)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		err := os.RemoveAll(dir)
		Expect(err).NotTo(HaveOccurred())
	})

	Context("Key operations", func() {
		It("should fail to retrieve key", func() {
			info, err := kr.Key("foo")
			Expect(err).To(HaveOccurred())
			Expect(info).To(BeNil())
		})
	})

	Context("NewMnemonic operation", func() {
		var (
			info     *keyring.Record
			mnemonic string
			err      error
		)

		BeforeEach(func() {
			registerCodec()
			mockIn.Reset("password\npassword\n")
			info, mnemonic, err = kr.NewMnemonic("foo", keyring.English, accounts.BIP44HDPath,
				keyring.DefaultBIP39Passphrase, hd.EthSecp256k1)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should create a new mnemonic and info", func() {
			Expect(mnemonic).NotTo(BeEmpty())
			Expect(info.Name).To(Equal("foo"))
			Expect(info.GetType().String()).To(Equal("local"))
			var pubKey cryptotypes.PubKey
			pubKey, err = info.GetPubKey()
			Expect(err).NotTo(HaveOccurred())
			Expect(pubKey.Type()).To(Equal(string(hd.EthSecp256k1Type)))
		})
	})

	Context("HD path operations", func() {
		var (
			mnemonic string
			bz       []byte
			err      error
		)

		BeforeEach(func() {
			registerCodec()
			mockIn.Reset("password\npassword\n")
			_, mnemonic, err = kr.NewMnemonic("foo", keyring.English, accounts.BIP44HDPath,
				keyring.DefaultBIP39Passphrase, hd.EthSecp256k1)
			Expect(err).NotTo(HaveOccurred())

			hdPath := accounts.BIP44HDPath

			bz, err = hd.EthSecp256k1.Derive()(mnemonic, keyring.DefaultBIP39Passphrase, hdPath)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should derive the correct HD path", func() {
			Expect(bz).NotTo(BeEmpty())
			var wrongBz []byte
			wrongBz, err = hd.EthSecp256k1.Derive()(mnemonic, keyring.DefaultBIP39Passphrase,
				"/wrong/hdPath")
			Expect(err).To(HaveOccurred())
			Expect(wrongBz).To(BeEmpty())
		})

		Context("Key generation and retrieval", func() {
			var (
				privkey cryptotypes.PrivKey
				addr    common.Address
			)

			BeforeEach(func() {
				registerCodec()
				mockIn.Reset("password\npassword\n")
				_, mnemonic, err = kr.NewMnemonic("foo", keyring.English, accounts.BIP44HDPath,
					keyring.DefaultBIP39Passphrase, hd.EthSecp256k1)
				Expect(err).NotTo(HaveOccurred())

				hdPath := accounts.BIP44HDPath

				bz, err = hd.EthSecp256k1.Derive()(mnemonic, keyring.DefaultBIP39Passphrase,
					hdPath)
				Expect(err).NotTo(HaveOccurred())

				privkey = hd.EthSecp256k1.Generate()(bz)
				addr = common.BytesToAddress(privkey.PubKey().Address().Bytes())
			})

			It("should generate and retrieve the correct private key and address", func() {
				Expect(addr.String()).To(Equal(common.Address(privkey.PubKey().
					Address()).String()))
			})
		})
	})
})
