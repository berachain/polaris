package main

import (
	"github.com/magefile/mage/mg"
)

type Git mg.Namespace

var (
	gitRecurse         = RunCmdV("git", "submodule", "update", "--recursive", "--remote")
	gitResetSubmodules = RunCmdV("git", "submodule", "foreach", "--recursive", "git", "reset", "--hard", "HEAD")
)

// ResetSubs runs `git submodule update --recursive --remote` and
// `git submodule foreach --recursive git reset --hard HEAD`.
func (g Git) ResetSubs() error {
	cmds := [](func(...string) error){gitResetSubmodules, gitRecurse}
	for _, cmd := range cmds {
		if err := cmd(); err != nil {
			return err
		}
	}
	return nil
}
