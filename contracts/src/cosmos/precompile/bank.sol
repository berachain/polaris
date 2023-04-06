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

/**
 * @dev Interface of all supported Cosmos events emitted by the bank module
 */
interface IBankModule {
    ////////////////////////////////////////// EVENTS /////////////////////////////////////////////
    
    /**
     * @dev Emitted by the bank module when `amount` tokens are sent to `recipient`
     */
    event Transfer(address indexed recipient, uint256 amount);

    /**
     * @dev Emitted by the bank module when `sender` sends some amount of tokens
     */
    event Message(address indexed sender);

    /**
     * @dev Emitted by the bank module when `amount` tokens are spent by `spender`
     */
    event CoinSpent(address indexed spender, uint256 amount);

    /**
     * @dev Emitted by the bank module when `amount` tokens are received by `receiver`
     */
    event CoinReceived(address indexed receiver, uint256 amount);

    /**
     * @dev Emitted by the bank module when `amount` tokens are minted by `minter`
     *
     * Note: "Coinbase" refers to the Cosmos event: EventTypeCoinMint. `minter` is a module
     * address.
     */
    event Coinbase(address indexed minter, uint256 amount);

    /**
     * @dev Emitted by the bank module when `amount` tokens are burned by `burner`
     *
     * Note: `burner` is a module address
     */
    event Burn(address indexed burner, uint256 amount);
    
    /////////////////////////////////////// READ METHODS //////////////////////////////////////////

    /**
     * @dev Returns the `amount` of account balance by address for a given denomination.
     */
    function getBalance(address accountAddress, string calldata denom) external view returns (uint256);

    /**
     * @dev Returns account balance by address for all denominations.
     */
    function getAllBalance(address accountAddress) external view returns (Coin[] memory);

    /**
     * @dev Returns the `amount` of account balance by address for a given denomination.
     */
    function getSpendableBalanceByDenom(address accountAddress, string calldata denom) external view returns (uint256);

    /**
     * @dev Returns account balance by address for all denominations.
     */
    function getSpendableBalances(address accountAddress) external view returns (Coin[] memory);

    /**
     * @dev Returns the total supply of a single coin.
     */
    function getSupplyOf(string calldata denom) external view returns (uint256);

    /**
     * @dev Returns the total supply of a all coins.
     */
    function getTotalSupply() external view returns (Coin[] memory);

    /**
     * @dev Returns the parameters of the bank module.
     */
    function getParams() external view returns (Param memory);

    /**
     * @dev Returns the denomination metadata
     */
    function getDenomMetadata(string calldata denom) external view returns (DenomMetadata memory);

    /**
     * @dev Returns all denominations metadata
     */
    function getDenomsMetadata() external view returns (DenomsMetadata memory);

    /**
     * @dev Returns if the denom is enabled to send
     */
    function getSendEnabled(string[] calldata denoms) external view returns (SendEnabled memory);

    ////////////////////////////////////// WRITE METHODS //////////////////////////////////////////

    /**
     * @dev Send coins from one address to another.
     */
    function send(address fromAddress, address toAddress, Coin calldata amount) external payable returns (bool);

    /**
     * @dev Send coins from one sender and to a series of different address. 
     * If any of the receiving addresses do not correspond to an existing account, 
     * a new account is created.
     * 
     * Inputs, despite being `repeated`, only allows one sender input. 
     */
    function multiSend(Input calldata input, Output[] memory outputs) external payable returns (bool);

    //////////////////////////////////////////// UTILS ////////////////////////////////////////////

    /**
     * @dev Represents a cosmos coin.
     * Note: this struct is generated as go struct that is then used in the precompile.
     */
    struct Coin {
        uint256 amount;
        string denom;
    }

    struct Param {
        bool defaultSendEnabled;
    }

    struct DenomUnit {
        string denom;
        string[] aliases;
        uint32 exponent;
    }
    struct DenomMetadata {
        string description;
        DenomUnit[] denomUnits;
        string base;
        string display;
        string name;
        string symbol;
    }

    struct DenomsMetadata {
        DenomMetadata[] metadatas;
    }

    struct SendEnabled {
        string denom;
        bool enabled;
    }

    struct Input {
        address addr;
        Coin[] coins;
    }
    
    struct Output {
        address addr;
        Coin[] coins;
    }
}
