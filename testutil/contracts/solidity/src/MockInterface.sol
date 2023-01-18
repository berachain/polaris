// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

interface MockInterface {
    struct Object {
        uint256 creationHeight;
        string timeStamp;
    }

    function getOutput(
        string calldata str
    ) external returns (Object[] calldata);

    function getOutputPartial() external returns (Object calldata);

    function contractFunc(address addr) external returns (int256 ans);

    function contractFunc(string calldata str) external;
}
