package utils

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestParseFlags(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "genbuild/utils")
}

var _ = Describe("utils", func() {
	It("should parse flag map into []string", func() {
		args := parseFlags(map[string]string{
			"home":          "~/.polard",
			"overwrite":     "",
			"chain-id":      "polaris-2061",
			"default-denom": "abera",
		})
		expectedArgs := []string{
			"--home", "~/.polard",
			"--overwrite",
			"--chain-id", "polaris-2061",
			"--default-denom", "abera",
		}
		Expect(args).To(Equal(expectedArgs))
	})
})
