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
	"errors"
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
		response, status, err := sendGraphQLRequest(`
		query {
			block {
				number	
			}
		}
		`)
		blockNumber := gjson.Get(response, "data.block.number").Int()

		Expect(status).To(Equal(200))
		Expect(err).ToNot(HaveOccurred())
		Expect(blockNumber).To(BeNumerically(">", 0))
	})

	It("should support eth_call", func() {

	})
	It("should support eth_estimateGas", func() {

	})
	It("should support eth_gasPrice", func() {
		response, status, err := sendGraphQLRequest(`
		query {
			gasPrice
		}
		`)
		gasPrice := gjson.Get(response, "data.gasPrice").String()
		Expect(status).To(Equal(200))
		Expect(err).ToNot(HaveOccurred())
		Expect(gasPrice).To(BeComparableTo("0x3b9aca07"))

	})

	It("should support eth_getBlockByHash", func() {
		response, status, err := sendGraphQLRequest(`
		query {
			block(hash:"0x1ddcdaaef4dc4b7ae80ce5f23383de2168311dfbba1fc2dd9a4fa4547d0264d6") {
			  transactionCount
			  baseFeePerGas
			  nextBaseFeePerGas
			  ommerCount
			}
		  }`)
		transactionCount := gjson.Get(response, "data.block.transactionCount").Int()
		ommerCount := gjson.Get(response, "data.block.ommerCount").Int()
		Expect(status).To((BeEquivalentTo(200)))
		Expect(err).ToNot(HaveOccurred())
		Expect(transactionCount).To(BeEquivalentTo(0))
		Expect(ommerCount).To(BeEquivalentTo(0))
	})
	It("should support eth_getBlockByNumber", func() {
		response, status, err := sendGraphQLRequest(`
		query {
			block(number:"0") {
			  transactionCount
			  baseFeePerGas
			  nextBaseFeePerGas
			  ommerCount
			}
		  }`)
		transactionCount := gjson.Get(response, "data.block.transactionCount").Int()
		ommerCount := gjson.Get(response, "data.block.ommerCount").Int()
		Expect(status).To((BeEquivalentTo(200)))
		Expect(err).ToNot(HaveOccurred())
		Expect(transactionCount).To(BeEquivalentTo(0))
		Expect(ommerCount).To(BeEquivalentTo(0))

	})

	It("should support eth_getBalance, eth_getCode, eth_getStorageAt, eth_getTransactionCount", func() {
		response, status, err := sendGraphQLRequest(`
		{	
			block {
			  account(address: "0x0000000000000000000000000000000000000000") {
				balance
				code
				storage(slot: "0x044852b2a670ade5407e78fb2863c51de9fcb96542a07186fe3aeda6bb8a116d")
				transactionCount
			  }
			}
		  }
	`)
		balance := gjson.Get(response, "data.block.account.balance").String()
		code := gjson.Get(response, "data.block.account.balance").String()
		storage := gjson.Get(response, "data.block.account.balance").String()
		transactionCount := gjson.Get(response, "data.block.account.balance").String()

		Expect(status).To(BeEquivalentTo(200))
		Expect(err).ToNot(HaveOccurred())
		Expect(balance).To(BeEquivalentTo("0"))
		Expect(code).To(BeEquivalentTo("0"))
		Expect(storage).To(BeEquivalentTo("0"))
		Expect(transactionCount).To(BeEquivalentTo("0"))

	})
	It("should support eth_getTransactionByBlockHashAndIndex", func() {
		response, status, err := sendGraphQLRequest(`
		{
			block(hash: "0x1ddcdaaef4dc4b7ae80ce5f23383de2168311dfbba1fc2dd9a4fa4547d0264d6") {
			  transactionAt(index: 0) {
				hash
				nonce
				index
				value
				gasPrice
				maxFeePerGas
				maxPriorityFeePerGas
				effectiveTip
				effectiveGasPrice
				gas
				inputData
			  }
			}
		  }
		  
	`)
		transactionAt := gjson.Get(response, "data.block.transactionAt").Exists()

		Expect(status).To(BeEquivalentTo(200))
		Expect(err).ToNot(HaveOccurred())
		Expect(transactionAt).To(BeNil())
	})
	It("should support eth_getTransactionByBlockNumberAndIndex", func() {
		response, status, err := sendGraphQLRequest(`
		{
			block(number: 0) {
			  transactionAt(index: 0) {
				hash
				nonce
				index
				value
				gasPrice
				maxFeePerGas
				maxPriorityFeePerGas
				effectiveTip
				effectiveGasPrice
				gas
				inputData
			  }
			}
		  }
		  
	`)
		transactionAt := gjson.Get(response, "data.block.transactionAt").Exists()

		Expect(status).To(BeEquivalentTo(200))
		Expect(err).ToNot(HaveOccurred())
		Expect(transactionAt).To(BeNil())
	})
	It("should support eth_getTransactionByHash", func() {
		response, status, err := sendGraphQLRequest(`
		{
			transaction(hash:"0x0000000000000000000000000000000000000000000000000000000000000000") {
			  index
			  maxFeePerGas
			  maxPriorityFeePerGas
			  effectiveTip
			  status
			  gasUsed
			  cumulativeGasUsed
			  effectiveGasPrice
			  type
			}
		  }`)

		transactionAt := gjson.Get(response, "data.block.transactionAt").Exists()

		Expect(status).To(BeEquivalentTo(200))
		Expect(err).ToNot(HaveOccurred())
		Expect(transactionAt).To(BeNil())
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

	It("should fail on a malformatted query", func() {
		_, status, _ := sendGraphQLRequest(`
		query {
			ooga
			booga
		}
		`)
		Expect(status).Should(Equal(400))
		//Expect(err).To(HaveOccurred())
	})

	It("should fail on a malformatted mutation", func() {

	})
})

func sendGraphQLRequest(query string) (string, int, error) {
	// 500 = random
	url := "http://localhost:8545/graphql"
	requestBody, err := json.Marshal(map[string]string{
		"query": query,
	})
	if err != nil {
		fmt.Println("Error while creating the request body:", err)
		return "", 500, err
	}

	// Create an HTTP client
	client := &http.Client{}

	// Create a POST request with the GraphQL endpoint URL and request body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error while creating the request:", err)
		return "", 500, err
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while sending the request:", err)
		return "", 500, err
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	responseStatusCode := resp.StatusCode
	if err != nil {
		fmt.Println("Error while reading the response:", err)
		return "", 500, err
	}

	// ugly asf
	ok := gjson.Get(string(responseBody), "data.errors")
	if ok.Exists() {
		return "", 400, errors.New(ok.String())
	}
	return string(responseBody), responseStatusCode, nil
}
