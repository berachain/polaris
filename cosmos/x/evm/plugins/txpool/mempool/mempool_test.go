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

package mempool

import (
	"bytes"
	"crypto/ecdsa"
	"math/big"
	"testing"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"

	"pkg.berachain.dev/polaris/cosmos/crypto/keys/ethsecp256k1"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/state"
	evmtypes "pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/crypto"
	"pkg.berachain.dev/polaris/eth/params"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestEthPool(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/x/evm/plugins/txpool/mempool")
}

var _ = Describe("EthTxPool", func() {
	var (
		ctx     sdk.Context
		sp      core.StatePlugin
		etp     *EthTxPool
		key1, _ = crypto.GenerateEthKey()
		addr1   = crypto.PubkeyToAddress(key1.PublicKey)
		key2, _ = crypto.GenerateEthKey()
		addr2   = crypto.PubkeyToAddress(key2.PublicKey)
	)

	BeforeEach(func() {
		sCtx, ak, bk, _ := testutil.SetupMinimalKeepers()
		sp = state.NewPlugin(ak, bk, testutil.EvmKey, &mockConfigurationPlugin{}, &mockPLF{})
		ctx = sCtx
		sp.Reset(ctx)
		sp.SetNonce(addr1, 1)
		sp.SetNonce(addr2, 2)
		sp.Finalize()
		sp.Reset(ctx)
		etp = NewEthereumTxPool()
		etp.SetNonceRetriever(sp)
	})

	Describe("All Cases", func() {
		It("should handle empty txs", func() {
			Expect(etp.Get(common.Hash{})).To(BeNil())
			emptyPending, emptyQueued := etp.Content()
			Expect(emptyPending).To(BeEmpty())
			Expect(emptyQueued).To(BeEmpty())
			Expect(etp.Nonce(addr1)).To(Equal(uint64(1)))
			Expect(etp.Nonce(addr2)).To(Equal(uint64(2)))
			Expect(etp.Nonce(common.HexToAddress("0x3"))).To(Equal(uint64(0)))
		})

		It("should return pending/queued txs with correct nonces", func() {
			ethTx1, tx1 := buildTx(key1, &coretypes.LegacyTx{Nonce: 1})
			ethTx2, tx2 := buildTx(key2, &coretypes.LegacyTx{Nonce: 2})

			Expect(etp.Insert(ctx, tx1)).ToNot(HaveOccurred())
			Expect(etp.Insert(ctx, tx2)).ToNot(HaveOccurred())

			Expect(etp.Nonce(addr1)).To(Equal(uint64(2)))
			Expect(etp.Nonce(addr2)).To(Equal(uint64(3)))

			Expect(etp.Get(ethTx1.Hash()).Hash()).To(Equal(ethTx1.Hash()))
			Expect(etp.Get(ethTx2.Hash()).Hash()).To(Equal(ethTx2.Hash()))

			pending, queued := etp.Content()
			lenP, lenQ := etp.Stats()

			Expect(pending).To(HaveLen(lenP))
			Expect(queued).To(HaveLen(lenQ))

			Expect(pending[addr1][0].Hash()).To(Equal(ethTx1.Hash()))
			Expect(pending[addr2][0].Hash()).To(Equal(ethTx2.Hash()))

			Expect(etp.Remove(tx2)).ToNot(HaveOccurred())
			Expect(etp.Get(ethTx2.Hash())).To(BeNil())
			p2, q2 := etp.ContentFrom(addr2)
			Expect(p2).To(BeEmpty())
			Expect(q2).To(BeEmpty())
			Expect(etp.Nonce(addr2)).To(Equal(uint64(2)))

			ethTx11, tx11 := buildTx(key1, &coretypes.LegacyTx{Nonce: 2})
			Expect(etp.Insert(ctx, tx11)).ToNot(HaveOccurred())
			Expect(etp.Nonce(addr1)).To(Equal(uint64(3)))
			p11, q11 := etp.ContentFrom(addr1)
			Expect(p11).To(HaveLen(2))
			Expect(p11[1].Hash()).To(Equal(ethTx11.Hash()))
			Expect(q11).To(HaveLen(0))
		})

		It("should handle replacement txs", func() {
			ethTx1, tx1 := buildTx(key1, &coretypes.LegacyTx{Nonce: 1, GasPrice: big.NewInt(1)})
			ethTx2, tx2 := buildTx(key1, &coretypes.LegacyTx{Nonce: 1, GasPrice: big.NewInt(2)})

			Expect(etp.Insert(ctx, tx1)).ToNot(HaveOccurred())
			Expect(etp.Insert(ctx, tx2)).ToNot(HaveOccurred())

			Expect(etp.Nonce(addr1)).To(Equal(uint64(2)))

			Expect(etp.Get(ethTx1.Hash())).To(BeNil())
			Expect(etp.Get(ethTx2.Hash()).Hash()).To(Equal(ethTx2.Hash()))

		})
	})
})

// MOCKS BELOW.

type mockConfigurationPlugin struct{}

func (mcp *mockConfigurationPlugin) GetEvmDenom() string {
	return "abera"
}

type mockPLF struct{}

func (mplf *mockPLF) Build(event *sdk.Event) (*coretypes.Log, error) {
	return &coretypes.Log{
		Address: common.BytesToAddress([]byte(event.Type)),
	}, nil
}

func buildTx(from *ecdsa.PrivateKey, txData *coretypes.LegacyTx) (*coretypes.Transaction, sdk.Tx) {
	signer := coretypes.LatestSignerForChainID(params.DefaultChainConfig.ChainID)
	signedEthTx := coretypes.MustSignNewTx(from, signer, txData)
	addr, _ := signer.Sender(signedEthTx)
	if !bytes.Equal(addr.Bytes(), crypto.PubkeyToAddress(from.PublicKey).Bytes()) {
		panic("sender mismatch")
	}
	pubKey := &ethsecp256k1.PubKey{Key: crypto.CompressPubkey(&from.PublicKey)}
	return signedEthTx, &mockSdkTx{
		signers: []sdk.AccAddress{cosmlib.AddressToAccAddress(addr)},
		msgs:    []sdk.Msg{evmtypes.NewFromTransaction(signedEthTx)},
		pubKeys: []cryptotypes.PubKey{pubKey},
		signatures: []signing.SignatureV2{
			{
				PubKey: pubKey,
				// NOTE: not including the signature data for the mock
				Sequence: txData.Nonce,
			},
		},
	}
}

type mockSdkTx struct {
	signers    []sdk.AccAddress
	msgs       []sdk.Msg
	pubKeys    []cryptotypes.PubKey
	signatures []signing.SignatureV2
}

func (m *mockSdkTx) ValidateBasic() error { return nil }

func (m *mockSdkTx) GetMsgs() []sdk.Msg { return m.msgs }

func (m *mockSdkTx) GetSigners() []sdk.AccAddress { return m.signers }

func (m *mockSdkTx) GetPubKeys() ([]cryptotypes.PubKey, error) { return m.pubKeys, nil }

func (m *mockSdkTx) GetSignaturesV2() ([]signing.SignatureV2, error) { return m.signatures, nil }
