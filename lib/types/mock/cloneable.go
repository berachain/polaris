// SPDX-License-Identifier: Apache-2.0
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package mock

import libtypes "github.com/berachain/polaris/lib/types"

//go:generate moq -out ./cloneable.mock.go -pkg mock ../ Cloneable

// WrappedCloneableMock is a mock for the `Cloneable` interface.
var _ libtypes.Cloneable[*WrappedCloneableMock] = &WrappedCloneableMock{}

// WrappedCloneableMock is a mock for the `Cloneable` interface.
// It wraps the `CloneableMock` and adds a `val` field.
type WrappedCloneableMock struct {
	CloneableMock[WrappedCloneableMock]
	val int
}

// NewWrappedCloneableMock returns a new `WrappedCloneableMock`.
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

// Clone returns a clone of the mock.
func (mco *WrappedCloneableMock) Clone() *WrappedCloneableMock {
	mco.CloneableMock.Clone()
	return &WrappedCloneableMock{
		val: mco.val,
		CloneableMock: CloneableMock[WrappedCloneableMock]{
			CloneFunc: mco.CloneableMock.CloneFunc,
		},
	}
}

// Val returns the value of the mock.
func (mco *WrappedCloneableMock) Val() int {
	return mco.val
}
