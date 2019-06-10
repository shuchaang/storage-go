package handler

import (
	"cloud-storage/db"
	"cloud-storage/util"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	salt = "your password"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		username := r.Form["username"][0]
		password := r.Form["password"][0]
		if len(username) < 5 || len(password) < 5 {
			io.WriteString(w, "长度不够")
			return
		}
		pwd := util.Sha1([]byte(password + salt))
		up := db.UserSignUp(username, pwd)
		if up {
			io.WriteString(w, "success")
		} else {
			io.WriteString(w, "fail to regist user")
		}
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		username := r.Form["username"][0]
		password := r.Form["password"][0]
		pwd := util.Sha1([]byte(password + salt))
		if db.UserLogIn(username, pwd) {
			token := genToken(username)
			if db.UpdateToken(username, token) {
				w.Write([]byte("登录成功:" + token))
				return
			} else {
				w.Write([]byte("token获取失败"))
				return
			}
		} else {
			w.Write([]byte("登录失败"))
			return
		}

	}
}

func genToken(username string) string {
	//md5(username+timestamp+salt)+timestamp[:8] 40bit
	ts := fmt.Sprintf("%x", time.Now().Unix())
	prefix := util.MD5([]byte(username + ts + "_token_salt"))
	return prefix + ts[:8]
}
