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
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	cmtclient "github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	ethabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	buildertypes "github.com/skip-mev/pob/x/builder/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	bindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"
	beracodec "pkg.berachain.dev/polaris/cosmos/crypto/codec"
	evmtypes "pkg.berachain.dev/polaris/cosmos/x/evm/types"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
)

type (
	Account struct {
		Address    common.Address
		PrivateKey *ecdsa.PrivateKey
	}

	EncodingConfig struct {
		InterfaceRegistry codectypes.InterfaceRegistry
		Codec             codec.Codec
		TxConfig          client.TxConfig
		Amino             *codec.LegacyAmino
	}

	ScriptConfig struct {
		// EthRPCURL is the URL of the Ethereum RPC endpoint
		EthRPCURL string
		// CosmosRPCURL is the URL of the Cosmos RPC endpoint
		CosmosRPCURL string
		// SearcherPrivateKey is the private key of the account that will init accounts and bid
		Searcher Account
		// ChainID is the chain ID of the BeraChain network
		ChainID int64
		// AuctionSmartContractAddress is the address of the auction smart contract
		AuctionSmartContractAddress string
		// NumAccounts is the number of accounts to init
		NumAccounts int
		// InitBalance is the initial balance of each account
		InitBalance *big.Int
		// TestAccounts
		TestAccounts []Account
		// EncodingConfig is the encoding config for the application
		EncodingConfig EncodingConfig
	}
)

var (
	CONFIG = DefaultConfig()
)

func main() {
	if err := initAccounts(); err != nil {
		panic(err)
	}

	params, err := getAuctionParams()
	if err != nil {
		panic(err)
	}
	reserveFee := params.ReserveFee.Amount.Int64()
	minBidIncrement := params.MinBidIncrement.Amount.Int64()

	// 1.
	wrapTestCase("Invalid auction bid with a low bid", func() {
		bundle := []*types.Transaction{
			createBasicEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, CONFIG.TestAccounts[0].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 1),
		}
		bidTx := createBidEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, big.NewInt(0), big.NewInt(params.ReserveFee.Amount.Int64()-1), bundle, 10000, 300000, 0)

		// We expect this to error out
		height, err := sendEthTx(bidTx)
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		}

		// No transactions should have been included in the block
		waitForABlock()
		displayExpectedOrder(bidTx, bundle, "Bid")
		displayBlock(height)
	})

	// 2.
	wrapTestCase("Invalid auction bid with too many auction transactions in bundle", func() {
		bundle := []*types.Transaction{
			createBasicEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, CONFIG.TestAccounts[0].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 1),
			createBasicEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, CONFIG.TestAccounts[0].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 2),
			createBasicEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, CONFIG.TestAccounts[0].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 3),
			createBasicEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, CONFIG.TestAccounts[0].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 4),
			createBasicEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, CONFIG.TestAccounts[0].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 5),
			createBasicEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, CONFIG.TestAccounts[0].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 6),
		}
		bidTx := createBidEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, big.NewInt(0), big.NewInt(params.ReserveFee.Amount.Int64()-1), bundle, 10000, 300000, 0)

		// We expect this to error out
		height, err := sendEthTx(bidTx)
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		}

		// No transactions should have been included in the block
		waitForABlock()
		displayExpectedOrder(bidTx, bundle, "Bid")
		displayBlock(height)
	})

	// 3.
	wrapTestCase("Valid auction transaction with a single bundle tx", func() {
		bundle := []*types.Transaction{
			createBasicEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, CONFIG.TestAccounts[0].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 1),
		}

		bidTx := createBidEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, big.NewInt(0), big.NewInt(params.ReserveFee.Amount.Int64()), bundle, 10000, 300000, 0)
		height, err := sendEthTx(bidTx)
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		}

		waitForABlock()
		displayExpectedOrder(bidTx, bundle, "Bid")
		displayBlock(height)
	})

	// 4.
	wrapTestCase("Invalid auction transaction that sends money to the auction smart contract", func() {
		bundle := []*types.Transaction{
			createBasicEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, CONFIG.TestAccounts[0].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 1),
		}

		// We expect this to error out
		bidTx := createBidEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, big.NewInt(1), big.NewInt(params.ReserveFee.Amount.Int64()), bundle, 10000, 300000, 0)
		height, err := sendEthTx(bidTx)
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		}

		waitForABlock()
		displayExpectedOrder(bidTx, bundle, "Bid")
		displayBlock(height)
	})

	// 5.
	wrapTestCase("Invalid auction transaction that has an invalid timeout set", func() {
		bundle := []*types.Transaction{
			createBasicEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, CONFIG.TestAccounts[0].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 1),
		}

		// We expect this to error out
		bidTx := createBidEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, big.NewInt(0), big.NewInt(params.ReserveFee.Amount.Int64()), bundle, 10, 300000, 0)
		height, err := sendEthTx(bidTx)
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		}

		waitForABlock()
		displayExpectedOrder(bidTx, bundle, "Bid")
		displayBlock(height)
	})

	// 6.
	wrapTestCase("Invalid auction transaction that has no bundles", func() {
		// We expect this to error out
		bundle := []*types.Transaction{}
		bidTx := createBidEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, big.NewInt(0), big.NewInt(params.ReserveFee.Amount.Int64()), bundle, 10000, 300000, 0)
		height, err := sendEthTx(bidTx)
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		}

		waitForABlock()
		displayExpectedOrder(bidTx, bundle, "Bid")
		displayBlock(height)
	})

	// 7.
	wrapTestCase("Multiple transactions with second bid being smaller than min bid increment", func() {
		// Create the first bid
		bundle := []*types.Transaction{
			createBasicEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, CONFIG.TestAccounts[0].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 1),
		}
		bidTx := createBidEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, big.NewInt(0), big.NewInt(reserveFee), bundle, 10000, 300000, 0)

		// Second bid
		nextBundle := []*types.Transaction{
			createBasicEthTx(CONFIG.TestAccounts[0].PrivateKey, CONFIG.TestAccounts[0].Address, CONFIG.TestAccounts[1].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 1),
		}
		losingBidTx := createBidEthTx(CONFIG.TestAccounts[0].PrivateKey, CONFIG.TestAccounts[0].Address, big.NewInt(0), big.NewInt(reserveFee), nextBundle, 10000, 300000, 0)

		// Send the first bid this should not error out
		height, err := sendEthTx(bidTx)
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		}

		// Send the second bid this should error out
		_, err = sendEthTx(losingBidTx)
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		}

		waitForABlock()
		displayExpectedOrder(bidTx, bundle, "First bid")
		displayExpectedOrder(losingBidTx, nextBundle, "Second bid")
		displayBlock(height)
	})

	// 8.
	wrapTestCase("Multiple transactions with increasing bids", func() {
		// Create the first bid
		firstBundle := []*types.Transaction{
			createBasicEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, CONFIG.TestAccounts[0].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 1),
		}
		firstBid := createBidEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, big.NewInt(0), big.NewInt(reserveFee), firstBundle, 10000, 300000, 0)

		// Second bid
		secondBundle := []*types.Transaction{
			createBasicEthTx(CONFIG.TestAccounts[0].PrivateKey, CONFIG.TestAccounts[0].Address, CONFIG.TestAccounts[1].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 1),
		}
		secondBid := createBidEthTx(CONFIG.TestAccounts[0].PrivateKey, CONFIG.TestAccounts[0].Address, big.NewInt(0), big.NewInt(reserveFee+minBidIncrement), secondBundle, 10000, 300000, 0)

		// Third bid
		thirdBundle := []*types.Transaction{
			createBasicEthTx(CONFIG.TestAccounts[1].PrivateKey, CONFIG.TestAccounts[1].Address, CONFIG.TestAccounts[2].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 1),
		}
		thirdBid := createBidEthTx(CONFIG.TestAccounts[1].PrivateKey, CONFIG.TestAccounts[1].Address, big.NewInt(0), big.NewInt(reserveFee+minBidIncrement*2), thirdBundle, 10000, 300000, 0)

		height, err := sendEthTx(firstBid)
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		}

		_, err = sendEthTx(secondBid)
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		}

		_, err = sendEthTx(thirdBid)
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		}

		waitForABlock()
		displayExpectedOrder(firstBid, firstBundle, "First bid")
		displayExpectedOrder(secondBid, secondBundle, "Second bid")
		displayExpectedOrder(thirdBid, thirdBundle, "Third bid")
		displayBlock(height)

		fmt.Println("Waiting for a block to be mined")
		waitForABlock()
		height = getCurrentBlockHeight()
		displayBlock(height)

		fmt.Println("Waiting for a block to be mined")
		waitForABlock()
		height = getCurrentBlockHeight()
		waitForABlock()
		displayBlock(height)
	})

	// 9.
	wrapTestCase("Searcher is attempting to include a transaction that was included in the previous blocks", func() {
		tx := createBasicEthTx(CONFIG.TestAccounts[0].PrivateKey, CONFIG.TestAccounts[0].Address, CONFIG.TestAccounts[1].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 10000)
		height, err := sendEthTx(tx)
		if err != nil {
			panic(err)
		}

		waitForABlock()
		displayBlock(height)
		waitForABlock()

		// Create the first bid
		bundle := []*types.Transaction{
			tx,
		}
		bid := createBidEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, big.NewInt(0), big.NewInt(reserveFee), bundle, 10000, 300000, 1000)

		height, err = sendEthTx(bid)
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		}

		waitForABlock()
		displayExpectedOrder(bid, bundle, "First bid")
		displayBlock(height)
	})

	// 10.
	wrapTestCase("Searcher is creating a bundle with a transaction that is already in the mempool", func() {
		tx := createBasicEthTx(CONFIG.TestAccounts[0].PrivateKey, CONFIG.TestAccounts[0].Address, CONFIG.TestAccounts[1].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 0)
		height, err := sendEthTx(tx)
		if err != nil {
			panic(err)
		}

		bundle := []*types.Transaction{
			tx,
		}
		bid := createBidEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, big.NewInt(0), big.NewInt(reserveFee), bundle, 10000, 300000, 0)

		height, err = sendEthTx(bid)
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		}

		waitForABlock()
		displayExpectedOrder(bid, bundle, "First bid")
		displayBlock(height)
	})

	// 11.
	wrapTestCase("Multiple searchers bid with overlapping transactions", func() {
		tx := createBasicEthTx(CONFIG.TestAccounts[0].PrivateKey, CONFIG.TestAccounts[0].Address, CONFIG.TestAccounts[1].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 0)
		height, err := sendEthTx(tx)
		if err != nil {
			panic(err)
		}

		firstBundle := []*types.Transaction{
			tx,
			createBasicEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, CONFIG.TestAccounts[1].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 1),
		}
		firstBid := createBidEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, big.NewInt(0), big.NewInt(reserveFee), firstBundle, 10000, 300000, 0)

		secondBundle := []*types.Transaction{
			tx,
			createBasicEthTx(CONFIG.TestAccounts[1].PrivateKey, CONFIG.TestAccounts[1].Address, CONFIG.TestAccounts[1].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 1),
		}
		secondBid := createBidEthTx(CONFIG.TestAccounts[1].PrivateKey, CONFIG.TestAccounts[1].Address, big.NewInt(0), big.NewInt(reserveFee+minBidIncrement), secondBundle, 10000, 300000, 0)

		thirdBundle := []*types.Transaction{
			tx,
			createBasicEthTx(CONFIG.TestAccounts[2].PrivateKey, CONFIG.TestAccounts[2].Address, CONFIG.TestAccounts[1].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 1),
		}
		thirdBid := createBidEthTx(CONFIG.TestAccounts[2].PrivateKey, CONFIG.TestAccounts[2].Address, big.NewInt(0), big.NewInt(reserveFee+minBidIncrement*2), thirdBundle, 10000, 300000, 0)

		height, err = sendEthTx(firstBid)
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		}

		_, err = sendEthTx(secondBid)
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		}

		_, err = sendEthTx(thirdBid)
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		}

		waitForABlock()
		displayExpectedOrder(firstBid, firstBundle, "First bid")
		displayExpectedOrder(secondBid, secondBundle, "Second bid")
		displayExpectedOrder(thirdBid, thirdBundle, "Third bid")
		displayBlock(height)

		fmt.Println("Waiting for a block to be mined")
		waitForABlock()
		height = getCurrentBlockHeight()
		displayBlock(height)

		fmt.Println("Waiting for a block to be mined")
		waitForABlock()
		height = getCurrentBlockHeight()
		displayBlock(height)
	})

	// 12.
	wrapTestCase("Multiple searchers bid with overlapping transactions that are already in the mempool", func() {
		tx := createBasicEthTx(CONFIG.TestAccounts[0].PrivateKey, CONFIG.TestAccounts[0].Address, CONFIG.TestAccounts[1].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 0)
		height, err := sendEthTx(tx)
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		}

		tx2 := createBasicEthTx(CONFIG.TestAccounts[0].PrivateKey, CONFIG.TestAccounts[0].Address, CONFIG.TestAccounts[1].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 1)
		height, err = sendEthTx(tx2)
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		}

		firstBundle := []*types.Transaction{
			tx,
			tx2,
			createBasicEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, CONFIG.TestAccounts[1].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 1),
		}
		firstBid := createBidEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, big.NewInt(0), big.NewInt(reserveFee), firstBundle, 10000, 300000, 0)

		secondBundle := []*types.Transaction{
			tx,
			createBasicEthTx(CONFIG.TestAccounts[2].PrivateKey, CONFIG.TestAccounts[2].Address, CONFIG.TestAccounts[1].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 1),
		}
		secondBid := createBidEthTx(CONFIG.TestAccounts[2].PrivateKey, CONFIG.TestAccounts[2].Address, big.NewInt(0), big.NewInt(reserveFee+minBidIncrement), secondBundle, 10000, 300000, 0)

		height, err = sendEthTx(firstBid)
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		}

		_, err = sendEthTx(secondBid)
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		}

		waitForABlock()
		displayExpectedOrder(firstBid, firstBundle, "First bid")
		displayExpectedOrder(secondBid, secondBundle, "Second bid")
		displayBlock(height)

		fmt.Println("Waiting for a block to be mined")
		waitForABlock()
		height = getCurrentBlockHeight()
		displayBlock(height)

		fmt.Println("Waiting for a block to be mined")
		waitForABlock()
		height = getCurrentBlockHeight()
		displayBlock(height)
	})

	// 13.
	wrapTestCase("Searcher makes multiple bids with the same nonce", func() {
		bundle := []*types.Transaction{
			createBasicEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, CONFIG.TestAccounts[1].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 1),
		}
		firstBid := createBidEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, big.NewInt(0), big.NewInt(reserveFee), bundle, 10000, 300000, 0)
		secondBid := createBidEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, big.NewInt(0), big.NewInt(reserveFee+minBidIncrement), bundle, 10000, 300000, 0)
		thirdBid := createBidEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, big.NewInt(0), big.NewInt(reserveFee+minBidIncrement*2), bundle, 10000, 300000, 0)

		height, err := sendEthTx(firstBid)
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		}

		_, err = sendEthTx(secondBid)
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		}

		_, err = sendEthTx(thirdBid)
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		}

		waitForABlock()
		displayExpectedOrder(firstBid, bundle, "First bid")
		displayExpectedOrder(secondBid, bundle, "Second bid")
		displayExpectedOrder(thirdBid, bundle, "Third bid")
		displayBlock(height)

		fmt.Println("Waiting for a block to be mined")
		waitForABlock()
		height = getCurrentBlockHeight()
		displayBlock(height)

		fmt.Println("Waiting for a block to be mined")
		waitForABlock()
		height = getCurrentBlockHeight()
		displayBlock(height)
	})

	// 14.
	wrapTestCase("Searcher is attempting to front-run another user", func() {
		bundle := []*types.Transaction{
			createBasicEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, CONFIG.TestAccounts[1].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 1),
			createBasicEthTx(CONFIG.TestAccounts[0].PrivateKey, CONFIG.TestAccounts[0].Address, CONFIG.TestAccounts[1].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 0),
		}
		bid := createBidEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, big.NewInt(0), big.NewInt(reserveFee), bundle, 10000, 300000, 0)

		height, err := sendEthTx(bid)
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		}

		waitForABlock()
		displayExpectedOrder(bid, bundle, "Bid")
		displayBlock(height)
	})

	// 15.
	wrapTestCase("Searcher attempts to bid more than their balance", func() {
		bundle := []*types.Transaction{
			createBasicEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, CONFIG.TestAccounts[1].Address, big.NewInt(1000000000000000000), []byte{}, 300000, 1),
		}
		balance, err := getBalanceOf(CONFIG.Searcher.Address)
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		}

		bidAmount := big.NewInt(0).Add(balance, big.NewInt(1))
		bid := createBidEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, big.NewInt(0), bidAmount, bundle, 10000, 300000, 0)

		height, err := sendEthTx(bid)
		fmt.Println(err)
		if err != nil {
			fmt.Println("Error sending transaction:", err)
		}

		waitForABlock()
		displayExpectedOrder(bid, bundle, "Bid")
		displayBlock(height)
	})
}

// wrapTestCase wraps the test case in a function that will be executed
func wrapTestCase(name string, testCase func()) {
	fmt.Println("--------------------------------------------------")
	waitForABlock()
	log := fmt.Sprintf("Running test case: %s", name)
	fmt.Println(log)
	testCase()
	fmt.Print("--------------------------------------------------\n\n\n\n")
}

// initAccounts initializes the accounts that will be used in the bidding simulation
func initAccounts() error {
	log := fmt.Sprintf("Initializing %d accounts...", CONFIG.NumAccounts)
	fmt.Println(log)

	for i := 0; i < CONFIG.NumAccounts; i++ {
		// Create a new account
		account := createAccount()

		// Send abera to the account
		sendTx := createBasicEthTx(CONFIG.Searcher.PrivateKey, CONFIG.Searcher.Address, account.Address, CONFIG.InitBalance, []byte{}, 300000, 0)

		// Broadcast the transaction
		if _, err := sendEthTx(sendTx); err != nil {
			return err
		}
		waitForABlock()
		if _, err := getBalanceOf(account.Address); err != nil {
			return err
		}

		log = fmt.Sprintf("Account %d initialized\n", i)
		fmt.Println(log)

		CONFIG.TestAccounts = append(CONFIG.TestAccounts, account)
	}

	return nil
}

// createAccount creates a new account and returns it
func createAccount() Account {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return Account{
		Address:    fromAddress,
		PrivateKey: privateKey,
	}
}

// createBasicEthTx will create a basic Ethereum transaction with the given parameters:
// - privateKey: the private key of the sender (passed in so the tx can be signed)
// - from: the address of the sender
// - to: the address of the recipient
// - value: the amount of abera to send
// - data: the data to send
// - gasLimit: the gas limit
// - nonceOffset: the nonce offset (used to create transactions that are not in sequence)
func createBasicEthTx(privateKey *ecdsa.PrivateKey, from, to common.Address, value *big.Int, data []byte, gasLimit, nonceOffset uint64) *types.Transaction {
	client, err := getEthClient(CONFIG.EthRPCURL)
	if err != nil {
		panic(err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), from)
	if err != nil {
		panic(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	tx := types.NewTransaction(nonce+nonceOffset, to, value, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(CONFIG.ChainID)), privateKey)
	if err != nil {
		panic(err)
	}

	return signedTx
}

// createBidEthTx will create an Ethereum transaction that will call the auction smart contract precompile
// with the given parameters:
// - privateKey: the private key of the sender (passed in so the tx can be signed)
// - from: the address of the sender
// - value: the amount of abera to send
// - bid: the amount of abera to bid
// - transactions: the transactions to send
// - gasLimit: the gas limit
// - nonceOffset: the nonce offset (used to create transactions that are not in sequence)
func createBidEthTx(privateKey *ecdsa.PrivateKey, from common.Address, value, bid *big.Int, transactions []*types.Transaction, timeout, gasLimit, nonceOffset uint64) *types.Transaction {
	bundle := make([][]byte, len(transactions))
	for i, tx := range transactions {
		txBz, err := tx.MarshalBinary()
		if err != nil {
			panic(err)
		}

		bundle[i] = txBz
	}

	// Create the data to send to the auction smart contract precompile
	data, err := encodeBidData(bid, bundle, timeout)
	if err != nil {
		panic(err)
	}

	// Create the basic Ethereum transaction
	return createBasicEthTx(privateKey, from, common.HexToAddress(CONFIG.AuctionSmartContractAddress), value, data, gasLimit, nonceOffset)
}

// encodeBidData will encode the given parameters into the data that will be sent to the auction smart contract precompile
// - bid: the amount of abera to bid
// - transactions: the transactions to send
// - timeout: the timeout for the auction bid
func encodeBidData(bid *big.Int, bundle [][]byte, timeout uint64) ([]byte, error) {
	// Create the data to send to the auction smart contract precompile
	abi, err := ethabi.JSON(strings.NewReader(bindings.BuilderModuleMetaData.ABI))
	if err != nil {
		return nil, err
	}

	data, err := abi.Pack("auctionBid", bid, bundle, timeout)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// getEthTxBytes will return the bytes of the given Ethereum transaction
func getEthTxBytes(tx *types.Transaction) []byte {
	data, err := tx.MarshalBinary()
	if err != nil {
		panic(err)
	}

	return data
}

// sendEthTx will send the given Ethereum transaction to the BeraChain network
func sendEthTx(tx *types.Transaction) (int64, error) {
	client, err := getEthClient(CONFIG.EthRPCURL)
	if err != nil {
		panic(err)
	}

	err = client.SendTransaction(context.Background(), tx)
	currentHeight := getCurrentBlockHeight()
	if err != nil {
		return currentHeight + 1, err
	}

	return currentHeight + 1, nil
}

// getEthClient returns an ethclient.Client that connects to the local
// Ethereum node.
func getEthClient(rpc string) (*ethclient.Client, error) {
	return ethclient.Dial(rpc)
}

// getCosmosClient returns an grpc.ClientConn that connects to the local
// Cosmos node.
func getCosmosClient(rpc string) (*grpc.ClientConn, error) {
	grpcConn, err := grpc.Dial(
		rpc,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}

	return grpcConn, nil
}

// getAuctionParams returns the current auction parameters
func getAuctionParams() (*buildertypes.Params, error) {
	// Get the grpc connection used to query account info and broadcast transactions
	grpcConn, err := getCosmosClient(CONFIG.CosmosRPCURL)
	if err != nil {
		return nil, err
	}

	// Query the current block height
	res, err := buildertypes.NewQueryClient(grpcConn).Params(context.Background(), &buildertypes.QueryParamsRequest{})
	if err != nil {
		panic(err)
	}

	return &res.Params, err
}

// getBalanceOf returns the balance of the given address
func getBalanceOf(address common.Address) (*big.Int, error) {
	client, err := getEthClient(CONFIG.EthRPCURL)
	if err != nil {
		return nil, err
	}

	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		return nil, err
	}

	log := fmt.Sprintf("Balance of %s: %s", address.Hex(), balance.String())
	fmt.Println(log)

	return balance, nil
}

// getCurrentBlockHeight returns the current block height
func getCurrentBlockHeight() int64 {
	grpcConn, err := getCosmosClient(CONFIG.CosmosRPCURL)
	if err != nil {
		panic(err)
	}

	// Query the current block height
	grpcRes, err := cmtclient.NewServiceClient(grpcConn).GetLatestBlock(context.Background(), &cmtclient.GetLatestBlockRequest{})
	if err != nil {
		panic(err)
	}

	return grpcRes.GetSdkBlock().Header.Height
}

// waitForABlock will wait for a block to be created
func waitForABlock() {
	curr := getCurrentBlockHeight()
	for {
		newHeight := getCurrentBlockHeight()
		if newHeight > curr {
			break
		}
		time.Sleep(time.Second)
	}
}

// displayBlock will display all of the transactions in the given block
func displayBlock(height int64) {
	grpcConn, err := getCosmosClient(CONFIG.CosmosRPCURL)
	if err != nil {
		panic(err)
	}

	// Query the current block height
	grpcRes, err := cmtclient.NewServiceClient(grpcConn).GetBlockByHeight(context.Background(), &cmtclient.GetBlockByHeightRequest{Height: height})
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nActual order for height %v:\n", height)
	txs := grpcRes.GetSdkBlock().GetData().Txs
	if len(txs) == 0 {
		fmt.Println("No transactions in block")
	}

	for index, txBz := range txs {
		tx, err := CONFIG.EncodingConfig.TxConfig.TxDecoder()(txBz)
		if err != nil {
			panic(err)
		}

		ethTx, err := getEthTransactionRequest(tx)
		if err != nil {
			panic(err)
		}

		txHash := sha256.Sum256(txBz)
		txHashStr := hex.EncodeToString(txHash[:])

		log := fmt.Sprintf("%d: ethHash: %s shaHash: %s", index+1, ethTx.Hash().String(), txHashStr)
		fmt.Println(log)
	}
}

// displayExpectedOrder will display the expected order of the transactions
func displayExpectedOrder(bid *types.Transaction, transactions []*types.Transaction, prefix string) {
	fmt.Printf("\n%s expected order:\n", prefix)
	log := fmt.Sprintf("%d: ethHash: %s", 1, bid.Hash().String())
	fmt.Println(log)
	for index, tx := range transactions {
		log = fmt.Sprintf("%d: ethHash: %s", index+2, tx.Hash().String())
		fmt.Println(log)
	}
}

// getEthTransactionRequest returns the EthTransactionRequest message from a
// sdk transaction. If the transaction is not an EthTransactionRequest, it returns
// nil.
func getEthTransactionRequest(tx sdk.Tx) (*coretypes.Transaction, error) {
	msgEthTx := make([]*coretypes.Transaction, 0)
	for _, msg := range tx.GetMsgs() {
		if ethTxMsg, ok := msg.(*evmtypes.EthTransactionRequest); ok {
			msgEthTx = append(msgEthTx, ethTxMsg.AsTransaction())
		}
	}

	switch {
	case len(msgEthTx) == 0:
		return nil, nil
	case len(msgEthTx) == 1 && len(tx.GetMsgs()) == 1:
		return msgEthTx[0], nil
	default:
		return nil, fmt.Errorf("invalid transaction: %T", tx)
	}
}

// createEncodingConfig creates a new EncodingConfig for testing purposes.
func createEncodingConfig() EncodingConfig {
	cdc := codec.NewLegacyAmino()
	interfaceRegistry := codectypes.NewInterfaceRegistry()

	evmtypes.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
	buildertypes.RegisterInterfaces(interfaceRegistry)
	beracodec.RegisterInterfaces(interfaceRegistry)

	codec := codec.NewProtoCodec(interfaceRegistry)

	return EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             codec,
		TxConfig:          tx.NewTxConfig(codec, tx.DefaultSignModes),
		Amino:             cdc,
	}
}

// Creates a default configuration for testing the script
func DefaultConfig() ScriptConfig {
	privateKey, err := crypto.HexToECDSA("90c77c6e96b76b75e9f641184f4b9f93887b347e2826639e2a312a946b7dc939")
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	searcher := Account{
		Address:    crypto.PubkeyToAddress(*publicKeyECDSA),
		PrivateKey: privateKey,
	}

	return ScriptConfig{
		EthRPCURL:                   "http://localhost:1317/eth/rpc",
		CosmosRPCURL:                "localhost:9090",
		Searcher:                    searcher,
		ChainID:                     69420,
		AuctionSmartContractAddress: "0xDf6B07176A9B17cC4C9AFC257bD404732E7d09B7",
		NumAccounts:                 3,
		InitBalance:                 big.NewInt(10000000000),
		TestAccounts:                []Account{},
		EncodingConfig:              createEncodingConfig(),
	}
}
