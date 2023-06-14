package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
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
		fmt.Println("sent transaction to network", txHash)
		txHashes = append(txHashes, txHash)
	}
	return txHashes
}

func generateQueries() []RPCRequest {
	var requests []RPCRequest
	for id, txHash := range txHashes {
		/*
		   sendTx() returns hash of the send transaction

		   GetReceiptsByHash() returns receipts of the transaction

		   get BlockNumber and BlockHash from the Receipt

		   Then call GetBlockByNumber() with the block number

		   Then call GetBlockByHash() with the block hash

		   Then call GetTransactionByHash() with the transaction hash

		   Then call GetReceiptsByHash() with the transaction hash

		   on the first run, these will all work beacuse of the cache

		   then when we stop the node, nuke the cache, and run again, these will all fail because no more cache and historical plugin gone

		*/
		request := RPCRequest{"2.0", "eth_getTransactionByHash", []interface{}{txHash.String()}, int64(id)}
		requests = append(requests, request)
	}

	return requests
}
