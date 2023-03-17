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
     * @param recipient The address of the recipient
     */
    function transferToCosmos(IERC20 token, address recipient, uint256 amount) public {
        // Transfer tokens to the Router.
        require(token.transferFrom(msg.sender, address(this), amount), "transfer failed");

        // Call the ERC20Module to handle the incoming transfer (mint bank module tokens to the user).
        (int8 handlerType, bool success) = erc20Module.handleIncoming(recipient, address(token), amount);
        require(success, "handle incoming failed");

        // If handlerType is 0 (mint), mint the tokens to the recipient.
        if (handlerType == 0) {
            PolarisERC20(address(token)).burn(amount);
        }
    }

    /**
     * @dev Transfer tokens to Cosmos
     * @param denom The denom to transfer.
     * @param recipient The address of the recipient
     * @param amount The amount of tokens to transfer
     */
    function transferToCosmos(string memory denom, address recipient, uint256 amount) public {
        address token = erc20Module.addressForDenom(denom);
        require(token != address(0), "unregistered denom");
        transferToCosmos(IERC20(token), recipient, amount);
    }

    /**
     * @dev Transfer tokens from Cosmos
     * @param denom The denom to transfer.
     * @param amount The amount of tokens to transfer
     * @param recipient The address of the recipient
     */
    function transferFromCosmos(string memory denom, address recipient, uint256 amount) public {
        address token;
        // Call the ERC20Module to handle the outgoing transfer (burn bank module tokens from the user).
        // If the ERC20Module returns true, it means that it requires that the shim deploy a new ERC20 token
        // to represent the bank module denom that we supplued.
        (int8 handlerType, bool requiresDeploy) = erc20Module.handleOutgoing(msg.sender, recipient, denom, amount);
        if (requiresDeploy) {
            // Deploy a new ERC20 token.
            token = address(new PolarisERC20(denom, denom));
            // If the ERC20Module fails to handle the post deploy request, revert.
            require(erc20Module.handleDeploy(address(token)), "handle deploy failed");
        }
        
        // If handlerType is 0 (mint), mint the tokens to the recipient.
        if (handlerType == 0) {
            PolarisERC20(token).mint(amount);
        }

        // Transfer tokens to the recipient.
        require(IERC20(token).transfer(recipient, amount), "transfer failed");
    }

    /**
     * @dev Transfer tokens from Cosmos
     * @param token The address of the token to transfer
     * @param amount The amount of tokens to transfer
     * @param recipient The address of the recipient
     */
    function transferFromCosmos(IERC20 token, address recipient, uint256 amount) public {
        string memory denom = erc20Module.denomForAddress(address(token));
        require(bytes(denom).length > 0, "unregistered token");
        transferFromCosmos(denom, recipient, amount);
    }
}
