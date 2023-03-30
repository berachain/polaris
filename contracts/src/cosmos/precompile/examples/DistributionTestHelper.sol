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

import {IDistributionModule} from "../distribution.sol";
import {ERC20} from "../../../../lib/ERC20.sol";

/**
 * @dev This contract is an example helper for calling the distribution precompile from another contract.
 */

contract DistributionTestHelper {
    // State
    IDistributionModule public distribution;

    // Errors
    error ZeroAddress();

    /**
     * @dev Constructor that sets the distribution precompile address.
     * @param _distributionprecompile The address of the staking precompile contract.
     */
    constructor(address _distributionprecompile) {
        if (_distributionprecompile == address(0)) {
            revert ZeroAddress();
        }

        distribution = IDistributionModule(_distributionprecompile);
    }

    function getWithdrawEnabled() external view returns (bool) {
        return distribution.getWithdrawEnabled();
    }
}
