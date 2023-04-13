// SPDX-License-Identifier: Apache-2.0
//

package encoding_test

import (
	"fmt"

	"github.com/holiman/uint256"

	enclib "pkg.berachain.dev/polaris/lib/encoding"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Salt", func() {
	It("should be unique and deterministic", func() {
		salts := make(map[uint256.Int]struct{})
		orderedSalts := make([]uint256.Int, 20000)

		for i := 0; i < 10000; i++ {
			salt := enclib.UniqueDeterminsticSalt([]byte("test"))
			_, found := salts[*salt]
			Expect(found).To(BeFalse())
			salts[*salt] = struct{}{}
			orderedSalts[i] = *salt
		}

		for i := 0; i < 10000; i++ {
			salt := enclib.UniqueDeterminsticSalt([]byte(fmt.Sprintf("test%d", i)))
			_, found := salts[*salt]
			Expect(found).To(BeFalse())
			salts[*salt] = struct{}{}
			orderedSalts[i+10000] = *salt
		}
	})
})
