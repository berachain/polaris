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
	"pkg.berachain.dev/stargazer/eth/crypto"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("PrivKey_PubKey", func() {
	var privKey *PrivKey

	BeforeEach(func() {
		var err error
		privKey, err = GenPrivKey()
		Expect(err).ToNot(HaveOccurred())
	})

	It("validates signing bytes", func() {
		msg := []byte("hello world")
		sigHash := crypto.Keccak256(msg)
		expectedSig, err := crypto.EthSecp256k1Sign(sigHash, privKey.Bytes())
		Expect(err).ToNot(HaveOccurred())

		sig, err := privKey.Sign(sigHash)
		Expect(err).ToNot(HaveOccurred())
		Expect(expectedSig).To(Equal(sig))
	})

	It("validates signature", func() {
		msg := []byte("hello world")
		sigHash := crypto.Keccak256(msg)
		sig, err := privKey.Sign(sigHash)
		Expect(err).ToNot(HaveOccurred())

		res := privKey.PubKey().VerifySignature(msg, sig)
		Expect(res).To(BeTrue())
	})
})
