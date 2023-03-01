pragma solidity ^0.8.4;

import "forge-std/Test.sol";
import "../src/examples/LiquidStaking.sol";
import {ERC20} from "solmate/tokens/ERC20.sol";
import {MockStaking} from "./MockStaking.sol";

contract Helper is Test {
    address ADMIN = address(0x1);
}

contract LiquidStakingTest is Helper {
    LiquidStaking public liquidStaking;
    IStakingModule public staking;

    function setUp() public {
        staking = new MockStaking();
        liquidStaking = new LiquidStaking(
            "Test Token",
            "TEST",
            address(staking),
            address(0x123)
        );
    }

    function testTotalAssets() public {
        uint256 totalAssets = liquidStaking.totalAssets();
        assertEq(totalAssets, 10 ether);
    }

    function testDelegate() public {
        vm.startPrank(ADMIN);
        uint256 depositAmount = 1 ether;
        vm.deal(ADMIN, depositAmount);
        liquidStaking.delegate{value: depositAmount}();
        assertEq(liquidStaking.balanceOf(ADMIN), depositAmount);
        vm.stopPrank();
    }

    function withdraw() public {
        vm.startPrank(ADMIN);
        uint256 depositAmount = 1 ether;
        vm.deal(ADMIN, depositAmount);
        liquidStaking.delegate{value: depositAmount}();
        liquidStaking.withdraw(depositAmount);
        assertEq(liquidStaking.balanceOf(ADMIN), 0);
        vm.stopPrank();
    }
}
