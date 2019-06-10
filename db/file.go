package db

import (
	"cloud-storage/db/mysql"
	"database/sql"
	"fmt"
)

func OnUploadFinish(filehash string, filename string, filesize int64, fileaddr string) bool {
	stmt, e := mysql.DBConn().Prepare("insert into tbl_file(file_sha1, file_name, file_addr, file_size,status)values (?,?,?,?,1)")
	if e != nil {
		fmt.Println(e.Error())
		return false
	}
	defer stmt.Close()
	_, e = stmt.Exec(filehash, filename, fileaddr, filesize)
	if e != nil {
		fmt.Println(e.Error())
		return false
	}
	return true
}

type FileTable struct {
	FileHash string
	FileName sql.NullString
	FilePath sql.NullString
	FileSize sql.NullInt64
}

func FindFileInfo(sha1 string) (*FileTable, error) {
	conn := mysql.DBConn()
	stmt, e := conn.Prepare("select file_sha1,file_name,file_size,file_addr from tbl_file where file_sha1=? and status=1 limit 1")
	if e != nil {
		return nil, e
	}
	defer stmt.Close()
	tfile := FileTable{}
	e = stmt.QueryRow(sha1).Scan(&tfile.FileHash, &tfile.FileName, &tfile.FileSize, &tfile.FilePath)
	if e != nil {
		return nil, e
	}
	return &tfile, nil
}

func UpdateFileInfo(sha string, name string) bool {
	conn := mysql.DBConn()
	stmt, e := conn.Prepare("update tbl_file set file_name=?,update_at=now() where file_sha1=?")
	if e != nil {
		fmt.Println(e)
		return false
	}
	defer stmt.Close()
	_, e = stmt.Exec(name, sha)
	if e != nil {
		fmt.Println(e)
		return false
	}
	return true
}

func DeleteFile(sha string) bool {
	conn := mysql.DBConn()
	stmt, e := conn.Prepare("update tbl_file set status=0,update_at=now() where file_sha1=?")
	if e != nil {
		fmt.Println(e)
		return false
	}
	defer stmt.Close()
	_, e = stmt.Exec(sha)
	if e != nil {
		fmt.Println(e)
		return false
	}
	return true
}
