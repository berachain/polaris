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

//nolint:ireturn // StackI uses generics.
package ds

// `StackI` is an interface that defines the methods that an items Stack must implement.
// items Stacks support holding cache entries and reverting to a certain index.
type StackI[Item any] interface {
	// `Peek` returns the Item at the top of the stack
	Peek() Item

	// `PeekAt` returns the Item at the given index.
	PeekAt(i int) Item

	// `Push` adds a new Item to the top of the stack. The Size method returns the current
	// number of entries in the items.
	Push(i Item)

	// `Pop` returns the Item at the top of the stack and removes it from the stack.
	Pop() Item

	// `PopToSize` discards all items entries after and including the given size.
	PopToSize(newSize int)

	// `Size` returns the current number of entries in the items.
	Size() int
}

// Compile-time check to ensure `Stack` implements `StackI`.
var _ StackI[any] = (*Stack[any])(nil)

// `Stack` is a struct that holds a slice of Items.
type Stack[Item any] struct {
	items []Item
}

// `NewStack` creates and returns a new `Stack` with an no items.
func NewStack[Item any]() *Stack[Item] {
	return &Stack[Item]{
		items: make([]Item, 0),
	}
}

// `Push` implements `StackI`.
func (s *Stack[Item]) Push(i Item) {
	s.items = append(s.items, i)
}

// `Size` implements `StackI`.
func (s *Stack[Item]) Size() int {
	return len(s.items)
}

// `Peek` implements `StackI`.
func (s *Stack[Item]) Peek() Item {
	return s.items[len(s.items)-1]
}

// `Peek` implements `StackI`.
func (s *Stack[Item]) PeekAt(i int) Item {
	return s.items[i]
}

// `Pop` implements `StackI`.
func (s *Stack[Item]) Pop() Item {
	newLen := len(s.items) - 1
	item := s.items[newLen]
	s.items = s.items[:newLen] // exclusive to chop off last item
	return item
}

// `PopToSize` implements `StackI`.
func (s *Stack[Item]) PopToSize(newSize int) {
	s.items = s.items[:newSize] // todo: help the GC?
}
