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
