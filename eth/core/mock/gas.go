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

package mock

import (
	"context"
	"errors"

	"github.com/berachain/stargazer/eth/core"
)

type GasPluginMock struct {
	txGasUsed     uint64
	blockGasUsed  uint64
	txGasLimit    uint64
	blockGasLimit uint64
}

func NewGasPluginMock() *GasPluginMock {
	return &GasPluginMock{}
}

func (w *GasPluginMock) Prepare(context.Context) {
	w.blockGasUsed = 0
}

func (w *GasPluginMock) Reset(context.Context) {
	w.txGasUsed = 0
}

func (w *GasPluginMock) TxConsumeGas(amount uint64) error {
	if w.txGasUsed+amount > w.txGasLimit {
		return errors.New("gas limit exceeded")
	}
	if w.blockGasUsed+amount > w.blockGasLimit {
		return core.ErrBlockOutOfGas
	}

	w.txGasUsed += amount
	return nil
}

func (w *GasPluginMock) CumulativeGasUsed() uint64 {
	return w.txGasUsed + w.blockGasUsed
}

func (w *GasPluginMock) TxGasRemaining() uint64 {
	return w.txGasLimit - w.txGasUsed
}

func (w *GasPluginMock) TxGasUsed() uint64 {
	return w.txGasUsed
}

func (w *GasPluginMock) TxRefundGas(amount uint64) {
	if w.txGasUsed < amount {
		w.txGasUsed = 0
	} else {
		w.txGasUsed -= amount
	}
}

func (w *GasPluginMock) SetTxGasLimit(amount uint64) error {
	w.txGasLimit = amount
	if w.txGasLimit < w.txGasUsed {
		return errors.New("gas limit is below currently consumed")
	}
	return nil
}

func (w *GasPluginMock) SetBlockGasLimit(amount uint64) {
	w.blockGasLimit = amount
}

func (w *GasPluginMock) BlockGasLimit() uint64 {
	return w.blockGasLimit
}

func (w *GasPluginMock) ConsumeGasToBlockLimit() error {
	delta := w.blockGasLimit - w.blockGasUsed
	if w.txGasUsed+delta > w.txGasLimit {
		return errors.New("tx gas limit exceeded")
	}

	w.txGasUsed += delta
	w.blockGasUsed += delta
	return nil
}
