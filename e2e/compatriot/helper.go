// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// setDirectory sets the directory to Polaris root
func setDirectory() error {
	if err := os.Chdir("../.."); err != nil {
		return fmt.Errorf("helper: An error occurred %v when changing directory\n", err)
	}
	return nil
}

// startNode starts the node and returns a pointer to the command
// for reference so that the node can be terminated later
func startNode(new, verbose bool) (*exec.Cmd, error) {
	var cmd *exec.Cmd
	if new {
		cmd = exec.Command("mage", "start")
	} else {
		cmd = exec.Command("./bin/polard", "start", "--home", "./.tmp/polard")
	}
	if verbose {
		cmd.Stdout = os.Stdout
	}
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("helper: An error occurred %v when starting chain\n", err)
	}
	wait()
	return cmd, nil
}

// stopNode stops the node and kills all subprocesses
func stopNode(nodeCmd *exec.Cmd) error {
	ps := exec.Command("ps")
	grep := exec.Command("grep", "./bin/polard")

	pipe, _ := ps.StdoutPipe()
	defer pipe.Close()

	grep.Stdin = pipe
	if err := ps.Start(); err != nil {
		return fmt.Errorf("helper: An error occurred %v when running ps\n", err)
	}
	output, err := grep.Output()
	if err != nil {
		return fmt.Errorf("helper: An error occurred %v when running grep\n", err)
	}

	// kill the subprocess
	exec.Command("kill", "-9", string(strings.Fields(string(output))[0])).Run()
	if err := nodeCmd.Process.Kill(); err != nil {
		return fmt.Errorf("helper: An error occurred %v when killing the process\n", err)
	}

	return nil
}

// wait waits for 5 seconds
// TODO: resolve hacky fix to wait for chain endpoints to be setup correctly
func wait() {
	time.Sleep(5 * time.Second)
}

// diff compares two files and saves the result to diff.txt
// TODO: upgrade this as to not store all of the files in memory
func diff(file1, file2 string) error {
	// compare file 1 and file 2
	diff := exec.Command("diff", cached, noncached)
	out, err := os.Create(diffFile)
	if err != nil {
		return fmt.Errorf("main: An error occurred %v when creating diff file\n", err)
	}
	defer out.Close()
	diff.Stdout = out
	if err := diff.Run(); err != nil {
		switch err.(type) {
		case *exec.ExitError:
			// this is just an exit code error, no worries
			// do nothing
		default: //couldn't run diff
			return fmt.Errorf("main: An error occurred %v when diffing\n", err)
		}
	}

	return nil
}
