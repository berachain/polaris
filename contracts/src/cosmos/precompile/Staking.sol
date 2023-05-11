// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

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
    event Delegate(address indexed validator, Coin[] amount);

    /**
     * @dev Emitted by the staking module when `amount` tokens are redelegated from
     * `sourceValidator` to `destinationValidator`
     */
    event Redelegate(address indexed sourceValidator, address indexed destinationValidator, Coin[] amount);

    /**
     * @dev Emitted by the staking module when `amount` tokens are used to create `validator`
     */
    event CreateValidator(address indexed validator, Coin[] amount);

    /**
     * @dev Emitted by the staking module when `amount` tokens are unbonded from `validator`
     */
    event Unbond(address indexed validator, Coin[] amount);

    /**
     * @dev Emitted by the staking module when `amount` tokens are canceled from `delegator`'s
     * unbonding delegation with `validator`
     */
    event CancelUnbondingDelegation(
        address indexed validator, address indexed delegator, Coin[] amount, int64 creationHeight
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
    function getDelegation(address delegatorAddress, address validatorAddress) external view returns (uint256);

    /**
     * @dev Returns the `amount` of tokens currently delegated by `delegatorAddress` to
     * `validatorAddress` (at hex bech32 address)
     */
    function getDelegation(string calldata delegatorAddress, string calldata validatorAddress)
        external
        view
        returns (uint256);

    /**
     * @dev Returns a time-ordered list of all UnbondingDelegationEntries between
     * `delegatorAddress` and `validatorAddress`
     */
    function getUnbondingDelegation(address delegatorAddress, address validatorAddress)
        external
        view
        returns (UnbondingDelegationEntry[] memory);

    /**
     * @dev Returns a time-ordered list of all UnbondingDelegationEntries between
     * `delegatorAddress` and `validatorAddress` (at hex bech32 address)
     */
    function getUnbondingDelegation(string calldata delegatorAddress, string calldata validatorAddress)
        external
        view
        returns (UnbondingDelegationEntry[] memory);

    /**
     * @dev Returns a list of `delegatorAddress`'s redelegating bonds from `srcValidator` to
     * `dstValidator`
     */
    function getRedelegations(address delegatorAddress, address srcValidator, address dstValidator)
        external
        view
        returns (RedelegationEntry[] memory);

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
    function delegate(address validatorAddress, uint256 amount) external payable returns (bool);

    /**
     * @dev msg.sender delegates the `amount` of tokens to `validatorAddress` (at hex bech32
     * address)
     */
    function delegate(string calldata validatorAddress, uint256 amount) external payable returns (bool);

    /**
     * @dev msg.sender undelegates the `amount` of tokens from `validatorAddress`
     */
    function undelegate(address validatorAddress, uint256 amount) external payable returns (bool);

    /**
     * @dev msg.sender undelegates the `amount` of tokens from `validatorAddress` (at hex bech32
     * address)
     */
    function undelegate(string calldata validatorAddress, uint256 amount) external payable returns (bool);

    /**
     * @dev msg.sender redelegates the `amount` of tokens from `srcValidator` to `validtorDstAddr`
     */
    function beginRedelegate(address srcValidator, address dstValidator, uint256 amount)
        external
        payable
        returns (bool);

    /**
     * @dev msg.sender redelegates the `amount` of tokens from `srcValidator` to `validtorDstAddr`
     * (at hex bech32 addresses)
     */
    function beginRedelegate(string calldata srcValidator, string calldata dstValidator, uint256 amount)
        external
        payable
        returns (bool);

    /**
     * @dev Cancels msg.sender's unbonding delegation with `validatorAddress` and delegates the
     * `amount` of tokens back to `validatorAddress`
     *
     * Provide the `creationHeight` of the original unbonding delegation
     */
    function cancelUnbondingDelegation(address validatorAddress, uint256 amount, int64 creationHeight)
        external
        payable
        returns (bool);

    /**
     * @dev Cancels msg.sender's unbonding delegation with `validatorAddress` and delegates the
     * `amount` of tokens back to `validatorAddress` (at hex bech32 addresses)
     *
     * Provide the `creationHeight` of the original unbonding delegation
     */
    function cancelUnbondingDelegation(string calldata validatorAddress, uint256 amount, int64 creationHeight)
        external
        payable
        returns (bool);

    //////////////////////////////////////////// UTILS ////////////////////////////////////////////
    /**
     * @dev Represents a cosmos coin.
     * Note: this struct is generated as go struct that is then used in the precompile.
     */
    struct Coin {
        uint256 amount;
        string denom;
    }

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
