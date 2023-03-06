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

pragma solidity ^0.8.4;

import "forge-std/Script.sol";
import "./LiquidStaking.sol";
import "../staking.sol";

contract DeployAndStake is Script {
    LiquidStaking public staking;

    // function run() public {
    //     vm.startBroadcast();

    //     address precompile = address(0x12);
    //     address validator = address(0x34);

    //     staking = new LiquidStaking("name", "SYMB", precompile, validator);

    //     vm.stopBroadcast();
    // }

    function run() public {
        vm.startBroadcast();
        // Precomile address
        address precompile = address(
            0xd9A998CaC66092748FfEc7cFBD155Aae1737C2fF
        );

        // Get the first validator from the precompile.
        // address validator = IStakingModule(precompile).getActiveValidators()[0];
        address[] memory validator = IStakingModule(precompile)
            .getActiveValidators();

        // Deploy the staking contract.
        // staking = new LiquidStaking("name", "SYMB", precompile, validator);

        // Stop the broadcast.
        vm.stopBroadcast();
    }
}
