package main

import (
	"flag"
	"fmt"
	"os"
)

var write *os.File

var (
	dirpath = flag.String("dirpath", ".", "directory path to scan")
	outfile = flag.String("outputfilepath", "chinese.txt", "output file path")
)

func main() {
	flag.Parse()

	defer write.Close()
	var err error
	write, err = os.OpenFile(*outfile, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println("output file read err:", err)
		return
	}
	// 	write.WriteString(
	// 		`package base

	// const (
	// `)
	ScanAll(*dirpath)
	// 	write.WriteString(
	// 		`
	// )
	// 	`)
}
