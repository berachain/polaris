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

import {IDistributionModule} from "../cosmos/precompile/Distribution.sol";
import {IStakingModule} from "../cosmos/precompile/Staking.sol";
import {ERC20} from "../../lib/ERC20.sol";

/**
 * @dev This contract is an example helper for calling the distribution precompile from another contract.
 */
contract DistributionWrapper {
    // State
    IDistributionModule public distribution;
    IStakingModule public staking;

    // Errors
    error ZeroAddress();

    /**
     * @dev Constructor that sets the distribution precompile address.
     * @param _distributionprecompile The address of the staking precompile contract.
     * @param _stakingprecompile The address of the staking precompile contract.
     */
    constructor(address _distributionprecompile, address _stakingprecompile) {
        if (_distributionprecompile == address(0) && _stakingprecompile == address(0)) {
            revert ZeroAddress();
        }

        distribution = IDistributionModule(_distributionprecompile);
        staking = IStakingModule(_stakingprecompile);
    }

    function getWithdrawEnabled() external view returns (bool) {
        return distribution.getWithdrawEnabled();
    }

    /**
     * @dev The caller (msg.sender) can set the address that will receive the delegation rewards.
     * @param _withdrawAddress The address to set as the withdraw address.
     */
    function setWithdrawAddress(address _withdrawAddress) external returns (bool) {
        return distribution.setWithdrawAddress(_withdrawAddress);
    }

    /**
     * @dev Withdraw the rewrads accumulated by the caller(msg.sender).
     * @param _delegatorAddress The address of the delegator.
     * @param _validatorAddress The address of the validator.
     */
    function withdrawRewards(address _delegatorAddress, address _validatorAddress) external {
        distribution.withdrawDelegatorReward(_delegatorAddress, _validatorAddress);
    }

    /**
     * @dev msg.sender delegates the `msg.value` of tokens to `_validator`
     * @param _validator The address of the validator.
     */
    function delegate(address _validator) external payable {
        staking.delegate(_validator, msg.value);
    }
}
