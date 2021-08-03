package Handlers

import (
	"NetworkDisk/Models/File"
	"NetworkDisk/Models/Redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Filelist(pool *Redis.RedisPool) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid,err:=getUid(c)
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"method":  "GET",
				"routing": "filelist",
				"error":err.Error(),
			})
			return
		}
		v,err:=pool.SMEMBERS(File.UserKey(uid))
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"method":  "GET",
				"routing": "filelist",
				"error":err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"files":v,
		})
	}
}
