package Handlers

import (
	"NetworkDisk/Models/File"
	"NetworkDisk/Models/Redis"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func Upload(db *gorm.DB,pool *Redis.RedisPool) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid,err:=getUid(c)//获取UID
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"method":  "POST",
				"routing": "upload",
				"error": err.Error(),
			})
			return
		}
		file,err:=c.FormFile("file")//获取上传的文件
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"method":  "POST",
				"routing": "upload",
				"error": err.Error(),
			})
			return
		}
		//保存文件到本地
		err = c.SaveUploadedFile(file,"./files/"+strconv.Itoa(uid)+"/"+file.Filename)
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"method":  "POST",
				"routing": "upload",
				"error": err.Error(),
			})
			return
		}
		//记录
		f:=File.File{
			PathName: "./files/"+strconv.Itoa(uid)+"/"+file.Filename,
			Name: file.Filename,
			Size: file.Size,
			Owner: uid,
		}
		p:=File.Privilege{
			PathName: f.PathName,
			Owner: f.Owner,
			User: uid,
		}
		//保存记录到MySQL
		err=f.Save(db)
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"method":  "POST",
				"routing": "upload",
				"error": err.Error(),
			})
			return
		}
		err=p.Save(db)
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"method":  "POST",
				"routing": "upload",
				"error": err.Error(),
			})
			return
		}

		//缓存记录到Redis
		err = pool.SADD(File.OwnerKey(p.Owner),p.PathName)
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"method":  "POST",
				"routing": "upload",
				"error": err.Error(),
			})
			return
		}
		//缓存记录到Redis
		err = pool.SADD(File.UserKey(p.User),p.PathName)
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"method":  "POST",
				"routing": "upload",
				"error": err.Error(),
			})
			return
		}
	}
}

