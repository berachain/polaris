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

package lib

import (
	"cosmossdk.io/core/address"

	"github.com/ethereum/go-ethereum/common"
)

/* -------------------------------------------------------------------------- */
/*                     AccAddress, ValAddress, ConsAddress                    */
/* -------------------------------------------------------------------------- */

// EthAddressFromString converts a Cosmos SDK address string to an Ethereum `Address`.
func EthAddressFromString(codec address.Codec, addr string) (common.Address, error) {
	bz, err := codec.StringToBytes(addr)
	if err != nil {
		return common.Address{}, err
	}
	return common.BytesToAddress(bz), nil
}

// MustEthAddressFromString converts a Cosmos SDK address string to an Ethereum `Address`. It
// panics if the conversion fails.
func MustEthAddressFromString(codec address.Codec, addr string) common.Address {
	address, err := EthAddressFromString(codec, addr)
	if err != nil {
		panic(err)
	}
	return address
}

// StringFromEthAddress converts an Ethereum `Address` to a Cosmos SDK address string.
func StringFromEthAddress(codec address.Codec, ethAddress common.Address) (string, error) {
	return codec.BytesToString(ethAddress.Bytes())
}

// MustStringFromEthAddress converts an Ethereum `Address` to a Cosmos SDK address string. It
// panics if the conversion fails.
func MustStringFromEthAddress(codec address.Codec, ethAddress common.Address) string {
	addr, err := StringFromEthAddress(codec, ethAddress)
	if err != nil {
		panic(err)
	}
	return addr
}
