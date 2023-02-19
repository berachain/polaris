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

package log

import "github.com/ethereum/go-ethereum/log"

type (
	// `Record` is a log record.
	Record = log.Record

	// `Logger` defines the logger interface.
	Logger = log.Logger
)

var (
	// `Root` is the root logger.
	Root = log.Root

	// `LvlTrace` is the trace log level.
	LvlTrace = log.LvlTrace

	// `LvlDebug` is the debug log level.
	LvlDebug = log.LvlDebug

	// `LvlInfo` is the info log level.
	LvlInfo = log.LvlInfo

	// `LvlWarn` is the warn log level.
	LvlWarn = log.LvlWarn

	// `LvlError` is the error log level.
	LvlError = log.LvlError

	// `LvlCrit` is the critical log level.
	LvlCrit = log.LvlCrit

	// `FuncHandler` is a log handler.
	FuncHandler = log.FuncHandler
)
