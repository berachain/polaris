// Copyright (C) 2022, Berachain Foundation. All rights reserved.
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

//nolint:ireturn // all `CacheEntries` must adhere to the same interface.
package state

import (
	"github.com/berachain/stargazer/core/state/store/journal"
	"github.com/berachain/stargazer/lib/common"
)

var (
	_ journal.CacheEntry = &RefundChange{}
	_ journal.CacheEntry = &AddLogChange{}
)

type (
	AddLogChange struct {
		sdb    *StateDB
		txHash common.Hash
	}
	RefundChange struct {
		sdb  *StateDB
		prev uint64
	}
)

// ==============================================================================
// AddLogChange
// ==============================================================================

// `Revert` implements `journal.CacheEntry`.
func (ce *AddLogChange) Revert() {
	sdb := ce.sdb
	logs := sdb.logs[ce.txHash]
	if len(logs) == 1 {
		delete(sdb.logs, ce.txHash)
	} else {
		sdb.logs[ce.txHash] = sdb.logs[ce.txHash][:len(logs)-1]
	}
	sdb.logSize--
}

// `Clone` implements `journal.CacheEntry`.
func (ce *AddLogChange) Clone() journal.CacheEntry {
	return &AddLogChange{
		sdb: ce.sdb,
	}
}

// ==============================================================================
// RefundChange
// ==============================================================================

// `Revert` implements `journal.CacheEntry`.
func (ce *RefundChange) Revert() {
	ce.sdb.refund = ce.prev
}

// `Clone` implements `journal.CacheEntry`.
func (ce *RefundChange) Clone() journal.CacheEntry {
	return &RefundChange{
		sdb:  ce.sdb,
		prev: ce.prev,
	}
}
