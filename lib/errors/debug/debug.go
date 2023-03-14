// SPDX-License-Identifier: Apache-2.0
//

package debug

import (
	"reflect"
	"runtime"
	"strings"
)

// GetFnName returns the name of a function `fn`.
func GetFnName(fn any) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	brokenUpName := strings.Split(fullName, ".") // guarantees len(brokenUpName) >= 1
	return brokenUpName[len(brokenUpName)-1]
}
