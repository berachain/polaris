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

import {IStakingModule} from "../src/staking.sol";

contract MockStaking is IStakingModule {
    constructor() {}

    /**
     * @dev Returns a list of active validators.
     */
    function getActiveValidators() external view returns (address[] memory) {
        address[] memory validators = new address[](1);
        validators[0] = address(0x1);
        return validators;
    }

    /**
     * @dev Returns the `amount` of tokens currently delegated by msg.sender to `validatorAddress`
     */
    function getDelegation(address validatorAddress)
        external
        view
        returns (uint256)
    {
        return 10 ether;
    }

    /**
     * @dev Returns the `amount` of tokens currently delegated by msg.sender to `validatorAddress`
     * (at hex bech32 address)
     */
    function getDelegation(string calldata validatorAddress)
        external
        view
        returns (uint256)
    {
        return 10 ether;
    }

    /**
     * @dev Returns a time-ordered list of all UnbondingDelegationEntries between msg.sender and
     * `validatorAddress`
     */
    function getUnbondingDelegation(address validatorAddress)
        external
        view
        returns (UnbondingDelegationEntry[] memory)
    {
        UnbondingDelegationEntry[]
            memory entries = new UnbondingDelegationEntry[](0);

        return entries;
    }

    /**
     * @dev Returns a time-ordered list of all UnbondingDelegationEntries between msg.sender and
     * `validatorAddress` (at hex bech32 address)
     */
    function getUnbondingDelegation(string calldata validatorAddress)
        external
        view
        returns (UnbondingDelegationEntry[] memory)
    {
        UnbondingDelegationEntry[]
            memory entries = new UnbondingDelegationEntry[](0);

        return entries;
    }

    /**
     * @dev Returns a list of the msg.sender's redelegating bonds from `srcValidator` to
     * `dstValidator`
     */
    function getRedelegations(address srcValidator, address dstValidator)
        external
        view
        returns (RedelegationEntry[] memory)
    {
        RedelegationEntry[] memory entries = new RedelegationEntry[](0);
        return entries;
    }

    /**
     * @dev Returns a list of the msg.sender's redelegating bonds from `srcValidator` to
     * `dstValidator` (at hex bech32 addresses)
     */
    function getRedelegations(
        string calldata srcValidator,
        string calldata dstValidator
    ) external view returns (RedelegationEntry[] memory) {
        RedelegationEntry[] memory entries = new RedelegationEntry[](0);
        return entries;
    }

    ////////////////////////////////////// WRITE METHODS //////////////////////////////////////////

    /**
     * @dev msg.sender delegates the `amount` of tokens to `validatorAddress`
     */
    function delegate(address validatorAddress, uint256 amount)
        external
        payable
    {
        require(msg.value == amount, "amount mismatch");
        payable(address(this)).transfer(msg.value);
    }

    /**
     * @dev msg.sender delegates the `amount` of tokens to `validatorAddress` (at hex bech32
     * address)
     */
    function delegate(string calldata validatorAddress, uint256 amount)
        external
        payable
    {
        require(msg.value == amount, "amount mismatch");
    }

    /**
     * @dev msg.sender undelegates the `amount` of tokens from `validatorAddress`
     */
    function undelegate(address validatorAddress, uint256 amount)
        external
        payable
    {
        require(msg.value == amount, "amount mismatch");
    }

    /**
     * @dev msg.sender undelegates the `amount` of tokens from `validatorAddress` (at hex bech32
     * address)
     */
    function undelegate(string calldata validatorAddress, uint256 amount)
        external
        payable
    {
        require(msg.value == amount, "amount mismatch");
    }

    /**
     * @dev msg.sender redelegates the `amount` of tokens from `srcValidator` to
     * `validtorDstAddr`
     */
    function beginRedelegate(
        address srcValidator,
        address dstValidator,
        uint256 amount
    ) external payable {
        require(msg.value == amount, "amount mismatch");
    }

    /**
     * @dev msg.sender redelegates the `amount` of tokens from `srcValidator` to
     * `validtorDstAddr` (at hex bech32 addresses)
     */
    function beginRedelegate(
        string calldata srcValidator,
        string calldata dstValidator,
        uint256 amount
    ) external payable {
        require(msg.value == amount, "amount mismatch");
    }

    /**
     * @dev Cancels msg.sender's unbonding delegation with `validatorAddress` and delegates the
     * `amount` of tokens back to `validatorAddress`
     *
     * Provide the `creationHeight` of the original unbonding delegation
     */
    function cancelUnbondingDelegation(
        address validatorAddress,
        uint256 amount,
        int64 creationHeight
    ) external payable {
        require(msg.value == amount, "amount mismatch");
    }

    /**
     * @dev Cancels msg.sender's unbonding delegation with `validatorAddress` and delegates the
     * `amount` of tokens back to `validatorAddress` (at hex bech32 addresses)
     *
     * Provide the `creationHeight` of the original unbonding delegation
     */
    function cancelUnbondingDelegation(
        string calldata validatorAddress,
        uint256 amount,
        int64 creationHeight
    ) external payable {
        require(msg.value == amount, "amount mismatch");
    }

    receive() external payable {}
}
