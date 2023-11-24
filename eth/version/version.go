// Package version implements reading of build version information.
package version

// Copyright 2022 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/ethereum/go-ethereum/params"
)

const (
	govcsTimeLayout = "2006-01-02T15:04:05Z"
	ourTimeLayout   = "20060102"
	ourPath         = "github.com/berachain/polaris/eth" // Path to our module
)

// These variables are set at build-time by the linker when the build is
// done by build/ci.go.
var gitCommit, gitDate string

// VCSInfo represents the git repository state.
type VCSInfo struct {
	Commit string // head commit hash
	Date   string // commit time in YYYYMMDD format
	Dirty  bool
}

// VCS returns version control information of the current executable.
func VCS() (VCSInfo, bool) {
	if gitCommit != "" {
		// Use information set by the build script if present.
		return VCSInfo{Commit: gitCommit, Date: gitDate}, true
	}
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		if buildInfo.Main.Path == ourPath {
			return buildInfoVCS(buildInfo)
		}
	}
	return VCSInfo{}, false
}

// ClientName creates a software name/version identifier according to common
// conventions in the Ethereum p2p network.
func ClientName(clientIdentifier string) string {
	git, _ := VCS()
	return fmt.Sprintf("%s/v%v/%v-%v/%v",
		cases.Title(language.English).String(clientIdentifier),
		params.VersionWithCommit(git.Commit, git.Date),
		runtime.GOOS, runtime.GOARCH,
		runtime.Version(),
	)
}

// Info returns build and platform information about the current binary.
//
// If the package that is currently executing is a prefixed by our go-ethereum
// module path, it will print out commit and date VCS information. Otherwise,
// it will assume it's imported by a third-party and will return the imported
// version and whether it was replaced by another module.
func Info() (string, string) {
	var ok bool
	var status VCSInfo
	var vcs string
	version := params.VersionWithMeta
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return version, ""
	}
	version = versionInfo(buildInfo)
	if status, ok = VCS(); ok {
		modified := ""
		if status.Dirty {
			modified = " (dirty)"
		}
		commit := status.Commit
		if len(commit) > 8 { //nolint:gomnd // its okay.
			commit = commit[:8]
		}
		vcs = commit + "-" + status.Date + modified
	}
	return version, vcs
}

// versionInfo returns version information for the currently executing
// implementation.
//
// Depending on how the code is instantiated, it returns different amounts of
// information. If it is unable to determine which module is related to our
// package it falls back to the hardcoded values in the params package.
func versionInfo(info *debug.BuildInfo) string {
	// If the main package is from our repo, prefix version with "geth".
	if strings.HasPrefix(info.Path, ourPath) {
		return fmt.Sprintf("geth %s", info.Main.Version)
	}
	// Not our main package, so explicitly print out the module path and
	// version.
	var version string
	if info.Main.Path != "" && info.Main.Version != "" {
		// These can be empty when invoked with "go run".
		version = fmt.Sprintf("%s@%s ", info.Main.Path, info.Main.Version)
	}
	mod := findModule(info, ourPath)
	if mod == nil {
		// If our module path wasn't imported, it's unclear which
		// version of our code they are running. Fallback to hardcoded
		// version.
		return version + fmt.Sprintf("geth %s", params.VersionWithMeta)
	}
	// Our package is a dependency for the main module. Return path and
	// version data for both.
	version += fmt.Sprintf("%s@%s", mod.Path, mod.Version)
	if mod.Replace != nil {
		// If our package was replaced by something else, also note that.
		version += fmt.Sprintf(" (replaced by %s@%s)", mod.Replace.Path, mod.Replace.Version)
	}
	return version
}

// findModule returns the module at path.
func findModule(info *debug.BuildInfo, path string) *debug.Module {
	if info.Path == ourPath {
		return &info.Main
	}
	for _, mod := range info.Deps {
		if mod.Path == path {
			return mod
		}
	}
	return nil
}

// buildInfoVCS returns VCS information of the build.
func buildInfoVCS(info *debug.BuildInfo) (VCSInfo, bool) {
	var s VCSInfo
	var ok bool
	for _, v := range info.Settings {
		switch v.Key {
		case "vcs.revision":
			s.Commit = v.Value
		case "vcs.modified":
			if v.Value == "true" {
				s.Dirty = true
			}
		case "vcs.time":
			t, err := time.Parse(govcsTimeLayout, v.Value)
			if err == nil {
				s.Date = t.Format(ourTimeLayout)
			}
		}
	}
	if s.Commit != "" && s.Date != "" {
		ok = true
	}
	return s, ok
}
