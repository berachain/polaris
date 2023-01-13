// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package crypto

import (
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
		expectedAddr := PubkeyToAddress(key.PublicKey)
		Expect(expectedAddr.Bytes()).To(Equal(addr.Bytes()))
	})

	It("validates signing bytes", func() {
		msg := []byte("hello world")
		sigHash := Keccak256Hash(msg)
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
		sigHash := Keccak256Hash(msg)
		sig, err := privKey.Sign(sigHash.Bytes())
		Expect(err).To(BeNil())

		res := privKey.PubKey().VerifySignature(msg, sig)
		Expect(res).To(BeTrue())
	})
})
