// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0;

import {IAuthModule} from "./precompile/Auth.sol";
import {IBankModule} from "./precompile/Bank.sol";
import {OFTCore} from "../../lib/layerzero-contracts/contracts/token/oft/OFTCore.sol";
import {IOFT} from "../../lib/layerzero-contracts/contracts/token/oft/IOFT.sol";

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
contract PolarisERC20 is OFTCore, IOFT {
    /*//////////////////////////////////////////////////////////////
                              ERC20 STORAGE
    //////////////////////////////////////////////////////////////*/

    string public denom;

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
        // TODO: Get the max decimals from the denom units? denomUnits[0] is not necessarily correct.
        return uint8(bank().getDenomMetadata(denom).denomUnits[0].exponent);
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
    // TODO: lzEndpointAddress should not be 0.
    constructor(string memory _denom) OFTCore(address(0x0)) {
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

    function allowance(address owner, address spender) public view virtual returns (uint256) {
        return authz().getSendAllowance(owner, spender, denom);
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
                            LAYER ZERO LOGIC
    //////////////////////////////////////////////////////////////*/

    /**
     * @dev returns the circulating amount of tokens on current chain
     */
    function circulatingSupply() external view override returns (uint256) {
        return totalSupply();
    }

    /**
     * @dev returns the address of the ERC20 token
     */
    function token() external view override returns (address) {
        return address(this);
    }

    /**
     * @dev Debits an amount from a specific address and sends it to another chain
     * @notice This function should only be called internally
     * @param _from The address to debit from
     * @param _dstChainId The ID of the destination chain where the amount will be sent
     * @param _toAddress The address (in bytes) that will receive the debited amount on the destination chain
     * @param _amount The amount to debit from the address
     * @return Returns the transaction status as a uint (e.g., 0 for failure, 1 for success)
     */
    function _debitFrom(address _from, uint16 _dstChainId, bytes memory _toAddress, uint256 _amount)
        internal
        virtual
        override
        returns (uint256)
    {
        // TODO: Implement
        _from;
        _dstChainId;
        _toAddress;
        _amount;
        return 0;
    }

    /**
     * @dev Credits an amount to a specific address from a source chain
     * @notice This function should only be called internally
     * @param _srcChainId The ID of the source chain from where the amount is sent
     * @param _toAddress The address that will receive the credited amount
     * @param _amount The amount to credit to the address
     * @return Returns the transaction status as a uint (e.g., 0 for failure, 1 for success)
     */
    function _creditTo(uint16 _srcChainId, address _toAddress, uint256 _amount)
        internal
        virtual
        override
        returns (uint256)
    {
        // TODO: Implement
        _srcChainId;
        _toAddress;
        _amount;
        return 0;
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
