package Handlers

import (
	"NetworkDisk/Models/File"
	"NetworkDisk/Models/Redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Getsharelinks(pool *Redis.RedisPool) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid,err:=getUid(c)
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"method":  "GET",
				"routing": "getsharelinks",
				"error":err.Error(),
			})
			return
		}

		path:=c.PostForm("filepath")
		name:=c.PostForm("filename")
		id,err:=strconv.Atoi(c.PostForm("uid"))//目标用户
		if err!=nil {
			c.JSON(http.StatusOK,gin.H{
				"method":  "GET",
				"routing": "getsharelinks",
				"error":"uid error",
			})
			return
		}

		v,err:=pool.SISMEMBER(File.OwnerKey(uid),path+name)//验证所有权
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"method":  "GET",
				"routing": "getsharelinks",
				"error":err.Error(),
			})
			return
		}
		if v==0 {
			c.JSON(http.StatusOK,gin.H{
				"method":  "GET",
				"routing": "getsharelinks",
				"error":"Unknown file",
			})
			return
		}

		link:=Encoding(uid,id,path,name)
		c.JSON(http.StatusOK,gin.H{
			"method":  "GET",
			"routing": "getsharelinks",
			"link":link,
		})
	}
}