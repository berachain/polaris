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
	"math/big"
	"testing"
	"time"

	"github.com/berachain/polaris/cosmos/config/mocks"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/config")
}

var _ = Describe("Parser", func() {
	var parser *AppOptionsParser
	var appOpts *mocks.AppOptions

	BeforeEach(func() {
		appOpts = new(mocks.AppOptions)
		parser = NewAppOptionsParser(appOpts)
	})

	It("should set and retrieve a string option", func() {
		value := "testValue"
		runTest(appOpts, parser.GetString, value)
	})

	It("should set and retrieve an integer option", func() {
		value := int(42)
		runTest(appOpts, parser.GetInt, value)
	})

	It("should handle an int64 option", func() {
		value := int64(42)
		runTest(appOpts, parser.GetInt64, value)
	})

	It("should set and retrieve a uint64 option", func() {
		value := uint64(42)
		runTest(appOpts, parser.GetUint64, value)
	})

	It("should set and retrieve a pointer to a uint64 option", func() {
		value := uint64(42)
		runTestWithOutput(appOpts, parser.GetUint64Ptr, "42", &value)
	})

	It("should set and retrieve a big.Int option", func() {
		value := new(big.Int).SetInt64(42)
		runTestWithOutput(appOpts, parser.GetBigInt, "42", value)
	})

	It("should set and retrieve a float64 option", func() {
		value := 3.14159
		runTest(appOpts, parser.GetFloat64, value)
	})

	It("should set and retrieve a boolean option", func() {
		value := true
		runTest(appOpts, parser.GetBool, value)
	})

	It("should set and retrieve a string slice option", func() {
		value := []string{"apple", "banana", "cherry"}
		runTest(appOpts, parser.GetStringSlice, value)
	})

	It("should set and retrieve a time.Duration option", func() {
		value := time.Second * 10
		runTest(appOpts, parser.GetTimeDuration, value)
	})

	It("should set and retrieve a common.Address option", func() {
		addressStr := "0x18df82c7e422a42d47345ed86b0e935e9718ebda"
		runTestWithOutput(
			appOpts, parser.GetCommonAddress, addressStr, common.HexToAddress(addressStr))
	})

	It("should set and retrieve a list of common.Address options", func() {
		addressStrs := []string{
			"0x20f33ce90a13a4b5e7697e3544c3083b8f8a51d4",
			"0x18df82c7e422a42d47345ed86b0e935e9718ebda",
		}
		expectedAddresses := []common.Address{
			common.HexToAddress(addressStrs[0]),
			common.HexToAddress(addressStrs[1]),
		}

		// Run the test using the runTest function
		runTestWithOutput(
			appOpts, parser.GetCommonAddressList, addressStrs, expectedAddresses)
	})

	It("should set and retrieve a hexutil.Bytes option", func() {
		bytesStr := "0x1234567890abcdef"
		expectedBytes := hexutil.MustDecode(bytesStr)

		// Run the test using the runTest function
		runTest(appOpts, parser.GetHexutilBytes, expectedBytes)
	})
})

// runTest handles testing of various types.
func runTest[A any](
	appOpts *mocks.AppOptions, parser func(string) (A, error), value A) {
	runTestWithOutput(appOpts, parser, value, value)
}

// runTest handles testing of various types.
func runTestWithOutput[A, B any](
	appOpts *mocks.AppOptions, parser func(string) (B, error), value A, output B) {
	// Set the value.
	appOpts.On("Get", "myTestKey").Return(value).Once()

	// Retrieve the option
	retrievedValue, err := parser("myTestKey")

	Expect(err).ToNot(HaveOccurred())
	Expect(retrievedValue).To(Equal(output))
}
