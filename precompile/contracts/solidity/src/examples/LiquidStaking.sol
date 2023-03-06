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

import {IStakingModule} from "../staking.sol";
import {ERC20} from "solmate/tokens/ERC20.sol";

/**
 * @dev LiquidStaking is a contract that allows users to delegate their Base Denom to a validator
 * and receive a liquid staking token in return. The liquid staking token can be redeemed for Base
 * Denom at any time.
 * Note: This is an example of how to delegate Base Denom to a validator.
 * Doing it this way is unsafe since the user can delegate more straight through precomile.
 * And withdraw via the precompile.
 */
contract LiquidStaking is ERC20 {
    // State
    IStakingModule public staking;
    address public validatorAddress;

    // Errors
    error ZeroAddress();
    error ZeroAmount();
    error InvalidValue();

    /**
     * @dev Constructor that sets the staking precompile address and the validator address.
     * @param _name The name of the token.
     * @param _symbol The symbol of the token.
     * @param _stakingprecompile The address of the staking precompile contract.
     * @param _validatorAddress The address of the validator to delegate to.
     */
    constructor(
        string memory _name,
        string memory _symbol,
        address _stakingprecompile,
        address _validatorAddress
    ) ERC20(_name, _symbol, 18) {
        if (_stakingprecompile == address(0)) revert ZeroAddress();
        if (_validatorAddress == address(0)) revert ZeroAddress();
        staking = IStakingModule(_stakingprecompile);
        validatorAddress = _validatorAddress;
    }

    /**
     * @dev Returns the total amount of assets delegated to the validator.
     * @return amount total amount of assets delegated to the validator.
     */
    function totalDelegated() public view returns (uint256 amount) {
        return staking.getDelegation(address(this), validatorAddress);
    }

    /**
     * @dev Delegates Base Denom to the validator.
     * @param amount amount of Base Denom to delegate.
     */
    function delegate(uint256 amount) public {
        if (amount == 0) revert ZeroAmount();

        // Delegate the amount to the validator.
        staking.delegate(validatorAddress, amount);
        _mint(msg.sender, amount);
    }

    /**
     * @dev Withdraws Base Denom from the validator.
     * @param amount amount of Base Denom to withdraw.
     */
    function withdraw(uint256 amount) public {
        if (amount == 0) revert ZeroAmount();
        _burn(msg.sender, amount);
        payable(msg.sender).transfer(amount);
    }
}
