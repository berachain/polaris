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

package mock

import libtypes "github.com/berachain/polaris/lib/types"

// Assert that `MockRegistrable` implements `Registrable`.
var _ libtypes.Registrable[string] = &Registrable{}

type Registrable struct {
	registerKey string
	data        string
}

func NewMockRegistrable(registerKey string, data string) *Registrable {
	return &Registrable{
		registerKey: registerKey,
		data:        data,
	}
}

func (m Registrable) RegistryKey() string {
	return m.registerKey
}

func (m Registrable) Data() string {
	return m.data
}
