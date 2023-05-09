// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0;

import {IERC20} from "../../lib/IERC20.sol";
import {IBankModule} from "./precompile/Bank.sol";
import {IERC20Module} from "./precompile/ERC20Module.sol";

abstract contract ERC20 is IERC20 {
    /*//////////////////////////////////////////////////////////////
                              Precompiles
    //////////////////////////////////////////////////////////////*/

    /**
     * @dev bank is a pure function for getting the address of the bank module precompile.
     * @return IBankModule the address of the bank module precompile.
     */
    function bank() internal pure returns (IBankModule) {
        return IBankModule(address(0x1));
    }

    /**
     * @dev erc20Module is a pure function for getting the address of the erc20 module precompile.
     * @return IERC20Module the address of the erc20 module precompile.
     */
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

    mapping(address => mapping(address => uint256)) public allowance;

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

    /*//////////////////////////////////////////////////////////////
                               CONSTRUCTOR
    //////////////////////////////////////////////////////////////*/

    constructor(string memory _denom) {
        denom = _denom;
    }

    /*//////////////////////////////////////////////////////////////
                               ERC20 LOGIC
    //////////////////////////////////////////////////////////////*/

    /**
     * @dev approve is a public method for approving a given address to spend a given amount of tokens.
     * @param spender the address to approve to spend tokens.
     * @param amount the amount of tokens to approve the given address to spend.
     * @return bool true if the approval was successful.
     */
    function approve(address spender, uint256 amount) public virtual returns (bool) {
        allowance[msg.sender][spender] = amount;

        emit Approval(msg.sender, spender, amount);

        return true;
    }

    /**
     * @dev transfer is a public method for transferring tokens to a given address.
     * @param to the address to transfer tokens to.
     * @param amount the amount of tokens to transfer.
     * @return bool true if the transfer was successful.
     */
    function transfer(address to, uint256 amount) public virtual returns (bool) {
        IBankModule.Coin[] memory coins = amountToCoins(amount);
        require(bank().send(msg.sender, to, coins), "PolarisERC20: failed to send tokens");

        emit Transfer(msg.sender, to, amount);
        return true;
    }

    /**
     * @dev transferFrom is a public method for transferring tokens from one address to another.
     * @param from the address to transfer tokens from.
     * @param to the address to transfer tokens to.
     * @param amount the amount of tokens to transfer.
     * @return bool true if the transfer was successful.
     */
    function transferFrom(address from, address to, uint256 amount) public virtual returns (bool) {
        // TODO: Use allowance once authz precompile is available.
        uint256 allowed = allowance[from][msg.sender]; // Saves gas for limited approvals.

        if (allowed != type(uint256).max) allowance[from][msg.sender] = allowed - amount;

        require(bank().send(from, to, amountToCoins(amount)), "PolarisERC20: failed to send bank tokens");

        emit Transfer(from, to, amount);
        return true;
    }

    /*//////////////////////////////////////////////////////////////
                               sdk.Coin helpers.
    //////////////////////////////////////////////////////////////*/

    /**
     * @dev amountToCoins is a helper function to convert an amount to sdk.Coin.
     * @param amount the amount to convert to sdk.Coin.
     * @return sdk.Coin[] the sdk.Coin representation of the given amount.
     */
    function amountToCoins(uint256 amount) internal view returns (IBankModule.Coin[] memory) {
        IBankModule.Coin[] memory coins = new IBankModule.Coin[](1);
        coins[0] = IBankModule.Coin({denom: denom, amount: amount});
        return coins;
    }
}
