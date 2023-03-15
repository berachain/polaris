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
