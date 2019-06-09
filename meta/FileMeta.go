package meta

import (
	"cloud-storage/db"
	"fmt"
)

//文件元信息结构体
type FileMeta struct {
	FileSha1   string
	FileName   string
	FileSize   int64
	FilePath   string
	UploadTime string
}

func InsertFileMetaInDB(f FileMeta) {
	db.OnUploadFinish(f.FileSha1, f.FileName, f.FileSize, f.FilePath)
}

func GetFileInfoFromDb(sha1 string) FileMeta {
	table, e := db.FindFileInfo(sha1)
	if e != nil {
		fmt.Println(e)
		return FileMeta{}
	}
	meta := FileMeta{
		FileSha1: sha1,
		FileName: table.FileName.String,
		FileSize: table.FileSize.Int64,
		FilePath: table.FilePath.String,
	}
	return meta
}

func UpdateFileNameFromDb(sha string, name string) bool {
	return db.UpdateFileInfo(sha, name)
}

func DeleteFileFromDB(sha string) bool {
	return db.DeleteFile(sha)
}
