package db

import (
	"cloud-storage/db/mysql"
	"fmt"
)

func UserSignUp(username string, pwd string) bool {
	stmt, e := mysql.DBConn().Prepare("insert into tbl_user(user_name,user_pwd)values(?,?)")
	if e != nil {
		fmt.Println(e)
		return false
	}
	defer stmt.Close()
	result, e := stmt.Exec(username, pwd)
	if e != nil {
		fmt.Println(e)
		return false
	}
	if i, e := result.RowsAffected(); e == nil && i > 0 {
		return true
	} else {
		fmt.Println(e)
		return false
	}
}

func UserLogIn(username string, pwd string) bool {
	stmt, e := mysql.DBConn().Prepare("select * from tbl_user where user_name=? and user_pwd=?")
	if e != nil {
		fmt.Println(e)
		return false
	}
	rows, e := stmt.Query(username, pwd)
	if e != nil {
		fmt.Println(e)
		return false
	}
	if rows == nil {
		fmt.Println(username + "不存在")
		return false
	}
	return true
}

func UpdateToken(username string, token string) bool {
	stmt, e := mysql.DBConn().Prepare("replace into tbl_user_token(user_name, user_token)values(?,?)")
	if e != nil {
		fmt.Println(e)
		return false
	}
	defer stmt.Close()
	_, e = stmt.Exec(username, token)
	if e != nil {
		fmt.Println(e)
		return false
	}

	return true
}
