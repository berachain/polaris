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

package block

import (
	"github.com/berachain/polaris/cosmos/x/evm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// headerHashKeySize is the number of bytes in the header hash key: 1 (prefix) + 8 (block height).
const headerHashKeySize = 9

// headerHashKeyForHeight returns the key for the hash of the header at the given height.
func headerHashKeyForHeight(number int64) []byte {
	bz := make([]byte, headerHashKeySize)
	copy(bz, []byte{types.HeaderHashKeyPrefix})
	copy(bz[1:], sdk.Uint64ToBigEndian(uint64(number%prevHeaderHashes)))
	return bz
}
