// SPDX-License-Identifier: Apache-2.0
//

package mock

import libtypes "pkg.berachain.dev/polaris/lib/types"

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
