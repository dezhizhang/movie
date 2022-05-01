package main

import (
	"github.com/julienschmidt/httprouter"
	"movie/controller"
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
	r := RegisterHandlers()

	http.ListenAndServe(":8000",r)

}