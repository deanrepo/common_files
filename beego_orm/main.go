package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

// Model Struct
type User struct {
	Id         int
	Name       string    `orm:"size(100)"`
	CreateTime time.Time `orm:"comment(创建时间)"`
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(User))

	// create table
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()

	ct := time.Now()
	ts := ct.Unix()
	ctU := time.Now().UTC()
	uts := ctU.Unix()
	fmt.Println(ts, uts)
	fmt.Println(ct, ctU)
	user := User{Name: "slene", CreateTime: ct}

	// insert
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	user.Name = "astaxie"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	u := User{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// // delete
	// num, err = o.Delete(&u)
	// fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
