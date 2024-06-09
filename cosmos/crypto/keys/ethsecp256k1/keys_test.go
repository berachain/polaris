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
