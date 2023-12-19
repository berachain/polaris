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

pragma solidity 0.8.23;

import {Cosmos} from "../CosmosTypes.sol";

/**
 * @dev Interface of the governance module's precompiled contract
 */
interface IGovernanceModule {
    ////////////////////////////////////////// Write Methods /////////////////////////////////////////////
    /**
     * @dev Submit a proposal to the governance module.
     * @notice Use the codec to marshal the proposal messages.
     * @param proposal The proposal to submit.
     * @return The id of the proposal.
     */
    function submitProposal(MsgSubmitProposal calldata proposal) external returns (uint64);

    /**
     * @dev Cancel a proposal.
     * @param proposalId The id of the proposal to cancel.
     * @return The time and block height the proposal was canceled.
     */
    function cancelProposal(uint64 proposalId) external returns (uint64, uint64);

    /**
     * @dev Vote on a proposal.
     * @param proposalId The id of the proposal to vote on.
     * @param option The option to vote on.
     * @param metadata The metadata to attach to the vote.
     */
    function vote(uint64 proposalId, int32 option, string memory metadata) external returns (bool);

    /**
     * @dev Vote on a proposal with weights.
     * @param proposalId The id of the proposal to vote on.
     * @param options The options to vote on.
     * @param metadata The metadata to attach to the vote.
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
     * @dev Get the deposits of the proposal with the given id.
     */
    function getProposalDeposits(uint64 proposalId) external view returns (Cosmos.Coin[] memory);

    /**
     * @dev Get the deposits of the proposal with the given id and depositor.
     */
    function getProposalDepositsByDepositor(uint64 proposalId, address depositor)
        external
        view
        returns (Cosmos.Coin[] memory);

    /**
     * @dev Get proposals with a given status.
     * @notice Accepts pagination request (empty == no pagination returned).
     * @param proposalStatus The status of the proposals to get.
     */
    function getProposals(int32 proposalStatus, Cosmos.PageRequest calldata pagination)
        external
        view
        returns (Proposal[] memory, Cosmos.PageResponse memory);

    /**
     * @dev Get the proposal tally result for the given id.
     * @param proposalId The id of the proposal to get the tally result for.
     */
    function getProposalTallyResult(uint64 proposalId) external view returns (TallyResult memory);

    /**
     * @dev Get the proposal votes with the given id.
     * @notice Accepts pagination request.
     * @param proposalId The id of the proposal to get the votes for.
     */
    function getProposalVotes(uint64 proposalId, Cosmos.PageRequest calldata pagination)
        external
        view
        returns (Vote[] memory, Cosmos.PageResponse memory);

    /**
     * @dev Get the proposal vote information with the given id and voter.
     * @param proposalId The id of the proposal to get the vote info for.
     * @param voter The address of the voter to get the vote info for.
     */
    function getProposalVotesByVoter(uint64 proposalId, address voter) external view returns (Vote memory);

    /**
     * @dev Get the governance module parameters.
     */
    function getParams() external view returns (Params memory);

    /**
     * @dev Get the governance module voting parameters.
     */
    function getVotingParams() external view returns (VotingParams memory);

    /**
     * @dev Get the governance module deposit parameters.
     */
    function getDepositParams() external view returns (DepositParams memory);

    /**
     * @dev Get the governance module tally parameters.
     */
    function getTallyParams() external view returns (TallyParams memory);

    /**
     * @dev Get the constitution of the chain.
     */
    function getConstitution() external view returns (string memory);

    ////////////////////////////////////////// Structs ///////////////////////////////////////////////////
    /**
     * @dev Represents a governance module `WeightedVoteOption`.
     */
    struct WeightedVoteOption {
        int32 voteOption;
        string weight;
    }

    /**
     * @dev Represents a governance module `Vote`.
     */
    struct Vote {
        uint64 proposalId;
        address voter;
        WeightedVoteOption[] options;
        string metadata;
    }

    /**
     * @dev Represents a governance module `MsgSubmitProposal`.
     */
    struct MsgSubmitProposal {
        Cosmos.CodecAny[] messages;
        Cosmos.Coin[] initialDeposit;
        address proposer;
        string metadata;
        string title;
        string summary;
        bool expedited;
    }

    /**
     * @dev Represents a governance module `Proposal`.
     */
    struct Proposal {
        uint64 id;
        Cosmos.CodecAny[] messages;
        int32 status;
        TallyResult finalTallyResult;
        uint64 submitTime;
        uint64 depositEndTime;
        Cosmos.Coin[] totalDeposit;
        uint64 votingStartTime;
        uint64 votingEndTime;
        string metadata;
        string title;
        string summary;
        address proposer;
    }

    /**
     * @dev Represents the governance module's parameters.
     */
    struct Params {
        Cosmos.Coin[] minDeposit;
        uint64 maxDepositPeriod;
        uint64 votingPeriod;
        string quorum;
        string threshold;
        string vetoThreshold;
        string minInitialDepositRatio;
        string proposalCancelRatio;
        string proposalCancelDest;
        uint64 expeditedVotingPeriod;
        string expeditedThreshold;
        Cosmos.Coin[] expeditedMinDeposit;
        bool burnVoteQuorum;
        bool burnProposalDepositPrevote;
        bool burnVoteVeto;
    }

    /**
     * @dev Represents the governance module's `VotingParams`.
     */
    struct VotingParams {
        uint64 votingPeriod;
    }

    /**
     * @dev Represents the governance module's `DepositParams`.
     */
    struct DepositParams {
        Cosmos.Coin[] minDeposit;
        uint64 maxDepositPeriod;
    }

    /**
     * @dev Represents the governance module's `TallyParams`.
     */
    struct TallyParams {
        string quorum;
        string threshold;
        string vetoThreshold;
    }

    /**
     * @dev Represents a governance module `TallyResult`.
     */
    struct TallyResult {
        string yesCount;
        string abstainCount;
        string noCount;
        string noWithVetoCount;
    }

    /**
     * @dev Emitted by the governance precompile when `submitProposal` is called.
     * @param proposalId The id of the proposal.
     * @param proposalSender The sender of the submit proposal.
     */
    event ProposalSubmitted(uint64 indexed proposalId, address indexed proposalSender);

    /**
     * @dev Emitted by the governance module when `submitProposal` is called.
     * @param proposalId The id of the proposal.
     * @param amount The amount of the deposit.
     */
    event ProposalDeposit(uint64 indexed proposalId, Cosmos.Coin[] amount);

    /**
     * @dev Emitted by the governance precompile when a proposal is voted on.
     * @param proposalVote The vote that was voted on for a proposal.
     */
    event ProposalVoted(Vote proposalVote);

    /**
     * @dev Emitted by the governance module when `cancelProposal` is called.
     * @param proposalId The id of the proposal.
     * @param sender The sender of the cancel proposal.
     */
    event CancelProposal(uint64 indexed proposalId, address indexed sender);
}
