// SPDX-License-Identifier: MIT
pragma solidity >=0.8.4;

contract RevertableTx {
    receive() external payable {
        revert("RevertableTx: revertTx");
    }
}