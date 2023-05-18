package core

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/rlp"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/common/hexutil"
	"pkg.berachain.dev/polaris/eth/params"
)

var DefaultGenesis = &Genesis{
	Config:     params.DefaultChainConfig,
	Nonce:      69,
	ExtraData:  hexutil.MustDecode("0x11bbe8db4e347b4e8c937c1c8370e4b5ed33adb3db69cbdb7a38e1e50b1b82fa"),
	GasLimit:   30_000_000,
	Difficulty: big.NewInt(69),
	Alloc:      GenesisAlloc{},
	// Alloc:      decodePrealloc("mainnetAllocData"),
	// For alloc, in the startup / initGenesis, we should allow the host chain to "fill in the data"
	// i.e in Cosmos, we let the AccountKeeper/EVMKeeper/BankKeeper fill in the Bank Data into the
	// genesis and then verify the equivalency later. This is to create an invariant that the bank
	// balances from the bank keeper and the token balances in the EVM are equiavlent at genesis.
}

func decodePrealloc(data string) GenesisAlloc {
	var p []struct{ Addr, Balance *big.Int }
	if err := rlp.NewStream(strings.NewReader(data), 0).Decode(&p); err != nil {
		panic(err)
	}
	ga := make(GenesisAlloc, len(p))
	for _, account := range p {
		ga[common.BigToAddress(account.Addr)] = GenesisAccount{Balance: account.Balance}
	}
	return ga
}
