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

package localnet

import (
	"context"
	gotesting "testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPolarisChain(t *gotesting.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "testing:integration")
}

var _ = Describe("Fixture", func() {
	var (
		// ctx context.Context
		c *LocalnetClient
	)

	// BeforeEach(func() {
	// 	ctx = context.Background()
	// })

	AfterEach(func() {
		if c != nil {
			Expect(c.Stop()).To(Succeed())
		}
	})

	It("should create a container", func() {
		var err error
		c, err := NewLocalnetClient(context.Background(), "something", "localhost:8545", "localhost:8546")
		Expect(err).ToNot(HaveOccurred())
		Expect(c).ToNot(BeNil())

		err = c.container.Start(context.Background())
		Expect(err).ToNot(HaveOccurred())
	})
})
