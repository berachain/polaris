// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

pragma solidity ^0.8.4;

interface IGovernanceModule {
    ////////////////////////////////////////// Write Methods /////////////////////////////////////////////
    function submitProposal(
        bytes calldata message,
        Coin[] calldata initialDeposit,
        string calldata metadata,
        string calldata title,
        string calldata summary,
        bool expedited
    ) external returns (uint64);

    function cancelProposal(
        uint256 proposalId
    ) external returns (uint256, uint256, uint256);

    function execLegacyContent(
        bytes calldata content,
        string calldata authority
    ) external;

    function vote(
        uint256 proposalId,
        int32 option,
        string memory metadata
    ) external;

    function voteWeighted(
        uint256 proposalId,
        WeightedVoteOption[] calldata options,
        string calldata metadata
    ) external;

    //////////////////////////////////////////// UTILS ////////////////////////////////////////////
    /**
     * @dev Represents a cosmos coin.
     * Note: this struct is generated in precompile/generated/i_staking_module.abigen.go
     */
    struct Coin {
        uint256 amount;
        string denom;
    }

    /**
     * @dev Represents a governance module `WeightedVoteOption`.
     * Note: this struct is generated in precompile/generated/i_staking_module.abigen.go
     */
    struct WeightedVoteOption {
        int32 voteOption;
        string weight;
    }
}
