// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package mock

import libtypes "github.com/berachain/stargazer/lib/types"

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
