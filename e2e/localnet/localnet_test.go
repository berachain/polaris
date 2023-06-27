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
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	dt "github.com/ory/dockertest"
)

func TestLocalnet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "testing:integration")
}

var _ = Describe("Fixture", func() {
	var (
		c   *LocalnetClient
		ctx context.Context
	)

	BeforeAll(func() {
		ctx = context.Background()
		pool, err := dt.NewPool("")
		Expect(err).ToNot(HaveOccurred())

		err = pool.Client.Ping()
		Expect(err).ToNot(HaveOccurred())

		err = buildDefault(pool)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		if c != nil {
			Expect(c.Stop()).To(Succeed())
		}
	})

	It("should create a container", func() {
		c, err := NewLocalnetClient(ctx, "polard/localnet:v0.0.0", "something", "8545/tcp", "8546/tcp")
		Expect(err).ToNot(HaveOccurred())
		Expect(c).ToNot(BeNil())

		err = c.Start(context.Background())
		Expect(err).ToNot(HaveOccurred())
	})
})
