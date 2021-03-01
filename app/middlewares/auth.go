package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var exceptionalURI = map[string] bool {
	"/api/login": true,
}

// 验证用户是否已登录
func Auth() gin.HandlerFunc {
	return func(c *gin.Context){
		processRequest(c)
		fmt.Println(c.FullPath())
		// 如果是例外的URI，则无需验证
		if exceptionalURI[c.FullPath()]{
			c.Next()
			processResponse(c)
			return
		}

		if token, _ := c.Cookie("token"); token != "abcd" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": 1, "msg": "Auth user login failed."})
			return
		}

		c.Next()
		processResponse(c)
	}
}

func processRequest(c *gin.Context) {
	fmt.Println("process request")
}

func processResponse(c *gin.Context) {
		fmt.Println("process response")
}