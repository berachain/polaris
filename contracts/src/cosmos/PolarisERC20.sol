// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0;

import {ERC20} from "../../lib/ERC20.sol";
import {Owned} from "../../lib/Owned.sol";

/**
 * @notice The PolarisERC20 token is used as the ERC20 token representation of IBC-originated SDK
 * coins.
 *
 * This implementation uses the Solmate ERC20 abstract contract. Only the deployer of the contract
 * is allowed to mint and burn tokens. The default value of {decimals} is 18.
 *
 * @author Berachain Team
 * @author Solmate (https://github.com/Rari-Capital/solmate/blob/main/src/tokens/ERC20.sol)
 */
contract PolarisERC20 is Owned, ERC20 {
    /**
     * @dev Sets the values for {name} and {symbol} to _denom and Owner to msg.sender.
     * @param _denom is the corresponding SDK Coin's denom.
     *
     * All of these values are immutable: they can only be set once during construction.
     */
    constructor(string memory _denom) Owned(msg.sender) ERC20(_denom, _denom, 18) {}

    /**
     * @dev Creates `amount` tokens and assigns them to `to`, increasing the total supply.
     *
     * Emits a {Transfer} event with `from` set to the zero address.
     */
    function mint(address to, uint256 amount) external onlyOwner {
        _mint(to, amount);
    }

    /**
     * @dev Destroys `amount` tokens from `from`, reducing the total supply.
     *
     * Emits a {Transfer} event with `to` set to the zero address.
     */
    function burn(address from, uint256 amount) external onlyOwner {
        _burn(from, amount);
    }
}
