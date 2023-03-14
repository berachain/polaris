// SPDX-License-Identifier: Apache-2.0
//

package mock

//go:generate moq -out ./controllable.mock.go -skip-ensure -pkg mock ../ Controllable

// ControllableMock is a mock for the `Controllable` interface.
func NewControllableMock1[K string]() *ControllableMock[K] {
	return &ControllableMock[K]{
		RevertToSnapshotFunc: func(n int) {},
		SnapshotFunc:         func() int { return 0 },
		RegistryKeyFunc:      func() K { return "object1" },
		FinalizeFunc:         func() {},
	}
}

func NewControllableMock2[K string]() *ControllableMock[K] {
	return &ControllableMock[K]{
		RevertToSnapshotFunc: func(n int) {},
		SnapshotFunc:         func() int { return 0 },
		RegistryKeyFunc:      func() K { return "object2" },
		FinalizeFunc:         func() {},
	}
}
