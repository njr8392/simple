package main

import "fmt"
import "unsafe"

func main() {
	vals := 0x1122334455667788
	start := unsafe.Pointer(&vals)
	for i := uintptr(0); i < unsafe.Sizeof(vals); i++ {
		item := *(*byte)(unsafe.Pointer(uintptr(start) + uintptr(i)))
		fmt.Printf("%x ", item)
	}

	fmt.Printf("\n")
}
