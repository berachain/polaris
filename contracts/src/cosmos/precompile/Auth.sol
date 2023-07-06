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

import {Cosmos} from "../CosmosTypes.sol";

/**
 * @dev Interface of the auth module precompiled contract
 */
interface IAuthModule {
    /**
     * @dev Returns the base account information for the given account address.
     */
    function getAccountInfo(
        address account
    ) external view returns (BaseAccount memory);

    /**
     * @dev setSendAllowance sets the send authorization (allowance) between owner and spender.
     * @param owner the account approving the allowance
     * @param spender the account being granted the allowance
     * @param amount the Coins of the allowance
     * @param expiration the expiration time of the grant (0 means no expiration)
     */
    function setSendAllowance(
        address owner,
        address spender,
        Cosmos.Coin[] calldata amount,
        uint256 expiration
    ) external returns (bool);

    /**
     * @dev getSendAllowance returns the send authorization (allowance) amount between owner and
     * spender.
     * @param owner the account that approved the allowance
     * @param spender the account that was granted the allowance
     * @param denom the denomination of the Coin that was allowed
     */
    function getSendAllowance(
        address owner,
        address spender,
        string calldata denom
    ) external view returns (uint256);

    //////////////////////////////////////////// UTILS ////////////////////////////////////////////

    /**
     * @dev Represents a Cosmos base account.
     */
    struct BaseAccount {
        address addr; // equivalent to the Address field of authtypes.BaseAccount
        bytes pubKey;
        uint64 accountNumber;
        uint64 sequence; // represents account nonce
    }
}
