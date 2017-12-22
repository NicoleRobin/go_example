package controllers

import (
	"auth/models"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"strconv"
	"strings"
)

type AuthController struct {
	beego.Controller
}

/*
Request:
GET /auth HTTP/1.0
Host: localhost

Auth-Method: plain # plain/apop/cram-md5/external
Auth-User: user
Auth-Pass: password
Auth-Protocol: imap # imap/pop3/smtp
Auth-Login-Attempt: 1
Client-IP: 192.0.2.42
Client-Host: client.example.org

当auth-method=none时，会有以下头部字段
Auth-Smtp-Helo:[mail.zy.com]
Auth-Smtp-From:[mail from:zjw@zy.com]
Auth-Smtp-To:[rcpt to:xl_test@xunlei.net]


Good Response:
HTTP/1.0 200 OK
Auth-Status: OK
Auth-Server: 198.51.100.1
Auth-Port: 143

Bad Response:
HTTP/1.0 200 OK
Auth-Status: Invalid login or password
Auth-Wait: 3
*/
func (this *AuthController) Get() {
	beego.Debug("In Get,", this.Ctx.Input.Context.Request)

	// 验证账号和密码
	var user string
	if this.Ctx.Input.Header("Auth-Protocol") == "smtp" {
		user = this.Ctx.Input.Header("Auth-User")
		pass := this.Ctx.Input.Header("Auth-Pass")
		if models.Validate(user, pass) == false {
			beego.Error("Validate failed, email:", user, ", pass:", pass)
			this.Ctx.Output.Header("Auth-Status", "Invalid login or password")
			this.Ctx.Output.Header("Auth-Wait", strconv.Itoa(3))
			return
		}
	} else if this.Ctx.Input.Header("Auth-Protocol") == "pop3" {
		user = this.Ctx.Input.Header("Auth-User")
		pass := this.Ctx.Input.Header("Auth-Pass")
		if models.Validate(user, pass) == false {
			beego.Error("Validate failed, email:", user, ", pass:", pass)
			this.Ctx.Output.Header("Auth-Status", "Invalid login or password")
			this.Ctx.Output.Header("Auth-Wait", strconv.Itoa(3))
			return
		}
	} else if this.Ctx.Input.Header("Auth-Protocol") == "imap" {
		user = this.Ctx.Input.Header("Auth-User")
		pass := this.Ctx.Input.Header("Auth-Pass")
		if models.Validate(user, pass) == false {
			beego.Error("Validate failed, email:", user, ", pass:", pass)
			this.Ctx.Output.Header("Auth-Status", "Invalid login or password")
			this.Ctx.Output.Header("Auth-Wait", strconv.Itoa(3))
			return
		}
	} else if this.Ctx.Input.Header("Auth-Protocol") == "none" {
		to_user := this.Ctx.Input.Header("Auth-Smtp-To")
		rcpt_to := strings.Split(to_user, ":")
		user = rcpt_to[1]
	}

	// 查询账号所属邮件分区
	mailRegion, err := models.GetMailRegionByEmail(user)
	if err != nil {
		beego.Error("In Get, GetMailRegionByUser failed for user:", user, " error:", err.Error())
		// 未查到账号--邮件分区信息，需要为该用户分配邮箱分区并建立对应关系
		mailRegion, err = models.InsertUserRegion(user)
		if err != nil {
			beego.Error("In Get, models.InsertUserRegion failed for user:", user, " error:", err.Error())
		}
	}

	if err == nil {
		beego.Debug("In Get, User:", user, " assign MailRegion ip1:", mailRegion.Ip1, " ip2:", mailRegion.Ip2)
		randIdx := rand.Intn(2) + 1
		beego.Debug("In Get, randIdx:", randIdx)
		if this.Ctx.Input.Header("Auth-Protocol") == "smtp" {
			this.Ctx.Output.Header("Auth-Status", "OK")
			if randIdx == 1 {
				this.Ctx.Output.Header("Auth-Server", mailRegion.Ip1)
				this.Ctx.Output.Header("Auth-Port", strconv.Itoa(mailRegion.SmtpPort1))
			} else if randIdx == 2 {
				this.Ctx.Output.Header("Auth-Server", mailRegion.Ip2)
				this.Ctx.Output.Header("Auth-Port", strconv.Itoa(mailRegion.SmtpPort2))
			}
		} else if this.Ctx.Input.Header("Auth-Protocol") == "pop3" {
			this.Ctx.Output.Header("Auth-Status", "OK")
			if randIdx == 1 {
				this.Ctx.Output.Header("Auth-Server", mailRegion.Ip1)
				this.Ctx.Output.Header("Auth-Port", strconv.Itoa(mailRegion.PopPort1))
			} else if randIdx == 2 {
				this.Ctx.Output.Header("Auth-Server", mailRegion.Ip2)
				this.Ctx.Output.Header("Auth-Port", strconv.Itoa(mailRegion.PopPort2))
			}
		} else if this.Ctx.Input.Header("Auth-Protocol") == "imap" {
			this.Ctx.Output.Header("Auth-Status", "OK")
			if randIdx == 1 {
				this.Ctx.Output.Header("Auth-Server", mailRegion.Ip1)
				this.Ctx.Output.Header("Auth-Port", strconv.Itoa(mailRegion.ImapPort1))
			} else if randIdx == 2 {
				this.Ctx.Output.Header("Auth-Server", mailRegion.Ip2)
				this.Ctx.Output.Header("Auth-Port", strconv.Itoa(mailRegion.ImapPort2))
			}
		} else if this.Ctx.Input.Header("Auth-Protocol") == "none" {
			this.Ctx.Output.Header("Auth-Status", "OK")
			if randIdx == 1 {
				this.Ctx.Output.Header("Auth-Server", mailRegion.Ip1)
				this.Ctx.Output.Header("Auth-Port", strconv.Itoa(mailRegion.ImapPort1))
			} else if randIdx == 2 {
				this.Ctx.Output.Header("Auth-Server", mailRegion.Ip2)
				this.Ctx.Output.Header("Auth-Port", strconv.Itoa(mailRegion.ImapPort2))
			}
		} else {
			// error Auth-Protocol
			beego.Error("Error Auth-Protocol:", this.Ctx.Input.Header("Auth-Protocol"))
			this.Ctx.Output.Header("Auth-Status", "Error Auth-Protocol")
			this.Ctx.Output.Header("Auth-Wait", strconv.Itoa(3))
		}
	} else {
		beego.Error("In Get, GetMailRegionByUser and InsertUserRegion both failed for user:", user)
		this.Ctx.Output.Header("Auth-Status", err.Error())
		this.Ctx.Output.Header("Auth-Wait", strconv.Itoa(3))
	}
}
