package main

import "github.com/magefile/mage/sh"

var (
	gitClean           = sh.RunCmd("git", "clean", "-fdx")
	gitSubmodule       = sh.RunCmd("git", "submodule")
	gitSumoduleForEach = sh.RunCmd("git", "submodule", "foreach", "--recursive")
)

// RUN git clean -xfd
// RUN git submodule foreach --recursive git clean -xfd
// RUN git reset --hard
// RUN git submodule foreach --recursive git reset --hard
// RUN git submodule update --init --recursive
