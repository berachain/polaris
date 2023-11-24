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

package ethsecp256k1

import (
	"testing"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestEthSecp256K1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/crypto/keys/ethsecp256k1")
}

var _ = Describe("PubPrivKey", func() {
	var privKey *PrivKey

	BeforeEach(func() {
		var err error
		privKey, err = GenPrivKey()
		Expect(err).ToNot(HaveOccurred())
	})

	It("validates type and equality", func() {
		Expect(privKey).To(BeAssignableToTypeOf((cryptotypes.PrivKey)(&PrivKey{})))
	})

	It("validates inequality", func() {
		privKey2, err := GenPrivKey()
		Expect(err).ToNot(HaveOccurred())
		Expect(privKey.Equals(privKey2)).To(BeFalse())
	})

	It("validates Ethereum address equality", func() {
		addr := privKey.PubKey().Address()
		key, err := privKey.ToECDSA()
		Expect(err).ToNot(HaveOccurred())
		expectedAddr := ethcrypto.PubkeyToAddress(key.PublicKey)
		Expect(expectedAddr.Bytes()).To(Equal(addr.Bytes()))
	})

	It("validates type", func() {
		pubKey := &PubKey{
			Key: privKey.PubKey().Bytes(),
		}
		Expect(pubKey).To(BeAssignableToTypeOf((cryptotypes.PubKey)(&PubKey{})))
	})

	It("validates equality", func() {
		privKey2, err := GenPrivKey()
		Expect(err).ToNot(HaveOccurred())
		Expect(privKey).ToNot(Equal(privKey2))
		Expect(privKey.PubKey()).ToNot(Equal(privKey2.PubKey()))
	})

})
