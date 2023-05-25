package encoding_test

import (
	enclib "pkg.berachain.dev/polaris/lib/encoding"

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
