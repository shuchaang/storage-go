package meta

import "fmt"

//文件元信息结构体
type FileMeta struct {
	FileSha1   string
	FileName   string
	FileSize   int64
	FilePath   string
	UploadTime string
}

var fileMetas map[string]FileMeta

func init() {
	fmt.Println("init fileMetas")
	fileMetas = make(map[string]FileMeta)
}

func UpdateFileMeta(f FileMeta) {
	fileMetas[f.FileSha1] = f
}

func GetFileMeta(sha1 string) FileMeta {
	return fileMetas[sha1]
}

func RemoveFileMeta(sha1 string) {
	delete(fileMetas, sha1)
}
