package crypto_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPrecompile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "eth/crypto")
}
