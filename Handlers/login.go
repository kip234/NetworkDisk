package Handlers

import (
	"NetworkDisk/Models/jwt"
	"NetworkDisk/Models/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Login(redis redis.RedisPool,template jwt.Jwt) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid,err:=getUid(c)//获取UID
		if err!=nil {
			c.JSON(http.StatusOK,gin.H{
				"method":  "POST",
				"routing": "login",
				"error":err.Error(),
			})
			return
		}
		template.Payload.Aud=uid//更新Payload值
		tocken := template.Encoding()//计算理应正确的JWT
		redis.SET(strconv.Itoa(uid),tocken)//放入Redis
		c.JSON(http.StatusOK,gin.H{//把令牌返回给客户
			"tocken":tocken,
		})
	}
}