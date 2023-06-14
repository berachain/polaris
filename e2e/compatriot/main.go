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
	"log"
	"os"
	"os/exec"
	"time"
)

const CACHED = "./cached.json"
const NONCACHED = "./noncached.json"

func main() {
	if err := os.Chdir("../.."); err != nil {
		log.Fatalf("main: An error occurred %v when changing directory\n", err)
	}

	startChain := exec.Command("./cosmos/init.sh")
	startChain.Stdout = os.Stdout
	if err := startChain.Start(); err != nil {
		log.Fatalf("main: An error occurred %v when starting chain\n", err)
	}
	// setup()

	print := exec.Command("cat", "magefiles/LICENSE.header")
	if err := print.Run(); err != nil {
		log.Fatalf("main: An error occurred %v when printing\n", err)
	}

	time.Sleep(10 * time.Second)

	// make queries and save results to file 1
	// TODO: figure out how to query the chain and output results after the endpoints are ready
	Query(CACHED)

	// kill the chain
	if err := startChain.Wait(); err != nil {
		log.Fatalf("main: An error occurred %v when waiting for start chain to finish\n", err)
	}

	if err := startChain.Process.Kill(); err != nil {
		log.Fatalf("main: An error occurred %v when killing the program\n", err)
	}

	// restart the chain
	restartChain := exec.Command("./bin/polard", "start", "--home", "$HOMEDIR")
	if err := restartChain.Run(); err != nil {
		log.Fatalf("main: An error occurred %v when restarting chain\n", err)
	}

	// make queries and save results to file 2
	Query(NONCACHED)

	// compare file 1 and file 2
	diff := exec.Command("diff", CACHED, NONCACHED)
	output, err := diff.Output()
	if err != nil {
		log.Fatalf("main: An error occurred %v when diffing\n", err)
	}
	fmt.Println(string(output))

	// run sanity checks
	if err := sanityCheck(CACHED); err != nil {
		log.Fatalf("main: An error occurred %v when sanity checking cached file\n", err)
	}

	if err := sanityCheck(NONCACHED); err != nil {
		log.Fatalf("main: An error occurred %v when sanity checking non-cached file\n", err)
	}
}
