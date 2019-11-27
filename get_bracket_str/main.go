package main

import (
	"flag"
	"fmt"

	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"unicode"
)

var (
	dirpath = flag.String("dirpath", ".", "directory path to scan")
	outfile = flag.String("outputfilepath", "text.txt", "output file path")
	write   *os.File
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

	ScanAll(*dirpath)

}

//ScanAll 扫描所有文件并提取每行中的所有中文字符
func ScanAll(dirpath string) {
	files, _ := ioutil.ReadDir(dirpath)
	outputFileName := path.Base(*outfile) //获取文件名带后缀
	// fmt.Println("output file name:", outputFileName)
	for _, fi := range files {
		// fmt.Println(fi.Name()) // debug
		if fi.Name() != outputFileName && !strings.Contains(fi.Name(), "XXX") && !strings.Contains(fi.Name(), "text") {
			if fi.IsDir() {
				ScanAll(dirpath + "/" + fi.Name())
			} else {
				//只扫描go文件
				if strings.Contains(fi.Name(), ".txt") {
					file, _ := os.Open(dirpath + "/" + fi.Name())
					buffer := bufio.NewReader(file)
					flag := true
					lineNum := 0
					for {
						lineNum++
						s, _, ok := buffer.ReadLine()
						canWrite := true
						if strings.Contains(string(s), "//") {
							canWrite = false
						}
						if strings.Contains(string(s), "/*") {
							flag = false
						}
						if ok == io.EOF {
							break
						}
						if flag && canWrite && ContainChineseChar(string(s)) {
							// // get text in the parenthesis
							// ret := getTextFromParenthese(string(s))
							// write.WriteString(ret + "\n")

							// get text that contain '('
							if strings.Contains(string(s), "(") {
								ret := string(s)
								write.WriteString(ret + "\n")
							}
						}
						if strings.Contains(string(s), "*/") {
							flag = true
						}
					}
					file.Close()
				}
			}
		}
	}
}

//GetAllChineseCharacter 提取每行中的所有中文字符
func GetAllChineseCharacter(line string) string {
	var r []rune
	if line != "" {
		r = []rune(line)
	}

	cnstr := ""
	for i := 0; i < len(r); i++ {
		if r[i] <= 40869 && r[i] >= 19968 {
			cnstr = cnstr + string(r[i])
		}
	}

	return cnstr
}

//ContainChineseChar 扫描字符串中是否含有中文字符
func ContainChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

// get string text from the parenthesis
func getTextFromParenthese(str string) string {
	sIdx := strings.Index(str, "(")
	eIdx := strings.Index(str, ")")
	ret := str[sIdx+1 : eIdx]
	return ret
}
