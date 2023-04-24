package go_file_util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/* 通用响应方法 */
func R(c *gin.Context, err error, msg interface{}, data interface{}) {
	// 如果有err,就说明是有错误,就返回错误响应(msg)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    msg,
		})
		return
	}
	// 返回正确响应(data)
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    data,
	})
	return
}
