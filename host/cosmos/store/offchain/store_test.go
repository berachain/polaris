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

package offchain

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/sims"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestOffchain(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "host/cosmos/store/offchain")
}

var (
	dbDir = sims.NewAppOptionsWithFlagHome("./tmp/berachain")
	name  = "indexer-test123"
)

var _ = Describe("offchainStore", func() {
	var (
		byte1 = []byte{1}
		byte2 = []byte{2}
		byte3 = []byte{3}
		byte4 = []byte{4}
		store = NewOffChainKVStore(name, dbDir)
	)
	It("checks for write to buffer", func() {
		store.Set(byte1, byte2)
		store.Set(byte3, byte4)
		Expect(store.Get(byte1)).To(Equal(byte2))
		Expect(store.Get(byte3)).To(Equal(byte4))
	})

	It("checks for write to disk", func() {
		store.Set(byte1, byte2)
		store.Set(byte2, byte3)
		store.Write()
		store.Delete(byte1)
		Expect(store.Has(byte1)).Should(BeFalse())
	})
})
