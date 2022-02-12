// @file: main.go
// @date: 2022/02/12

package main

// #include "hello.h"
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var contentAddr *C.char

	length := C.readFile(C.CString("config.yml"), &contentAddr)
	var content []byte
	var contentHdr = (*reflect.SliceHeader)(unsafe.Pointer(&content))
	contentHdr.Data = uintptr(unsafe.Pointer(contentAddr))
	contentHdr.Len = int(length)
	contentHdr.Cap = int(length)
	fmt.Println(string(content))
}
