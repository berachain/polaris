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

package config

import (
	"fmt"
	"math/big"
	"time"

	"github.com/spf13/cast"

	servertypes "github.com/cosmos/cosmos-sdk/server/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// baseTen is for the big.Int string conversation.
const baseTen = 10

// AppOptions is for generating a mock.
type AppOptions interface {
	servertypes.AppOptions
}

// AppOptionsParser is a struct that holds the application options for parsing.
type AppOptionsParser struct {
	servertypes.AppOptions
}

// NewAppOptionsParser creates a new instance of AppOptionsParser with the provided
// AppOptions.
func NewAppOptionsParser(opts servertypes.AppOptions) *AppOptionsParser {
	return &AppOptionsParser{opts}
}

// GetCommonAddress returns the common.Address for the provided key.
func (c *AppOptionsParser) GetCommonAddress(key string) (common.Address, error) {
	addressStr, err := c.GetString(key)
	if err != nil {
		return common.Address{}, err
	}
	if !common.IsHexAddress(addressStr) {
		return common.Address{}, fmt.Errorf("invalid address: %s flag %s", addressStr, key)
	}
	return common.HexToAddress(addressStr), nil
}

// GetCommonAddressList retrieves a list of common.Address from a configuration key.
func (c *AppOptionsParser) GetCommonAddressList(key string) ([]common.Address, error) {
	addresses := make([]common.Address, 0)
	addressStrs := cast.ToStringSlice(c.Get(key))
	for _, addressStr := range addressStrs {
		address := common.HexToAddress(addressStr)
		if !common.IsHexAddress(addressStr) {
			return nil, fmt.Errorf("invalid address in list: %s flag %s", addressStr, key)
		}
		addresses = append(addresses, address)
	}
	return addresses, nil
}

// GetHexutilBytes returns a hexutil.Bytes from a configuration key.
func (c *AppOptionsParser) GetHexutilBytes(key string) (hexutil.Bytes, error) {
	bytesStr, err := c.GetString(key)
	if err != nil {
		return hexutil.Bytes{}, err
	}
	return hexutil.Decode(bytesStr)
}

// GetString retrieves a string value from a configuration key.
func (c *AppOptionsParser) GetString(key string) (string, error) {
	return handleError(c, cast.ToStringE, key)
}

// GetInt retrieves an integer value from a configuration key.
func (c *AppOptionsParser) GetInt(key string) (int, error) {
	return handleError(c, cast.ToIntE, key)
}

// GetInt64 retrieves a int64 value from a configuration key.
func (c *AppOptionsParser) GetInt64(key string) (int64, error) {
	return handleError(c, cast.ToInt64E, key)
}

// GetUint64 retrieves a uint64 value from a configuration key.
func (c *AppOptionsParser) GetUint64(key string) (uint64, error) {
	return handleError(c, cast.ToUint64E, key)
}

// GetUint64Ptr retrieves a pointer to a uint64 value fro	m a configuration key.
func (c *AppOptionsParser) GetUint64Ptr(key string) (*uint64, error) {
	return handleErrorPtr(c, cast.ToUint64E, key)
}

// GetBigInt retrieves a big.Int value from a configuration key.
func (c *AppOptionsParser) GetBigInt(key string) (*big.Int, error) {
	return handleErrorPtr(c, handleBigInt, key)
}

// GetFloat64 retrieves a float64 value from a configuration key.
func (c *AppOptionsParser) GetFloat64(key string) (float64, error) {
	return handleError(c, cast.ToFloat64E, key)
}

// GetBool retrieves a boolean value from a configuration key.
func (c *AppOptionsParser) GetBool(key string) (bool, error) {
	return handleError(c, cast.ToBoolE, key)
}

// GetStringSlice retrieves a slice of strings from a configuration key.
func (c *AppOptionsParser) GetStringSlice(key string) ([]string, error) {
	return handleError(c, cast.ToStringSliceE, key)
}

// GetTimeDuration retrieves a time.Duration value from a configuration key.
func (c *AppOptionsParser) GetTimeDuration(key string) (time.Duration, error) {
	return handleError(c, cast.ToDurationE, key)
}

// isNilRepresentation returns true if the provided value is "<nil>", "null" or "".
// it is used to determine when we need to initialize a nil ptr for a value to
// prevent the sdk from panicking on startup due to weird value.
func (c *AppOptionsParser) isNilRepresentation(value string) bool {
	return value == "<nil>" || value == "null" || value == ""
}

// handleError handles an error for a given flag in the AppOptionsParser.
// It attempts to cast the flag's value using the provided castFn and returns the result.
// If the cast fails, it returns an error.
func handleError[T any](
	c *AppOptionsParser, castFn func(interface{}) (T, error), flag string) (T, error) {
	var val T
	var err error
	if val, err = castFn(c.Get(flag)); err != nil {
		var t T
		return t, fmt.Errorf(
			"error while reading configuration: %w flag: %s", err, flag)
	}
	return val, nil
}

// handleErrorPtr handles an error for a given flag in the AppOptionsParser.
// It attempts to cast the flag's value using the provided castFn and returns a pointer to
// the result. If the cast fails, it returns an error. If the flag's value is empty,
// it returns a nil pointer.
func handleErrorPtr[T any](
	c *AppOptionsParser, castFn func(interface{}) (T, error), flag string) (*T, error) {
	var num string
	var err error
	if num, err = cast.ToStringE((c.Get(flag))); err != nil {
		return nil, fmt.Errorf("error while reading string: %w flag: %s", err, flag)
	} else if c.isNilRepresentation(num) {
		return (*T)(nil), nil
	}
	var val T
	if val, err = castFn(num); err != nil {
		return nil, fmt.Errorf("error while converting to value: %w flag: %s", err, flag)
	}
	return &val, nil
}

// handleBigInt handles a big.Int value from the given numStr interface.
// It attempts to parse the numStr as a big.Int and returns the result.
// If parsing fails, it returns an error.
func handleBigInt(numStr interface{}) (big.Int, error) {
	num, ok := new(big.Int).SetString(numStr.(string), baseTen)
	if !ok {
		return big.Int{}, fmt.Errorf("invalid big.Int string: %s", numStr.(string))
	}
	return *num, nil
}
