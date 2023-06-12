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

package mock

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/cosmos/testing/types/mock/interfaces/mock"
)

// FakeMsg is a mock implementation of sdk.Msg for testing purposes.
func NewMsg() *mock.MsgMock {
	mockedMsg := &mock.MsgMock{
		ProtoMessageFunc: func() {
			panic("mock out the ProtoMessage method")
		},
		ResetFunc: func() {
			panic("mock out the Reset method")
		},
		StringFunc: func() string {
			panic("mock out the String method")
		},
	}
	return mockedMsg
}

// FakeMsg is a mock implementation of sdk.Msg for testing purposes.
func NewTx() *mock.TxMock {
	// make and configure a mocked interfaces.Tx
	mockedTx := &mock.TxMock{
		GetMsgsFunc: func() []sdk.Msg {
			panic("mock out the GetMsgs method")
		},
	}
	return mockedTx
}
