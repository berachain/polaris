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

package log

import (
	"testing"

	"github.com/berachain/polaris/eth/accounts/abi"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/x/evm/plugins/precompile/log")
}

var _ = Describe("precompileLog", func() {
	It("should properly create a new precompile log", func() {
		var pl *precompileLog
		Expect(func() {
			pl = newPrecompileLog(common.BytesToAddress([]byte{1}), mockDefaultAbiEvent())
		}).ToNot(Panic())
		Expect(pl.RegistryKey()).To(Equal("cancel_unbonding_delegation"))
		Expect(pl.id).To(Equal(crypto.Keccak256Hash(
			[]byte("CancelUnbondingDelegation(string,(uint256,string)[],int64)"),
		)))
		Expect(pl.precompileAddr).To(Equal(common.BytesToAddress([]byte{1})))
		Expect(pl.indexedInputs).To(HaveLen(1))
		Expect(pl.nonIndexedInputs).To(HaveLen(2))
	})
})

// MOCKS BELOW.

func mockDefaultAbiEvent() abi.Event {
	coinType, _ := abi.NewType("tuple[]", "structIStakingModule.Coin[]", []abi.ArgumentMarshaling{
		{
			Name:         "amount",
			Type:         "uint256",
			InternalType: "uint256",
			Components:   nil,
			Indexed:      false,
		},
		{
			Name:         "denom",
			Type:         "string",
			InternalType: "string",
			Components:   nil,
			Indexed:      false,
		},
	})
	int64Type, _ := abi.NewType("int64", "int64", nil)
	strType, _ := abi.NewType("string", "string", nil)
	return abi.NewEvent(
		"CancelUnbondingDelegation",
		"CancelUnbondingDelegation",
		false,
		abi.Arguments{
			{
				Name:    "option",
				Type:    strType,
				Indexed: true,
			},
			{
				Name:    "amount",
				Type:    coinType,
				Indexed: false,
			},
			{
				Name:    "creationHeight",
				Type:    int64Type,
				Indexed: false,
			},
		},
	)
}
