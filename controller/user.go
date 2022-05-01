package controller

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

type User struct {

}

//创建用户

func CreateUser(w http.ResponseWriter, r *http.Request,p httprouter.Params)  {
	io.WriteString(w,"create user controller")
}

// 用户登录

func Login(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	username := p.ByName("username")
	password := p.ByName("password")

	fmt.Printf("username=%s,password=%s",username,password)
	io.WriteString(w,"用户登录成功")
}

