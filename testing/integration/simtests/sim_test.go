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

package simtests

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"

	network "pkg.berachain.dev/stargazer/testing/utils/network"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestNetwork(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "integration")
}

var _ = Describe("SimulationTests", func() {
	var net *network.Network
	// var client *ethclient.Client
	//test

	BeforeEach(func() {
		cfg := network.ConfigWithTestAccount()
		net = network.New(GinkgoT(), cfg)
		_, err := net.WaitForHeightWithTimeout(1, 15*time.Second)
		Expect(err).ToNot(HaveOccurred())
		_, err = ethclient.Dial(net.Validators[0].APIAddress + "/eth/rpc")
		Expect(err).ToNot(HaveOccurred())

	})

	It("should be able to send a transaction and verify it's been received with receipt", func() {
		//TODO: implement
		Expect(true).To(BeTrue())
	})
})
