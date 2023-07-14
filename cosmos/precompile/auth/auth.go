//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package auth

import (
	"context"
	"math/big"
	"time"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/auth"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/core/vm"
)

// Contract is the precompile contract for the auth(z) module.
type Contract struct {
	ethprecompile.BaseContract

	authQueryServer authtypes.QueryServer
	msgServer       authz.MsgServer
	queryServer     authz.QueryServer
}

// NewPrecompileContract returns a new instance of the auth(z) module precompile contract. Uses the
// auth module's account address as the contract address.
func NewPrecompileContract(
	authQueryServer authtypes.QueryServer,
	authzMsgServer authz.MsgServer,
	authzQueryServer authz.QueryServer,
) *Contract {
	return &Contract{
		BaseContract: ethprecompile.NewBaseContract(
			generated.AuthModuleMetaData.ABI,
			cosmlib.AccAddressToEthAddress(authtypes.NewModuleAddress(authtypes.ModuleName)),
		),
		authQueryServer: authQueryServer,
		msgServer:       authzMsgServer,
		queryServer:     authzQueryServer,
	}
}

// GetAccountInfoStringInput implements `getAccountInfo(address)`.
func (c *Contract) GetAccountInfo(
	ctx context.Context,
	account common.Address,
) (generated.IAuthModuleBaseAccount, error) {
	return c.accountInfoHelper(ctx, cosmlib.Bech32FromEthAddress(account))
}

// SetSendAllowance sends a send authorization message to the authz module.
func (c *Contract) SetSendAllowance(
	ctx context.Context,
	spender common.Address,
	amount any,
	expiration *big.Int,
) (bool, error) {
	amt, err := cosmlib.ExtractCoinsFromInput(amount)
	if err != nil {
		return false, err
	}

	polarCtx := vm.UnwrapPolarContext(ctx)

	return c.setSendAllowanceHelper(
		ctx,
		time.Unix(int64(polarCtx.Block().Time), 0),
		cosmlib.AddressToAccAddress(polarCtx.MsgSender()),
		cosmlib.AddressToAccAddress(spender),
		amt,
		expiration,
	)
}

// GetSendAllowance returns the amount of tokens that the spender is allowd to spend.
func (c *Contract) GetSendAllowance(
	ctx context.Context,
	owner common.Address,
	spender common.Address,
	denom string,
) (*big.Int, error) {
	return c.getSendAllownaceHelper(
		ctx,
		time.Unix(int64(vm.UnwrapPolarContext(ctx).Evm().GetContext().Time), 0),
		cosmlib.AddressToAccAddress(owner),
		cosmlib.AddressToAccAddress(spender),
		denom,
	)
}

// getHighestAllowance returns the highest allowance for a given coin denom.
func getHighestAllowance(sendAuths []*banktypes.SendAuthorization, coinDenom string) *big.Int {
	// Init the max to 0.
	var max = big.NewInt(0)
	// Loop through the send authorizations and find the highest allowance.
	for _, sendAuth := range sendAuths {
		// Get the spendable limit for the coin denom that was specified.
		amount := sendAuth.SpendLimit.AmountOf(coinDenom)
		// If not set, the current is the max, if set, compare the current with the max.
		if max == nil || amount.BigInt().Cmp(max) > 0 {
			max = amount.BigInt()
		}
	}
	return max
}
