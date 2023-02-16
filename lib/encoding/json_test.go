// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

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
