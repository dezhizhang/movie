package model

// 返回的数据类型

type Response struct {
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}
