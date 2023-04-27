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

package erc20

import (
	"context"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"

	cpbindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"
	pbindings "pkg.berachain.dev/polaris/contracts/bindings/polaris"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/precompile"
	erc20types "pkg.berachain.dev/polaris/cosmos/x/erc20/types"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/precompile/log"
	"pkg.berachain.dev/polaris/eth/accounts/abi"
	"pkg.berachain.dev/polaris/eth/common"
	ethprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/lib/utils"
)

// Contract is the precompile contract for the auth module.
type Contract struct {
	ethprecompile.BaseContract

	bk cosmlib.BankKeeper
	em ERC20Module

	polarisERC20ABI abi.ABI
}

// NewPrecompileContract returns a new instance of the auth module precompile contract.
func NewPrecompileContract(bk cosmlib.BankKeeper, em ERC20Module) ethprecompile.StatefulImpl {
	return &Contract{
		BaseContract: ethprecompile.NewBaseContract(
			cpbindings.ERC20ModuleMetaData.ABI,
			// cosmlib.AccAddressToEthAddress(
			// 	authtypes.NewModuleAddress(erc20types.ModuleName),
			// ),
			common.HexToAddress("0x696969"), // TODO: module addresses are broken
		),
		bk:              bk,
		em:              em,
		polarisERC20ABI: abi.MustUnmarshalJSON(pbindings.PolarisERC20MetaData.ABI),
	}
}

// CustomValueDecoders implements StatefulImpl.
func (c *Contract) CustomValueDecoders() ethprecompile.ValueDecoders {
	return ethprecompile.ValueDecoders{
		erc20types.AttributeKeyToken:     ConvertCommonHexAddress,
		erc20types.AttributeKeyDenom:     log.ReturnStringAsIs,
		erc20types.AttributeKeyOwner:     ConvertCommonHexAddress,
		erc20types.AttributeKeyRecipient: ConvertCommonHexAddress,
	}
}

// PrecompileMethods implements StatefulImpl.
func (c *Contract) PrecompileMethods() ethprecompile.Methods {
	return ethprecompile.Methods{
		{
			AbiSig:  "coinDenomForERC20Address(address)",
			Execute: c.CoinDenomForERC20AddressAddrInput,
		},
		{
			AbiSig:  "coinDenomForERC20Address(string)",
			Execute: c.CoinDenomForERC20AddressStringInput,
		},
		{
			AbiSig:  "erc20AddressForCoinDenom(string)",
			Execute: c.ERC20AddressForCoinDenom,
		},
		{
			AbiSig:  "convertCoinToERC20(string,uint256)",
			Execute: c.ConvertCoinToERC20,
		},
		{
			AbiSig:  "convertCoinToERC20From(string,address,address,uint256)",
			Execute: c.ConvertCoinToERC20FromAddrInput,
		},
		{
			AbiSig:  "convertCoinToERC20From(string,string,string,uint256)",
			Execute: c.ConvertCoinToERC20FromStringInput,
		},
		{
			AbiSig:  "convertCoinToERC20To(string,address,uint256)",
			Execute: c.ConvertCoinToERC20ToAddrInput,
		},
		{
			AbiSig:  "convertCoinToERC20To(string,string,uint256)",
			Execute: c.ConvertCoinToERC20ToStringInput,
		},
		{
			AbiSig:  "convertERC20ToCoin(address,uint256)",
			Execute: c.ConvertERC20ToCoin,
		},
		{
			AbiSig:  "convertERC20ToCoinFrom(address,address,address,uint256)",
			Execute: c.ConvertERC20ToCoinFromAddrInput,
		},
		{
			AbiSig:  "convertERC20ToCoinFrom(address,string,string,uint256)",
			Execute: c.ConvertERC20ToCoinFromStringInput,
		},
		{
			AbiSig:  "convertERC20ToCoinTo(address,address,uint256)",
			Execute: c.ConvertERC20ToCoinToAddrInput,
		},
		{
			AbiSig:  "convertERC20ToCoinTo(address,string,uint256)",
			Execute: c.ConvertERC20ToCoinToStringInput,
		},
	}
}

// CoinDenomForERC20AddressAddrInput returns the SDK coin denomination for the given ERC20 address.
func (c *Contract) CoinDenomForERC20AddressAddrInput(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	addr, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}

	resp, err := c.em.CoinDenomForERC20Address(
		ctx,
		&erc20types.CoinDenomForERC20AddressRequest{
			Token: cosmlib.AddressToAccAddress(addr).String(),
		},
	)
	if err != nil {
		return nil, err
	}

	return []any{resp.Denom}, nil
}

// CoinDenomForERC20AddressStringInput returns the SDK coin denomination for the given ERC20 address.
func (c *Contract) CoinDenomForERC20AddressStringInput(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	addr, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}

	resp, err := c.em.CoinDenomForERC20Address(
		ctx,
		&erc20types.CoinDenomForERC20AddressRequest{
			Token: addr,
		},
	)
	if err != nil {
		return nil, err
	}

	return []any{resp.Denom}, nil
}

// ERC20AddressForCoinDenom returns the ERC20 address for the given SDK coin denomination.
func (c *Contract) ERC20AddressForCoinDenom(
	ctx context.Context,
	_ ethprecompile.EVM,
	_ common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	denom, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}

	resp, err := c.em.ERC20AddressForCoinDenom(
		ctx,
		&erc20types.ERC20AddressForCoinDenomRequest{
			Denom: denom,
		},
	)
	if err != nil {
		return nil, err
	}

	tokenAddr := common.Address{}
	if resp.Token != "" {
		var tokenAccAddr sdk.AccAddress
		if tokenAccAddr, err = sdk.AccAddressFromBech32(resp.Token); err != nil {
			return nil, err
		}
		tokenAddr = cosmlib.AccAddressToEthAddress(tokenAccAddr)
	}

	return []any{tokenAddr}, nil
}

// ConvertCoinToERC20 converts SDK coins to ERC20 tokens for msg.sender.
func (c *Contract) ConvertCoinToERC20(
	ctx context.Context,
	evm ethprecompile.EVM,
	caller common.Address,
	value *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	denom, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	amount, ok := utils.GetAs[*big.Int](args[1])
	if !ok {
		return nil, precompile.ErrInvalidBigInt
	}

	err := c.convertCoinToERC20(ctx, evm, value, denom, caller, caller, amount)
	return []any{err == nil}, err
}

// ConvertCoinToERC20FromAddrInput converts SDK coins to ERC20 tokens from owner to recipient.
func (c *Contract) ConvertCoinToERC20FromAddrInput(
	ctx context.Context,
	evm ethprecompile.EVM,
	_ common.Address,
	value *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	denom, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	owner, ok := utils.GetAs[common.Address](args[1])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	recipient, ok := utils.GetAs[common.Address](args[2])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	amount, ok := utils.GetAs[*big.Int](args[3])
	if !ok {
		return nil, precompile.ErrInvalidBigInt
	}

	err := c.convertCoinToERC20(ctx, evm, value, denom, owner, recipient, amount)
	return []any{err == nil}, err
}

// ConvertCoinToERC20FromStringInput converts SDK coins to ERC20 tokens from owner to recipient.
func (c *Contract) ConvertCoinToERC20FromStringInput(
	ctx context.Context,
	evm ethprecompile.EVM,
	_ common.Address,
	value *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	denom, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	ownerBech32, ok := utils.GetAs[string](args[1])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	recipientBech32, ok := utils.GetAs[string](args[2])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	amount, ok := utils.GetAs[*big.Int](args[3])
	if !ok {
		return nil, precompile.ErrInvalidBigInt
	}

	owner, err := sdk.AccAddressFromBech32(ownerBech32)
	if err != nil {
		return nil, err
	}
	recipient, err := sdk.AccAddressFromBech32(recipientBech32)
	if err != nil {
		return nil, err
	}

	err = c.convertCoinToERC20(
		ctx, evm, value, denom,
		cosmlib.AccAddressToEthAddress(owner), cosmlib.AccAddressToEthAddress(recipient),
		amount,
	)
	return []any{err == nil}, err
}

// ConvertCoinToERC20ToAddrInput converts SDK coins to ERC20 tokens from msg.sender to recipient.
func (c *Contract) ConvertCoinToERC20ToAddrInput(
	ctx context.Context,
	evm ethprecompile.EVM,
	caller common.Address,
	value *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	denom, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	recipient, ok := utils.GetAs[common.Address](args[1])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	amount, ok := utils.GetAs[*big.Int](args[2])
	if !ok {
		return nil, precompile.ErrInvalidBigInt
	}

	err := c.convertCoinToERC20(ctx, evm, value, denom, caller, recipient, amount)
	return []any{err == nil}, err
}

// ConvertCoinToERC20ToStringInput converts SDK coins to ERC20 tokens from msg.sender to recipient.
func (c *Contract) ConvertCoinToERC20ToStringInput(
	ctx context.Context,
	evm ethprecompile.EVM,
	caller common.Address,
	value *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	denom, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	recipientBech32, ok := utils.GetAs[string](args[1])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	amount, ok := utils.GetAs[*big.Int](args[2])
	if !ok {
		return nil, precompile.ErrInvalidBigInt
	}

	recipient, err := sdk.AccAddressFromBech32(recipientBech32)
	if err != nil {
		return nil, err
	}

	err = c.convertCoinToERC20(
		ctx, evm, value, denom,
		caller, cosmlib.AccAddressToEthAddress(recipient),
		amount,
	)
	return []any{err == nil}, err
}

// ConvertERC20ToCoin converts ERC20 tokens to SDK coins for msg.sender.
func (c *Contract) ConvertERC20ToCoin(
	ctx context.Context,
	evm ethprecompile.EVM,
	caller common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	token, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	amount, ok := utils.GetAs[*big.Int](args[1])
	if !ok {
		return nil, precompile.ErrInvalidBigInt
	}

	err := c.convertERC20ToCoin(ctx, caller, evm, token, caller, caller, amount)
	return []any{err == nil}, err
}

// ConvertERC20ToCoinFromAddrInput converts ERC20 tokens to SDK coins from owner to recipient.
func (c *Contract) ConvertERC20ToCoinFromAddrInput(
	ctx context.Context,
	evm ethprecompile.EVM,
	caller common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	token, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	owner, ok := utils.GetAs[common.Address](args[1])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	recipient, ok := utils.GetAs[common.Address](args[2])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	amount, ok := utils.GetAs[*big.Int](args[3])
	if !ok {
		return nil, precompile.ErrInvalidBigInt
	}

	err := c.convertERC20ToCoin(ctx, caller, evm, token, owner, recipient, amount)
	return []any{err == nil}, err
}

// ConvertERC20ToCoinFromStringInput converts ERC20 tokens to SDK coins from owner to recipient.
func (c *Contract) ConvertERC20ToCoinFromStringInput(
	ctx context.Context,
	evm ethprecompile.EVM,
	caller common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	token, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	ownerBech32, ok := utils.GetAs[string](args[1])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	recipientBech32, ok := utils.GetAs[string](args[2])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	amount, ok := utils.GetAs[*big.Int](args[3])
	if !ok {
		return nil, precompile.ErrInvalidBigInt
	}

	owner, err := sdk.AccAddressFromBech32(ownerBech32)
	if err != nil {
		return nil, err
	}
	recipient, err := sdk.AccAddressFromBech32(recipientBech32)
	if err != nil {
		return nil, err
	}

	err = c.convertERC20ToCoin(
		ctx, caller, evm, token,
		cosmlib.AccAddressToEthAddress(owner), cosmlib.AccAddressToEthAddress(recipient),
		amount,
	)
	return []any{err == nil}, err
}

// ConvertERC20ToCoinToAddrInput converts ERC20 tokens to SDK coins from msg.sender to recipient.
func (c *Contract) ConvertERC20ToCoinToAddrInput(
	ctx context.Context,
	evm ethprecompile.EVM,
	caller common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	token, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	recipient, ok := utils.GetAs[common.Address](args[1])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	amount, ok := utils.GetAs[*big.Int](args[2])
	if !ok {
		return nil, precompile.ErrInvalidBigInt
	}

	err := c.convertERC20ToCoin(ctx, caller, evm, token, caller, recipient, amount)
	return []any{err == nil}, err
}

// ConvertERC20ToCoinToStringInput converts ERC20 tokens to SDK coins from msg.sender to recipient.
func (c *Contract) ConvertERC20ToCoinToStringInput(
	ctx context.Context,
	evm ethprecompile.EVM,
	caller common.Address,
	_ *big.Int,
	_ bool,
	args ...any,
) ([]any, error) {
	token, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	recipientBech32, ok := utils.GetAs[string](args[1])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	amount, ok := utils.GetAs[*big.Int](args[2])
	if !ok {
		return nil, precompile.ErrInvalidBigInt
	}

	recipient, err := sdk.AccAddressFromBech32(recipientBech32)
	if err != nil {
		return nil, err
	}

	err = c.convertERC20ToCoin(
		ctx, caller, evm, token,
		caller, cosmlib.AccAddressToEthAddress(recipient),
		amount,
	)
	return []any{err == nil}, err
}
