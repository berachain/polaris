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

	"pkg.berachain.dev/polaris/cosmos/testing/types/mock"
	mocktypes "pkg.berachain.dev/polaris/cosmos/testing/types/mock/interfaces/mock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// TestAnteDecorator is a mock implementation of sdk.AnteDecorator for testing purposes.
type TestAnteDecorator struct{}

// AnteHandle returns a custom error if called.
func (f TestAnteDecorator) AnteHandle(
	ctx sdk.Context, _ sdk.Tx, _ bool, _ sdk.AnteHandler,
) (sdk.Context, error) {
	return ctx, errors.New("ante_handle_called")
}

func TestAnteLib(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/lib/ante")
}

var _ = Describe("IgnoreDecorator", func() {
	var (
		ignoreDecorator   *IgnoreDecorator[TestAnteDecorator, *mocktypes.MsgMock]
		fakeAnteDecorator TestAnteDecorator
	)

	BeforeEach(func() {
		fakeAnteDecorator = TestAnteDecorator{}
		ignoreDecorator = NewIgnoreDecorator[TestAnteDecorator, *mocktypes.MsgMock](fakeAnteDecorator)
	})

	// Test case when the transaction contains the specified message type.
	Context("when the transaction contains the specified message type", func() {
		It("should bypass the wrapped decorator", func() {
			tx := mock.NewTx()
			tx.GetMsgsFunc = func() []sdk.Msg {
				return []sdk.Msg{mock.NewMsg()}
			}
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
			tx := mock.NewTx()
			tx.GetMsgsFunc = func() []sdk.Msg {
				return []sdk.Msg{}
			}
			ctx := sdk.Context{}
			next := func(sdk.Context, sdk.Tx, bool) (sdk.Context, error) {
				return ctx, nil
			}

			_, err := ignoreDecorator.AnteHandle(ctx, tx, false, next)
			Expect(err).To(MatchError("ante_handle_called"))
		})
	})
})
