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
	"os"
	"testing"

	dt "github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLocalnet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "testing:integration")
}

var _ = Describe("Fixture", func() {
	var (
		c          *LocalnetClient
		ctx        context.Context
		localImage string
	)

	BeforeEach(func() {
		ctx = context.Background()

		fmt.Println("Started 1")

		pool, err := dt.NewPool("")
		Expect(err).ToNot(HaveOccurred())

		err = pool.Client.Ping()
		Expect(err).ToNot(HaveOccurred())

		baseBuildArgs := []docker.BuildArg{
			{
				Name:  "GO_VERSION",
				Value: "1.20.4",
			},
			{
				Name:  "PRECOMPILE_CONTRACTS_DIR",
				Value: "./contracts",
			},
			{
				Name:  "GOOS",
				Value: "linux",
			},
			{
				Name:  "GOARCH",
				Value: "arm64",
			},
		}
		baseImageName := "polard/base:v0.0.0"

		baseBuildOpts := docker.BuildImageOptions{
			Name:         baseImageName,
			ContextDir:   "../../",
			Dockerfile:   "./cosmos/docker/base.Dockerfile",
			BuildArgs:    baseBuildArgs,
			OutputStream: os.Stdout,
		}

		err = pool.Client.BuildImage(baseBuildOpts)
		Expect(err).ToNot(HaveOccurred())

		fmt.Println("Started 2")

		localBuildArgs := []docker.BuildArg{
			{
				Name:  "BASE_IMAGE",
				Value: baseImageName,
			},
		}
		localImage = "polard/localnet:v0.0.0"
		localBuildOpts := docker.BuildImageOptions{
			Name:         localImage,
			ContextDir:   "./",
			Dockerfile:   "Dockerfile",
			BuildArgs:    localBuildArgs,
			OutputStream: os.Stdout,
		}
		err = pool.Client.BuildImage(localBuildOpts)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		if c != nil {
			Expect(c.Stop()).To(Succeed())
		}
	})

	It("should create a container", func() {
		fmt.Println("Started 4")
		c, err := NewLocalnetClient(ctx, localImage, "something", "localhost:8545", "localhost:8546")
		Expect(err).ToNot(HaveOccurred())
		Expect(c).ToNot(BeNil())

		err = c.Start(context.Background())
		Expect(err).ToNot(HaveOccurred())
	})
})
