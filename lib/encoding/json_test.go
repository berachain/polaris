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

package encoding_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	enclib "github.com/berachain/stargazer/lib/encoding"
)

var _ = Describe("MustMarshalJSON", func() {
	type testStruct struct {
		Foo string `json:"foo"`
		Bar int    `json:"bar"`
	}

	It("marshals JSON successfully", func() {
		test := testStruct{Foo: "hello", Bar: 123}
		expectedJSON := []byte(`{"foo":"hello","bar":123}`)

		resultJSON := enclib.MustMarshalJSON(test)
		Expect(resultJSON).To(Equal(expectedJSON))
	})

	It("panics when unable to marshal JSON", func() {
		test := make(chan int)
		Expect(func() {
			enclib.MustMarshalJSON(test)
		}).To(Panic())
	})
})

var _ = Describe("MustUnmarshalJSON", func() {
	type testStruct struct {
		Foo string `json:"foo"`
		Bar int    `json:"bar"`
	}

	It("unmarshals JSON successfully", func() {
		expected := testStruct{Foo: "hello", Bar: 123}
		testJSON := []byte(`{"foo":"hello","bar":123}`)

		var result testStruct
		resultPtr := enclib.MustUnmarshalJSON[testStruct](testJSON)
		result = *resultPtr
		Expect(result).To(Equal(expected))
	})

	It("panics when unable to unmarshal JSON", func() {
		testJSON := []byte(`{"foo":"hello","bar":"world"}`)

		var result *testStruct
		Expect(func() {
			result = enclib.MustUnmarshalJSON[testStruct](testJSON)
		}).To(Panic())

		Expect(result).To(BeNil())
	})
})
