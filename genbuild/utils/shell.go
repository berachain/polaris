package utils

import (
	"fmt"
	"github.com/magefile/mage/sh"
)

func RunV(cmd string, args []string, flags map[string]string) error {
	return sh.RunV(cmd, append(args, parseFlags(flags)...)...)
}

func parseFlags(flags map[string]string) []string {
	var args []string
	for k, v := range flags {
		args = append(args, fmt.Sprintf("--%s", k))
		if v != "" {
			args = append(args, v)
		}
	}
	return args
}
