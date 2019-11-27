package main

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/* var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:123456tcp(192.168.5.201:3306)/test?collation=utf8_general_ci")
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
} */

//蓝钻礼包领取状态 0-表示未领取
type BlueDiamondGiftPackDb struct {
	Newbie         int
	DailyGift      int
	DailyYearGift  int
	DailySuperGift int
	LevelGift      map[int]int
	Special        int
}

func (this BlueDiamondGiftPackDb) MarshalJSON() ([]byte, error) {
	b := struct {
		Newbie         int
		DailyGift      int
		DailyYearGift  int
		DailySuperGift int
		LevelGift      map[int]int
		Special        int
	}{
		Newbie:         this.Newbie,
		DailyGift:      this.DailyGift,
		DailyYearGift:  this.DailyYearGift,
		DailySuperGift: this.DailySuperGift,
		LevelGift:      this.LevelGift,
		Special:        this.Special,
	}
	return json.Marshal(b)
}

func (this *BlueDiamondGiftPackDb) UnmarshalJSON(data []byte) error {
	b := &struct {
		Newbie         int
		DailyGift      int
		DailyYearGift  int
		DailySuperGift int
		LevelGift      map[int]int
		Special        int
	}{}

	err := json.Unmarshal(data, b)
	if err != nil {
		log.Println("err:", err)
	}

	this.Newbie = b.Newbie
	this.DailyGift = b.DailyGift
	this.DailyYearGift = b.DailyYearGift
	this.DailySuperGift = b.DailySuperGift
	this.LevelGift = b.LevelGift
	this.Special = b.Special

	return err
}

func (this *BlueDiamondGiftPackDb) Scan(value interface{}) error {
	return this.UnmarshalJSON(value.([]byte))
}

func (this BlueDiamondGiftPackDb) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func main() {
	// insert()
	query()

	/* 	m := make(map[int]int)
	   	m[10] = 1
	   	m[20] = 1
	   	b := BlueDiamondGiftPackDb{
	   		Newbie:    1,
	   		DailyGift: 1,
	   		LevelGift: m,
	   		Special:   1,
	   	}

	   	bstr, err := b.MarshalJSON()
	   	checkErr(err)
	   	fmt.Printf("bstr len: %d\n", len(bstr)) */

}

//查询demo
func query() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?collation=utf8_general_ci")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM blue_test")
	checkErr(err)

	//普通demo
	//for rows.Next() {
	//    var userId int
	//    var userName string
	//    var userAge int
	//    var userSex int

	//    rows.Columns()
	//    err = rows.Scan(&userId, &userName, &userAge, &userSex)
	//    checkErr(err)

	//    fmt.Println(userId)
	//    fmt.Println(userName)
	//    fmt.Println(userAge)
	//    fmt.Println(userSex)
	//}

	//字典类型
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
}

//插入demo
func insert() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?collation=utf8_general_ci")
	checkErr(err)

	m := make(map[int]int)
	m[10] = 1
	m[20] = 1
	b := BlueDiamondGiftPackDb{
		Newbie:    1,
		DailyGift: 1,
		LevelGift: m,
		Special:   1,
	}

	bstr, err := b.MarshalJSON()
	checkErr(err)
	fmt.Printf("bstr len: %d\n", len(bstr))

	stmt, err := db.Prepare(`INSERT blue_test (id,info) values (?,?)`)
	checkErr(err)
	res, err := stmt.Exec(1, bstr)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
