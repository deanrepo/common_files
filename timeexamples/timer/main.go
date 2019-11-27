package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	test5()
}

func test1() {
	input := make(chan interface{})

	//producer - produce the messages
	go func() {
		for i := 0; i < 5; i++ {
			input <- i
		}
		input <- "hello, world"
	}()

	t1 := time.NewTimer(time.Second * 5)
	t2 := time.NewTimer(time.Second * 10)

	for {
		select {
		//consumer - consume the messages
		case msg := <-input:
			fmt.Println(msg)

		case <-t1.C:
			println("5s timer")
			t1.Reset(time.Second * 5)

		case <-t2.C:
			println("10s timer")
			t2.Reset(time.Second * 10)
		}
	}
}

func test2() {

	requests := make(chan int, 5)
	for i := 1; i <= 2; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests { //会循环两次，前面往requests channel中发送了两个值
		<-limiter //执行到这里，需要隔 200毫秒才继续往下执行，time.Tick(timer)上面已定义
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now() //这里burstyLimiter channel 连续发送了三次值
	}

	go func() {
		for t := range time.Tick(time.Second * 2) {
			burstyLimiter <- t // 这里往burstyLimiter channel 循环发送time.Tick
		}
	}()
	burstyRequests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		burstyRequests <- i
		<-burstyLimiter //前三次没有速度限制，会直接打印出后面的println的内容
		fmt.Println("request", i, time.Now())
	}

	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter //继续接收burstyLimiter值，除了前三次，后面的都是time.Tick进行速度限制,2秒打印一次，直到此次循环结束
		fmt.Println("request", req, time.Now())
	}
}

// var tPtr *time.Timer

var tPtr *time.Ticker

func CheckoutTimer() bool {
	<-tPtr.C
	return true
}

func test3() {
	// tPtr = time.NewTimer(time.Duration(int(time.Second) * 5))
	tPtr = time.NewTicker(time.Duration(int(time.Second) * 5))
	go func() {
		if CheckoutTimer() {
			fmt.Println("时间到")
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
}

func test4() {
	c := time.Tick(1 * time.Second)
	for now := range c {
		fmt.Printf("%v\n", now)
	}
}

// test for timer and select
func test5() {
	t1 := time.NewTimer(time.Second * 5)
	// block until receive timer singal
	select {
	case <-t1.C:
		println("10s timer")
	}
	fmt.Println("timer end")
}
