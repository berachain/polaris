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

	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/auth"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/precompile"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/lib/utils"
)

// requiredGas is the amount required to convert between eth addresses and bech32 cosmos addresses.
const requiredGas = 1000

// Contract is the precompile contract for the auth(z) module.
type Contract struct {
	ethprecompile.BaseContract

	authQueryServer authtypes.QueryServer

	msgServer   authz.MsgServer
	queryServer authz.QueryServer
}

// NewPrecompileContract returns a new instance of the auth(z) module precompile contract. Uses the
// auth module's account address as the contract address.
func NewPrecompileContract(m authz.MsgServer, q authz.QueryServer) *Contract {
	return &Contract{
		BaseContract: ethprecompile.NewBaseContract(
			generated.AuthModuleMetaData.ABI,
			cosmlib.AccAddressToEthAddress(
				authtypes.NewModuleAddress(authtypes.ModuleName),
			),
		),
		msgServer:   m,
		queryServer: q,
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
			AbiSig:  "setSendAllowance(address,address,(uint256,string)[],uint256)",
			Execute: c.SetSendAllowance,
		},
		{
			AbiSig:  "getSendAllowance(address,address,string)",
			Execute: c.GetSendAllowance,
		},
		{
			AbiSig:  "getAccountInfo(address)",
			Execute: c.GetAccountInfoAddrInput,
		},
		{
			AbiSig:  "getAccountInfo(string)",
			Execute: c.GetAccountInfoStringInput,
		},
	}
}

// ConvertHexToBech32 converts a common.Address to a bech32 string.
func (c *Contract) ConvertHexToBech32(
	_ context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
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
	_ context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
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

// SetSendAllowance sends a send authorization message to the authz module.
func (c *Contract) SetSendAllowance(
	ctx context.Context,
	evm ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	owner, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	spender, ok := utils.GetAs[common.Address](args[1])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	amount, err := cosmlib.ExtractCoinsFromInput(args[2])
	if err != nil {
		return nil, err
	}
	expiration, ok := utils.GetAs[*big.Int](args[3])
	if !ok {
		return nil, precompile.ErrInvalidBigInt
	}

	return c.setSendAllowanceHelper(
		ctx,
		time.Unix(int64(evm.GetContext().Time), 0),
		cosmlib.AddressToAccAddress(owner),
		cosmlib.AddressToAccAddress(spender),
		amount,
		expiration,
	)
}

// GetSendAllowance returns the amount of tokens that the spender is allowd to spend.
func (c *Contract) GetSendAllowance(
	ctx context.Context,
	evm ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	owner, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	spender, ok := utils.GetAs[common.Address](args[1])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	denom, ok := utils.GetAs[string](args[2])
	if !ok {
		return nil, precompile.ErrInvalidString
	}

	return c.getSendAllownaceHelper(
		ctx,
		time.Unix(int64(evm.GetContext().Time), 0),
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

// GetAccountInfoAddrInput implements `getAccountInfo(address)`.
func (c *Contract) GetAccountInfoAddrInput(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	acc, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	return c.accountInfoHelper(ctx, cosmlib.Bech32FromEthAddress(acc))
}

// GetAccountInfoStringInput implements `getAccountInfo(string)`.
func (c *Contract) GetAccountInfoStringInput(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	acc, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	return c.accountInfoHelper(ctx, acc)
}
