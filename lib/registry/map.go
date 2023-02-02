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

package registry

import (
	libtypes "github.com/berachain/stargazer/lib/types"
)

// `mapRegistry` is a simple implementation of `Registry` that uses a map as the underlying data
// structure.
type mapRegistry[K comparable, T libtypes.Registrable[K]] struct {
	// items is the map of items in the registry.
	items map[K]T
}

// `NewMap` creates and returns a new `mapRegistry`.
//
//nolint:revive // only used as Registry interface.
func NewMap[K comparable, T libtypes.Registrable[K]]() *mapRegistry[K, T] {
	return &mapRegistry[K, T]{
		items: make(map[K]T),
	}
}

// `Get` returns an item using its ID.
func (mr *mapRegistry[K, T]) Get(id K) T {
	return mr.items[id]
}

// `Register` adds an item to the registry.
func (mr *mapRegistry[K, T]) Register(item T) error {
	mr.items[item.RegistryKey()] = item
	return nil
}

// `Remove` removes an item from the registry.
func (mr *mapRegistry[K, T]) Remove(id K) {
	delete(mr.items, id)
}

// `Exists` returns true if the item exists in the registry.
func (mr *mapRegistry[K, T]) Exists(id K) bool {
	_, ok := mr.items[id]
	return ok
}

// `Iterate` returns the underlying map.
func (mr *mapRegistry[K, T]) Iterate() map[K]T {
	return mr.items
}
