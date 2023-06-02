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

package main

// tests defined per sim.
var (
	ethTests = []testSpec{{Name: "http/ChainIDSupport", Run: chainIDSupport}, {Name: "http/GasPriceSupport", Run: gasPriceSupport}, {Name: "http/BlockNumberSupport", Run: blockNumberSupport}, {Name: "http/GetBalanceSupport", Run: getBalanceSupport}, {Name: "http/EstimateGasSupport", Run: estimateGasSupport}, {Name: "http/GetTransactionByHash", Run: getTransactionByHash}, {Name: "ws/ChainIDSupport", Run: chainIDSupport}, {Name: "ws/GasPriceSupport", Run: gasPriceSupport}, {Name: "ws/BlockNumberSupport", Run: blockNumberSupport}, {Name: "ws/GetBalanceSupport", Run: getBalanceSupport}, {Name: "ws/EstimateGasSupport", Run: estimateGasSupport}, {Name: "ws/GetTransactionByHash", Run: getTransactionByHash}} //nolint: lll // auto-generated
)

var tests = ethTests
