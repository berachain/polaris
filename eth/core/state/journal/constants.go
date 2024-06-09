// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package journal

const (
	// initCapacity is the initial capacity of the journals.
	// TODO: determine appropriate value.
	initCapacity = 32
	// refundRegistryKey is the registry key for the refund journal.
	refundRegistryKey = `refund`
	// logsRegistryKey is the registry key for the logs journal.
	logsRegistryKey = `logs`
	// accessListRegistryKey is the registry key for the access list journal.
	accessListRegistryKey = `accessList`
	// `suicidesRegistryKey` is the registry key for the suicides journal.
	suicidesRegistryKey = `suicides`
	// `transientRegistryKey` is the registry key for the transient journal.
	transientRegistryKey = `transient`
)
