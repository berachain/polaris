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

package types

import (
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/eth/accounts"
)

var initConfig sync.Once

// SetupCosmosConfig sets up the Cosmos SDK configuration to be compatible with the
// semantics of etheruem.
func SetupCosmosConfig() {
	initConfig.Do(func() {
		// set the address prefixes
		config := sdk.GetConfig()
		SetBip44CoinType(config)
		config.Seal()
	})
}

// SetBip44CoinType sets the global coin type to be used in hierarchical deterministic wallets.
func SetBip44CoinType(config *sdk.Config) {
	config.SetCoinType(accounts.Bip44CoinType)
	config.SetPurpose(sdk.Purpose)
}
