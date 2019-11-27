package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"
)

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
	m := make(map[int]int)
	m[10] = 1
	m[20] = 1
	b := BlueDiamondGiftPackDb{
		Newbie:    1,
		DailyGift: 1,
		LevelGift: m,
		Special:   1,
	}

	// d, err := json.Marshal(b)
	d, err := b.MarshalJSON()
	CheckErr(err)
	fmt.Println(string(d))

	e := &BlueDiamondGiftPackDb{}

	err = e.UnmarshalJSON(d)
	CheckErr(err)
	fmt.Printf("%+v\n", e)

	err = b.Scan(d)
	CheckErr(err)

	v, err := b.Value()
	CheckErr(err)
	fmt.Printf("%+v\n", v)

}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
