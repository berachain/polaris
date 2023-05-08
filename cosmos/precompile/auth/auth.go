// SPDX-License-Identifier: BUSL-1.1
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

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/precompile"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/lib/utils"
)

const requiredGas = 1000

// Contract is the precompile contract for the auth module.
type Contract struct {
	ethprecompile.BaseContract

	msgServer authz.MsgServer
}

// NewPrecompileContract returns a new instance of the auth module precompile contract.
func NewPrecompileContract(m authz.MsgServer) *Contract {
	return &Contract{
		BaseContract: ethprecompile.NewBaseContract(
			generated.AuthModuleMetaData.ABI,
			cosmlib.AccAddressToEthAddress(
				authtypes.NewModuleAddress(authtypes.ModuleName),
			),
		),
		msgServer: m,
	}
}

// PrecompileMethods implements StatefulImpl.
func (c *Contract) PrecompileMethods() ethprecompile.Methods {
	return ethprecompile.Methods{
		{
			AbiSig:      "convertHexToBech32(address)",
			Execute:     c.ConvertHexToBech32,
			RequiredGas: requiredGas,
		},
		{
			AbiSig:      "convertBech32ToHexAddress(string)",
			Execute:     c.ConvertBech32ToHexAddress,
			RequiredGas: requiredGas,
		},
		{
			AbiSig:  "sendGrant(address,address,(uint256,string)[],uint256)",
			Execute: c.SendGrant,
		},
	}
}

// ConvertHexToBech32 converts a common.Address to a bech32 string.
func (c *Contract) ConvertHexToBech32(
	ctx context.Context,
	_ ethprecompile.EVM,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	hexAddr, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}

	// try val address first
	valAddr, err := sdk.ValAddressFromHex(hexAddr.String())
	if err == nil {
		return []any{valAddr.String()}, nil
	}

	// try account address
	accAddr, err := sdk.AccAddressFromHexUnsafe(hexAddr.String())
	if err == nil {
		return []any{accAddr.String()}, nil
	}

	return nil, precompile.ErrInvalidHexAddress
}

// ConvertBech32ToHexAddress converts a bech32 string to a common.Address.
func (c *Contract) ConvertBech32ToHexAddress(
	ctx context.Context,
	_ ethprecompile.EVM,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	bech32Addr, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}

	// try account address first
	accAddr, err := sdk.AccAddressFromBech32(bech32Addr)
	if err == nil {
		return []any{cosmlib.AccAddressToEthAddress(accAddr)}, nil
	}

	// try validator address
	valAddr, err := sdk.ValAddressFromBech32(bech32Addr)
	if err == nil {
		return []any{cosmlib.ValAddressToEthAddress(valAddr)}, nil
	}

	return nil, precompile.ErrInvalidBech32Address
}

// SendGrant sends a send grant message to the authz module.
func (c *Contract) SendGrant(
	ctx context.Context,
	evm ethprecompile.EVM,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	granter, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}

	grantee, ok := utils.GetAs[common.Address](args[1])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}

	limit, err := extractCoinsFromInput(args[2])
	if err != nil {
		return nil, err
	}

	expiration, ok := utils.GetAs[*big.Int](args[3])
	if !ok {
		return nil, precompile.ErrInvalidBigInt
	}

	// Get the block time from the EVM.
	blockTime := time.Unix(int64(evm.GetContext().Time), 0)

	return c.sendGrant(
		ctx, blockTime, cosmlib.AddressToAccAddress(granter), cosmlib.AddressToAccAddress(grantee), limit, expiration)
}

// extractCoinsFromInput converts coins from input (of type any) into sdk.Coins.
func extractCoinsFromInput(coins any) (sdk.Coins, error) {
	// note: we have to use unnamed struct here, otherwise the compiler cannot cast
	// the any type input into IBankModuleCoin.
	amounts, ok := utils.GetAs[[]struct {
		Amount *big.Int `json:"amount"`
		Denom  string   `json:"denom"`
	}](coins)
	if !ok {
		return nil, precompile.ErrInvalidCoin
	}

	sdkCoins := sdk.NewCoins()
	for _, evmCoin := range amounts {
		sdkCoins = sdkCoins.Add(
			sdk.Coin{
				Amount: sdk.NewIntFromBigInt(evmCoin.Amount),
				Denom:  evmCoin.Denom,
			},
		)
	}
	return sdkCoins, nil
}

// sendGrant is the helper method to call the grant method on the msgServer, with a send authorization.
func (c *Contract) sendGrant(
	ctx context.Context,
	blocktime time.Time,
	granter, grantee sdk.AccAddress,
	limit sdk.Coins,
	expiration *big.Int,
) ([]any, error) {
	var (
		grant authz.Grant
		err   error
	)

	// Create the send authorization.
	sendAuth := banktypes.NewSendAuthorization(limit, []sdk.AccAddress{grantee})

	// If the expiration is 0, then the grant is valid forever, and can be nil.
	if expiration == big.NewInt(0) {
		grant, err = authz.NewGrant(blocktime, sendAuth, nil)
	} else {
		expirationTime := time.Unix(expiration.Int64(), 0)
		grant, err = authz.NewGrant(blocktime, sendAuth, &expirationTime)
	}

	// Assert that the grant is valid.
	if err != nil {
		return nil, err
	}

	_, err = c.msgServer.Grant(ctx, &authz.MsgGrant{
		Granter: granter.String(),
		Grantee: grantee.String(),
		Grant:   grant,
	})

	return []any{err == nil}, err
}
