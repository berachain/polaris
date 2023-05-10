// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0;

import {IERC20} from "../../lib/IERC20.sol";
import {IAuthModule} from "./precompile/Auth.sol";
import {IBankModule} from "./precompile/Bank.sol";

/**
 * @notice Polaris implementation of ERC20 + EIP-2612.
 *
 * The PolarisERC20 token is used as the ERC20 token representation of IBC-originated coins on
 * Cosmos SDK Polaris chains. Uses the bank module to actually hold account balances and execute
 * transfers. The authz module is used to set approvals and permissions for spends.
 *
 * @author Berachain Team
 * @author Solmate (https://github.com/Rari-Capital/solmate/blob/main/src/tokens/ERC20.sol)
 */
contract PolarisERC20 is IERC20 {
    /*//////////////////////////////////////////////////////////////
                              ERC20 STORAGE
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
                            EIP-2612 STORAGE
    //////////////////////////////////////////////////////////////*/

    uint256 internal immutable INITIAL_CHAIN_ID;

    bytes32 internal immutable INITIAL_DOMAIN_SEPARATOR;

    mapping(address => uint256) public nonces;

    /*//////////////////////////////////////////////////////////////
                               CONSTRUCTOR
    //////////////////////////////////////////////////////////////*/

    /// @param _denom is the corresponding SDK Coin's denom.
    constructor(string memory _denom) {
        denom = _denom;

        INITIAL_CHAIN_ID = block.chainid;
        INITIAL_DOMAIN_SEPARATOR = computeDomainSeparator();
    }

    /*//////////////////////////////////////////////////////////////
                               ERC20 LOGIC
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

    /**
     * @dev approve is a public method for approving a given address to spend a given amount of tokens.
     * @param spender the address to approve to spend tokens.
     * @param amount the amount of tokens to approve the given address to spend.
     * @return bool true if the approval was successful.
     */
    function approve(address spender, uint256 amount) public virtual returns (bool) {
        require(
            authz().setSendAllowance(msg.sender, spender, amountToAuthCoins(amount), 0),
            "PolarisERC20: failed to approve spend"
        );

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
        require(bank().send(msg.sender, to, amountToBankCoins(amount)), "PolarisERC20: failed to send tokens");

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
        require(amount <= authz().getSendAllowance(from, msg.sender, denom), "PolarisERC20: insufficient approval");
        require(bank().send(from, to, amountToBankCoins(amount)), "PolarisERC20: failed to send bank tokens");

        emit Transfer(from, to, amount);
        return true;
    }

    /*//////////////////////////////////////////////////////////////
                             EIP-2612 LOGIC
    //////////////////////////////////////////////////////////////*/

    function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s)
        public
        virtual
    {
        require(deadline >= block.timestamp, "PolarisERC20: PERMIT_DEADLINE_EXPIRED");

        // Unchecked because the only math done is incrementing
        // the owner's nonce which cannot realistically overflow.
        unchecked {
            address recoveredAddress = ecrecover(
                keccak256(
                    abi.encodePacked(
                        "\x19\x01",
                        DOMAIN_SEPARATOR(),
                        keccak256(
                            abi.encode(
                                keccak256(
                                    "Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)"
                                ),
                                owner,
                                spender,
                                value,
                                nonces[owner]++,
                                deadline
                            )
                        )
                    )
                ),
                v,
                r,
                s
            );

            require(recoveredAddress != address(0) && recoveredAddress == owner, "PolarisERC20: INVALID_SIGNER");

            require(
                authz().setSendAllowance(recoveredAddress, spender, amountToAuthCoins(value), 0),
                "PolarisERC20: failed to approve spend"
            );
        }

        emit Approval(owner, spender, value);
    }

    function DOMAIN_SEPARATOR() public view virtual returns (bytes32) {
        return block.chainid == INITIAL_CHAIN_ID ? INITIAL_DOMAIN_SEPARATOR : computeDomainSeparator();
    }

    function computeDomainSeparator() internal view virtual returns (bytes32) {
        return keccak256(
            abi.encode(
                keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"),
                keccak256(bytes(name())),
                keccak256("1"),
                block.chainid,
                address(this)
            )
        );
    }

    /*//////////////////////////////////////////////////////////////
                              SDK HELPERS
    //////////////////////////////////////////////////////////////*/

    /**
     * @dev bank is a pure function for getting the address of the bank module precompile.
     * @return IBankModule the address of the bank module precompile.
     */
    function bank() internal pure returns (IBankModule) {
        return IBankModule(address(0x4381dC2aB14285160c808659aEe005D51255adD7));
    }

    /**
     * @dev authz is a pure function for getting the address of the auth(z) module precompile.
     * @return IAuthModule the address of the auth(z) module precompile.
     */
    function authz() internal pure returns (IAuthModule) {
        return IAuthModule(address(0xBDF49C3C3882102fc017FFb661108c63a836D065));
    }

    /**
     * @dev amountToBankCoins is a helper function to convert an amount to sdk.Coin.
     * @param amount the amount to convert to sdk.Coin.
     * @return sdk.Coin[] the sdk.Coin representation of the given amount.
     */
    function amountToBankCoins(uint256 amount) internal view returns (IBankModule.Coin[] memory) {
        IBankModule.Coin[] memory coins = new IBankModule.Coin[](1);
        coins[0] = IBankModule.Coin({denom: denom, amount: amount});
        return coins;
    }

    /**
     * @dev amountToAuthCoins is a helper function to convert an amount to sdk.Coin.
     * @param amount the amount to convert to sdk.Coin.
     * @return sdk.Coin[] the sdk.Coin representation of the given amount.
     */
    function amountToAuthCoins(uint256 amount) internal view returns (IAuthModule.Coin[] memory) {
        IAuthModule.Coin[] memory coins = new IAuthModule.Coin[](1);
        coins[0] = IAuthModule.Coin({denom: denom, amount: amount});
        return coins;
    }
}
