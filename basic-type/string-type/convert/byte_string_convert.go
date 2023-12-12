package convert

import (
	"unsafe"
)

/**
*string-type(bytes)
*[]byte(str) 性能不佳
 */

func toBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func toString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func SliceByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StringToSliceByte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
