// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0;

import {ERC20} from "../../lib/ERC20.sol";

/**
 * @notice The PolarisERC20 token is used as the ERC20 token representation of IBC-originated coins on
 * Cosmos SDK Polaris chains.
 *
 * @author Berachain Team
 * @author Solmate (https://github.com/Rari-Capital/solmate/blob/main/src/tokens/ERC20.sol)
 */
contract PolarisERC20 is ERC20 {
    /*//////////////////////////////////////////////////////////////
                               CONSTRUCTOR
    //////////////////////////////////////////////////////////////*/

    /// @param _denom is the corresponding SDK Coin's denom.
    constructor(string memory _denom) ERC20(_denom, _denom, 18) {}
}
