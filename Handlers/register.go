//暂时没弄
package Handlers

import (
	"NetworkDisk/Models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strconv"
)

//注册成狗后会返回UID
func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Models.User
		err:=c.ShouldBind(&user)//获取用户相关信息
		if err!=nil {//绑定出错
			c.JSON(http.StatusBadRequest,gin.H{
				"method":  "POST",
				"routing": "register",
				"Error":err.Error(),
			})
			return
		}
		err=user.Save(db)
		if err!=nil{//存入数据库出错
			c.JSON(http.StatusBadRequest,gin.H{
				"method":  "POST",
				"routing": "register",
				"Error":err.Error(),
			})
			return
		}

		//为新注册的用户创建文件夹
		//由于文件夹包含UID，所以放在存入数据库之后,以便获得数据库中的ID
		s:="./files/"+strconv.Itoa(user.Uid)+"/"
		err=os.MkdirAll(s,0777)
		if err!=nil {//创建文件夹出错
			c.JSON(http.StatusBadRequest,gin.H{
				"method":  "POST",
				"routing": "register",
				"Error":err.Error(),
			})
			return
		}

		//反馈
		c.JSON(http.StatusOK,gin.H{
				"method":  "POST",
				"routing": "register",
				"AddUserID":user.Uid,
		})
	}
}

