package models

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math/rand"
)

type UserRegion struct {
	RegionId int
	User     string
}

type MailRegion struct {
	Id        int
	Ip1       string
	SmtpPort1 int
	PopPort1  int
	ImapPort1 int
	Ip2       string
	SmtpPort2 int
	PopPort2  int
	ImapPort2 int
	UserNum   int
	MaxUser   int
}

type User struct {
	Email string
	Count int
}

// func init() {
// 	orm.RegisterModel(new(MailRegion))
// 	orm.RegisterModel(new(UserRegion))
// }

// SELECT mail_region.* from mail_region, user_region where mail_region.ID = user_region.region_id and user_region.user = ?

func GetMailRegionByEmail(email string) (*MailRegion, error) {
	beego.Debug("In GetMailRegionByEmail, email:", email)
	var mailRegion MailRegion
	sql := "SELECT mail_region.* from mail_region, user_region where mail_region.ID = user_region.region_id and user_region.email = ?"
	beego.Debug("In GetMailRegionByEmail, sql:", sql)
	o := orm.NewOrm()
	err := o.Raw(sql, email).QueryRow(&mailRegion)
	if err != nil {
		beego.Error("orm.Raw().QueryRow failed, error:", err.Error())
		return nil, err
	}
	beego.Debug("In GetMailRegionByEmail, sql execute success, email:", email)
	return &mailRegion, nil
}

func InsertUserRegion(email string) (*MailRegion, error) {
	beego.Debug("In InsertUserRegion, email:", email)

	// 查询未满邮件分区
	var mailRegions []MailRegion
	sql := "select mail_region.* from (SELECT region_id, count(user_region.email) as user_num from user_region group by user_region.region_id) as t, mail_region where t.region_id=mail_region.id and t.user_num < mail_region.max_user"
	beego.Debug("In InsertUserRegion, sql:", sql)
	o := orm.NewOrm()
	num, err := o.Raw(sql).QueryRows(&mailRegions)
	if err != nil {
		beego.Error("orm.Raw().QueryRow failed, error:", err.Error())
		return nil, err
	}

	if num <= 0 {
		beego.Error("orm.Raw().QueryRow num:", num)
		return nil, errors.New("No valid mail_region!")
	}
	beego.Debug("In InsertUserRegion, mailRegion nums:", num)
	var i int64
	for i = 0; i < num; i = i + 1 {
		beego.Debug("In InsertUserRegion, SQL:", sql, " success, return mailRegion:ID:", mailRegions[i].Id, " Ip1:", mailRegions[i].Ip1, " Ip2:", mailRegions[i].Ip2, " MaxUser:", mailRegions[i].MaxUser)
	}
	randIdx := rand.Int63n(num)
	beego.Debug("In InsertUserRegion, randIdx:", randIdx)

	// 建立账号--邮件分区对应关系
	var mailRegion = mailRegions[randIdx]
	sql = "INSERT INTO user_region(region_id, email) VALUES(?, ?)"
	res, err := o.Raw(sql, mailRegion.Id, email).Exec()
	if err != nil {
		beego.Error("orm.Raw().Exec() failed, error:", err.Error())
		return nil, err
	}
	affectRow, _ := res.RowsAffected()
	beego.Debug("SQL:", sql, " execute success, affect row:", affectRow)

	return &mailRegion, nil
}

func Validate(email string, pass string) bool {
	beego.Debug("In Validate, email:", email)
	var user User
	sql := "select email, count(*) as count from virtual_users where email=? and md5(?)=password group by email"
	beego.Debug("In Validate, sql:", sql)
	o := orm.NewOrm()
	err := o.Raw(sql, email, pass).QueryRow(&user)
	if err != nil {
		beego.Error("orm.Raw().QueryRow failed, error:", err.Error())
		return false
	}
	if user.Count == 1 {
		beego.Error("In Validate, email:", email, ", correct password")
		return true
	} else {
		beego.Error("In Validate, email:", email, ", wrong password")
		return false
	}
}
