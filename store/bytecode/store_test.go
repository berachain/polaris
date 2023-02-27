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

package bytecode

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/sims"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/crypto"
)

func TestByteCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "store/bytecode")
}

var _ = Describe("bytecodeStore", func() {
	var (
		addr  = common.BytesToAddress([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
		code1 = []byte{1}
		dbDir = sims.NewAppOptionsWithFlagHome("/tmp/berachain")
		store = NewByteCodeStore(dbDir)
	)

	It("should set and get byte code", func() {
		store.StoreByteCode(addr, code1)
		codeHash := crypto.Keccak256Hash(code1)
		code, err := store.GetByteCode(addr, codeHash)
		Expect(err).To(BeNil())
		Expect(code).To(Equal(code1))
	})

	It("should fail to get byte code if the code hash does not match", func() {
		store.StoreByteCode(addr, code1)
		codeHash := crypto.Keccak256Hash([]byte{2})
		code, err := store.GetByteCode(addr, codeHash)
		Expect(err).To(Equal(ErrByteCodeDoesNotMatch))
		Expect(code).To(BeNil())
	})

	It("should iterate over byte code", func() {
		log := make([]byte, 0)

		store.StoreByteCode(addr, code1)
		store.IterateByteCode(nil, nil, func(addr common.Address, code []byte) bool {
			log = append(log, code...)
			return true // break the iteration
		})

		Expect(log).To(Equal(code1))
	})

	It("should set and get version", func() {
		store.SetVersion(1)
		Expect(store.GetVersion()).To(Equal(int64(1)))
	})
})
