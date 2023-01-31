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

package snapshot

import (
	"github.com/berachain/stargazer/lib/ds"
	"github.com/berachain/stargazer/lib/ds/stack"
	libtypes "github.com/berachain/stargazer/lib/types"
)

// `initJournalCapacity` is the initial capacity of the `journal` stack.
const initJournalCapacity = 1000000

// `Controller` implements the `libtypes.Snapshottable` interface.
var _ libtypes.Snapshottable = (*Controller[libtypes.Snapshottable])(nil)

// `Controller` conforms to the `libtypes.Snapshottable` interface and is used to sync
// snapshotting across multiple `libtypes.Snapshottable` objects.
type Controller[T libtypes.Snapshottable] struct {
	keyToSnapshottable map[string]T
	journal            ds.Stack[map[string]int]
}

// `NewController` returns a new `Controller` object.
func NewController[T libtypes.Snapshottable]() *Controller[T] {
	return &Controller[T]{
		keyToSnapshottable: make(map[string]T),
		journal:            stack.New[map[string]int](initJournalCapacity),
	}
}

// `Register` adds a `libtypes.Snapshottable` object to the `Controller`.
func (c *Controller[T]) Register(key string, object T) error {
	if _, ok := c.keyToSnapshottable[key]; ok {
		return ErrObjectAlreadyExists
	}
	c.keyToSnapshottable[key] = object
	return nil
}

// `Get` returns the `libtypes.Snapshottable` object with the given `id`.
func (c *Controller[T]) Get(key string) T {
	return c.keyToSnapshottable[key]
}

// `Snapshot` returns the current snapshot number.
func (c *Controller[T]) Snapshot() int {
	snap := make(map[string]int)
	for id, store := range c.keyToSnapshottable {
		snap[id] = store.Snapshot()
	}
	c.journal.Push(snap)

	return c.journal.Size()
}

// `RevertToSnapshot` reverts all `libtypes.Snapshottable` objects to the snapshot with
// the given `snap` number.
func (c *Controller[T]) RevertToSnapshot(id int) {
	lastestSnapshot := c.journal.Peek()
	for id, store := range c.keyToSnapshottable {
		// Only revert if exists. This is to handle the case where a
		// `libtypes.Snapshottable` object is added after a snapshot has been taken.
		if objRevision, ok := lastestSnapshot[id]; ok {
			store.RevertToSnapshot(objRevision)
		}
	}
	c.journal.PopToSize(id)
}

// `Finalize` is a no-op and is left to be extended by an implementation.
func (c *Controller[T]) Finalize() {}
