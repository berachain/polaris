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

import {IBankModule} from "../bank.sol";
import {ERC20} from "../../../../lib/ERC20.sol";

/**
 * @dev LiquidStaking is a contract that allows users to delegate their Base Denom to a validator
 * and receive a liquid staking token in return. The liquid staking token can be redeemed for Base
 * Denom at any time.
 * Note: This is an example of how to delegate Base Denom to a validator.
 * Doing it this way is unsafe since the user can delegate more straight through precomile.
 * And withdraw via the precompile.
 */
contract Fundraiser is ERC20 {
    address public owner;
    uint256 public raisedAmount = 0;
    mapping(string => uint256) public raisedFund;

    // State
    IBankModule public immutable bank = IBankModule(0x4381dC2aB14285160c808659aEe005D51255adD7);

    event Success(bool indexed success);
    event Data(bytes data);

    // Errors
    error ZeroAddress();
    error ZeroAmount();
    error InvalidValue();

    /**
     * @dev Constructor that sets the staking precompile address and the validator address.
     * @param _name The name of the token.
     * @param _symbol The symbol of the token.
     */
    constructor(string memory _name, string memory _symbol) ERC20(_name, _symbol, 18) {
        owner = msg.sender;
    }

    function withdrawDonations() external {
        require(msg.sender == owner, "Funds will only be released to the owner");
        
        payable(owner).transfer(raisedAmount);
        bank.send(this, owner, raisedAmount);
    }

    receive() external payable {
        // this built-in function doesn't require any calldata,
        // it will get called if the data field is empty and 
        // the value field is not empty.
        // this allows the smart contract to receive ether just like a 
        // regular user account controlled by a private key would.
        raisedAmount += msg.value;
    }
}
