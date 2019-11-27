package main

import (
	"fmt"
	"log"

	"github.com/tealeg/xlsx"
)

func main() {
	userScore := make(map[int]int)
	excelFileName := "D:/goPath/src/common_files/excelop/foo.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Println(err)
		return
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			// for _, cell := range row.Cells {
			// 	text := cell.String()
			// 	fmt.Printf("%s\n", text)
			// }
			id, _ := row.Cells[0].Int()
			score, _ := row.Cells[1].Int()
			userScore[id] = score
		}
	}
	for k, v := range userScore {
		fmt.Printf("userId: %v, score: %v\n", k, v)
	}
}
