// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0;

import {IERC20} from "../../lib/IERC20.sol";
import {IBankModule} from "./precompile/Bank.sol";
import {IERC20Module} from "./precompile/ERC20Module.sol";

/// @notice Modern and gas efficient ERC20 + EIP-2612 implementation.
/// @author Solmate (https://github.com/Rari-Capital/solmate/blob/main/src/tokens/ERC20.sol)
/// @author Modified from Uniswap (https://github.com/Uniswap/uniswap-v2-core/blob/master/contracts/UniswapV2ERC20.sol)
/// @dev Do not manually set balances without updating totalSupply, as the sum of all user balances must not exceed it.
abstract contract ERC20 is IERC20 {
    /*//////////////////////////////////////////////////////////////
                              Precompiles
    //////////////////////////////////////////////////////////////*/

    function bank() public view returns (IBankModule) {
        return IBankModule(address(0x1));
    }


    function erc20Module() public view returns (IERC20Module) {
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
    function name() public view returns (string memeory) {
        return bank().getDenomMetadata(denom).display;
    } 



    /**
     * @dev symbol is a public view method for reading the `sdk.Coin` symbol for this erc20.
     * @return string the sdk.Coin symbol for this erc20.
     */
    function symbol () public view returns (string memory) {
        return bank().getDenomMetadata(denom).symbol;
    }

    /**
     * @dev decimals is a public view method for reading the `sdk.Coin` decimals for this erc20.
     * @return uint8 the sdk.Coin decimals for this erc20.
     */
    function decimals() public view returns (uint8) {
        return bank().getDenomMetadata(name()).denomUnits[0].exponent;
    }

    string immutable public denom;

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
        bank().getSpendableBalance(user, name());
    }

    // mapping(address => mapping(address => uint256)) public allowance;
    // 

    /*//////////////////////////////////////////////////////////////
                            EIP-2612 STORAGE
    //////////////////////////////////////////////////////////////*/

    uint256 internal immutable INITIAL_CHAIN_ID;

    bytes32 internal immutable INITIAL_DOMAIN_SEPARATOR;

    // mapping(address => uint256) public nonces; //TODO Permit.

    /*//////////////////////////////////////////////////////////////
                               CONSTRUCTOR
    //////////////////////////////////////////////////////////////*/

    constructor(
        string memory _denom
    ) {
        denom = _denom;        

        INITIAL_CHAIN_ID = block.chainid;
        // INITIAL_DOMAIN_SEPARATOR = computeDomainSeparator();
    }

    /*//////////////////////////////////////////////////////////////
                               ERC20 LOGIC
    //////////////////////////////////////////////////////////////*/

    // function approve(address spender, uint256 amount) public virtual returns (bool) {
    //     // TODO:
    //     // allowance[msg.sender][spender] = amount;
    //     // emit Approval(msg.sender, spender, amount);
    //     // return true;
    // }

    function transfer(address to, uint256 amount) public virtual returns (bool) {
        // Create an array of length one that holds struct Coin.
        Coin[] memory coins = new Coin[](1);
        coins[0] = sdk.Coin({
            denom: denom,
            amount: amount
        });
        bank().send(msg.sender, to, coins);
        emit Transfer(msg.sender, to, amount);
        return true;
    }

    function transferFrom(
        address from,
        address to,
        uint256 amount
    ) public virtual returns (bool) {
        // TODO:
        // uint256 allowed = allowance[from][msg.sender]; // Saves gas for limited approvals.\
        // if (allowed != type(uint256).max) allowance[from][msg.sender] = allowed - amount;
        // balanceOf[from] -= amount;
        // // Cannot overflow because the sum of all user
        // // balances can't exceed the max uint256 value.
        // unchecked {
        //     balanceOf[to] += amount;
        // }
        // emit Transfer(from, to, amount);
        // return true;

        // IBankModule(0x696969).ge
    }

    /*//////////////////////////////////////////////////////////////
                             EIP-2612 LOGIC
    //////////////////////////////////////////////////////////////*/

    function permit(
        address owner,
        address spender,
        uint256 value,
        uint256 deadline,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) public virtual {
        // TODO: 
        // require(deadline >= block.timestamp, "PERMIT_DEADLINE_EXPIRED");

        // // Unchecked because the only math done is incrementing
        // // the owner's nonce which cannot realistically overflow.
        // unchecked {
        //     address recoveredAddress = ecrecover(
        //         keccak256(
        //             abi.encodePacked(
        //                 "\x19\x01",
        //                 DOMAIN_SEPARATOR(),
        //                 keccak256(
        //                     abi.encode(
        //                         keccak256(
        //                             "Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)"
        //                         ),
        //                         owner,
        //                         spender,
        //                         value,
        //                         nonces[owner]++,
        //                         deadline
        //                     )
        //                 )
        //             )
        //         ),
        //         v,
        //         r,
        //         s
        //     );

        //     require(recoveredAddress != address(0) && recoveredAddress == owner, "INVALID_SIGNER");

        //     allowance[recoveredAddress][spender] = value;
        // }

        // emit Approval(owner, spender, value);
    }

    // function DOMAIN_SEPARATOR() public view virtual returns (bytes32) {
    //     return block.chainid == INITIAL_CHAIN_ID ? INITIAL_DOMAIN_SEPARATOR : computeDomainSeparator();
    // }

    // function computeDomainSeparator() internal view virtual returns (bytes32) {
    //     return
    //         keccak256(
    //             abi.encode(
    //                 keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"),
    //                 keccak256(bytes(name)),
    //                 keccak256("1"),
    //                 block.chainid,
    //                 address(this)
    //             )
    //         );
    // }



        struct Coin {
        uint256 amount;
        string denom;
    }

}