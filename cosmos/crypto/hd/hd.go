// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
