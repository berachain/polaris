// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;

import {ERC20} from "../lib/ERC20.sol";

contract TestERC20 is ERC20 {
    constructor() ERC20("Token", "TK", 18) {}

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }
}