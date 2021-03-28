package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Code    int         `json:"code"`
	Result  interface{} `json:"result"`
	Message string      `json:"message"`
}

// // 第一种形式 方法
// type Context struct {
// 	Ctx *gin.Context
// }
// func (c *Context) Response(code int, data interface{}, message interface{}) {
// 	c.Ctx.JSON(http.StatusOK, response{
// 		Code:    code,
// 		Data:    data,
// 		Message: message,
// 	})
// }
// // 调用
// ctx := handler.Context{Ctx: c}
// ctx.Response(200, "health", "success")

// 第二种形式 函数
func ResponseOK(c *gin.Context, code int, result interface{}, message string) {
	c.JSON(http.StatusOK, response{
		Code:    code,
		Result:  result,
		Message: message,
	})
}

func ResponseUnauthorized(c *gin.Context, code int, result interface{}, message string) {
	c.JSON(http.StatusUnauthorized, response{
		Code:    code,
		Result:  result,
		Message: message,
	})
}

func ResponseBadRequest(c *gin.Context, code int, result interface{}, message string) {
	c.JSON(http.StatusBadRequest, response{
		Code:    code,
		Result:  result,
		Message: message,
	})
}

func ResponseFound(c *gin.Context, code int, result interface{}, message string) {
	c.JSON(http.StatusFound, response{
		Code:    code,
		Result:  result,
		Message: message,
	})
}

func ResponseError(c *gin.Context, code int, result interface{}, message string) {
	c.JSON(http.StatusInternalServerError, response{
		Code:    code,
		Result:  result,
		Message: message,
	})
}

// http.StatusBadRequest
// http.StatusOK
// http.StatusUnauthorized
// http.StatusInternalServerError
// http.StatusServiceUnavailable
// http.StatusMovedPermanently
// http.StatusFound
