// SPDX-License-Identifier: Apache-2.0
//

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
