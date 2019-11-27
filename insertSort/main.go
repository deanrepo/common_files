package main

import "fmt"

func main() {
	s := []int{9, 0, 6, 5, 8, 2, 1, 7, 4, 3}
	fmt.Println(s)
	InsertionSort(s)
	fmt.Println(s)
}

//插入排序
func InsertionSort(s []int) {
	n := len(s)
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0 && s[j] > s[j-1]; j-- {
			swap(s, j, j-1)
		}
	}
}

func swap(slice []int, i int, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
