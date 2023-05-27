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
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	cbindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos"
	cpbindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile/erc20"
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

	bk bankkeeper.Keeper
	em ERC20Module

	polarisERC20ABI abi.ABI
	polarisERC20Bin string
}

// NewPrecompileContract returns a new instance of the auth module precompile contract.
func NewPrecompileContract(bk bankkeeper.Keeper, em ERC20Module) ethprecompile.StatefulImpl {
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
		polarisERC20ABI: abi.MustUnmarshalJSON(cbindings.PolarisERC20MetaData.ABI),
		polarisERC20Bin: cbindings.PolarisERC20MetaData.Bin,
	}
}

// CustomValueDecoders implements StatefulImpl.
func (c *Contract) CustomValueDecoders() ethprecompile.ValueDecoders {
	return ethprecompile.ValueDecoders{
		erc20types.AttributeKeyToken:     TransferCommonHexAddress,
		erc20types.AttributeKeyDenom:     log.ReturnStringAsIs,
		erc20types.AttributeKeyOwner:     TransferCommonHexAddress,
		erc20types.AttributeKeyRecipient: TransferCommonHexAddress,
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
			AbiSig:  "transferCoinToERC20(string,uint256)",
			Execute: c.TransferCoinToERC20,
		},
		{
			AbiSig:  "transferCoinToERC20From(string,address,address,uint256)",
			Execute: c.TransferCoinToERC20FromAddrInput,
		},
		{
			AbiSig:  "transferCoinToERC20From(string,string,string,uint256)",
			Execute: c.TransferCoinToERC20FromStringInput,
		},
		{
			AbiSig:  "transferCoinToERC20To(string,address,uint256)",
			Execute: c.TransferCoinToERC20ToAddrInput,
		},
		{
			AbiSig:  "transferCoinToERC20To(string,string,uint256)",
			Execute: c.TransferCoinToERC20ToStringInput,
		},
		{
			AbiSig:  "transferERC20ToCoin(address,uint256)",
			Execute: c.TransferERC20ToCoin,
		},
		{
			AbiSig:  "transferERC20ToCoinFrom(address,address,address,uint256)",
			Execute: c.TransferERC20ToCoinFromAddrInput,
		},
		{
			AbiSig:  "transferERC20ToCoinFrom(address,string,string,uint256)",
			Execute: c.TransferERC20ToCoinFromStringInput,
		},
		{
			AbiSig:  "transferERC20ToCoinTo(address,address,uint256)",
			Execute: c.TransferERC20ToCoinToAddrInput,
		},
		{
			AbiSig:  "transferERC20ToCoinTo(address,string,uint256)",
			Execute: c.TransferERC20ToCoinToStringInput,
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
			Token: cosmlib.Bech32FromEthAddress(addr),
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

// TransferCoinToERC20 transfers SDK coins to ERC20 tokens for msg.sender.
func (c *Contract) TransferCoinToERC20(
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

	err := c.transferCoinToERC20(ctx, evm, value, denom, caller, caller, amount)
	return []any{err == nil}, err
}

// TransferCoinToERC20FromAddrInput transfers SDK coins to ERC20 tokens from owner to recipient.
func (c *Contract) TransferCoinToERC20FromAddrInput(
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

	err := c.transferCoinToERC20(ctx, evm, value, denom, owner, recipient, amount)
	return []any{err == nil}, err
}

// TransferCoinToERC20FromStringInput transfers SDK coins to ERC20 tokens from owner to recipient.
func (c *Contract) TransferCoinToERC20FromStringInput(
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

	err = c.transferCoinToERC20(
		ctx, evm, value, denom,
		cosmlib.AccAddressToEthAddress(owner), cosmlib.AccAddressToEthAddress(recipient),
		amount,
	)
	return []any{err == nil}, err
}

// TransferCoinToERC20ToAddrInput transfers SDK coins to ERC20 tokens from msg.sender to recipient.
func (c *Contract) TransferCoinToERC20ToAddrInput(
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

	err := c.transferCoinToERC20(ctx, evm, value, denom, caller, recipient, amount)
	return []any{err == nil}, err
}

// TransferCoinToERC20ToStringInput transfers SDK coins to ERC20 tokens from msg.sender to recipient.
func (c *Contract) TransferCoinToERC20ToStringInput(
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

	err = c.transferCoinToERC20(
		ctx, evm, value, denom,
		caller, cosmlib.AccAddressToEthAddress(recipient),
		amount,
	)
	return []any{err == nil}, err
}

// TransferERC20ToCoin transfers ERC20 tokens to SDK coins for msg.sender.
func (c *Contract) TransferERC20ToCoin(
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

	err := c.transferERC20ToCoin(ctx, caller, evm, token, caller, caller, amount)
	return []any{err == nil}, err
}

// TransferERC20ToCoinFromAddrInput transfers ERC20 tokens to SDK coins from owner to recipient.
func (c *Contract) TransferERC20ToCoinFromAddrInput(
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

	err := c.transferERC20ToCoin(ctx, caller, evm, token, owner, recipient, amount)
	return []any{err == nil}, err
}

// TransferERC20ToCoinFromStringInput transfers ERC20 tokens to SDK coins from owner to recipient.
func (c *Contract) TransferERC20ToCoinFromStringInput(
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

	err = c.transferERC20ToCoin(
		ctx, caller, evm, token,
		cosmlib.AccAddressToEthAddress(owner), cosmlib.AccAddressToEthAddress(recipient),
		amount,
	)
	return []any{err == nil}, err
}

// TransferERC20ToCoinToAddrInput transfers ERC20 tokens to SDK coins from msg.sender to recipient.
func (c *Contract) TransferERC20ToCoinToAddrInput(
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

	err := c.transferERC20ToCoin(ctx, caller, evm, token, caller, recipient, amount)
	return []any{err == nil}, err
}

// TransferERC20ToCoinToStringInput transfers ERC20 tokens to SDK coins from msg.sender to recipient.
func (c *Contract) TransferERC20ToCoinToStringInput(
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

	err = c.transferERC20ToCoin(
		ctx, caller, evm, token,
		caller, cosmlib.AccAddressToEthAddress(recipient),
		amount,
	)
	return []any{err == nil}, err
}

// ==============================================================================
// Event Attribute Value Decoders
// ==============================================================================

// TransferCommonHexAddress is a value decoder.
var _ ethprecompile.ValueDecoder = TransferCommonHexAddress

// TransferCommonHexAddress transfers a common hex address attribute to a common.Address and returns
// it as type any.
func TransferCommonHexAddress(attributeValue string) (any, error) {
	return common.HexToAddress(attributeValue), nil
}
