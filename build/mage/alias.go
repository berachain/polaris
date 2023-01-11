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
	mi "github.com/berachain/stargazer/build/mage/internal"
)

// Go Aliases.
var (
	goInstall       = mi.RunCmdV("go", "install", "-mod=readonly")
	goBuild         = mi.RunCmdV("go", "build", "-mod=readonly")
	goRun           = mi.RunCmdV("go", "run")
	goTest          = mi.RunCmdV("go", "test", "-mod=readonly")
	ginkgoTest      = mi.RunCmdV("ginkgo", "-r", "--randomize-all", "--fail-on-pending", "-trace")
	ginkgoCoverArgs = []string{"--junit-report", "out.xml", "--cover",
		"--coverprofile", "testUnitCover.txt", "--covermode", "atomic"}
	goGenerate  = mi.RunCmdV("go", "generate")
	goModVerify = mi.RunCmdV("go", "mod", "verify")
	goModTidy   = mi.RunCmdV("go", "mod", "tidy")
)

// Forge Aliases.
var (
	forgeBuild = mi.RunCmdV("forge", "build", "--extra-output-files", "abi")
	forgeClean = mi.RunCmdV("forge", "clean")
	forgeTest  = mi.RunCmdV("forge", "test")
	forgeFmt   = mi.RunCmdV("forge", "fmt")
)

// Buf Aliases.
var (
	bufRepo = "github.com/bufbuild/buf/cmd/buf"
	// bufBuild  = mi.RunCmdV("go", "run", bufRepo, "build").
	bufFormat = mi.RunCmdV("go", "run", bufRepo, "format", "-w")
	bufLint   = mi.RunCmdV("go", "run", bufRepo, "lint", "--error-format=json")
)
