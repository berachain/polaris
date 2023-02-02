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
	var h2 = common.BytesToHash([]byte{2})
	var a1 = common.BytesToAddress([]byte{3})
	var a2 = common.BytesToAddress([]byte{4})
	var ti = uint(10)

	BeforeEach(func() {
		l = utils.MustGetAs[*logs](NewLogs())
		l.Prepare(h1, ti)
		Expect(l.currenTxIndex).To(Equal(ti))
		Expect(l.currentTxHash).To(Equal(h1))
		Expect(l.txHashToLogs[h1].Capacity()).To(Equal(32))
	})

	It("should have the correct registry key", func() {
		Expect(l.RegistryKey()).To(Equal("logs"))
	})

	When("adding logs", func() {
		BeforeEach(func() {
			l.AddLog(&coretypes.Log{Address: a1})
			logs := l.GetLogs(h1, h2)
			Expect(len(logs)).To(Equal(1))
			Expect(logs[0].Address).To(Equal(a1))
			Expect(logs[0].BlockHash).To(Equal(h2))
		})

		It("should correctly snapshot and revert", func() {
			id := l.Snapshot()

			l.AddLog(&coretypes.Log{Address: a2})
			logs := l.GetLogs(h1, h2)
			Expect(len(logs)).To(Equal(2))
			Expect(logs[1].Address).To(Equal(a2))
			Expect(logs[1].BlockHash).To(Equal(h2))

			l.RevertToSnapshot(id)
			logs = l.GetLogs(h1, h2)
			Expect(len(logs)).To(Equal(1))
		})

		It("should corrctly finalize", func() {
			Expect(func() { l.Finalize() }).ToNot(Panic())
		})
	})
})
