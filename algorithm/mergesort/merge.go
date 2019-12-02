package main

import (
	"fmt"
	"math/rand"
)

func main() {
	test := make([]int, 10)
	for i := 0; i < len(test); i++ {
		test[i] = rand.Intn(100)
	}
	for idx, v := range test {
		fmt.Println(idx, v)
	}
	fmt.Println("================")
	ret := mergesort(test)
	for idx, v := range ret {
		fmt.Println(idx, v)
	}

}

func mergesort(data []int) []int {
	le := len(data)
	if le < 2 {
		return data
	}
	mid := le / 2
	lo := data[:mid]
	hi := data[mid:]
	lo = mergesort(lo)
	hi = mergesort(hi)
	return merge(lo, hi)
}

func merge(lo, hi []int) []int {
	l1 := len(lo)
	l2 := len(hi)
	newLe := l1 + l2
	ret := make([]int, newLe)
	i, j, k := 0, 0, 0
	for i < l1 && j < l2 {
		if lo[i] <= hi[j] {
			ret[k] = lo[i]
			i++
		} else {
			ret[k] = hi[j]
			j++
		}
		k++
	}
	for i < l1 {
		ret[k] = lo[i]
		i++
		k++
	}
	for j < l2 {
		ret[k] = hi[j]
		j++
		k++
	}
	return ret
}
