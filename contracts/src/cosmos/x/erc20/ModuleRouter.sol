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

pragma solidity ^0.8.19;

import {IERC20Module} from "../../precompile/erc20.sol";
import {PolarisERC20} from "./PolarisERC20.sol";


interface IERC20 {
    function transfer(address recipient, uint256 amount) external returns (bool);
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
}

// The ERC20ModuleRouter is a contract that routes calls to the
// ERC20Module. It is responsible managing the state of the ERC20
contract ERC20ModuleRouter {
    IERC20Module public erc20Module;


    mapping(string => address) public denomToToken;
    mapping(address => string) public tokenToDenom;

    /**
     * @dev Constructor function
     * @param _erc20Module The address of the ERC20Module
     */
    constructor(IERC20Module _erc20Module) {
        erc20Module = _erc20Module;
    }

    /**
     * @dev Transfer tokens to Cosmos
     * @param token The address of the token to transfer
     * @param amount The amount of tokens to transfer
     * @param receiver The address of the receiver
     */
    function transferToCosmos(IERC20 token, address receiver, uint256 amount) external {
        // Transfer tokens to the ERC20Module.
        token.transferFrom(msg.sender, address(this), amount);

        // If the token is not registered, register it.
        if (keccak256(abi.encodePacked(tokenToDenom[address(token)])) == keccak256(abi.encodePacked(""))) {
            string memory denom = string(abi.encode("polaris/", address(token)));
            tokenToDenom[address(token)] = denom;
            denomToToken[denom] = address(token);
        }

        // Call the ERC20Module to handle the incoming transfer (mint bank module tokens to the user).
        erc20Module.handleIncoming(receiver, address(token), amount);
    }

    /**
     * @dev Transfer tokens from Cosmos
     * @param token The address of the token to transfer
     * @param amount The amount of tokens to transfer
     * @param receiver The address of the receiver    
     */
    function transferFromCosmos(IERC20 token, uint256 amount, address receiver) external {
        // Call the ERC20Module to handle the outgoing transfer (burn bank module tokens from the user).
        erc20Module.handleOutgoing(msg.sender, receiver, address(token), amount);

        // Transfer tokens to the receiver.
        token.transfer(receiver, amount);
    }

    /**
     * @dev Transfer tokens from Cosmos
     * @param denom The denom of the token to transfer
     * @param amount The amount of tokens to transfer
     * @param receiver The address of the receiver    
     */
    function transferFromCosmos(string memory denom, uint256 amount, address receiver) external {
        // Call the ERC20Module to handle the outgoing transfer.
        (address te, bool needsDeploy) = erc20Module.handleOutgoingString(msg.sender, receiver, denom, amount);

        // // Get or deploy the ERC20 contract.
        address token = denomToToken[denom];
        // if (token) {
        //     token = address(new PolarisERC20(denom, denom, 18));
        //     denomToToken[denom] = token;
        //     tokenToDenom[token] = denom;
        // }

        // Transfer tokens to the receiver.
        IERC20(token).transfer(receiver, amount);
    }
}