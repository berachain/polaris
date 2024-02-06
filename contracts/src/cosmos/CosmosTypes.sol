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

/**
 * @dev This library contains types used by the Cosmos module.
 */
library Cosmos {
    /**
     * @dev Represents a cosmos coin.
     */
    struct Coin {
        uint256 amount;
        string denom;
    }

    struct PageRequest {
        string key;
        uint64 offset;
        uint64 limit;
        bool countTotal;
        bool reverse;
    }

    struct PageResponse {
        string nextKey;
        uint64 total;
    }

    /**
     * @dev Represents a Cosmos SDK `codectypes.Any`.
     */
    struct CodecAny {
        string typeURL;
        bytes value;
    }
}

/**
 * @dev This contract uses types in the Cosmos library.
 */
contract CosmosTypes {
    function coin(Cosmos.Coin calldata) public pure {}

    function pageRequest(Cosmos.PageRequest calldata) public pure {}

    function pageResponse(Cosmos.PageResponse calldata) public pure {}

    function codecAny(Cosmos.CodecAny calldata) public pure {}
}
