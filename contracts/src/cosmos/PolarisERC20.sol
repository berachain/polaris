// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0;

import {IERC20} from "../../lib/IERC20.sol";
import {IBankModule} from "./precompile/Bank.sol";
import {IERC20Module} from "./precompile/ERC20Module.sol";

abstract contract ERC20 is IERC20 {
    /*//////////////////////////////////////////////////////////////
                              Precompiles
    //////////////////////////////////////////////////////////////*/

    function bank() internal pure returns (IBankModule) {
        return IBankModule(address(0x1));
    }

    function erc20Module() internal pure returns (IERC20Module) {
        return IERC20Module(address(0x0));
    }

    /*//////////////////////////////////////////////////////////////
                                 EVENTS
    //////////////////////////////////////////////////////////////*/

    event Transfer(address indexed from, address indexed to, uint256 amount);

    event Approval(address indexed owner, address indexed spender, uint256 amount);

    /*//////////////////////////////////////////////////////////////
                            METADATA STORAGE
    //////////////////////////////////////////////////////////////*/

    /**
     * @dev name is a public view method for reading the `sdk.Coin` name for this erc20.
     * @return string the sdk.Coin name for this erc20.
     */
    function name() public view returns (string memory) {
        return bank().getDenomMetadata(denom).display;
    }

    /**
     * @dev symbol is a public view method for reading the `sdk.Coin` symbol for this erc20.
     * @return string the sdk.Coin symbol for this erc20.
     */
    function symbol() public view returns (string memory) {
        return bank().getDenomMetadata(denom).symbol;
    }

    /**
     * @dev decimals is a public view method for reading the `sdk.Coin` decimals for this erc20.
     * @return uint8 the sdk.Coin decimals for this erc20.
     */
    function decimals() public view returns (uint8) {
        return uint8(bank().getDenomMetadata(denom).denomUnits[0].exponent);
    }

    string public denom;

    /*//////////////////////////////////////////////////////////////
                              ERC20 STORAGE
    //////////////////////////////////////////////////////////////*/

    /**
     * @dev totalSupply is a public view method for reading the `sdk.Coin` total supply for this erc20.
     * @return uint256 the sdk.Coin total supply for this erc20.
     */
    function totalSupply() public view returns (uint256) {
        return bank().getSupply(denom);
    }

    /**
     * @dev balanceOf is a public view method for reading the `sdk.Coin` balance of a given address for this erc20.
     * @param user the address of the user to get the balance of.
     * @return uint256 the sdk.Coin balance of the given address for this erc20.
     */
    function balanceOf(address user) public view returns (uint256) {
        return bank().getSpendableBalance(user, denom);
    }

    //TODO:
    mapping(address => mapping(address => uint256)) public allowance;

    /*//////////////////////////////////////////////////////////////
                               CONSTRUCTOR
    //////////////////////////////////////////////////////////////*/

    constructor(string memory _denom) {
        denom = _denom;
    }

    /*//////////////////////////////////////////////////////////////
                               ERC20 LOGIC
    //////////////////////////////////////////////////////////////*/

    //TODO:
    function approve(address spender, uint256 amount) public virtual returns (bool) {
        allowance[msg.sender][spender] = amount;

        emit Approval(msg.sender, spender, amount);

        return true;
    }

    function transfer(address to, uint256 amount) public virtual returns (bool) {
        IBankModule.Coin[] memory coins = amountToCoins(amount);
        bank().send(msg.sender, to, coins);

        emit Transfer(msg.sender, to, amount);
        return true;
    }

    function transferFrom(address from, address to, uint256 amount) public virtual returns (bool) {
        // TODO: Use allowance once authz precompile is available.
        uint256 allowed = allowance[from][msg.sender]; // Saves gas for limited approvals.

        if (allowed != type(uint256).max) allowance[from][msg.sender] = allowed - amount;

        bank().send(from, to, amountToCoins(amount));

        emit Transfer(from, to, amount);
        return true;
    }

    /*//////////////////////////////////////////////////////////////
                               sdk.Coin helpers.
    //////////////////////////////////////////////////////////////*/

    function amountToCoins(uint256 amount) internal view returns (IBankModule.Coin[] memory) {
        IBankModule.Coin[] memory coins = new IBankModule.Coin[](1);
        coins[0] = IBankModule.Coin({denom: denom, amount: amount});
        return coins;
    }
}
