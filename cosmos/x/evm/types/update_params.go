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

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	errorslib "pkg.berachain.dev/polaris/lib/errors"
)

// Compile-time interface assertions.
var _ sdk.Msg = (*UpdateParamsRequest)(nil)

// GetSigners returns the expected signers for the request.
func (m UpdateParamsRequest) GetSigners() []sdk.AccAddress {
	//#nosec G703: // purposely leave error unhandled.
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

// ValidateBasic validates the request.
func (m *UpdateParamsRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorslib.Wrap(err, "invalid authority address")
	}

	return m.Params.ValidateBasic()
}
