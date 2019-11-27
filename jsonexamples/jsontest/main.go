package main

import (
	"encoding/json"
	"fmt"
)

type PostResust struct {
	Code   int    `json:"code"` // 上传数据返回的结果状态码，
	Msg    string `josn:"msg"`  // 上传数据返回的消息
	Nums   []int  `json:"nums"`
	Person `json:"person"`
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	format := `{%s}`
	tStr := fmt.Sprintf(format, "")
	fmt.Println(tStr)

	p := PostResust{
		Code: 0,
		Nums: []int{1, 2, 4, 5},
	}

	s, _ := json.Marshal(p)
	fmt.Println(string(s))

	s1 := `{"code": 12,"person":{}}`
	p1 := &PostResust{}
	err := json.Unmarshal([]byte(s1), p1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("p1: %+v\n", *p1)
}
