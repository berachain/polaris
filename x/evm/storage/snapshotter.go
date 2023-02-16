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

package keeper

import (
	snapshot "cosmossdk.io/store/snapshots/types"
	storetypes "cosmossdk.io/store/types"

	"github.com/berachain/stargazer/eth/core/state"
	"github.com/berachain/stargazer/x/evm/types"
)

// Compile-time interface assertion.
var _ snapshot.ExtensionSnapshotter = &BytecodeSnapshotter{}

// SnapshotFormat format 1 is just gzipped wasm byte code for each item payload. No protobuf envelope, no metadata.
const SnapshotFormat = 1

type BytecodeSnapshotter struct {
	cms storetypes.MultiStore
	sp  state.Plugin
}

// `NewBytecodeSnapshotter` creates a new BytecodeSnapshotter.
func NewBytecodeSnapshotter(cms storetypes.MultiStore, sp state.Plugin) *BytecodeSnapshotter {
	return &BytecodeSnapshotter{
		cms: cms,
		sp:  sp,
	}
}

func (bys *BytecodeSnapshotter) SnapshotName() string {
	return types.ModuleName
}

func (bys *BytecodeSnapshotter) SnapshotFormat() uint32 {
	return SnapshotFormat
}

func (bys *BytecodeSnapshotter) SupportedFormats() []uint32 {
	return []uint32{SnapshotFormat}
}

func (bys *BytecodeSnapshotter) SnapshotExtension(
	height uint64, payloadWriter snapshot.ExtensionPayloadWriter,
) error {
	// cacheMS, err := bys.cms.CacheMultiStoreWithVersion(int64(height))
	// if err != nil {
	// 	return err
	// }

	// ctx := sdk.NewContext(cacheMS, tmproto.Header{}, false, log.NewNopLogger())
	// seenBefore := make(map[string]bool)
	var err error

	// Iterate over all bytecode

	// bys.sp.IterateCodeInfos(ctx, func(id uint64, info types.CodeInfo) bool {
	// 	// Many code ids may point to the same code hash... only sync it once
	// 	hexHash := hex.EncodeToString(info.CodeHash)
	// 	// if seenBefore, just skip this one and move to the next
	// 	if seenBefore[hexHash] {
	// 		return false
	// 	}
	// 	seenBefore[hexHash] = true

	// 	// load code and abort on error
	// 	wasmBytes, err := bys.wasm.GetByteCode(ctx, id)
	// 	if err != nil {
	// 		rerr = err
	// 		return true
	// 	}

	// 	compressedWasm, err := ioutils.GzipIt(wasmBytes)
	// 	if err != nil {
	// 		rerr = err
	// 		return true
	// 	}

	// 	err = snapshot.WriteExtensionItem(protoWriter, compressedWasm)
	// 	if err != nil {
	// 		rerr = err
	// 		return true
	// 	}

	// 	return false
	// })

	return err
}

func (bys *BytecodeSnapshotter) RestoreExtension(
	height uint64, format uint32, reader snapshot.ExtensionPayloadReader) error {
	// if format == SnapshotFormat {
	// 	// return bys.processAllItems(height, protoReader, restoreV1, finalizeV1)
	// }
	return snapshot.ErrUnknownFormat
}

// func restoreV1(ctx sdk.Context, k *Keeper, compressedCode []byte) error {
// 	if !ioutils.IsGzip(compressedCode) {
// 		return types.ErrInvalid.Wrap("not a gzip")
// 	}
// 	wasmCode, err := ioutils.Uncompress(compressedCode, uint64(types.MaxWasmSize))
// 	if err != nil {
// 		return sdkerrors.Wrap(types.ErrCreateFailed, err.Error())
// 	}

// 	// FIXME: check which codeIDs the checksum matches??
// 	_, err = k.wasmVM.Create(wasmCode)
// 	if err != nil {
// 		return sdkerrors.Wrap(types.ErrCreateFailed, err.Error())
// 	}
// 	return nil
// }

// func finalizeV1(ctx sdk.Context, k *Keeper) error {
// 	// FIXME: ensure all codes have been uploaded?
// 	return k.InitializePinnedCodes(ctx)
// }

// func (bys *BytecodeSnapshotter) processAllItems(
// 	height uint64,
// 	protoReader protoio.Reader,
// 	cb func(sdk.Context, *Keeper, []byte) error,
// 	finalize func(sdk.Context, *Keeper) error,
// ) (snapshot.SnapshotItem, error) {
// 	ctx := sdk.NewContext(bys.cms, tmproto.Header{Height: int64(height)}, false, log.NewNopLogger())

// 	// keep the last item here... if we break, it will either be empty (if we hit io.EOF)
// 	// or contain the last item (if we hit payload == nil)
// 	var item snapshot.SnapshotItem
// 	for {
// 		item = snapshot.SnapshotItem{}
// 		err := protoReader.ReadMsg(&item)
// 		if err == io.EOF {
// 			break
// 		} else if err != nil {
// 			return snapshot.SnapshotItem{}, sdkerrors.Wrap(err, "invalid protobuf message")
// 		}

// 		// if it is not another ExtensionPayload message, then it is not for us.
// 		// we should return it an let the manager handle this one
// 		payload := item.GetExtensionPayload()
// 		if payload == nil {
// 			break
// 		}

// 		if err := cb(ctx, bys.wasm, payload.Payload); err != nil {
// 			return snapshot.SnapshotItem{}, sdkerrors.Wrap(err, "processing snapshot item")
// 		}
// 	}

// 	return item, finalize(ctx, bys.wasm)
// }
