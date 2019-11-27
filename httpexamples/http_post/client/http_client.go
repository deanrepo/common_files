package main

import (
	"fmt"
	"io/ioutil"

	//    "log"
	"net/http"
	"net/url"

	//    "strings"
	"bytes"
	"encoding/json"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// test1()
	// test2()
	httpPostForm()
}

func test1() {
	resp, _ := http.Get("http://127.0.0.1:8080/?a=123456&b=aaa&b=bbb")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	var user User
	user.Name = "aaa"
	user.Age = 99
	if bs, err := json.Marshal(user); err == nil {
		//        fmt.Println(string(bs))
		req := bytes.NewBuffer([]byte(bs))
		tmp := `{"name":"junneyang", "age": 88}`
		req = bytes.NewBuffer([]byte(tmp))

		body_type := "application/json;charset=utf-8"
		resp, _ = http.Post("http://127.0.0.1:8080/test/", body_type, req)
		body, _ = ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	} else {
		fmt.Println(err)
	}

	client := &http.Client{}
	request, _ := http.NewRequest("GET", "http://127.0.0.1:8080/?a=123456&b=aaa&b=bbb", nil)
	request.Header.Set("Connection", "keep-alive")
	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	}

	req := `{"name":"junneyang", "age": 88}`
	req_new := bytes.NewBuffer([]byte(req))
	request, _ = http.NewRequest("POST", "http://127.0.0.1:8080/test/", req_new)
	request.Header.Set("Content-type", "application/json")
	response, _ = client.Do(request)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	}
}

func test2() {

	// POST方法一
	var user User
	user.Name = "张飞"
	user.Age = 100
	if bs, err := json.Marshal(user); err == nil {
		//        fmt.Println(string(bs))
		req := bytes.NewBuffer([]byte(bs))
		// tmp := `{"name":"junneyang", "age": 88}`
		// req = bytes.NewBuffer([]byte(tmp))

		body_type := "application/json;charset=utf-8"
		resp, _ := http.Post("http://127.0.0.1:8080/test/", body_type, req)
		body, _ := ioutil.ReadAll(resp.Body)
		// fmt.Println(string(body))

		if err := json.Unmarshal(body, &user); err == nil {
			fmt.Println(user.Name)
			user.Age += 100
			fmt.Println(user.Age)
		}

	} else {
		fmt.Println(err)
	}

	// POST方法二
	// client := &http.Client{}
	// req := `{"name":"junneyang", "age": 88}`
	// req_new := bytes.NewBuffer([]byte(req))
	// request, _ := http.NewRequest("POST", "http://127.0.0.1:8080/test/", req_new)
	// request.Header.Set("Content-type", "application/json;charset=utf-8")
	// response, _ := client.Do(request)
	// if response.StatusCode == 200 {
	// 	body, _ := ioutil.ReadAll(response.Body)
	// 	fmt.Println(string(body))
	// }
}

func httpPostForm() {
	resp, err := http.PostForm("http://127.0.0.1:8080/",
		url.Values{"key": {"李四"}, "id": {"123"}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
