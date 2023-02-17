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

package abi

import (
	"github.com/berachain/stargazer/eth/common"
	"github.com/berachain/stargazer/eth/common/hexutil"
)

// `CompiliedContract` is a contract that has been compiled.
type CompiliedContract struct {
	ABI ABI
	Bin hexutil.Bytes
}

// `BuildCompiledContract` builds a `CompiledContract` from an ABI string and a bytecode string.
func BuildCompiledContract(abiStr, bytecode string) CompiliedContract {
	var parsedAbi ABI
	if err := parsedAbi.UnmarshalJSON([]byte(abiStr)); err != nil {
		panic(err)
	}
	return CompiliedContract{
		ABI: parsedAbi,
		Bin: common.Hex2Bytes(bytecode),
	}
}
