package router

import (
	"github.com/julienschmidt/httprouter"
	"movie/controller"
)

func InitRouter() *httprouter.Router {
	router := httprouter.New()

	// 用户注册
	router.POST("/api/register",controller.Register)
	// 用户登录
	router.POST("/api/login",controller.Login)




	router.GET("/api/user",controller.CreateUser)
	return router
}
