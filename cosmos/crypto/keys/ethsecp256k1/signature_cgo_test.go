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

package ethsecp256k1

import (
	"crypto/ecdsa"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("PrivKey_PubKey", func() {
	var privKey *PrivKey
	var ecdsaPrivKey *ecdsa.PrivateKey

	BeforeEach(func() {
		var err error
		privKey, err = GenPrivKey()
		Expect(err).ToNot(HaveOccurred())
		ecdsaPrivKey, err = privKey.ToECDSA()
		Expect(err).ToNot(HaveOccurred())
	})

	It("validates signing bytes", func() {
		msg := []byte("hello world")
		// for the eth case, we have to manually hash in the test.
		sigHash := ethcrypto.Keccak256(msg)

		expectedSig, err := ethcrypto.Sign(sigHash, ecdsaPrivKey)
		Expect(err).ToNot(HaveOccurred())

		sig, err := privKey.Sign(msg)
		Expect(err).ToNot(HaveOccurred())
		Expect(expectedSig).To(Equal(sig))
	})

	It("validates signature", func() {
		msg := []byte("hello world")
		sigHash := ethcrypto.Keccak256(msg)
		sig, err := privKey.Sign(sigHash)
		Expect(err).ToNot(HaveOccurred())

		res := privKey.PubKey().VerifySignature(sigHash, sig)
		Expect(res).To(BeTrue())
	})
})
