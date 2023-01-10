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

import (
	"fmt"
	"os"

	mi "github.com/berachain/stargazer/build/mage/internal"
)

var (
	protoGenDockerPath = "./build/docker/proto-gen.Dockerfile"
	// TODO: remove once https://github.com/cosmos/cosmos-sdk/pull/13960 is merged
	protoImageName    = "berachain-proto"
	protoImageVersion = "0.11.2"
	protoDir          = "proto"
)

func dockerRunProtoImage(pwd string) func(args ...string) error {
	return mi.RunCmdV("docker",
		"run", "--rm", "-v", pwd+":/workspace",
		"--workdir", "/workspace",
		protoImageName+":"+protoImageVersion)
}

// Run all proto related targets.
func Proto() error {
	cmds := []func() error{ProtoFormat, ProtoLint, ProtoGen}
	for _, cmd := range cmds {
		if err := cmd(); err != nil {
			return err
		}
	}
	return nil
}

// Generate protobuf source files.
func ProtoGen() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	if err = ProtoDockerBuild(); err != nil {
		return err
	}

	return dockerRunProtoImage(dir)(
		"sh", "./build/scripts/proto/proto_generate.sh",
	)
}

// Check that the generated protobuf source files are up to date.
func ProtoGenCheck() error {
	if err := ProtoGen(); err != nil {
		return err
	}
	if err := gitDiff(); err != nil {
		return fmt.Errorf("generated files are out of date: %w", err)
	}
	return nil
}

// Format .proto files.
func ProtoFormat() error {
	return bufWrapper(bufFormat)
}

// Lint .proto files.
func ProtoLint() error {
	return bufWrapper(bufLint)
}

// Build the proto-gen docker image.
func ProtoDockerBuild() error {
	return dockerBuild(
		"--pull",
		"-f", protoGenDockerPath,
		"-t", protoImageName+":"+protoImageVersion,
		"build/docker")
}

// Wraps buf commands with the proper directory change.
func bufWrapper(bufFunc func(args ...string) error) error {
	rootCwd, _ := os.Getwd()
	// Change to the directory where the *.proto's are.
	if err := os.Chdir(protoDir); err != nil {
		return err
	}
	// Run the buf command.
	if err := bufFunc(); err != nil {
		return err
	}
	// Go back to the starting directory.
	if err := os.Chdir(rootCwd); err != nil {
		return err
	}
	return nil
}
