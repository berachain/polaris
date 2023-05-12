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

package misc_test

import (
	"math/big"
	"reflect"
	"testing"

	"pkg.berachain.dev/polaris/cosmos/testing/integration"
	"pkg.berachain.dev/polaris/eth/accounts/abi"
	"pkg.berachain.dev/polaris/eth/common/hexutil"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type MsgSendEnergy struct {
	To     *big.Int `json:"to"`
	From   *big.Int `json:"from"`
	Amount *big.Int `json:"amount"`
}

func TestDecoding(t *testing.T) {
	RegisterTestingT(t)
	// bytes from abi.encoding the following solidity struct:
	//    struct MsgSendEnergy {
	//        uint to;
	//        uint from;
	//        uint amount;
	//    }
	// with values: 3293, 2385, 23509
	expectedMsg := MsgSendEnergy{
		To:     big.NewInt(3293),
		From:   big.NewInt(2385),
		Amount: big.NewInt(23509),
	}
	bzStr := "0x0000000000000000000000000000000000000000000000000000000000000cdd00000000000000000000000000000000000000000000000000000000000009510000000000000000000000000000000000000000000000000000000000005bd5"
	bz, err := hexutil.Decode(bzStr)
	Expect(err).To(BeNil())

	// make MsgSendEnergy abi type.
	msgSendEnergyType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "to", Type: "uint256"},
		{Name: "from", Type: "uint256"},
		{Name: "amount", Type: "uint256"},
	})
	Expect(err).To(BeNil())
	msgSendEnergyType.TupleType = reflect.TypeOf(MsgSendEnergy{})

	args := abi.Arguments{{Type: msgSendEnergyType}}
	unpacked, err := args.Unpack(bz)
	Expect(err).To(BeNil())

	unpackedMsg, ok := unpacked[0].(MsgSendEnergy)
	Expect(ok).To(Equal(true))
	Expect(unpackedMsg).To(Equal(expectedMsg))

	// check if we can pack a go struct and unpack it back
	packedBz, err := args.Pack(unpackedMsg)
	Expect(err).To(BeNil())
	// bytes should be the same
	Expect(packedBz).To(Equal(bz))

	unpacked, err = args.Unpack(packedBz)
	Expect(err).To(BeNil())

	unpackedMsg, ok = unpacked[0].(MsgSendEnergy)
	Expect(err).To(BeNil())
	Expect(unpackedMsg).To(Equal(expectedMsg))
}

func TestMiscellaneousPrecompile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/testing/integration/precompile/misc")
}

var tf *integration.TestFixture

var _ = SynchronizedBeforeSuite(func() []byte {
	// Setup the network and clients here.
	tf = integration.NewTestFixture(GinkgoT())
	return nil
}, func(data []byte) {})
