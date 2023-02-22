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
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Genesis", func() {
	It("fail if genesis is invalid", func() {
		params := DefaultParams()
		params.EvmDenom = ""
		state := NewGenesisState(*params, nil, nil)
		err := ValidateGenesis(*state)
		Expect(err).To(HaveOccurred())
	})

	It("should create new genesis state", func() {
		crs := []CodeRecord{
			{
				Address: "address",
				Code:    []byte("code"),
			},
		}
		srs := []StateRecord{
			{
				Address: "address",
				Slot:    []byte("slot"),
				Value:   []byte("value"),
			},
		}

		state := NewGenesisState(*DefaultParams(), crs, srs)
		Expect(state.CodeRecords).To(Equal(crs))
		Expect(state.StateRecords).To(Equal(srs))

		defaultGenesis := DefaultGenesis()
		Expect(defaultGenesis.Params).To(Equal(*DefaultParams()))
	})
})
