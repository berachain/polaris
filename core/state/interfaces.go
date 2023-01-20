package state

import (
	"math/big"

	coretypes "github.com/berachain/stargazer/core/types"
	"github.com/berachain/stargazer/lib/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// `StargazerStateDB` defines an extension to the interface provided by go-ethereum to
// support additional state transition functionalities that are useful in a Cosmos SDK context.
type StargazerStateDB interface {
	BaseStateDB
	PrecompileStateDB

	Logs() []*coretypes.EthLog

	TransferBalance(from, to common.Address, amount *big.Int)
}

// `PrecompileStateDB` defines an extension to the interface provided by go-ethereum to
// support additional state transition functionalities that are useful in a Cosmos SDK context.
type PrecompileStateDB interface {
	// `AddLog` adds a log to the statedb.
	AddLog(*coretypes.EthLog)

	// `GetContext` returns the cosmos sdk context with the statedb multistore attached.
	GetContext() sdk.Context
}
