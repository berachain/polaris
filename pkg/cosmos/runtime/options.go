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

package runtime

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec/types"

	cryptocodec "pkg.berachain.dev/polaris/cosmos/crypto/codec"
)

// `PolarisAppOptions` is a list of `func(*baseapp.BaseApp)` that are used to configure the baseapp.
func PolarisAppOptions(
	interfaceRegistry types.InterfaceRegistry, baseAppOptions ...func(*baseapp.BaseApp),
) []func(*baseapp.BaseApp) {
	polarisAppOptions := baseAppOptions
	polarisAppOptions = append(
		polarisAppOptions,
		[]func(bApp *baseapp.BaseApp){
			func(bApp *baseapp.BaseApp) {
				cryptocodec.RegisterInterfaces(interfaceRegistry)
			},
		}...,
	)
	return polarisAppOptions
}
