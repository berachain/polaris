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

package txpool

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	"cosmossdk.io/log"

	"github.com/berachain/polaris/cosmos/runtime/txpool/mocks"

	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTxpool(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/runtime/txpool")
}

var _ = Describe("", func() {
	var h *handler

	var subscription *mocks.Subscription
	var serializer *mocks.TxSerializer
	var broadcaster *mocks.TxBroadcaster
	var subprovider *mocks.TxSubProvider

	BeforeEach(func() {
		t := GinkgoT()
		defer GinkgoRecover()
		subscription = mocks.NewSubscription(t)
		subscription.On("Err").Return(nil)
		subscription.On("Unsubscribe").Return()
		broadcaster = mocks.NewTxBroadcaster(t)
		subprovider = mocks.NewTxSubProvider(t)
		subprovider.On("SubscribeTransactions", mock.Anything, mock.Anything).Return(subscription)
		serializer = mocks.NewTxSerializer(t)
		h = newHandler(broadcaster, subprovider, serializer, newCometRemoteCache(), log.NewTestLogger(t))
		err := h.Start()
		Expect(err).NotTo(HaveOccurred())
		for !h.Running() {
			// Wait for handler to start.
			time.Sleep(50 * time.Millisecond)
		}
		Expect(h.Running()).To(BeTrue())
	})

	AfterEach(func() {
		err := h.Stop()
		Expect(err).NotTo(HaveOccurred())
		for h.Running() {
			// Wait for handler to start.
			time.Sleep(50 * time.Millisecond)
		}
		Expect(h.Running()).To(BeFalse())
	})

	When("", func() {
		It("should handle 1 tx", func() {
			defer GinkgoRecover()
			serializer.On("ToSdkTxBytes", mock.Anything, mock.Anything).Return([]byte{123}, nil).Once()
			broadcaster.On("BroadcastTxSync", []byte{123}).Return(nil, nil).Once()

			h.txsCh <- core.NewTxsEvent{
				Txs: []*ethtypes.Transaction{ethtypes.NewTx(&ethtypes.LegacyTx{Nonce: 5, Gas: 100})},
			}
		})

		It("should handle multiple tx", func() {
			defer GinkgoRecover()
			serializer.On("ToSdkTxBytes", mock.Anything, mock.Anything).Return([]byte{123}, nil).Twice()
			broadcaster.On("BroadcastTxSync", []byte{123}).Return(nil, nil).Twice()

			h.txsCh <- core.NewTxsEvent{Txs: []*ethtypes.Transaction{
				ethtypes.NewTx(&ethtypes.LegacyTx{Nonce: 5, Gas: 10}),
				ethtypes.NewTx(&ethtypes.LegacyTx{Nonce: 6, Gas: 10}),
			}}
		})
	})

})
