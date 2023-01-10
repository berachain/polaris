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
	"strings"

	"github.com/magefile/mage/sh"
)

var (
	sdkRepo        = "github.com/cosmos/cosmos-sdk"
	version        = "0.0.0"
	commit, _      = sh.Output("git", "log", "-1", "--format='%H'")
	defaultDB      = "pebbledb"
	ledgerEnabled  = true
	appName        = "berachain"
	executableName = "berad"
)

// generateCmdToBuild returns the command to build a given command.
func generateCmdToBuild(cmd string) string {
	return "./cmd/" + cmd
}

// generateOutDirectory returns the output directory for a given command.
func generateOutDirectory(cmd string) string {
	return outdir + "/" + cmd
}

// generateBuildTags returns the build tags to be used when building the binary.
func generateBuildTags() string {
	tags := []string{defaultDB}
	if ledgerEnabled {
		tags = append(tags, "ledger")
	}
	return "-tags='" + strings.Join(tags, " ") + "'"
}

// generateLinkerFlags returns the linker flags to be used when building the binary.
func generateLinkerFlags(production, statically bool) string {
	baseFlags := []string{
		"-X ", sdkRepo + "/version.Name=" + executableName,
		" -X ", sdkRepo + "/version.AppName=" + appName,
		" -X ", sdkRepo + "/version.Version=" + version,
		" -X ", sdkRepo + "/version.Commit=" + commit,
		// TODO: Refactor versioning more broadly.
		// " \"-X " + sdkRepo + "/version.BuildTags=" + strings.Join(generateBuildTags(), ",") +
		" -X ", sdkRepo + "/version.DBBackend=" + defaultDB,
	}

	if production {
		baseFlags = append(baseFlags, "-w", "-s")
	}

	if statically {
		baseFlags = append(
			baseFlags,
			"-linkmode=external",
			"-extldflags",
			"\"-Wl,-z,muldefs -static\"",
		)
	}

	return "-ldflags=" + strings.Join(baseFlags, " ")
}
