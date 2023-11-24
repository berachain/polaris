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
