// SPDX-License-Identifier: Apache-2.0
//

package debug_test

import (
	"testing"

	"pkg.berachain.dev/polaris/lib/errors/debug"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDebug(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "lib/errors/debug")
}

func Hello() error {
	return nil
}

type GoodBye struct{}

func (g GoodBye) GoodBye() error {
	return nil
}

func (g *GoodBye) GoodByePtr() error {
	return nil
}

type GoodBye2 struct{}

func (g GoodBye2) GoodBye2() error {
	return nil
}

var _ = Describe("TestFnName", func() {
	It("should return the name of the function", func() {
		Expect(debug.GetFnName(Hello)).Should(Equal("Hello"))
	})

	It("should return the name of the function for struct functions", func() {
		gb := GoodBye{}
		Expect(debug.GetFnName(gb.GoodBye)).Should(Equal("GoodBye-fm"))
		Expect(debug.GetFnName(gb.GoodByePtr)).Should(Equal("GoodByePtr-fm"))
		Expect(debug.GetFnName(GoodBye2{}.GoodBye2)).Should(Equal("GoodBye2-fm"))
	})
})
