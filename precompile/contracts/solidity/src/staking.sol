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

/**
 * @dev Interface of the staking module's precompiled contract
 */
interface IStakingModule {
    ////////////////////////////////////////// EVENTS /////////////////////////////////////////////

    /**
     * @dev Emitted by the staking module when `amount` tokens are delegated to
     * `validator`
     */
    event Delegate(address indexed validator, uint256 amount);

    /**
     * @dev Emitted by the staking module when `amount` tokens are redelegated from
     * `sourceValidator` to `destinationValidator`
     */
    event Redelegate(
        address indexed sourceValidator,
        address indexed destinationValidator,
        uint256 amount
    );

    /**
     * @dev Emitted by the staking module when `amount` tokens are used to create `validator`
     */
    event CreateValidator(address indexed validator, uint256 amount);

    /**
     * @dev Emitted by the staking module when `amount` tokens are unbonded from `validator`
     */
    event Unbond(address indexed validator, uint256 amount);

    /**
     * @dev Emitted by the staking module when `amount` tokens are canceled from `delegator`'s
     * unbonding delegation with `validator`
     */
    event CancelUnbondingDelegation(
        address indexed validator,
        address indexed delegator,
        uint256 amount,
        int64 creationHeight
    );

    /////////////////////////////////////// READ METHODS //////////////////////////////////////////

    /**
     * @dev Returns a list of active validators.
     */
    function getActiveValidators() external view returns (address[] memory);

    /**
     * @dev Returns the `amount` of tokens currently delegated by `delegatorAddress` to
     * `validatorAddress`
     */
    function getDelegation(
        address delegatorAddress,
        address validatorAddress
    ) external view returns (uint256);

    /**
     * @dev Returns the `amount` of tokens currently delegated by `delegatorAddress` to
     * `validatorAddress` (at hex bech32 address)
     */
    function getDelegation(
        string calldata delegatorAddress,
        string calldata validatorAddress
    ) external view returns (uint256);

    /**
     * @dev Returns a time-ordered list of all UnbondingDelegationEntries between
     * `delegatorAddress` and `validatorAddress`
     */
    function getUnbondingDelegation(
        address delegatorAddress,
        address validatorAddress
    ) external view returns (UnbondingDelegationEntry[] memory);

    /**
     * @dev Returns a time-ordered list of all UnbondingDelegationEntries between
     * `delegatorAddress` and `validatorAddress` (at hex bech32 address)
     */
    function getUnbondingDelegation(
        string calldata delegatorAddress,
        string calldata validatorAddress
    ) external view returns (UnbondingDelegationEntry[] memory);

    /**
     * @dev Returns a list of `delegatorAddress`'s redelegating bonds from `srcValidator` to
     * `dstValidator`
     */
    function getRedelegations(
        address delegatorAddress,
        address srcValidator,
        address dstValidator
    ) external view returns (RedelegationEntry[] memory);

    /**
     * @dev Returns a list of `delegatorAddress`'s redelegating bonds from `srcValidator` to
     * `dstValidator` (at hex bech32 addresses)
     */
    function getRedelegations(
        string calldata delegatorAddress,
        string calldata srcValidator,
        string calldata dstValidator
    ) external view returns (RedelegationEntry[] memory);

    ////////////////////////////////////// WRITE METHODS //////////////////////////////////////////

    /**
     * @dev msg.sender delegates the `amount` of tokens to `validatorAddress`
     */
    function delegate(
        address validatorAddress,
        uint256 amount
    ) external payable;

    /**
     * @dev msg.sender delegates the `amount` of tokens to `validatorAddress` (at hex bech32
     * address)
     */
    function delegate(
        string calldata validatorAddress,
        uint256 amount
    ) external payable;

    /**
     * @dev msg.sender undelegates the `amount` of tokens from `validatorAddress`
     */
    function undelegate(
        address validatorAddress,
        uint256 amount
    ) external payable;

    /**
     * @dev msg.sender undelegates the `amount` of tokens from `validatorAddress` (at hex bech32
     * address)
     */
    function undelegate(
        string calldata validatorAddress,
        uint256 amount
    ) external payable;

    /**
     * @dev msg.sender redelegates the `amount` of tokens from `srcValidator` to `validtorDstAddr`
     */
    function beginRedelegate(
        address srcValidator,
        address dstValidator,
        uint256 amount
    ) external payable;

    /**
     * @dev msg.sender redelegates the `amount` of tokens from `srcValidator` to `validtorDstAddr`
     * (at hex bech32 addresses)
     */
    function beginRedelegate(
        string calldata srcValidator,
        string calldata dstValidator,
        uint256 amount
    ) external payable;

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
    ) external payable;

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
    ) external payable;

    //////////////////////////////////////////// UTILS ////////////////////////////////////////////

    /**
     * @dev Represents one entry of an unbonding delegation
     *
     * Note: the field names of the native struct should match these field names (by camelCase)
     * Note: we are using the types in precompile/generated
     */
    struct UnbondingDelegationEntry {
        // creationHeight is the height which the unbonding took place
        int64 creationHeight;
        // completionTime is the unix time for unbonding completion, formatted as a string
        string completionTime;
        // initialBalance defines the tokens initially scheduled to receive at completion
        uint256 initialBalance;
        // balance defines the tokens to receive at completion
        uint256 balance;
        // unbondingingId incrementing id that uniquely identifies this entry
        uint64 unbondingId;
    }

    /**
     * @dev Represents a redelegation entry with relevant metadata
     *
     * Note: the field names of the native struct should match these field names (by camelCase)
     * Note: we are using the types in precompile/generated
     */
    struct RedelegationEntry {
        // creationHeight is the height which the redelegation took place
        int64 creationHeight;
        // completionTime is the unix time for redelegation completion, formatted as a string
        string completionTime;
        // initialBalance defines the initial balance when redelegation started
        uint256 initialBalance;
        // sharesDst is the amount of destination-validatorAddress shares created by redelegation
        uint256 sharesDst;
        // unbondingId is the incrementing id that uniquely identifies this entry
        uint64 unbondingId;
    }
}
