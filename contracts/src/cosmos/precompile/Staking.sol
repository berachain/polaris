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

pragma solidity 0.8.23;

import {Cosmos} from "../CosmosTypes.sol";

/**
 * @dev Interface of the staking module's precompiled contract
 */
interface IStakingModule {
    ////////////////////////////////////////// EVENTS /////////////////////////////////////////////

    /**
     * @dev Emitted by the staking module when `amount` tokens are delegated to
     * `validator`
     * @param validator The validator operator address
     * @param amount The amount of tokens delegated
     */
    event Delegate(address indexed validator, Cosmos.Coin[] amount);

    /**
     * @dev Emitted by the staking module when `amount` tokens are redelegated from
     * `sourceValidator` to `destinationValidator`
     * @param sourceValidator The source validator operator address
     * @param destinationValidator The destination validator operator address
     * @param amount The amount of tokens redelegated
     */
    event Redelegate(address indexed sourceValidator, address indexed destinationValidator, Cosmos.Coin[] amount);

    /**
     * @dev Emitted by the staking module when `amount` tokens are used to create `validator`
     * @param validator The validator operator address
     * @param amount The amount of tokens used to create the validator
     */
    event CreateValidator(address indexed validator, Cosmos.Coin[] amount);

    /**
     * @dev Emitted by the staking module when `amount` tokens are unbonded from `validator`
     * @param validator The validator operator address
     * @param amount The amount of tokens unbonded
     */
    event Unbond(address indexed validator, Cosmos.Coin[] amount);

    /**
     * @dev Emitted by the staking module when `amount` tokens are canceled from `delegator`'s
     * unbonding delegation with `validator`
     * @param validator The validator operator address
     * @param delegator The delegator address
     * @param amount The amount of tokens canceled
     * @param creationHeight The height at which the unbonding delegation was created
     */
    event CancelUnbondingDelegation(
        address indexed validator, address indexed delegator, Cosmos.Coin[] amount, int64 creationHeight
    );

    /////////////////////////////////////// READ METHODS //////////////////////////////////////////

    /**
     * @dev Returns the operator address of the validator for the given consensus address.
     * @param consAddr The consensus address (as bytes) of the validator
     */
    function getValAddressFromConsAddress(bytes calldata consAddr) external pure returns (address);

    /**
     * @dev Returns a list of all active validators.
     * @notice Accepts pagination request (empty == no pagination returned).
     */
    function getValidators(Cosmos.PageRequest calldata pagination)
        external
        view
        returns (Validator[] memory, Cosmos.PageResponse memory);

    /**
     * @dev Returns a list of bonded validator (operator) addresses.
     * @notice Accepts pagination request (empty == no pagination returned).
     */
    function getBondedValidators(Cosmos.PageRequest calldata pagination)
        external
        view
        returns (Validator[] memory, Cosmos.PageResponse memory);

    /**
     * @dev Returns a list of bonded validator (operator) addresses, sorted by power (stake) in
     * descending order.
     */
    function getBondedValidatorsByPower() external view returns (address[] memory);

    /**
     * @dev Returns the validator at the given address.
     * @param validatorAddress The validator operator address
     */
    function getValidator(address validatorAddress) external view returns (Validator memory);

    /**
     * @dev Returns all the validators delegated to by the given delegator.
     * @notice Accepts pagination request (empty == no pagination returned).
     * @param delegatorAddress The delegator address to query validators for.
     */
    function getDelegatorValidators(address delegatorAddress, Cosmos.PageRequest calldata pagination)
        external
        view
        returns (Validator[] memory, Cosmos.PageResponse memory);

    /**
     * @dev Returns all the delegations delegated to the given validator.
     * @notice Accepts pagination request (empty == no pagination returned).
     * @param validatorAddress The validator operator address to query delegations for.
     */
    function getValidatorDelegations(address validatorAddress, Cosmos.PageRequest calldata pagination)
        external
        view
        returns (Delegation[] memory, Cosmos.PageResponse memory);

    /**
     * @dev Returns the `amount` of tokens currently delegated by `delegatorAddress` to
     * `validatorAddress`
     * @param delegatorAddress The delegator address
     * @param validatorAddress The validator operator address
     */
    function getDelegation(address delegatorAddress, address validatorAddress) external view returns (uint256);

    /**
     * @dev Returns a time-ordered list of all UnbondingDelegationEntries between
     * `delegatorAddress` and `validatorAddress`
     * @param delegatorAddress The delegator address
     * @param validatorAddress The validator operator address
     */
    function getUnbondingDelegation(address delegatorAddress, address validatorAddress)
        external
        view
        returns (UnbondingDelegationEntry[] memory);

    /**
     * @dev Returns a list of all unbonding delegations for a given delegator
     * @notice Accepts pagination request (empty == no pagination returned).
     * @param delegatorAddress The delegator address
     */
    function getDelegatorUnbondingDelegations(address delegatorAddress, Cosmos.PageRequest calldata pagination)
        external
        view
        returns (UnbondingDelegation[] memory, Cosmos.PageResponse memory);

    /**
     * @dev Returns a list of `delegatorAddress`'s redelegating bonds from `srcValidator` to
     * `dstValidator`
     * @notice Accepts pagination request (empty == no pagination returned).
     * @param delegatorAddress The delegator address
     * @param srcValidator The source validator operator address
     * @param dstValidator The destination validator operator address
     */
    function getRedelegations(
        address delegatorAddress,
        address srcValidator,
        address dstValidator,
        Cosmos.PageRequest calldata pagination
    ) external view returns (RedelegationEntry[] memory, Cosmos.PageResponse memory);

    ////////////////////////////////////// WRITE METHODS //////////////////////////////////////////

    /**
     * @dev msg.sender delegates the `amount` of tokens to `validatorAddress`
     * @param validatorAddress The validator operator address
     * @param amount The amount of tokens to delegate
     */
    function delegate(address validatorAddress, uint256 amount) external payable returns (bool);

    /**
     * @dev msg.sender undelegates the `amount` of tokens from `validatorAddress`
     * @param validatorAddress The validator operator address
     * @param amount The amount of tokens to undelegate
     */
    function undelegate(address validatorAddress, uint256 amount) external payable returns (bool);

    /**
     * @dev msg.sender redelegates the `amount` of tokens from `srcValidator` to `validtorDstAddr`
     * @param srcValidator The source validator operator address
     * @param dstValidator The destination validator operator address
     * @param amount The amount of tokens to redelegate
     */
    function beginRedelegate(address srcValidator, address dstValidator, uint256 amount)
        external
        payable
        returns (bool);

    /**
     * @dev Cancels msg.sender's unbonding delegation with `validatorAddress` and delegates the
     * `amount` of tokens back to `validatorAddress`
     * @param validatorAddress The validator operator address
     * @param amount The amount of tokens to cancel
     * @param creationHeight The height at which the unbonding delegation was created
     */
    function cancelUnbondingDelegation(address validatorAddress, uint256 amount, int64 creationHeight)
        external
        payable
        returns (bool);

    //////////////////////////////////////////// UTILS ////////////////////////////////////////////

    /**
     * @dev Represents a validator.
     */
    struct Validator {
        address operatorAddr;
        bytes consAddr;
        bool jailed;
        string status;
        uint256 tokens;
        uint256 delegatorShares;
        Description description;
        int64 unbondingHeight;
        string unbondingTime;
        Commission commission;
        uint256 minSelfDelegation;
        int64 unbondingOnHoldRefCount;
        uint64[] unbondingIds;
    }

    /**
     * @dev Represents the commission parameters for a given validator.
     */
    struct Commission {
        CommissionRates commissionRates;
        string updateTime;
    }

    /**
     * @dev Represents the initial commission rates to be used for creating a validator.
     */
    struct CommissionRates {
        uint256 rate;
        uint256 maxRate;
        uint256 maxChangeRate;
    }

    /**
     * @dev Represents a validator description.
     */
    struct Description {
        string moniker;
        string identity;
        string website;
        string securityContact;
        string details;
    }

    /**
     * @dev Represents one entry of an unbonding delegation
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
     * @dev Represents all unbonding bonds of a single delegator with relevant metadata
     */
    struct UnbondingDelegation {
        address delegatorAddress;
        address validatorAddress;
        UnbondingDelegationEntry[] entries;
    }

    /**
     * @dev Represents a redelegation entry with relevant metadata
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

    /**
     * @dev Represents a single delegation.
     */
    struct Delegation {
        address delegator;
        // tokens
        uint256 balance;
        // shares
        uint256 shares;
    }
}
