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

import (
	"github.com/berachain/stargazer/eth/common"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("StorageKeyFor", func() {
	It("returns a prefix to iterate over a given account storage", func() {
		address := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
		prefix := StorageKeyFor(address)
		Expect(prefix).To(HaveLen(1 + common.AddressLength))
		Expect(prefix[0]).To(Equal(keyPrefixStorage))
		Expect(prefix[1:]).To(Equal(address.Bytes()))
	})
})

var _ = Describe("SlotKeyFor", func() {
	It("returns a storage key for a given account and storage slot", func() {
		address := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
		slot := common.HexToHash("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef")
		key := SlotKeyFor(address, slot)
		Expect(key).To(HaveLen(1 + common.AddressLength + common.HashLength))
		Expect(key[0]).To(Equal(keyPrefixStorage))
		Expect(key[1 : 1+common.AddressLength]).To(Equal(address.Bytes()))
		Expect(key[1+common.AddressLength:]).To(Equal(slot.Bytes()))
	})
})

var _ = Describe("CodeHashKeyFo or a given account", func() {
	address := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
	key := CodeHashKeyFor(address)
	Expect(key).To(HaveLen(1 + common.AddressLength))
	Expect(key[0]).To(Equal(keyPrefixCode))
	Expect(key[1:]).To(Equal(address.Bytes()))
})

var _ = Describe("CodeKeyFor", func() {
	It("returns a code key for a given account", func() {
		address := common.HexToHash("0x1234567890abcdef1234567890abcdef12345678")
		key := CodeKeyFor(address)
		Expect(key).To(HaveLen(1 + common.HashLength))
		Expect(key[0]).To(Equal(keyPrefixCode))
		Expect(key[1:]).To(Equal(address.Bytes()))
	})
})
