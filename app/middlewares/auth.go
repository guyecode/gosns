package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context){
		// 验证用户是否已登录
		fmt.Println("process request")
		c.Next()
		fmt.Println("process response")
	}
}