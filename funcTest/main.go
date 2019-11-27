package main

import "fmt"

func main() {
	i1, i2 := f2()
	fmt.Println(i1, i2)

}

func f1() (int, int) {
	return 1, 3
}

func f2() (int, int) {
	return f1()
}
