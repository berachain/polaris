// SPDX-License-Identifier: Apache-2.0
//

package encoding

import (
	"crypto/sha256"
	"encoding/binary"

	"github.com/holiman/uint256"
)

// globalNonce is used as a counter across generation of all salts.
var globalNonce uint64 = 0

// UniqueDeterministicSalt returns a unique and deterministic salt for the given input bytes. Uses
// sha256 to hash the input bytes and the global nonce and returns the salt as a *uint256.Int.
func UniqueDeterminsticSalt(input []byte) *uint256.Int {
	h := sha256.New()
	h.Write(input)
	binary.Write(h, binary.BigEndian, globalNonce)
	globalNonce++
	return new(uint256.Int).SetBytes(h.Sum(nil))
}
