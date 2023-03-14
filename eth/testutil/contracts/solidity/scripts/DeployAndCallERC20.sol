pragma solidity ^0.8.17;

import "../lib/forge-std/src/Script.sol";
import "../src/SolmateERC20.sol";

contract DeployAndCallERC20 is Script {
    function run() public {
        address dropAddress = address(12);
        uint256 quantity = 50000;

        vm.startBroadcast();
        SolmateERC20 drop = new SolmateERC20();

        for (uint256 i = 0; i < 66; i++) {
            drop.mint(dropAddress, quantity);
        }

        vm.stopBroadcast();
    }
}
