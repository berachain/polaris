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

const (
	// DefaultEvmDenom is the default EVM denom.
	DefaultEvmDenom = "abera"
)

var (
	// DefaultExtraEIPs is the default extra EIPs.
	DefaultExtraEIPs = []int64{}
)

// DefaultParams contains the default values for all parameters.
func DefaultParams() *Params {
	return &Params{
		EvmDenom:  DefaultEvmDenom,
		ExtraEIPs: DefaultExtraEIPs,
	}
}

// EthereumChainConfig returns the chain config as a struct.
// ValidateBasic is used to validate the parameters.
func (p *Params) ValidateBasic() error {
	if p.EvmDenom == "" {
		return ErrNoEvmDenom
	}
	if p.ExtraEIPs == nil {
		return ErrNoExtraEIPs
	}
	return nil
}
