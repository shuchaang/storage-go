package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var db *sql.DB

func init() {
	db, _ := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/storage")
	db.SetMaxOpenConns(1000)
	e := db.Ping()
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}
	fmt.Println("init db conn")
}

func DBConn() *sql.DB {
	return db
}
