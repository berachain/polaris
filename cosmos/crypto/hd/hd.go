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

//
//nolint:gomnd // ignore magic numbers
package hd

import (
	"crypto/ecdsa"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/cosmos/go-bip39"
)

func GenerateWallet(mnemonic string) (*ecdsa.PrivateKey, *string, error) {
	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, "")
	// Generate a new master node using the seed.
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, nil, err
	}
	// This gives the path: m/44H
	acc44H, err := masterKey.Derive(hdkeychain.HardenedKeyStart + 44)
	if err != nil {
		return nil, nil, err
	}
	// This gives the path: m/44H/60H
	acc44H60H, err := acc44H.Derive(hdkeychain.HardenedKeyStart + 60)
	if err != nil {
		return nil, nil, err
	}
	// This gives the path: m/44H/60H/0H
	acc44H60H0H, err := acc44H60H.Derive(hdkeychain.HardenedKeyStart + 0)
	if err != nil {
		return nil, nil, err
	}
	// This gives the path: m/44H/60H/0H/0
	acc44H60H0H0, err := acc44H60H0H.Derive(0)
	if err != nil {
		return nil, nil, err
	}
	// This gives the path: m/44H/60H/0H/0/0
	acc44H60H0H00, err := acc44H60H0H0.Derive(0)
	if err != nil {
		return nil, nil, err
	}
	btcecPrivKey, err := acc44H60H0H00.ECPrivKey()
	if err != nil {
		return nil, nil, err
	}
	privateKey := btcecPrivKey.ToECDSA()
	path := "m/44H/60H/0H/0/0"
	return privateKey, &path, nil
}
