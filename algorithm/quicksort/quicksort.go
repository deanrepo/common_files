package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var data [100]int
	for i := 0; i < 100; i++ {
		data[i] = rand.Intn(100)
	}
	fmt.Println(data)
	quicksort(data[:])
	fmt.Println(data)

}

func quicksort(data []int) {
	if len(data) <= 1 {
		return
	}

	left, right := 0, len(data)-1
	pivot := rand.Int() % len(data)

	data[pivot], data[right] = data[right], data[pivot]

	for i, v := range data {
		if v < data[right] {
			data[i], data[left] = data[left], data[i]
			left++
		}
	}

	data[left], data[right] = data[right], data[left]

	quicksort(data[:left])
	quicksort(data[left+1:])
}
