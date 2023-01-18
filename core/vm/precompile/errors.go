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

package precompile

import "errors"

var (
	// `ErrNoPrecompileMethodForABIMethod` is returned when no precompile method is provided for a
	// corresponding ABI method.
	ErrNoPrecompileMethodForABIMethod = errors.New("this ABI method does not have a corresponding precompile method")

	// `ErrWrongContainerFactory` is returned when the wrong precompile container factory is used
	// to build a precompile contract.
	ErrWrongContainerFactory = errors.New("this container factory does not support the given precompile contract")

	// `ErrNotStatelessPrecompile` is returned when a base contract implementation does not
	// properly implement the `StatelessContractImpl` interface.
	ErrNotStatelessPrecompile = errors.New("this precompile contract does not implement `StatelessContractImpl`")
)
