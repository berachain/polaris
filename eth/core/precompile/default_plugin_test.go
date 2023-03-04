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

package precompile_test

import (
	"github.com/ethereum/go-ethereum/core/vm"

	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/core/precompile"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

//nolint:lll // test data.
const precompInput = `a8b53bdf3306a35a7103ab5504a0c9b492295564b6202b1942a84ef300107281000000000000000000000000000000000000000000000000000000000000001b307835653165303366353363653138623737326363623030393366663731663366353366356337356237346463623331613835616138623838393262346538621122334455667788991011121314151617181920212223242526272829303132`

var _ = Describe("Default Plugin", func() {
	var dp precompile.Plugin

	BeforeEach(func() {
		dp = precompile.NewDefaultPlugin()
	})

	When("running a stateless contract", func() {
		It("should run out of gas", func() {
			ret, remainingGas, err := dp.Run(nil, &mockStateless{&mockBase{}}, nil, common.Address{}, nil, 5, false)
			Expect(ret).To(BeNil())
			Expect(remainingGas).To(Equal(uint64(0)))
			Expect(err.Error()).To(Equal("out of gas"))
		})

		It("should run a geth contract", func() {
			pc := vm.PrecompiledContractsHomestead[common.BytesToAddress([]byte{1})]
			_, remainingGas, err := dp.Run(nil, pc, []byte(precompInput), common.Address{}, nil, 3000, true)
			Expect(remainingGas).To(Equal(uint64(0)))
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
