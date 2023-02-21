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

import libtypes "pkg.berachain.dev/stargazer/lib/types"

//go:generate moq -out ./cloneable.mock.go -pkg mock ../ Cloneable

// `WrappedCloneableMock` is a mock for the `Cloneable` interface.
var _ libtypes.Cloneable[*WrappedCloneableMock] = &WrappedCloneableMock{}

// `WrappedCloneableMock` is a mock for the `Cloneable` interface.
// It wraps the `CloneableMock` and adds a `val` field.
type WrappedCloneableMock struct {
	CloneableMock[WrappedCloneableMock]
	val int
}

// `NewWrappedCloneableMock` returns a new `WrappedCloneableMock`.
func NewWrappedCloneableMock[T any](val int) *WrappedCloneableMock {
	return &WrappedCloneableMock{
		CloneableMock: CloneableMock[WrappedCloneableMock]{
			CloneFunc: func() WrappedCloneableMock {
				return WrappedCloneableMock{}
			},
		},
		val: val,
	}
}

// `Clone` returns a clone of the mock.
func (mco *WrappedCloneableMock) Clone() *WrappedCloneableMock {
	mco.CloneableMock.Clone()
	return &WrappedCloneableMock{
		val: mco.val,
		CloneableMock: CloneableMock[WrappedCloneableMock]{
			CloneFunc: mco.CloneableMock.CloneFunc,
		},
	}
}

// `Val` returns the value of the mock.
func (mco *WrappedCloneableMock) Val() int {
	return mco.val
}
