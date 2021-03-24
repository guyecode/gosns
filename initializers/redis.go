package initializers

import (
	"fmt"
	"gosns/global"
	"github.com/garyburd/redigo/redis"
)

func InitReids() (redis.Conn, error){
	m := global.CONFIG.Redis
	dsn := fmt.Sprintf("%s:%d", m.Host, m.Port)
	c, err := redis.Dial("tcp", dsn)
	if err != nil {
		return nil, err
	}
	redis.String(c.Do("SET", "name", "GUYE"))
	global.REDIS = c
	return c, nil
}