package xstring

import (
	"reflect"
	"unsafe"
)

// 参考文章： https://segmentfault.com/a/1190000037679588。

// ToString 更加快速的将[]byte转换为字符串。
func ToString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

// ToBytes 更加高效的将字符串转换字节数组。
func ToBytes(str string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&str))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
