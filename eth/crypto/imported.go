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

package crypto

import (
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	// GenerateEthKey is a function that generates a new Ethereum key.
	GenerateEthKey = crypto.GenerateKey
	// Keccak256 is a function that computes and returns the Keccak256 hash of the input data.
	Keccak256 = crypto.Keccak256
	// Keccak256Hash is a function that computes and returns the Keccak256 hash of the input data,
	// but the return type is Hash.
	Keccak256Hash = crypto.Keccak256Hash
	// PubkeyToAddress is a function that derives the Ethereum address from the given public key.
	PubkeyToAddress = crypto.PubkeyToAddress
	// LoadECDSA is a function that loads a private key from a given file.
	LoadECDSA = crypto.LoadECDSA
)
