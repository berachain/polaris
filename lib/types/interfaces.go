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

package types

// `Cloneable` is an interface that defines a `Clone` method.
type Cloneable[T any] interface {
	Clone() T
}

// `Snapshottable` is an interface that defines methods for snapshotting and reverting
// a logical unit of data.
type Snapshottable interface {
	// `Snapshot` returns an identifier for the current revision of the data.
	Snapshot() int

	// `RevertToSnapshot` reverts the data to a previous version
	RevertToSnapshot(int)
}

// `Registrable` is an interface that all objects that can be registered in a
// `Registry` must implement.
type Registrable[K comparable] interface {
	// `RegistryKey` returns the key that will be used to register the object.
	RegistryKey() K
}

// `Registry` is an interface that all objects that can be used as a registry
// must implement.
type Registry[K comparable, T Registrable[K]] interface {
	// Get return an item using its ID. It returns nil if the ID does not exist.
	Get(K) T

	// Register adds an item to the registry.
	Register(T) error

	// Remove removes an item from the registry.
	Remove(K)

	// Exists returns true if the item exists in the registry.
	Exists(K) bool

	// Iterate returns an iterable map of the registry.
	Iterate() map[K]T
}

// `Controllable` defines a type which can be controlled.
type Controllable[K comparable] interface {
	Snapshottable
	Registrable[K]

	Write()
}

// `Controller` is an interface for controller types.
type Controller[K comparable, T Controllable[K]] interface {
	Snapshottable
	Registry[K, T]

	Finalize()
}
