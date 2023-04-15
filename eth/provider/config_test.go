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

package provider

import (
	"math/big"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	filepath = ".config.example.toml"
)

func TestProvider(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "eth/provider")
}

var _ = Describe("TestConfig", func() {
	It("should read in the config correctly", func() {
		config, err := ReadConfigFile(filepath)
		Expect(err).ToNot(HaveOccurred())
		Expect(config.ChainConfig.ChainID).To(Equal(big.NewInt(69420)))
		Expect(config.ChainConfig.HomesteadBlock).To(Equal(big.NewInt(0)))
		Expect(config.NodeConfig.UserIdent).To(Equal("my-identity"))
		Expect(config.NodeConfig.DataDir).To(Equal("/var/data/my-node"))
		Expect(config.NodeConfig.HTTPHost).To(Equal("localhost"))
		Expect(config.NodeConfig.HTTPPort).To(BeNumerically("==", 8545))
		Expect(config.NodeConfig.HTTPCors).To(ContainElement("*"))
		Expect(config.NodeConfig.HTTPVirtualHosts).To(ContainElement("localhost"))
		Expect(config.NodeConfig.HTTPModules).To(ConsistOf("eth", "net"))
		Expect(config.NodeConfig.AuthAddr).To(Equal("localhost"))
		Expect(config.NodeConfig.AuthPort).To(BeNumerically("==", 8546))
		Expect(config.NodeConfig.AuthVirtualHosts).To(ContainElement("localhost"))
		Expect(config.NodeConfig.WSHost).To(Equal("localhost"))
		Expect(config.NodeConfig.WSPort).To(BeNumerically("==", 8546))
		Expect(config.NodeConfig.WSOrigins).To(ContainElement("*"))
		Expect(config.NodeConfig.WSModules).To(ConsistOf("eth", "net"))
		Expect(config.NodeConfig.GraphQLCors).To(ContainElement("*"))
		Expect(config.NodeConfig.GraphQLVirtualHosts).To(ContainElement("localhost"))
		Expect(config.RPCConfig.RPCGasCap).To(BeNumerically("==", 10000000))
		Expect(config.RPCConfig.RPCEVMTimeout).To(Equal(10 * time.Second))
		Expect(config.RPCConfig.RPCTxFeeCap).To(Equal(float64(1)))
		Expect(config.RPCConfig.GPO.Blocks).To(BeNumerically("==", 10))
		Expect(config.RPCConfig.GPO.Percentile).To(BeNumerically("==", 50))
		Expect(config.RPCConfig.GPO.MaxHeaderHistory).To(BeNumerically("==", 192))
		Expect(config.RPCConfig.GPO.MaxBlockHistory).To(BeNumerically("==", 5000))
		Expect(config.RPCConfig.GPO.Default).To(Equal(big.NewInt(1000000000)))
		Expect(config.RPCConfig.GPO.MaxPrice).To(Equal(big.NewInt(100000000000)))
		Expect(config.RPCConfig.GPO.IgnorePrice).To(Equal(big.NewInt(0)))
	})
})
