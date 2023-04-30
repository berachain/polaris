package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/lib/utils"
)

// IgnoreDecorator is an AnteDecorator that wraps an existing AnteDecorator. It allows
// EthTransactions to skip said Decorator by checking to see if the transaction
// contains a message of the given type.
type IgnoreDecorator[D sdk.AnteDecorator, M sdk.Msg] struct {
	decorator D
}

// NewIgnoreDecorator returns a new IgnoreDecorator[D, M] instance.
func NewIgnoreDecorator[D sdk.AnteDecorator, M sdk.Msg](decorator D) *IgnoreDecorator[D, M] {
	return &IgnoreDecorator[D, M]{
		decorator: decorator,
	}
}

// AnteHandle implements the sdk.AnteDecorator interface, it is handle the
// type check for the message type.
func (sd IgnoreDecorator[D, M]) AnteHandle(
	ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler,
) (sdk.Context, error) {
	for _, msg := range tx.GetMsgs() {
		if _, ok := utils.GetAs[M](msg); ok {
			return next(ctx, tx, simulate)
		}
	}

	return sd.decorator.AnteHandle(ctx, tx, simulate, next)
}
