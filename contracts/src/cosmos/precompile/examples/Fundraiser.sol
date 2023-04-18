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
import {Owned} from "../../../../lib/Owned.sol";

/**
 * @dev Fundraiser is a contract that allows users to donate tokens in any denom.
 * Only the owner can withdraw the funds.
 * Note: This is an example of how to use the bank precompile.
 */
contract Fundraiser is Owned {
    // State
    IBankModule public immutable bank = IBankModule(0x4381dC2aB14285160c808659aEe005D51255adD7);

    constructor() Owned(msg.sender) {}

    function withdrawDonations() external onlyOwner {
        require(msg.sender == owner, "Funds will only be released to the owner");
        bank.send(address(this), owner, GetRaisedAmounts());
    }

    function Donate(IBankModule.Coin[] calldata coins) external {
        bank.send(msg.sender, address(this), coins);
    }

    function GetRaisedAmounts() public view returns (IBankModule.Coin[] memory) {
        return bank.getAllBalances(address(this));
    }
}
