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

pragma solidity ^0.8.17;

import {IStakingModule} from "../staking.sol";
import {ERC20} from "../../../../lib/ERC20.sol";

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
    constructor(string memory _name, string memory _symbol, address _stakingprecompile, address _validatorAddress)
        ERC20(_name, _symbol, 18)
    {
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
