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

package debug

import (
	"reflect"
	"runtime"
	"strings"
)

const (
	// hyphen is included in the name of all runtime functions as [name]-fm.
	hyphen = `-`
	// dot is included in the name of all runtime functions as [pkg_name].[name]-fm.
	dot = `.`
)

// GetFnName returns the name of a function `fn`.
func GetFnName(fn any) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	brokenUpNames := strings.Split(fullName, dot) // guarantees len(brokenUpName) >= 1
	brokenUpName := brokenUpNames[len(brokenUpNames)-1]

	dehyphenatedName := strings.Split(brokenUpName, hyphen)
	if len(dehyphenatedName) > 0 {
		return dehyphenatedName[0]
	}

	return brokenUpName
}
