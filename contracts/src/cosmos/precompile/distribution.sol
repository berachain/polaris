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

import {IBankModule} from "./bank.sol";

/**
 * @dev Interface of the distribution module's precompiled contract
 */
interface IDistributionModule {
    /**
     * @dev The caller (msg.sender) can set the address that will receive the deligation rewards.
     */
    function setWithdrawAddress(address withdrawAddress) external returns (bool);

    /**
     * @dev The caller (msg.sender) can set the address that will receive the deligation rewards.
     * Howver taking in a bech32 address.
     */
    function setWithdrawAddress(string calldata withdrawAddress) external returns (bool);

    /**
     * @dev Withdraw the rewrads accumilated by the caller(msg.sender). Returns the rewards claimed.
     */
    function withdrawDelegatorReward(address delegator, address validator)
        external
        returns (IBankModule.Coin[] memory);

    /**
     * @dev Withdraw the rewrads accumilated by the delegator from the validagor. Returns the rewards claimed.
     * However taking in a bech32 address.
     */
    function withdrawDelegatorReward(string calldata delegator, string calldata validator)
        external
        returns (IBankModule.Coin[] memory);

    /**
     * @dev Emitted by the distribution module when `amount` is withdrawn from a delegation with
     * `validator` as rewards.
     */
    event WithdrawRewards(address indexed validator, uint256 amount);
}
