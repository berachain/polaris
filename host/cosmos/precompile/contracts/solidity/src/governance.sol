pragma solidity ^0.8.4;

interface IGovernanceModule {
    ////////////////////////////////////////// Write Methods /////////////////////////////////////////////
    /**
     *@dev Submit a proposal to the governance module. Returns the proposal id.
     */
    function submitProposal(
        bytes calldata message,
        Coin[] calldata initialDeposit,
        string calldata metadata,
        string calldata title,
        string calldata summary,
        bool expedited
    ) external returns (uint64);

    /**
     *@dev Cancel a proposal. Returns the cancled time and height.
      burned.
     */
    function cancelProposal(
        uint64 proposalId
    ) external returns (uint64, uint64);

    /**
     * @dev Represents a cosmos coin.
     * Note: this struct is generated as go struct that is then used in the precompile.
     */
    struct Coin {
        uint64 amount;
        string denom;
    }
}
