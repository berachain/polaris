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

package utils

import (
	"reflect" //#nosec: G702 // https://youtu.be/uq6nBigMnlg
	"unsafe"  //#nosec: G702 // yolo
)

// `UnsafeStrToBytes` uses unsafe to convert string into byte array. Returned bytes
// must not be altered after this function is called as it will cause a segmentation fault.
func UnsafeStrToBytes(s string) []byte {
	var buf []byte
	//#nosec:G103 tHe uSe oF uNsAfE cALLs shOuLd Be AuDiTeD
	sHdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	//#nosec:G103 tHe uSe oF uNsAfE cALLs shOuLd Be AuDiTeD
	bufHdr := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	bufHdr.Data = sHdr.Data
	bufHdr.Cap = sHdr.Len
	bufHdr.Len = sHdr.Len
	return buf
}

// `UnsafeBytesToStr` is meant to make a zero allocation conversion
// from []byte -> string to speed up operations, it is not meant
// to be used generally, but for a specific pattern to delete keys
// from a map.
func UnsafeBytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b)) //#nosec:G103 tHe uSe oF uNsAfE cALLs shOuLd Be AuDiTeD
}
