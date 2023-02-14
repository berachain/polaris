// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;

interface MockPrecompileInterface {
    struct Object {
        uint256 creationHeight;
        string timeStamp;
    }

    function getOutput(
        string calldata str
    ) external returns (Object[] calldata);

    function getOutputPartial() external returns (Object calldata);

    function contractFunc(address addr) external returns (uint256 ans);

    function contractFuncStr(string calldata str) external;
}
