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

package lib_test

import (
	"fmt"

	"github.com/holiman/uint256"

	storetypes "cosmossdk.io/store/types"

	"pkg.berachain.dev/polaris/cosmos/lib"
	testutil "pkg.berachain.dev/polaris/cosmos/testing/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Salt", func() {
	var nonceStore storetypes.BasicKVStore

	BeforeEach(func() {
		nonceStore = testutil.NewContext().KVStore(testutil.ERC20Key)
	})

	It("should be unique and deterministic", func() {
		salts := make(map[uint256.Int]struct{})
		orderedSalts := make([]uint256.Int, 20000)

		for i := 0; i < 10000; i++ {
			salt := lib.UniqueDeterminsticSalt(nonceStore, []byte("test"))
			_, found := salts[*salt]
			Expect(found).To(BeFalse())
			salts[*salt] = struct{}{}
			orderedSalts[i] = *salt
		}

		for i := 0; i < 10000; i++ {
			salt := lib.UniqueDeterminsticSalt(nonceStore, []byte(fmt.Sprintf("test%d", i)))
			_, found := salts[*salt]
			Expect(found).To(BeFalse())
			salts[*salt] = struct{}{}
			orderedSalts[i+10000] = *salt
		}
	})
})
