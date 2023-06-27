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
	"fmt"
	"testing"

	tc "github.com/testcontainers/testcontainers-go"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLocalnet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "testing:integration")
}

var (
	contractsDir = "./contracts"
	goVersion    = "1.20.4"
	goOS         = "linux"
	goArch       = "arm64"
)

var _ = Describe("Fixture", func() {
	var (
		c         *LocalnetClient
		ctx       context.Context
		container tc.Container
	)

	BeforeEach(func() {
		ctx = context.Background()

		fmt.Println("Started 1")
		baseImage := tc.ContainerRequest{
			FromDockerfile: tc.FromDockerfile{
				Context:    "../../",
				Dockerfile: "./cosmos/docker/base.Dockerfile",
				BuildArgs: map[string]*string{
					"GO_VERSION":               &goVersion,
					"PRECOMPILE_CONTRACTS_DIR": &contractsDir,
					"GOOS":                     &goOS,
					"GOARCH":                   &goArch,
				},
			},
		}

		_, err := tc.GenericContainer(ctx,
			tc.GenericContainerRequest{
				ContainerRequest: baseImage,
				Started:          false,
				Reuse:            false,
			})
		Expect(err).ToNot(HaveOccurred())

		fmt.Println("Started 2")
		localnetImage := tc.ContainerRequest{
			FromDockerfile: tc.FromDockerfile{
				Context:    "./",
				Dockerfile: "Dockerfile",
			},
		}
		localnetContainer, err := tc.GenericContainer(ctx, tc.GenericContainerRequest{
			ContainerRequest: localnetImage,
			Started:          false,
			Reuse:            false,
		})
		Expect(err).ToNot(HaveOccurred())
		container = localnetContainer
		fmt.Println("Started 3")
	})

	AfterEach(func() {
		if c != nil {
			Expect(c.Stop()).To(Succeed())
		}
	})

	It("should create a container", func() {
		var err error
		name, err := container.Name(ctx)
		Expect(err).ToNot(HaveOccurred())
		fmt.Println("name: ", name)
		fmt.Println("Started 4")
		c, err := NewLocalnetClient(ctx, name[1:], "something", "localhost:8545", "localhost:8546")
		Expect(err).ToNot(HaveOccurred())
		Expect(c).ToNot(BeNil())

		err = c.Start(context.Background())
		Expect(err).ToNot(HaveOccurred())
	})
})
