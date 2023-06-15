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

package main

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"pkg.berachain.dev/polaris/eth/core/types"
)

const POLARIS_RPC = "http://localhost:8545"
const TESTS = "./e2e/compatriot/tests.json"

var client *ethclient.Client
var txHashes []common.Hash

// setup starts up the chain and spams the transactions
func setup() error {
	txHashes = submitTransactionsToNetwork()
	requests = generateQueries()
	return nil
}

// connectToClient connects to an Ethereum client and returns the client instance
func connectToClient(url string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// SignTransaction signs the given transaction with the provided private key and returns the signed transaction object
func signTransaction(tx *types.Transaction, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {

	signedTx, err := types.SignTx(tx, coretypes.NewEIP155Signer(big.NewInt(2061)), privateKey)
	if err != nil {
		panic(err)
	}

	return signedTx, nil
}

func buildTx(address common.Address, nonce uint64) (*coretypes.Transaction, error) {

	toAddress := common.HexToAddress("0x00000000000000000000000000000000DeaDBeef") // Replace with the recipient's Ethereum address
	value := big.NewInt(1000000000000000000)                                       // 1 ETH in wei
	gasLimit := uint64(21000)                                                      // Standard gas limit for a simple transaction
	data := []byte{}                                                               // Optional data for contract interactions

	return coretypes.NewTransaction(nonce, toAddress, value, gasLimit, big.NewInt(0), data), nil
}

// sendTx sends a transaction to the deadbeef address and returns its hash
func sendTx(nonce uint64) (common.Hash, error) {
	client, err := connectToClient("http://localhost:8545") // Replace with your Ethereum client URL
	if err != nil {
		panic(err)
	}

	tx, err := buildTx(common.HexToAddress("0x20f33CE90A13a4b5E7697E3544c3083B8F8A51D4"), nonce)
	if err != nil {
		panic(err)
	}

	// Sign the transaction
	privKey, err := crypto.ToECDSA(common.Hex2Bytes("fffdbb37105441e14b0ee6330d855d8504ff39e705c3afa8f859ac9865f99306"))
	if err != nil {
		panic(err)
	}
	signedTx, err := signTransaction(tx, privKey)
	if err != nil {
		panic(err)
	}
	// Send the transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		panic(err)
	}
	return signedTx.Hash(), nil
}

// submitTransactionsToNetwork submits transactions to the network and returns all the txHashes
func submitTransactionsToNetwork() []common.Hash {
	for i := 0; i < 100; i++ {

		txHash, err := sendTx(uint64(i))
		if err != nil {
			panic(err)
		}
		// fmt.Println("sent transaction to network", txHash)
		txHashes = append(txHashes, txHash)
	}
	return txHashes
}

// generateQueries generates the queries to be sent to the chain for every transaction
func generateQueries() []RPCRequest {
	var requests []RPCRequest
	id := 0
	for _, txHash := range txHashes {
		hash := txHash.String()

		transactionByHashRequest := RPCRequest{"2.0", "eth_getTransactionByHash", []interface{}{hash}, int64(id)}

		// blockByNumberRequest := RPCRequest{"2.0", "eth_getBlockByNumber", []interface{}{hash, false}, int64(id+1)}
		// blockByHashRequest := RPCRequest{"2.0", "eth_getBlockByHash", []interface{}{hash, false}, int64(id+2)}
		// receiptsByHashRequest := RPCRequest{"2.0", "eth_getTransactionReceipt", []interface{}{hash}, int64(id+3)}

		id += 1
		requests = append(requests, transactionByHashRequest)
	}

	return requests
}
