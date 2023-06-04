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
	"sync"
	"testing"
	"time"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"

	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/trie"

	"pkg.berachain.dev/polaris/cosmos/crypto/keys/ethsecp256k1"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	evmtypes "pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/core/mock"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/core/vm"
	vmmock "pkg.berachain.dev/polaris/eth/core/vm/mock"
	"pkg.berachain.dev/polaris/eth/crypto"
	"pkg.berachain.dev/polaris/eth/params"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestEthPool(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/x/evm/plugins/txpool/mempool")
}

var _ = Describe("WrappedGethTxPool", func() {
	var (
		ctx        sdk.Context
		sdb        *vmmock.PolarisStateDBMock
		etp        *WrappedGethTxPool
		key1, _    = crypto.GenerateEthKey()
		addr1      = crypto.PubkeyToAddress(key1.PublicKey)
		key2, _    = crypto.GenerateEthKey()
		addr2      = crypto.PubkeyToAddress(key2.PublicKey)
		addr1Nonce uint64
		addr2Nonce uint64
	)

	BeforeEach(func() {
		etp = NewWrappedGethTxPool()

		cp := mock.NewConfigurationPluginMock()

		sdb = vmmock.NewEmptyStateDB()
		sdb.GetNonceFunc = func(addr common.Address) uint64 {
			if addr == addr1 {
				return addr1Nonce
			}
			if addr == addr2 {
				return addr2Nonce
			}
			return 0
		}
		sdb.GetBalanceFunc = func(addr common.Address) *big.Int {
			return big.NewInt(1000000000000000000)
		}
		bc := newMockBlockChain(sdb)
		txp := txpool.NewTxPool(txpool.DefaultConfig, cp.ChainConfig(), bc)
		etp.SetTxPool(txp)
		etp.SetGasPrice(big.NewInt(1)) // TODO: set this for real int he real app.
		addr1Nonce = 1
		addr2Nonce = 2
		etp.Setup(cp, &mockSerializer{})
		etp.Prepare(bc.CurrentBlock())
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

		It("should error with low nonces", func() {
			_, tx1 := buildTx(key1, &coretypes.LegacyTx{Nonce: 0, GasPrice: big.NewInt(100), Gas: 100000})
			err := etp.Insert(ctx, tx1)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("nonce too low"))
		})

		It("should return pending/queued txs with correct nonces", func() {
			ethTx1, tx1 := buildTx(key1, &coretypes.LegacyTx{Nonce: 1, GasPrice: big.NewInt(100), Gas: 100000})
			ethTx2, tx2 := buildTx(key2, &coretypes.LegacyTx{Nonce: 2, GasPrice: big.NewInt(100), Gas: 100000})

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

			Expect(isPendingTx(etp, ethTx1)).To(BeTrue())
			Expect(isPendingTx(etp, ethTx2)).To(BeTrue())

			Expect(etp.Remove(tx2)).ToNot(HaveOccurred())
			Expect(etp.Get(ethTx2.Hash())).To(BeNil())
			p2, q2 := etp.ContentFrom(addr2)
			Expect(p2).To(BeEmpty())
			Expect(q2).To(BeEmpty())
			Expect(etp.Nonce(addr2)).To(Equal(uint64(2)))

			ethTx11, tx11 := buildTx(key1, &coretypes.LegacyTx{Nonce: 2, GasPrice: big.NewInt(100), Gas: 100000})
			Expect(etp.Insert(ctx, tx11)).ToNot(HaveOccurred())
			Expect(etp.Nonce(addr1)).To(Equal(uint64(3)))
			p11, q11 := etp.ContentFrom(addr1)
			Expect(p11).To(HaveLen(2))

			Expect(isPendingTx(etp, ethTx11)).To(BeTrue())
			Expect(q11).To(BeEmpty())
		})

		It("should handle replacement txs", func() {
			ethTx1, tx1 := buildTx(key1, &coretypes.LegacyTx{Nonce: 1, GasPrice: big.NewInt(90), Gas: 10000000})
			ethTx2, tx2 := buildTx(key1, &coretypes.LegacyTx{Nonce: 1, GasPrice: big.NewInt(100), Gas: 10000000})

			Expect(etp.Insert(ctx, tx1)).ToNot(HaveOccurred())
			Expect(etp.Insert(ctx, tx2)).ToNot(HaveOccurred())

			Expect(etp.Nonce(addr1)).To(Equal(uint64(2)))

			Expect(etp.Get(ethTx1.Hash())).To(BeNil())
			Expect(etp.Get(ethTx2.Hash()).Hash()).To(Equal(ethTx2.Hash()))

		})
		It("should enqueue transactions with out of order nonces then poll from queue when inorder nonce tx is received",
			func() {
				_, tx1 := buildTx(key1, &coretypes.LegacyTx{Nonce: 1, GasPrice: big.NewInt(100), Gas: 100000})
				ethtx3, tx3 := buildTx(key1, &coretypes.LegacyTx{Nonce: 3, GasPrice: big.NewInt(100), Gas: 100000})

				Expect(etp.Insert(ctx, tx1)).ToNot(HaveOccurred())
				Expect(etp.Insert(ctx, tx3)).ToNot(HaveOccurred())

				Expect(isQueuedTx(etp, ethtx3)).To(BeTrue())

				_, tx2 := buildTx(key1, &coretypes.LegacyTx{Nonce: 2, GasPrice: big.NewInt(1000), Gas: 100000})
				Expect(etp.Insert(ctx, tx2)).ToNot(HaveOccurred())

				_, queuedTransactions := etp.ContentFrom(addr1)
				Expect(queuedTransactions).To(BeEmpty())
				Expect(etp.Nonce(addr1)).To(Equal(uint64(4)))
			})

		It("should not allow replacement txs with gas increase < 10%", func() {
			_, tx1 := buildTx(key1, &coretypes.LegacyTx{Nonce: 1, GasPrice: big.NewInt(99), Gas: 10000000})
			_, tx2 := buildTx(key1, &coretypes.LegacyTx{Nonce: 1, GasPrice: big.NewInt(100), Gas: 10000000})
			_, tx3 := buildTx(key1, &coretypes.LegacyTx{Nonce: 1, GasPrice: big.NewInt(99), Gas: 10000000})

			Expect(etp.Insert(ctx, tx1)).ToNot(HaveOccurred())
			Expect(etp.Insert(ctx, tx2)).To(HaveOccurred())
			Expect(etp.Insert(ctx, tx3)).To(HaveOccurred()) // should skip the math for replacement
		})

		It("should handle spam txs and prevent DOS attacks", func() {
			for i := 1; i < 1000; i++ {
				_, tx := buildTx(key1, &coretypes.LegacyTx{Nonce: uint64(i), GasPrice: big.NewInt(100), Gas: 100000})
				Expect(etp.Insert(ctx, tx)).ToNot(HaveOccurred())
			}
			// probably more stuff down here...
		})

		It("should be able to fetch transactions from the cache", func() {

			var txHashes []common.Hash
			for i := 1; i < 100; i++ {
				ethTx, tx := buildTx(key1, &coretypes.LegacyTx{Nonce: uint64(i), GasPrice: big.NewInt(100), Gas: 100000})
				Expect(etp.Insert(ctx, tx)).ToNot(HaveOccurred())
				txHashes = append(txHashes, ethTx.Hash())
			}
			for _, txHash := range txHashes {
				Expect(etp.Get(txHash).Hash()).To(Equal(txHash))
			}

		})

		It("should allow resubmitting a transaction with same nonce but different fields", func() {
			_, tx := buildTx(key1, &coretypes.LegacyTx{Nonce: 1, GasPrice: big.NewInt(100), Gas: 100000})
			_, tx2 := buildTx(key1, &coretypes.LegacyTx{Nonce: 1,
				GasPrice: big.NewInt(500), Data: []byte("blahblah"), Gas: 100000})

			Expect(etp.Insert(ctx, tx)).ToNot(HaveOccurred())
			Expect(etp.Insert(ctx, tx2)).ToNot(HaveOccurred())
		})

		It("should prioritize transactions first by nonce, then priority", func() {
			_, tx := buildTx(key1, &coretypes.LegacyTx{Nonce: 1, GasPrice: big.NewInt(10), Gas: 100000})
			_, tx2 := buildTx(key1, &coretypes.LegacyTx{Nonce: 2, GasPrice: big.NewInt(50), Gas: 100000})
			_, tx3 := buildTx(key1, &coretypes.LegacyTx{Nonce: 3, GasPrice: big.NewInt(30), Gas: 100000})
			_, tx31 := buildTx(key1, &coretypes.LegacyTx{Nonce: 3, GasPrice: big.NewInt(50), Gas: 100000})

			Expect(etp.Insert(ctx, tx)).ToNot(HaveOccurred())
			Expect(etp.Insert(ctx, tx2)).ToNot(HaveOccurred())
			Expect(etp.Insert(ctx, tx3)).ToNot(HaveOccurred())
			Expect(etp.Insert(ctx, tx31)).ToNot(HaveOccurred())

			var prevTx *coretypes.Transaction
			for _, txs := range etp.Pending(false) {
				for _, tx := range txs {
					if prevTx == nil { // for the first transaction from a sender
						prevTx = tx
					} else { // for the rest of the transactions
						Expect(tx.Nonce()).To(Equal(prevTx.Nonce() + 1))
						prevTx = tx

					}
					// NOTE: replacement transactions are not handled because the old tx is removed from the pool
				}
			}
			pending, _ := etp.Stats()
			Expect(pending).To(Equal(3))
		})
		It("should handle many pending txs", func() {
			ethTx1, tx1 := buildTx(key1, &coretypes.LegacyTx{Nonce: 1, GasPrice: big.NewInt(100), Gas: 100000})
			ethTx2, tx2 := buildTx(key1, &coretypes.LegacyTx{Nonce: 2, GasPrice: big.NewInt(200), Gas: 100000})
			ethTx3, tx3 := buildTx(key1, &coretypes.LegacyTx{Nonce: 3, GasPrice: big.NewInt(300), Gas: 100000})
			Expect(etp.Insert(ctx, tx1)).ToNot(HaveOccurred())
			Expect(etp.Insert(ctx, tx2)).ToNot(HaveOccurred())
			Expect(etp.Insert(ctx, tx3)).ToNot(HaveOccurred())

			expected := []common.Hash{ethTx1.Hash(), ethTx2.Hash(), ethTx3.Hash()}
			found := etp.Pending(false)[addr1]
			Expect(found).To(HaveLen(3))
			for i, ethTx := range found {
				Expect(ethTx.Hash()).To(Equal(expected[i]))
			}
		})

		It("should not return pending when queued", func() {
			_, tx2 := buildTx(key1, &coretypes.LegacyTx{Nonce: 2, GasPrice: big.NewInt(200), Gas: 100000})
			_, tx3 := buildTx(key1, &coretypes.LegacyTx{Nonce: 3, GasPrice: big.NewInt(300), Gas: 100000})
			Expect(etp.Insert(ctx, tx2)).ToNot(HaveOccurred())
			Expect(etp.Insert(ctx, tx3)).ToNot(HaveOccurred())

			// TODO: Check Content
			Expect(etp.Pending(false)[addr1]).To(BeEmpty())
			// Expect(etp.Content()[addr1]).To(HaveLen(2))
			pending, queued := etp.Stats()
			Expect(pending).To(Equal(0))
			Expect(queued).To(Equal(2))
		})

		// TODO THESE ARE HOOD AS FUCK TESTS
		It("should handle concurrent additions", func() {

			// apologies in advance for this test, it's not great.

			// var wg sync.WaitGroup

			// wg.Add(1)
			// go func(etp *WrappedGethTxPool) {
			// 	defer wg.Done()
			// 	for i := 1; i <= 10; i++ {
			// 		_, tx := buildTx(key1, &coretypes.LegacyTx{Nonce: uint64(i)})
			// 		Expect(etp.Insert(ctx, tx)).ToNot(HaveOccurred())
			// 	}
			// }(etp)

			// wg.Add(1)
			// go func(etp *WrappedGethTxPool) {
			// 	defer wg.Done()
			// 	for i := 2; i <= 11; i++ {
			// 		_, tx := buildTx(key2, &coretypes.LegacyTx{Nonce: uint64(i)})
			// 		Expect(etp.Insert(ctx, tx)).ToNot(HaveOccurred())
			// 	}
			// }(etp)

			// wg.Wait()
			// lenPending, _ := etp.Stats()
			// Expect(lenPending).To(BeEquivalentTo(20))
		})
		It("should handle concurrent reads", func() {

			readsFromA := 0
			readsFromB := 0

			// fill mempoopl with transactions
			var wg sync.WaitGroup

			for i := 1; i < 10; i++ {
				_, tx := buildTx(key1, &coretypes.LegacyTx{Nonce: uint64(i), GasPrice: big.NewInt(100), Gas: 100000})
				Expect(etp.Insert(ctx, tx)).ToNot(HaveOccurred())
			}

			// concurrently read mempool from Peer A ...
			wg.Add(1)
			go func(etp *WrappedGethTxPool) {
				defer wg.Done()
				for _, txs := range etp.Pending(false) {
					for range txs {
						readsFromA++
					}
				}
			}(etp)

			// ... and peer B
			wg.Add(1)
			go func(etp *WrappedGethTxPool) {
				defer wg.Done()
				for _, txs := range etp.Pending(false) {
					for range txs {
						readsFromB++
					}
				}
			}(etp)

			wg.Wait()
			Expect(readsFromA).To(BeEquivalentTo(readsFromB))
		})

		// It("should be able to return the transaction priority for a Cosmos tx and effective gas tip value", func() {
		// 	ethTx1, tx1 := buildTx(key1, &coretypes.DynamicFeeTx{
		// 		Nonce: 1, GasTipCap: big.NewInt(1), GasFeeCap: big.NewInt(10000)})
		// 	ethTx2, tx2 := buildTx(key2, &coretypes.DynamicFeeTx{
		// 		Nonce: 2, GasTipCap: big.NewInt(2), GasFeeCap: big.NewInt(200)})

		// 	// Test that the priority policy is working as expected.
		// 	tpp := EthereumTxPriorityPolicy{baseFee: big.NewInt(69)}
		// 	Expect(tpp.GetTxPriority(ctx, tx1)).To(Equal(ethTx1.EffectiveGasTipValue(tpp.baseFee)))
		// 	Expect(tpp.GetTxPriority(ctx, tx2)).To(Equal(ethTx2.EffectiveGasTipValue(tpp.baseFee)))

		// 	// Test live mempool
		// 	err := etp.Insert(ctx, tx1)
		// 	Expect(err).ToNot(HaveOccurred())
		// 	err = etp.Insert(ctx, tx2)
		// 	Expect(err).ToNot(HaveOccurred())

		// 	// Test that the priority policy is working as expected.
		// 	iter := etp.Select(context.TODO(), nil)
		// 	higherPriorityTx := evmtypes.GetAsEthTx(iter.Tx())
		// 	lowerPriorityTx := evmtypes.GetAsEthTx(iter.Next().Tx())
		// 	Expect(higherPriorityTx.Hash()).To(Equal(ethTx2.Hash()))
		// 	Expect(lowerPriorityTx.Hash()).To(Equal(ethTx1.Hash()))
		// })

		It("should throw when attempting to remove a transaction that doesn't exist", func() {
			_, tx := buildTx(key1, &coretypes.LegacyTx{Nonce: 1, GasPrice: big.NewInt(100), Gas: 100000})
			Expect(etp.Insert(ctx, tx)).ToNot(HaveOccurred())
			// TODO: @calbera we might want to call with sync here? idk if this could cause problems.
			// TODO: since we are calling Insert with sync=false rn, this creates a race condition with the line
			// TODO: below, we temporarily add a sleep to avoid this, but we should fix this.
			time.Sleep(200 * time.Millisecond)
			Expect(etp.Remove(tx)).ToNot(HaveOccurred())
			Expect(etp.Remove(tx)).To(HaveOccurred())
		})

		It("should return StateDB's nonce when seeing nonce gap on first lookup", func() {
			ethTx, tx := buildTx(key1, &coretypes.LegacyTx{Nonce: 3, GasPrice: big.NewInt(100), Gas: 100000})

			Expect(etp.Insert(ctx, tx)).ToNot(HaveOccurred())

			sdbNonce := sdb.GetNonce(addr1)
			txNonce := ethTx.Nonce()
			Expect(sdbNonce).ToNot(BeEquivalentTo(txNonce))
			Expect(sdbNonce).To(BeEquivalentTo(1))
			Expect(txNonce).To(BeEquivalentTo(3))
			Expect(etp.Nonce(addr1)).To(BeEquivalentTo(sdbNonce))

		})
		// TODO: SDK TRANSACTIONS BIG BROKEN.
		// It("should break out of func Nonce(addr) when seeing a noncontigious nonce gap", func() {
		// 	_, tx1 := buildTx(key1, &coretypes.LegacyTx{Nonce: 1, GasPrice: big.NewInt(100), Gas: 100000})
		// 	tx2 := buildSdkTx(key1, 2)
		// 	_, tx3 := buildTx(key1, &coretypes.LegacyTx{Nonce: 3, GasPrice: big.NewInt(100), Gas: 100000})
		// 	_, tx10 := buildTx(key1, &coretypes.LegacyTx{Nonce: 10, GasPrice: big.NewInt(100), Gas: 100000})

		// 	Expect(etp.Insert(ctx, tx1)).ToNot(HaveOccurred())
		// 	Expect(etp.Nonce(addr1)).To(BeEquivalentTo(2))

		// 	Expect(etp.Insert(ctx, tx2)).ToNot(HaveOccurred())
		// 	Expect(etp.Nonce(addr1)).To(BeEquivalentTo(3))

		// 	Expect(etp.Insert(ctx, tx3)).ToNot(HaveOccurred())
		// 	Expect(etp.Nonce(addr1)).To(BeEquivalentTo(4))

		// 	Expect(etp.Insert(ctx, tx10)).ToNot(HaveOccurred())
		// 	Expect(etp.Nonce(addr1)).To(BeEquivalentTo(4)) // should not be 10
		// })

	})
})

// MOCKS BELOW.

type mockBlockChain struct {
	sdb           vm.GethStateDB
	chainHeadFeed *event.Feed
}

func newMockBlockChain(sdb vm.GethStateDB) *mockBlockChain {
	return &mockBlockChain{
		sdb:           sdb,
		chainHeadFeed: new(event.Feed),
	}
}

func (bc *mockBlockChain) CurrentBlock() *coretypes.Header {
	return &coretypes.Header{
		Number:   new(big.Int),
		BaseFee:  big.NewInt(1),
		GasLimit: 1000000000,
	}
}

func (bc *mockBlockChain) GetBlock(_ common.Hash, _ uint64) *coretypes.Block {
	return coretypes.NewBlock(bc.CurrentBlock(), nil, nil, nil, trie.NewStackTrie(nil))
}

func (bc *mockBlockChain) StateAt(common.Hash) (vm.GethStateDB, error) {
	return bc.sdb, nil
}

func (bc *mockBlockChain) StateAtHeader(*coretypes.Header) (vm.GethStateDB, error) {
	return bc.sdb, nil
}

func (bc *mockBlockChain) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	return bc.chainHeadFeed.Subscribe(ch)
}

func isQueuedTx(mempool *WrappedGethTxPool, tx *coretypes.Transaction) bool {
	_, queued := mempool.Content()

	for _, list := range queued {
		for _, ethTx := range list {
			if tx.Hash() == ethTx.Hash() {
				return true
			}
		}
	}
	return false
}

func isPendingTx(mempool *WrappedGethTxPool, tx *coretypes.Transaction) bool {
	pending, _ := mempool.Content()

	for _, list := range pending {
		for _, ethTx := range list {
			if tx.Hash() == ethTx.Hash() {
				return true
			}
		}
	}
	return false
}

type mockSerializer struct{}

func (ms *mockSerializer) SerializeToSdkTx(signedTx *coretypes.Transaction) (sdk.Tx, error) {
	signer := coretypes.LatestSignerForChainID(params.DefaultChainConfig.ChainID)
	addr, err := signer.Sender(signedTx)
	if err != nil {
		return nil, err
	}
	pk, err := coretypes.PubkeyFromTx(signedTx, signer)
	if err != nil {
		return nil, err
	}
	pubKey := &ethsecp256k1.PubKey{Key: pk}
	return &mockSdkTx{
		signers: []sdk.AccAddress{cosmlib.AddressToAccAddress(addr)},
		msgs:    []sdk.Msg{evmtypes.NewFromTransaction(signedTx)},
		pubKeys: []cryptotypes.PubKey{pubKey},
		signatures: []signing.SignatureV2{
			{
				PubKey: pubKey,
				// NOTE: not including the signature data for the mock
				Sequence: signedTx.Nonce(),
			},
		},
	}, nil
}

//nolint:unused // will be used later.
func buildSdkTx(from *ecdsa.PrivateKey, nonce uint64) sdk.Tx {
	pubKey := &ethsecp256k1.PubKey{Key: crypto.CompressPubkey(&from.PublicKey)}
	signer := crypto.PubkeyToAddress(from.PublicKey)
	return &mockSdkTx{
		signers: []sdk.AccAddress{cosmlib.AddressToAccAddress(signer)},
		msgs:    []sdk.Msg{},
		pubKeys: []cryptotypes.PubKey{pubKey},
		signatures: []signing.SignatureV2{
			{
				PubKey: pubKey,
				// NOTE: not including the signature data for the mock
				Sequence: nonce,
			},
		},
	}
}

func buildTx(from *ecdsa.PrivateKey, txData coretypes.TxData) (*coretypes.Transaction, sdk.Tx) {
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
				Sequence: signedEthTx.Nonce(),
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
