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

package params

import (
	"math/big"

	"github.com/berachain/stargazer/eth/common"
)

const DefaultEIP155ChainID = 42069

var DefaultChainConfig = &ChainConfig{
	ChainID:                       big.NewInt(DefaultEIP155ChainID),
	HomesteadBlock:                big.NewInt(0),
	DAOForkBlock:                  big.NewInt(0),
	DAOForkSupport:                true,
	EIP150Block:                   big.NewInt(0),
	EIP150Hash:                    common.Hash{},
	EIP155Block:                   big.NewInt(0),
	EIP158Block:                   big.NewInt(0),
	ByzantiumBlock:                big.NewInt(0),
	ConstantinopleBlock:           big.NewInt(0),
	PetersburgBlock:               big.NewInt(0),
	IstanbulBlock:                 big.NewInt(0),
	BerlinBlock:                   big.NewInt(0),
	LondonBlock:                   big.NewInt(0),
	ArrowGlacierBlock:             big.NewInt(0),
	GrayGlacierBlock:              big.NewInt(0),
	MergeNetsplitBlock:            big.NewInt(0),
	TerminalTotalDifficulty:       big.NewInt(0),
	TerminalTotalDifficultyPassed: true,
	Ethash:                        nil,
	Clique:                        nil,
}
