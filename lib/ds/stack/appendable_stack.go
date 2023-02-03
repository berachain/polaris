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

import "github.com/berachain/stargazer/lib/ds"

type aStack[T any] struct {
	head int // should always be size - 1

	buf []T
}

func NewA[T any]() ds.Stack[T] {
	return &aStack[T]{
		head: -1,
	}
}

func (a *aStack[T]) Peek() T {
	return a.buf[a.head]
}

func (a *aStack[T]) PeekAt(index int) T {
	if index < 0 || index > a.head {
		panic("index out of bounds")
	}
	return a.buf[index]
}

func (a *aStack[T]) Push(i T) int {
	a.buf = append(a.buf, i)
	a.head++
	return a.head + 1
}

func (a *aStack[T]) Size() int {
	return a.head + 1
}

// same as size, capacity not supported
func (a *aStack[T]) Capacity() int {
	return a.Size()
}

func (a *aStack[T]) Pop() T {
	if a.head == -1 {
		var t T
		return t
	}
	a.head--
	return a.buf[a.head+1]
}

func (a *aStack[T]) PopToSize(newSize int) T {
	if newSize < 0 || newSize > a.head+1 {
		panic("newSize out of bounds")
	}
	a.head = newSize - 1
	return a.buf[newSize]
}
