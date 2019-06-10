package handler

import (
	"cloud-storage/meta"
	"cloud-storage/util"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//返回html页面
		fmt.Printf("welcome to index")
		bytes, e := ioutil.ReadFile("./static/view/index.html")
		if e != nil {
			io.WriteString(w, "read index.html error")
			return
		}
		io.WriteString(w, string(bytes))
	} else {
		//接收文件流
		file, header, e := r.FormFile("file")
		if e != nil {
			io.WriteString(w, "read file error "+e.Error())
			return
		}
		defer file.Close()

		fileMeta := meta.FileMeta{
			FileName:   header.Filename,
			FilePath:   "/Users/sun7ay/Desktop/" + header.Filename,
			UploadTime: time.Now().Format("2006-01-02 15:04:05"),
		}

		create, e := os.Create(fileMeta.FilePath)
		if e != nil {
			fmt.Println(e)
			io.WriteString(w, "创建临时文件失败")
			return
		}
		defer create.Close()

		fileMeta.FileSize, e = io.Copy(create, file)
		if e != nil {
			io.WriteString(w, "复制文件失败")
			return
		}
		create.Seek(0, 0)
		fileMeta.FileSha1 = util.Filesha1(create)
		meta.InsertFileMetaInDB(fileMeta)
		fmt.Println(fileMeta.FileSha1)
		io.WriteString(w, "upload success")
	}
}

func FileInfoHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fileHash := r.Form["hash"][0]
	fMeta := meta.GetFileInfoFromDb(fileHash)
	res, e := json.Marshal(fMeta)
	if e == nil {
		w.Write(res)
	} else {
		io.WriteString(w, "json转换失败")
	}
}

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fileHash := r.Form["hash"][0]
	fMeta := meta.GetFileInfoFromDb(fileHash)
	file, e := os.Open(fMeta.FilePath)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()
	bytes, e := ioutil.ReadAll(file)

	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("content-disposition", "attachment;filename=\""+fMeta.FileName+"\"")
	w.Write(bytes)
}

func RenameFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form["name"][0]
	hash := r.Form["hash"][0]
	op := r.Form["op"][0]
	if op != "0" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fileMeta := meta.GetFileInfoFromDb(hash)
	fileMeta.FileName = name
	meta.UpdateFileNameFromDb(fileMeta.FileSha1, fileMeta.FileName)
	bytes, err := json.Marshal(fileMeta)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	fmt.Println(fileMeta)
	w.Write(bytes)
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	hash := r.Form["hash"][0]

	fileMeta := meta.GetFileInfoFromDb(hash)

	os.Remove(fileMeta.FilePath)

	meta.DeleteFileFromDB(hash)
}
