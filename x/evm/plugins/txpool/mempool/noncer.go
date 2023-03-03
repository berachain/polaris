package mempool

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

// `NonceRetriever` is an interface that allows for the
// TxPool plugin to retrieve the nonce of an account.
type NonceRetriever interface {
	GetNonce(addr common.Address) uint64
}

// `noncer` is a struct that implements the NonceRetriever interface
// and caches the nonce of an account.
type noncer struct {
	fallback NonceRetriever
	nonces   map[common.Address]uint64
	lock     sync.Mutex
}

// `newNoncer` returns a new noncer.
func newNoncer(nr NonceRetriever) *noncer {
	return &noncer{
		fallback: nr,
		nonces:   make(map[common.Address]uint64),
	}
}

// `GetNonce` returns the nonce of an account.
func (txn *noncer) get(addr common.Address) uint64 {
	txn.lock.Lock()
	defer txn.lock.Unlock()

	if _, ok := txn.nonces[addr]; !ok {
		if nonce := txn.fallback.GetNonce(addr); nonce != 0 {
			txn.nonces[addr] = nonce
		}
	}
	return txn.nonces[addr]
}

// `SetNonce` sets the nonce of an account.
func (txn *noncer) set(addr common.Address, nonce uint64) {
	txn.lock.Lock()
	defer txn.lock.Unlock()

	txn.nonces[addr] = nonce
}

// `SetIfLower` sets the nonce of an account if the nonce is lower than the
// current nonce.
func (txn *noncer) setIfLower(addr common.Address, txNonce uint64) {
	txn.lock.Lock()
	defer txn.lock.Unlock()

	if _, ok := txn.nonces[addr]; !ok {
		if sdbNonce := txn.fallback.GetNonce(addr); sdbNonce != 0 {
			txn.nonces[addr] = sdbNonce
		}
	}
	if txn.nonces[addr] <= txNonce {
		return
	}
	txn.nonces[addr] = txNonce
}
