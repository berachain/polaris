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

interface IGovernanceModule {
    ////////////////////////////////////////// Write Methods /////////////////////////////////////////////
    /**
     * @dev Submit a proposal to the governance module. Returns the proposal id.
     */
    function submitProposal(bytes calldata proposal, bytes calldata message) external returns (uint64);

    /**
     * @dev Cancel a proposal. Returns the cancled time and height.
     *   burned.
     */
    function cancelProposal(uint64 proposalId) external returns (uint64, uint64);

    /**
     * @dev Vote on a proposal.
     */
    function vote(uint64 proposalId, int32 option, string memory metadata) external returns (bool);

    /**
     * @dev Vote on a proposal with weights.
     */
    function voteWeighted(uint64 proposalId, WeightedVoteOption[] calldata options, string calldata metadata)
        external
        returns (bool);

    ////////////////////////////////////////// Read Methods /////////////////////////////////////////////

    /**
     * @dev Get the proposal with the given id.
     */
    function getProposal(uint64 proposalId) external view returns (Proposal memory);

    /**
     * @dev Get proposals with a given status.
     */
    function getProposals(int32 proposalStatus) external view returns (Proposal[] memory);

    ////////////////////////////////////////// Structs ///////////////////////////////////////////////////

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

    /**
     * @dev Emitted by the governance module when `submitProposal` is called.
     */
    event SubmitProposal(uint64 indexed proposalId, string indexed proposalMessages);

    /**
     * @dev Emitted by the governance module when `submitProposal` is called.
     */
    event SubmitProposal(uint64 indexed votingPeriodStart);

    /**
     * @dev Emitted by the governance module when `submitProposal` is called.
     */
    event ProposalDeposit(string indexed amount, uint64 indexed proposalId);

    /**
     * @dev Emitted by the governance module when `AddVote` is called in the msg server.
     */
    event ProposalVote(string indexed option, uint64 indexed proposalId);

    /**
     * @dev Emitted by the governance module when `cancelProposal` is called.
     */
    event CancelProposal(string indexed sender, uint64 indexed proposalId);
}
