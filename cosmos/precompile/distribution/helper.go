package distribution

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
)

// `setWithdrawAddressHelper` is a helper function for the `SetWithdrawAddress` method.
func (c *Contract) setWithdrawAddressHelper(ctx context.Context, delegator, withdrawer sdk.AccAddress) ([]any, error) {
	_, err := c.msgServer.SetWithdrawAddress(ctx, &distributiontypes.MsgSetWithdrawAddress{
		DelegatorAddress: delegator.String(),
		WithdrawAddress:  withdrawer.String(),
	})
	if err != nil {
		return nil, err
	}

	return []any{}, nil
}
