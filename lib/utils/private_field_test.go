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

package utils_test

import (
	"github.com/berachain/stargazer/lib/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type A struct {
	privateField  string
	privateField2 *A
}

var _ = Describe("Get Private Field", func() {
	var a1 *A
	var a2 *A

	BeforeEach(func() {
		a1 = &A{
			privateField: "a1",
		}
		a2 = &A{
			privateField:  "a2",
			privateField2: a1,
		}
	})

	It("should correctly return a primitive private field", func() {
		Expect(utils.MustGetPrivateFieldByName[string](a1, "privateField")).To(Equal("a1"))
		Expect(utils.MustGetPrivateFieldByName[*A](a1, "privateField2")).To(BeNil())
	})

	It("should correctly return an object private field", func() {
		Expect(utils.MustGetPrivateFieldByName[string](a2, "privateField")).To(Equal("a2"))
		a1Ptr, ok := utils.GetAs[*A](utils.MustGetPrivateFieldByName[*A](a2, "privateField2"))
		Expect(ok).To(BeTrue())
		Expect(a1Ptr.privateField).To(Equal("a1"))
		Expect(a1Ptr.privateField2).To(BeNil())
	})

	It("should panic when called on non-struct", func() {
		Expect(func() { utils.MustGetPrivateFieldByName[string]("non-struct", "") }).To(Panic())
	})

	It("should panic if incorrect field name is called", func() {
		Expect(func() { utils.MustGetPrivateFieldByName[string](a1, "privateField3") }).To(Panic())
	})
})
