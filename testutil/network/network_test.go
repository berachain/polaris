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

package network_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"pkg.berachain.dev/stargazer/testutil/network"
)

func TestNetwork(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "testutil/network:integration")
}

var _ = Describe("NetworkRPC", func() {
	var net *network.Network
	BeforeEach(func() {
		net = network.New(GinkgoT(), network.DefaultConfig())
		_, err := net.WaitForHeightWithTimeout(3, 30*time.Second)
		Expect(err).To(BeNil())

	})

	It("eth_chainId", func() {
		client, err := ethclient.Dial(net.Validators[0].APIAddress + "/eth/rpc")
		Expect(err).To(BeNil())
		chainID, err := client.ChainID(context.Background())
		Expect(err).To(BeNil())
		Expect(chainID.String()).To(Equal("42069"))
	})

	It("eth_blockNumber", func() {
		// Dial GetBlock
		client, err := ethclient.Dial(net.Validators[0].APIAddress + "/eth/rpc")
		Expect(err).To(BeNil())
		blockNumber, err := client.BlockNumber(context.Background())
		Expect(err).To(BeNil())
		Expect(blockNumber).To(BeNumerically(">", 0))
	})

	It("eth_getBlockByNumber", func() {
		client, err := ethclient.Dial(net.Validators[0].APIAddress + "/eth/rpc")
		Expect(err).To(BeNil())
		blockNumber, err := client.BlockNumber(context.Background())
		Expect(err).To(BeNil())
		block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
		Expect(err).To(BeNil())
		Expect(block.Number().Uint64()).To(Equal(blockNumber))
	})

	It("eth_getBlockByHash", func() {

		// TODO: expected failure because offchain kv is not working yet, this reads from offchain kv
		client, err := ethclient.Dial(net.Validators[0].APIAddress + "/eth/rpc")
		Expect(err).To(BeNil())
		block, err := client.BlockByNumber(context.Background(), big.NewInt(1))
		Expect(err).To(BeNil())
		hash := block.Hash()
		blockByHash, err := client.BlockByHash(context.Background(), hash)
		Expect(err).To(BeNil())
		Expect(blockByHash.Hash()).To(Equal(hash))
	})

	// AfterEach(func() {
	// 	net = network.New(GinkgoT(), network.DefaultConfig())
	// 	_, err := net.(3, 15*time.Second)
	// 	net.
	// 	Expect(err).To(BeNil())

	// })
})
