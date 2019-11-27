package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// rand.Seed(time.Now().UnixNano())
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		x := rand.Intn(100)
		fmt.Println(x)
	}

	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())

	rand.Seed(12)
	for i := 0; i < 10; i++ {
		x := rand.Intn(100)
		fmt.Println(x)
	}

	randS := rand.NewSource(time.Now().UnixNano())
	fmt.Println(randS.Int63())

}
