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
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

const CACHED = "./e2e/compatriot/cached.json"
const NONCACHED = "./e2e/compatriot/noncached.json"
const diffFile = "./e2e/compatriot/diff.txt"

func main() {
	// set the directory
	if err := os.Chdir("../.."); err != nil {
		log.Fatalf("main: An error occurred %v when changing directory\n", err)
	}

	// start the chain
	startChain := exec.Command("./cosmos/init.sh")
	startChain.Stdout = os.Stdout
	if err := startChain.Start(); err != nil {
		log.Fatalf("main: An error occurred %v when starting chain\n", err)
	}
	// setup()

	time.Sleep(10 * time.Second) // hacky fix to wait for chain endpoints to be setup correctly

	// make queries and save results to file 1
	Query(CACHED)

	// kill the chain
	ps := exec.Command("ps")
	grep := exec.Command("grep", "./bin/polard")

	pipe, _ := ps.StdoutPipe()
	defer pipe.Close()

	grep.Stdin = pipe
	ps.Start()
	output, _ := grep.Output()

	// kill the subprocess
	exec.Command("kill", string(strings.Fields(string(output))[0])).Run()

	if err := startChain.Process.Kill(); err != nil {
		log.Fatalf("main: An error occurred %v when killing the program\n", err)
	}

	// restart the chain
	restartChain := exec.Command("./bin/polard", "start")
	if err := restartChain.Start(); err != nil {
		log.Fatalf("main: An error occurred %v when restarting chain\n", err)
	}

	// make queries and save results to file 2
	Query(NONCACHED)

	if err := restartChain.Process.Kill(); err != nil {
		log.Fatalf("main: An error occurred %v when killing the restarted chain\n", err)
	}

	// compare file 1 and file 2
	diff := exec.Command("diff", CACHED, NONCACHED)
	diff.Stdout = os.NewFile(3, diffFile)
	if err := diff.Run(); err != nil {
		switch err.(type) {
		case *exec.ExitError:
			// this is just an exit code error, no worries
			// do nothing
		default: //couldnt run diff
			log.Fatalf("main: An error occurred %v when diffing\n", err)
		}
	}

	// run sanity checks
	if err := sanityCheck(CACHED); err != nil {
		log.Fatalf("main: An error occurred %v when sanity checking cached file\n", err)
	}

	if err := sanityCheck(NONCACHED); err != nil {
		log.Fatalf("main: An error occurred %v when sanity checking non-cached file\n", err)
	}
}
