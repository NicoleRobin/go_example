package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"

func main() {
	fmt.Println("vim-go")
	db, err := sql.Open("mysql", "mailserver:mailserver@tcp(10.10.33.36)/mailserver")
	if err != nil {
		fmt.Printf("sql.Open failed, error:%s\n", err.Error())
		return
	}
	defer db.Close()

	stmtOut, err := db.Prepare("SELECT id FROM virtual_users where email=?")
	if err != nil {
		fmt.Printf("db.Prepare failed, error:%s\n", err.Error())
		return
	}
	defer stmtOut.Close()

	var squareNum int
	err = stmtOut.QueryRow("xl_test@xunlei.net").Scan(&squareNum)
	if err != nil {
		fmt.Println("stmtOut.QueryRow failed, error:%s\n", err.Error())
		return
	}
	fmt.Printf("The square number of 5 is: %d\n", squareNum)
}
