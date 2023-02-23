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

package solidity

//go:generate abigen --pkg generated --abi ./out/SolmateERC20.sol/SolmateERC20.abi.json --bin ./out/SolmateERC20.sol/SolmateERC20.bin --out ./generated/solmate_erc20.abigen.go --type SolmateERC20
//go:generate abigen --pkg generated --abi ./out/MockPrecompileInterface.sol/MockPrecompileInterface.abi.json --out ./generated/mock_precompile_interface.abigen.go --type MockPrecompile
//go:generate abigen --pkg generated --abi ./out/NonRevertableTx.sol/NonRevertableTx.abi.json --bin ./out/NonRevertableTx.sol/NonRevertableTx.bin --out ./generated/non_revertable_tx.abigen.go --type NonRevertableTx
