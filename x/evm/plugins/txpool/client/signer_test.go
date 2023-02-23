package client

import (
	"math/big"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"pkg.berachain.dev/stargazer/crypto"
	"pkg.berachain.dev/stargazer/eth/core/types"
	ethcrypto "pkg.berachain.dev/stargazer/eth/crypto"
)

var _ = Describe("signer.PubKey", func() {
	var (
		key, _ = ethcrypto.GenerateEthKey()

		signer   = types.NewLondonSigner(new(big.Int).SetInt64(420))
		signedTx *types.Transaction
	)
	When("we are testing the code in geth", func() {
		It("should recover the public key from a valid signature", func() {
			txData := &types.LegacyTx{
				Nonce: 0,
				Gas:   10000000,
				Data:  []byte("abcdef"),
			}
			signedTx = types.MustSignNewTx(key, signer, txData)
			Expect(signer.PubKey(signedTx)).To(Equal(ethcrypto.FromECDSAPub(&key.PublicKey)))
		})

		It("should recover the public key from a valid signature", func() {
			txData := &types.DynamicFeeTx{
				Nonce: 0,
				Gas:   10000000,
				Data:  []byte("abcdef"),
			}
			signedTx = types.MustSignNewTx(key, signer, txData)
			Expect(signer.PubKey(signedTx)).To(Equal(ethcrypto.FromECDSAPub(&key.PublicKey)))
		})

		It("should recover the public key from a valid signature", func() {
			txData := &types.AccessListTx{
				Nonce: 0,
				Gas:   10000000,
				Data:  []byte("abcdef"),
			}
			signedTx = types.MustSignNewTx(key, signer, txData)
			Expect(signer.PubKey(signedTx)).To(Equal(ethcrypto.FromECDSAPub(&key.PublicKey)))
		})
	})

	When("we are testing the cosmos type key", func() {
		BeforeEach(func() {
			txData := &types.LegacyTx{
				Nonce: 0,
				Gas:   10000000,
				Data:  []byte("abcdef"),
			}
			signedTx = types.MustSignNewTx(key, signer, txData)
			Expect(signer.PubKey(signedTx)).To(Equal(ethcrypto.FromECDSAPub(&key.PublicKey)))
			bz, err := signer.PubKey(signedTx)
			Expect(err).To(BeNil())
			pk, err := crypto.NewPubKeyFromBytes(bz)
			Expect(err).To(BeNil())
			Expect(pk.Address()).To(Equal(ethcrypto.PubkeyToAddress(key.PublicKey)))
			Expect(PubkeyFromTx(signedTx, signer)).To(Equal(pk))
		})
	})
})
