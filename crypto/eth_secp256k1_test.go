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

package crypto

import (
	"github.com/berachain/stargazer/eth/crypto"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("PrivKey", func() {
	var privKey *EthSecp256K1PrivKey

	BeforeEach(func() {
		var err error
		privKey, err = GenerateKey()
		Expect(err).To(BeNil())
	})

	It("validates type and equality", func() {
		Expect(privKey).To(BeAssignableToTypeOf((cryptotypes.PrivKey)(&EthSecp256K1PrivKey{})))
	})

	It("validates inequality", func() {
		privKey2, err := GenerateKey()
		Expect(err).To(BeNil())
		Expect(privKey.Equals(privKey2)).To(BeFalse())
	})

	It("validates Ethereum address equality", func() {
		addr := privKey.PubKey().Address()
		key, err := privKey.ToECDSA()
		Expect(err).To(BeNil())
		expectedAddr := crypto.PubkeyToAddress(key.PublicKey)
		Expect(expectedAddr.Bytes()).To(Equal(addr.Bytes()))
	})

	It("validates signing bytes", func() {
		msg := []byte("hello world")
		sigHash := crypto.Keccak256Hash(msg)
		expectedSig, err := secp256k1.Sign(sigHash.Bytes(), privKey.Bytes())
		Expect(err).To(BeNil())

		sig, err := privKey.Sign(sigHash.Bytes())
		Expect(err).To(BeNil())
		Expect(expectedSig).To(Equal(sig))
	})
})

var _ = Describe("PrivKey_PubKey", func() {
	var privKey *EthSecp256K1PrivKey

	BeforeEach(func() {
		var err error
		privKey, err = GenerateKey()
		Expect(err).To(BeNil())
	})

	It("validates type", func() {
		pubKey := &EthSecp256K1PubKey{
			Key: privKey.PubKey().Bytes(),
		}
		Expect(pubKey).To(BeAssignableToTypeOf((cryptotypes.PubKey)(&EthSecp256K1PubKey{})))
	})

	It("validates equality", func() {
		privKey2, err := GenerateKey()
		Expect(err).To(BeNil())
		Expect(privKey).ToNot(Equal(privKey2))
		Expect(privKey.PubKey()).ToNot(Equal(privKey2.PubKey()))
	})

	It("validates signature", func() {
		msg := []byte("hello world")
		sigHash := crypto.Keccak256Hash(msg)
		sig, err := privKey.Sign(sigHash.Bytes())
		Expect(err).To(BeNil())

		res := privKey.PubKey().VerifySignature(msg, sig)
		Expect(res).To(BeTrue())
	})
})
