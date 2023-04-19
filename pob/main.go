package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	bindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"
)

func main() {
	client, err := ethclient.Dial("http://localhost:1317/eth/rpc")
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.HexToECDSA("90c77c6e96b76b75e9f641184f4b9f93887b347e2826639e2a312a946b7dc939")
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(69420))
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0xDf6B07176A9B17cC4C9AFC257bD404732E7d09B7")
	instance, err := bindings.NewBuilderModule(address, client)
	if err != nil {
		panic(err)
	}

	tx, err := instance.AuctionBid(auth, bindings.IBuilderModuleCoin{Denom: "abera", Amount: 100}, [][]byte{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
}
