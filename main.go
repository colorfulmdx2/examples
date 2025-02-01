package main

import (
	"fmt"
	"unsafe"
)

type Secret struct {
	a int
	b int
}

func main() {
	s := Secret{10, 20}
	// Get pointer to struct
	ptr := unsafe.Pointer(&s)
	fmt.Println(&s.a)

	fmt.Println(ptr)

	// Get pointer to field `b` by adding an offset
	pb := (*string)(unsafe.Add(ptr, unsafe.Offsetof(s.b)))

	// Modify `b` directly in memory
	*pb = "hhhh"

	fmt.Println(s) // Output: {10 99}
}
