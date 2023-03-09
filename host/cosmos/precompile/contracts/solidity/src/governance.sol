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
     *@dev Vote on a proposal.
     */
    function vote(
        uint64 proposalId,
        int32 option,
        string memory metadata
    ) external;

    /**
     *@dev Vote on a proposal with weights.
     */
    function voteWeighted(
        uint64 proposalId,
        WeightedVoteOption[] calldata options,
        string calldata metadata
    ) external;

    ////////////////////////////////////////// Read Methods /////////////////////////////////////////////

    /**
     *@dev Get the proposal with the given id.
     */
    function getProposal(
        uint256 proposalId
    ) external view returns (Proposal memory);

    /**
     *@dev Get proposals with a given status.
     */
    function getProposalsStringAddr(
        int32 proposalStatus
    ) external view returns (Proposal[] memory);

    /**
     *@dev Get proposals with a given status, voter, and depositor, using eth addresses.
     */
    function getProposalsAddr(
        int32 proposalStatus
    ) external view returns (Proposal[] memory);

    ////////////////////////////////////////// Utils  ///////////////////////////////////////////////////

    /**
     * @dev Represents a cosmos coin.
     * Note: this struct is generated as go struct that is then used in the precompile.
     */
    struct Coin {
        uint64 amount;
        string denom;
    }

    /**
     * @dev Represents a governance module `WeightedVoteOption`.
     * Note: this struct is generated in generated/i_staking_module.abigen.go
     */
    struct WeightedVoteOption {
        int32 voteOption;
        string weight;
    }

    /**
     * @dev Represents a governance module `Proposal`.
     * Note: this struct is generated in generated/i_staking_module.abigen.go
     */
    struct Proposal {
        uint64 id;
        bytes message;
        int32 status;
        TallyResult finalTallyResult;
        uint64 submitTime;
        uint64 depositEndTime;
        Coin[] totalDeposit;
        uint64 votingStartTime;
        uint64 votingEndTime;
        string metadata;
        string title;
        string summary;
        string proposer;
    }

    /**
     * @dev Represents a governance module `TallyResult`.
     * Note: this struct is generated in generated/i_staking_module.abigen.go
     */
    struct TallyResult {
        string yesCount;
        string abstainCount;
        string noCount;
        string noWithVetoCount;
    }
}
