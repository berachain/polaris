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

package queue

import (
	"sync/atomic"
)

// LockfreeQueue represents a FIFO structure with operations to enqueue
// and dequeue generic values.
// Reference: https://www.cs.rochester.edu/research/synchronization/pseudocode/queues.html
type LockFreeQueue[T any] struct {
	head *atomic.Pointer[node[T]]
	tail *atomic.Pointer[node[T]]
}

// node represents a node in the queue.
type node[T any] struct {
	value T
	next  atomic.Pointer[node[T]]
}

// newNode creates and initializes a node.
func newNode[T any](v T) *node[T] {
	return &node[T]{value: v}
}

// NewQueue creates and initializes a LockFreeQueue.
func NewLockFreeQueue[T any]() *LockFreeQueue[T] {
	var head atomic.Pointer[node[T]]
	var tail atomic.Pointer[node[T]]
	var n = node[T]{}
	head.Store(&n)
	tail.Store(&n)
	return &LockFreeQueue[T]{
		head: &head,
		tail: &tail,
	}
}

// Enqueue adds a series of Request to the queue.
func (q *LockFreeQueue[T]) Enqueue(v T) {
	n := newNode(v)
	for {
		tail := q.tail.Load()
		next := tail.next.Load()
		if tail == q.tail.Load() {
			if next == nil {
				if tail.next.CompareAndSwap(next, n) {
					q.tail.CompareAndSwap(tail, n)
					return
				}
			} else {
				q.tail.CompareAndSwap(tail, next)
			}
		}
	}
}

// Dequeue removes a Request from the queue.
//
//nolint:nestif // its okay.
func (q *LockFreeQueue[T]) Dequeue() T {
	var t T
	for {
		head := q.head.Load()
		tail := q.tail.Load()
		next := head.next.Load()
		if head == q.head.Load() {
			if head == tail {
				if next == nil {
					return t
				}
				q.tail.CompareAndSwap(tail, next)
			} else {
				v := next.value
				if q.head.CompareAndSwap(head, next) {
					return v
				}
			}
		}
	}
}

// Check if the queue is empty.
func (q *LockFreeQueue[T]) IsEmpty() bool {
	return q.head.Load() == q.tail.Load()
}
