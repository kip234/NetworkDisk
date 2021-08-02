package Handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

func Download() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid,err:=getUid(c)
		path:="./files/"+strconv.Itoa(uid)+c.PostForm("filepath")//文件路径
		name:=c.PostForm("filename")//文件名
		file,err:=os.Open(path+name)//打开文件
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