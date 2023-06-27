package precompile

import (
	"context"
	"math/big"

	"pkg.berachain.dev/polaris/eth/common"
)

type (
	// PolarContext defines the fields that each Precompile implementation should have access to.
	// It contains a context,
	// an EVM to execute,
	// an Address for msg.sender,
	// and a *big.Int for msg.value
	PolarContext struct {
		Ctx       context.Context
		Evm       EVM
		MsgSender common.Address
		MsgValue  *big.Int
	}
)
