package ds_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDS(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "lib/ds")
}
