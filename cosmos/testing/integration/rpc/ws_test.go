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

package jsonrpc

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"

	geth "github.com/ethereum/go-ethereum"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	gethrpc "github.com/ethereum/go-ethereum/rpc"

	"pkg.berachain.dev/polaris/cosmos/testing/network"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "pkg.berachain.dev/polaris/cosmos/testing/integration/utils"
)

var _ = Describe("Network with WS", func() {
	var wsClient *ethclient.Client
	var net *network.Network
	var ctx context.Context
	BeforeEach(func() {
		net, _ = StartPolarisNetwork(GinkgoT())
	})

	AfterEach(func() {
		// TODO: FIX THE OFFCHAIN DB
		os.RemoveAll("data")
	})
	Context("checking websockets", func() {
		BeforeEach(func() {
			ctx = context.Background()
			// Dial an Ethereum RPC Endpoint
			wsaddr := "ws:" + strings.TrimPrefix(net.Validators[0].APIAddress+"/eth/rpc/ws/", "http:")
			// rpcaddr := net.Validators[0].APIAddress + "/eth/rpc"
			ws, err := gethrpc.DialWebsocket(ctx, wsaddr, "*")
			Expect(err).ToNot(HaveOccurred())
			wsClient = ethclient.NewClient(ws)
			Expect(err).ToNot(HaveOccurred())
		})
		It("should connect", func() {
			// Dial an Ethereum websocket Endpoint
			wsaddr := "ws:" + strings.TrimPrefix(net.Validators[0].APIAddress+"/eth/rpc/ws/", "http:")
			// rpcaddr := net.Validators[0].APIAddress + "/eth/rpc"
			ws, err := gethrpc.DialWebsocket(ctx, wsaddr, "*")
			Expect(err).ToNot(HaveOccurred())
			wsClient = ethclient.NewClient(ws)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should subscribe to new heads", func() {
			// Subscribe to new heads
			sub, err := wsClient.SubscribeNewHead(ctx, make(chan *gethtypes.Header))
			Expect(err).ToNot(HaveOccurred())
			Expect(sub).ToNot(BeNil())
		})

		It("should subscribe to logs", func() {
			// Subscribe to logs
			sub, err := wsClient.SubscribeFilterLogs(ctx, geth.FilterQuery{}, make(chan gethtypes.Log))
			Expect(err).ToNot(HaveOccurred())
			Expect(sub).ToNot(BeNil())
		})
		It("should get recent blocks", func() {
			headers := make(chan *gethtypes.Header)
			sub, _ := wsClient.SubscribeNewHead(ctx, headers)
			fmt.Print("Listening...")
			select {
			case err := <-sub.Err():
				Fail(fmt.Sprintf("Error in subscription for recent blocks: %v", err))
			case header := <-headers:
				fmt.Printf("New block: %v", header.Number.Uint64())
				_, err := wsClient.BlockByNumber(ctx, big.NewInt(header.Number.Int64()))
				Expect(err).ToNot(HaveOccurred())
			}
		})
	})

})
