package Handlers

import (
	"NetworkDisk/Models/Redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Logout(pool *Redis.RedisPool) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid,err:=getUid(c)//获取UID
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"method":  "POST",
				"routing": "CheakJWT",
				"error": err.Error(),
			})
			return
		}

		err = pool.DEL(strconv.Itoa(uid))//清除Redis记录
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"method":  "POST",
				"routing": "CheakJWT",
				"error": err.Error(),
			})
			return
		}
	}
}
