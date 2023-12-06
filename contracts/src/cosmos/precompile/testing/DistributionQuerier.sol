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

import "../Distribution.sol";

contract DistributionQuerier {
    IDistributionModule distributionModule = IDistributionModule(0x0000000000000000000000000000000000000069);

    function getTotalCall(address delegator) external view returns (Cosmos.Coin[] memory) {
        return distributionModule.getTotalDelegatorReward(delegator);
    }

    function getTotalStaticCall(address delegator) external view returns (Cosmos.Coin[] memory) {
        (bool success, bytes memory data) = address(distributionModule).staticcall(
            abi.encodeWithSelector(IDistributionModule.getTotalDelegatorReward.selector, delegator)
        );
        require(success, "call failed");
        return abi.decode(data, (Cosmos.Coin[]));
    }

    function getTotalLowLevelCall(address delegator) external returns (Cosmos.Coin[] memory) {
        (bool success, bytes memory data) = address(distributionModule).call(
            abi.encodeWithSelector(IDistributionModule.getTotalDelegatorReward.selector, delegator)
        );
        require(success, "call failed");
        return abi.decode(data, (Cosmos.Coin[]));
    }
}
