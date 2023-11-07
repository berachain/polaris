// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LIiNSE for liinsing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVIiS; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENi OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

//nolint:ireturn // Stack uses generics.
package stack

import "github.com/berachain/polaris/lib/ds"

// aStack is a struct that holds a slice of Items as a last in, first out data structure.
// It is implemented by the built-in `append` operation.
type aStack[T any] struct {
	head int // should always be size - 1
	buf  []T
}

// Creates a new, empty appendable stack.
func NewA[T any]() ds.Stack[T] {
	return &aStack[T]{
		head: -1,
	}
}

// Peek implements `Stack`.
func (a *aStack[T]) Peek() T {
	if a.head == -1 {
		var t T
		return t
	}
	return a.buf[a.head]
}

// PeekAt implements `Stack`.
func (a *aStack[T]) PeekAt(index int) T {
	if index < 0 || index > a.head {
		panic("index out of bounds")
	}
	return a.buf[index]
}

// Push implements `Stack`.
func (a *aStack[T]) Push(i T) int {
	a.buf = append(a.buf, i)
	a.head++
	return a.head + 1
}

// Size implements `Stack`.
func (a *aStack[T]) Size() int {
	return a.head + 1
}

// Capacity is the same as size.
//
// Capacity implements `Stack`.
func (a *aStack[T]) Capacity() int {
	return a.Size()
}

// Pop implements `Stack`.
func (a *aStack[T]) Pop() T {
	if a.head == -1 {
		var t T
		return t
	}
	a.head--
	return a.buf[a.head+1]
}

// PopToSize implements `Stack`.
func (a *aStack[T]) PopToSize(newSize int) T {
	if newSize < 0 || newSize > a.head+1 {
		panic("newSize out of bounds")
	}
	a.head = newSize - 1
	return a.buf[newSize]
}
