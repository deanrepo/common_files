package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// 获取现在的时间，秒数设为零
	fmt.Println("set second to zero")
	tstr1 := now.Format("2006-01-02 15:04:00")
	fmt.Println(tstr1)
	t1, err := time.Parse("2006-01-02 15:04:05", tstr1)
	if err != nil {
		panic(err)
	}
	fmt.Println("t1=>", t1)

	// 获取现在的时间，分钟设为零
	fmt.Println("set minute to zero")
	tstr2 := now.Format("2006-01-02 15:00:05")
	fmt.Println(tstr2)
	t2, err := time.Parse("2006-01-02 15:04:05", tstr2)
	if err != nil {
		panic(err)
	}
	fmt.Println("t2=>", t2)

	// 获取现在的时间，小时设为零
	fmt.Println("set hour to zero")
	tstr3 := now.Format("2006-01-02 00:04:05")
	fmt.Println(tstr3)
	t3, err := time.Parse("2006-01-02 15:04:05", tstr3)
	if err != nil {
		panic(err)
	}
	fmt.Println("t3=>", t3)

	// 获取现在的时间，小时和秒数设为零
	fmt.Println("set hour and second to zero")
	tstr4 := now.Format("2006-01-02 00:04:00")
	fmt.Println(tstr4)
	t4, err := time.Parse("2006-01-02 15:04:05", tstr4)
	if err != nil {
		panic(err)
	}
	fmt.Println("t4=>", t4)

	fmt.Println("##############################")
	fmt.Println(now.Hour(), now.Minute(), now.Second())

}
