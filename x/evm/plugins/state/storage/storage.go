// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package storage

import (
	"fmt"

	"pkg.berachain.dev/polaris/lib/errors"
	libtypes "pkg.berachain.dev/polaris/lib/types"
)

// Compile-time type assertions.
var _ libtypes.Cloneable[Storage] = Storage{}
var _ fmt.Stringer = Storage{}

// `Storage` represents the account Storage map as a slice of single key-value
// Slot pairs. This helps to ensure that the Storage map can be iterated over
// deterministically.
type Storage []*Slot

// `ValidateBasic` performs basic validation of the Storage data structure.
// It checks for duplicate keys and calls `ValidateBasic` on each `State`.
func (s Storage) ValidateBasic() error {
	seenSlots := make(map[string]struct{})
	for i, slot := range s {
		if _, found := seenSlots[slot.Key]; found {
			return errors.Wrapf(ErrInvalidState, "duplicate state key %d: %s", i, slot.Key)
		}

		if err := slot.ValidateBasic(); err != nil {
			return err
		}

		seenSlots[slot.Key] = struct{}{}
	}
	return nil
}

// `String` implements `fmt.Stringer`.
func (s Storage) String() string {
	var str string
	for _, slot := range s {
		str += fmt.Sprintf("%s\n", slot.String())
	}

	return str
}

// `Clone` implements `types.Cloneable`.
func (s Storage) Clone() Storage {
	cpy := make(Storage, len(s))
	copy(cpy, s)

	return cpy
}
