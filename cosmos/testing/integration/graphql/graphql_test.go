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
	"io"
	"math/big"
	"net/http"
	"strconv"
	"testing"

	"github.com/tidwall/gjson"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"

	"pkg.berachain.dev/polaris/cosmos/testing/integration"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "pkg.berachain.dev/polaris/cosmos/testing/integration/utils"
)

var (
	tf     *integration.TestFixture
	client *ethclient.Client
)

func TestRpc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/testing/graphql:integration")
}

var _ = SynchronizedBeforeSuite(func() []byte {
	// Setup the network and clients here.
	tf = integration.NewTestFixture(GinkgoT())
	client = tf.EthClient
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
		_, addr := DeployERC20(tf.GenerateTransactOpts("alice"), client)
		// function selector for decimals() padded to 32 bytes
		calldata := "0x313ce56700000000000000000000000000000000000000000000000000000000"
		query := fmt.Sprintf(`
		query { 
			block(number:1) { 
				call(data: { to: "%s", data: "%s" }) {
					data
					status
					gasUsed
				} 
			} 
		}`, addr.String(), calldata)
		_, status, err := sendGraphQLRequest(query)

		Expect(status).To(Equal(200))
		Expect(err).ToNot(HaveOccurred())
	})
	It("should support eth_estimateGas", func() {
		alice := tf.Address("alice")
		response, status, err := sendGraphQLRequest(fmt.Sprintf(
			`query { 
				block(number: "1") { 
					estimateGas( data: { to: "%s" }) 
					} 
			}`, alice.String()))

		estimateGas := gjson.Get(response, "data.block.estimateGas").String()
		Expect(status).To(Equal(200))
		Expect(strconv.ParseUint(estimateGas, 10, 64)).To(BeNumerically(">=", 21000))
		Expect(err).ToNot(HaveOccurred())
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
		Expect(strconv.ParseUint(gasPrice[2:], 16, 64)).To(BeNumerically(">", 0))

	})

	It(`should support eth_getBlockByHash, eth_getBlockTransactionCountByHash, 
	eth_getUncleByBlockHashAndIndex, eth_getUncleCountByBlockHash`, func() {

		mostRecentBlockHash, err := getMostRecentBlockHash()
		Expect(err).ToNot(HaveOccurred())
		query := fmt.Sprintf(`
		query {
			block(hash: "%s") {
				transactionCount
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
				ommerCount
				ommerAt(index: 0) {
				  number
				  hash
				  rawHeader
				}	
			}
		}`, mostRecentBlockHash)
		response, status, err := sendGraphQLRequest(query)
		transactionCount := gjson.Get(response, "data.block.transactionCount").Int()
		ommerCount := gjson.Get(response, "data.block.ommerCount").Int()
		Expect(status).To((BeEquivalentTo(200)))
		Expect(err).ToNot(HaveOccurred())
		Expect(transactionCount).To(BeNumerically(">=", 0))
		Expect(ommerCount).To(BeNumerically(">=", 0))
	})
	It(`should support eth_getBlockByNumber, eth_getBlockTransactionCountByNumber, 
	eth_getUncleCountByBlockNumber, eth_getUncleByBlockNumberAndIndex`, func() {

		response, status, err := sendGraphQLRequest(`
		query {
			block(number: 0) {
			  transactionCount
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
			  ommerCount
			  ommerAt(index: 0) {
				number
				hash
				rawHeader
			  }
			}
		  }`)
		transactionCount := gjson.Get(response, "data.block.transactionCount").Int()
		transactionAt := gjson.Get(response, "data.block.transactionAt").String()
		ommerCount := gjson.Get(response, "data.block.ommerCount").Int()
		ommerAt := gjson.Get(response, "data.block.ommerAt").String()
		Expect(err).ToNot(HaveOccurred())
		Expect(status).To((BeEquivalentTo(200)))

		Expect(transactionCount).To(BeEquivalentTo(0))
		Expect(ommerCount).To(BeEquivalentTo(0))

		Expect(transactionAt).ToNot(BeNil())
		Expect(ommerAt).ToNot(BeNil())
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
		Expect(balance).To(BeEquivalentTo("0x0"))
		Expect(code).To(BeEquivalentTo("0x0"))
		Expect(storage).To(BeEquivalentTo("0x0"))
		Expect(transactionCount).To(BeEquivalentTo("0x0"))

	})
	It("should support eth_getTransactionByBlockHashAndIndex", func() {
		mostRecentBlockHash, err := getMostRecentBlockHash()
		Expect(err).ToNot(HaveOccurred())
		query := fmt.Sprintf(`
		query {
			block(hash: "%s") {
				transactionCount
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
				ommerCount
				ommerAt(index: 0) {
				  number
				  hash
				  rawHeader
				}	
			}
		}`, mostRecentBlockHash)
		response, status, err := sendGraphQLRequest(query)
		transactionAt := gjson.Get(response, "data.block.transactionAt").Exists()

		Expect(status).To(BeEquivalentTo(200))
		Expect(err).ToNot(HaveOccurred())
		Expect(transactionAt).To(BeTrue())
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
		Expect(transactionAt).To(BeTrue())
	})
	It("should support eth_getTransactionByHash and eth_getTransactionReceipt", func() {
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

		transactionAt := gjson.Get(response, "data.transaction").Exists()

		Expect(status).To(BeEquivalentTo(200))
		Expect(err).ToNot(HaveOccurred())
		Expect(transactionAt).To(BeTrue())
	})

	It("should support eth_getLogs", func() {
		response, status, err := sendGraphQLRequest(`query {
					logs(filter: {
					  topics: []
					}) {
					  index
					}
				  }`)

		logs := gjson.Get(response, "data.logs").Exists()

		Expect(logs).To(BeTrue())
		Expect(status).Should(Equal(200))
		Expect(err).ToNot(HaveOccurred())
	})

	It("should support eth_protocolVersion", func() {
		// I don't even think Geth supports this
		// even though it says it does here:
		// https://eips.ethereum.org/EIPS/eip-1767
	})
	It("should support eth_sendRawTransaction", func() {
		alicePrivKey := tf.PrivKey("alice")
		bob := tf.Address("bob")
		tx := types.NewTransaction(
			0, // Nonce
			bob,
			big.NewInt(5),  // Value
			uint64(22000),  // Gas limit
			big.NewInt(10), // Gas price
			nil,
		)
		signed, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(69420)), alicePrivKey)
		Expect(err).ToNot(HaveOccurred())
		rlpEncoded, err := rlp.EncodeToBytes(signed)
		Expect(err).ToNot(HaveOccurred())
		data := fmt.Sprintf("mutation { sendRawTransaction(data: \"0x%x\") }", rlpEncoded)
		_, status, err := sendGraphQLRequest(data)
		Expect(err).ToNot(HaveOccurred())
		Expect(status).Should(Equal(200))
	})

	It("should support eth_syncing", func() {
		response, status, err := sendGraphQLRequest(`
		query {
			syncing {
			  startingBlock
			}
		  }`)
		syncing := gjson.Get(response, "data.syncing").Value()

		Expect(status).To(BeEquivalentTo(200))
		Expect(err).ToNot(HaveOccurred())
		Expect(syncing).To(BeNil())
	})
	It("should fail on a malformatted query", func() {
		response, status, _ := sendGraphQLRequest(`
		query {
			ooga
			booga
		}
		`)
		errorMessage := gjson.Get(response, "data.errors.message")
		Expect(errorMessage).ToNot(BeNil())
		Expect(status).Should(Equal(400))
	})

	It("should fail on a malformatted mutation", func() {
		response, status, _ := sendGraphQLRequest(`
		mutation {
			ooga
			booga
		}
		`)
		errorMessage := gjson.Get(response, "data.errors.message")
		Expect(errorMessage).ToNot(BeNil())
		Expect(status).Should(Equal(400))
	})
})

func getMostRecentBlockHash() (string, error) {
	err := tf.Network.WaitForNextBlock()
	Expect(err).ToNot(HaveOccurred())
	mostRecentBlockHashQueryResponse, _, err := sendGraphQLRequest(`
	query {
		block {
		  hash
		}
	  }`)
	mostRecentBlockHash := gjson.Get(mostRecentBlockHashQueryResponse, "data.block.hash").String()

	if err != nil {
		return "", err
	}

	return mostRecentBlockHash, err
}

func sendGraphQLRequest(query string) (string, int, error) {
	url := "http://localhost:8545/graphql"
	requestBody, err := json.Marshal(map[string]string{
		"query": query,
	})
	if err != nil {
		return "", 500, err
	}

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", 500, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", 500, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	responseStatusCode := resp.StatusCode
	if err != nil {
		return "", 500, err
	}

	return string(responseBody), responseStatusCode, nil
}
