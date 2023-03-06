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

package address

import (
	"context"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/eth/accounts/abi"
	"pkg.berachain.dev/polaris/eth/common"
	coreprecompile "pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/eth/params"
	"pkg.berachain.dev/polaris/lib/utils"
	"pkg.berachain.dev/polaris/precompile"
	"pkg.berachain.dev/polaris/precompile/contracts/solidity/generated"
	evmutils "pkg.berachain.dev/polaris/x/evm/utils"
)

// `Contract` is the precompile contract for the address util.
type Contract struct {
	contractAbi abi.ABI
}

// `NewPrecompileContract` creates a new contract instance that implements the `precompile.StatefulImpl` interface.
func NewPrecompileContract() coreprecompile.StatefulImpl {
	var contractAbi abi.ABI
	if err := contractAbi.UnmarshalJSON([]byte(generated.AddressMetaData.ABI)); err != nil {
		panic(err)
	}
	return &Contract{
		contractAbi: contractAbi,
	}
}

// `RegisterKey` implements `precompile.StatelessImpl`.
func (c *Contract) RegistryKey() common.Address {
	// Contract Address: 0x19
	return common.BytesToAddress([]byte{19})
}

// `ABIMethods` implements StatefulImpl.
func (c *Contract) ABIMethods() map[string]abi.Method {
	return c.contractAbi.Methods
}

// `PrecompileMethods` implements StatefulImpl.
func (c *Contract) PrecompileMethods() coreprecompile.Methods {
	return coreprecompile.Methods{
		{
			AbiSig:      "convertHexToBech32(address)",
			Execute:     c.ConvertHexToBech32,
			RequiredGas: params.IdentityBaseGas,
		},
		{
			AbiSig:      "convertBech32ToHexAddress(string)",
			Execute:     c.ConvertBech32ToHexAddress,
			RequiredGas: params.IdentityBaseGas,
		},
	}
}

// `ABIEvents` implements StatefulImpl.
func (c *Contract) ABIEvents() map[string]abi.Event {
	return c.contractAbi.Events
}

// `CustomValueDecoders` implements StatefulImpl.
func (c *Contract) CustomValueDecoders() coreprecompile.ValueDecoders {
	return nil
}

// `ConvertHexToBech32` converts a hex string to a bech32 string.
func (c *Contract) ConvertHexToBech32(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	addr, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, precompile.ErrInvalidHexAddress
	}
	return []any{evmutils.AddressToAccAddress(addr)}, nil
}

// `ConvertBech32ToHexAddress` converts a bech32 string to a hex string.
func (c *Contract) ConvertBech32ToHexAddress(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	addr, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, precompile.ErrInvalidString
	}
	accAddr, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return nil, err
	}
	return []any{evmutils.AccAddressToEthAddress(accAddr)}, nil
}
