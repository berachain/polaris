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

package state

import "github.com/berachain/stargazer/lib/ds"

type Controller struct {
	stores      map[string]Store
	snapTracker ds.Stack[map[string]int]
}

func NewController() *Controller {
	return &Controller{
		stores: make(map[string]Store),
	}
}

func (ctrl *Controller) AddStore(store Store) {
	ctrl.stores[store.Name()] = store
}

func (ctrl *Controller) GetStore(name string) Store {
	return ctrl.stores[name]
}

func (ctrl *Controller) Snapshot() int {
	snap := make(map[string]int)
	for name, store := range ctrl.stores {
		snap[name] = store.Snapshot()
	}
	ctrl.snapTracker.Push(snap)
	return ctrl.snapTracker.Size()
}

func (ctrl *Controller) RevertToSnapshot(snap int) {
	top := ctrl.snapTracker.Peek()
	for name, store := range ctrl.stores {
		store.RevertToSnapshot(top[name])
	}
	ctrl.snapTracker.PopToSize(snap)
}
