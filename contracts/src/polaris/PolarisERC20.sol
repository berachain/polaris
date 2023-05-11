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

import {ERC20} from "../../lib/ERC20.sol";
import {Owned} from "../../lib/Owned.sol";

/**
 * @dev Implementation of the {IERC20} interface.
 *
 * The PolarisERC20 token is used as the ERC20 token representation of IBC-originated SDK coins.
 *
 * This implementation uses the Solmate ERC20 abstract contract. Only the deployer of the contract
 * is allowed to mint and burn tokens. The default value of {decimals} is 18.
 */
contract PolarisERC20 is Owned, ERC20 {
    /**
     * @dev Sets the values for {name} and {symbol} and Owner to msg.sender.
     *
     * All of these values are immutable: they can only be set once during construction.
     */
    constructor(string memory name, string memory symbol) Owned(msg.sender) ERC20(name, symbol, 18) {}

    /**
     * @dev Creates `amount` tokens and assigns them to `to`, increasing the total supply.
     *
     * Emits a {Transfer} event with `from` set to the zero address.
     */
    function mint(address to, uint256 amount) external onlyOwner {
        _mint(to, amount);
    }

    /**
     * @dev Destroys `amount` tokens from `from`, reducing the total supply.
     *
     * Emits a {Transfer} event with `to` set to the zero address.
     */
    function burn(address from, uint256 amount) external onlyOwner {
        _burn(from, amount);
    }
}
