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

package plugin

import (
	coretypes "github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Logs", func() {
	var l *logs
	var h1 = common.BytesToHash([]byte{1})
	var a1 = common.BytesToAddress([]byte{3})
	var a2 = common.BytesToAddress([]byte{4})

	BeforeEach(func() {
		l = utils.MustGetAs[*logs](NewLogs())
		Expect(l.Capacity()).To(Equal(32))
		l.PrepareForTx(h1)
		Expect(l.currentTxHash).To(Equal(h1))
	})

	It("should have the correct registry key", func() {
		Expect(l.RegistryKey()).To(Equal("logs"))
	})

	When("adding logs", func() {
		BeforeEach(func() {
			l.AddLog(&coretypes.Log{Address: a1})
			Expect(l.Size()).To(Equal(1))
			Expect(l.PeekAt(0).Address).To(Equal(a1))
		})

		It("should correctly snapshot and revert", func() {
			id := l.Snapshot()

			l.AddLog(&coretypes.Log{Address: a2})
			Expect(l.Size()).To(Equal(2))
			Expect(l.PeekAt(1).Address).To(Equal(a2))

			l.RevertToSnapshot(id)
			Expect(l.Size()).To(Equal(1))
		})

		It("should corrctly finalize", func() {
			Expect(func() { l.Finalize() }).ToNot(Panic())
		})
	})
})
