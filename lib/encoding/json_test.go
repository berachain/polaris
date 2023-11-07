// SPDX-License-Identifier: Apache-2.0
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package encoding_test

import (
	enclib "github.com/berachain/polaris/lib/encoding"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
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
