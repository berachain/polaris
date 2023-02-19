// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;

contract RevertableTx {
    fallback() external {
        revert("RevertableTx: revertTx");
    }
}