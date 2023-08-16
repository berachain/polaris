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

package governance

import (
	"testing"

	"cosmossdk.io/math"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGovernancePrecompile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/precompile/governance")
}

var _ = Describe("Governance Precompile", func() {
	It("Should be able to marshal and unmarshal msgBz", func() {
		msg := banktypes.MsgSend{
			FromAddress: sdk.AccAddress([]byte("from")).String(),
			ToAddress:   sdk.AccAddress([]byte("from")).String(),
			Amount:      sdk.NewCoins(sdk.NewCoin("abera", math.NewInt(100))),
		}

		msgAny, err := codectypes.NewAnyWithValue(&msg)
		Expect(err).ToNot(HaveOccurred())

		msgBz, err := msgAny.Marshal()
		Expect(err).ToNot(HaveOccurred())

		anys, err := UnmarshalAnyBzSlice([][]byte{msgBz})
		Expect(err).ToNot(HaveOccurred())

		typeURL := anys[0].GetTypeUrl()
		Expect(typeURL).To(Equal("/cosmos.bank.v1beta1.MsgSend"))
	})

})
