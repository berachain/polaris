package mock

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

//go:generate moq -out ./statedb.mock.go -pkg mock ../ StargazerStateDB

// `NewEmptyStateDB` creates a new `StateDBMock` instance.
func NewEmptyStateDB() *StargazerStateDBMock {
	ssdb := &StargazerStateDBMock{}

	ssdb.SetCodeFunc = func(addr common.Address, code []byte) {
	}

	ssdb.SetStateFunc = func(addr common.Address, key, value common.Hash) {
	}

	ssdb.SetNonceFunc = func(addr common.Address, nonce uint64) {
	}

	ssdb.GetNonceFunc = func(addr common.Address) uint64 {
		return 0
	}

	ssdb.GetBalanceFunc = func(addr common.Address) *big.Int {
		return big.NewInt(0)
	}

	ssdb.GetCodeSizeFunc = func(addr common.Address) int {
		return 0
	}

	ssdb.GetCodeFunc = func(addr common.Address) []byte {
		return nil
	}

	ssdb.GetCodeHashFunc = func(addr common.Address) common.Hash {
		return common.Hash{}
	}

	ssdb.GetRefundFunc = func() uint64 {
		return 0
	}

	return ssdb
}
