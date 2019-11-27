package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomNumberGenerator(seed int64) *rand.Rand {
	// s1 := rand.NewSource(time.Now().UnixNano())
	s1 := rand.NewSource(seed)
	r1 := rand.New(s1)
	return r1
}

func main() {
	N := 10
	for i := 0; i < N; i++ {
		rng := RandomNumberGenerator(int64(i))
		fmt.Println(rng.Int())
	}

	fmt.Println("================")
	fmt.Println(time.Now().Unix())
}
