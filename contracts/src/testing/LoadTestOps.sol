pragma solidity 0.8.22;

/**
 * @title LoadTestOps
 * @dev This contract is used to benchmark I/O and computation complex operations
 */
contract LoadTestOps {
    uint256 private data = 1;

    /**
     * @dev Loads data into memory and performs various operations
     */
    function loadData() public {
        assembly {
            // mload
            let m := mload(0x40)
            // mstore
            mstore(m, 0x60)
            // sstore and keccak256
            sstore(data.slot, keccak256(m, 0x20))
        }
    }
}
