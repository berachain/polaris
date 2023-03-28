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

interface IERC20 {
    function approve(address spender, uint256 amount) external returns (bool);
    function transfer(address to, uint256 amount) external returns (bool);
    function transferFrom(address from, address to, uint256 amount) external returns (bool);
}

/**
 * @dev Interface of the erc20 module's precompiled contract
 */
interface IERC20Module {
    ////////////////////////////////////////// EVENTS /////////////////////////////////////////////

    /**
     * @dev Emitted by the erc20 module when `amount` tokens are converted from ERC20 (of address 
     * `token`) to Cosmos SDK coin (of denomination `denom`).
     */
    event TransferFromERC20ToCosmos(address indexed token, string indexed denom, uint256 amount);

    /**
     * @dev Emitted by the erc20 module when `amount` tokens are converted from Cosmos SDK coin (of 
     * denomination `denom`) to ERC20 (of address `token`).
     */
    event TransferFromCosmosToERC20(string indexed denom, address indexed token, uint256 amount);

    ////////////////////////////////////// WRITE METHODS //////////////////////////////////////////

    /**
     * @dev transferFromCosmosToERC20 converts Cosmos SDK coins to ERC20 tokens.
     * @param token the ERC20 token being transfered to
     * @param to the address to transfer to
     * @param amount the amount of tokens to transfer
     */
    function transferFromCosmosToERC20(IERC20 token, address to, uint256 amount) external;

    /**
     * @dev transferFromCosmosToERC20 converts Cosmos SDK coins to ERC20 tokens.
     * @param denom the denomination of the Cosmos SDK coin
     * @param to the address to transfer to
     * @param amount the amount of tokens to transfer
     */
    function transferFromCosmosToERC20(string calldata denom, address to, uint256 amount) external;

    /**
     * @dev transferFromERC20ToCosmos converts ERC20 tokens Cosmos SDK coins.
     * @param token the ERC20 token being transfered from
     * @param to the address to transfer to
     * @param amount the amount of tokens to transfer
     */
    function transferFromERC20ToCosmos(IERC20 token, address to, uint256 amount) external;

    /**
     * @dev transferFromERC20ToCosmos converts ERC20 tokens Cosmos SDK coins.
     * @param token the ERC20 token being transfered from
     * @param to the bech32 address to transfer to
     * @param amount the amount of tokens to transfer
     */
    function transferFromERC20ToCosmos(IERC20 token, string calldata to, uint256 amount) external;

    /**
     * @dev transferFromERC20ToCosmos converts ERC20 tokens Cosmos SDK coins.
     * @param denom the denomination of the Cosmos SDK coin
     * @param to the address to transfer to
     * @param amount the amount of tokens to transfer
     */
    function transferFromERC20ToCosmos(string calldata denom, address to, uint256 amount) external;

    /**
     * @dev transferFromERC20ToCosmos converts ERC20 tokens Cosmos SDK coins.
     * @param denom the denomination of the Cosmos SDK coin
     * @param to the bech32 address to transfer to
     * @param amount the amount of tokens to transfer
     */
    function transferFromERC20ToCosmos(string calldata denom, string calldata to, uint256 amount) external;

    /////////////////////////////////////// READ METHODS //////////////////////////////////////////

    /**
     * @dev denomForAddress returns the x/bank module denomination for the given ERC20 address.
     */
    function denomForAddress(address token) external view returns (string memory);

    /**
     * @dev denomForAddress returns the x/bank module denomination for the given ERC20 address 
     * `token` (in string bech32 format).
     */
    function denomForAddress(string calldata token) external view returns (string memory);

    /**
     * @dev addressForDenom returns the ERC20 address for the given x/bank module denomination.
     */
    function addressForDenom(string calldata denom) external view returns (address);
}
