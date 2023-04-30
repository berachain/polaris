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

package antelib

import (
	"errors"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// FakeAnteDecorator is a mock implementation of sdk.AnteDecorator for testing purposes.
type FakeAnteDecorator struct{}

// AnteHandle returns a custom error if called.
func (f FakeAnteDecorator) AnteHandle(
	ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler,
) (sdk.Context, error) {
	return ctx, errors.New("ante_handle_called")
}

// FakeMsg is a mock implementation of sdk.Msg for testing purposes.
type FakeMsg struct{}

func (f FakeMsg) Route() string                { return "" }
func (f FakeMsg) Type() string                 { return "fake_msg" }
func (f FakeMsg) ValidateBasic() error         { return nil }
func (f FakeMsg) GetSignBytes() []byte         { return []byte{} }
func (f FakeMsg) GetSigners() []sdk.AccAddress { return nil }
func (f FakeMsg) ProtoMessage()                {}
func (f FakeMsg) Reset()                       {}
func (f FakeMsg) String() string               { return "fake bing bong" }

// FakeTx is a mock implementation of sdk.Tx for testing purposes.
type FakeTx struct {
	msgs []sdk.Msg
}

func (t FakeTx) GetMsgs() []sdk.Msg {
	return t.msgs
}

func (t FakeTx) ValidateBasic() error {
	return nil
}

func TestAnteLib(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/lib/ante")
}

var _ = Describe("IgnoreDecorator", func() {
	var (
		ignoreDecorator   *IgnoreDecorator[FakeAnteDecorator, FakeMsg]
		fakeAnteDecorator FakeAnteDecorator
	)

	BeforeEach(func() {
		fakeAnteDecorator = FakeAnteDecorator{}
		ignoreDecorator = NewIgnoreDecorator[FakeAnteDecorator, FakeMsg](fakeAnteDecorator)
	})

	// Test case when the transaction contains the specified message type.
	Context("when the transaction contains the specified message type", func() {
		It("should bypass the wrapped decorator", func() {
			tx := FakeTx{[]sdk.Msg{FakeMsg{}}}
			ctx := sdk.Context{}
			next := func(sdk.Context, sdk.Tx, bool) (sdk.Context, error) {
				return ctx, nil
			}

			newCtx, err := ignoreDecorator.AnteHandle(ctx, tx, false, next)
			Expect(newCtx).To(Equal(ctx))
			Expect(err).ToNot(HaveOccurred())
		})
	})

	// Test case when the transaction does not contain the specified message type.
	Context("when the transaction does not contain the specified message type", func() {
		It("should call the wrapped decorator's AnteHandle", func() {
			tx := FakeTx{[]sdk.Msg{}}
			ctx := sdk.Context{}
			next := func(sdk.Context, sdk.Tx, bool) (sdk.Context, error) {
				return ctx, nil
			}

			_, err := ignoreDecorator.AnteHandle(ctx, tx, false, next)
			Expect(err).To(MatchError("ante_handle_called"))
		})
	})
})
