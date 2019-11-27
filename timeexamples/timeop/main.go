package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	// tstr := "2019-03-01T07:00:00+08:00"
	// tstr := "2019-03-01 T 07:00:00 00"
	// t, _ := GetTime(tstr)
	// fmt.Println(tstr)
	// fmt.Println(t)

	// t1, err := time.Parse("2006-01-02T15:04:05+08:00", tstr)
	// t1, err := time.Parse(tstr, tstr)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(t1)

	// t := time.Now().Format("2006-01-02 15:04:05")
	// tstr := time.Now().Format("2006-01-02")
	// fmt.Println(tstr)
	// t, _ := time.Parse("2006-01-02", tstr)
	// fmt.Println(t)
	// tt := time.Unix(1555459200, 0)
	// fmt.Println(tt)

	// t1, err := time.Parse("2006-01-02", t.String())
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// t := time.Now()
	// t1 := time.Now().In(time.Local)
	// fmt.Println(t)
	// fmt.Println(t1)
	// fmt.Println(t1.Location())
	// test3()
	// time.Sleep(20 * time.Second)
	test4()
}

// func GetTime(timeStr string) (time.Time, error) {
// 	loc, _ := time.LoadLocation("Local")
// 	// time, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
// 	time, _ := time.ParseInLocation(timeStr, timeStr, loc)
// 	return time, nil
// }

func test1() {
	test := "12:02:00"
	ts := strings.Split(test, ":")
	for _, v := range ts {
		temp, _ := strconv.Atoi(v)
		fmt.Println(temp)
	}
}

func test2() {
	test := "2007-01-02 14:12:09"
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation("2006-01-02 15:04:05", test, loc)
	fmt.Println(t, err)

	now := time.Now()
	fmt.Println(now)

	d := t.Sub(now)
	fmt.Println(d)

	d1 := now.Sub(t)
	fmt.Println(d1)
}

func test3() {
	d := time.Second * 2
	time.AfterFunc(d, printTest)
}

func printTest() {
	fmt.Println("this is a test for after func")
}

func test4() {
	limitTS := 1<<31 - 1
	fmt.Println(limitTS)
	t := time.Unix(int64(limitTS), 0)
	fmt.Println(t)
}
