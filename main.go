package main

import (
	"github.com/julienschmidt/httprouter"
	"movie/controller"
	"movie/driver"
	"movie/router"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/user",controller.CreateUser)

	// 用户登录
	router.POST("/api/login",controller.Login)
	return router
}


func main()  {
	// 初始化数据库
	driver.InitDB()

	// 初始化路由
	r := router.InitRouter()


	http.ListenAndServe(":8000",r)

}