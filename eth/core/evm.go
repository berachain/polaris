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

package core

import (
	"math/big"

	"github.com/berachain/stargazer/eth/common"
	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/lib/utils"
)

// Compile-time type assertion.
var _ vm.CanTransferFunc = canTransfer
var _ vm.TransferFunc = transfer

// `canTransfer` checks whether there are enough funds in the address' account to make a transfer.
// NOTE: This does not take the necessary gas in to account to make the transfer valid.
func canTransfer(sdb vm.GethStateDB, addr common.Address, amount *big.Int) bool {
	return sdb.GetBalance(addr).Cmp(amount) >= 0
}

// `transfer` subtracts amount from sender and adds amount to recipient using a `vm.GethStateDB`.
func transfer(sdb vm.GethStateDB, sender, recipient common.Address, amount *big.Int) {
	utils.MustGetAs[vm.StargazerStateDB](sdb).TransferBalance(sender, recipient, amount)
}
