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

import (
	"github.com/berachain/stargazer/lib/common"
)

const (
	keyPrefixCode byte = iota
	keyPrefixHash
	keyPrefixStorage
)

// NOTE: we use copy to build keys for max performance: https://github.com/golang/go/issues/55905

// AddressStoragePrefix returns a prefix to iterate over a given account storage.
func AddressStoragePrefix(address common.Address) []byte {
	bz := make([]byte, 1+common.AddressLength)
	copy(bz, []byte{keyPrefixStorage})
	copy(bz[1:], address[:])
	return bz
}

// `StateKeyFor` defines the full key under which an account state is stored.
func StateKeyFor(address common.Address, key common.Hash) []byte {
	bz := make([]byte, 1+common.AddressLength+common.HashLength)
	copy(bz, []byte{keyPrefixStorage})
	copy(bz[1:], address[:])
	copy(bz[1+common.AddressLength:], key[:])
	return bz
}

// `CodeHashKeyFor` defines the full key under which an addreses codehash is stored.
func CodeHashKeyFor(address common.Address) []byte {
	bz := make([]byte, 1+common.AddressLength)
	copy(bz, []byte{keyPrefixCode})
	copy(bz[1:], address[:])
	return bz
}

// `CodeKeyFor` defines the full key for which a codehash's corresponding code is stored.
func CodeKeyFor(codeHash common.Hash) []byte {
	bz := make([]byte, 1+common.HashLength)
	copy(bz, []byte{keyPrefixCode})
	copy(bz[1:], codeHash[:])
	return bz
}
