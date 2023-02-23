package crypto_test

import (
	"math/big"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/eth/crypto"
)

var _ = Describe("signer.PubKey", func() {
	var (
		key, _   = crypto.GenerateEthKey()
		signer   = types.NewLondonSigner(new(big.Int).SetInt64(420))
		signedTx *types.Transaction
	)
	It("should recover the public key from a valid signature", func() {
		txData := &types.LegacyTx{
			Nonce: 0,
			Gas:   10000000,
			Data:  []byte("abcdef"),
		}
		signedTx = types.MustSignNewTx(key, signer, txData)
		signer.PubKey(signedTx)
		Expect(signer.PubKey(signedTx)).To(Equal(crypto.FromECDSAPub(&key.PublicKey)))
	})

	It("should recover the public key from a valid signature", func() {
		txData := &types.DynamicFeeTx{
			Nonce: 0,
			Gas:   10000000,
			Data:  []byte("abcdef"),
		}
		signedTx = types.MustSignNewTx(key, signer, txData)
		signer.PubKey(signedTx)
		Expect(signer.PubKey(signedTx)).To(Equal(crypto.FromECDSAPub(&key.PublicKey)))
	})

	It("should recover the public key from a valid signature", func() {
		txData := &types.AccessListTx{
			Nonce: 0,
			Gas:   10000000,
			Data:  []byte("abcdef"),
		}
		signedTx = types.MustSignNewTx(key, signer, txData)
		signer.PubKey(signedTx)
		Expect(signer.PubKey(signedTx)).To(Equal(crypto.FromECDSAPub(&key.PublicKey)))
	})
})
