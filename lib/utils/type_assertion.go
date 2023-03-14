// SPDX-License-Identifier: Apache-2.0
//

package utils

// `GetAs` returns `obj` as type `T`. Uses Golang built-in type assertion.
func GetAs[T any](obj any) (T, bool) {
	casted, ok := obj.(T)
	return casted, ok
}

// `MustGetAs` returns `obj` as type `T`. Will panic if `obj` is not of type `T`. Uses Golang
// built-in type assertion.
func MustGetAs[T any](obj any) T {
	return obj.(T)
}

// `Implements` returns whether `obj` implements `T`. Uses Golang built-in type assertion.
func Implements[T any](obj any) bool {
	_, ok := obj.(T)
	return ok
}
