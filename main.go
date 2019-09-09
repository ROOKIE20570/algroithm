package main

import (
	"C"
	"sync"
)
//#include <stdio.h>
var total struct {
	sync.Mutex
	value int64
}

func main() {

	println("hello cgo")

	C.puts(C.CString("Hello, World\n"))

}
