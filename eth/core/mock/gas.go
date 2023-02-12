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
)

type GasPluginMock struct {
	gasUsed  uint64
	gasLimit uint64
}

func NewGasPluginMock(gasLimit uint64) *GasPluginMock {
	return &GasPluginMock{
		gasLimit: gasLimit,
	}
}

func (w *GasPluginMock) Reset(context.Context) {
	w.gasUsed = 0
}

func (w *GasPluginMock) ConsumeGas(amount uint64) error {
	if w.gasUsed+amount > w.gasLimit {
		return errors.New("gas limit exceeded")
	}
	w.gasUsed += amount
	return nil
}

func (w *GasPluginMock) CumulativeGasUsed() uint64 {
	return w.gasUsed
}

func (w *GasPluginMock) GasRemaining() uint64 {
	return w.gasLimit - w.gasUsed
}

func (w *GasPluginMock) GasUsed() uint64 {
	return w.gasUsed
}

func (w *GasPluginMock) RefundGas(amount uint64) {
	if w.gasUsed < amount {
		w.gasUsed = 0
	} else {
		w.gasUsed -= amount
	}
}

func (w *GasPluginMock) SetGasLimit(amount uint64) error {
	w.gasLimit = amount
	if w.gasLimit < w.gasUsed {
		return errors.New("gas limit is below currently consumed")
	}
	return nil
}
