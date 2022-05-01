# 视频网站

### httprouter使用
```go
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/user",controller.CreateUser)
	
	return router
}


func main()  {
	r := RegisterHandlers()

	http.ListenAndServe(":8000",r)

}
```