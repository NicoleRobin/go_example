package main

import "github.com/astaxie/beego/orm"
import _ "github.com/go-sql-driver/mysql"
import "fmt"

type MailRegion struct {
	Region_id int
	Email     string
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "mailserver:mailserver@tcp(10.10.33.36)/mailserver")
}
func main() {
	fmt.Println("vim-go")
	o := orm.NewOrm()
	var mailRegions []MailRegion
	r := o.Raw("SELECT * FROM user_region where email=?", "x_test@xunlei.net")
	num, err := r.QueryRows(&mailRegions)
	if err != nil {
		fmt.Printf("orm.Raw().QueryRow() failed, error:%s\n", err.Error())
		return
	}
	fmt.Println("Num:", num)

	// fmt.Printf("MailRegion.RegionID:%d, MailRegion.Email:%s\n", mailRegion.Region_id, mailRegion.Email)
}
