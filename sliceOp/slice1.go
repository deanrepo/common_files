package main

import "fmt"

func main() {
	// testSubtract()
	test := make([]int, 1)
	fmt.Println(test)
	fmt.Printf("%+v\n", &test)
}

func testAddOne() {
	var buffer [256]byte

	slice := buffer[10:20]
	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}
	fmt.Println("before", slice)
	AddOneToEachElement(slice)
	fmt.Println("after", slice)
}

func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}

func testSubtract() {
	var buffer [256]byte

	slice := buffer[10:20]
	fmt.Println("Before: len(slice) =", len(slice))
	newSlice := SubtractOneFromLength(slice)
	fmt.Println("After:  len(slice) =", len(slice))
	fmt.Println("After:  len(newSlice) =", len(newSlice))
}

func SubtractOneFromLength(slice []byte) []byte {
	slice = slice[0 : len(slice)-1]
	return slice
}
