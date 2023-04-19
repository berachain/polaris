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

import {IERC20} from "../../../../lib/IERC20.sol";
import {IERC20Module} from "../ERC20Module.sol";

// Swapper is an example smart contract that uses the erc20 module precompile to swap/convert
// between SDK coins and ERC20 tokens.
contract Swapper {
    IERC20Module public immutable erc20Module = IERC20Module(0x0000000000000000000000000000000000696969);

    // converts ERC20 --> SDK coin
    // owner must first grant this contract to spend owner's tokens if the token is originally
    // an ERC20 token
    function swap(IERC20 token, uint256 amount) external {
        bool converted = erc20Module.convertERC20ToCoin(token, msg.sender, amount);
        require(converted, "Swapper: convertERC20ToCoin failed");
    }

    // converts SDK coin --> ERC20
    function swap(string calldata denom, uint256 amount) external {
        bool converted = erc20Module.convertCoinToERC20(denom, msg.sender, amount);
        require(converted, "Swapper: convertCoinToERC20 failed");
    }

    // gets the Polaris ERC20 token for a given SDK coin denomination
    function getPolarisERC20(string calldata denom) external view returns (IERC20) {
        return erc20Module.erc20AddressForCoinDenom(denom);
    }
}
