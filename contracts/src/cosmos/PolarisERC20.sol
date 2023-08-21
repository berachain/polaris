// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0;

import {IERC20} from "../../lib/IERC20.sol";
import {IAuthModule} from "./precompile/Auth.sol";
import {IBankModule} from "./precompile/Bank.sol";
import {IERC20Module} from "./precompile/ERC20Module.sol";
import {Cosmos} from "./CosmosTypes.sol";

/**
 * @notice Polaris implementation of ERC20 + EIP-2612.
 *
 * The PolarisERC20 token is used as the ERC20 token representation of IBC-originated coins on
 * Cosmos SDK Polaris chains. Uses the bank module to actually hold account balances and execute
 * transfers. Approvals are handled by the ERC20 contract itself and NOT the `authz` module.
 *
 * @author Berachain Team
 * @author Solmate (https://github.com/Rari-Capital/solmate/blob/main/src/tokens/ERC20.sol)
 */
contract PolarisERC20 is IERC20 {
    /*//////////////////////////////////////////////////////////////
                              ERC20 STORAGE
    //////////////////////////////////////////////////////////////*/

    string public denom;
    mapping(address => mapping(address => uint256)) public allowance;

    /**
     * @dev name is a public view method for reading the `sdk.Coin` name for this erc20.
     * @return string the sdk.Coin name for this erc20.
     */
    function name() public view returns (string memory) {
        // TODO: Get the name/display from the denom metadata.
        return denom;
    }

    /**
     * @dev symbol is a public view method for reading the `sdk.Coin` symbol for this erc20.
     * @return string the sdk.Coin symbol for this erc20.
     */
    function symbol() public view returns (string memory) {
        // TODO: Get the symbol from the denom metadata.
        return denom;
    }

    /**
     * @dev decimals is a public view method for reading the `sdk.Coin` decimals for this erc20.
     * @return uint8 the sdk.Coin decimals for this erc20.
     */
    function decimals() public pure returns (uint8) {
        // TODO: Get the max decimals from the denom units.
        return 18;
    }

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
        require(erc20Module().performBankTransfer(msg.sender, to, amount), "PolarisERC20: failed to send tokens");

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
        uint256 allowed = allowance[from][msg.sender]; // Saves gas for limited approvals.

        if (allowed != type(uint256).max) {
            allowance[from][msg.sender] = allowed - amount;
        }

        require(erc20Module().performBankTransfer(msg.sender, to, amount), "PolarisERC20: failed to send bank tokens");

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
        require(deadline >= block.timestamp, "PERMIT_DEADLINE_EXPIRED");

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

            require(recoveredAddress != address(0) && recoveredAddress == owner, "INVALID_SIGNER");

            allowance[recoveredAddress][spender] = value;
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
                keccak256(bytes(denom)),
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
     * @dev erc20Module is a pure function for getting the address of the erc20 module precompile.
     * @return IERC20Module the address of the bank module precompile.
     */
    function erc20Module() internal pure returns (IERC20Module) {
        return IERC20Module(address(0x696969));
    }
}
