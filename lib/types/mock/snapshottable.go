package mock

//go:generate moq -out ./snapshottable.mock.go -pkg mock ../ Snapshottable

// SnapshottableMock is a mock for the `Snapshottable` interface.
func NewSnapshottableMock() *SnapshottableMock {
	return &SnapshottableMock{
		RevertToSnapshotFunc: func(n int) {},
		SnapshotFunc:         func() int { return 0 },
	}
}
