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
