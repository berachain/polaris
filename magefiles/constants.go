package main

import "os"

/* -------------------------------------------------------------------------- */
/*                                  Commands                                  */
/* -------------------------------------------------------------------------- */

var (
	// Forge
	forgeBuild = RunCmdV("forge", "build", "--extra-output-files", "bin", "--extra-output-files", "abi", "--silent")
	forgeClean = RunCmdV("forge", "clean")
	forgeTest  = RunCmdV("forge", "test")
	forgeFmt   = RunCmdV("forge", "fmt")

	// Docker
	dockerBuild  = RunCmdV("docker", "build", "--rm=false")
	dockerBuildX = RunCmdV("docker", "buildx", "build", "--rm=false")
	dockerRun    = RunCmdV("docker", "run")

	// Buf
	bufCommand = RunCmdV("buf")

	// Testing
	goTest     = RunCmdV("go", "test", "-mod=readonly")
	ginkgoTest = RunCmdV("ginkgo", "-r", "--randomize-all", "--fail-on-pending", "-trace")

	// Toolchain
	goInstall  = RunCmdV("go", "install", "-mod=readonly")
	goBuild    = RunCmdV("go", "build", "-mod=readonly")
	goRun      = RunCmdV("go", "run")
	goGenerate = RunCmdV("go", "generate")
	goModTidy  = RunCmdV("go", "mod", "tidy")
	goWorkSync = RunCmdV("go", "work", "sync")
	gitDiff    = RunCmdV("git", "diff", "--stat", "--exit-code", ".",
		"':(exclude)*.mod' ':(exclude)*.sum'")
)

/* -------------------------------------------------------------------------- */
/*                             Packages & Modules                             */
/* -------------------------------------------------------------------------- */
var (
	repoModuleDirs = readGoModulesFromGoWork("go.work")
)

const (
	cosmosSDK  = "github.com/cosmos/cosmos-sdk"
	moq        = "github.com/matryer/moq"
	golangCi   = "github.com/golangci/golangci-lint/cmd/golangci-lint"
	golines    = "github.com/segmentio/golines"
	gosec      = "github.com/securego/gosec/v2/cmd/gosec"
	addlicense = "github.com/google/addlicense"
)

/* -------------------------------------------------------------------------- */
/*                                   Docker                                   */
/* -------------------------------------------------------------------------- */

var (
	baseImageVersion  = "polard/base:v0.0.0"
	protoImageName    = "ghcr.io/cosmos/proto-builder"
	protoImageVersion = "0.13.5"
	protoDir          = "cosmos/proto"
)

/* -------------------------------------------------------------------------- */
/*                                 Directories                                */
/* -------------------------------------------------------------------------- */

const (
	outdir             = "./bin"
	baseHiveDockerPath = "./e2e/hive/"
)

var (
	hiveClone      = os.Getenv("GOPATH") + "/src/"
	clonePath      = hiveClone + ".hive-e2e/"
	simulatorsPath = clonePath + "simulators/polaris/"
	clientsPath    = clonePath + "clients/polard/"
)
