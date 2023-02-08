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

import "errors"

var (
	// `ErrEthEventNotRegistered` is returned when an incoming event is not mapped to any
	// registered Ethereum event.
	ErrEthEventNotRegistered = errors.New("no Ethereum event was registered for this Cosmos event")
	// `ErrNoAttributeKeyFound` is returned when no Cosmos event attribute is provided for a
	// certain Ethereum event's argument.
	ErrNoAttributeKeyFound = errors.New("this Ethereum event argument has no matching Cosmos attribute key")
	// `ErrNotEnoughAttributes` is returned when a Cosmos event does not have enough attributes for
	// its corresponding Ethereum event; there are less Cosmos event attributes than Ethereum event
	// arguments.
	ErrNotEnoughAttributes = errors.New("not enough event attributes provided")
	// `ErrNoValueDecoderFunc` is returned when a Cosmos event's attribute key is not mapped to any
	// attribute value decoder function.
	ErrNoValueDecoderFunc = errors.New("no value decoder function is found for event attribute key")
)
