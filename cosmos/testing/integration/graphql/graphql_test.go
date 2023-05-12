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
package graphql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"

	"pkg.berachain.dev/polaris/cosmos/testing/integration"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tidwall/gjson"
)

var (
	tf       *integration.TestFixture
	client   *ethclient.Client
	wsclient *ethclient.Client
)

func TestRpc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/testing/jsonrpc:integration")
}

var _ = SynchronizedBeforeSuite(func() []byte {
	// Setup the network and clients here.
	tf = integration.NewTestFixture(GinkgoT())
	client = tf.EthClient
	wsclient = tf.EthWsClient
	return nil
}, func(data []byte) {})

var _ = Describe("GraphQL", func() {

	It("should support eth_blockNumber", func() {
		response, err := sendGraphQLRequest(`
		query {
			block {
				number	
			}
		}
		`)
		blockNumber := gjson.Get(string(response), "data.block.number").Int()
		Expect(err).ToNot(HaveOccurred())
		// mashallah
		Expect(blockNumber).To(BeNumerically(">=", 3349))
	})

	It("should support eth_call", func() {
		success, err := sendGraphQLRequest(`
		query {
			call(data: {
				to: "0x00000000..",
				data: "0x000000..."
			}) {
				data
				status
				gasUsed
			}
		}
		`)
		Expect(err).ToNot(HaveOccurred())
		Expect(success).ToNot(BeNil())

	})
	It("should support eth_estimateGas", func() {
		response, err := sendGraphQLRequest(`
		query {
			eth_estimateGas
		}
		`)
		Expect(err).ToNot(HaveOccurred())
		Expect(response).ToNot(BeNil())
	})
	It("should support eth_gasPrice", func() {
		response, err := sendGraphQLRequest(`
		query {
			gasPrice
		}
		`)
		Expect(err).ToNot(HaveOccurred())
		Expect(response).ToNot(BeNil())

	})
	It("should support eth_getBalance", func() {
		balance, err := sendGraphQLRequest(`
		query {
			account(address: "0x000000000000") {
				balance
			}
		}
		`)
		Expect(err).ToNot(HaveOccurred())
		Expect(balance).ToNot(BeNil())

	})
	It("should support eth_getBlockByHash", func() {

	})
	It("should support eth_getBlockByNumber", func() {

	})
	It("should support eth_getBlockTransactionCountByHash", func() {

	})
	It("should support eth_getBlockTransactionCountByNumber", func() {

	})
	It("should support eth_getCode", func() {

	})
	It("should support eth_getLogs", func() {

	})
	It("should support eth_getStorageAt", func() {

	})
	It("should support eth_getTransactionByBlockHashAndIndex", func() {

	})
	It("should support eth_getTransactionByBlockNumberAndIndex", func() {

	})
	It("should support eth_getTransactionByHash", func() {

	})
	It("should support eth_getTransactionCount", func() {

	})
	It("should support eth_getTransactionReceipt", func() {

	})
	It("should support eth_getUncleByBlockHashAndIndex", func() {

	})
	It("should support eth_getUncleByBlockNumberAndIndex", func() {

	})
	It("should support eth_getUncleCountByBlockHash", func() {

	})
	It("should support eth_getUncleCountByBlockNumber", func() {

	})
	It("should support eth_protocolVersion", func() {

	})
	It("should support eth_sendRawTransaction", func() {

	})

	It("should support eth_syncing", func() {

	})

	It("should subscribe to the chain", func() {

	})

	It("should fail on a malformatted query", func() {

	})

	It("should fail on a malformatted mutation", func() {

	})

})

func sendGraphQLRequest(query string) ([]byte, error) {
	url := "http://localhost:8545/graphql"
	requestBody, err := json.Marshal(map[string]string{
		"query": query,
	})
	if err != nil {
		fmt.Println("Error while creating the request body:", err)
		return nil, err
	}

	// Create an HTTP client
	client := &http.Client{}

	// Create a POST request with the GraphQL endpoint URL and request body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error while creating the request:", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while sending the request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading the response:", err)
		return nil, err
	}

	// ugly asf
	ok := gjson.Get(string(responseBody), "data.errors")
	if ok.Exists() {
		panic("Malformatted request.")
	}
	return responseBody, nil
}
