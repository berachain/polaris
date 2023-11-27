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

package stack_test

import (
	"testing"

	"github.com/berachain/polaris/lib/ds/stack"
)

const (
	// initCapacity should be close to the predicted size of the stack to avoid manual resizing.
	initCapacity = 500
	numPushes    = 500
	numPopToSize = 10
)

// Benchmarks (of pushing to the stack and popping the stack to a size) show that the
// appendable-stack is narrowly slower than the regular stack, which uses pre-allocated buffer with
// an initial capacity and manual resizing.

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
