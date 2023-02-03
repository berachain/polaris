package stack_test

import (
	"testing"

	"github.com/berachain/stargazer/lib/ds/stack"
)

const (
	initCapacity = 500
	numPushes    = 500
	numPopToSize = 10
)

func BenchmarkAStack(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		s := stack.NewA[*data]()
		for p := 0; p < numPushes; p++ {
			s.Push(newData())
		}

		ratio := numPushes / numPopToSize
		for j := numPopToSize; j > 0; j-- {
			s.PopToSize((j - 1) * ratio)
		}
	}
}

func BenchmarkStack(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		s := stack.New[*data](initCapacity)
		for p := 0; p < numPushes; p++ {
			s.Push(newData())
		}

		ratio := numPushes / numPopToSize
		for j := numPopToSize; j > 0; j-- {
			s.PopToSize((j - 1) * ratio)
		}
	}
}

// MOCKS BELOW.

type data struct {
	a int
	b string
	c uint
	d []byte
	e [20]byte
}

func newData() *data {
	return &data{
		a: 10023,
		b: "dfsafasd3",
		c: 4589403,
		d: []byte{0x4, 0x42, 0xfe},
		e: [20]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
	}
}
