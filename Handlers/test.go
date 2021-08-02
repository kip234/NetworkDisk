package Handlers

import (
	"NetworkDisk/Models"
	"NetworkDisk/Models/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJraXAiLCJleHAiOjAsInN1YiI6Ik5ldHdvcmtEaXNrIiwiYXVkIjoxMjMsIm5kZiI6MCwiaWF0IjowLCJqdGkiOjB9.uf/3D4ZBeevMMA4NUaTV2PIH2Rfy01Dz6zQsnl4pv5g=
func Test(pool redis.RedisPool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var u Models.User
		err:=c.ShouldBind(&u)
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"error":err.Error(),
			})
			return
		}
		err = pool.SET(strconv.Itoa(u.Uid),u.Pwd)
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"error":err.Error(),
			})
			return
		}
		Pwd,err:=pool.GET(strconv.Itoa(u.Uid))
		c.JSON(http.StatusOK,gin.H{
			"Uid":u.Uid,
			"Pwd":u.Pwd,
			"redis":Pwd,
		})

	}
}