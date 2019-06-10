package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var conn *sql.DB

func init() {
	conn, _ = sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/storage")
	conn.SetMaxOpenConns(1000)
	e := conn.Ping()
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(1)
	}
	fmt.Println("init db conn")
}

func DBConn() *sql.DB {
	return conn
}
