package core

import (
	"github.com/ethereum/go-ethereum/core"
)

type Message = core.Message

var (
	NewEVMTxContext                 = core.NewEVMTxContext
	ErrIntrinsicGas                 = core.ErrIntrinsicGas
	EthIntrinsicGas                 = core.IntrinsicGas
	ErrInsufficientFundsForTransfer = core.ErrInsufficientFundsForTransfer
	ErrInsufficientFunds            = core.ErrInsufficientFunds
	ErrGasUintOverflow              = core.ErrGasUintOverflow
)
