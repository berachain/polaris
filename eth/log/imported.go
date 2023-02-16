// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package log

import "github.com/ethereum/go-ethereum/log"

type (
	// `Record` is a log record.
	Record = log.Record
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
