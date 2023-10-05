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

	"pkg.berachain.dev/polaris/cosmos/config/mocks"

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
		appOpts = new(mocks.AppOptions) // Initialize the mock
		parser = NewAppOptionsParser(appOpts)
	})

	// GetString
	It("should set and retrieve a string option", func() {
		value := "testValue"
		runTest(appOpts, parser.GetString, value)
	})

	// GetInt
	It("should set and retrieve an integer option", func() {
		value := int(42)
		runTest(appOpts, parser.GetInt, value)
	})

	// GetInt64
	It("should handle an int64 option", func() {
		value := int64(42)
		runTest(appOpts, parser.GetInt64, value)
	})

	// GetUint64
	It("should set and retrieve a uint64 option", func() {
		value := uint64(42)
		runTest(appOpts, parser.GetUint64, value)
	})

	// GetUint64Ptr
	It("should set and retrieve a pointer to a uint64 option", func() {
		value := uint64(42)
		runTestUint64Ptr(appOpts, parser.GetUint64Ptr, value)
	})

	// GetBigInt
	It("should set and retrieve a big.Int option", func() {
		value := new(big.Int).SetInt64(42)
		runTestBigInt(appOpts, parser.GetBigInt, *value)
	})

	// GetFloat64
	It("should set and retrieve a float64 option", func() {
		value := 3.14159
		runTest(appOpts, parser.GetFloat64, value)
	})

	// GetBool
	It("should set and retrieve a boolean option", func() {
		value := true
		runTest(appOpts, parser.GetBool, value)
	})

	// GetStringSlice
	It("should set and retrieve a string slice option", func() {
		value := []string{"apple", "banana", "cherry"}
		runTest(appOpts, parser.GetStringSlice, value)
	})

	// GetTimeDuration
	It("should set and retrieve a time.Duration option", func() {
		value := time.Second * 10
		runTest(appOpts, parser.GetTimeDuration, value)
	})
})

// runTest handles testing of various types.
func runTest[A any](
	appOpts *mocks.AppOptions, parser func(string) (A, error), value A) {
	// Set the value.
	appOpts.On("Get", "myTestKey").Return(value).Once()

	// Retrieve the option
	retrievedValue, err := parser("myTestKey")

	Expect(err).ToNot(HaveOccurred())
	Expect(retrievedValue).To(Equal(value))
}

// runTestUint64Ptr handles testing of uint64 pointers.
func runTestUint64Ptr(
	appOpts *mocks.AppOptions, parser func(string) (*uint64, error), value uint64) {
	// Set the value.
	appOpts.On("Get", "myTestKey").Return("42").Once()

	// Retrieve the option
	retrievedValue, err := parser("myTestKey")

	Expect(err).ToNot(HaveOccurred())
	Expect(*retrievedValue).To(Equal(value))
}

// runTestBigInt handles testing of big.Int values.
func runTestBigInt(
	appOpts *mocks.AppOptions, parser func(string) (*big.Int, error), value big.Int) {
	// Set the value.
	appOpts.On("Get", "myTestKey").Return("42").Once()

	// Retrieve the option
	retrievedValue, err := parser("myTestKey")

	Expect(err).ToNot(HaveOccurred())
	Expect(*retrievedValue).To(Equal(value))
}
