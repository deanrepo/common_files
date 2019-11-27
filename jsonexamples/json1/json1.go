package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	test := make(map[int]*Person)
	p := &Person{}
	p.Name = "Dean"
	p.Age = 20

	p1 := &Person{}
	p1.Name = "jerry"
	p1.Age = 30

	test[1] = p
	test[2] = p1

	t, err := json.Marshal(test)

	fmt.Println(string(t), err)

	t1 := make(map[int]*Person)

	err = json.Unmarshal(t, &t1)
	if err != nil {
		panic(err)
	}
	for k, v := range t1 {
		fmt.Println(k, v)
	}

}
