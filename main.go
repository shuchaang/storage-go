package main

import (
	"cloud-storage/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/info", handler.FileInfoHandler)
	http.HandleFunc("/file/download", handler.DownloadFile)
	http.HandleFunc("/file/rename", handler.RenameFile)
	http.HandleFunc("/file/delete", handler.DeleteFile)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
