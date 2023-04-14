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
     * `token`) to an SDK coin (of denomination `denom`).
     */
    event ConvertErc20ToCoin(address indexed token, string indexed denom, uint256 amount);

    /**
     * @dev Emitted by the erc20 module when `amount` tokens are converted from SDK coin (of
     * denomination `denom`) to ERC20 (of address `token`).
     */
    event ConvertCoinToErc20(string indexed denom, address indexed token, uint256 amount);

    /////////////////////////////////////// READ METHODS //////////////////////////////////////////

    /**
     * @dev coinDenomForERC20Address returns the SDK coin denomination for the given ERC20 address.
     */
    function coinDenomForERC20Address(IERC20 token) external view returns (string memory);

    /**
     * @dev coinDenomForERC20Address returns the SDK coin denomination for the given ERC20 address
     * `token` (in string bech32 format).
     */
    function coinDenomForERC20Address(string calldata token) external view returns (string memory);

    /**
     * @dev erc20AddressForCoinDenom returns the ERC20 address for the given SDK coin denomination.
     */
    function erc20AddressForCoinDenom(string calldata denom) external view returns (IERC20);

    ////////////////////////////////////// WRITE METHODS //////////////////////////////////////////

    /**
     * @dev convertCoinToERC20 converts `amount` SDK coins to ERC20 tokens for `owner`.
     * @param denom the denomination of the SDK coin being converted from
     * @param owner the account to convert for
     * @param amount the amount of tokens to convert
     */
    function convertCoinToERC20(string calldata denom, address owner, uint256 amount) external returns (bool);

    /**
     * @dev convertCoinToERC20 converts `amount` SDK coins to ERC20 tokens for `owner`.
     * @param denom the denomination of the SDK coin being converted from
     * @param owner the account to convert for (bech32 address)
     * @param amount the amount of tokens to convert
     */
    function convertCoinToERC20(string calldata denom, string calldata owner, uint256 amount) external returns (bool);

    /**
     * @dev convertERC20ToCoin converts `amount` ERC20 tokens to SDK coins for `owner`.
     * @param token the ERC20 token being converted from
     * @param owner the account to convert for
     * @param amount the amount of tokens to transfer
     */
    function convertERC20ToCoin(IERC20 token, address owner, uint256 amount) external returns (bool);

    /**
     * @dev convertERC20ToCoin converts `amount` ERC20 tokens to SDK coins for `owner`.
     * @param token the ERC20 token being transfered from
     * @param owner the account to convert for (bech32 address)
     * @param amount the amount of tokens to transfer
     */
    function convertERC20ToCoin(IERC20 token, string calldata owner, uint256 amount) external returns (bool);
}
