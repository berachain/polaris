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
)

// useful files
const (
	cached    = "./e2e/compatriot/cached.json"
	noncached = "./e2e/compatriot/noncached.json"
	diffFile  = "./e2e/compatriot/diff.txt"
)

func main() {
	flags := os.Args[1:]

	var verbose bool
	// set the verbose flag
	if len(flags) > 0 {
		switch flags[0] {
		case "-v":
			verbose = true
		}
	}

	// set the directory
	if err := setDirectory(); err != nil {
		log.Fatalf("main: An error occurred %v when setting directory\n", err)
	}

	var node *exec.Cmd

	// TODO: improve so it tries again
	// catch panics
	// requires graceful shutdown of the node on process termination (localhost:8545 is still being run)
	defer func() {
		if err := recover(); err != nil {
			// kill the chain
			if err := stopNode(node); err != nil {
				log.Fatalf("main: An error occurred %v when stopping the node\n", err)
			}
			log.Fatalf("main: An error occurred %v and was caught\n", err)
		}
	}()

	// start the chain
	node, err := startNode(true, verbose)
	if err != nil {
		log.Fatalf("main: An error occurred %v when starting the node\n", err)
	}

	// chain setup
	if err := setup(); err != nil {
		// kill the chain
		if err := stopNode(node); err != nil {
			log.Fatalf("main: An error occurred %v when stopping the node\n", err)
		}
		log.Fatalf("main: An error occurred %v when setting up chain\n", err)
	}

	// make queries and save results to file 1
	if err := query(cached); err != nil {
		// kill the chain
		if err := stopNode(node); err != nil {
			log.Fatalf("main: An error occurred %v when stopping the node\n", err)
		}
		log.Fatalf("main: An error occurred %v when querying chain\n", err)
	}

	// kill the chain
	if err := stopNode(node); err != nil {
		log.Fatalf("main: An error occurred %v when stopping the node\n", err)
	}

	// restart the chain
	node, err = startNode(false, verbose)
	if err != nil {
		log.Fatalf("main: An error occurred %v when restarting the node\n", err)
	}

	// make queries and save results to file 2
	query(noncached)

	// kill the chain again
	if err := stopNode(node); err != nil {
		log.Fatalf("main: An error occurred %v when stopping the node\n", err)
	}

	// compare the two files
	if err := diff(cached, noncached); err != nil {
		log.Fatalf("main: An error occurred %v when diffing files\n", err)
	}
}
