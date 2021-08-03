package Handlers

import (
	"NetworkDisk/Models/File"
	"NetworkDisk/Models/Redis"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Usesharedlinks(db *gorm.DB,pool *Redis.RedisPool) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid,err:=getUid(c)
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"method":  "GET",
				"routing": "usesharedlinks",
				"error":err.Error(),
			})
			return
		}

		link:=c.PostForm("link")
		owner,u,path,name,err:=Decoding(link)
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"method":  "GET",
				"routing": "usesharedlinks",
				"error":err.Error(),
			})
			return
		}
		//判断目标用户是否正确
		if uid!=u&&u!=0 {
			c.JSON(http.StatusOK,gin.H{
				"method":  "GET",
				"routing": "usesharedlinks",
				"error":"You cannot use this link",
			})
			return
		}
		//判断链接正确性
		v,err:=pool.SISMEMBER(File.OwnerKey(owner),path+name)
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"method":  "GET",
				"routing": "usesharedlinks",
				"error":err.Error(),
			})
			return
		}
		if v==0 {
			c.JSON(http.StatusOK,gin.H{
				"method":  "GET",
				"routing": "usesharedlinks",
				"error": "link error",
				"owner":owner,
				"u":u,
				"path":path,
				"name":name,
			})
			return
		}
		//后续操作
		pool.SADD(File.UserKey(uid),path+name)//刷新缓存
		p:=File.Privilege{PathName: path+name,Owner: owner}
		//db.Where(&p).Find(&p)//查找文件记录
		p.User=uid
		//p.Pri=0
		//fmt.Println(f)
		err = p.Save(db)//保存记录
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"method":  "GET",
				"routing": "usesharedlinks",
				"error":err.Error(),
			})
			return
		}
	}
}
