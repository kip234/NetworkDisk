package Handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Upload() gin.HandlerFunc {
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

		err = c.SaveUploadedFile(file,"./files/"+strconv.Itoa(uid)+"/"+file.Filename)//保存
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

