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
