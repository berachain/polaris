// SPDX-License-Identifier: Apache-2.0
//

package errors

import "fmt"

func Wrap(err error, desc string) error {
	return fmt.Errorf("%s: %w", desc, err)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), err)
}
