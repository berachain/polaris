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

/**
 * @dev Interface of all supported Cosmos events emitted by the bank module
 */
interface IBankModule {
    ////////////////////////////////////////// EVENTS /////////////////////////////////////////////
    
    /**
     * @dev Emitted by the bank module when `amount` tokens are sent to `recipient`
     */
    event Transfer(address indexed recipient, uint256 amount);

    /**
     * @dev Emitted by the bank module when `sender` sends some amount of tokens
     */
    event Message(address indexed sender);

    /**
     * @dev Emitted by the bank module when `amount` tokens are spent by `spender`
     */
    event CoinSpent(address indexed spender, uint256 amount);

    /**
     * @dev Emitted by the bank module when `amount` tokens are received by `receiver`
     */
    event CoinReceived(address indexed receiver, uint256 amount);

    /**
     * @dev Emitted by the bank module when `amount` tokens are minted by `minter`
     *
     * Note: "Coinbase" refers to the Cosmos event: EventTypeCoinMint. `minter` is a module
     * address.
     */
    event Coinbase(address indexed minter, uint256 amount);

    /**
     * @dev Emitted by the bank module when `amount` tokens are burned by `burner`
     *
     * Note: `burner` is a module address
     */
    event Burn(address indexed burner, uint256 amount);
    
    /////////////////////////////////////// READ METHODS //////////////////////////////////////////

    // func (k BaseKeeper) Balance(ctx context.Context, req *types.QueryBalanceRequest) (*types.QueryBalanceResponse, error) {
        // The Balance endpoint allows users to query account balance by address for a given denomination.
    function getBalance(address accountAddress, string denom) external view returns (uint256);

// func (k BaseKeeper) AllBalances(ctx context.Context, req *types.QueryAllBalancesRequest) (*types.QueryAllBalancesResponse, error) {
    // The AllBalances endpoint allows users to query account balance by address for all denominations.
    function getAllBalance(address accountAddress) external view returns (uint256);

// func (k BaseKeeper) TotalSupply(ctx context.Context, req *types.QueryTotalSupplyRequest) (*types.QueryTotalSupplyResponse, error) {
//     function getTotalSupply() external view returns (uint256);

// func (k BaseKeeper) SupplyOf(c context.Context, req *types.QuerySupplyOfRequest) (*types.QuerySupplyOfResponse, error) {
//     function geSupplyOf() external view returns (uint256);

    ////////////////////////////////////// WRITE METHODS //////////////////////////////////////////

    /**
     * @dev msg.sender delegates the `amount` of tokens to `validatorAddress`
     */
    function delegate(address validatorAddress, uint256 amount) external payable returns (bool);


    //////////////////////////////////////////// UTILS ////////////////////////////////////////////


}
