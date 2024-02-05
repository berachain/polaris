// SPDX-License-Identifier: MIT
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

pragma solidity 0.8.23;

import {Cosmos} from "../CosmosTypes.sol";

/**
 * @dev Interface of the distribution module's precompiled contract
 */
interface IDistributionModule {
    ////////////////////////////////////////// EVENTS /////////////////////////////////////////////
    /**
     * @dev Emitted by the distribution module when `amount` is withdrawn from a delegation with
     * `validator` as rewards.
     * @param validator The validator address to withdraw the rewards from.
     * @param amount The amount of rewards withdrawn.
     */
    event WithdrawRewards(address indexed validator, Cosmos.Coin[] amount);

    /**
     * @dev Emitted by the distribution module when `withdrawAddress` is set to receive rewards
     * upon withdrawal.
     * @param withdrawAddress The address to set as the withdraw address.
     */
    event SetWithdrawAddress(address indexed withdrawAddress);

    /////////////////////////////////////// READ METHODS //////////////////////////////////////////

    /**
     * @dev Returns whether withdrawing delegation rewards is enabled.
     */
    function getWithdrawEnabled() external view returns (bool);

    /**
     * @dev Returns the address that will receive the delegation rewards.
     * @param delegator the delegator for which the withdraw address is returned.
     */
    function getWithdrawAddress(address delegator) external view returns (address);

    /**
     * @dev Returns the rewards accumulated by the delegator for the validator.
     * @param delegator The delegator to retrieve the rewards for.
     * @param validator The validator (operator address) to retrieve the rewards for.
     */
    function getDelegatorValidatorReward(address delegator, address validator)
        external
        view
        returns (Cosmos.Coin[] memory);

    /**
     * @dev Returns the all rewards accumulated by the delegator.
     * @param delegator The delegator to retrieve the totalRewards for.
     */
    function getAllDelegatorRewards(address delegator) external view returns (ValidatorReward[] memory);

    /**
     * @dev Returns the total rewards accumulated by the delegator.
     * @param delegator The delegator to retrieve the totalRewards for.
     */
    function getTotalDelegatorReward(address delegator) external view returns (Cosmos.Coin[] memory);

    ////////////////////////////////////// WRITE METHODS //////////////////////////////////////////

    /**
     * @dev The caller (msg.sender) can set the address that will receive the delegation rewards.
     * @param withdrawAddress The address to set as the withdraw address.
     */
    function setWithdrawAddress(address withdrawAddress) external returns (bool);

    /**
     * @dev Withdraw the rewrads accumulated by the caller(msg.sender). Returns the rewards claimed.
     * @param delegator The delegator to withdraw the rewards from.
     * @param validator The validator (operator address) to withdraw the rewards from.
     */
    function withdrawDelegatorReward(address delegator, address validator) external returns (Cosmos.Coin[] memory);

    //////////////////////////////////////////// UTILS ////////////////////////////////////////////

    /**
     * @dev Represents a delegator's rewards for one particular validator.
     */
    struct ValidatorReward {
        address validator;
        Cosmos.Coin[] rewards;
    }
}
