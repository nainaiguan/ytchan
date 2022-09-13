package util

import "unsafe"

func StringToByte(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func ByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
