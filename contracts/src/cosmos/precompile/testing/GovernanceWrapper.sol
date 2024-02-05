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

pragma solidity ^0.8.17;

import {IGovernanceModule} from "../Governance.sol";
import {IBankModule} from "../Bank.sol";
import {Cosmos} from "../../CosmosTypes.sol";

contract GovernanceWrapper {
    // State
    IGovernanceModule public governanceModule;
    IBankModule public immutable bank = IBankModule(0x4381dC2aB14285160c808659aEe005D51255adD7);

    // Errors
    error ZeroAddress();

    /**
     * @dev Constructor.
     * @param _governanceModule The address of the governance module.
     */
    constructor(address _governanceModule) {
        if (_governanceModule == address(0)) {
            revert ZeroAddress();
        }
        governanceModule = IGovernanceModule(_governanceModule);
    }

    /**
     * @dev Submit a proposal.
     * @param proposal The proposal.
     */
    function submit(IGovernanceModule.MsgSubmitProposal calldata proposal, string calldata denom, uint256 amount)
        external
        payable
        returns (uint64)
    {
        // Send the deposit amount to the contract.
        Cosmos.Coin[] memory coins = new Cosmos.Coin[](1);
        coins[0].denom = denom;
        coins[0].amount = amount;
        return governanceModule.submitProposal(proposal);
    }

    /**
     * @dev get a proposal.
     * @param proposalId The proposal id.
     */
    function getProposal(uint64 proposalId) external view returns (IGovernanceModule.Proposal memory) {
        return governanceModule.getProposal(proposalId);
    }

    /**
     * @dev get proposals.
     * @param proposalStatus The proposal status.
     */
    function getProposals(int32 proposalStatus) external view returns (IGovernanceModule.Proposal[] memory) {
        Cosmos.PageRequest memory pageReq;
        (IGovernanceModule.Proposal[] memory proposals,) = governanceModule.getProposals(proposalStatus, pageReq);
        return proposals;
    }

    /**
     * @dev vote.
     * @param proposalId The proposal id.
     * @param option The option.
     * @param metadata The metadata.
     */
    function vote(uint64 proposalId, int32 option, string memory metadata) external returns (bool) {
        return governanceModule.vote(proposalId, option, metadata);
    }

    /**
     * @dev Cancel a proposal. Returns the canceled time and height.
     *   burned.
     * @param proposalId The id of the proposal to cancel.
     */
    function cancelProposal(uint64 proposalId) external returns (uint64, uint64) {
        return governanceModule.cancelProposal(proposalId);
    }

    // Fallback function for receiving funds.
    receive() external payable {}
}
