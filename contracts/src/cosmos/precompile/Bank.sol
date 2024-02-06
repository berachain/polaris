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

pragma solidity 0.8.23;

import {Cosmos} from "../CosmosTypes.sol";

/**
 * @dev Interface of all supported Cosmos events emitted by the bank module
 */
interface IBankModule {
    ////////////////////////////////////////// EVENTS /////////////////////////////////////////////

    /**
     * @dev Emitted by the bank module when `amount` tokens are sent to `recipient`
     * @param recipient The recipient address
     * @param amount The amount of Cosmos coins sent
     */
    event Transfer(address indexed recipient, Cosmos.Coin[] amount);

    /**
     * @dev Emitted by the bank module when `sender` sends some amount of tokens
     * @param sender The sender address
     */
    event Message(address indexed sender);

    /**
     * @dev Emitted by the bank module when `amount` tokens are spent by `spender`
     * @param spender The spender address
     * @param amount The amount of Cosmos coins spent
     */
    event CoinSpent(address indexed spender, Cosmos.Coin[] amount);

    /**
     * @dev Emitted by the bank module when `amount` tokens are received by `receiver`
     * @param receiver The receiver address
     * @param amount The amount of Cosmos coins received
     */
    event CoinReceived(address indexed receiver, Cosmos.Coin[] amount);

    /**
     * @dev Emitted by the bank module when `amount` tokens are minted by `minter`
     * @param minter The minter address
     * @param amount The amount of Cosmos coins minted
     * @notice "Coinbase" refers to the Cosmos event: `EventTypeCoinMint`
     * @notice `minter` is always a module address
     */
    event Coinbase(address indexed minter, Cosmos.Coin[] amount);

    /**
     * @dev Emitted by the bank module when `amount` tokens are burned by `burner`
     * @param burner The burner address
     * @param amount The amount of Cosmos coins burned
     * @notice `burner` is always a module address
     */
    event Burn(address indexed burner, Cosmos.Coin[] amount);

    /////////////////////////////////////// READ METHODS //////////////////////////////////////////

    /**
     * @dev Returns the `amount` of account balance by address for a given coin denomination
     * @notice If the denomination is not found, returns 0
     */
    function getBalance(address accountAddress, string calldata denom) external view returns (uint256);

    /**
     * @dev Returns account balance by address for all denominations
     * @notice If the account address is not found, returns an empty array
     */
    function getAllBalances(address accountAddress) external view returns (Cosmos.Coin[] memory);

    /**
     * @dev Returns the `amount` of account balance by address for a given coin denomination
     * @notice If the denomination is not found, returns 0
     */
    function getSpendableBalance(address accountAddress, string calldata denom) external view returns (uint256);

    /**
     * @dev Returns account balance by address for all coin denominations
     * @notice If the account address is not found, returns an empty array
     */
    function getAllSpendableBalances(address accountAddress) external view returns (Cosmos.Coin[] memory);

    /**
     * @dev Returns the total supply of a single coin
     */
    function getSupply(string calldata denom) external view returns (uint256);

    /**
     * @dev Returns the total supply of a all coins
     */
    function getAllSupply() external view returns (Cosmos.Coin[] memory);

    ////////////////////////////////////// WRITE METHODS //////////////////////////////////////////

    /**
     * @dev Send coins from msg.sender to another
     * @param toAddress The recipient address
     * @param amount The amount of Cosmos coins to send
     * @notice If the sender does not have enough balance, returns false
     */
    function send(address toAddress, Cosmos.Coin[] calldata amount) external payable returns (bool);

    //////////////////////////////////////////// UTILS ////////////////////////////////////////////

    /**
     * @dev Represents a denom unit in the bank module
     * @notice this struct is generated in generated/i_bank_module.abigen.go
     */
    struct DenomUnit {
        string denom;
        string[] aliases;
        uint32 exponent;
    }

    /**
     * @dev Represents a denom metadata in the bank module
     * @notice this struct is generated in generated/i_bank_module.abigen.go
     */
    struct DenomMetadata {
        string description;
        DenomUnit[] denomUnits;
        string base;
        string display;
        string name;
        string symbol;
    }
}
