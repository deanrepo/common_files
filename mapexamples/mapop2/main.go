package main

// map可以边遍历边删除其中的元素

import (
	"fmt"
)

func main() {
	query := map[string]string{}

	query["test0"] = "0"
	query["test1"] = "1"
	query["test2"] = "2"

	i := 0
	for k, v := range query {
		delete(query, "test2")
		fmt.Println(query, k, v)
		i++
	}

}
