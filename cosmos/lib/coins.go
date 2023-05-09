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

package lib

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"

	generated "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"
	"pkg.berachain.dev/polaris/cosmos/precompile"
	"pkg.berachain.dev/polaris/lib/utils"
)

// SdkCoinsToEvmCoins converts sdk.Coins into []generated.IBankModuleCoin.
// The []generated.IBankModuleCoin is just a representation of []Coin from the generated solidity bindings
// and is equivalent to any of the other []Coin types generated from their respective solidity bindings.
// i.e. []generated.IERC20Coin.
func SdkCoinsToEvmCoins(sdkCoins sdk.Coins) []generated.IBankModuleCoin {
	evmCoins := make([]generated.IBankModuleCoin, len(sdkCoins))
	for i, coin := range sdkCoins {
		evmCoins[i] = generated.IBankModuleCoin{
			Amount: coin.Amount.BigInt(),
			Denom:  coin.Denom,
		}
	}
	return evmCoins
}

// ExtractCoinsFromInput converts coins from input (of type any) into sdk.Coins.
func ExtractCoinsFromInput(coins any) (sdk.Coins, error) {
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
