package main

/*
#include <stdio.h>
*/
import "C"
import "unsafe"

func main() {
	buf := NewMyBuffer(1024)
	defer buf.Delete()

	copy(buf.Data(), "hello\x00")
	C.puts((*C.char)(unsafe.Pointer(&(buf.Data()[0]))))
}
