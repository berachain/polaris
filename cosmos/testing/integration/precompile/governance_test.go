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

package precompile

import (
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	bindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"

	// . "pkg.berachain.dev/polaris/cosmos/testing/integration/utils"
	"pkg.berachain.dev/polaris/cosmos/testing/network"
	"pkg.berachain.dev/polaris/eth/common"
)

var _ = Describe("Governance", func() {
	It("should call functions on the precompile directly", func() {

		// Setup the governance msg.
		initDeposit := sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 1))
		govAcctAddr := common.HexToAddress("0x7b5Fe22B5446f7C62Ea27B8BD71CeF94e03f3dF2")
		govAccAddr := cosmlib.AddressToAccAddress(govAcctAddr)
		callerAccAddres := cosmlib.AddressToAccAddress(network.TestAddress)
		message := &banktypes.MsgSend{
			FromAddress: govAccAddr.String(),
			ToAddress:   callerAccAddres.String(),
			Amount:      initDeposit,
		}
		metadata := "metadata"
		title := "title"
		summary := "summary"
		msg, err := codectypes.NewAnyWithValue(message)
		Expect(err).ToNot(HaveOccurred())

		// Call the precompile.
		txr := tf.GenerateTransactOpts("")
		txr.Value = big.NewInt(100)
		msgBz, err := msg.Marshal()
		Expect(err).ToNot(HaveOccurred())
		tx, err := governancePrecompile.SubmitProposal(
			txr,
			msgBz, // TODO: How to handle any type?.
			[]bindings.IGovernanceModuleCoin{
				{
					Amount: 100,
					Denom:  "usdc",
				},
			},
			metadata,
			title,
			summary,
			false,
		)
		Expect(err).ToNot(HaveOccurred())
		fmt.Println(tx)
	})
})
