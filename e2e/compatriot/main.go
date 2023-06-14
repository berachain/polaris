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

	"github.com/magefile/mage/sh"
)

const CACHED = "./cached.json"
const NONCACHED = "./noncached.json"

func main() {
	setup()

	// make queries and save results to file 1
	Query(CACHED)

	// kill the chain

	// make queries and save results to file 2
	Query(NONCACHED)

	// compare file 1 and file 2
	err := sh.Run("diff", CACHED, NONCACHED)
	if err != nil {
		log.Fatalf("main: An error occurred %v when diffing\n", err)
	}

	// run sanity checks
	if err := sanityCheck(CACHED); err != nil {
		log.Fatalf("main: An error occurred %v when sanity checking cached file\n", err)
	}

	if err := sanityCheck(NONCACHED); err != nil {
		log.Fatalf("main: An error occurred %v when sanity checking non-cached file\n", err)
	}
}
