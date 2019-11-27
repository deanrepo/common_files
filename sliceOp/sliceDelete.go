package main

//遍历slice时删除其中的元素

import "fmt"

func main() {

}

func m1() {
	var x = []int{90, 15, 81, 87, 47, 59, 81, 18, 25, 40, 56, 8}
	y := x[:0]
	for _, n := range x {
		if n%2 == 0 {
			y = append(y, n)
		}
	}
	fmt.Println(x)
}

func m2() {
	var x = []int{90, 15, 81, 87, 47, 59, 81, 18, 25, 40, 56, 8}

	i := 0
	l := len(x)
	for i < l {
		if x[i]%2 != 0 {
			x = append(x[:i], x[i+1:]...)
			l--
		} else {
			i++
		}
	}
	x = x[:i]

	fmt.Println(x)
	// [90 18 40 56 8]

}
