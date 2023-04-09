package cmd

import "pkg.berachain.dev/polaris/magefiles/cmd/utils"

var (
	// Commands.
	goInstall  = utils.RunCmdV("go", "install", "-mod=readonly")
	goBuild    = utils.RunCmdV("go", "build", "-mod=readonly")
	goRun      = utils.RunCmdV("go", "run")
	goGenerate = utils.RunCmdV("go", "generate")
	goModTidy  = utils.RunCmdV("go", "mod", "tidy")
	goWorkSync = utils.RunCmdV("go", "work", "sync")
	goTest     = utils.RunCmdV("go", "test", "-mod=readonly")
	ginkgoTest = utils.RunCmdV("ginkgo", "-r", "--randomize-all", "--fail-on-pending", "-trace")
)

// TestUnit runs all unit tests in the given path.
func TestUnit(path string) error {
	return ginkgoTest("--skip", ".*integration.*", "./"+path+"/...")
}

// TestIntegration runs all integration tests in the given path.
func TestIntegration(path string) error {
	return ginkgoTest("--skip", ".*unit.*", "./"+path+"/...")
}
