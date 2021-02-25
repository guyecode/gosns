package routes

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

func loadTestRoutes(engine *gin.Engine){
	engine.GET("ping", Ping)
}

var count int

func Ping(c *gin.Context) {
	count++
	var msg struct {
		Code int            `json:"code"`
		Msg  string         `json:"msg"`
		Data map[string]int `json:"data"`
	}
	//data :=
	msg.Code = 0
	msg.Msg = "成功"
	msg.Data = map[string]int{"count": count}
	msg.Data["lucky_number"] = rand.Intn(100 * 100)
	c.JSON(http.StatusOK, msg)
}