// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

pragma solidity ^0.8.4;

import "../../../../lib/forge-std/src/Script.sol";
import "../../../../lib/forge-std/src/console2.sol";
import "../staking.sol";
import "./LiquidStaking.sol";

contract Deploy is Script {
    address precompile =
        address(0xd9A998CaC66092748FfEc7cFBD155Aae1737C2fF);
    IStakingModule staking = IStakingModule(precompile);

    function run() public {
        vm.startBroadcast();
        
        LiquidStaking ls = new LiquidStaking(
            "hello",
            "sss",
            precompile,
            address(0xE77B9d929c8599b811265145e397AcA50591b246)
        );


        // require(x != bytes(""), "Failed to get active validators");
        // // staking.getDelegation(0x20f33CE90A13a4b5E7697E3544c3083B8F8A51D4, 0x60445cEEc8f3239524c03ba79117c8c343e8D2E3);
        // address[] memory vals = staking.getActiveValidators();
        // (bool success1, bytes memory data1) = address(ls).call(
        //     abi.encodeWithSignature("getActiveValidators()")
        // );
        // console.logString("IN DEPLOY");
        // console.logBytes(data1);
        // require(success1, "Failed to get active validators");   

        (bool success12, bytes memory data12) = address(ls).call(
            abi.encodeWithSignature("getActiveValidatorsMock()")
        );
        console.logString("IN DEPLOY2");
        console.logBytes(data12);
        require(success12, "Failed to get active validators");   
        // require(success1, "Failed to get active validators");   
        // console.logBytes(data1);
        // address validator = abi.decodeWithSignature("getActiveValidators()", data1, (address[]))[0];
        // console.logAddress(validator);

        // (bool success2, bytes memory data2) = precompile.call(
        //     abi.encodeWithSignature("delegate(address,uint256)", address(0x4ca86164278f898bA3C07a7179c07c7A4d2619D7), 889000000000000)
        // );
        // console.logBool(success2);
        // console.logBytes(data2);
        // staking.delegate(address(0x4ca86164278f898bA3C07a7179c07c7A4d2619D7), 889000000000000);

        vm.stopBroadcast();
    }
}
