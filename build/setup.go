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

//nolint:forbidigo // its okay.
package main

import (
	"errors"
	"fmt"

	"github.com/carolynvs/magex/pkg"
	"github.com/magefile/mage/sh"
)

var (
	// tools.
	buf              = "github.com/bufbuild/buf/cmd/buf"
	gosec            = "github.com/cosmos/gosec/v2/cmd/gosec"
	golangcilint     = "github.com/golangci/golangci-lint/cmd/golangci-lint"
	addlicense       = "github.com/google/addlicense"
	goimportsReviser = "github.com/incu6us/goimports-reviser/v3"
	moq              = "github.com/matryer/moq"
	ginkgo           = "github.com/onsi/ginkgo/v2/ginkgo"
	golines          = "github.com/segmentio/golines"
	goimports        = "golang.org/x/tools/cmd/goimports"
	// gopls            = "golang.org/x/tools/gopls".
	rlpgen = "github.com/ethereum/go-ethereum/rlp/rlpgen"

	allTools = []string{buf, gosec, golangcilint, addlicense,
		goimportsReviser, moq, ginkgo, golines, goimports, rlpgen /*gopls*/}
)

// Setup runs the setup script for the current OS.
func main() {
	var err error

	// Ensure Mage is installed and available on the $PATH.
	if err = pkg.EnsureMage(""); err != nil {
		panic(err)
	}

	// Coming soon
	// // Run the setup script for the current OS.
	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	err = setupMac()
	// case "linux":
	// 	err = setupLinux()
	// default:
	// 	err = fmt.Errorf("ngmi unsupported OS")
	// }

	if err = setupGoDeps(); err != nil {
		panic(err)
	}
}

func setupGoDeps() error {
	for _, tool := range allTools {
		fmt.Println("Installing ", tool)
		if err := sh.RunCmd("go", "install", "-mod=readonly", tool); err() != nil {
			return errors.New("failed to install " + tool + ": " + err().Error())
		}
	}
	fmt.Println("\n==============================================================")
	fmt.Println("Tools installed successful! Ensure $GOPATH/bin is on your $PATH!")
	fmt.Println("==============================================================")
	return nil
}

// // setupMac runs the setup script for macOS.
// func setupMac() error {
// 	return fmt.Errorf("mac setup coming soon")
// }

// // setupLinux runs the setup script for Linux.
// func setupLinux() error {
// 	return fmt.Errorf("linux setup coming soon")
// }
