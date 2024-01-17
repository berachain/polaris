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
