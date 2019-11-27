package main

import "fmt"

func main() {
	var buffer [256]byte

	slice := buffer[10:20]
	fmt.Println("Before: len(slice) =", len(slice))
	PtrSubtractOneFromLength(&slice)
	fmt.Println("After:  len(slice) =", len(slice))
}

func PtrSubtractOneFromLength(slicePtr *[]byte) {
	slice := *slicePtr
	*slicePtr = slice[0 : len(slice)-1]
}
