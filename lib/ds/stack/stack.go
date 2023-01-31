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

const (
	resizeRatio = 2
	two         = 2
)

// `Stack` is a struct that holds a slice of Items.
// Last in, first out data structure.
type stack[T any] struct {
	size     int
	capacity int

	buf []T
}

// Creates a new, empty stack.
func New[T any](capacity int) ds.Stack[T] {
	return &stack[T]{
		capacity: capacity,
		buf:      make([]T, capacity),
	}
}

// `Peek` implements `Stack`.
func (s *stack[T]) Peek() T {
	if s.size == 0 {
		var t T
		return t
	}
	return s.buf[s.size-1]
}

// `PeekAt` implements `Stack`.
func (s *stack[T]) PeekAt(index int) T {
	if index < 0 || index >= s.size {
		panic("index out of bounds")
	}
	return s.buf[index]
}

// `Push` implements `Stack`.
func (s *stack[T]) Push(i T) {
	s.expandIfRequired()
	s.buf[s.size] = i
	s.size++
}

// `Size` implements `Stack`.
func (s *stack[T]) Size() int {
	return s.size
}

// `Capacity` implements `Stack`.
func (s *stack[T]) Capacity() int {
	return s.capacity
}

// `Pop` implements `Stack`.
func (s *stack[T]) Pop() T {
	s.size--
	s.shrinkIfRequired()
	return s.buf[s.size]
}

// `PopToSize` implements `Stack`.
func (s *stack[T]) PopToSize(newSize int) {
	if newSize < 0 || newSize > s.size {
		panic("newSize out of bounds")
	}

	s.size = newSize
	s.shrinkIfRequired()
}

// `expandIfRequired` expands the stack if the size is equal to the capacity.
func (s *stack[T]) expandIfRequired() {
	if s.size < s.capacity {
		return
	}

	newBuf := make([]T, (s.capacity*resizeRatio)/two)
	s.buf = append(s.buf, newBuf...)
	s.capacity *= resizeRatio
}

// `shrinkIfRequired` shrinks the stack if the size is less than the capacity/resizeRatio.
func (s *stack[T]) shrinkIfRequired() {
	if newCap := s.capacity / resizeRatio; s.size < newCap {
		newBuf := make([]T, newCap)
		copy(newBuf, s.buf)
		s.buf = newBuf
		s.capacity = newCap
	}
}
