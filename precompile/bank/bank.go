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

package bank

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"pkg.berachain.dev/polaris/eth/accounts/abi"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/precompile"
	"pkg.berachain.dev/polaris/precompile/contracts/solidity/generated"
	evmutils "pkg.berachain.dev/polaris/x/evm/utils"
)

// `Contract` is the precompile contract for the bank module.
type Contract struct {
	contractAbi *abi.ABI
}

// `NewPrecompileContract` returns a new instance of the bank precompile contract.
func NewPrecompileContract() precompile.StatefulImpl {
	var contractAbi abi.ABI
	if err := contractAbi.UnmarshalJSON([]byte(generated.BankEventsMetaData.ABI)); err != nil {
		panic(err)
	}
	return &Contract{
		contractAbi: &contractAbi,
	}
}

// `RegistryKey` implements the `precompile.StatefulImpl` interface.
func (c *Contract) RegistryKey() common.Address {
	// Contract Address: 0x4381dC2aB14285160c808659aEe005D51255adD7
	return evmutils.AccAddressToEthAddress(authtypes.NewModuleAddress(banktypes.ModuleName))
}

// `AbiMethods` implements the `precompile.StatefulImpl` interface.
func (c *Contract) ABIMethods() map[string]abi.Method {
	return nil
}

// `PrecompileMethods` implements the `precompile.StatefulImpl` interface.
func (c *Contract) PrecompileMethods() precompile.Methods {
	return nil
}

// `AbiEvents` implements the `precompile.StatefulImpl` interface.
func (c *Contract) ABIEvents() map[string]abi.Event {
	return c.contractAbi.Events
}

// `CustomValueDecoders` implements StatefulImpl.
func (c *Contract) CustomValueDecoders() precompile.ValueDecoders {
	return nil
}
