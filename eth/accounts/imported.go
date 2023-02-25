package accounts

import (
	"github.com/ethereum/go-ethereum/accounts"
)

type (
	DerivationPath = accounts.DerivationPath
	HDPathIterator = func() DerivationPath
)

var (
	DefaultBaseDerivationPath = accounts.DefaultBaseDerivationPath
	DefaultIterator           = accounts.DefaultIterator
	LedgerLiveIterator        = accounts.LedgerLiveIterator
	ParseDerivationPath       = accounts.ParseDerivationPath
)
