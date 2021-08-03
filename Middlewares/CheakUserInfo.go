package Middlewares

import (
	"NetworkDisk/Models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CheakUserInfo(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Models.User
		err:=c.ShouldBind(&user)
		if err!=nil{//绑定失败
			c.JSON(http.StatusOK,gin.H{
				"method":  "POST",
				"Middlewares": "CheakUserInfo",
				"error":err.Error(),
			})
			c.Abort()
			return
		}
		tmp:= Models.User{}
		if err:=tmp.Load(db,user.Uid);err!=nil{//查找失败
			c.JSON(http.StatusOK,gin.H{
				"method":  "POST",
				"Middlewares": "CheakUserInfo",
				"error":err.Error(),
			})
			c.Abort()
			return
		}
		if tmp.Pwd!=user.Pwd{//密码对不上
			c.JSON(http.StatusOK,gin.H{
				"method":  "POST",
				"Middlewares": "CheakUserInfo",
				"error":"password wrong !",
			})
			c.Abort()
			return
		}
		c.Set("uid",tmp.Uid)//验证通过,存入UID
		//设置cookie
		//c.SetCookie("Uid",strconv.Itoa(user.Uid),AvailableLimit,"/","localhost",false,true)
	}
}