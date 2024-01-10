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

package runtime

import (
	"context"

	"golang.org/x/exp/slog"

	"cosmossdk.io/log"
)

// ethHandler implements the slog.Handler interface.
// It is used to handle logging for the Ethereum module.
var _ slog.Handler = (*ethHandler)(nil)

// ethHandler is a struct that contains a logger.
type ethHandler struct {
	logger log.Logger
}

// newEthHandler is a constructor function for ethHandler.
// It takes a logger as an argument and returns a slog.Handler.
func newEthHandler(_logger log.Logger) slog.Handler {
	return &ethHandler{logger: _logger}
}

// With is a method on ethHandler that returns a new ethHandler with
// additional context.
func (h *ethHandler) With(ctx ...interface{}) slog.Handler {
	return &ethHandler{logger: h.logger.With(ctx...)}
}

// Handle is a method on ethHandler that logs a message at the
// appropriate level with context key/value pairs.
func (h *ethHandler) Handle(_ context.Context, r slog.Record) error {
	polarisGethHandler := h.logger.With("module", "polaris-geth")
	x := r.NumAttrs()
	attrs := make([]interface{}, 0, x*2) //nolint:gomnd // 2 times.
	r.Attrs(func(a slog.Attr) bool {
		attrs = append(attrs, a.Key)
		attrs = append(attrs, a.Value)
		x--
		return x != -1
	})
	switch r.Level { //nolint:nolintlint,exhaustive // linter is bugged.
	case slog.LevelDebug:
		polarisGethHandler.Debug(r.Message, attrs...)
	case slog.LevelInfo:
		polarisGethHandler.Info(r.Message, attrs...)
	case slog.LevelWarn:
		polarisGethHandler.Error(r.Message, attrs...)
	case slog.LevelError:
		polarisGethHandler.Error(r.Message, attrs...)
	}
	return nil
}

// WithAttrs is a method on ethHandler that returns a new ethHandler with
// additional context provided by a set of slog.Attr.
func (h *ethHandler) WithAttrs(as []slog.Attr) slog.Handler {
	newLogger := h.logger
	for _, a := range as {
		newLogger = newLogger.With(a.Key, a.Value)
	}
	return &ethHandler{logger: newLogger}
}

// WithGroup is a method on ethHandler that returns a new ethHandler with
// additional context provided by a group name.
func (h *ethHandler) WithGroup(name string) slog.Handler {
	return h.WithAttrs([]slog.Attr{{Key: "group", Value: slog.StringValue(name)}})
}

// Enabled reports whether l emits log records at the given context and level.
func (*ethHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return true
}
