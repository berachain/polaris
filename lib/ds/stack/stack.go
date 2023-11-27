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

import (
	"github.com/berachain/polaris/lib/ds"
)

const (
	resizeRatio = 2
	two         = 2
)

// stack is a struct that holds a slice of Items as a last in, first out data structure.
// It is implemented by pre-allocating a buffer with a capacity.
type stack[T any] struct {
	size            int
	capacity        int
	initialCapacity int
	buf             []T
}

// Creates a new, empty stack with the given initial capacity.
func New[T any](initialCapacity int) ds.Stack[T] {
	return &stack[T]{
		capacity:        initialCapacity,
		buf:             make([]T, initialCapacity),
		initialCapacity: initialCapacity,
	}
}

// Peek implements `Stack`.
func (s *stack[T]) Peek() T {
	if s.size == 0 {
		var t T
		return t
	}
	return s.buf[s.size-1]
}

// PeekAt implements `Stack`.
func (s *stack[T]) PeekAt(index int) T {
	if index >= s.size {
		panic("index out of bounds")
	}
	return s.buf[index]
}

// Push implements `Stack`.
func (s *stack[T]) Push(i T) int {
	s.expandIfRequired()
	s.buf[s.size] = i
	s.size++
	return s.size
}

// Size implements `Stack`.
func (s *stack[T]) Size() int {
	return s.size
}

// Capacity implements `Stack`.
func (s *stack[T]) Capacity() int {
	return s.capacity
}

// Pop implements `Stack`.
func (s *stack[T]) Pop() T {
	if s.size == 0 {
		var t T
		return t
	}
	s.size--
	s.shrinkIfRequired()
	return s.buf[s.size]
}

// PopToSize implements `Stack`.
func (s *stack[T]) PopToSize(newSize int) T {
	if newSize > s.size {
		panic("newSize out of bounds")
	} else if newSize == s.size {
		t := new(T)
		return *t
	}
	s.size = newSize
	lastElemPopped := s.buf[s.size]
	s.shrinkIfRequired()
	return lastElemPopped
}

// expandIfRequired expands the stack if the size is equal to the capacity.
func (s *stack[T]) expandIfRequired() {
	if s.size < s.capacity {
		return
	}
	newCapacity := max(s.initialCapacity, (s.capacity*resizeRatio)/two)
	s.buf = append(s.buf, make([]T, newCapacity)...)
	s.capacity *= resizeRatio
}

// shrinkIfRequired shrinks the stack if the size is less than the capacity/resizeRatio.
func (s *stack[T]) shrinkIfRequired() {
	if newCapacity := max(s.initialCapacity, s.capacity/resizeRatio); s.size < newCapacity {
		newBuf := make([]T, newCapacity)
		copy(newBuf, s.buf)
		s.buf = newBuf
		s.capacity = newCapacity
	}
}
