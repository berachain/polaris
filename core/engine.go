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

package core

// // `EngineAPI` is the StateProcessor.
// type EngineAPI interface {
// 	// Set the stateDB factory
// 	SetStateDBFactory(StateDBFactory)

// 	Prepare(context.Context, common.Address, common.Hash)
// 	Finalize(context.Context)
// }

// type interface {
// 	Coinbase() common.Address
// 	Number() *big.Int
// }

// type StateDBFactory interface {
// 	BuildStateDB(ctx context.Context) vm.StargazerStateDB
// }

// type Engine struct {
// 	// currentBlock is the current block being processed.
// 	currentBlock EthereumBlock
// }

// func (e *Engine) EthHeader(ctx sdk.Context) *coretypes.Header {
// 	return &coretypes.Header{

// }

// func NewEngine() *Engine {
// 	return &Engine{}
// }

// func (e *Engine) PrepareForBlock(ctx context.Context, eb EthereumBlock) {

// }

// func (e *Engine) FinalizeBlock(ctx context.Context) (types.Receipts, []*types.Log, uint64, error) {
// 	return nil, nil, 0, nil
// }
