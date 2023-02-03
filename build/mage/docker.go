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

package mage

import mi "github.com/berachain/stargazer/build/mage/internal"

var (
	// Commands.
	dockerBuild = mi.RunCmdV("docker", "build", "--rm=false")

	// Variables.
	baseDockerPath         = "./build/docker/"
	beradDockerPath        = baseDockerPath + "berad.Dockerfile"
	jsonrpcDockerPath      = "./jsonrpc/Dockerfile"
	imageName              = "berachain-node"
	testImageVersion       = "e2e-test-dev"
	goVersion              = "1.19.4"
	debianStaticImage      = "gcr.io/distroless/static-debian11"
	golangAlpine           = "golang:1.19-alpine3.17"
	precompileContractsDir = ""
)

// Build a lightweight docker image for berad.
func DockerGen() error {
	return dockerBuildBeradWith(goVersion, debianStaticImage, version)
}

// Build a debuggable docker image for berad.
func DockerDebug() error {
	return dockerBuildBeradWith(goVersion, golangAlpine, version)
}

// Build a docker image for berad with e2e test dependencies.
func DockerE2eTest() error {
	return dockerBuildBeradWith(goVersion, golangAlpine, testImageVersion)
}

func DockerBuildJSONRPCServer() error {
	return dockerBuild(
		"-f", jsonrpcDockerPath,
		"--build-arg", "GO_VERSION="+goVersion,
		"--build-arg", "RUNNER_IMAGE="+debianStaticImage,
		"-t", "jsonrpc-server",
		".",
	)
}

// Build a docker image for berad with the supplied arguments.
func dockerBuildBeradWith(goVersion, runnerImage, imageVersion string) error {
	return dockerBuild(
		"--build-arg", "GO_VERSION="+goVersion,
		"--build-arg", "RUNNER_IMAGE="+runnerImage,
		"--build-arg", "PRECOMPILE_CONTRACTS_DIR="+precompileContractsDir,
		"-f", beradDockerPath,
		"-t", imageName+":"+imageVersion,
		".",
	)
}
