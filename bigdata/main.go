package main

import (
	"flag"
	"fmt"
	"strconv"
)

var bigData = flag.Bool("bigdata", false, "open big data analysis module")
var done = make(chan int)

func main() {
	flag.Parse()

	// 是否开启大数据分析模块
	if *bigData {
		fmt.Println("开启大数据分析模块")
		BigData.IsOpen = true
		go BigData.PostData()
	}

	for i := 0; i < 10000; i++ {
		BigData.PostChan <- []byte("test " + strconv.Itoa(i))
	}

	done <- 0
}
