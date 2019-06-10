package db

import (
	"cloud-storage/db/mysql"
	"fmt"
	"time"
)

type UserFile struct {
	UserName   string
	FileHash   string
	FileName   string
	FileSize   int
	UploadAt   string
	LastUpdate string
}

func InsertUserFile(username, filename, hash string, size int) bool {
	stmt, e := mysql.DBConn().Prepare("insert into tbl_user_file(user_name, file_sha1, file_size, file_name, upload_at, status) values (?,?,?,?,?,1)")
	if e != nil {
		fmt.Println(e)
		return false
	}
	defer stmt.Close()
	_, e = stmt.Exec(username, hash, size, filename, time.Now())
	if e != nil {
		fmt.Println(e)
		return false
	}
}
