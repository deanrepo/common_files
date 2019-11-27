package main

import (
	"fmt"
	"sync"
)

type SM struct {
	sm sync.Map
}

func newSm() *SM {
	return &SM{}
}

func main() {
	s := newSm()
	s.sm.Store("one", "one test")
	s.sm.Store("two", "two test")
	s.sm.Store("three", "three test")
	v, ok := s.sm.Load("one")
	if ok {
		fmt.Println(v)
	}
	fmt.Println()
	s.sm.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		if k.(string) == "two" {
			return false
		}
		return true
	})
}
