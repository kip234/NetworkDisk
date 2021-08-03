package Handlers

import (
	"NetworkDisk/Models/File"
	"NetworkDisk/Models/Redis"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Download(pool *Redis.RedisPool) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid,err:=getUid(c)//获取UID
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"method":  "GET",
				"routing": "register",
				"error":err.Error(),
			})
			return
		}

		path:=c.PostForm("filepath")//文件路径
		name:=c.PostForm("filename")//文件名
		//验证权限
		v,err:=pool.SISMEMBER(File.UserKey(uid),path+name)
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"method":  "GET",
				"routing": "register",
				"error":err.Error(),
			})
			return
		}
		if v==0{
			c.JSON(http.StatusOK,gin.H{
				"method":  "GET",
				"routing": "register",
				"error":"Unknown file",
			})
			return
		}
		//打开文件
		file,err:=os.Open(path+name)
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"method":  "GET",
				"routing": "register",
				"error":err.Error(),
			})
			return
		}
		defer file.Close()
		//准备传输
		c.Writer.Header().Add("Content-Disposition",fmt.Sprintf("attachment;filename=%s",name))
		c.File(path+name)
	}
}