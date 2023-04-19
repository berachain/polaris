// SPDX-License-Identifier: Apache-2.0
//

package encoding

import (
	"crypto/sha256"
	"encoding/binary"

	"github.com/holiman/uint256"
	"golang.org/x/crypto/sha3"
)

// TODO: get rid of this. global nonce must be on disk and consensus synced.

// globalNonce is used as a counter across generation of all salts and hashes.
var globalNonce uint64

// UniqueDeterministicSalt returns a unique and deterministic salt for the given input bytes. Uses
// sha256 to hash the input bytes and the global nonce and returns the salt as a *uint256.Int.
func UniqueDeterminsticSalt(input []byte) *uint256.Int {
	h := sha256.New()
	h.Write(input)
	err := binary.Write(h, binary.BigEndian, globalNonce)
	if err != nil {
		panic(err)
	}
	globalNonce++
	return new(uint256.Int).SetBytes(h.Sum(nil))
}

// UniqueDeterministicHash returns a unique and deterministic hash of the input bytes. Uses sha1 to
// hash the input bytes and the global nonce and returns the hash as a []byte.
func UniqueDeterministicHash(input []byte) []byte {
	h := sha3.New256()
	h.Write(input)
	err := binary.Write(h, binary.BigEndian, globalNonce)
	if err != nil {
		panic(err)
	}
	globalNonce++
	return h.Sum(nil)
}
