package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main func run")
	go dbDataOp()
	time.Sleep(24 * time.Second)

}

func dbDataOp() {
	fmt.Println("db data op run")
	tickerOp := time.NewTicker(5 * time.Second)
	tickerOp3 := time.NewTicker(3 * time.Second)

	for {
		select {
		case <-tickerOp.C:
			fmt.Println("5 sec")
		case <-tickerOp3.C:
			fmt.Println("3 sec")
		}
	}
}
