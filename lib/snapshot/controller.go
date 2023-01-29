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

// `initTrackerStackCapacity` is the initial capacity of the `snapTracker` stack.
const initTrackerStackCapacity = 16

// `Controller` implements the `lib.Snapshottable` interface.
var _ libtypes.Snapshottable = (*Controller)(nil)

// `Controller` conforms to the `lib.Snapshottable` interface and is used to sync
// snapshotting across multiple `lib.Snapshottable` objects.
type Controller struct {
	tracked     map[string]libtypes.Snapshottable
	snapTracker ds.Stack[map[string]int]
}

// `NewController` returns a new `Controller` object.
func NewController() *Controller {
	return &Controller{
		snapTracker: stack.New[map[string]int](initTrackerStackCapacity),
		tracked:     make(map[string]libtypes.Snapshottable),
	}
}

// `Control` adds a `lib.Snapshottable` object to the `Controller`.
func (ctrl *Controller) Control(id string, object libtypes.Snapshottable) error {
	if _, ok := ctrl.tracked[id]; ok {
		return ErrObjectAlreadyExists
	}
	ctrl.tracked[id] = object
	return nil
}

// `Get` returns the `lib.Snapshottable` object with the given `id`.
func (ctrl *Controller) Get(id string) libtypes.Snapshottable {
	return ctrl.tracked[id]
}

// `Snapshot` returns the current snapshot number.
func (ctrl *Controller) Snapshot() int {
	snap := make(map[string]int)
	for id, store := range ctrl.tracked {
		snap[id] = store.Snapshot()
	}
	ctrl.snapTracker.Push(snap)

	return ctrl.snapTracker.Size()
}

// `RevertToSnapshot` reverts all `lib.Snapshottable` objects to the snapshot with
// the given `snap` number.
func (ctrl *Controller) RevertToSnapshot(revision int) {
	lastestSnapshot := ctrl.snapTracker.Peek()
	for id, store := range ctrl.tracked {
		// Only revert if exists. This is to handle the case where a
		// `lib.Snapshottable` object is added after a snapshot has been taken.
		if objRevision, ok := lastestSnapshot[id]; ok {
			store.RevertToSnapshot(objRevision)
		}
	}
	ctrl.snapTracker.PopToSize(revision)
}

// `Revision` returns a specific set of snapshot numbers for all `lib.Snapshottable`
// that are being tracked by the `Controller` at that revision number.
func (ctrl *Controller) Revision(revision int) map[string]int {
	// 1st revision is the 0th index.
	return ctrl.snapTracker.PeekAt(revision - 1)
}

// `LatestRevision` returns the current snapshot numbers for all `lib.Snapshottable`
// that are being tracked by the `Controller`.
func (ctrl *Controller) LatestRevision() map[string]int {
	return ctrl.snapTracker.Peek()
}

// `Finalize` is a no-op and is left to be extended by an implementation.
func (ctrl *Controller) Finalize() {}
