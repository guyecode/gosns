package routes

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"gosns/global"
	"math/rand"
	"net/http"
	"strconv"
)

func loadTestRoutes(engine *gin.Engine){
	engine.GET("/ping", ping)
	engine.GET("/calc", calc)
}

var count int
type Msg struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

func ping(c *gin.Context) {
	count++
	//data :=
	msg := Msg{}
	msg.Code = 0
	msg.Msg = "成功"
	msg.Data = make(map[string]interface{})
	msg.Data["count"] = count
	msg.Data["text"] = "测试Test。。。"
	msg.Data["lucky_number"] = rand.Intn(100 * 100)
	msg.Data["list"] = [...]string{"January", "February"}
	c.JSON(http.StatusOK, msg)
	fmt.Println(redis.String(global.REDIS.Do("GET", "name")))
	fmt.Println("ping ")
}

func calc(c *gin.Context) {
	i, _ := strconv.Atoi("0")
	result := 1 / i
	c.JSON(http.StatusOK, gin.H{"result": result})
}